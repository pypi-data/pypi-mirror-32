# pylint: disable=unused-import,redefined-outer-name
"""Unit test fixtures"""
from os.path import join, dirname

import pytest

TEST_DATA_FILE = join(dirname(__file__), 'test_data.json')


@pytest.fixture
def test_data():
    """yield a temporary test cache file"""
    import tempfile
    import os
    test_data = None
    with open(TEST_DATA_FILE, 'r') as test_data_file_obj:
        test_data = test_data_file_obj.read()

    file_handle, test_file = tempfile.mkstemp()
    with open(test_file, 'w') as test_file_obj:
        test_file_obj.write(test_data)

    yield test_file

    os.close(file_handle)
    os.remove(test_file)


@pytest.fixture
def bkend(test_data):
    """create a backend with test data"""
    from reentry.jsonbackend import JsonBackend
    test_file = test_data
    return JsonBackend(datafile=test_file)


@pytest.fixture
def manager(bkend):
    from reentry.default_manager import PluginManager
    manager = PluginManager(backend=bkend)
    yield manager


@pytest.fixture
def noscan_manager(bkend):
    from reentry.default_manager import PluginManager
    manager = PluginManager(backend=bkend, scan_for_not_found=False)
    yield manager
