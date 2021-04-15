  
import aiohttp
from aiohttp import web
from asyncio import create_task

async def callback(msg):
    print(msg)


async def websocket(session):
    print('xxx')
    async with session.ws_connect('wss://api.huobi.de.com/ws') as ws:
        print(ws)
        async for msg in ws:
            print(msg)
            if msg.type == aiohttp.WSMsgType.TEXT:
                await callback(msg.data)
            elif msg.type == aiohttp.WSMsgType.CLOSED:
                break
            elif msg.type == aiohttp.WSMsgType.ERROR:
                break


async def init(app):
    session = aiohttp.ClientSession()
    app['websocket_task'] = create_task(websocket(session))
    print(app['websocket_task'])

app = web.Application()
app.on_startup.append(init)

web.run_app(app)