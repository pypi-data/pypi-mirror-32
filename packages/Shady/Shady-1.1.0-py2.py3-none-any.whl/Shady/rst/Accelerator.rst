The Shady "Accelerator" (`ShaDyLib`)
====================================

Shady works by harnessing the graphics processor (GPU) to perform, in parallel, most
of the pixel-by-pixel operations entailed in signal generation, contrast modulation,
windowing, linearization and dithering.  For most stimulus arrangements, this leaves
relatively little for the CPU to do on each frame: it just has to issue the OpenGL
commands that clear the screen and then, for each stimulus in turn, transfer a small
set of variable values from CPU to GPU.  Nonetheless this may amount to a few hundred
separate operations per frame---we'll call these the "CPU housekeeping" operations
(and we'll consider them separate from the optional further computations that can be
performed between frames to :doc:`animate <MakingPropertiesDynamic>` stimuli.

Early versions of Shady used the third-party package `pyglet <https://pypi.org/project/pyglet/>`_ for two purposes:
first, to initialize the windowing system, open a window and thereby create an OpenGL
context for a new `World`;  and second, to provide Python wrappers for the OpenGL calls
that are required for initialization and frame-to-frame operation. Thus, we were able
to implement the CPU housekeeping operations in Python (specifically, in the
`Shady.PyEngine` sub-module). This worked well... except when it didn't.  

The problem is that Python, being a high-level dynamic interpreted language, is
inefficient for performing large numbers of simple operations. Relative to the equivalent
operations compiled from, say C or C++, Python not only requires extra time but,
critically, adds a large amount of *variability* to the time taken when running
on a modern operating system like Windows.  DOC-TODO

DOC-TODO