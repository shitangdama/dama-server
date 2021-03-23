from tortoise import Model, fields

class CoinTicker(Model):
    id = fields.IntField(pk=True)
    amount = fields.CharField(50)

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