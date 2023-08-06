=============================
Importanize (import organize)
=============================

.. image:: https://badge.fury.io/py/importanize.png
    :target: http://badge.fury.io/py/importanize

.. image:: https://travis-ci.org/miki725/importanize.png?branch=master
    :target: https://travis-ci.org/miki725/importanize

.. image:: https://coveralls.io/repos/miki725/importanize/badge.png?branch=master
    :target: https://coveralls.io/r/miki725/importanize?branch=master

Utility for organizing Python imports using PEP8 or custom rules

* Free software: MIT license
* GitHub: https://github.com/miki725/importanize

Installing
----------

You can install ``importanize`` using pip::

    $ pip install importanize

Why?
----

I think imports are important in Python. I also think PEP8 is awesome
(if you disagree, read some PHP) and there are many tools to help
developers reformat code to match PEP8. There are however fewer tools
for organizing imports either by following PEP8 or custom rules.
There is `isort <http://isort.readthedocs.org/en/latest/>`_
(which unfortunately I found out about after writing this lib)
however it seems to do lots of magic to determine which packages
are 3rd party, local packages, etc. I wanted the imports configuration
to be simple and explicit.
This is where ``importanize`` comes in. It allows to organize
Python imports using PEP8 or your custom rules. Read on for
more information.

Using
-----

Using ``importanize`` is super easy. Just run::

    $ importanize file_to_organize.py

That will re-format all imports in the given file.
As part of the default configuration, ``importanize`` will try
it's best to organize imports to follow PEP8 however that is a rather
challenging task, since it is difficult to determine all import groups
as suggested by `PEP8 <http://legacy.python.org/dev/peps/pep-0008/#imports>`_:

1) standard library imports
2) related third party imports
3) local application/library specific imports

To help ``importanize`` distinguish between different import groups in most
cases it would be recommended to use custom config file::

    $ importanize file_to_organize.py config.json

Config file is simply a ``json`` file like this::

    {
        "formatter": "grouped",
        "groups": [
            {
                "type": "stdlib"
            },
            {
                "type": "sitepackages"
            },
            {
                "type": "remainder"
            },
            {
                "type": "packages",
                "packages": [
                    "my_favorite_package"
                ]
            },
            {
                "type": "local"
            }
        ]
    }

Default config looks something like::

    {
        "groups": [
            {
                "type": "stdlib"
            },
            {
                "type": "sitepackages"
            },
            {
                "type": "remainder"
            },
            {
                "type": "local"
            }
        ]
    }

Currently the only required key is ``"groups"`` which must be an array
of group definitions. ``importanize`` will use these group definitions
to organize imports and will output import groups in the same order
as defined in the config file. These are the supported group types:

* ``stdlib`` - standard library imports including ``__future__``
* ``sitepackages`` - imports coming from the ``site-packages`` directory
* ``local`` - local imports which start with ``"."``. for example
  ``from .foo import bar``
* ``packages`` - if this group is specified, additional key ``packages``
  is required within import group definition which should list
  all Python packages (root level) which should be included in that group::

      {
          "type": "packages",
          "packages": ["foo", "bar"]
      }

* ``remaining`` - all remaining imports which did not satisfy requirements
  of all other groups will go to this group.

You can use the config file by specifying it in the ``importanize``
command as shown above however you can also create an ``.importanizerc``
file and commit that to your repository. As a matter of fact,
you can see the
`.importanizerc <https://github.com/miki725/importanize/blob/master/.importanizerc>`_
config file used for the importanize repository itself.
Additionally multiple configurations are supported within a single repository
via sub-configurations. Simply place ``.importanizerc`` within a sub-folder
and all imports will be reconfigured under that folder.

You can also choose the formatter used to organize long multiline imports.
Currently, there are two formatters available:

* ``grouped`` (default)
* ``inline-grouped``

It can be set using the formatter config value, or the formatter option, for
example::

    $ importanize --formatter=inline-group --print tests/test_data/input.txt


Finally, you can see all other available ``importanize`` cli options::

    $ importanize --help

Not all configurations can be provided via cli.
Additional available configurations in configuration file:

* ``length`` - line length after which the formatter will split imports
* ``exclude`` - list of glob patterns of files which should be excluded from organizing.
  For example::

        "exclude": [
            "path/to/file",
            "path/to/files/ignore_*.py"
        ]

* ``after_imports_new_lines`` - number of lines to be included after imports
* ``add_imports`` - list of imports to add to every file.
  For example::

        "add_imports": [
            "from __future__ import absolute_import, print_function, unicode_literals"
        ]

It integrates with pre-commit_. You can use the following config

::

    repos:
    - repo: https://github.com/miki725/importanize/
      rev: 'master'
      hooks:
      - id: importanize
        args: [--verbose]

Example
-------

Here is a before and after using the default formatter(on hypothetical file):

Before
~~~~~~

::

    from __future__ import unicode_literals, print_function
    import os.path as ospath
    import datetime
    from package.subpackage.module.submodule import CONSTANT, Klass, foo, bar, rainbows
    from .module import foo, bar
    from ..othermodule import rainbows

After
~~~~~

::

    from __future__ import print_function, unicode_literals
    import datetime
    from os import path as ospath

    from package.subpackage.module.submodule import (
        CONSTANT,
        Klass,
        bar,
        foo,
        rainbows,
    )

    from ..othermodule import rainbows
    from .module import bar, foo

Here is what ``importanize`` did:

