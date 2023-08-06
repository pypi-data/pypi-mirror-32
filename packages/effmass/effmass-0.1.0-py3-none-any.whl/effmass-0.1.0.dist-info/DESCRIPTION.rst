# effmass

[![Build Status](https://travis-ci.com/lucydot/effmass.svg?branch=master)](https://travis-ci.com/lucydot/effmass)
[![Test Coverage](https://codeclimate.com/github/lucydot/effmass/badges/coverage.svg)](https://codeclimate.com/github/lucydot/effmass/coverage)
![Documentation](https://readthedocs.org/projects/pip/badge/?version=latest)

`effmass` is a Python package for calculating various definitions of effective mass from the electronic bandstructure of a semiconducting material. It consists of a core class that calculates the effective mass and other associated properties of selected bandstructure segments. The module also contains functions for locating bandstructure extrema and plotting approximations to the dispersion.

Examples are provided in a Jupyter notebook [here](https://nbviewer.jupyter.org/github/lucydot/effmass/blob/master/paper/notebook.ipynb).
API documentation is [here](https://effmass.readthedocs.io/en/latest/).
Source code is available as a git repository at [https://github.com/lucydot/effmass](https://github.com/lucydot/effmass).

## Features

`effmass` can:

**Read in a bandstructure:**
This requires the `VASP` output files `PROCAR` and `OUTCAR`. It is assumed you have walked through a 1D slice of the Brillouin Zone, capturing the maxima and minima of interest. `effmass` uses the Python package [vasppy](https://github.com/bjmorgan/vasppy) for parsing `VASP` output.

**Locate extrema:**
These correspond to the valence band maxima and conduction band minima. Maxima and minima within a certain energy range can also be located.

**Calculate curvature, transport and optical effective masses:**
The curvature (aka inertial), transport and optical effective masses are calculated using the derivatives of a fitted polynomial function. The optical effective mass can also be calculated assuming a Kane dispersion.

**Assess the extent of non-parabolicity:**
Parameters of the Kane quasi-linear dispersion are calculated to quantify the extent of non-parabolicity over a given energy range. add

**Calculate the quasi-fermi level for a given carrier concentration:**
This requires the `VASP` output file `DOSCAR`. Using density-of-states data and assuming no thermal smearing, `effmass` can calculate the energy to which states are occupied. This is a useful approximation to the quasi-Fermi level.

**Plot fits to the dispersion:**
Selected bandstructure segments and approximations to the dispersion (assuming a Kane, quadratic, or higher order fit) can be visualised.

## Development

Development is in progress and hosted on [Github](https://github.com/lucydot/effmass). Please use the [issue tracker](https://github.com/lucydot/effmass/issues/) for feature requests, bug reports and more general questions. If you would like to contribute, please do so via a pull request.

## Installation

TODO - register on PyPI

## Tests

Automated testing of the latest commit happens [here](https://travis-ci.com/lucydot/effmass).

Manual tests can be run using 
```
python -m pytest
```

This code has been tested with Python versions 3.6.

## Documentation

An overview of the features of effmass along with example code is contained in a [Jupyter notebook](https://nbviewer.jupyter.org/github/lucydot/effmass/blob/master/paper/notebook.ipynb), which is available in the `paper` directory.

API documentation is available [here](https://effmass.readthedocs.io/en/latest/).

## Citing

TODO - register on Zenodo for DOI.
TODO - submit to JOSS for DOI.


