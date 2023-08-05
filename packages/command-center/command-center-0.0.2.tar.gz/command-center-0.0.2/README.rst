=====================
Python Command Center
=====================

GUI window to easily run your own frequently used commands

Installation
============
Install from pypi using pip::

   pip install command-center

Usage
=====
Simply run pcc in the terminal::

   $ pcc

Options
=======

Columns
-------
You can change the number of columns that command center will display by
including a ``--col=N`` argument.::

   $ pcc --col=4

Extra Commands
--------------
You can include additional commands for only the next execution by including
key value pair arguments such as ``key=value``. Note that unless your key and
value do not include spacing you will need to double quote the argument.::

   $ pcc "pytest abc=pytest a/b/c.py -x"
