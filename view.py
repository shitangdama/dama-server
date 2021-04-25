import json
import gzip
import aiohttp
import aiohttp_jinja2
from aiohttp import web
from models import CoinTicker, CoinSubscribe

async def callback(msg):
    print(msg)

class Handler:
    def  __init__(self, ws):
        self.ws = ws

# 这里处理下
    async def connect(self):
        print(44444)
        while True:
            msg = await self.ws.receive()
            print(2222)
            if msg.type == aiohttp.WSMsgType.TEXT:
                await callback(msg.data)
            elif msg.type == aiohttp.WSMsgType.BINARY:
                data = json.loads(gzip.decompress(msg.data).decode())
                # 响应websocket server ping 保持连接
                
                print(data)
                ping = data.get('ping')
                print("{'pong': %s}"% (ping))
                if ping:
                    await ws.send_str("{'pong': %s}"% (ping))
                else:
                    # 正常处理消息
                    await callback(data)
            elif msg.type == aiohttp.WSMsgType.CLOSED:
                await callback(msg)
                break
            elif msg.type == aiohttp.WSMsgType.ERROR:
                await callback(msg)
                break



    async def subscribe_coin(self, request):
        data = await request.json()
        print(data["symbol"])
        symbol = data["symbol"]
        period = "1min"
        channel = "{'sub': 'market.%s.kline.%s', 'id': 'id1'}" % (symbol, period)
        # await self.ws.send_str(channel)
        await self.ws.send_str("""{"sub": "market.ethusdt.kline.1min","id": "id10"}""")
        await self.ws.send_str("""{"req": "market.ethusdt.kline.1min","id": "id10"}""")
        print(33333)
        return web.json_response({'coins': "Hello, world"})

    async def test(self):
        return {'coins': "Hello, world"}

    @aiohttp_jinja2.template('index.html')
    async def index(self, request):
        coins = await CoinTicker.all().order_by("-vol")
        return {'coins': coins}

    @aiohttp_jinja2.template('subscribe.html')
    async def subscribe(self, request):
        coins = await CoinSubscribe.all().order_by("id")
        return {'coins': coins}

    # @aiohttp_jinja2.template('subscribe.html')
    # async def subscribe(self, request):
    #     coins = await CoinSubscribe.all().order_by("id")
    #     return {'coins': coins}
