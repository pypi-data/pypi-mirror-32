import logging
import asyncio
import traceback
from functools import partial
import aioamqp.protocol
import aioamqp.channel
import aioamqp.exceptions
import aioamqp.envelope
import aioamqp.properties as amqp_prop
import aioamqp
from typing import Optional, List, Callable
from .app import Component
from .error import PrepareError
from .misc import mask_url_pwd, async_call
from .tracer import (Span, CLIENT, SERVER, SPAN_TYPE, SPAN_KIND,
                     SPAN_TYPE_AMQP, SPAN_KIND_AMQP_IN, SPAN_KIND_AMQP_OUT,
                     SPAN_KIND_AMQP_ACK, SPAN_KIND_AMQP_NACK)

#
aioamqp.channel.logger.level = logging.CRITICAL
aioamqp.protocol.logger.level = logging.CRITICAL

STOP_TIMEOUT = 5


class AmqpTracerConfig:
    def on_publish_start(self, context_span: 'Span',
                         channel: 'aioamqp.channel.Channel', payload: bytes,
                         exchange_name: str, routing_key: str,
                         properties: Optional[dict], mandatory: bool,
                         immediate: bool) -> None:
        pass

    def on_publish_end(self, context_span: 'Span',
                       channel: 'aioamqp.channel.Channel',
                       err: Optional[Exception]) -> None:
        if err:
            context_span.tag('error.message', str(err))
            context_span.annotate(traceback.format_exc())

    def on_ack_start(self, span: 'Span', channel: 'aioamqp.channel.Channel',
                     delivery_tag: str, multiple: bool) -> None:
        pass

    def on_ack_end(self, span: 'Span', channel: 'aioamqp.channel.Channel',
                   err: Optional[Exception]) -> None:
        if err:
            span.tag('error.message', str(err))
            span.annotate(traceback.format_exc())

    def on_nack_start(self, span: 'Span', channel: 'aioamqp.channel.Channel',
                      delivery_tag: str, multiple: bool) -> None:
        pass

    def on_nack_end(self, span: 'Span', channel: 'aioamqp.channel.Channel',
                    err: Optional[Exception]) -> None:
        if err:
            span.tag('error.message', str(err))
            span.annotate(traceback.format_exc())


