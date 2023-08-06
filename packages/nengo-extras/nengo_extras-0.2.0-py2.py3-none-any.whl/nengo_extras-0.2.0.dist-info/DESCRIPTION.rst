************
Nengo extras
************

Extra utilities and add-ons for Nengo.

This repository contains utilities that occupy
a liminal space not quite generic enough for inclusion in Nengo_,
but useful enough that they should be publicly accessible.

Some of these utilities may eventually migrate to Nengo_,
and others may be split off into their own separate repositories.

.. _Nengo: https://github.com/nengo/nengo

Installation
============

To install Nengo extras, we recommend using ``pip``.

.. code:: bash

   pip install nengo-extras

Usage
=====

Example notebooks can be found
in the ``docs/examples`` directory.

For a listing of the contents of this repository,
and information on how to use it,
see the `full documentation <https://www.nengo.ai/nengo-extras>`_.

Development
===========

To run the unit tests:

   pytest nengo_extras [--plots]

To run the static checks:

.. code-block:: bash

   .ci/static.sh run

To build the documentation:

.. code-block:: bash

   sphinx-build docs docs/_build

***************
Release history
***************

.. Changelog entries should follow this format:

   version (release date)
   ======================

   **section**

   - One-line description of change (link to Github issue/PR)

.. Changes should be organized in one of several sections:

   - Added
   - Changed
   - Deprecated
   - Removed
   - Fixed

0.2.0 (May 31, 2018)
====================

**Added**

- Added the association matrix learning rule (AML)
  to learn associations from cue vectors to target vectors
  in a one-shot fashion without catastrophic forgetting.
  (`#72 <https://github.com/nengo/nengo-extras/pull/72>`_)
- Added classes to convert Nengo models to GEXF for visualization with Gephi.
  (`#54 <https://github.com/nengo/nengo-extras/pull/54>`_)
- Added a ``Camera`` process to stream images from a camera to Nengo.
  (`#61 <https://github.com/nengo/nengo-extras/pull/61>`_)

0.1.0 (March 14, 2018)
======================

Initial release of Nengo Extras!
Tested with Nengo 2.7.0, but should work with earlier versions.
If you run into any issues, please
`file a bug report <https://github.com/nengo/nengo-extras/issues/new>`_.


