Making Properties Dynamic
=========================

Individual Properties
---------------------

A key feature of Shady is the ability to make `World` and `Stimulus` properties
dynamic, by setting them to functions of time. This can be done simply by
assigning a function object to the property instead of a static value or array. The
function should take one argument, `t` (for time), and return whatever value
(or array of values) you want that property to have at time `t`. On every frame
callback, Shady will run any dynamic property functions you have assigned
using the current world time and update the value of the corresponding properties.

Note that there are multiple ways to create callable objects in Python. The most
flexible way is to define a function using `def`::

    def simple_acceleration( t ):
        return t ** 2
		
    stim.x = simple_acceleration   # note no parentheses - we are assigning the function itself
        
Another way is to specify an anonymous function in-line using `lambda`::

    stim.x = lambda t: t ** 2
    
Both methods are completely valid, but `lambda` functions are restricted to 
containing just the single expression you want to return from the function.

Note that `t` is measured in seconds, and by default it is the number of
seconds elapsed since the `World` first started rendering stimuli. So if it
has been a long time since the `World` started, your `Stimulus` will likely be
off the screen in the above example.  However, each `Stimulus` can have an
independent "time zero", and this can be reset to the current time using the
call::

    stim.ResetClock()

There are no unusual restrictions on your dynamic functions, provided that
they take exactly one argument and return a value or array that is appropriate
for that property. Any Python variables or objects that are accessible in the
same namespace can be used and modified::

    # Horizontal oscillator
    import time
    import numpy
    # ...
    amplitude = 200
    stim.x = lambda t: amplitude * numpy.sin( 2 * numpy.pi * t )
    stim.ResetClock(); time.sleep( 0.5 ); print( stim.x )
    # --> 200   (maximum amplitude)
    amplitude = 350
    print( stim.x )
    amplitude = 350   (no time has elapsed, still at maximum amplitude)
    
Note that changing the `amplitude` variable in the previous example
completely erases all memory of the original dynamic. The stimulus will
act as though it had been oscillating with amplitude 300 since world
time began, and will therefore appear to abruptly jump across the screen
at the moment the dynamic is changed and start somewhere in the middle
of its 300-pixel oscillation.

To make sure the stimulus starts its new oscillation from the middle of
the screen (`x == 0`), we need to *resets its clock*. The time argument
`t` that is used to evaluate a stimulus's dynamic properties is the amount
of time that has passed since that stimulus's clock began. By default,
every stimulus's clock is synchronized to the world's clock (which started
at the moment the `World()` was created), even if the stimulus was created
more recently. To reset the stimulus clock to zero, use `ResetClock()`::

    import time
    import Shady
    world = Shady.World()
    time.sleep( 50 )
    stim = Shady.Stimulus( ... )
    stim.x = lambda t: 10 * t   # stimulus will move to the right at ten pixels/second
    print( stim.position )
    # --> 500, 0
    stim.ResetClock(); print( stim.position )
    # --> 0, 0
    time.sleep( 20 ); print( stim.position )
    # --> 200, 0

Dynamic functions do not even need to use the time variable. You can make
the properties of your stimulus dependent on whatever variables you want::

    import Shady
    world = Shady.World( width=800, height=600 )
    stimulus = world.Stimulus()
    def grow_to_the_right( _ ):
        normalized_x = stim.x / world.width + ( world.width / 2 )   # == 0 on left, == 1 on right 
        return 10 * normalized_x
    stim.scale = grow_to_the_right
    stim.x = -200
    print( stim.scale )
    # --> 2.5
    stim.x = 200
    print( stim.scale )
    # --> 7.5  
    
The function `grow_to_the_right` in this example inspects the stimulus's
`x`-coordinate on every frame and uses it compute its scale, independently
of time. (Note that we still have to give our dynamic function exactly one
argument so that Shady can pass in the stimulus's clock, but we use an underscore
`_` to indicate that this argument is not used.)

As these examples have shown, whenever you attempt to access a managed 
property, its current static value or array of values will be returned, 
even if the property is dynamic. If you want to retrieve the actual function 
object being used to calculate its dynamics, use the `Shady.Stimulus.GetDynamic`
method::

    import time
    # ...
    stim.alpha = lambda t: ( 100 - t ) / 100   # stimulus will fade out over 100 seconds
    stim.ResetClock(); time.sleep( 30 ); print( stim.alpha )
    # --> 0.7
    print( stim.GetDynamic( 'alpha' ) )
    # --> <function <lambda> at 0x00........>
    stim.color = stim.alpha   
    # this will set the color to a static mid-gray value of 0.7
    print( stim.GetDynamic( 'color' ) )
    # --> None
    stim.color = stim.GetDynamic( 'alpha' )  
    # this will make the stimulus brighten at the same rate as it fades out
    print( stim.GetDynamic( 'color' ) )
    # --> <function <lambda> at 0x00........>

In general, you can set a Shady property to any callable object that takes
exactly one argument. This includes any instance of a class with a `__call__`
method defined, provided the call takes one argument. The optional `Shady.Dynamics`
submodule offers several useful classes designed to be used as dynamic properties
in Shady, such as `integrating values over time <Integral>` and `smoothly transitioning
between a start and end value <Transition>`.

