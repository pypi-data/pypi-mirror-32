Installing Shady
================

The programmer's interface to Shady is in Python.  Python is a wonderfully clean, 
powerful, easy-to-program, easy-to-read language. But it has its problems.  The main 
difficulty, for most non-trivial pieces of Python software, is getting off the ground in 
the first place:  particularly in scientific application frameworks, the hardest program
to write is  often "hello world". If you have worked a lot with Python, you'll probably
identify with the xkcd.com cartoon from 2018-04-30:

.. image:: https://imgs.xkcd.com/comics/python_environment.png
   :alt: http://xkcd.com/1987
	
If you have (some version of) this problem, we cannot pretend to offer the solution here,
but we do sympathize. What we can do is describe relatively-simple setup procedures that
have worked for us.

Installation TL;DR
------------------

#. Install an `Anaconda <https://www.anaconda.com/download/>`_ distribution of Python.
   We recommend Python 3, unless you already know a particular reason why you'll need 
   Python 2---in any case, Shady should be fine with either 2.7 or 3.4+.
  
#. Do whatever it takes to ensure that, when you type `python` at a system command prompt, 
   the version of Python that starts is the one with which you intend to use Shady.
   Whenever we give you a command line that starts with `python`, we'll assume this is
   already straightened out. It's worth taking some time to make sure this happens when
   you want it to (and not when you don't---it's possible that you'll need to keep
   a previous installation of Python as the default for more general non-scientific
   purposes).  Depending on how you want to manage your system, this may involve one
   of the following:
   
       * on a session-by-session basis, calling the `activate` script provided by
         Anaconda before doing anything else;
       * manipulating the `PATH` environment variable so that it includes the parent
         directory of Anaconda's `python` or `python.exe` binary: this can be done
         system-wide, or on a user-specific basis, or (on non-Windows systems) in one
         of the startup scripts such as ``~/.bash_profile`` that run only when you begin
         a new shell session; 
       * (on Windows) allowing the Anaconda installer to manipulate the registry such
         that the newly installed Python is the default Python system-wide;
       * your own custom solution.
   
   Windows example (from the Command Prompt, assuming you have chosen to install Anaconda
   in `C:\Anaconda`):
   
   .. code-block:: batch
   
       call "c:\Anaconda\Scripts\activate.bat"
    
   Non-Windows example (from the Terminal application, assuming you're using the `bash`
   shell and have installed Anaconda in ``~/anaconda``):
   
   .. code-block:: bash
   
       source ~/anaconda/bin/activate
   
#. If you didn't want to install Anaconda's full (very large) default set of third-party
   packages, you can start with `the minimal "miniconda" subset <https://conda.io/miniconda.html>`_  and 
   then use the `conda` package manager to install the following recommended packages:
  
    - `numpy <https://pypi.org/project/numpy/>`_
    - `pillow <https://pypi.org/project/pillow/>`_
    - `matplotlib <https://pypi.org/project/matplotlib/>`_
    - `ipython <https://pypi.org/project/ipython/>`_
    
   We do it as follows:  first install miniconda; then open the Command Prompt (on
   Windows) or Terminal (MacOS); then call the `activate` script; then install::
  
       python -m conda install numpy pillow matplotlib ipython

   If you're using some other (non-Anaconda) installation of Python you can install
   these via `pip` instead of `conda`.
  
#. An additional nice-to-have package is `opencv-python <https://pypi.org/project/opencv-python/>`_ which is not yet (at the time of
   writing) managed by `conda`, but which can instead be installed via `pip`.
  
       python -m pip install opencv-python

#. Install Shady itself, again via `pip`::

      python -m pip install shady

#. Test the various features of Shady.  The following command runs the example script
   :doc:`examples/showcase.py <examples_showcase>` interactively::
   
      python -m Shady demo showcase

Hardware and Operating System Compatibility
-------------------------------------------

Shady was conceived to play a modular role in larger, more complex multi-modal neuroscience
applications that may include novel human interface devices and/or specialized neuroscientific
equipment, such as eye-trackers and EEG amplifiers. Manufacturers of such equipment are
overwhelmingly more likely to support Windows than anything else.

Hence, the Windows platform is where we aim to optimize performance, and most of our experience
in doing so has been with Windows 10. On other platforms, it should not be any harder to get
Shady running, but it may be harder to get it to perform really well, so we had better describe
our support for non-Windows platforms as experimental. However, our experiences so far (with
macos 10.8 through 10.13, and with Ubuntu 14 in a VirtualBox) indicate that both the C++ code
and CMake files for the accelerator, and the Python code of the rest of the module, are
cross-platform compatible.

Shady probably will not work on big-endian hardware. Since most commercial CPUs are little-
endian, at least by default, we have had no opportunity to test it on big-endian systems and
little interest in doing so.


