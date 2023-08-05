==========
PyForecast
==========


.. image:: https://img.shields.io/pypi/v/pyforecast.svg
        :target: https://pypi.python.org/pypi/pyforecast

.. image:: https://img.shields.io/travis/vafliik/pyforecast.svg
        :target: https://travis-ci.org/vafliik/pyforecast

.. image:: https://readthedocs.org/projects/pyforecast/badge/?version=latest
        :target: https://pyforecast.readthedocs.io/en/latest/?badge=latest
        :alt: Documentation Status




Python binding for Forecast Harverst API - **Work In Progress**


* Free software: Apache Software License 2.0
* Documentation: https://pyforecast.readthedocs.io.


Features
--------

Implemented
^^^^^^^^^^^
- Client
- Project
- Person

TO DO
^^^^^^^^^^^
- Who am I
- Assignments
- Milestones
- User Connections
- Placeholders
- All PUT/POST requests (inserting/modifying data)

=====
Usage
=====

Create an Authorization token in Forecast App: https://id.getharvest.com/developers

Create an instance of the ``forecast.Api`` using the Account ID and Authorization token::

    >>> import forecast
    >>> api = forecast.Api(account_id='account_id', authorization_token='authorization_token')


To fetch a single user's public status messages, where ``user`` is a Twitter user's screen name::

    >>> for project in api.get_projects():
    >>>    print(project.name, project.id)
    Demo Project 101234
    Killer App 106555

    >>> person = api.get_person(42)
    >>> print(person.first_name, person.last_name, person.email)
    Pavel Pribyl pribyl.pavel@gmail.com

Credits
-------

This package was created with Cookiecutter_ and the `audreyr/cookiecutter-pypackage`_ project template.

.. _Cookiecutter: https://github.com/audreyr/cookiecutter
.. _`audreyr/cookiecutter-pypackage`: https://github.com/audreyr/cookiecutter-pypackage
