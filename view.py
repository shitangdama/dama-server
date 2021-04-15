



class Handler:
    def  __init__(self, ws):
        self._ws = ws
    
    async def listen_ws(self):
        async for msg in self._ws:
            if msg.type == WSMsgType.TEXT:
                await print(msg)
            elif msg.type == WSMsgType.CLOSED:
                break
            elif msg.type == WSMsgType.ERROR:
                break
    async def test(self):
        return {'coins': "Hello, world"}


    # @aiohttp_jinja2.template('subscribe.html')
    # async def subscribe(self, request):
    #     coins = await CoinSubscribe.all().order_by("id")
    #     return {'coins': coins}
