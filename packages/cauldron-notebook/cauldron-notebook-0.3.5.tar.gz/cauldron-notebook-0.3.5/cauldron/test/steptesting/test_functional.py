import os
from unittest.mock import MagicMock
from unittest.mock import patch

import pytest

import cauldron as cd
from cauldron import steptest
from cauldron.steptest import CauldronTest


@pytest.fixture(name='tester')
def tester_fixture() -> CauldronTest:
    """Create the Cauldron project test environment"""
    tester = CauldronTest(project_path=os.path.dirname(__file__))
    tester.setup()
    yield tester
    tester.tear_down()


def test_first_step(tester: CauldronTest):
    """Should not be any null/NaN values in df"""
    assert cd.shared.fetch('df') is None
    step = tester.run_step('S01-first.py')
    df = cd.shared.df
    assert not df.isnull().values.any()

    error_echo = step.echo_error()
    assert error_echo == ''


def test_second_step(tester: CauldronTest):
    """
    Should fail without exception because of an exception raised in the
    source but failure is allowed
    """
    step = tester.run_step('S02-errors.py', allow_failure=True)
    assert not step.success

    error_echo = step.echo_error()
    assert 0 < len(error_echo)


def test_second_step_strict(tester: CauldronTest):
    """
    Should fail because of an exception raised in the source when strict
    failure is enforced
    """
    with pytest.raises(Exception):
        tester.run_step('S02-errors.py', allow_failure=False)


@patch('_testlib.patching_test')
def test_second_step_with_patching(
        patching_test: MagicMock,
        tester: CauldronTest
):
    """Should override the return value with the patch"""
    patching_test.return_value = 12
    cd.shared.value = 42

    tester.run_step('S03-lib-patching.py')
    assert 12 == cd.shared.result


def test_second_step_without_patching(tester: CauldronTest):
    """Should succeed running the step without patching"""
    cd.shared.value = 42
    tester.run_step('S03-lib-patching.py')
    assert 42 == cd.shared.result


def test_to_strings(tester: CauldronTest):
    """Should convert list of integers to a list of strings"""
    before = [1, 2, 3]
    step = tester.run_step('S01-first.py')
    after = step.local.to_strings(before)
    assert step.success
    assert ['1', '2', '3'] == after


def test_modes(tester: CauldronTest):
    """Should be testing and not interactive or single run"""
    step = tester.run_step('S01-first.py')
    assert step.success
    assert step.local.is_testing
    assert not step.local.is_interactive
    assert not step.local.is_single_run


def test_find_in_current_path():
    """Should find a project in this file's directory"""
    directory = os.path.dirname(os.path.realpath(__file__))
    result = steptest.find_project_directory(directory)
    assert directory == result


def test_find_in_parent_path():
    """Should find a project in the parent directory"""
    directory = os.path.dirname(os.path.realpath(__file__))
    subdirectory = os.path.join(directory, 'fake')
    result = steptest.find_project_directory(subdirectory)
    assert directory == result


def test_find_in_grandparent_path():
    """Should find a project in the grandparent directory"""
    directory = os.path.dirname(os.path.realpath(__file__))
    subdirectory = os.path.join(directory, 'fake', 'fake')
    result = steptest.find_project_directory(subdirectory)
    assert directory == result


def test_find_failed_at_root():
    """Should raise FileNotFoundError if top-level directory has no project"""
    directory = os.path.dirname(os.path.realpath(__file__))
    subdirectory = os.path.join(directory, 'fake')

    with patch('os.path.dirname', return_value=subdirectory) as func:
        with pytest.raises(FileNotFoundError):
            steptest.find_project_directory(subdirectory)
        func.assert_called_once_with(subdirectory)


def test_make_temp_path(tester: CauldronTest):
    """Should make a temp path for testing"""
    temp_path = tester.make_temp_path('some-id', 'a', 'b.test')
    assert temp_path.endswith('b.test')


def test_no_such_step(tester: CauldronTest):
    """Should fail if no such step exists"""
    with pytest.raises(Exception):
        tester.run_step('FAKE-STEP.no-exists')


def test_no_such_project(tester: CauldronTest):
    """Should fail if no project exists"""
    project = cd.project.internal_project
    cd.project.load(None)

    with pytest.raises(Exception):
        tester.run_step('FAKE')

    cd.project.load(project)


def test_open_project_fails(tester: CauldronTest):
    """Should raise Assertion error after failing to open the project"""
    with patch('cauldron.steptest.support.open_project') as open_project:
        open_project.side_effect = RuntimeError('FAKE')
        with pytest.raises(AssertionError):
            tester.open_project()
