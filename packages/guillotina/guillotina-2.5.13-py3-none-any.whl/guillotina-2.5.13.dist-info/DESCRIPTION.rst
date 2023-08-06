Introduction
============

.. image:: https://img.shields.io/badge/docs-latest-brightgreen.svg?style=flat
   :target: http://guillotina.readthedocs.io/en/latest/

.. image:: https://travis-ci.org/plone/guillotina.svg?branch=master
   :target: https://travis-ci.org/plone/guillotina

.. image:: https://coveralls.io/repos/github/plone/guillotina/badge.svg?branch=master
   :target: https://coveralls.io/github/plone/guillotina?branch=master
   :alt: Test Coverage

.. image:: https://img.shields.io/pypi/pyversions/guillotina.svg
   :target: https://pypi.python.org/pypi/guillotina/
   :alt: Python Versions

.. image:: https://img.shields.io/pypi/v/guillotina.svg
   :target: https://pypi.python.org/pypi/guillotina

.. image:: https://img.shields.io/pypi/l/guillotina.svg
   :target: https://pypi.python.org/pypi/guillotina/
   :alt: License

.. image:: https://badges.gitter.im/plone/guillotina.png
   :target: https://gitter.im/plone/guillotina
   :alt: Chat

Please `read the detailed docs <http://guillotina.readthedocs.io/en/latest/>`_


This is the working project of the next generation Guillotina server based on asyncio.


Dependencies
------------

* python >= 3.6
* postgresql >= 9.6


Quickstart
----------

We use pip::

  pip install guillotina


Run postgresql
--------------

If you don't have a postgresql server to play with, you can run one easily
with docker.

Download and start the docker container by running::

  make run-postgres



Run the server
--------------

To run the server::

    g


Then...

    curl http://localhost:8080


Or, better yet, use postman to start playing with API.


Getting started with development
--------------------------------

Using pip::

  ./bin/pip install requirements.txt
  ./bin/pip install -e .[test]


Run tests
---------

We're using pytest::

    ./bin/pytest guillotina

and for test coverage::

    ./bin/pytest --cov=guillotina guillotina/

With file watcher...

    ./bin/ptw guillotina --runner=./bin/py.test


To run tests with cockroach db::

   USE_COCKROACH=true ./bin/pytest guillotina

Default
-------

Default root access can be done with AUTHORIZATION header : Basic root:root


Docker
------

You can also run Guillotina with Docker!


First, run postgresql::

    docker run --rm \
        -e POSTGRES_DB=guillotina \
        -e POSTGRES_USER=guillotina \
        -p 127.0.0.1:5432:5432 \
        --name postgres \
        postgres:9.6

Then, run guillotina::

    docker run --rm -it \
        --link=postgres -p 127.0.0.1:8080:8080 \
        guillotina/guillotina:latest \
        g -c '{"databases": [{"db": {"storage": "postgresql", "dsn": "postgres://guillotina:@postgres/guillotina"}}], "root_user": {"password": "root"}}'


This assumes you have a config.yaml in your current working directory


Chat
----

Join us to talk about Guillotina at https://gitter.im/plone/guillotina


2.5.13 (2018-06-08)
-------------------

- Fix use of AllowSingle when overriding permissions
  [vangheem]


2.5.12 (2018-04-02)
-------------------

- Do not swallow any exceptions on commit.
  [vangheem]


2.5.11 (2018-03-28)
-------------------

- Fix error loading settings
  [vangheem]


2.5.10 (2018-03-26)
-------------------

- Make sure to clear commit hook on tcp begin
  [vangheem]

- Add save method to upload data manager so saving data can be defered to
  after commit hook
  [vangheem]


2.5.9 (2018-03-22)
------------------

- return 404 when attempting to download file that is missing
  [vangheem]


2.5.8 (2018-03-22)
------------------

- Fix getting filename for file downloads
  [vangheem]


2.5.7 (2018-03-21)
------------------

- Only do commit and voting if we have objects to do it with
  [vangheem]


2.5.6 (2018-03-21)
------------------

- Make sure to set size when deserializing files from base64
  [vangheem]


2.5.5 (2018-03-20)
------------------

- Fix TUS upload with zero length files
  [vangheem]


2.5.4 (2018-03-19)
------------------

- `save_file` can provide size value
  [vangheem]


2.5.3 (2018-03-19)
------------------

- IFile.content_type should be string, not bytes
  [vangheem]

- `UploadDataManager.finish()` should return the file object created
  [vangheem]


2.5.2 (2018-03-19)
------------------

- Fix `@duplicate` endpoint when no destination is provided
  [vangheem]


2.5.1 (2018-03-19)
------------------

- Be able to not automatically serialize behaviors
  [vangheem]


2.5.0 (2018-03-19)
------------------

- normalize file manager api so we can have more simple integrations with s3/gcloud
  [vangheem]


2.4.7 (2018-03-17)
------------------

- Be able to safely PATCH with same payload from GET
  [vangheem]


2.4.6 (2018-03-17)
------------------

- Updated docs
  [vangheem]


2.4.5 (2018-03-15)
------------------

- `BucketListValue.iter_buckets` returns annotation object
  [vangheem]


2.4.4 (2018-03-15)
------------------

- fix `BucketListValue.iter_buckets` to correctly load uncached annotations
  from database
  [vangheem]


2.4.3 (2018-03-14)
------------------

- New `PatchField`
  [vangheem]

- New `BucketListField`
  [vangheem]

...

You are seeing a truncated changelog.

You can read the `changelog file <https://github.com/plone/guillotina/blob/master/CHANGELOG.rst>`_
for a complete list.



