.. -*- coding: utf-8 -*-

Changes
-------

1.7 (2018-05-24)
~~~~~~~~~~~~~~~~

* Avoid prettification of BindParameters


1.6 (2018-05-23)
~~~~~~~~~~~~~~~~

* Use Sphinx 1.7+ API to install the extension


1.5 (2017-08-10)
~~~~~~~~~~~~~~~~

* New option ``autodoc_sa_prettifier_options`` to pass arbitrary keyword options to the
  prettifier function


1.4 (2017-08-09)
~~~~~~~~~~~~~~~~

* Replace the dynamic argument placeholders injected by SA with their literal values, leaving
  the developer's explicit bindparams alone


1.3 (2017-08-08)
~~~~~~~~~~~~~~~~

* Handle also the `pg_query`__ SQL prettifier

* New options, ``autodoc_sa_prettifier`` and ``autodoc_pygments_lang``

__ https://pypi.python.org/pypi/pg_query


1.2 (2017-03-22)
~~~~~~~~~~~~~~~~

* Minor tweak, no externally visible changes


1.1 (2017-01-17)
~~~~~~~~~~~~~~~~

* First release on PyPI


1.0 (unreleased)
~~~~~~~~~~~~~~~~

* Polished, tested and extended to support class' attributes as well

* Extracted from metapensiero.sphinx.patchdb
