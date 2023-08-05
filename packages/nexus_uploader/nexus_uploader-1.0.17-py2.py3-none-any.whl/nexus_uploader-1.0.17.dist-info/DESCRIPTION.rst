|pypi_version_img| |pypi_license_img| |travis_build_status|

Python tools to help with the development & deployment of
company-private Python packages.

It was developped to use Sonatype Nexus as a private Pypi mirror (until
`it supports this
natively <https://issues.sonatype.org/browse/NEXUS-6037>`__), but should
be adaptable to any repository supporting artifact upload with HTTP.

It is composed of 2 distincts tools :

-  ``nexus_uploader`` : a simple module to be used as a
   ``setup_requires`` entry in ``setup.py``, in order to easily upload
   Python packages onto a Nexus.
-  ``pyRequirements2nexus`` : a CLI tool to convert standard Python
   ``requirements.txt`` into ``nexus-requirements.txt`` files, made of
   URLs pointing to installable packages that ``pyRequirements2nexus``
   mirrored on your Nexus.

Table of Contents
================

.. contents::

Workflow diagram
================

.. figure:: https://raw.githubusercontent.com/voyages-sncf-technologies/nexus_uploader/master/docs/PythonPackaging.png
   :alt:

Features
========

- fully compatible Python 2 & 3

nexus\_uploader features
------------------------

-  easy integration in ``setup.py``
-  HTTP BasicAuth to connect to Sonatype Nexus

pyRequirements2nexus features
-----------------------------

-  full dependency resolution of Python packages, working with both Pypi
   public ones & private ones on a Nexus
-  to install packages, the end machine only require a connexion to the
   Nexus host, not to the Internet
-  support ``-e`` editable packages in ``requirements.txt``
-  support package URL fallbacks as comments in ``requirements.txt``
-  can select only Python 2 packages with `--allowed-pkg-classifiers py2.py3-none-any,py2-none-any`
-  a list of packages already included in the base environment can be
   specified (e.g. if you are using Anaconda), so that they will be
   excluded from the final ``nexus-requirements.txt``

Limitations

-  only support ``==`` version locking in ``requirements.txt`` (not
   ``>=``). This is intentional, to ensure package versions do not
   change unexpectedly.

How to upload home-made Python packages to my Nexus ?
=====================================================

Here is a handy recipe you can put in your ``setup.py`` :

::

    setup(
        ...
        setup_requires=['nexus_uploader']
    )

Usage:

::

    $ python setup.py sdist nexus_upload --repository http://nexus/content/repositories/Snapshots/com/vsct/my_project --username $REPOSITORY_USER --password $REPOSITORY_PASSWORD

Important note
--------------

Contrary to ``requirements.txt`` files, there is no way to include URLs
to dependencies in ``setup.py`` ``install_requires`` (nor
``setup_requires``). To work around this limitation,
``pyrequirements2nexus`` takes advantage of the deprecated
"dependency\_links" mechanism. Using them, it is able to discover the
full chain of Python packages dependencies, both from Pypi **AND** from
your local Nexus.

To easily generate the "dependency\_links.txt" in your Nexus-hosted
private package out of your ``requirements.txt``, use the following
recipe in your its ``setup.py`` :

::

    with open(os.path.join(os.path.dirname(os.path.realpath(__file__)), 'requirements.txt')) as requirements_file:
        requirements = requirements_file.readlines()
        dependency_links = [req for req in requirements if req.startswith('http')]
        install_requires = [req for req in requirements if not req.startswith('http')]

    setup(
        ...
        install_requires=install_requires,
        dependency_links=dependency_links,
    )

Because "dependency\_links" are not supported since `pip
1.6 <https://github.com/pypa/pip/pull/1519/commits/95ac4c16f544dcc4282d2a4245aba0384f7e629a>`__,
they will **NOT** be installed by ``pip install`` normal recursive
dependency-retrieval algorithm. You should always use the flat
``nexus-requirements.txt`` with ``pip install``.

pyRequirements2nexus usage
==========================

::

    pip install nexus_uploader
    pyRequirements2nexus --help

Also take a look at ``jenkins-install-python-requirements.sh`` for an
example of how we use it on our Jenkins.

Installation of nexus-requirements.txt on an end machine
========================================================

::

    pip install --user --no-index --no-deps --no-cache-dir --upgrade --requirement nexus-requirements.txt

Supported requirements.txt format
=================================

::

    http://nexus/content/repositories/repo_id/my/project/group/mypkgname/1.0/mypkgname-1.0-py2.py3-none-any.tar.gz
    nose==1.3.7   # -> transformed into an URL like this: http://nexus/content/repositories/repo_id/my/project/group/...
    -e ../my/relative/path  # http://nexus/content/repositories/...fallback_url...

Contributing
============

`pre-commit hooks <http://pre-commit.com>`__ installation:

::

    pip install -r dev-requirements
    pre-commit install

Unit tests:

::

    py.test tests/

Smoke tests using Pypi:

::

    ipython3 --pdb tests/smoke_test_extract_classifier_and_extension.py 200

FAQ
===

pip install - Download error on https://pypi.python.org / Couldn't find index page for
--------------------------------------------------------------------------------------

The stack trace :