class Channel:
    name: Optional[str] = None
    amqp: Optional['Amqp'] = None
    channel: 'aioamqp.channel.Channel' = None
    _cons_cnt: int = 0
    _cons_fut: asyncio.Future = None
    _cons_tag: str = None
    _stopping: bool = False

    async def open(self):
        try:
            await self.close()
        except Exception as err:
            if self.amqp:
                self.amqp.app.log_err(err)
        self.channel = await self.amqp._protocol.channel()

    async def close(self):
        if self.channel:
            await self.channel.close()

    async def publish(self, context_span: Span, payload: bytes,
                      exchange_name: str, routing_key: str,
                      properties: Optional[dict] = None,
                      mandatory: bool = False, immediate: bool = False,
                      tracer_config: Optional[AmqpTracerConfig] = None,
                      propagate_trace: bool = True, retry: bool=True):
        span = None
        if context_span:
            span = context_span.new_child(
                'amqp:publish {} {}'.format(exchange_name, routing_key),
                CLIENT
            )
            span.metrics_tag(SPAN_TYPE, SPAN_TYPE_AMQP)
            span.metrics_tag(SPAN_KIND, SPAN_KIND_AMQP_OUT)
            if propagate_trace:
                headers = context_span.make_headers()
                properties = properties or {}
                if 'headers' not in properties:
                    properties['headers'] = {}
                properties['headers'].update(headers)
            span.start()
            if tracer_config:
                tracer_config.on_publish_start(span, self.channel, payload,
                                               exchange_name, routing_key,
                                               properties, mandatory,
                                               immediate)
        try:
            try:
                await self.channel.basic_publish(payload, exchange_name,
                                                 routing_key,
                                                 properties=properties,
                                                 mandatory=mandatory,
                                                 immediate=immediate)
            except Exception as e:
                span.tag('error', 'true')
                span.tag('error.message', str(e))
                span.annotate(traceback.format_exc())
                await self.open()
                if retry:
                    await self.publish(context_span, payload, exchange_name,
                                       routing_key, properties, mandatory,
                                       immediate, tracer_config,
                                       propagate_trace, retry=False)

            if span:
                if tracer_config:
                    tracer_config.on_publish_end(span, self.channel, None)
                span.finish()
        except Exception as err:
            if span:
                if tracer_config:
                    tracer_config.on_publish_end(span, self.channel, err)
                span.finish(exception=err)
            raise

    async def consume(self, fn, queue_name='', consumer_tag='', no_local=False,
                      no_ack=False, exclusive=False, no_wait=False,
                      arguments=None):
        if not asyncio.iscoroutinefunction(fn):
            raise UserWarning()

        callback = partial(self._consume_callback, fn)
        self._cons_fut = asyncio.Future(loop=self.amqp.loop)
        res = await self.channel.basic_consume(
            callback, queue_name=queue_name, consumer_tag=consumer_tag,
            no_local=no_local, no_ack=no_ack, exclusive=exclusive,
            no_wait=no_wait, arguments=arguments)
        self._cons_tag = res['consumer_tag']
        return res

    async def ack(self, context_span: Span, delivery_tag: str,
                  multiple: bool = False,
                  tracer_config: Optional[AmqpTracerConfig] = None) -> None:
        await self._ack_nack(context_span, True, delivery_tag, multiple,
                             tracer_config)

    async def nack(self, context_span: Span, delivery_tag: str,
                   multiple: bool = False,
                   tracer_config: Optional[AmqpTracerConfig] = None) -> None:
        await self._ack_nack(context_span, False, delivery_tag, multiple,
                             tracer_config)

    async def _ack_nack(self, context_span: Span, is_ack: bool,
                        delivery_tag: str, multiple: bool = False,
                        tracer_config: Optional[AmqpTracerConfig] = None):
        span = None
        if context_span:
            tracer_config = tracer_config or AmqpTracerConfig()
            span = context_span.new_child('amqp:ack', CLIENT)
            span.metrics_tag(SPAN_TYPE, SPAN_TYPE_AMQP)
            span.start()
        try:
            if is_ack:
                if span:
                    span.metrics_tag(SPAN_KIND, SPAN_KIND_AMQP_ACK)
                    tracer_config.on_ack_start(span, self.channel,
                                               delivery_tag, multiple)
                await self.channel.basic_client_ack(delivery_tag=delivery_tag,
                                                    multiple=multiple)
                if span:
                    tracer_config.on_ack_end(span, self.channel, None)
            else:
                span.metrics_tag(SPAN_KIND, SPAN_KIND_AMQP_NACK)
                if span:
                    tracer_config.on_nack_start(span, self.channel,
                                                delivery_tag, multiple)
                await self.channel.basic_client_nack(delivery_tag=delivery_tag,
                                                     multiple=multiple)
                if span:
                    tracer_config.on_nack_end(span, self.channel, None)
            if span:
                span.finish()
        except Exception as err:
            if span:
                if is_ack:
                    tracer_config.on_ack_end(span, self.channel, err)
                else:
                    tracer_config.on_nack_end(span, self.channel, err)
                span.finish(exception=err)
            raise

    async def _consume_callback_handler(self, callback: Callable,
                                        channel: aioamqp.channel.Channel,
                                        body: bytes,
                                        envelope: aioamqp.envelope.Envelope,
                                        properties: amqp_prop.Properties):
        async_call(self.amqp.loop,
                   partial(
                       self._consume_callback, callback, channel, body,
                       envelope, properties))

    async def _consume_callback(self, callback: Callable,
                                channel: aioamqp.channel.Channel, body: bytes,
                                envelope: aioamqp.envelope.Envelope,
                                properties: amqp_prop.Properties):
        if not channel.is_open:
            return

        self._cons_cnt += 1
        span = None
        try:
            if self.amqp.app.tracer:
                span = self.amqp.app.tracer.new_trace_from_headers(
                    properties.headers)
                span.name('amqp:message')
                span.metrics_tag(SPAN_TYPE, SPAN_TYPE_AMQP)
                span.metrics_tag(SPAN_KIND, SPAN_KIND_AMQP_IN)
                span.kind(SERVER)
                if envelope.routing_key is not None:
                    span.tag('amqp.routing_key', envelope.routing_key)
                if envelope.exchange_name is not None:
                    span.tag('amqp.exchange_name', envelope.exchange_name)
                if properties.delivery_mode is not None:
                    span.tag('amqp.delivery_mode',
                             properties.delivery_mode)
                if properties.expiration is not None:
                    span.tag('amqp.expiration', properties.expiration)
                span.start()

            await callback(span, channel, body, envelope, properties)

            if span:
                span.finish()
        except Exception as err:
            if span:
                span.tag('error.message', str(err))
                span.annotate(traceback.format_exc())
                span.finish(exception=err)
            self.amqp.app.log_err(err)
            raise

        finally:
            self._cons_cnt -= 1
            if self._stopping and self._cons_cnt == 0 and self._cons_fut:
                self._cons_fut.set_result(1)

    async def start(self):
        self._stopping = False
        await self.open()

    async def stop(self):
        self._stopping = True
        if self._cons_tag:
            try:
                await self.channel.basic_cancel(self._cons_tag)
            except Exception as err:
                self.amqp.app.log_err(err)
            finally:
                self._cons_tag = None
        try:
            if self._cons_cnt > 0 and self._cons_fut:
                await asyncio.wait_for(self._cons_fut, timeout=STOP_TIMEOUT)
        finally:
            await self.close()

    async def _safe_declare_queue(self, queue_name=None, passive=False,
                                  durable=False, exclusive=False,
                                  auto_delete=False, no_wait=False,
                                  arguments=None) -> Optional[dict]:
        ch = await self.amqp._protocol.channel()
        try:
            res = await ch.queue_declare(queue_name=queue_name,
                                         passive=passive, durable=durable,
                                         exclusive=exclusive,
                                         auto_delete=auto_delete,
                                         no_wait=no_wait, arguments=arguments)
            return res
        except aioamqp.exceptions.ChannelClosed as e:
            if e.code == 406:
                # ignore error if attributes not match
                return None
            else:
                raise
        finally:
            if ch.is_open:
                await ch.close()

    async def _safe_declare_exchange(self, exchange_name, type_name,
                                     passive=False, durable=False,
                                     auto_delete=False, no_wait=False,
                                     arguments=None) -> Optional[dict]:
        ch = await self.amqp._protocol.channel()
        try:
            res = await ch.exchange_declare(exchange_name=exchange_name,
                                            type_name=type_name,
                                            passive=passive, durable=durable,
                                            auto_delete=auto_delete,
                                            no_wait=no_wait,
                                            arguments=arguments)
            return res
        except aioamqp.exceptions.ChannelClosed as e:
            if e.code == 406:
                # ignore error if attributes not match
                return None
            else:
                raise
        finally:
            if ch.is_open:
                await ch.close()


