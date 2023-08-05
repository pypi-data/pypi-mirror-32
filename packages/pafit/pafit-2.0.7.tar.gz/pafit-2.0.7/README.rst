The PaFit Package
=================

**Kinematic Position Angle Fitting for Galaxies**

.. image:: https://img.shields.io/pypi/v/pafit.svg
        :target: https://pypi.org/project/pafit/
.. image:: https://img.shields.io/badge/arXiv-astroph:0512200-orange.svg
    :target: https://arxiv.org/abs/astro-ph/0512200
.. image:: https://img.shields.io/badge/DOI-10.1111/...-blue.svg
        :target: https://doi.org/10.1111/j.1365-2966.2005.09902.x

The PaFit package is a Python implementation of the algorithm to fit the global
kinematic position angle of galaxies, described in Appendix C of
`Krajnovic et al. (2006) <http://adsabs.harvard.edu/abs/2006MNRAS.366..787K>`_.

Attribution
-----------

If you use this software for your research, please cite `Krajnovic et al. (2006)
<http://adsabs.harvard.edu/abs/2006MNRAS.366..787K>`_.
The BibTeX entry for the paper is::

    @Article{Krajnovic2006,
        title = {Kinemetry: a generalization of photometry to the higher moments
            of the line-of-sight velocity distribution},
        author = {{Krajnovi{\'c}}, D. and {Cappellari}, M. and {de Zeeuw}, P.~T.
            and {Copin}, Y.},
        journal = {MNRAS},
        eprint = {arXiv:astro-ph/0512200}
        year = {2006},
        pages = {787-802},
        volume = {366},
        doi = {10.1111/j.1365-2966.2005.09902.x}
    }

Installation
------------

install with::

    pip install pafit

Without writing access to the global ``site-packages`` directory, use::

    pip install --user pafit

Documentation
-------------

See the ``pafit/fit_kinematic_pa.py`` file header.

License
-------

Copyright (c) 2005-2018 Michele Cappellari

This software is provided as is without any warranty whatsoever.
Permission to use, for non-commercial purposes is granted.
Permission to modify for personal or internal use is granted,
provided this copyright and disclaimer are included in all
copies of the software. All other rights are reserved.
In particular, redistribution of the code is not allowed.

