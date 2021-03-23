from tortoise import Model, fields

class CoinTicker(Model):
    id = fields.IntField(pk=True)
    symbol = fields.TextField()
    name = fields.TextField()
    amount = fields.FloatField()
    count = fields.FloatField()
    open = fields.FloatField()
    close = fields.FloatField()
    high = fields.FloatField()
    low = fields.FloatField()
    vol = fields.FloatField()
    bid = fields.FloatField()
    bidSize = fields.FloatField()
    ask = fields.FloatField()
    askSize = fields.FloatField()

    self_trade = fields.JSONField()
    contrast_trade = fields.JSONField()

    def __str__(self):
        return f"User {self.id}: {self.name}"

    class Meta:
        table = "coin_ticker"
        # PrintBasic.print_basic(self.amount, format_data + "Amount")
        # PrintBasic.print_basic(self.count, format_data + "Count")
        # PrintBasic.print_basic(self.open, format_data + "Opening Price")
        # PrintBasic.print_basic(self.close, format_data + "Last Price")
        # PrintBasic.print_basic(self.low, format_data + "Low Price")
        # PrintBasic.print_basic(self.high, format_data + "High Price")
        # PrintBasic.print_basic(self.vol, format_data + "Vol")
        # PrintBasic.print_basic(self.symbol, format_data + "Trading Symbol")
        # PrintBasic.print_basic(self.bid, format_data + "Best Bid Price")
        # PrintBasic.print_basic(self.bidSize, format_data + "Best Bid Size")
        # PrintBasic.print_basic(self.ask, format_data + "Best Ask Price")
        # PrintBasic.print_basic(self.askSize, format_data + "Best Ask Size")



# symbol	string	true	NA	交易对	btcusdt, ethbtc等（如需获取杠杆ETP净值K线，净值symbol = 杠杆ETP交易对symbol + 后缀‘nav’，例如：btc3lusdtnav）
# period	string	true	NA	返回数据时间粒度，也就是每根蜡烛的时间区间	1min, 5min, 15min, 30min, 60min, 4hour, 1day, 1mon, 1week, 1year
# size	integer	false	150	返回 K 线数据条数	[1, 2000]
# id	long	调整为新加坡时间的时间戳，单位秒，并以此作为此K线柱的id
# amount	float	以基础币种计量的交易量
# count	integer	交易次数
# open	float	本阶段开盘价
# close	float	本阶段收盘价
# low	float	本阶段最低价
# high	float	本阶段最高价
# vol	float	以报价币种计量的交易量