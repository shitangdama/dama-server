import logging
import sys
import pathlib
import asyncio
# import uvloop
# from huobi.client.market import MarketClient

import aiohttp_jinja2
import jinja2
from aiohttp import web
from apscheduler.schedulers.asyncio import AsyncIOScheduler
from tortoise.contrib.aiohttp import register_tortoise

from models import CoinTicker
from job import kline_job, info_job

PROJECT_ROOT = pathlib.Path(__file__).parent

# asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
loop = asyncio.get_event_loop()


@aiohttp_jinja2.template('index.html')
async def index(request):
    coins = await CoinTicker.all().order_by("-vol")
    print(coins)
    return {'coins': coins}

@aiohttp_jinja2.template('subscribe.html')
async def subscribe(request):
    # coins = await CoinTicker.all().order_by("-vol")
    # print(coins)
    return {'coins': 1111}

def setup_routes(app):
    app.router.add_get('/', index)
    app.router.add_get('/subscribe', subscribe)

async def init_app():
    app = web.Application()
    # app['config'] = config

    setup_routes(app)
    aiohttp_jinja2.setup(
        app,
        loader=jinja2.FileSystemLoader(str(PROJECT_ROOT / 'templates'))
    )

    register_tortoise(
        app, db_url="postgres://postgres:changeme@localhost:5432/postgres", modules={"models": ["models"]}, generate_schemas=True
    )

    app.router.add_static('/static/', path=PROJECT_ROOT / 'static', name='static')

    return app

def main():
    scheduler = AsyncIOScheduler()
    # scheduler.add_job(update, trigger='cron', minute='*/1')
    scheduler.add_job(kline_job, trigger='cron', minute='*/1')
    scheduler.start()

    logging.basicConfig(level=logging.DEBUG)
    app = init_app()
    web.run_app(app, port=8080)

if __name__ == '__main__':
    main()