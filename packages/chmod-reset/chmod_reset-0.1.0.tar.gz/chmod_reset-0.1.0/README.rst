===============================
CHMOD reset
===============================
|Build| |Coverage| |PyPI| |Status| |Version| |Python| |License|

.. |Build| image:: https://travis-ci.org/seignovert/python-chmod-reset.svg?branch=master
        :target: https://travis-ci.org/seignovert/python-chmod-reset
.. |Coverage| image:: https://coveralls.io/repos/github/seignovert/python-chmod-reset/badge.svg?branch=master
        :target: https://coveralls.io/github/seignovert/python-chmod-reset?branch=master
.. |PyPI| image:: https://img.shields.io/badge/PyPI-chmod_reset-blue.svg
        :target: https://pypi.python.org/project/chmod_reset
.. |Status| image:: https://img.shields.io/pypi/status/chmod_reset.svg?label=Status
.. |Version| image:: https://img.shields.io/pypi/v/chmod_reset.svg?label=Version
.. |Python| image:: https://img.shields.io/pypi/pyversions/chmod_reset.svg?label=Python
.. |License| image:: https://img.shields.io/pypi/l/chmod_reset.svg?label=License

*Python package to reset chmod recursively*

Install
-------
With ``pip``:

.. code:: bash

    $ pip install chmod_reset

With the ``source files``:

.. code:: bash

    $ git clone https://github.com/seignovert/python-chmod-reset.git
    $ cd chmod_reset ; python setup.py install

Usage
------

.. code:: bash

    reset-chmod -h # For help
    reset-chmod [-q] <FOLDER_NAME> # [-q] `quiet` optional

.. code:: python

    >>> import chmod_reset

    >>> chmod_reset.scan('.')
