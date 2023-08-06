pysetenv
========

**pysetenv** is a simple Python-based program to allow users to set
environment variables before executing a command. While it's
cross-platform, it's primarily designed to be used on Windows, where
``cmd.exe`` makes this considerably more difficult than on POSIX
systems.

Why?
----

pysetenv is designed as a support package to help authors of other
Python packages create command-line strings like you'd expect from
``sh`` or the ``env`` command. Generally, pysetenv will only be
installed on Windows systems, with POSIX systems using ``sh`` or
``env``. For example, in your ``setup.py``, you would write:

.. code:: python

    setup(
        # ...
        install_requires=['pysetenv;platform_system=="Windows"'],
    )

License
-------

This project is licensed under the `BSD 3-clause license <LICENSE>`__.


