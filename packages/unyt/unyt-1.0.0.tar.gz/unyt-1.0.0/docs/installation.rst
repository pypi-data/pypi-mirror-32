.. highlight:: shell

============
Installation
============


Stable release
--------------

To install unyt, run this command in your terminal:

.. code-block:: console

    $ pip install unyt

This is the preferred method to install unyt, as it will always install the most recent stable release.

If you don't have `pip`_ installed, this `Python installation guide`_ can guide
you through the process.

.. _pip: https://pip.pypa.io
.. _Python installation guide: http://docs.python-guide.org/en/latest/starting/installation/


From source
-----------

The sources for unyt can be downloaded from the `Github repo`_.

You can either clone the public repository:

.. code-block:: console

    $ git clone git://github.com/yt-project/unyt

Or download the `tarball`_:

.. code-block:: console

    $ curl  -OL https://github.com/yt-project/unyt/tarball/master

Once you have a copy of the source, you can install it by navigating to the root of the installation and issuing the following command:

.. code-block:: console

    $ pip install .

If you would like to make an "editable" where you can directly edit the
Python source files of the installed version of ``unyt``, then you can do:

.. code-block:: console

    $ pip install -e .

.. _Github repo: https://github.com/yt-project/unyt
.. _tarball: https://github.com/yt-project/unyt/tarball/master
