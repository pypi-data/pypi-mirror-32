.. image:: https://travis-ci.org/gocept/zodb.py3migrate.svg?branch=master
        :target: https://travis-ci.org/gocept/zodb.py3migrate.svg

.. image:: https://readthedocs.org/projects/zodbpy3migrate/badge/?version=latest
        :target: https://zodbpy3migrate.readthedocs.io
        :alt: Documentation Status

.. image:: https://img.shields.io/pypi/v/zodb.py3migrate.svg
        :target: https://pypi.org/project/zodb.py3migrate/
        :alt: PyPI

.. image:: https://img.shields.io/pypi/pyversions/zodb.py3migrate.svg
        :target: https://pypi.org/project/zodb.py3migrate/
        :alt: Python versions


===================================
zodb.py3migrate - ZODB and Python 3
===================================

If you have a ZODB_ database written using Python 2.x this package helps you to
get your database ready for using it with Python 3. It is able to:

* display objects which need to be converted,

* do most of the the conversion for you,

* and switch the database file to be readable in a process running Python 3.

This package is compatible with Python version 2.7.

The documentation is at https://zodbpy3migrate.readthedocs.io.

.. _ZODB : http://zodb.org


==========
Developing
==========

:Author:
    `gocept <http://gocept.com/>`_ <mail@gocept.com>

:Online documentation:
    http://zodbpy3migrate.readthedocs.io/

:PyPI page:
    https://pypi.python.org/pypi/zodb.py3migrate

:Issues:
    https://github.com/gocept/zodb.py3migrate/issues

:Source code:
    https://github.com/gocept/zodb.py3migrate

:Current change log:
    https://raw.githubusercontent.com/gocept/zodb.py3migrate/master/CHANGES.rst


==========
Change log
==========

0.6 (2018-06-05)
================

- Update requirements to ``ZODB >= 4`` as with an older version the migration
  cannot be done successfully.
  (`#13 <https://github.com/gocept/zodb.py3migrate/issues/13>`_)


0.5 (2017-01-17)
================

- Release as wheel and include all files in release.

- Ensure compatibility with ``setuptools >= 30``.


0.4 (2016-10-29)
================

- Fix brown bag release.


0.3 (2016-10-29)
================

- Fixes for issues #4 and #5: Converted ZODB ist now actually saved,
  using additional subtransactions improves the memory footprint.


0.2 (2016-06-08)
================

- Split up the two functions previously united in the script
  ``bin/zodb-py3migrate`` into ``bin/zodb-py3migrate-analyze`` resp.
  ``bin/zodb-py3migrate-convert``.

- Add new options to the analysis script:

  - ``--start`` to start the analysis with a predefined OID.

  - ``--limit`` to stop the analysis after a certain amount of seen OIDs.

0.1 (2016-05-19)
================

* Initial release.


