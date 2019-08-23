{
  "req": "market.btcusdt.kline.1min",
  "id": "id10"
}

#1.KLine
sub/req
market.$symbol.kline.$period
K线 数据，包含单位时间区间的开盘价、收盘价、最高价、最低价、成交量、成交额、成交笔数等数据 $period 可选值：{ 1min, 5min, 15min, 30min, 60min, 4hour,1day, 1mon, 1week, 1year }	

#2.Market Depth
sub/req
market.$symbol.depth.$type

盘口深度，按照不同 step 聚合的买一、买二、买三等和卖一、卖二、卖三等数据 $type 可选值：{ step0, step1, step2, step3, step4, step5, percent10 } （合并深度0-5）；step0时，不合并深度

#3.Trade Detail
sub/req
market.$symbol.trade.detail

成交记录，包含成交价格、成交量、成交方向等信息

#4.Market Detail
sub/req
market.$symbol.detail

#5.Market Tickers
sub
market.tickers
所有对外公开交易对的 日K线、最近24小时成交量等信息


####一下需要验证
#6.Accounts
sub, unsub
accounts
订阅账户资产变更
#7.Orders
orders.$symbol.update
sub,unsub
订阅订单变更（新）
#8.Accounts list
req
accounts.list
请求账户资产信息

#Order list
req	
accounts.list
请求订单信息	

#order detail
req
orders.detail
请求某个订单明细	