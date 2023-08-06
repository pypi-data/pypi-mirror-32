pysourceinfo
==============

<img align="right" src="docsrc/pysourceinfo-64x64.png?raw=true"/>

The *pysourceinfo* package provides basic runtime information on executed 
sourcefiles based on *inspect*, *sys*, *os*, and additional sources.
The covered objects include packages, modules/files and functions/methods/scripts. ::


    +--------------------------------------------------------+
    |                     pysourceinfo                       |
    +--------------------------------------------------------+
          |              |              |               |
          V              V              V               V
    +----------+    +---------+    +---------+    +----------+
    | inspect  |    |   sys   |    |   os    |    |   dis    |
    +----------+    +---------+    +---------+    +----------+


The The *pysourceinfo* package provides the modules:

* *pysourceinfo.fileinfo* - [File System Names](https://pythonhosted.org/pysourceinfo/namebinding.html#FILENAMEBINDING)
* *pysourceinfo.objectinfo* - [Object Identifier](https://pythonhosted.org/pysourceinfo/namebinding.html#OBJECTNAMEBINDING)
* *pysourceinfo.bininfo* - [Compiled Binaries](https://pythonhosted.org/pysourceinfo/namebinding.html#OBJECTNAMEBINDING)
* *pysourceinfo.infolists* - [Sets and Entities](https://pythonhosted.org/pysourceinfo/namebinding.html#SETSANDENTITIES)
* *pysourceinfo.helper* - [Helper](https://pythonhosted.org/pysourceinfo/helper.html#)

The extension packackage *stackinfo* provides for object addresses within modules - Object Identifier OID. 

* *stackinfo* - *PyStackInfo* \[[doc](https://pythonhosted.org/pystackinfo/)] \[[src](https://pypi.python.org/pypi/pystackinfo/)]

The old interface is now deprecated, kept for some releases for compatibility:

* *pysourceinfo.PySourceInfo* - deprecated old module for compatibility, going to be canceled


For code examples refer to the source package 'pysourceinfo.UseCases' and 'pysourceinfo.tests'.

**Downloads**:

* Sourceforge.net: https://sourceforge.net/projects/pysourceinfo/files/

* Github: https://github.com/ArnoCan/pysourceinfo/

**Online documentation**:

* https://pypi.python.org/pypi/pysourceinfo/
* https://pythonhosted.org/pysourceinfo/

setup.py
--------

The installer adds a few options to the standard setuptools options.

* *build_doc*: Creates *Sphinx* based documentation with embeded javadoc-style API documentation by *Epydoc*, html only.

* *build_sphinx*: Creates Sphinx part of the documentation as standalone html. Calls 'callDocSphinx.sh'.

* *build_epydoc*: Creates Epydoc part of the documentation as standalone html.

* *install_project_doc*: Install a local copy into the doc directory of the project.

* *instal_doc*: Install a local copy of the previously build documents in accordance to PEP-370.

* *test*: Runs PyUnit tests by discovery.

* *usecase*: Runs PyUnit tests on UseCases subdirectory by discovery.


* *--help-pysourceinfo*: Displays this help.

* *--no-install-required*: Suppresses installation dependency checks, requires appropriate PYTHONPATH.

* *--offline*: Sets online dependencies to offline, or ignores online dependencies.

* *--exit*: Exit 'setup.py'.


After successful installation the 'selftest' verifies basic checks by:

  *pysourceinfo --selftest*

with the exit value '0' when OK.

The option '-v' raises the degree of verbosity for inspection

  *pysourceinfo --selftest -v -v -v -v*
 

Project Data
------------

* PROJECT: 'pysourceinfo'

* MISSION: Extend the standard PyUnit package for arbitrary ExecUnits.

* VERSION: 00.01

* RELEASE: 00.01.020

* STATUS: alpha

* AUTHOR: Arno-Can Uestuensoez

* COPYRIGHT: Copyright (C) 2010,2011,2015-2017 Arno-Can Uestuensoez @Ingenieurbuero Arno-Can Uestuensoez

* LICENSE: Artistic-License-2.0 + Forced-Fairplay-Constraints
  Refer to enclose documents:
  
    *  ArtisticLicense20.html - for base license: Artistic-License-2.0 

    *  licenses-amendments.txt - for amendments: Forced-Fairplay-Constraints

VERSIONS and RELEASES
---------------------

**Planned Releases:**

* RELEASE: 00.00.00x - Pre-Alpha: Extraction of the features from hard-coded application into a reusable package.

* RELEASE: 00.01.00x - Alpha: Completion of basic features. 

* RELEASE: 00.02.00x - Alpha: Completion of features, stable interface. 

* RELEASE: 00.03.00x - Beta: Accomplish test cases for medium to high complexity.

* RELEASE: 00.04.00x - Production: First production release. 

* RELEASE: 00.05.00x - Production: Various performance enhancements.


**Current Release: 00.01.020 - Alpha:**

Python support: 

* Standard Python(CPython) - Python2.7, and Python3.5+ 

* PyPy - 5.10+ - Python2.7+, and Python3.5+

OS-Support:

* Linux: Fedora, CentOS-6, CentOS-7, Debian-9, RHEL-6, RHEL-7 - others should work, Python-Release 

* Windows: Win7, Win10 - others see Cygwin

* Mac-OS: Snow Leopard - others should work too.

* Cygwin: 2.874/64 bit - 32bit should just work

* BSD: OpenBSD-6 - others should work

* Raspbian: on RaspberryPI(2,3), Asus-Tinker-Board

* OpenWRT: on RaspberryPI(2)

Major Changes:

* Python2.6 support dropped.

* Python3.5+ support introduced.

* PyPy tests added.

* Changed module structure and fitting names into overall category based naming schema.

* Split more stack related parts into package *PyStackInfo* \[[doc](https://pythonhosted.org/pystackinfo/)] \[[src](https://pypi.python.org/pypi/pystackinfo/)]

* Added special support for decorators.

* Added support for some special cases of Python syntax elements, e.g. support of nested-classes.

* Added several new functions.

* Enhanced documentation
