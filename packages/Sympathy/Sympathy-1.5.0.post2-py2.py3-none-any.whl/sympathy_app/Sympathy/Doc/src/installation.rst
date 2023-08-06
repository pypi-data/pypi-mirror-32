.. This file is part of Sympathy for Data.
..
..  Copyright (c) 2017 System Engineering Software Society
..
..     Sympathy for Data is free software: you can redistribute it and/or modify
..     it under the terms of the GNU General Public License as published by
..     the Free Software Foundation, either version 3 of the License, or
..     (at your option) any later version.
..
..     Sympathy for Data is distributed in the hope that it will be useful,
..     but WITHOUT ANY WARRANTY; without even the implied warranty of
..     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
..     GNU General Public License for more details.
..     You should have received a copy of the GNU General Public License
..     along with Sympathy for Data. If not, see <http://www.gnu.org/licenses/>.

Installation instructions
=========================

For Windows
-----------
Download the latest version of Sympathy from the `official
homepage <https://www.sympathyfordata.com/>`_. If you are using any
custom node libraries then make sure to select the same Python version
(Python 2 or Python 3) as the libraries have been written for.

After downloading, run the installer and follow the
instructions. This will install Sympathy as well as a custom
Python version with all dependencies for it.


For Mac OS
----------

These instructions are written for MacOS X 10.11.6 using MacPorts.
They have also been tested on 10.13.0.

Start by installing Xcode from the App Store (that will download an XCode
installer, so this is a two-stage process).

.. code-block:: bash

   sudo xcode-select --install
   sudo xcodebuild -license


You can install Sympathy either for Python 3 (recommended) or
Python 2.7. The common installation steps are needed for both cases.

Installing Python 2 environment
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Download and install `MacPorts <http://www.macports.org>`__.
Before continuing, it is recommended to run

.. code-block:: bash

   sudo port selfupdate

Then install the dependencies

.. code-block:: bash

   sudo port install python27 py27-zmq py27-ply py27-pyside py27-pyodbc py27-psutil py27-spyder py27-sphinx py27-pandas py27-numpy py27-scipy py27-matplotlib py27-pyflakes py27-pylint py27-ipython py27-h5py py27-lxml py27-xlwt py27-xlrd py27-pip graphviz py27-XlsxWriter py27-scikit-learn py27-scikit-image py27-mock


Make PySide detected as installed to avoid rebuild.

.. code-block:: bash

   sudo bash -c "cat >/opt/local/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/site-packages/PySide-1.2.2.egg-info <<EOF
   Metadata-Version: 1.1
   Name: PySide
   Version: 1.2.2
   EOF"

Installing Python 3 environment
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Download and install `MacPorts <http://www.macports.org>`__.
Before continuing, it is recommended to run

.. code-block:: bash

   sudo port selfupdate

Then install the dependencies

.. code-block:: bash

   sudo port install python36 py36-zmq py36-ply py36-pyside py36-pyodbc py36-psutil py36-spyder py36-sphinx py36-pandas py36-numpy py36-scipy py36-matplotlib py36-pyflakes py36-pylint py36-ipython py36-h5py py36-lxml py36-xlwt py36-xlrd py36-pip graphviz py36-XlsxWriter py36-scikit-learn py36-scikit-image py36-mock


Make PySide detected as installed to avoid rebuild.

.. code-block:: bash

   sudo bash -c "cat >/opt/local/Library/Frameworks/Python.framework/Versions/3.6/lib/python3.6/site-packages/PySide-1.2.2.egg-info <<EOF
   Metadata-Version: 1.1
   Name: PySide
   Version: 1.2.2
   EOF"


Common installation steps
~~~~~~~~~~~~~~~~~~~~~~~~~

Finally we can download and install the Sympathy python "wheel" file, see
:ref:`whl_install_unix`.

Now we are ready to run Sympathy! See :ref:`whl_run_unix`.


For Linux
---------
These installation instructions have been written for Ubuntu 16.04
which is the only officially supported Linux distribution for Sympathy
for Data. Nonetheless, these instructions should also serve as a
starting point for later versions of Ubuntu or other Linux
distributions.

Before you start either installation, make sure that your computer is
internet connected and has the latest version of all packages. If
unsure, run the commands:

.. code-block:: bash

   sudo apt-get update
   sudo apt-get dist-upgrade


You can install Sympathy either for Python 3 (recommended) or
Python 2.7. The common installation steps are needed for both cases.


