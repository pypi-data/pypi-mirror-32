from typing import Optional, List, Union, Dict, Callable, AnyStr, IO, \
    DefaultDict
import asyncio
import abc
import itertools
import logging
import time
import random
from collections import defaultdict

import aiohttp
from aiohttp.client import DEFAULT_TIMEOUT
from aiohttp import ClientSession
from yarl import URL
from tenacity import retry
from tenacity.retry import retry_if_result, retry_if_exception_type
from tenacity.wait import wait_fixed
from tenacity.stop import stop_after_attempt

from .pipelines import Pipeline
from .things import Request, Response, Item, Things
from .pool import Pool
from .exceptions import ThingDropped

__all__ = ['Ant']


class Ant(abc.ABC):
    response_pipelines: List[Pipeline] = []
    request_pipelines: List[Pipeline] = []
    item_pipelines: List[Pipeline] = []
    request_cls = Request
    response_cls = Response
    request_timeout = DEFAULT_TIMEOUT.total
    request_retries = 3
    request_retry_delay = 5
    request_proxies: List[Union[str, URL]] = []
    request_max_redirects = 10
    request_allow_redirects = True
    response_in_stream = False
    connection_limit = 100  # see "TCPConnector" in "aiohttp"
    connection_limit_per_host = 0
    pool_limit = 100
    pool_raise_exception = False

    def __init__(self):
        self.logger = logging.getLogger(self.__class__.__name__)
        # report var
        self._reports: DefaultDict[str, List[int, int]] = defaultdict(
            lambda: [0, 0])
        self._drop_reports: DefaultDict[str, List[int, int]] = defaultdict(
            lambda: [0, 0])
        self._start_time = time.time()
        self._last_time = self._start_time
        self._report_slot = 60  # report once after one minute by default
        self.session: aiohttp.ClientSession = ClientSession(
            response_class=self.response_cls, request_class=self.request_cls,
            connector=aiohttp.TCPConnector(
                limit=self.connection_limit,
                enable_cleanup_closed=True,
                limit_per_host=self.connection_limit_per_host)
        )
        self.pool: Pool = Pool(
            limit=self.pool_limit, raise_exception=self.pool_raise_exception)

    @property
    def name(self):
        return self.__class__.__name__

    async def request(self, url: Union[str, URL], method: str = 'GET',
                      params: Optional[dict] = None,
                      headers: Optional[dict] = None,
                      cookies: Optional[dict] = None,
                      data: Optional[Union[AnyStr, Dict, IO]] = None,
                      proxy: Optional[Union[str, URL]] = None,
                      timeout: Optional[Union[int, float]] = None,
                      retries: Optional[int] = None,
                      response_in_stream: Optional[bool] = None
                      ) -> Response:
        if not isinstance(url, URL):
            url = URL(url)
        if proxy and not isinstance(proxy, URL):
            proxy = URL(proxy)
        elif proxy is None:
            proxy = self.get_proxy()
        if timeout is None:
            timeout = self.request_timeout
        if retries is None:
            retries = self.request_retries
        if response_in_stream is None:
            response_in_stream = self.response_in_stream

        req = self.request_cls(method, url, timeout=timeout, params=params,
                               headers=headers, cookies=cookies, data=data,
                               proxy=proxy,
                               response_in_stream=response_in_stream)
        req = await self._handle_thing_with_pipelines(
            req, self.request_pipelines)
        self.report(req)

        if retries > 0:
            res = await self.make_retry_decorator(
                retries, self.request_retry_delay)(self._request)(req)
        else:
            res = await self._request(req)

        res = await self._handle_thing_with_pipelines(
            res, self.response_pipelines)
        self.report(res)
        return res

    async def get(self, url: Union[str, URL], params: Optional[dict] = None,
                  headers: Optional[dict] = None,
                  cookies: Optional[dict] = None,
                  data: Optional[Union[AnyStr, Dict, IO]] = None,
                  proxy: Optional[Union[str, URL]] = None,
                  timeout: Optional[Union[int, float]] = None,
                  retries: Optional[int] = None) -> Response:
        kwargs = locals()
        kwargs.pop('self')
        return await self.request(method='GET', **kwargs)

    async def post(self, url: Union[str, URL], params: Optional[dict] = None,
                   headers: Optional[dict] = None,
                   cookies: Optional[dict] = None,
                   data: Optional[Union[AnyStr, Dict, IO]] = None,
                   proxy: Optional[Union[str, URL]] = None,
                   timeout: Optional[Union[int, float]] = None,
                   retries: Optional[int] = None) -> Response:
        kwargs = locals()
        kwargs.pop('self')
        return await self.request(method='POST', **kwargs)

    async def put(self, url: Union[str, URL], params: Optional[dict] = None,
                  headers: Optional[dict] = None,
                  cookies: Optional[dict] = None,
                  data: Optional[Union[AnyStr, Dict, IO]] = None,
                  proxy: Optional[Union[str, URL]] = None,
                  timeout: Optional[Union[int, float]] = None,
                  retries: Optional[int] = None) -> Response:
        kwargs = locals()
        kwargs.pop('self')
        return await self.request(method='PUT', **kwargs)

    async def patch(self, url: Union[str, URL], params: Optional[dict] = None,
                    headers: Optional[dict] = None,
                    cookies: Optional[dict] = None,
                    data: Optional[Union[AnyStr, Dict, IO]] = None,
                    proxy: Optional[Union[str, URL]] = None,
                    timeout: Optional[Union[int, float]] = None,
                    retries: Optional[int] = None) -> Response:
        kwargs = locals()
        kwargs.pop('self')
        return await self.request(method='PATCH', **kwargs)

    async def delete(self, url: Union[str, URL], params: Optional[dict] = None,
                     headers: Optional[dict] = None,
                     cookies: Optional[dict] = None,
                     data: Optional[Union[AnyStr, Dict, IO]] = None,
                     proxy: Optional[Union[str, URL]] = None,
                     timeout: Optional[Union[int, float]] = None,
                     retries: Optional[int] = None) -> Response:
        kwargs = locals()
        kwargs.pop('self')
        return await self.request(method='DELETE', **kwargs)

    async def head(self, url: Union[str, URL], params: Optional[dict] = None,
                   headers: Optional[dict] = None,
                   cookies: Optional[dict] = None,
                   data: Optional[Union[AnyStr, Dict, IO]] = None,
                   proxy: Optional[Union[str, URL]] = None,
                   timeout: Optional[Union[int, float]] = None,
                   retries: Optional[int] = None) -> Response:
        kwargs = locals()
        kwargs.pop('self')
        return await self.request(method='HEAD', **kwargs)

    async def collect(self, item: Item) -> None:
        self.logger.debug('Collect item: ' + str(item))
        await self._handle_thing_with_pipelines(item, self.item_pipelines)
        self.report(item)

    async def open(self) -> None:
        self.logger.info('Opening')
        for pipeline in itertools.chain(self.item_pipelines,
                                        self.response_pipelines,
                                        self.request_pipelines):
            obj = pipeline.on_spider_open()
            if asyncio.iscoroutine(obj):
                await obj

    async def close(self) -> None:
        for pipeline in itertools.chain(self.item_pipelines,
                                        self.response_pipelines,
                                        self.request_pipelines):
            obj = pipeline.on_spider_close()
            if asyncio.iscoroutine(obj):
                await obj

        await self.session.close()
        await self.pool.close()

        self.logger.info('Closed')

    @abc.abstractmethod
    async def run(self) -> None:
        """App custom entrance"""

    async def main(self) -> None:
        try:
            await self.open()
            await self.run()
            # wait scheduled coroutines before "self.close" coroutine running
            await self.pool.wait_scheduled_coroutines()
            await self.close()
        except Exception as e:
            self.logger.exception(
                'Run ant with ' + e.__class__.__name__)
        await self.pool.wait_scheduled_coroutines()
        # total report
        for name, counts in self._reports.items():
            self.logger.info('Get {:d} {:s} in total'.format(counts[1], name))
        for name, counts in self._drop_reports.items():
            self.logger.info('Drop {:d} {:s} in total'.format(counts[1], name))
        self.logger.info(
            'Run {:s} in {:f} seconds'.format(self.__class__.__name__,
                                              time.time() - self._start_time))

    @staticmethod
    def make_retry_decorator(retries: int, delay: float
                             ) -> Callable[[Callable], Callable]:
        return retry(wait=wait_fixed(delay),
                     retry=(retry_if_result(lambda res: res.status >= 500) |
                            retry_if_exception_type()),
                     stop=stop_after_attempt(retries + 1))

    def get_proxy(self) -> Optional[URL]:
        """Chose a proxy, default by random"""
        try:
            return URL(random.choice(self.request_proxies))
        except IndexError:
            return None

    async def _handle_thing_with_pipelines(
            self, thing: Things, pipelines: List[Pipeline]) -> Things:
        """Process thing one by one, break the process chain when get
        exception.

        :raise ThingDropped"""
        self.logger.debug('Process thing: ' + str(thing))
        raw_thing = thing
        for pipeline in pipelines:
            try:
                thing = pipeline.process(thing)
                if asyncio.iscoroutine(thing):
                    thing = await thing
            except Exception as e:
                if isinstance(e, ThingDropped):
                    self.report(raw_thing, dropped=True)
                raise e
        return thing

    async def _request(self, req: Request) -> Response:
        proxy = req.proxy
        # cookies in headers, params in url
        req_kwargs = dict(method=req.method, url=req.url, headers=req.headers,
                          data=req.data, timeout=req.timeout, proxy=proxy,
                          max_redirects=self.request_max_redirects,
                          allow_redirects=self.request_allow_redirects)
        # proxy auth not work in one session with many requests,
        # add auth header to fix it
        if proxy is not None:
            if proxy.scheme == 'http' and proxy.user is not None:
                req_kwargs['headers'][aiohttp.hdrs.PROXY_AUTHORIZATION] = \
                    aiohttp.BasicAuth.from_url(proxy).encode()

        response = await self.session._request(**req_kwargs)

        if not req.response_in_stream:
            await response.read()
            response.close()
            await response.wait_for_close()
        return response

    def report(self, thing: Things, dropped: bool = False) -> None:
        now_time = time.time()
        if now_time - self._last_time > self._report_slot:
            self._last_time = now_time
            for name, counts in self._reports.items():
                count = counts[1] - counts[0]
                counts[0] = counts[1]
                self.logger.info(
                    'Get {:d} {:s} in total with {:d}/{:d}s rate'.format(
                        counts[1], name, count, self._report_slot))
            for name, counts in self._drop_reports.items():
                count = counts[1] - counts[0]
                counts[0] = counts[1]
                self.logger.info(
                    'Drop {:d} {:s} in total with {:d}/{:d} rate'.format(
                        counts[1], name, count, self._report_slot))
        report_type = thing.__class__.__name__
        if dropped:
            reports = self._drop_reports
        else:
            reports = self._reports
        counts = reports[report_type]
        counts[1] += 1
