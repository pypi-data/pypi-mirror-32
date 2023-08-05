import asyncio
import signal
import logging
from typing import Dict, Optional
from .error import PrepareError, GracefulExit
from .tracer import Tracer

logger = logging.getLogger('aioapp')


def _raise_graceful_exit():  # pragma: no cover
    raise GracefulExit()


class Component(object):
    def __init__(self) -> None:
        super(Component, self).__init__()
        self.loop: asyncio.AbstractEventLoop = None
        self.app: 'Application' = None

    async def prepare(self) -> None:
        raise NotImplementedError()

    async def start(self) -> None:
        raise NotImplementedError()

    async def stop(self) -> None:
        raise NotImplementedError()


class Application(object):
    def __init__(self, loop=None) -> None:
        super(Application, self).__init__()
        self.loop = loop or asyncio.get_event_loop()
        self._components: Dict[str, Component] = {}
        self._stop_deps: dict = {}
        self._stopped: list = []
        self.tracer: Tracer = Tracer(self, self.loop)

    def add(self, name: str, comp: Component,
            stop_after: list = None):
        if not isinstance(comp, Component):
            raise UserWarning()
        if name in self._components:
            raise UserWarning()
        if stop_after:
            for cmp in stop_after:
                if cmp not in self._components:
                    raise UserWarning('Unknown component %s' % cmp)
        comp.loop = self.loop
        comp.app = self
        self._components[name] = comp
        self._stop_deps[name] = stop_after

    def __getattr__(self, item: str) -> Component:
        if item not in self._components:
            raise AttributeError
        return self._components[item]

    def log_err(self, err):
        if not err:
            return
        if isinstance(err, BaseException):
            logging.exception(err)
        else:
            logging.error(err)

    def log_warn(self, warn):
        logging.warning(warn)

    def log_info(self, info):
        logging.info(info)

    def log_debug(self, debug):
        logging.debug(debug)

    def setup_logging(self, tracer_driver=None, tracer_addr=None,
                      tracer_name=None, tracer_sample_rate=1.0,
                      tracer_send_inteval=3,
                      tracer_default_sampled: Optional[bool] = None,
                      tracer_default_debug: Optional[bool] = None,
                      metrics_driver=None, metrics_addr=None,
                      metrics_name=None):
        if tracer_driver:
            self.tracer.setup_tracer(tracer_driver, tracer_name, tracer_addr,
                                     tracer_sample_rate, tracer_send_inteval,
                                     tracer_default_sampled,
                                     tracer_default_debug)
        if metrics_driver:
            self.tracer.setup_metrics(metrics_driver, metrics_addr,
                                      metrics_name)

    async def _shutdown_tracer(self):
        if self.tracer:
            self.log_info("Shutting down tracer")
            await self.tracer.close()

    def run(self) -> int:
        try:
            self.loop.run_until_complete(self.run_prepare())
        except PrepareError as e:
            self.log_err(e)
            return 1
        except KeyboardInterrupt:  # pragma: no cover
            return 1
        self.run_loop()
        self.loop.run_until_complete(self.run_shutdown())
        print("Bye")
        if hasattr(self.loop, 'shutdown_asyncgens'):
            self.loop.run_until_complete(self.loop.shutdown_asyncgens())
        self.loop.close()
        return 0

    async def run_prepare(self):
        self.log_info('Prepare for start')

        await asyncio.gather(*[comp.prepare()
                               for comp in self._components.values()],
                             loop=self.loop)

        self.log_info('Starting...')
        await asyncio.gather(*[comp.start()
                               for comp in self._components.values()],
                             loop=self.loop)

        self.log_info('Running...')

    def run_loop(self):
        try:
            self.loop.add_signal_handler(signal.SIGINT, _raise_graceful_exit)
            self.loop.add_signal_handler(signal.SIGTERM, _raise_graceful_exit)
        except NotImplementedError:  # pragma: no cover
            # add_signal_handler is not implemented on Windows
            pass
        try:
            self.loop.run_forever()
        except GracefulExit:  # pragma: no cover
            pass

    async def run_shutdown(self):
        self.log_info('Shutting down...')
        for comp_name in self._components:
            await self._stop_comp(comp_name)
        await self._shutdown_tracer()

    async def _stop_comp(self, name):
        if name in self._stopped:
            return
        if name in self._stop_deps and self._stop_deps[name]:
            for dep_name in self._stop_deps[name]:
                await self._stop_comp(dep_name)
        await self._components[name].stop()
        self._stopped.append(name)
