
MR SPRINT
=========

**Magnetic Resonance Experiment Simulator and Visualization Tool**

MRSPRINT is a visual magnetic resonance simulator where you can simulate
a magnetic resonance experiment - spectroscopy or imaging. Its main goal is
provide an education tool that "help" the student/staff to understand,
interpret and explore magnetic resonance phenomena.

This tool is totally free (see license), but if you are making use of it,
you need to cite us using both citations. This is very important to us.

Article:
Software:

What you can see
----------------


* Precession: spins precessing in static magnetic field;
* Resonance: resonance when an RF pulse is applied;
* Contrast: T1, T2, and density of spins;
* Field inhomogeneity: isochromates can be shown with their dispersion;
* Evolution: magnetization with intensity, frequency and phase.
* FID: free induction decay, the signal;
* Echo: spin or gradient echo (rephasing/dephasing);

Main screens
------------

Context editors are specialized to edit each step of experiment, see below.


.. image:: screenshots/context-menu.png
   :align: center
   :alt: Contexts Menu


Sample
^^^^^^

You can open and/or create samples to run with MR experiment. Each sample
element have three characteristics: t1,t2 and density of spins.


.. image:: screenshots/sample-screen.png
   :align: center
   :alt: Sample Screen


System
^^^^^^

Here you can set static magnetic field and add inhomogeneity to it.


.. image:: screenshots/system-screen.png
   :align: center
   :alt: System Screen


Sequence
^^^^^^^^

You can choose a pulse sequence for your experiment, that includes the RF
and gradient pulses. The sequence is programmed in Python at this moment.


.. image:: screenshots/sequence-screen.png
   :align: center
   :alt: Sequence Screen


Simulator
^^^^^^^^^

At this point you can set simulation mode and other details about
simulation , e.g. time resolution.


.. image:: screenshots/simulator-screen.png
   :align: center
   :alt: Simulator Screen


Processing
^^^^^^^^^^

Finally you can process your data to plot the spectrum, imaging or
other features.


.. image:: screenshots/processing-screen.png
   :align: center
   :alt: Processing Screen


2D Editor
---------

This editor provides a table that represents the current selected
slice of the 3D view. In this table you can edit the values of each
element property(ies). Colors also help you to pre visualise the intensity
of chosen value.

3D View
-------

3D view is used to show contexts objects such as sample, magnet field
inhomogeneity, and the evolution of magnetization. You can move the cam
to adjust the perspective.

Data
----

Data can be saved using HDF5 files. For each context we adopted an
extension to distinguish from each other. HDF5 files can be easily
modified with other external tool if necessary.

Future planned features
-----------------------


* K-Space visualisation;
* Graphical sequence editor;
* Chemical interactions;
* Flux (spins not fixed in positions).

A brief history of this universe
--------------------------------

Being short, when difficulties appears to understand magnetic
resonance phenomena, we will search tools that could help us.
At that time we have found nice explanations on
`Brian site <http://mrsrl.stanford.edu/~brian/bloch/>`_.
As we are fan of Python, we discovered this implementation from
`Neji <https://github.com/neji49/bloch-simulator-python>`_. Yet, it was
difficult to change parameters and create new graphics.
So, this work starts. From that, many ideas have shining, and, here we are.

The main developer is PhD Daniel C. Pizetta, at University of São Paulo,
São Carlos, Brazil. A lot of work was developed by bachelor Victor M. Souza
and now Clara Vidor is continuing the development. We also have a head
professor PhD Fernando Fernandes Paiva.

Download binaries - click-and-run
---------------------------------

Installing from PyPI - stable, end-user
---------------------------------------

Simple run

``$ pip install mrsprint``

This must install all necessary dependencies then the code.

Installing from code - up to date, developers
---------------------------------------------

Clone from GitLab the `latest code <http://gitlab.com/dpizeetta/mrsimulator>`_.
Enter into the folder containing ``setup.py`` file.

``$ cd mrsprint``

Run the install command using developer mode option (\ ``-e``\ )

``$ pip install -e .``

This command will install all necessary dependencies and the code will be
installed in the same place it is, so it is easy to update doing ``git pull``.

Dependencies
------------


* NumPy: Numerical mathematical library;
* SciPy: Scientific library;
* NMRGlue: NMR processing library;
* PyQtGraph: Data visualization library;
* PyQt/Pyside: Graphical framework.

Problems on Windows
-------------------

Windows users may have some problem with some C necessary extensions. We
recommend to visit `Windows Compilers <https://wiki.python.org/moin/WindowsCompilers>`_\ ,
but the first trial should be installing `Visual Studio Build Tools <http://landinghub.visualstudio.com/visual-cpp-build-tools>`_. Sometimes this
is necessary even though using distributions like Anaconda. This include
solutions for vcvarshall.bat missing.
