import json
import gzip
import asyncio
import aiohttp
from aiohttp import web

# ws://hq.sinajs.cn/wskt?list=s_sh000001

global client_ws

async def client():
    async with aiohttp.ClientSession() as session:
        async with session.ws_connect('wss://api.huobi.de.com/ws',
                                      heartbeat=1) as ws:
            global client_ws
            client_ws = ws
            try:
                async for msg in ws:
                    if msg.type == aiohttp.WSMsgType.TEXT:
                        print(msg.data)
                    elif msg.type == aiohttp.WSMsgType.BINARY:
                        data = json.loads(gzip.decompress(msg.data).decode())
                        # 响应websocket server ping 保持连接
                        
                        print(data)
                        ping = data.get('ping')
                        
                        if ping:
                            print(json.dumps({'pong': ping}))
                            await ws.send_str(json.dumps({"pong": ping}))
                        else:
                            # 正常处理消息
                            print(data)
                    elif msg.type == aiohttp.WSMsgType.CLOSED:
                        print(msg)
                        break
                    elif msg.type == aiohttp.WSMsgType.ERROR:
                        print(msg)
                        break
            except Exception as e:
                print(f"ERROR: {type(e)}:{e}")
                raise e


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


async def server():
    app = web.Application()

    app.add_routes([web.get('/ws', handle)])
    app.add_routes([web.get('/', test)])
    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, 'localhost', 8080)
    await site.start()


loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.gather(
    server(),
    client()
))