Installing Python 3 environment
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Start by installing the required prerequisites

.. code-block:: bash

  sudo apt-get install build-essential cmake qt4-default python3-pip python3-pyodbc
  sudo -H pip3 install scikit-image
  sudo -H pip3 install scikit-learn

Since modern Ubuntu has a later version of Python 3 (3.5 or later) not
directly supported by PySide (and it therefore cannot be built like it can for
Python 2) we need to use the version given by the
distribution. If you are installing under a non-supported Linux system
you can try without this step if your ``python3 --version`` shows 3.4
or earlier.

.. code-block:: bash

   sudo apt-get install python3-pyside
   cd /usr/lib/python3/dist-packages
   sudo bash -c "cat >PySide-1.2.2.egg-info <<EOF
   Metadata-Version: 1.1
   Name: PySide
   Version: 1.2.2
   EOF"

Installing Python 2 environment
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Start by installing the required prerequisites

.. code-block:: bash

   sudo apt-get install build-essential cmake qt4-default python-pip python-pyodbc
   sudo -H pip install scikit-image
   sudo -H pip install scikit-learn

As an optional step you can use the distributions version of PySide (a
wrapper library for Qt). If you do not perform this step then the
installation will recompile a fresh version of PySide -- a process
which takes about 20 minutes.

.. code-block:: bash

   sudo apt-get install python-pyside
   cd /usr/lib/python2.7/dist-packages
   sudo bash -c "cat >PySide-1.2.2.egg-info <<EOF
   Metadata-Version: 1.1
   Name: PySide
   Version: 1.2.2
   EOF"


Common installation steps
~~~~~~~~~~~~~~~~~~~~~~~~~

Finally we can download and install the Sympathy python "wheel" file, see
:ref:`whl_install_unix`.

For other Linux distributions than Ubuntu 16.04: if you see any text
in red during the execution of above command, some package may be missing. Read
the part in red and install the required package before trying
again. The preference is always to use a package provided by your
distribution (eg.  ``sudo apt-get install python-xxx``), or at second
hand, one using pip directly (eg. ``sudo -H pip install xxx``).

Now we are ready to run Sympathy! See :ref:`whl_run_unix`.

Linux specific troubleshooting
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

If Sympathy hangs when you try to start it with *python3 -m sympathy_app syg* as
a normal user, then it is possible that you have run it once with *sudo* without
the *-H* flag. This leads to root owning all the cache files. The best way out
of this is to run the following commands:

.. code-block:: bash

  cd
  sudo chown -R MYNAME:MYGROUP .cache/Sympathy\ for\ Data/
  sudo chown -R MYNAME:MYGROUP .local/share/data/Sympathy\ for\ Data/

Where you need to replace *MYNAME* and *MYGROUP* with your username and group
(often the same as the username).  If this does not solve the problem, try
launching Sympathy using the *strace* command

.. code-block:: bash

  strace python3 -m sympathy_app syg

This will make alot of printouts of all system calls, you can break with Ctrl-C
and look for any *permission denied* printouts.

Unix
----

This sections applies to Linux and MacOS.


.. _whl_install_unix:

Install Sympathy wheel
~~~~~~~~~~~~~~~~~~~~~~

We can download the Sympathy python wheel file from the
`official homepage <https://www.sympathyfordata.com/>`_.  Assuming
that you have downloaded it as the file "Sympathy-VERSION.whl" you can
install it by running the following commands *from the folder where you
downloaded it*:

In place of python, use the python environment that was installed for use with
Sympathy.  For example, python, python2, python3, python2.7 or python3.6.

.. code-block:: bash

   sudo -H python -m pip install Sympathy-VERSION.whl
   sudo -H python -m sympathy_app install


Note that the last command launches Sympathy graphically as root so
that the installation can be finished. Close without doing anything
else and continue by launching Sympathy as a normal user (see below).
If you have not done so take a look at the :ref:`quick_start` pages in
the documentation.


.. _whl_run_unix:

Running Sympathy
~~~~~~~~~~~~~~~~

You can run Sympathy either with a GUI (first command below), or for data
processing applications in head-less mode (second command)

In place of python, use the python environment that was installed for use with
Sympathy.  For example, python, python2, python3, python2.7 or python3.6.

.. code-block:: bash

  python -m sympathy_app syg
  python -m sympathy_app sy <my workflow>
