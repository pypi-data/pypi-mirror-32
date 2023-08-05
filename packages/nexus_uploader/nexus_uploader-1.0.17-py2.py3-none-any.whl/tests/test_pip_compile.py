from nexus_uploader.pip_compile import pip_compile
from nexus_uploader.requirements import read_requirements

try: # Python 3
    from unittest.mock import patch
except ImportError: # Python 2
    from mock import patch
from piptools.exceptions import NoCandidateFound
import pytest
from .pytest_markers import pypi_integration


@pypi_integration  # only ran with --run-pypi-integration-tests
def test_pipcompile_with_real_repository():
    reqs_gen = pip_compile(['nexus_uploader'],
                           nexus_hostname='dummy-nexus.wtf',
                           append_egg_hash_to_url_if_need_be=raise_if_called)
    requirements_read = read_requirements(reqs_gen)
    assert len(requirements_read) == 10


def test_pipcompile_with_mocked_repository(FakePypiRepository):
    with patch('nexus_uploader.pip_compile.PyPIRepository', FakePypiRepository):
        reqs_gen = pip_compile(['lib1'],
                               nexus_hostname='dummy-nexus.wtf',
                               append_egg_hash_to_url_if_need_be=raise_if_called)
        assert read_requirements(reqs_gen) == \
            [('version-locked', ('lib1', '1.2.3', '')), ('version-locked', ('lib2', '4.5.7', 'via lib1'))]


@pypi_integration  # only ran with --run-pypi-integration-tests
def test_pipcompile_conflicting_dependencies():
    with pytest.raises(NoCandidateFound) as excinfo:
        pip_compile(['nexus_uploader==1.0.0','nexus_uploader==1.0.1'],
                    nexus_hostname='dummy-nexus.wtf',
                    append_egg_hash_to_url_if_need_be=raise_if_called)
    assert 'Could not find a version that matches nexus_uploader==1.0.0,==1.0.1' in str(excinfo.value)


def raise_if_called(*args, **kwargs):
    raise NotImplementedError
