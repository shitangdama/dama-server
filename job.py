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
    for item in list_obj:
        coin, created = await Coinkline.get_or_create(symbol=symbol, name=symbol[0:-4], time_id=item.id, period=interval)
        await Coinkline.filter(id=coin.id).update(                
            count=item.count,
            open=item.open,
            close=item.close,
            low=item.low,
            high=item.high,
            vol=item.vol,
        )

async def info_job():
    market_client = MarketClient(init_log=True, url="https://api.huobi.be")
    list_obj = market_client.get_market_tickers()

    btc = {}
    items = []
    for item in list_obj:
        if item.symbol.endswith("usdt"):
            tmp = {
                "amount": item.amount,
                "count": item.count,
                "open": item.open,
                "close": item.close,
                "low": item.low,
                "high":item.high,
                "vol":item.vol,
                "symbol":item.symbol,
                "name": item.symbol[0:-4]
            }
            items.append(tmp)
            if item.symbol == "btcusdt":
                btc = tmp
    btc["self_trade"] = ((btc["close"] - btc["open"])/btc["open"]) * 100
    for coin in items:
        if coin["symbol"] == "btcusdt":
            coin["self_trade"] = btc["self_trade"]
            coin["contrast_trade"] = 0.0
        else:
            coin["self_trade"] = ((coin["close"] - coin["open"])/coin["open"]) * 100
            coin["contrast_trade"] = coin["self_trade"] - btc["self_trade"]
        coin_data, created = await CoinTicker.get_or_create(symbol=coin["symbol"],name=coin["name"])
        await CoinTicker.filter(symbol=coin_data.symbol).update(                
                amount=coin["amount"],
                count=coin["count"],
                open=coin["open"],
                close=coin["close"],
                low=coin["low"],
                high=coin["high"],
                vol=coin["vol"],
                self_trade= coin["self_trade"],
                contrast_trade = coin["contrast_trade"],
            )