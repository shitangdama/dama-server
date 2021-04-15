from huobi.client.market import MarketClient
from huobi.constant import *
from huobi.exception.huobi_api_exception import HuobiApiException
from huobi.model.market.candlestick_event import CandlestickEvent


def callback(obj_event: 'MarketDetailEvent'):
    obj_event.print_object()
    print()


def error(e: 'HuobiApiException'):
    print(e.error_code + e.error_message)

market_client = MarketClient(url='wss://api.huobi.de.com/swap-ws')
market_client.sub_market_detail("btcusdt", callback)
