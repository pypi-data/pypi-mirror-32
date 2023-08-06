pysourceinfo
============

The 'pysourceinfo' package provides basic runtime information on executed 
sourcefiles based on 'inspect', 'sys', 'os', and additional sources.
The covered objects include packages, modules/files and functions/methods/scripts. 

The supported platforms are:

* Linux, BSD, Unix, OS-X, Cygwin, and Windows

* Python2.7+, Python3.5+


**Online documentation**:

* https://pysourceinfo.sourceforge.io/

**Runtime-Repository**:

* PyPI: https://pypi.org/project/pysourceinfo/

  Install: *pip install pysourceinfo*, see also 'Install'.


**Downloads**:

* Sourceforge.net: https://sourceforge.net/projects/pysourceinfo/files/

* bitbucket.org: https://bitbucket.org/acue/pysourceinfo

* github: https://github.com/ArnoCan/pysourceinfo/

* PyPI: https://pypi.org/project/pysourceinfo/
 

Project Data
------------

* PROJECT: 'pysourceinfo'

* MISSION: Extend the standard PyUnit package for arbitrary ExecUnits.

* VERSION: 00.01

* RELEASE: 00.01.027

* STATUS: alpha

* AUTHOR: Arno-Can Uestuensoez

* COPYRIGHT: Copyright (C) 2010,2011,2015-2017 Arno-Can Uestuensoez @Ingenieurbuero Arno-Can Uestuensoez

* LICENSE: Artistic-License-2.0 + Forced-Fairplay-Constraints


  Refer to enclosed documents ArtisticLicense20.html and licenses-amendments.txt:
  
  * bitbucket: https://bitbucket.org/acue/pyplatformids/src
  * github.com: https://github.com/ArnoCan/pyplatformids/ArtisticLicense20.html
  * sourceforge.net: https://pyplatformids.sourceforge.io/_static/ArtisticLicense20.html

  * pythonhosted.org: https://bitbucket.org/acue/pyplatformids/src
  * github.com: https://github.com/ArnoCan/pyplatformids/licenses-amendments.txt
  * sourceforge.net: https://pyplatformids.sourceforge.io/_static/licenses-amendments.txt

VERSIONS and RELEASES
---------------------

**Current Release: 00.01.027 - Alpha:**

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

* Split more stack related parts into package *PyStackInfo*

* Added special support for decorators.

* Added support for some special cases of Python syntax elements, e.g. support of nested-classes.

* Added several new functions.

* Enhanced documentation
