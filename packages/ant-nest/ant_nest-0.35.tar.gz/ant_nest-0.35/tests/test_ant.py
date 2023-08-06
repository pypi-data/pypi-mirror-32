import asyncio
import os

import pytest
from tenacity import RetryError
import asyncio
from yarl import URL

from ant_nest import *
from .test_things import fake_response


@pytest.mark.asyncio
async def test_ant():
    class TestPipeline(Pipeline):
        def __init__(self):
            super().__init__()
            self.count = 0

        def process(self, thing):
            self.count += 1
            return thing

    class TestAnt(Ant):
        item_pipelines = [TestPipeline()]
        request_retries = 0

        async def run(self):
            await self.request('http://test.com')
            await self.collect(object())
            assert self.item_pipelines[0].count == 1

        async def _request(self, req: Request):
            return fake_response(b'')

    ant = TestAnt()
    assert ant.name == 'TestAnt'
    await ant.main()

    class Test2Ant(TestAnt):
        async def run(self):
            raise Exception('Test exception')

    await Test2Ant().main()


@pytest.mark.asyncio
async def test_ant_with_retry():
    class Test2Ant(Ant):
        request_retries = 2
        request_retry_delay = 0.1

        def __init__(self):
            super().__init__()
            self.min_retries = 2
            self.retries = 0

        async def run(self):
            return None

        async def _request(self, req: Request):
            if self.retries < self.min_retries:
                self.retries += 1
                raise IOError()
            else:
                res = fake_response(b'')
                res.status = 200
                return res

    await Test2Ant().request('https://www.test.com')
    ant = Test2Ant()
    ant.request_retries = 1
    with pytest.raises(RetryError):
        await ant.request('https://www.test.com')
    # with params
    ant.retries = 0
    ant.request_retries = 3
    await ant.request('https://www.test.com')
    ant.retries = 0
    with pytest.raises(IOError):
        await ant.request('https://www.test.com', retries=0)

    await ant.main()
    # retry condition

    class Test3Ant(Test2Ant):
        async def _request(self, req: Request):
            if self.retries < self.min_retries:
                self.retries += 1
                res = fake_response(b'')
                res.status = 500
                return res
            else:
                res = fake_response(b'')
                res.status = 200
                return res

    await Test3Ant().request('https://www.test.com')
    ant = Test3Ant()
    ant.request_retries = 1
    with pytest.raises(RetryError):
        await ant.request('https://www.test.com')

    await ant.main()


@pytest.mark.asyncio
async def test_pipelines():

    class TestPipeline(Pipeline):
        async def process(self, thing):
            raise TypeError('Test error')

        def on_spider_open(self):
            raise TypeError('Test error')

        def on_spider_close(self):
            raise TypeError('Test error')

    class TestAnt(Ant):
        item_pipelines = [TestPipeline()]

        async def run(self):
            return None

    pls = [Pipeline() for x in range(10)]
    thing = object()
    ant = TestAnt()
    assert thing is await ant._handle_thing_with_pipelines(thing, pls)
    # with exception
    pls[5] = TestPipeline()
    with pytest.raises(TypeError):
        await ant._handle_thing_with_pipelines(thing, pls)
    # exception will be ignored in "open" and "close" method
    with pytest.raises(TypeError):
        await ant.open()
    with pytest.raises(TypeError):
        await ant.close()


@pytest.mark.asyncio
async def test_ant_report():

    class FakePipeline(Pipeline):
        def process(self, thing: Request):
            if thing.host == 'error':
                raise ThingDropped
            return thing

    class FakeAnt(Ant):
        request_pipelines = [FakePipeline()]
        request_retries = 0

        async def run(self):
            pass

        async def _request(self, req: Request):
            return fake_response(b'')

    ant = FakeAnt()
    # request report
    await ant.request('http://test.com')
    assert ant._reports['Request'][1] == 1
    assert ant._reports['Request'][0] == 0
    with pytest.raises(ThingDropped):
        await ant.request('http://error')
    assert ant._drop_reports['Request'][1] == 1

    ant._last_time -= ant._last_time + 1  # report
    ant.report(Request('GET', URL('http://test')))
    assert ant._reports['Request'][0] == 1
    ant.report(Request('GET', URL('http://test')), dropped=True)
    assert ant._drop_reports['Request'][0] == 1

    await ant.main()


