Flask-Multipass
===============

.. image:: https://raw.githubusercontent.com/indico/flask-multipass/master/artwork/flask-multipass.svg
    :width: 200px

.. image:: https://readthedocs.org/projects/flask-multipass/badge/?version=latest
    :target: https://flask-multipass.readthedocs.org/
.. image:: https://travis-ci.org/indico/flask-multipass.svg
    :target: https://travis-ci.org/indico/flask-multipass
.. image:: https://coveralls.io/repos/indico/flask-multipass/badge.svg
    :target: https://coveralls.io/r/indico/flask-multipass

Flask-Multipass provides Flask with a user authentication/identity
system which can use different backends (such as local users,
LDAP and OAuth) simultaneously.

It was developed at CERN and is currently used in production by `Indico <https://github.com/indico/indico>`_.

There are bult-in authentication and identity providers for:

 * `Static (hardcode) credentials <https://github.com/indico/flask-multipass/blob/master/flask_multipass/providers/static.py>`_
 * `Local (SQLAlchemy DB) authentication <https://github.com/indico/flask-multipass/blob/master/flask_multipass/providers/sqlalchemy.py>`_
 * `OAuth <https://github.com/indico/flask-multipass/blob/master/flask_multipass/providers/oauth.py>`_
 * `Shibboleth <https://github.com/indico/flask-multipass/blob/master/flask_multipass/providers/shibboleth.py>`_
 * `LDAP <https://github.com/indico/flask-multipass/blob/master/flask_multipass/providers/shibboleth.py>`_

Those can be used simultaneously and interchangeably (e.g. authenticate with OAuth and search users with LDAP).

Documentation is available at https://flask-multipass.readthedocs.org


