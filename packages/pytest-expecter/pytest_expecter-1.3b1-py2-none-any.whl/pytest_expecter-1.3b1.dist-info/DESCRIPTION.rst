pytest-expecter
===============

A ``pytest`` plugin for
`garybernhardt/expecter <https://github.com/garybernhardt/expecter>`__
that hides the internal stacktrace.

| |Build Status|
| |PyPI Version|

Overview
--------

This lets you write tests (optionally using
`ropez/pytest-describe <https://github.com/ropez/pytest-describe>`__)
like this:

.. code:: python

    def describe_foobar():

        def it_can_pass(expect):
            expect(2 + 3) == 5

        def it_can_fail(expect):
            expect(2 + 3) == 6

and get output like this:

.. code:: sh

    =================================== FAILURES ===================================
    _________________________ describe_foobar.it_can_fail __________________________

        def it_can_fail(expect):
    >       expect(2 + 3) == 6
    E       AssertionError: Expected 6 but got 5

    test_foobar.py:7: AssertionError
    ====================== 1 failed, 1 passed in 2.67 seconds ======================

Installation
------------

.. code:: sh

    pip install pytest-expecter

.. |Build Status| image:: http://img.shields.io/travis/jacebrowning/pytest-expecter/plugin.svg
   :target: https://travis-ci.org/jacebrowning/pytest-expecter
.. |PyPI Version| image:: http://img.shields.io/pypi/v/pytest-expecter.svg
   :target: https://pypi.python.org/pypi/pytest-expecter

Revision History
================

1.3 (unreleased)
----------------

-  Added ``startswith``, ``endswith``, ``istartswith``, and
   ``iendswith`` helpers.

1.2 (2018/03/17)
----------------

-  Added ``icontains`` method to check for containment ignoring case.
-  Added ``iexcludes`` method to check for exclusion ignoring case.

1.1 (2018/02/21)
----------------

-  Added ``expect`` fixture to use directly in tests.

1.0 (2017/12/03)
----------------

-  Initial stable release.

0.2.2.post7 (2017/12/02)
------------------------

-  Added automatic conversion from ``OrderedDict`` to ``dict`` on Python
   3.6 to create readable diffs.


