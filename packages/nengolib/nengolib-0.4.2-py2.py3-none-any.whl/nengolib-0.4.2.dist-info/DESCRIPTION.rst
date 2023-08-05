.. image:: https://arvoelke.github.io/nengolib-docs/_static/logo.png
   :width: 64
   :height: 64
   :target: https://github.com/arvoelke/nengolib
   :alt: Nengolib Logo

.. image:: https://travis-ci.org/arvoelke/nengolib.svg?branch=master
   :target: https://travis-ci.org/arvoelke/nengolib
   :alt: Build Status

.. image:: https://codecov.io/github/arvoelke/nengolib/coverage.svg?branch=master
   :target: https://codecov.io/github/arvoelke/nengolib?branch=master
   :alt: Code Coverage

import nengolib
===============

Additional extensions and tools for modelling dynamical systems in
`Nengo <https://github.com/nengo/nengo>`__.


`Documentation <https://arvoelke.github.io/nengolib-docs/>`__
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

This project's documentation is hosted on GitHub.IO:
https://arvoelke.github.io/nengolib-docs/.


Development
~~~~~~~~~~~

To install the development version of nengolib::

    git clone https://github.com/arvoelke/nengolib
    cd nengolib
    python setup.py develop

Notebooks can be run manually in ``docs/notebooks`` by running::

    pip install jupyter
    jupyter notebook

***************
Release History
***************

0.4.2 (May 18, 2018)
====================

Tested against Nengo versions 2.1.0-2.7.0.

**Added**

- Solving for connection weights by accounting for the neural
  dynamics. To use, pass in ``nengolib.Temporal()`` to
  ``nengo.Connection`` for the ``solver`` parameter.
  Requires ``nengo>=2.5.0``.
  (`#137 <https://github.com/arvoelke/nengolib/pull/137>`_)

0.4.1 (December 5, 2017)
========================

Tested against Nengo versions 2.1.0-2.6.0.

**Fixed**

- Compatible with newest SciPy release (1.0.0).
  (`#130 <https://github.com/arvoelke/nengolib/pull/130>`_)

0.4.0b (June 7, 2017)
=====================

Initial beta release of nengolib.
Tested against Nengo versions 2.1.0-2.4.0.


