==========
simplepath
==========

.. image:: https://badge.fury.io/py/simplepath.png
    :target: http://badge.fury.io/py/simplepath

.. image:: https://travis-ci.org/dealertrack/simplepath.png?branch=master
    :target: https://travis-ci.org/dealertrack/simplepath

.. image:: https://coveralls.io/repos/dealertrack/simplepath/badge.svg
    :target: https://coveralls.io/r/dealertrack/simplepath


``simplepath`` is a library for data-structure lookups
using super simple expressions with performance in mind.
*"simplepath"* is a word play on some other ``*path`` technologies
such as ``xpath``, ``jsonpath``, ``jpath``, etc.

* Free software: MIT license
* GitHub: https://github.com/dealertrack/simplepath

Inspiration
-----------

The inspiration for ``simplepath`` was performance. Many other
libraries focus on making single lookups, however they fall 
short when a lot of data needs to be queried.

For example if a dictionary with some structure needs to be converted
into another dictionary with a different structure, a simple and
configurable way of doing that might be to define a configuration
dictionary where the keys will be the keys of the output dictionary, 
and values will be lookup expressions to get appropriate data::

    {
        "greetings": "foo.greeting",
        "planet": "foo.[0].planet",
        ...
    }

The above approach is easy to implement, however is not very performant
since for each lookup the lookup expression will have to be evaluated.
At dealertrack, we needed to do something similar at some point and
tried `jsonpath-rw <https://pypi.python.org/pypi/jsonpath-rw>`_
which would sometimes take 15 seconds to map dictionaries with only
a couple hundred expressions. Upon some investigation, most of the
time was being spent in `ply <https://pypi.python.org/pypi/ply>`_.
Unfortunately we did not find another comparable library which
accomplished everything we needed, and satisfied our performance
requirements, so ``simplepath`` was born.

Installing
----------

You can install ``simplepath`` using pip::

    $ pip install simplepath

Quick Guide
-----------

Here is a quick example.

::

    from simplepath.mapper import Mapper

    class MyMapper(Mapper):
        config = {
            'greetings': 'example.greetings',
            'to': 'example.planets.<find:planet=Earth>.residents',
        }

    data = {
        'example': {
            'greetings': 'Hello',
            'planets': [
                {
                    'planet': 'Mars',
                    'residents': 'martians',
                },
                {
                    'planet': 'Earth',
                    'residents': 'people',
                },
                {
                    'planet': 'Space',
                    'residents': 'aliens',
                },
            ]
        }
    }

    MyMapper.map_data(data) == {
        'greetings': 'Hello',
        'to': 'people',
    }

Testing
-------

To run the tests you need to install testing requirements first::

    $ make install

Then to run tests, you can use ``nosetests`` or simply use Makefile command::

    $ nosetests -sv
    # or
    $ make test




History
-------

0.3.5 (2018-06-14)
~~~~~~~~~~~~~~~~~~~~~

* Fixed bug in which ``TypeError`` would result if a ``None`` value were mapped to a ``ListConfig``


0.3.4 (2017-07-28)
~~~~~~~~~~~~~~~~~~~~~

* Using wheel for distribution
* Removed tests from being packaged
* Switched to using Python 3.6 vs 3.5 for running Travis builds

0.3.3 (2016-05-15)
~~~~~~~~~~~~~~~~~~~~~

* Fixed bug where global LUT would leak data when calling expressions
  within custom lookups. See `#11 <https://github.com/dealertrack/simplepath/issues/11>`_.
* Switched to using Python 3.5 vs 3.4 for running Travis builds

0.3.2 (2015-09-14)
~~~~~~~~~~~~~~~~~~~~~

* Registered ``AsTypeLookup`` and ``ArithmeticLookup`` as ``as_type`` and ``arith`` lookups
  in default lookup registry

0.3.1 (2015-08-28)
~~~~~~~~~~~~~~~~~~~~~

* Added the ``AsTypeLookup`` and ``ArithmeticLookup``

0.3.0 (2015-07-15)
~~~~~~~~~~~~~~~~~~~~~

* Added ability to use lists in the simplepath config which will generate a list in the mapped data

0.2.0 (2015-06-26)
~~~~~~~~~~~~~~~~~~~~~

* Added ``deepvars`` utility which is useful when using simplepath with objects

0.1.1 (2015-03-31)
~~~~~~~~~~~~~~~~~~~~~

* Fixed a link to the repo in the package description

0.1.0 (2015-01-08)
~~~~~~~~~~~~~~~~~~~~~

* First release


Credits
-------

This utility was created at `dealertrack technologies`_
(`dealertrack GitHub`_) for our internal use so thank you
dealertrack for allowing to contribute the utility
to the open-source community.

Development Lead
~~~~~~~~~~~~~~~~

* Miroslav Shubernetskiy  - https://github.com/miki725

Contributors
~~~~~~~~~~~~

* Ndubisi Onuora - https://github.com/NdubisiOnuora


.. _dealertrack GitHub: https://github.com/Dealertrack
.. _dealertrack technologies: https://www.dealertrack.com


License
-------

The MIT License (MIT)

Copyright (c) 2014, dealertrack technologies

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.