NOTE: Be wary when using the same public variable to define multiple dynamics functions
in a row. Because of how functions interact with their namespace in Python, the
*current* (i.e. last set) value of that variable will be used when the dynamics are
evaluated on each Shady frame callback. This includes simple looping variables! If you
want to 'freeze' the value of a public variable when defining a dynamic function, you
will need to separate it from that variable's namespace, e.g. by using a nested function::

    ### WRONG ###
    import numpy
    import Shady
    world = Shady.World()
    stimuli = []
    amplitudes = [100, 200, 300]
    for a in amplitudes:
        stim = world.Stimulus()
        stim.x = lambda t: a * numpy.sin( 2 * numpy.pi * t )
        stimuli.append( stimulus )
    # all three stimuli will use a == 300 when their dynamics are evaluated!

    ### WRONG ###
    import numpy
    import Shady
    world = Shady.World()
    stimuli = []
    amplitudes = [100, 200, 300]
    for i in range( 3 ):
        stim = world.Stimulus()
        stim.x = lambda t: amplitudes[i] * numpy.sin( 2 * numpy.pi * t )
        stimuli.append( stimulus )
    # all three stimuli will use i == 2, i.e. amplitudes[2]!
    
    ### RIGHT ###
    import numpy
    import Shady
    
    def create_oscillation_dynamic( amplitude )
        # the argument `amplitude` is 'baked in' to each dynamic
        return lambda t: amplitude * numpy.sin( 2 * numpy.pi * t )
    
    world = Shady.World()
    stimuli = []
    amplitudes = [100, 200, 300]
    for amplitude in amplitudes:
        stim = world.Stimulus()
        stim.x = create_oscillation_dynamic( amplitude )
        stimuli.append( stim )

Also note that properties of your `World` instance can be made
dynamic using all of the methods described above. For example, to
create a world whose background color oscillates between black and
white::

    import numpy
    import Shady
    world = Shady.World( clearColor=lambda t: 0.5 + 0.5 * numpy.sin( 2 * numpy.pi * t ) )

The world's dynamics will be updated before any of the stimuli it contains,
and its stimuli are updated according to their draw order (i.e. `.z`).
Stimuli with the same `z`-value will be drawn in the order they were
created.

The Animate Method
------------------

As the behavior of your stimulus grows more complex and its
properties become more interdependent, you may begin to find that relying
on individual property dynamics becomes unwieldy. In this case, you will
likely want to use the stimulus's `Animate()` method, which is evaluated
before any property dynamics on each Shady frame callback.

The only practical difference between the `Animate()` method and
any dynamic properties is that `Animate()` takes a `self` argument,
which makes it easier to refer to the stimulus in your logic (e.g.
for checking and modifying its state). The function does not need
to return any value, which means that you will most likely want to
create it using the standard `def`. Once created, pass the function
object to the `.SetAnimationCallback()` method to properly bind it to
the stimulus::

    import time
    import numpy
    # ...
    is_bouncing = False
    time_started_bouncing = None

    def bounce( self, t ):
        if is_bouncing:
            if time_started_bouncing is None:
                time_started_bouncing = t
                # Note use of `_t` in the lambda to distinguish it from the bounce() argument `t`.
                self.y = lambda _t: 100 * numpy.abs(
                    numpy.sin( 2 * numpy.pi * (_t - time_started_bouncing ) )
        else:
            if time_started_bouncing is not None:
                time_started_bouncing = None
                self.y = 0

    stim.SetAnimationCallback( bounce )   # again, note that function object is assigned
    is_bouncing = True   # stimulus starts bouncing
    time.sleep( 5 )
    is_bouncing = False   # stimulus stops bouncing
    time.sleep( 5 )

This example is a little more complex than any of the examples in
the previous section, but that's exactly why the `Animate()` method
is useful. The `bounce()` function assigns a bouncing dynamic to
the stimulus's y-coordinate whenever `is_bouncing` is set to `True`,
making sure that the stimulus only starts bouncing at that moment.
It abruptly resets the y-coordinate to zero whenever `is_bouncing`
is set to False. (The optional `Shady.Dynamics` submodule contains a
`StateMachine` class that makes it easier to switch your stimuli
between different modes of behavior like this.)

If your animation callback has two arguments (i.e. a `self` as well
as just a `t`) then you *must* use the `.SetAnimationCallback()` helper
to properly bind your function as the `.Animate()` method of the
instance, so that Python knows that the Stimulus instance should be
passed in as the `self` argument. The following will **not** work::

    ### WRONG ###
    # ...
    stim.Animate = bounce

If your callback has only one argument, it is interpreted as time
`t`---in this case, you can use `.SetAnimationCallback()` or just
directly assign `stim.Animate = func`.

As with dynamics, instances of the `World` class can have an
`.Animate()` method set in the exact same way as instances of
the `Stimulus` class.

Note that that `Stimulus` and `World` instances provide have an
attribute `AnimationCallback` which can be used as a decorator,
as a syntactic alternative to calling `.SetAnimationCallback()`::

	@stim.AnimationCallback
	def bounce( self, t ):
		# ...

Order of Dynamic Evaluations
----------------------------

Shady evaluates property dynamics and `Animate()` methods in the
following order on each frame:

    1. `World.Animate()`
    
    2. `World` dynamic properties
    
    3. Each `Stimulus` (sorted first by `.z` and second by time of
       creation:
       
          a. Stimulus.Animate()
          b. Stimulus dynamic properties

For each `World` or `Stimulus` instance, the dynamics are evaluated
in a fixed order that may seem arbitrary---it is not recommended to
make dynamic properties that use the values of other dynamic properties,
thereby relying on an assumption that certain properties are evaluated
before others in a given frame. If you need to do this, you can use
the `Animate()` method to set them procedurally in the order you need
them calculated.
