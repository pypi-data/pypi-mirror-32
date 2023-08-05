#!/usr/bin/env python
import ipaddress
import os
import threading
import logging
try:
    from functools import lru_cache
except ImportError:
    from backports.functools_lru_cache import lru_cache

import six
from tinydb import TinyDB, Query
from tinydb.middlewares import CachingMiddleware

from iptocc.json_storage_read_only import JSONStorageReadOnly

__author__ = "Ronie Martinez"
__copyright__ = "Copyright 2017, Ronie Martinez"
__credits__ = ["Ronie Martinez"]
__license__ = "MIT"
__version__ = "1.0.2"
__maintainer__ = "Ronie Martinez"
__email__ = "ronmarti18@gmail.com"
__status__ = "Production"

dir_path = os.path.dirname(os.path.realpath(__file__))
caching_middleware = CachingMiddleware(JSONStorageReadOnly)
database = TinyDB(os.path.join(dir_path, 'rir_statistics_exchange.json'), storage=caching_middleware)
query = Query()

lock = threading.Lock()
logger = logging.getLogger(__name__)


@lru_cache(maxsize=100000)
def ipv4_get_country_code(ip_address):
    with lock:
        for record in database.search(query.type == 'ipv4'):
            start_address = ipaddress.IPv4Address(record.get('start'))
            if start_address <= ip_address < start_address + record.get('value'):
                country_code = record.get('country_code')
                if six.PY2:
                    country_code = str(country_code)
                logger.debug('Country code for ip=%s is %s.', ip_address, country_code)
                return country_code
        logger.debug('Cannot find country code for ip=%s', ip_address)
        return None


@lru_cache(maxsize=100000)
def ipv6_get_country_code(ip_address):
    with lock:
        for record in database.search(query.type == 'ipv6'):
            network = ipaddress.IPv6Network('{}/{}'.format(record.get('start'), record.get('value')))
            if ip_address in network:
                country_code = record.get('country_code')
                if six.PY2:
                    country_code = str(country_code)
                logger.debug('Country code for ip=%s is %s.', ip_address, country_code)
                return country_code
        logger.debug('Cannot find country code for ip=%s', ip_address)
        return None


def get_country_code(ip_address):
    # convert string to ipaddress.IPv4Address or ipaddress.IPv6Address
    if isinstance(ip_address, six.text_type):
        ip_address = ipaddress.ip_address(ip_address)
    if six.PY2 and isinstance(ip_address, six.string_types):
        ip_address = ipaddress.ip_address(unicode(ip_address))
    if isinstance(ip_address, ipaddress.IPv4Address):
        return ipv4_get_country_code(ip_address)  # IPv4
    return ipv6_get_country_code(ip_address)  # IPv6