Known Issues
------------

Non-forward-compatible shader code (Mac):

	Our random-number generator (for additive noise and for dithering) is of poorer
	quality on the Mac.  The reason is as follows: our shader code, written in OpenGL
	Shading Language (GLSL) is backwardly compatible with old legacy versions of the
	language (GLSL 1.2) corresponding to OpenGL 2.1). However, we use one or two features
	from later versions (GLSL 3.3+, corresponding to OpenGL 3.3+) when they are available,
	and these features allow us to improve the quality of the random number generator. On
	our Windows systems this has worked just fine: legacy GLSL can be mixed with newer
	features.  But on macos this is not allowed: one has to choose either old or new GLSL,
	and cannot mix features from one while remaining compatible with the other. For
	historical reasons (`pyglet` compatibility, in the absence of our binary accelerator),
	we have stuck with the old version.  In future releases we intend to migrate to
	modern OpenGL/GLSL, to ensure compatibility with future graphics cards that may drop
	legacy GLSL support.


Software Prerequisites
----------------------

Scientific software packages in Python have an unfortunate tendency to rely on a
"house of cards" made up of specific versions of other third-party packages. Somewhere in
the hierarchy of dependencies, sooner or later you end up locked into a legacy version of
something you don't want. With this in mind, we limited Shady's dependencies to a small
number of well established, very widely used, and actively developed general-purpose
packages. We test it with 5-year-old versions of its principal dependencies as well as
current versions. We aim to ensure that Shady's functionality degrades gracefully even in
their absence.

Shady supports Python versions 2 and 3 (specifically CPython, which is the standard, most
prevalent implementation). Shady doesn't have *hard* dependencies on third-party packages
beyond that.  On any CPython implementation of version 2.7.x, or 3.4 and up, some of
Shady's core functionality should be available. This claim comes with two caveats.

The first caveat is that we are assuming availability of the ShaDyLib :doc:`accelerator <Accelerator>`
which is a compiled binary (dynamic library).  Compiled binaries for 32-bit Windows,
64-bit Windows and 64-bit MacOS are bundled as part of the Shady distribution. If you are
using a different OS (e.g. some flavour of Linux) or if the dynamic library fails to load
for any reason (some form of dynamic-library dependency hell, no doubt) then you may need
to (re)compile the accelerator.  *Without* the accelerator, you can still run Shady, but:

- you will need to install another third-party package, `pyglet <https://pypi.org/project/pyglet>`_, which can be done
  via `pip`::

      python -m pip install pyglet

- and real-time performance will likely be very poor (perhaps only sporadically, perhaps
  consistently) and you may also find that some things are not supported (e.g. Mac Retina
  displays). Either way, this route is not recommended if you need production quality. 
	   
The second caveat is that there are a few *recommended* packages, without which Shady's
functionality is relatively limited. Without any third-party packages, you can display
rectangular or oval patches, with or without a customizable 2-D carrier signal function,
a 2-D contrast-modulation function, a spatial windowing function, colour, and dynamic
pixel noise. You will also have :doc:`powerful tools <MakingPropertiesDynamic>` for governing the 
way these properties change over time. The following third-party packages, if available,
add specific extra types of functionality over and above the core:

	`numpy <https://pypi.org/project/numpy>`_:
	
		If you have `numpy`, you will also be able to create or render arbitrary textures
		which you have defined as pixel arrays.  There are also subtler (but potentially
		very powerful) advantages, such as the fact that dynamic objects like the
		`Shady.Dynamics.Integral` can be multi-dimensional (an example in which this is
		leveraged to the max is :doc:`examples/dots4.py <examples_dots4>`).

	`pillow <https://pypi.org/project/pillow>`_:
	
		With `numpy` **and** `pillow` working together, you will also be able to:
	
			- load your texture data from common image formats on disk,
			- save screen capture data to disk in common image formats, and
			- render text stimuli, in a monospace font.

	`matplotlib <https://pypi.org/project/matplotlib>`_:

		With `numpy` and `pillow` **and** `matplotlib` working together, you should be
		able to render text stimuli in any of the fonts installed on your system.
		Furthermore, `matplotlib` enables you to plot timing diagnostics and image
		histograms.

	`opencv-python <https://pypi.org/project/opencv-python>`_:

		With the `cv2` module (part of the `opencv-python` package) working together 
		with `numpy`, you will be able to save stimulus sequences as movies, and also
		display stimuli from video files or live camera feeds.

	`ipython <https://pypi.org/project/ipython>`_:

		IPython is not integral to Shady, but it provides great advantages for
		interactive configuration of Shady stimuli. By improving the user experience
		at the command prompt (with things like tab completion, dynamic object
		introspection, and cross-session command history) IPython makes it much easier
		to explore and understand Shady interactively.
		
