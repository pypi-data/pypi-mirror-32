#!/usr/bin/env python3
import sys
import logging
import random

from caproto.benchmarking import set_logging_level
from caproto.curio.server import start_server
from caproto.server import (pvproperty, PVGroup, pvfunction)


logger = logging.getLogger(__name__)


class MyPVGroup(PVGroup):
    'Example group of PVs, where the prefix is defined on instantiation'
    # PV #1: {prefix}random - defaults to dtype of int
    @pvproperty
    async def fixed_random(self, instance):
        'Random integer between 1 and 100'
        logger.debug('read random')
        return random.randint(1, 100)

    @pvfunction(default=[0])
    async def get_random(self,
                         low: int=100,
                         high: int=1000) -> int:
        'A configurable random number'
        low, high = low[0], high[0]
        return random.randint(low, high)


if __name__ == '__main__':
    import curio
    from pprint import pprint

    try:
        prefix = sys.argv[1]
    except IndexError:
        prefix = 'rpc:'

    macros = {}

    set_logging_level(logging.DEBUG)
    logger.setLevel(logging.DEBUG)
    logging.basicConfig()

    logger.info('Starting up: prefix=%r macros=%r', prefix, macros)
    ioc = MyPVGroup(prefix=prefix, macros=macros)

    # here's what accessing a pvproperty descriptor looks like:
    print(f'fixed_random using the descriptor getter is: {ioc.fixed_random}')
    print(f'get_random using the descriptor getter is: {ioc.get_random}')
    print('get_random is an autogenerated subgroup with PVs:')
    for pvspec in ioc.get_random.pvdb.items():
        print(f'\t{pvspec!r}')

    # here is the auto-generated pvdb:
    pprint(ioc.pvdb)

    curio.run(start_server, ioc.pvdb)
