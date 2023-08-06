import os
import time
from datetime import datetime
import io
from unittest import mock

import pytest
from yarl import URL
import aiofiles

from ant_nest import *
from .test_things import fake_response


@pytest.mark.asyncio
async def test_pipeline():
    pl = Pipeline()
    await pl.process(Request('GET', URL('https://test.com')))


def test_response_filter_error_pipeline():
    pl = ResponseFilterErrorPipeline()
    res = fake_response(b'')
    err_res = fake_response(b'')
    res.status = 200
    err_res.status = 403
    assert res is pl.process(res)
    with pytest.raises(ThingDropped):
        pl.process(err_res)


def test_request_duplicate_filter_pipeline():
    pl = RequestDuplicateFilterPipeline()
    req = Request('GET', URL('http://test.com'))
    assert pl.process(req) is req
    with pytest.raises(ThingDropped):
        pl.process(req)


def test_item_print_pipeline(item_cls):
    pl = ItemPrintPipeline()
    item = item_cls()
    item.count = 3
    item.info = 'hi'
    assert pl.process(item) is item


def test_item_filed_replace_pipeline(item_cls):
    pl = ItemFieldReplacePipeline(['info'])
    item = item_cls()
    item.info = 'hi\n,\t\r ant\n'
    pl.process(item)
    assert item.info == 'hi, ant'


@pytest.mark.asyncio
async def test_item_base_file_dump_pipeline():
    pl = ItemBaseFileDumpPipeline()
    await pl.dump('/dev/null', 'Hello World')
    await pl.dump('/dev/null', b'Hello World')
    await pl.dump('/dev/null', io.StringIO('Hello World'))
    await pl.dump('/dev/null', io.BytesIO(b'Hello World'))
    await pl.dump('/dev/null', open('./tests/test.html'), buffer_size=4)
    async with aiofiles.open('./tests/test.html') as f:
        await pl.dump('/dev/null', f)
    async with aiofiles.open('./tests/test.html', 'rb') as f:
        await pl.dump('/dev/null', f, buffer_size=4)

    with pytest.raises(ValueError):
        await pl.dump('/dev/null', None)


@pytest.mark.asyncio
async def test_item_json_dump_pipeline(item_cls):
    pl = ItemJsonDumpPipeline(to_dict=lambda x: x)
    item = item_cls()
    item.count = 1
    assert pl.process(item) is item
    item = item_cls()
    item.info = 'hi'
    pl.process(item)
    await pl.on_spider_close()

    # clean file
    ci = os.getenv('TEST_HOST', 'localhost')
    if ci == 'localhost':
        os.remove('./Item.json')


def test_request_user_agent_pipeline():
    pl = RequestUserAgentPipeline(user_agent='ant')
    req = Request('GET', URL('https://www.hi.com'))
    assert pl.process(req) is req
    assert req.headers['User-Agent'] == 'ant'

    req.headers['User-Agent'] = 'custom'
    assert pl.process(req).headers['User-Agent'] == 'custom'


def test_request_random_user_agent_pipeline():
    pl = RequestRandomUserAgentPipeline()
    req = Request('GET', URL('https://www.hi.com'))
    assert pl.process(req) is req
    assert req.headers.get('User-Agent') is not None

    req.headers['User-Agent'] = 'custom'
    assert pl.process(req).headers['User-Agent'] == 'custom'

    with pytest.raises(ValueError):
        RequestRandomUserAgentPipeline(system='something')

    with pytest.raises(ValueError):
        RequestRandomUserAgentPipeline(browser='something')

    pl = RequestRandomUserAgentPipeline(system='UnixLike', browser='Firefox')
    user_agent = pl.create()
    assert 'X11' in user_agent
    assert 'Firefox' in user_agent


@pytest.mark.asyncio
async def test_item_email_pipeline(item_cls):
    class FakeSMTP:
        async def send_message(self, msg):
            pass

        async def connect(self, *args, **kwargs):
            pass

        async def starttls(self):
            pass

        async def login(self, *args, **kwargs):
            pass

        def close(self):
            pass

    pl = ItemEmailPipeline('test', 'a@b.c', 'letmein', 'localhost', 25,
                           recipients=['b@a.c', 'c@b.a'])

    with mock.patch('aiosmtplib.SMTP', new=FakeSMTP):
        fsmtp = await pl.create_smtp()
        pl.starttls = True
        fsmtp = await pl.create_smtp()

        await pl.send(fsmtp, 'test', 'test')
        with open('./tests/test.html') as f:
            await pl.send(fsmtp, 'test', 'test', attachments=[f])

        # with item
        item = item_cls()
        item.info = 'hi'
        pl.process(item)
        await pl.on_spider_close()


