from pybithumb.client import Bithumb


def get_tickers():
    return Bithumb.get_tickers()


def get_market_detail(currency):
    return Bithumb.get_market_detail(currency)


def get_current_price(currency):
    return Bithumb.get_current_price(currency)


def get_orderbook(currency):
    return Bithumb.get_orderbook(currency)