class Amqp(Component):

    def __init__(self, url: Optional[str] = None,
                 channels: List[Channel] = None,
                 heartbeat: int = 5,
                 connect_max_attempts: int = 10,
                 connect_retry_delay: float = 1.0) -> None:
        super().__init__()
        self.url = url
        self.connect_max_attempts = connect_max_attempts
        self.connect_retry_delay = connect_retry_delay
        self.heartbeat = heartbeat
        self._started = False
        self._shutting_down = False
        self._consuming = False
        self._transport = None
        self._protocol: aioamqp.protocol.AmqpProtocol = None
        self._channels = channels
        if channels:
            names = [ch.name for ch in channels if ch.name is not None]
            if len(names) != len(set(names)):
                raise UserWarning('There are not unique names in the channel '
                                  'names: %s' % (','.join(names)))

    @property
    def _masked_url(self) -> Optional[str]:
        return mask_url_pwd(self.url)

    async def prepare(self) -> None:
        self._connecting = True
        for i in range(self.connect_max_attempts):
            try:
                await self._connect()
                return
            except Exception as e:
                self.app.log_err(str(e))
                await asyncio.sleep(self.connect_retry_delay,
                                    loop=self.app.loop)
        raise PrepareError("Could not connect to %s" % self._masked_url)

    async def start(self) -> None:
        self._started = True
        await self._start_channels()

    async def stop(self) -> None:
        self._started = False
        self._shutting_down = True
        await self._stop_channels()
        await self._cleanup()

    async def _connect(self):
        await self._cleanup()
        self.app.log_info("Connecting to %s" % self._masked_url)
        (self._transport,
         self._protocol) = await aioamqp.from_url(self.url,
                                                  on_error=self._con_error,
                                                  heartbeat=self.heartbeat)
        self.app.log_info("Connected to %s" % self._masked_url)
        self._connecting = False

        if self._started:
            await self._start_channels()

    async def _con_error(self, error):
        if error and not self._shutting_down:
            self.app.log_err(error)
        if self._shutting_down or not self._started:
            return
        if self._connecting:
            return
        self._connecting = True

        async_call(self.loop, self._reconnect,
                   delay=self.connect_retry_delay)

    async def _reconnect(self):
        try:
            await self._connect()
        except Exception as e:
            self.app.log_err(e)
            async_call(self.loop, self._reconnect,
                       delay=self.connect_retry_delay)

    async def _cleanup(self):
        if self._protocol:
            try:
                await self._protocol.close()
            except Exception as e:
                self.app.log_err(e)
            self._protocol = None
            self._transport = None

    async def _start_channels(self):
        self._consuming = True
        if self._channels:
            for ch in self._channels:
                ch.amqp = self
                await ch.start()

    async def _stop_channels(self):
        self._consuming = False
        if self._channels:
            for ch in reversed(self._channels):
                await ch.stop()

    def channel(self, name: str) -> Optional['Channel']:
        if self._channels:
            for ch in self._channels:
                if ch.name is not None and ch.name == name:
                    return ch
        return None