::

    Collecting http://nexus/content/repositories/pip/com/vsct/pip/jsonschema/2.5.1/jsonschema-2.5.1-py2.py3-none-any.tar.gz (from -r scripts/requirements.pip (line 12))
      Downloading http://nexus/content/repositories/pip/com/vsct/pip/jsonschema/2.5.1/jsonschema-2.5.1-py2.py3-none-any.tar.gz (50kB)
        Complete output from command python setup.py egg_info:
        Download error on https://pypi.python.org/simple/vcversioner/: [Errno -2] Name or service not known -- Some packages may not be found!
        Couldn't find index page for 'vcversioner' (maybe misspelled?)
        Download error on https://pypi.python.org/simple/: [Errno -2] Name or service not known -- Some packages may not be found!
        No local packages or download links found for vcversioner

Explanation : https://github.com/Julian/jsonschema/issues/276

Solution :

::

    $ cat <<EOF > ~/.pydistutils.cfg
    [easy_install]
    allow_hosts = nexus
    find_links = http://nexus/content/repositories/pip/com/vsct/pip/vcversioner/2.14.0.0/
    EOF

How to generate a "--default-packages" file out of an Anaconda .sh installer
----------------------------------------------------------------------------

::

    grep -aF 'extract_dist ' Anaconda3-2.4.1-Linux-x86_64.sh \
        | perl -p -e 's/extract_dist (.+?[0-9])[^.]*$/\1\n/;' -e 's/^(.+)-(.+)$/\1 == \2/;' \
        | grep -vE '^(_cache|_license|anaconda|python) ' > anaconda3-2.4.1_included_packages.txt

Tip for easily removing packages from your nexus
------------------------------------------------

::

    pip install --user repositorytools
    export REPOSITORY_USER=...
    export REPOSITORY_PASSWORD=
    artifact delete http://nexus/content/repositories/pip/com/vsct/pip/ultrajson/1.35/ultrajson-1.35-macosx-10.6-intel.tar.gz

ToDo
====

-  detect if package releases on Pypi require gcc compilation (are they
   using setuptools/distutils ``Extension`` in ``setup.py`` ?)
-  classifier-based selection of Python packages
-  add support for md5 & sha1 upload/checks

.. |pypi_version_img| image:: https://img.shields.io/pypi/v/nexus_uploader.svg?style=flat
   :target: https://pypi.python.org/pypi/nexus_uploader
.. |pypi_license_img| image:: https://img.shields.io/pypi/l/nexus_uploader.svg?style=flat
   :target: https://pypi.python.org/pypi/nexus_uploader
.. |travis_build_status| image:: https://travis-ci.org/voyages-sncf-technologies/nexus_uploader.svg?branch=master
    :target: https://travis-ci.org/voyages-sncf-technologies/nexus_uploader


Change Log
==========
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

[1.0.17] - 2018-05-16
---------------------
Fixed
~~~~~
- Using HTTPS url to connect to Pypi by default to avoid "HTTPError: 403 Client Error: SSL is required"

[1.0.16] - 2018-04-09
---------------------
Fixed
~~~~~
- support for X.Y.Z-SNAPSHOT (uppercase) package versions

[1.0.15] - 2018-02-21
---------------------
Added
~~~~~
- Now supporting unlocked versions using `>=` / `<=` / `>` / `<` specifiers

[1.0.14] - 2017-07-12
---------------------
Fixed
~~~~~
- Tuple unpacking issue on py3.5 for fixed-version requirements

[1.0.13] - 2017-07-03
---------------------
Fixed
~~~~~
- The resolver now handles properly uppercase versions (like X.Y-SNAPSHOT)

[1.0.12] - 2017-07-03
---------------------
Fixed
~~~~~
- `Changelog.md` wasn't included in tarball (due to `MANIFEST.in`)

[1.0.10] - 2017-07-03
---------------------
Fixed
~~~~~
- a bug introduced in v1.0.9: `piptools.exceptions.UnsupportedConstraint: pip-compile does not support URLs as packages, unless they are editable`
- changelog rendering in Pypi rst description

[1.0.9] - 2017-06-30
--------------------
Fixed
~~~~~
- Upgrading pip-tools dependency to 1.9.0 to include commit b8043be (support for pip 8.1.2) -> https://github.com/voyages-sncf-technologies/nexus_uploader/commit/38f031b
- Fixing support for classifiers to be able to only retrieve Python2-compatible pkgs -> https://github.com/voyages-sncf-technologies/nexus_uploader/commit/35e78fc
- Fixing repository files permissions -> https://github.com/voyages-sncf-technologies/nexus_uploader/commit/38f031b

Added
~~~~~
- Made the lib compatible for Python 2.7 -> https://github.com/voyages-sncf-technologies/nexus_uploader/commit/5833c5f
- More tests

[1.0.8] - 2016-12-13
--------------------
Added
~~~~~
- `--pypi-json-api-url` parameter : source of Python packages to feed the Nexus with
- `--allowed-pkg-classifiers` parameter : when no `source` release is available for a given `package==version`,
  select a package matching on of those classifiers as a 2nd choice
  (introduced to handle `docutils==0.13.1` that only exist as a wheel).

Changed
~~~~~~~
- The `--allowed-pkg-classifiers` comes with a default value of `py3-none-any` changing the existing behaviour.
  To keep the previous behaviour you'll need to pass an empty value to this parameter.
  An `PypiQueryError` will then be raised if no `source` package is available.

