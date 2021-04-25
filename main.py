import logging
import sys
import pathlib
import asyncio
import jinja2
# import uvloop


import aiohttp_jinja2

from aiohttp import web, ClientSession
from apscheduler.schedulers.asyncio import AsyncIOScheduler
from tortoise.contrib.aiohttp import register_tortoise

from job import kline_job, info_job
from view import Handler
from asyncio import create_task



from huobi.client.market import MarketClient
from huobi.constant import *
from huobi.exception.huobi_api_exception import HuobiApiException
from huobi.model.market.candlestick_event import CandlestickEvent


PROJECT_ROOT = pathlib.Path(__file__).parent

# asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
loop = asyncio.get_event_loop()

async def init(app):
    pass
async def init_app():
    app = web.Application()

    # app.add_routes([web.get('/', h.test)])
    app.on_startup.append(init)


    return app

async def main():
    # scheduler = AsyncIOScheduler()
    # scheduler.add_job(info_job, trigger='cron', minute='*/1')
    # scheduler.add_job(kline_job, trigger='cron', minute='*/1')
    # scheduler.start()
    app = web.Application()

    aiohttp_jinja2.setup(
        app,
        loader=jinja2.FileSystemLoader(str(PROJECT_ROOT / 'templates'))
    )

    register_tortoise(
        app, db_url="postgres://postgres:changeme@localhost:5432/postgres", modules={"models": ["models"]}, generate_schemas=True
    )

    app.router.add_static('/static/', path=PROJECT_ROOT / 'static', name='static')

    
    session = ClientSession()
    ws = await session.ws_connect('wss://api.huobi.de.com/ws', timeout=30)
    h = Handler(ws)
    task1 = create_task(h.connect())

    app.add_routes([web.get('/', h.index),
                web.get('/subscribe', h.subscribe),
                web.post('/subscribe', h.subscribe_coin)])


    logging.basicConfig(level=logging.DEBUG)
    # app = init_app()s
    task2 = web.run_app(app, port=8080)

    loop.run_until_complete(asyncio.gather(task1, task2))
if __name__ == '__main__':
   asyncio.run(main())