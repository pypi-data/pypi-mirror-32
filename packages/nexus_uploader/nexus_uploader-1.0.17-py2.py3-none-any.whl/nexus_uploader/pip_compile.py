#
# Copyright (C) 2016 VSCT
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# pylint: disable=deprecated-module

from tempfile import NamedTemporaryFile
try: # Python 3
    from tempfile import TemporaryDirectory
except ImportError: # Python 2
    from contextlib import contextmanager
    import shutil
    from tempfile import mkdtemp
    @contextmanager
    def TemporaryDirectory(): # pylint: disable=invalid-name
        tmp_dir_path = mkdtemp()
        try:
            yield tmp_dir_path
        finally:
            shutil.rmtree(tmp_dir_path)

import optparse, os, pip
from pip.download import unpack_url
from pip.req import parse_requirements
from pip.req.req_install import InstallRequirement
from pip.req.req_set import make_abstract_dist, Installed
from pip.download import PipSession
from piptools.scripts.compile import PipCommand
from piptools.repositories import LocalRequirementsRepository, PyPIRepository
from piptools.resolver import Resolver
from piptools.utils import assert_compatible_pip_version, key_from_ireq
from piptools.writer import OutputWriter

from .utils import aslist


def pip_compile(reqfile_lines, nexus_hostname, append_egg_hash_to_url_if_need_be):
    assert_compatible_pip_version()
    in_reqfile_lines = []
    for reqfile_line in reqfile_lines:
        if reqfile_line.startswith('http'):
            if ' #' in reqfile_line:
                reqfile_line = reqfile_line.partition(' #')[0]
            reqfile_line = append_egg_hash_to_url_if_need_be(reqfile_line)
        in_reqfile_lines.append(reqfile_line)
    constraints = _constraints_from_reqfile_lines(in_reqfile_lines)
    with TemporaryDirectory() as tmp_dir_path:
        return _pip_compile(constraints, nexus_hostname, append_egg_hash_to_url_if_need_be, tmp_dir_path)

