#!/usr/bin/env python

import os
import sys


filepath = os.path.realpath(__file__)
filepath = "/".join(filepath.split("/")[0:-2])
sys.path.append(filepath)

from InventoryFilters import DummyInventory  # noqa


class FilterModule(object):
    ''' A filter to fix interface's name format '''

    def filters(self):

        inv = DummyInventory.DummyInventory()
        return {
            'dummy_inv': inv.get_inventory
        }
