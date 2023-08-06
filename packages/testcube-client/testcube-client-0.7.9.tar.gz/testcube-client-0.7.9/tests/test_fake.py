#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
Tests for `business` module. Requires a testcube server.
"""

import unittest
from os import chdir

from faker import Faker

from testcube_client import business
from testcube_client.request_helper import *
from testcube_client.settings import enable_debug_log

fake = Faker()

server = 'http://127.0.0.1:8000'
result_dir = r'C:\temp'


class TestCases(unittest.TestCase):
    def setUp(self):
        enable_debug_log()
        register_client(server, force=True)
        self.team = business.get_or_create_team('Core')
        self.product = business.get_or_create_product('TestCube', self.team)
        chdir(result_dir)

    def tearDown(self):
        pass

    def test_add_run_steps(self):
        run_url = business.start_run(team_name='Core',
                                     product_name='TestCube',
                                     run_name=fake.text(100))

        logging.info(run_url)

        business.finish_run('re*.xml')

    def test_add_run_one(self):
        business.run(team_name='Core',
                     product_name='TestCube',
                     run_name=fake.text(100),
                     result_xml_pattern='*.xml')

        # start second time
        business.run(team_name='Core',
                     product_name='TestCube',
                     run_name=fake.text(100),
                     result_xml_pattern='*.xml')

    def test_cleanup_runs(self):
        business.cleanup_runs(days=60)


if __name__ == '__main__':
    unittest.main()
