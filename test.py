from huobi.client.market import MarketClient

market_client = MarketClient(url="https://api.huobi.be")
print(11111)
list_obj = market_client.get_market_tickers()
for item in list_obj:
    print(item.print_object())