@aslist
def _pip_compile(constraints, nexus_hostname, append_egg_hash_to_url_if_need_be, html_index_dir_path):
    dependency_links_requirements = {str(c.req).lower(): c for c in constraints if c.link}

    class CustomResolver(Resolver):
        @staticmethod
        def check_constraints(constraints):
            # Overrides Resolver._check_constraints
            # We allow non-editable URLs as packages
            pass

        def _iter_dependencies(self, ireq):
            # Overrides Resolver._iter_dependencies
            # We use our cached InstallRequirements that have a .link, instead of the new ones built by piptools.resolver.Resolver._iter_dependencies
            for dep_ireq in super(CustomResolver, self)._iter_dependencies(ireq):
                if str(dep_ireq.req) == 'setuptools':
                    continue
                if str(dep_ireq.req) in dependency_links_requirements:
                    yield dependency_links_requirements[str(dep_ireq.req)]
                else:
                    yield dep_ireq

    class CustomPyPIRepository(PyPIRepository):
        def get_dependencies(self, ireq):
            # Overrides PyPIRepository.get_dependencies
            dependencies = set(super(CustomPyPIRepository, self).get_dependencies(ireq))
            dist = self.get_dist(ireq)
            dep_links_ireqs = self.get_dep_links_ireqs(dist)
            return dependencies | dep_links_ireqs

        @staticmethod
        def get_dep_links_ireqs(dist):
            # pylint: disable=protected-access
            dependency_links = list(dist._get_metadata('dependency_links.txt'))
            dependency_links = [append_egg_hash_to_url_if_need_be(url) for url in dependency_links]
            dependency_links = [url for url in dependency_links if url]
            dep_links_ireqs = set(InstallRequirement.from_line(url) for url in dependency_links)
            # We cache those dependency_links so that they can be accessed from CustomResolver._iter_dependencies, self.find_all_candidates and at the end of _pip_compile
            for dep_ireq in dep_links_ireqs:
                dep_ireq.remove_temporary_source()
                dependency_links_requirements[str(dep_ireq.req)] = dep_ireq
                if str(dep_ireq.req).lower() != str(dep_ireq.req):
                    # This is required for dependencies with a version like X.Y.Z-SNAPSHOT (notice the uppercase).
                    # They need to exist BOTH as key in this dict, otherwise:
                    # * if only the lowercase is there, `get_dependencies` raise a pip.exceptions.DistributionNotFound
                    # * if only the uppercase is there, an erroneous "version-locked" pkg version gets out of this module, and we get a:
                    # requests.exceptions.HTTPError: 404 Client Error: Not Found (no releases) for url: https://pypi.python.org/pypi/$pkg/json
                    dependency_links_requirements[str(dep_ireq.req).lower()] = dep_ireq
            return dep_links_ireqs

        def get_dist(self, ireq):
            # Reproduce pip.req.req_set.RequirementSet._prepare_file code, called by RequirementSet.prepare_files
            ireq.check_if_exists()
            if ireq.satisfied_by is not None:
                abstract_dist = Installed(ireq)
            else:
                ireq.ensure_has_source_dir(self.source_dir)
                ireq.populate_link(self.finder, upgrade=False, require_hashes=False)
                assert ireq.link
                unpack_url(ireq.link, ireq.source_dir, self._download_dir, only_download=True, session=self.session)
                abstract_dist = make_abstract_dist(ireq)
                abstract_dist.prep_for_dist()
            return abstract_dist.dist(self.finder)

        def find_all_candidates(self, req_name):
            # Overrides PyPIRepository.find_all_candidates
            # We update our --find-links HTML file
            _write_find_links_html_filename(html_index_dir_path, [dep_ireq.link.url for dep_ireq in dependency_links_requirements.values()])
            return super(CustomPyPIRepository, self).find_all_candidates(req_name)

    # Reproduce piptools.scripts.compile.cli code
    pip_options, session = _get_pip_options(['--trusted-host', nexus_hostname, '--find-links', html_index_dir_path])
    repository = LocalRequirementsRepository(existing_pins=dict(), proxied_repository=CustomPyPIRepository(pip_options, session))
    resolver = CustomResolver(constraints, repository, clear_caches=True)
    results = resolver.resolve()

    out_reqfile_lines = _capture_annotated_out_reqlines(results, resolver=resolver, constraints=constraints, format_control=repository.finder.format_control)

    for requirement in out_reqfile_lines:
        eol_comment = ''
        if ' #' in requirement and not requirement.startswith('#'):
            requirement, eol_comment = requirement.split(' #')
        requirement, eol_comment = requirement.strip(), eol_comment.strip()
        # !warning! piptools OutwputWriter now lowercases versions
        if requirement.lower() in dependency_links_requirements:
            yield dependency_links_requirements[requirement.lower()].link.url + (' # ' + eol_comment if eol_comment else '')
        else:
            yield requirement + (' # ' + eol_comment if eol_comment else '')

def _constraints_from_reqfile_lines(reqfile_lines):
    with NamedTemporaryFile('w') as tmp_file:
        tmp_file.write('\n'.join(reqfile_lines))
        tmp_file.flush()
        return list(parse_requirements(tmp_file.name, session=PipSession()))

def _get_pip_options(pip_args):
    # Reproduce piptools.scripts.compile.cli code
    pip_command = PipCommand()
    pip.cmdoptions.make_option_group(pip.cmdoptions.index_group, pip_command.parser)
    pip_command.parser.add_option(optparse.Option('--pre', action='store_true', default=False))
    pip_options, _ = pip_command.parse_args(pip_args)
    # pylint: disable=protected-access
    session = pip_command._build_session(pip_options)
    return pip_options, session

def _write_find_links_html_filename(html_index_dir_path, find_links_urls):
    html_content = '\n'.join("<a href='{}'></a>".format(url) for url in find_links_urls)
    with open(os.path.join(html_index_dir_path, 'index.html'), 'w') as html_index_file:
        html_index_file.write(html_content)

def _capture_annotated_out_reqlines(results, resolver, constraints, format_control):
    with NamedTemporaryFile(delete=False) as tmp_file:
        tmp_file_name = tmp_file.name
    try:
        writer = OutputWriter(src_files=None, dst_file=tmp_file_name, dry_run=False,
                              emit_header=False, emit_index=False, emit_trusted_host=False,
                              annotate=True, generate_hashes=False, default_index_url=None, index_urls=None,
                              trusted_hosts=(), format_control=format_control)
        writer.write(results=results, reverse_dependencies=resolver.reverse_dependencies(results),
                     unsafe_requirements=resolver.unsafe_constraints,
                     primary_packages={key_from_ireq(ireq) for ireq in constraints},
                     markers={}, hashes=None)
        with open(tmp_file_name, 'r') as tmp_file:
            return tmp_file.readlines()
    finally:
        os.unlink(tmp_file_name)
