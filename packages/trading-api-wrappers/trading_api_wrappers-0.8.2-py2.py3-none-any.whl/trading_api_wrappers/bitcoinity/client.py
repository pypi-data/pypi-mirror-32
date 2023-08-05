# local
from ..base import Client, Server

# API Server
PROTOCOL = 'https'
HOST = 'bitcoinity.org'


class Bitcoinity(Client):
    def __init__(self, timeout=15):
        super().__init__(Server(PROTOCOL, HOST), timeout)

    def ticker(self,
               currency: str,
               exchange: str,
               span: str):
        params = {
            'currency': currency,
            'exchange': exchange,
            'span': span,
        }
        url = self.url_for('markets/get_ticker')
        data = self.get(url, params=params)
        return data