@pytest.mark.asyncio
async def test_with_real_request():
    httpbin_base_url = os.getenv('TEST_HTTPBIN', 'http://localhost:8080/')

    class TestAnt(Ant):
        async def run(self):
            return None

    ant = TestAnt()
    res = await ant.request(httpbin_base_url)
    assert res.status == 200
    # method
    for method in ['GET', 'POST', 'DELETE', 'PUT', 'PATCH', 'HEAD']:
        res = await ant.request(httpbin_base_url + 'anything', method=method)
        assert res.status == 200
        if method != 'HEAD':
            assert res.simple_json['method'] == method
        else:
            assert res.simple_text == ''
        # short way
        res = await getattr(ant, method.lower())(httpbin_base_url + 'anything')
        assert res.status == 200
        if method != 'HEAD':
            assert res.simple_json['method'] == method
        else:
            assert res.simple_text == ''
    # params
    res = await ant.request(httpbin_base_url + 'get?k1=v1&k2=v2')
    assert res.status == 200
    assert res.simple_json['args']['k1'] == 'v1'
    assert res.simple_json['args']['k2'] == 'v2'
    res = await ant.get(httpbin_base_url + 'get', params={'k1': 'v1', 'k2': 'v2'})
    assert res.status == 200
    assert res.simple_json['args']['k1'] == 'v1'
    assert res.simple_json['args']['k2'] == 'v2'
    # data with str
    res = await ant.post(httpbin_base_url + 'post', data='Test data')
    assert res.status == 200
    assert res.simple_json['data'] == 'Test data'
    # data with dict
    res = await ant.request(httpbin_base_url + 'post', method='POST', data={'k1': 'v1'})
    assert res.status == 200
    assert res.simple_json['form']['k1'] == 'v1'
    # data with bytes
    res = await ant.request(httpbin_base_url + 'post', method='POST', data=b'12345')
    assert res.status == 200
    assert res.simple_json['data'] == '12345'
    # data with file
    with open('tests/test.html', 'r') as f:
        file_content = f.read()
        f.seek(0)
        res = await ant.request(httpbin_base_url + 'post', method='POST', data=f)
        assert res.status == 200
        assert res.simple_json['data'] == file_content
    # headers
    res = await ant.request(httpbin_base_url + 'headers', headers={'Custom-Key': 'test'})
    assert res.status == 200
    assert res.headers['Content-Type'] == 'application/json'
    assert res.simple_json['headers']['Custom-Key'] == 'test'
    # cookies
    res = await ant.request(httpbin_base_url + 'cookies', cookies={'Custom-Key': 'test'})
    assert res.status == 200
    assert res.simple_json['cookies']['Custom-Key'] == 'test'
    # get cookies
    ant.request_allow_redirects = False
    res = await ant.request(httpbin_base_url + 'cookies/set?k1=v1&k2=v2')
    assert res.status == 302
    assert res.cookies['k1'].value == 'v1'
    assert res.cookies['k2'].value == 'v2'
    # redirects
    ant.request_allow_redirects = True
    ant.request_max_redirects = 3
    res = await ant.request(httpbin_base_url + 'redirect/2')
    assert res.status == 200
    # with http proxy
    proxy = os.getenv('TEST_HTTP_PROXY', 'http://bugman:letmein@localhost:3128')
    ant.request_proxies.append(proxy)
    res = await ant.request('http://httpbin.org/anything')
    assert res.status == 200
    # no proxy anymore
    ant.request_proxies.pop()
    res = await ant.request('http://httpbin.org/anything')
    assert res.status == 200
    # set proxy by request
    res = await ant.get('http://httpbin.org/anything', proxy=proxy)
    assert res.status == 200
    # with stream
    ant.response_in_stream = True
    res = await ant.request(httpbin_base_url + 'anything')
    assert res.status == 200
    with pytest.raises(ValueError):
        res.simple_text
    while True:
        chunk = await res.content.read(10)
        if len(chunk) == 0:
            break
    # set streaming by request
    res = await ant.request(httpbin_base_url + 'anything',
                            response_in_stream=False)
    assert res.status == 200
    assert res.simple_text is not None

    await ant.close()
