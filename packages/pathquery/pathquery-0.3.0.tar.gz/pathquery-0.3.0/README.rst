PathQuery
=========

PathQuery is a tool to declaratively define file searches that returns a list
of `path.py <https://github.com/jaraco/path.py>`_ Path objects.

Example
-------

Search for all files recursively except in the node_modules folder and change its perms:

.. code-block:: python

    from pathquery import pathquery

    for path in pathquery("yourdir").ext("js") - pathquery("yourdir/node_modules"):
        path.chmod(0755)

Install
-------

To use::

  $ pip install pathquery

API
---

Path properties can be inspected as part of the query:

.. code-block:: python

    pathquery("yourdir").is_dir()
    pathquery("yourdir").is_not_dir()
    pathquery("yourdir").is_symlink()
    pathquery("yourdir").is_not_symlink()

Queries are also chainable:

.. code-block:: python

    for path in pathquery("yourdir").ext("pyc").is_symlink() - pathq("yourdir/node_modules"):
        path.remove()
