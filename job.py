from huobi.client.market import MarketClient
from huobi.constant import *
from huobi.utils import *
from models import Coinkline

async def kline_job():
    market_client = MarketClient(init_log=True, url="https://api.huobi.be")
    interval = CandlestickInterval.MIN5
    symbol = "ethusdt"
    # 200
    list_obj = market_client.get_candlestick(symbol, interval, 10)
    print(interval)
    print(len(list_obj))
    for item in list_obj:
        print(item.id)
        coin, created = await Coinkline.get_or_create(symbol=symbol, name=symbol[0:-4], time_id=item.id, period=interval)
        print(coin)
        await Coinkline.filter(id=coin.id).update(                
            count=item.count,
            open=item.open,
            close=item.close,
            low=item.low,
            high=item.high,
            vol=item.vol,
        )