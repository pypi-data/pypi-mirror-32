import asyncio
import functools
from aiossdb.pool import create_pool


class Client:
    def __init__(self, host='127.0.0.1', port=8888, password=None, timeout=None, max_connection=100, loop=None,
                 encoding='utf-8'):
        self.host = host
        self.port = port
        self.password = password
        self.timeout = timeout
        self.max_connection = max_connection
        self.encoding = encoding

        if loop is None:
            # loop = asyncio.new_event_loop()
            # asyncio.set_event_loop(loop)
            loop = asyncio.get_event_loop()

        self.loop = loop

        self._pool = None

    async def get_pool(self):
        if self._pool is None:
            self._pool = await create_pool(
                (self.host, self.port), password=self.password, loop=self.loop,
                timeout=self.timeout, maxsize=self.max_connection, encoding=self.encoding
            )
        return self._pool

    async def execute(self, cmd, *args, **kwargs):
        pool = await self.get_pool()
        res = await pool.execute(cmd, *args, **kwargs)
        return res

    def __getattr__(self, item):
        if item not in self.__dict__:
            self.__dict__[item] = functools.partial(self.execute, item)

        return self.__dict__[item]

    async def close(self):
        if self._pool:
            self._pool.close()
            await self._pool.wait_closed()
            self._pool = None

    async def __aenter__(self):
        return self

    async def __aexit__(self, exc_type, exc_val, exc_tb):
        await self.close()
