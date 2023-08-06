from ..base import Server

# API Server
PROTOCOL = 'https'
HOST = 'www.bitstamp.net/api'


# Bitstamp API server
class BitstampServer(Server):

    def __init__(self):
        super().__init__(PROTOCOL, HOST)
