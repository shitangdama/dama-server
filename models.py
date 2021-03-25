from tortoise import Model, fields

class CoinTicker(Model):
    id = fields.IntField(pk=True)
    symbol = fields.CharField(unique=True,max_length=128)
    name = fields.CharField(unique=True,max_length=128)
    amount = fields.FloatField(default=0.0)
    count = fields.FloatField(default=0.0)
    open = fields.FloatField(default=0.0)
    close = fields.FloatField(default=0.0)
    high = fields.FloatField(default=0.0)
    low = fields.FloatField(default=0.0)
    vol = fields.FloatField(default=0.0)
    self_trade = fields.FloatField(default=0.0)
    contrast_trade = fields.FloatField(default=0.0)

    def __str__(self):
        return f"User {self.id}: {self.symbol}"

    class Meta:
        table = "coin_ticker"

class Coinkline(Model):
    id = fields.IntField(pk=True)
    symbol = fields.CharField(unique=True,max_length=128)
    name = fields.CharField(unique=True,max_length=128)
    period = fields.CharField(max_length=128)
    time_id = fields.BigIntField()
    count = fields.BigIntField(default=0)
    open = fields.FloatField(default=0.0)
    close = fields.FloatField(default=0.0)
    high = fields.FloatField(default=0.0)
    low = fields.FloatField(default=0.0)
    vol = fields.FloatField(default=0.0)

    def __str__(self):
        return f"User {self.id}: {self.symbol}"

    class Meta:
        table = "coin_kline"


class CoinSubscribe(Model):
    id = fields.IntField(pk=True)
    symbol = fields.CharField(unique=True,max_length=128)
    name = fields.CharField(unique=True,max_length=128)
    day_trend = fields.JSONField()
    min_trend = fields.JSONField()
    hour_trend = fields.JSONField()
    day_trend = fields.JSONField()
    min_trend = fields.JSONField()
    hour_trend = fields.JSONField()
    class Meta:
        table = "coin_subscribe"

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