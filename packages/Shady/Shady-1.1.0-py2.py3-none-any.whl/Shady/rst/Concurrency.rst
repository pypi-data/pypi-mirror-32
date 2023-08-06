Concurrency
===========

.. contents:: :local:


Running single-threaded
-----------------------

Here's how you can use Shady in a single-threaded way::

    import Shady

    w = Shady.World( threaded=False )
    # This may (depending on platform) open a window already, but
    # if so it will be inactive.

    s = w.Stimulus( Shady.EXAMPLE_MEDIA.alien1 )
    # create a Stimulus...
    
    s.frame = Shady.Integral( 16 )
    # ...and configure it as desired
    
    @w.AnimationCallback
    def EachFrame( t ):
        # ...  any code you write here will be called on every
        # frame. The callback can have the prototype `f(self, t)`
        # or just `f(t)`, where `t` is time in seconds since the
        # `World` began. Note that each `Stimulus` instance can
        # have its own animation callback too.
        pass

    w.Run()
    # This is a synchronous call - it returns only when the window closes.
    # It renders stimuli dynamically in the window and allows the window to
    # respond to mouse and keyboard activity (with the default event-handler
    # in place, you can press Q or escape to close the window).

In the above example, `World` construction, rendering, and all animation and
event-handling callbacks in the main thread.  You should not try to type the
above commands line-by-line into an interactive prompt, because the first
line may (on some platforms) create a frozen full-screen window that may
then obscure your console window and, because it is not processing events,
may not respond to your attempts to alt-tab away from it.

A slightly different way to organize the above would be to put the
stimulus-initialization code in the `Prepare()` method of a `World`
subclass::

    import Shady

    class MyWorld( Shady.World ):

        def Prepare( self, walkingSpeed=20 ):
            self.Stimulus(
                Shady.EXAMPLE_MEDIA.alien1,
                frame = Shady.Integral( walkingSpeed ),
            )
        
        def Animate( self, t ):
            # ... the `.Animate()` method will be used as the
            # animation callback unless you replace it using the 
            # `@w.AnimationCallback` decorator or (equivalently) the
            # `w.SetAnimationCallback()` method.
            pass

    w = MyWorld( threaded=False, walkingSpeed=16 )
    # The `walkingSpeed` argument, unrecognized by the constructor, is
    # simply passed through to the `.Prepare()` method (the prototype for
    # which may have any arguments you like after `self`).

    w.Run()
    # As before, because the `World` was created with `threaded=False`,
    # the window will be inactive until you call `.Run()`

Running the Shady engine in a background thread (Windows only)
--------------------------------------------------------------

The following has worked nicely for us on Windows systems::

    import Shady
    w = Shady.World()   # threaded=True is the default
    # the `World` starts rendering and processing events immediately,
    # in a background thread

    w.Stimulus( Shady.EXAMPLE_MEDIA.alien1, frame=Shady.Integral( 16 ) )
    # thread-sensitive operations like this are automatically deferred
    # and will be called in the `World`'s rendering thread at the end
    # of the next frame.

    @w.AnimationCallback
    def DoSomething( self, t ):
        # ... you can set the animation callback as before
        # (with or without `self`)
        pass

In this case, a synchronous call to `w.Run()` is optional: all that would do
is cause your main thread to sleep until the `World` has finished.

This relies on using the binary "ShaDyLib" :doc:`accelerator <Accelerator>` as the `Shady.Rendering.BackEnd()`.
Without the accelerator (using, for example, `pyglet` as the back-end) you
may find that some functionality (such as event handling) does not work
properly when `the Shady.World` is in a background thread.

It also relies on Windows.  On other platforms, the graphical toolkit
GLFW, which underlies the ShaDyLib windowing back-end, insists on being in
the main thread (nearly all windowing/GUI toolboxes seem to do this,
myopically enough).   If you try to create a `Shady.World` on non-Windows
platforms without saying `threaded=False`, it will automatically revert
to `threaded=False` and issue a warning, together with a reminder that
you will have to call `.Run()` explicitly.  Unless, of course, you
use a sneaky workaround...


Emulating background-thread operation on non-Windows platforms
--------------------------------------------------------------

It is convenient and readable, and especially conducive to *interactive*
construction of a `World` and its stimuli, to be able to say::

    import Shady
    w = Shady.World()
    # ...

and have the `World` immediately start running in a different thread,
while you continue to issue commands from the main thread to update its
content and behavior.  However, as explained above, you can only do
this on Windows: on other platforms, the `World` will only run in the
main thread.

There is a workaround, implemented in the utility function
`Shady.Utilities.RunShadyScript()`, which is used when you start an
interactive session with the `-m Shady` flag::

    python -m Shady

or when you invoke your python script with the same flag::

    python -m Shady my_script.py

(In the latter case the `run` subcommand is assumed by default, so this
is actually a shorthand for::

    python -m Shady run my_script.py

There are other subcommands, such as `demo` which allows you to
run scripts interactively if they are specially formatted, as many
of our :doc:`example scripts <ExampleScripts>` are.)

Starting Python with `-m Shady` (or equivalently, calling
`RunShadyScript()` from within Python) starts an queue of operations
in the main thread, to which thread-sensitive `Shady.World` operations
will automatically be directed. It then redirects everything *else*
(either the interactive shell prompt, or the rest of your script) to
a subsidiary thread.

For many intents and purposes, this is just like starting the
`Shady.World` in a background thread, one of the main advantages
being that it allows you to build and test your `World` interactively
on the command line.  It has its limitations, however. For one thing,
you can only create one `World` per session this way (on Windows, you
can create one threaded `World` after another, and can even have two
running at the same time---although I have no data and only pessimistic
suspicions about the performance of the latter idea).  The fun also
comes to a crashing end when you to try do something else that requires
a solipsistic graphical toolbox, like plot a matplotlib graph.


Limitations on multi-threaded performance in Python
---------------------------------------------------

So far, we have found that our multi-threaded `Shady` applications
have generally worked well on Windows. This is largely because
most of the rendering effort is performed on the GPU, and most
of the remaining CPU work is carried out (at least by default
if you have the ShaDyLib :doc:`accelerator <Accelerator>`) in compiled C++ code
rather than Python. Very very little is actually done in Python on
each frame.

However, as soon as your Python code (animation callbacks, dynamic
property assignments, and event handlers) reaches a certain critical
level of complexity, you should be aware of the possibility that
Python itself may cause multi-threaded performance to be significantly
worse than single-threaded. This is because the Python interpreter
itself cannot run in more than one thread at a time, and multi-threading
is actually achieved by deliberately, cooperatively switching between
threads at (approximately) regular intervals, mutexing the entire
Python interpreter and saving/restoring its state on each switch. This
is Python's notorious Global Interpreter Lock or GIL, and a lot has been
written/ranted about it on the Internet, so we will not go into the
details here.  Just be aware that it exists, and that consequently it is
often better to divide concurrent operations between *processes* (e.g.
using the standard `multiprocessing` module) rather than between threads.
You might decide to design your system such that all your `Shady` stuff,
and *only* your `Shady` stuff, runs in a single dedicated process. That
process would then use the tools in `multiprocessing`, or other
inter-process communication methods, to talk to the other parts of the
system.
