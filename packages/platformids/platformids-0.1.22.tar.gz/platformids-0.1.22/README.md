platformids
===========

<img align="right" src="docsrc/pyplatformids-64x64.png?raw=true"/>

The ‘platformids‘ package provides the enumeration of runtime platforms. This extends the standard Python facilities by

* more specific canonical platform enumerations
* provides additional hierarchical bitmasks for faster processing
* provides mapping of string and numeric representation for human display
* provides a boolean flag V3K for Python3
* provides a bitmask Vxyz for the complete Python version for faster operations

The supported platforms are:

* Linux, BSD, Unix, OS-X, Cygwin, and Windows

* Python2.7+, Python3.5+


**Downloads**:

* Sourceforge.net: https://sourceforge.net/projects/pyplatformids/files/

* Github: https://github.com/ArnoCan/pyplatformids/

**Online documentation**:

* https://pypi.python.org/pypi/pyplatformids/
* https://pythonhosted.org/pyplatformids/

setup.py
--------

The installer adds a few options to the standard setuptools options.

* *build_sphinx*: Creates documentation for runtime system by Sphinx, html only. Calls 'callDocSphinx.sh'.

* *build_epydoc*: Creates documentation for runtime system by Epydoc, html only. Calls 'callDocEpydoc.sh'.

* *instal_doc*: Install a local copy of the previously build documents in accordance to PEP-370.

* *test*: Runs PyUnit tests by discovery.

* *--help-platformids*: Displays this help.

* *--no-install-required*: Suppresses installation dependency checks, requires appropriate PYTHONPATH.

* *--offline*: Sets online dependencies to offline, or ignores online dependencies.

* *--exit*: Exit 'setup.py'.
 

Project Data
------------

* PROJECT: 'platformids'

* MISSION: Extend the standard PyUnit package for arbitrary ExecUnits.

* VERSION: 00.01

* RELEASE: 00.01.020

* STATUS: alpha

* AUTHOR: Arno-Can Uestuensoez

* COPYRIGHT: Copyright (C) 2010,2011,2015-2018 Arno-Can Uestuensoez @Ingenieurbuero Arno-Can Uestuensoez

* LICENSE: Artistic-License-2.0 + Forced-Fairplay-Constraints

  Refer to enclosed documents:
  
  *  'ArtisticLicense20.html' - for base license: Artistic-License-2.0 [[pythonhosted.org](https://pythonhosted.org/pyplatformids/_static/ArtisticLicense20.html)] [[github.com](https://github.com/ArnoCan/pyplatformids/ArtisticLicense20.html)] [[sourceforge.net](https://sourceforge.net/projects/pyplatformids/files/ArtisticLicense20.html)]

  *  'licenses-amendments.txt' - for amendments: Forced-Fairplay-Constraints [[pythonhosted.org](https://pythonhosted.org/pyplatformids/_static/licenses-amendments.txt)] [[github.com](https://github.com/ArnoCan/pyplatformids/licenses-amendments.txt)] [[sourceforge.net](https://sourceforge.net/projects/pyplatformids/files/licenses-amendments.txt)]

VERSIONS and RELEASES
---------------------

**Planned Releases:**

* RELEASE: 00.00.00x - Pre-Alpha: Extraction of the features from hard-coded application into a reusable package.

* RELEASE: 00.01.00x - Alpha: Completion of basic features. 

* RELEASE: 00.02.00x - Alpha: Completion of features, stable interface. 

* RELEASE: 00.03.00x - Beta: Accomplish test cases for medium to high complexity.

* RELEASE: 00.04.00x - Production: First production release.

* RELEASE: 00.05.00x - Production: Various performance enhancements.

* RELEASE: 00.06.00x - Production: Security review.

**Current Release: 00.01.020 - Alpha:**

Python support:

*  Python2.7, and Python3.5+

OS-Support:

* Linux: Fedora, CentOS, RHEL - others should work, ToDo: Debian, and SuSE 

* BSD - remote Eclipse by ePyUnit: OpenBSD - FreeBSD and others should work

* Mac-OS: Snow Leopard - others should work too

* Windows: Win7, Win10 - others see Cygwin

* Cygwin: 2.874/64 bit


Major Changes:

* Initial version.

