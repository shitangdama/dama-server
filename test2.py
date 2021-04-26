import json
import gzip
import logging
import sys
import pathlib
import asyncio
import aiohttp
import jinja2
import aiohttp_jinja2
from aiohttp import web
from tortoise.contrib.aiohttp import register_tortoise

from models import CoinTicker, CoinSubscribe
from apscheduler.schedulers.asyncio import AsyncIOScheduler
from job import kline_job, info_job

PROJECT_ROOT = pathlib.Path(__file__).parent
global client_ws

async def handle(request):
    ws = web.WebSocketResponse(
            autoping=True, heartbeat=10.0, compress=True, max_msg_size=0)
    print(request.match_info.get('name', "nocilantro"))
    await ws.prepare(request)
    for i in range(10):
        await ws.send_str(str(i))
        await asyncio.sleep(1)
    async for msg in ws:
        if msg.type == aiohttp.WSMsgType.TEXT:
            data = msg.data
            if data:
                await ws.send_str('recevive' + data)
    return ws

async def test(request):
    global client_ws
    print(client_ws)
    await client_ws.send_str("""{"sub": "market.ethusdt.kline.1min","id": "id10"}""")
    return web.json_response({'coins': "Hello, world"})

async def client():
    async with aiohttp.ClientSession() as session:
        async with session.ws_connect('wss://api.huobi.de.com/ws', heartbeat=1) as ws:
            global client_ws
            client_ws = ws
            try:
                async for msg in ws:
                    if msg.type == aiohttp.WSMsgType.TEXT:
                        print("TEXT====")
                        print(msg)
                    elif msg.type == aiohttp.WSMsgType.BINARY:
                        data = json.loads(gzip.decompress(msg.data).decode())
                        print("BINARY====")
                        # 响应websocket server ping 保持连接
                        ping = data.get('ping')
                        print(ping)
                        if ping:
                            await ws.send_str(json.dumps({"pong": ping}))
                        else:
                            print(data)
                    elif msg.type == aiohttp.WSMsgType.CLOSED:
                        print("CLOSED====", msg)
                        break
                    elif msg.type == aiohttp.WSMsgType.ERROR:
                        print("ERROR====", msg)
                        break
            except Exception as e:
                print(f"ERROR: {type(e)}:{e}")
                raise e


@aiohttp_jinja2.template('index.html')
async def index(request):
    coins = await CoinTicker.all().order_by("-vol")
    return {'coins': coins}

@aiohttp_jinja2.template('subscribe.html')
async def subscribe(request):
    coins = await CoinSubscribe.all().order_by("id")
    return {'coins': coins}

async def server():
    app = web.Application()
    # scheduler = AsyncIOScheduler()
    # scheduler.add_job(info_job, trigger='cron', minute='*/1')
    # # scheduler.add_job(kline_job, trigger='cron', minute='*/1')
    # scheduler.start()

    aiohttp_jinja2.setup(
        app,
        loader=jinja2.FileSystemLoader(str(PROJECT_ROOT / 'templates'))
    )

    register_tortoise(
        app, db_url="postgres://postgres:kbr199sd5shi@localhost:5432/postgres", modules={"models": ["models"]}, generate_schemas=True
    )

    app.router.add_static('/static/', path=PROJECT_ROOT / 'static', name='static')

    app.add_routes([web.get('/ws', handle)])
    app.add_routes([web.get('/', index)])
    app.add_routes([web.get('/subscribe', subscribe)])
    
    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, 'localhost', 8080)
    print("server start")
    await site.start()


loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.gather(
    server(),
    client()
))