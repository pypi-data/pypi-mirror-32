Property Sharing
================

World and Stimulus properties in Shady (including those with single
value) are stored in arrays, which are transferred to the rendering
engine on every frame for drawing. By default, every new `Stimulus`
object you create has a fresh new set of property arrays created for
it, initialized to their default values. However, it is possible to
have multiple stimuli use the exact same arrays in memory for one or
more properties. Sharing properties in this way will cause those
stimuli to be linked until you explicitly unlink them. This allows
your program to be less complex and more CPU-efficient, since you
will only have to change the property value of one stimulus and the
change will affect all the others.

The simplest way to share a property between two stimuli is to
use the shorthand convention of assigning an actual `Stimulus`
instance to the relevant property value of another stimulus::

    # ...
    stim1 = world.Stimulus( position=( 0, 200 ) )
    stim2 = world.Stimulus( position=stim1 )
    # now, any change to either stimulus's position will affect both
    stim1.x = 500
    print( stim2.position )
    # --> 500, 200

This technique is just a shortcut for the `.ShareProperties` method
of the first stimulus::

    # ...
    stim1.ShareProperties('position', stim2)

This more powerful method can be used to share multiple properties at
a time between multiple stimuli::

    # ...
    master_stimulus = world.Stimulus( size=50, y=100 )
    stimuli = [ world.Stimulus( size=50, x=x ) for x in range( -400, 400, 100 ) ]
    master_stimulus.ShareProperties( 'rotation scale', stimuli )

Note that the name `master_stimulus` here is slightly misleading, because
sharing is fully symmetric: any change to the `.rotation`, or `.scale` of
any of the `Stimulus` instances in `stimuli` will affect the remaining
contents of `stimuli` *and* our `master_stimulus`. That said, it may be
useful to designate one stimulus as the master to indicate a convention
that this stimulus should be used to control the others, especially if
you are going to use `dynamic property assignment <MakingPropertiesDynamic>`::

	master_stimulus.rotation = lambda t: t * 20

`.ShareProperties` also allows you to set property values at the same time
as sharing them::

    # ...
    stim1 = world.Stimulus()
    stim2 = world.Stimulus()
    # share color and set that shared color to red
    stim1.ShareProperties( stim2, color=( 1, 0, 0 ) )

If you want a stimulus to stop sharing properties, you can again use the
shorthand of `Stimulus`-instance assignment::

	stim2.color = stim2      # be yourself
	
...which is really a shortcut for `stim2.MakePropertiesIndependent( 'color' )`.
The `MakePropertiesIndependent` method can also simultaneously change the
value(s) of one or more properties as it unlinks them - here's another example::

    # ...
    stim1.ShareProperties( stim2, position=( -100, 200 ), alpha=0.5, scale=2.5 )
    stim1.Set( position=600 )
    print( stim2.position )
    # --> 600, 600
    stim2.MakePropertiesIndependent( scale=7 )
    print( stim1.scale )
    # --> 2.5
    stim2.alpha = 0.3
    print( stim1.alpha )
    # --> 0.3

One final warning: property sharing does **not** work with property index
shortcuts, as two stimuli cannot share just part of a full property array.
If you want to share specific property dimensions such as `.x` or `.red`,
but not the other dimensions of that property, you should use a dynamic
function instead to ensure it is continually updated::

    ### WRONG ###
    # ...
    stim2.x = stim1
    # --> ValueError: x is the name of a shortcut, not a fully-fledged
    #     property - cannot link it across objects

    ### RIGHT ###
    # ...
    stim2.x = lambda t: stim1.x   # (although it comes at a small CPU cost)
