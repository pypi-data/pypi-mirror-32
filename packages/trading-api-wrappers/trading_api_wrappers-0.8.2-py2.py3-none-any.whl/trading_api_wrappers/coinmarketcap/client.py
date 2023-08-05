# local
from ..base import Client, Server
from ..common import clean_parameters

# API Server
PROTOCOL = 'https'
HOST = 'api.coinmarketcap.com'
VERSION = 'v1'


class CoinMarketCap(Client):

    def __init__(self, timeout: int=120):
        super().__init__(Server(PROTOCOL, HOST, VERSION), timeout)
        self._currencies = None

    def ticker(self,
               currency: str=None,
               convert: str=None,
               start: int=None,
               limit: int=None):
        params = {
            'start': start,
            'limit': limit,
            'convert': convert,
        }
        if currency:
            if len(currency) == 3:
                currency = self._get_symbol(currency)['value']
            url = self.url_for('ticker/%s/', currency)
            data = self.get(url, params=params)[0]
        else:
            url = self.url_for('ticker/')
            data = self.get(url, params=params)
        return data

    def price(self,
              currency: str,
              convert: str=None):
        ticker = self.ticker(currency, convert)
        return float(ticker[f"price_{convert or 'usd'}".lower()])

    def stats(self, convert: str=None):
        params = {
            'convert': convert,
        }
        url = self.url_for('global/')
        data = self.get(url, params=params)
        return data

    def _get_currencies(self):
        ticker = self.ticker()
        return {currency['symbol']: dict(value=currency['id'], decimals=8)
                for currency in ticker}

    def _get_symbol(self, currency: str):
        if self._currencies is None:
            self._currencies = self._get_currencies()
        return self._currencies[currency.upper()]

    def get(self, url: str, headers: dict=None, params: dict=None):
        clean_params = clean_parameters(params)
        return super(CoinMarketCap, self).get(url, params=clean_params)
