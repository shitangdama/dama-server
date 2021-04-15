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

from models import CoinTicker, CoinSubscribe
from job import kline_job, info_job
from view import Handler
from aiohttp import web, ClientSession
from asyncio import create_task

from huobi.client.market import MarketClient
from huobi.constant import *
from huobi.exception.huobi_api_exception import HuobiApiException
from huobi.model.market.candlestick_event import CandlestickEvent


# def callback(candlestick_event: 'CandlestickEvent'):
#     candlestick_event.print_object()
#     print("\n")


# def error(e: 'HuobiApiException'):
#     print(e.error_code + e.error_message)

# def test(loop):
#     market_client = MarketClient(url='wss://api.huobi.de.com/swap-ws')
#     market_client.sub_candlestick("btcusdt,ethusdt", CandlestickInterval.MIN1, callback, error)

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
    coins = await CoinSubscribe.all().order_by("id")
    return {'coins': coins}

def setup_routes(app):
    app.router.add_get('/', index)
    app.router.add_get('/subscribe', subscribe)

async def callback(msg):
    print(msg)

async def websocket(session):
    print(22222)
    async with session.ws_connect('wss://api.huobi.de.com/swap-ws') as ws:
        print(33333)
        await ws.send_json({"sub": "market.btcusdt","id": "id1"})
        print(1111)
        async for msg in ws:
            if msg.type == aiohttp.WSMsgType.TEXT:
                await callback(msg.data)
                break
            elif msg.type == aiohttp.WSMsgType.BINARY:
                data = json.loads(gzip.decompress(msg.data).decode())
                if isinstance(msg, dict) and 'ping' in msg.keys():
                    await ws.send_json({"pong": msg.get("ping")})
                    

                else:
                    # 正常处理消息
                    await callback(data)
            elif msg.type == aiohttp.WSMsgType.CLOSED:
                await callback(msg.data)
                break
            elif msg.type == aiohttp.WSMsgType.ERROR:
                await callback(msg.data)
                break

async def init(app):
    session = ClientSession()
    app['websocket_task'] = create_task(websocket(session))

async def init_app():
    app = web.Application()
    print(1111) 
    # session = ClientSession()
    # ws = await session.ws_connect('wss://api.huobi.de.com/swap-ws')
    # h = Handler(ws)
    # h.listen_ws()
    # setup_routes(app)

    # app.add_routes([web.get('/', h.test)])
    app.on_startup.append(init)
    # aiohttp_jinja2.setup(
    #     app,
    #     loader=jinja2.FileSystemLoader(str(PROJECT_ROOT / 'templates'))
    # )

    # register_tortoise(
    #     app, db_url="postgres://postgres:changeme@localhost:5432/postgres", modules={"models": ["models"]}, generate_schemas=True
    # )

    # app.router.add_static('/static/', path=PROJECT_ROOT / 'static', name='static')

    return app

def main():
    # scheduler = AsyncIOScheduler()
    # scheduler.add_job(info_job, trigger='cron', minute='*/1')
    # scheduler.add_job(kline_job, trigger='cron', minute='*/1')
    # scheduler.start()

    logging.basicConfig(level=logging.DEBUG)
    app = init_app()
    web.run_app(app, port=8080)

if __name__ == '__main__':
    main()