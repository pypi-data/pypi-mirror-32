mir.dlsite README
=================

.. image:: https://circleci.com/gh/darkfeline/mir.dlsite.svg?style=shield
   :target: https://circleci.com/gh/darkfeline/mir.dlsite
   :alt: CircleCI
.. image:: https://codecov.io/gh/darkfeline/mir.dlsite/branch/master/graph/badge.svg
   :target: https://codecov.io/gh/darkfeline/mir.dlsite
   :alt: Codecov
.. image:: https://badge.fury.io/py/mir.dlsite.svg
   :target: https://badge.fury.io/py/mir.dlsite
   :alt: PyPI Release
.. image:: https://readthedocs.org/projects/mir-dlsite/badge/?version=latest
   :target: http://mir-dlsite.readthedocs.io/en/latest/
   :alt: Latest Documentation

API for DLsite.

Before running any other make command, run::

  $ pipenv install --dev

To build an installable wheel, run::

  $ pipenv run make wheel

To build a source distribution, run::

  $ pipenv run make sdist

To run tests, run::

  $ pipenv run make check

To build docs, run::

  $ pipenv run make html

To build a TAGS file, run::

  $ make TAGS
