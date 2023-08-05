"""
    This file contains all unit tests for the monitor-rules-table in the database. (Corresponding to the file:
    'flask_monitoringdashboard/database/monitor-rules.py')
    See __init__.py for how to run the test-cases.
"""

import unittest

from flask_monitoringdashboard.database import session_scope
from flask_monitoringdashboard.test.utils import set_test_environment, clear_db, add_fake_data, NAME, TIMES


class TestMonitorRule(unittest.TestCase):

    def setUp(self):
        set_test_environment()
        clear_db()
        add_fake_data()

    def test_get_monitor_rules(self):
        """
            Test whether the function returns the right values.
        """
        from flask_monitoringdashboard.database.monitor_rules import get_monitor_rules
        from flask_monitoringdashboard import config
        with session_scope() as db_session:
            result = get_monitor_rules(db_session)
            self.assertEqual(len(result), 1)
            self.assertEqual(result[0].endpoint, NAME)
            self.assertTrue(result[0].monitor)
            self.assertEqual(result[0].version_added, config.version)
            self.assertEqual(result[0].last_accessed, TIMES[0])

    def test_get_monitor_data(self):
        """
            Test whether the function returns the right values.
        """
        from flask_monitoringdashboard.database.monitor_rules import get_monitor_rules, get_monitor_data
        # since all monitor-rules in the test-database have the 'monitor'-variable set to True, the outcome of both
        # functions is equivalent
        with session_scope() as db_session:
            result1 = get_monitor_data(db_session)
            result2 = get_monitor_rules(db_session)
            self.assertEqual(len(result1), len(result2))
            self.assertEqual(result1[0].endpoint, result2[0].endpoint)
            self.assertEqual(result1[0].last_accessed, result2[0].last_accessed)
            self.assertEqual(result1[0].monitor, result2[0].monitor)
            self.assertEqual(result1[0].time_added, result2[0].time_added)
            self.assertEqual(result1[0].version_added, result2[0].version_added)
