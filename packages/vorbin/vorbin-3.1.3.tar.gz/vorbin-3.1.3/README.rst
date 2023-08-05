The VorBin package
==================

**Adaptive Voronoi Binning of Two Dimensional Data**

.. image:: https://img.shields.io/pypi/v/vorbin.svg
        :target: https://pypi.org/project/vorbin/
.. image:: https://img.shields.io/badge/arXiv-astroph:0302262-orange.svg
    :target: https://arxiv.org/abs/astro-ph/0302262
.. image:: https://img.shields.io/badge/DOI-10.1046/...-green.svg
        :target: https://doi.org/10.1046/j.1365-8711.2003.06541.x

This VorBin package is a Python implementation of the two-dimensional adaptive
spatial binning method of `Cappellari & Copin (2003)
<http://adsabs.harvard.edu/abs/2003MNRAS.342..345C>`_. It uses Voronoi
tessellations to bin data to a given minimum signal-to-noise ratio.

.. contents::

Attribution
-----------

If you use this software for your research, please cite
`Cappellari & Copin (2003) <http://adsabs.harvard.edu/abs/2003MNRAS.342..345C>`_.
The BibTeX entry for the paper is::

    @ARTICLE{Cappellari2003,
        author = {{Cappellari}, M. and {Copin}, Y.},
        title = "{Adaptive spatial binning of integral-field spectroscopic
            data using Voronoi tessellations}",
        journal = {MNRAS},
        eprint = {astro-ph/0302262},
        year = 2003,
        volume = 342,
        pages = {345-354},
        doi = {10.1046/j.1365-8711.2003.06541.x}
    }

Installation
------------

install with::

    pip install vorbin

Without writing access to the global ``site-packages`` directory, use::

    pip install --user vorbin

Documentation
-------------

A usage example is provided by the procedure ``voronoi_2d_binning_example.py``.

Perform the following simple steps to bin you own 2D data with minimal Python interaction:

1. Write your data vectors [X, Y, Signal, Noise] in the text file
   ``voronoi_2d_binning_example.txt``, following the example provided;

2. Change the line ``targetSN = 50.0`` in the procedure ``voronoi_2d_binning_example.py``,
   to specify the desired target S/N of your final bins;

3. Run the program ``voronoi_2d_binning_example`` and wait for the final plot to appear.
   The output is saved in the text file ``voronoi_2d_binning_output.txt``. The
   last column BIN_NUM in the file is *all* that is needed to actually bin the data;

4. Read the documentation at the beginning of the file ``voronoi_2d_binning.py`` to
   fully understand the meaning of the various optional output parameters.

When some pixels have no signal
-------------------------------

Binning should not be used blindly when some pixels contain significant noise
but virtually no signal. This situation may happen e.g. when extracting the gas
kinematics from observed galaxy spectra. One way of using voronoi_2d_binning
consists of first selecting the pixels with S/N above a minimum threshold and
then binning each set of connected pixels *separately*. Alternatively one may
optimally weight the pixels before binning. For details, see Sec. 2.1 of
`Cappellari & Copin (2003) <http://adsabs.harvard.edu/abs/2003MNRAS.342..345C>`_.

Binning X-ray data
------------------

For X-ray data, or other data coming from photon-counting devices the noise is
generally accurately Poissonian. In the Poissonian case, the S/N in a bin can
never decrease by adding a pixel (see Sec.2.1 of
`Cappellari & Copin 2003 <http://adsabs.harvard.edu/abs/2003MNRAS.342..345C>`_),
and it is preferable to bin the data *without* first removing the observed pixels
with no signal.

Binning very big images
-----------------------

Computation time in voronoi_2d_binning scales nearly as npixels^1.5, so it may
become a problem for large images (e.g. at the time of writing npixels > 1000x1000).
Let's assume that we really need to bin the image as a whole and that the S/N in
a significant number of pixels is well above our target S/N. As for many other
computational problems, a way to radically decrease the computation time consists
of proceeding in a hierarchical manner. Suppose for example we have a 4000x4000
pixels image, we can do the following:

1. Rebin the image regularly (e.g. in groups of 8x8 pixels) to a manageable
   size of 500x500 pixels;
2. Apply the standard Voronoi 2D-binning procedure to the 500x500 image;
3. Transform all unbinned pixels (which already have enough S/N) of the
   500x500 Voronoi 2D-binned image back into their original individual
   full-resolution pixels;
4. Now apply Voronoi 2D-binning only to the connected regions of
   full-resolution pixels;
5. Merge the set of lower resolution bins with the higher resolution ones.

License
-------

Copyright (c) 2001-2018 Michele Cappellari

This software is provided as is without any warranty whatsoever.
Permission to use, for non-commercial purposes is granted.
Permission to modify for personal or internal use is granted,
provided this copyright and disclaimer are included in all
copies of the software. All other rights are reserved.
In particular, redistribution of the code is not allowed.