* alphabetical sort, even inside import line (look at ``__future__``)
* normalized ``import .. as ..`` into ``from .. import .. as ..``
* broke long import (>80 chars) which has more than one import
  into multiple lines
* reordered some imports (e.g. local imports ``..`` should be before ``.``)

Testing
-------

To run the tests you need to install testing requirements first::

    $ make install

Then to run tests, you can use ``nosetests`` or simply use Makefile command::

    $ nosetests -sv
    # or
    $ make test

.. _pre-commit: https://pre-commit.com/




History
-------

0.7.0 (2018-06-06)
~~~~~~~~~~~~~~~~~~

* Fixed removing first line in files without imports.
* Added ``--list`` option to list all found imports grouped by same packages as in config.

0.6.4 (2018-05-29)
~~~~~~~~~~~~~~~~~~

* Added support for custom line length.

0.6.3 (2018-01-27)
~~~~~~~~~~~~~~~~~~

* Fixed (again) importanize hanging when provided relative file path when finding sub-configurations.

0.6.2 (2018-01-08)
~~~~~~~~~~~~~~~~~~

* Fixed importanize failing on empty files.
  Thanks `Milind <https://github.com/milin>`_.
* Fixed importanize hanging when provided relative file path when finding sub-configurations.
  Thanks `Milind <https://github.com/milin>`_.

0.6.1 (2017-10-06)
~~~~~~~~~~~~~~~~~~

* Fixed sub-configrations. They are searched when organizing individual files now.

0.6.0 (2017-10-06)
~~~~~~~~~~~~~~~~~~

* Added support for sub-configurations when ``.importanize`` is found.
* Added support for ``add_imports`` in configuration.

0.5.3 (2017-06-06)
~~~~~~~~~~~~~~~~~~

* Added support to customize number of new lines added after imports
  via ``after_imports_new_lines`` configuration.
  Useful when using auto formatters such as ``yapf``.

0.5.2 (2017-05-18)
~~~~~~~~~~~~~~~~~~

* Skipping directories which makes skipping subfolders much faster
* Fixed bug which incorrectly skipped files

0.5.1 (2017-05-09)
~~~~~~~~~~~~~~~~~~

* Fixed bug which incorrectly removed duplicate leafless imports which had different ``as`` names

0.5.0 (2017-05-03)
~~~~~~~~~~~~~~~~~~

* Added ``--ci`` flag to validate import organization in files
* Added ``sitepackages`` import group. Thanks `Pamela <https://github.com/PamelaM>`_.
  See ``README`` for more info
* Added pipe handling (e.g. ``cat foo.py | importanize``)
* Fixed bug which incorrectly sorted imports with aliases (e.g. ``import foo as bar``)
* Files are not overridden when imports are already organized.
  Useful in precommit hooks which detect changed files.
* Released as Python `wheel <http://pythonwheels.com/>`_

0.4.1 (2015-07-28)
~~~~~~~~~~~~~~~~~~

* Fixed a bug where ``importanize`` did not correctly detect stdlibs on Windows
  (see `#29 <https://github.com/miki725/importanize/issues/29/>`_)
* Removed ``future`` dependency since ``six>=1.9`` includes all the used features
* Fixed tests to be executable on Windows

0.4 (2015-04-13)
~~~~~~~~~~~~~~~~

* Added multiple formatter options. Can be used using ``--formatter``
  flag or can be set in the configuration file.
* Fixes a bug in parsing imports when encountering both ``\`` and ``()``
  (see `#26 <https://github.com/miki725/importanize/issues/26>`_ for example)
* Fixes a bug where wildcard leaf imports were combined with other others
  (see `#25 <https://github.com/miki725/importanize/issues/25/>`_ for example)

0.3 (2015-01-18)
~~~~~~~~~~~~~~~~

* Using tokens to parse Python files. As a result this allows to
  fix how comments are handled
  (see `#21 <https://github.com/miki725/importanize/issues/21>`_ for example)

0.2 (2014-10-30)
~~~~~~~~~~~~~~~~

* New "exclude" config which allows to skip files
* Presetving origin file new line characters
* Traversing parent paths to find importanize config file

0.1.4 (2014-10-12)
~~~~~~~~~~~~~~~~~~

* Multiple imports (e.g. ``import a, b``) are normalized
  instead of exiting
* Multiple imports with the same stem are combined into
  single import statement
  (see `#17 <https://github.com/miki725/importanize/issues/17>`_ for example)

0.1.3 (2014-09-15)
~~~~~~~~~~~~~~~~~~

* Fixed where single line triple-quote docstrings would cause
  none of the imports to be recognized

0.1.2 (2014-09-15)
~~~~~~~~~~~~~~~~~~

* Fixed where import leafs were not properly sorted for
  mixed case (aka CamelCase)

0.1.1 (2014-09-07)
~~~~~~~~~~~~~~~~~~

* Ignoring comment blocks when parsing for imports
* Fixed bug when imports start on a first line,
  extra lines were being added to the file.

0.1.0 (2014-09-07)
~~~~~~~~~~~~~~~~~~

* First release on PyPI.


Credits
-------

Development Lead
~~~~~~~~~~~~~~~~

* Miroslav Shubernetskiy  - https://github.com/miki725

Contributors
~~~~~~~~~~~~

* Benjamin Abel  - https://github.com/benjaminabel
* Pamela McA'Nulty - https://github.com/PamelaM
* Milind Shakya - https://github.com/milin
* Serkan Hosca - https://github.com/shosca


License
-------

The MIT License (MIT)

Copyright (c) 2014, Miroslav Shubernetskiy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.