@pytest.mark.asyncio
async def test_item_mysql_pipeline():
    mysql_server = os.getenv('TEST_MYSQL_SERVER', 'localhost')
    mysql_port = int(os.getenv('TEST_MYSQL_PORT', 3306))
    mysql_user = os.getenv('TEST_MYSQL_USER', 'root')
    mysql_password = os.getenv('TEST_MYSQL_PASSWORD', 'letmein')

    bpl = ItemBaseMysqlPipeline(host=mysql_server, port=mysql_port,
                                user=mysql_user, password=mysql_password,
                                database='mysql', table='',
                                to_dict=lambda x: x,
                                buffer_length=1)
    await bpl.on_spider_open()
    await bpl.push_data('''DROP DATABASE IF EXISTS test;
                           CREATE DATABASE test;''')
    await bpl.push_data(
        '''CREATE TABLE test.test (
        `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
        `test` TEXT DEFAULT NULL,
        `test_bool` BOOL DEFAULT NULL,
        `test_int` INT DEFAULT NULL,
        `test_float` FLOAT DEFAULT NULL,
        `test_bytes` BLOB DEFAULT NULL,
        `test_datetime` DATETIME DEFAULT NULL,
        PRIMARY KEY (`id`)
        ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;''')

    test_item = dict(test='I ant', test_bool=False, test_int=1, test_float=0.3,
                     test_bytes=b'\xf0\x9f\x91\x8d',
                     test_datetime=datetime.now())
    ibpl = ItemMysqlInsertPipeline(host=mysql_server, port=mysql_port,
                                   user=mysql_user, password=mysql_password,
                                   database='test', table='test',
                                   to_dict=lambda x: x,
                                   buffer_length=1)
    await ibpl.on_spider_open()
    assert test_item is await ibpl.process(test_item)
    data = await ibpl.pull_data('SELECT * FROM test')
    assert test_item['test'] == data[0]['test']

    ubpl = ItemMysqlUpdatePipeline(host=mysql_server, port=mysql_port,
                                   user=mysql_user, password=mysql_password,
                                   database='test', table='test',
                                   to_dict=lambda x: x,
                                   primary_key='id', buffer_length=1)
    await ubpl.on_spider_open()
    test_item['id'] = data[0]['id']
    test_item['test'] = 'I ANT'
    assert test_item is await ubpl.process(test_item)
    data = await ubpl.pull_data('SELECT * FROM test')
    assert test_item['test'] == data[0]['test']
    test_item['test'] = None
    assert test_item is await ubpl.process(test_item)
    data = await ubpl.pull_data('SELECT * FROM test')
    assert test_item['test'] == data[0]['test']

    iubpl = ItemMysqlInsertUpdatePipeline(
        update_keys=['test'],
        host=mysql_server, port=mysql_port, user=mysql_user,
        password=mysql_password,
        to_dict=lambda x: x,
        database='test', table='test', buffer_length=1)
    await iubpl.on_spider_open()
    test_item['test'] = 'I love ant!'
    assert test_item is await iubpl.process(test_item)
    data = await iubpl.pull_data('SELECT * FROM test')
    assert test_item['test'] == data[0]['test']

    await ubpl.on_spider_close()
    await ibpl.on_spider_close()
    await iubpl.on_spider_close()
    await bpl.push_data('DROP TABLE test.test;DROP DATABASE test')


@pytest.mark.asyncio
async def test_redis_pipeline():
    redis_address = os.getenv('TEST_REDIS_ADDRESS', 'redis://localhost:6379')
    pl = ItemBaseRedisPipeline(redis_address)
    pool = await pl.create_redis()
    with await pool as conn:
        value = 'value'
        await conn.set('key', value)
        assert await conn.get('key') == value
    pool.close()
    await pool.wait_closed()
