# -*- coding: utf-8 -*-
import sys

sys.path.append('..')
sys.path.append('../..')

from utils.config import Config

class test_config():

    wallet_A_address = Config.NODES[8]["address"]
    wallet_B_address = Config.NODES[7]["address"]

    node_B_puiblic_key = Config.NODES[7]["pubkey"]

    vote_price = "3000"

    blocks_per_round = 5

    punish_ratio = 0.5


