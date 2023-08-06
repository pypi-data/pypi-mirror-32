=============================
dmango
=============================

.. image:: https://badge.fury.io/py/dmango.svg
    :target: https://badge.fury.io/py/dmango

.. image:: https://travis-ci.org/simkimsia/dmango.svg?branch=master
    :target: https://travis-ci.org/simkimsia/dmango

.. image:: https://codecov.io/gh/simkimsia/dmango/branch/master/graph/badge.svg
    :target: https://codecov.io/gh/simkimsia/dmango

Your project description goes here

Documentation
-------------

The full documentation is at https://dmango.readthedocs.io.

Quickstart
----------

Install dmango::

    pip install dmango

Add it to your `INSTALLED_APPS`:

.. code-block:: python

    INSTALLED_APPS = (
        ...
        'dmango.apps.DmangoConfig',
        ...
    )

Add dmango's URL patterns:

.. code-block:: python

    from dmango import urls as dmango_urls


    urlpatterns = [
        ...
        url(r'^', include(dmango_urls)),
        ...
    ]

Features
--------

* TODO

Running Tests
-------------

Does the code actually work?

::

    source <YOURVIRTUALENV>/bin/activate
    (myenv) $ pip install tox
    (myenv) $ tox

Credits
-------

Tools used in rendering this package:

*  Cookiecutter_
*  `cookiecutter-djangopackage`_

.. _Cookiecutter: https://github.com/audreyr/cookiecutter
.. _`cookiecutter-djangopackage`: https://github.com/pydanny/cookiecutter-djangopackage
