from huobi.client.market import MarketClient
from huobi.constant import *
from huobi.utils import *
from models import Coinkline, CoinSubscribe, CoinTicker

async def kline_job():
    market_client = MarketClient(init_log=True, url="https://api.huobi.de.com")
    interval = CandlestickInterval.MIN5
    eth_symbol = "ethusdt"
    btc_symbol = "btcusdt"
    btc_result = await get_kline(market_client, btc_symbol, interval)
    eth_result = await get_kline(market_client, eth_symbol, interval)

    eth_coin_subscribe, created = await CoinSubscribe.get_or_create(symbol=eth_symbol, name=eth_symbol[0:-4])
    btc_coin_subscribe, created = await CoinSubscribe.get_or_create(symbol=btc_symbol, name=btc_symbol[0:-4])

    
    btc_time_list = sorted([i for i in eth_result.keys()])
    eth_time_list = sorted([i for i in eth_result.keys()])
    print(btc_time_list)
    btc_vol = []
    btc_close = []
    btc_one=btc_result[btc_time_list[0]]["close"]
    eth_one=eth_result[eth_time_list[0]]["close"]
    for num in btc_time_list:
        if btc_result[num]:
            btc_vol.append(btc_result[num]["vol"])
            btc_close.append(((btc_result[num]["close"] - btc_one)/btc_one)*100)

    eth_vol = []
    eth_close = []
    eth_contrast = []
    for num in eth_time_list:
        eth_vol.append(eth_result[num]["vol"])
        eth_close.append(((eth_result[num]["close"]- btc_one)/btc_one)*100)
        if btc_result[num] and eth_result[num]:
            eth_contrast.append(((eth_result[num]["close"] - btc_result[num]["close"])/btc_one)*100)

    await CoinSubscribe.filter(id=btc_coin_subscribe.id).update(                
        trend_5min={
            "vol": btc_vol,
            "close": btc_close
        },
    )

    await CoinSubscribe.filter(id=eth_coin_subscribe.id).update(                
        trend_5min={
            "vol": btc_vol,
            "close": btc_close,
            "contrast": eth_contrast
        },
    )


    # asyncio.gather(ss
    # 200
    # list_obj = market_client.get_candlestick(symbol, interval, 200)
    # for item in list_obj:
    #     coin, created = await Coinkline.get_or_create(symbol=symbol, name=symbol[0:-4], time_id=item.id, period=interval)
    #     await Coinkline.filter(id=coin.id).update(                
    #         count=item.count,
    #         open=item.open,
    #         close=item.close,
    #         low=item.low,
    #         high=item.high,
    #         vol=item.vol,
    #     )

    # 这里先生成三个走势图 走势图 相对走势图 成交量

async def get_kline(client, symbol, interval):
    list_obj = client.get_candlestick(symbol, interval, 20)
    result = {}
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
        result[coin.time_id] = {
            "close": item.close,
            "vol": item.vol,
        }
    return result

async def info_job():
    market_client = MarketClient(init_log=True, url="https://api.huobi.de.com")
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