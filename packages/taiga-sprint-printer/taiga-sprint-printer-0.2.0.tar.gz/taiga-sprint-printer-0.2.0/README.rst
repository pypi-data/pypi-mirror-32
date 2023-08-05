Taiga Sprint Printer
====================

|Build Status| |codecov|

📃 A simple tool for printing your sprint from Taiga

Install
-------

🐍 You need ``Python3`` to run this program

::

    pip install taiga-sprint-printer

And launch it from the command line
-----------------------------------

::

    sprint-printer

Set the color for user stories and tasks
----------------------------------------

Taiga sprint printer has two default colors, 🔴\ ``red`` for user stories
and 🔵\ ``blue`` for tasks. If you want to change these values run:

::

    sprint-printer colors

Reset account configuration
---------------------------

Taiga Sprint Printer asks you the server location and your username only
the first time. If you want to change these settings launch
sprint-printer with the ``new`` option:

::

    sprint-printer new

Contribute
----------

Feel free to send suggestions, open issues and create a PR to improve
this software!

If you want to start developing, create a virtualenv with Python3 and
install requirements

::

    pip install -r requirements-dev.txt

License
-------

MIT - Lotrèk 2018

.. |Build Status| image:: https://travis-ci.org/lotrekagency/taiga-sprint-printer.svg?branch=master
   :target: https://travis-ci.org/lotrekagency/taiga-sprint-printer
.. |codecov| image:: https://codecov.io/gh/lotrekagency/taiga-sprint-printer/branch/master/graph/badge.svg
   :target: https://codecov.io/gh/lotrekagency/taiga-sprint-printer
