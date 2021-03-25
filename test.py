
from huobi.client.market import MarketClient
from huobi.constant import *
from huobi.utils import *
from models import Coinkline


market_client = MarketClient(init_log=True, url="https://api.huobi.be")
interval = CandlestickInterval.MIN5
symbol = "ethusdt"
# 200
list_obj = market_client.get_candlestick(symbol, interval, 10)
LogInfo.output("---- {interval} candlestick for {symbol} ----".format(interval=interval, symbol=symbol))
LogInfo.output_list(list_obj)

for item in list_obj:
    coin = await Coinkline.get_or_create(symbol=item.symbol, name=item.symbol[0:-4], time_id=item.id)
    await Coinkline.filter(id=item.id).update(                
        count=item.count,
        open=item.open,
        close=item.close,
        low=item.low,
        high=item.high,
        vol=item.vol,
        period=CandlestickInterval.MIN5
    )
# Coinkline