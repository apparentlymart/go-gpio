// Package gpio provides generic interfaces for interacting with GPIO
// pins. It doesn't actually have any implementation of GPIO usage for
// any particular hardware; other packages (such as linuxgpio) will
// provide various implementations of these interfaces.
package gpio

// Value is an enumeration type for the two possible GPIO values Low and High.
type Value int

// Direction is an enumeration type for the two possible GPIO directions
// In and Out.
type Direction int

// EdgeSensitivity is an enumeration type for different events that a
// Waiter can monitor for.
type EdgeSensitivity int

const (
	// Low is a GPIO Value that sets the GPIO pin to logic low.
	Low Value = 0

	// High is a GPIO Value that sets the GPIO pin to logic high.
	High Value = 1
)

const (
	// In is a GPIO direction for reading data (measuring a signal).
	In Direction = 0

	// Out is a GPIO direction for writing data (asserting a signal).
	Out Direction = 1
)

const (
	// NoEdges disables all triggers on a Waiter, meaning that
	// it would wait indefinitely.
	NoEdges EdgeSensitivity = 0

	// RisingEdge makes a Waiter sensitive to changes from Low to High.
	RisingEdge EdgeSensitivity = 1

	// FallingEdge makes a Waiter sensitive to changes from High to Low
	FallingEdge EdgeSensitivity = 2

	// BothEdges makes a Waiter sensitive to changes in either direction.
	BothEdges EdgeSensitivity = 3
)

// A ValueSetter can have a GPIO value set on it.
type ValueSetter interface {
	SetValue(value Value) (err error)
}

// A ValueGetter can have a GPIO value read from it.
type ValueGetter interface {
	Value() (value Value, err error)
}

// A DirectionSetter can have a GPIO data direction set on it.
type DirectionSetter interface {
	SetDirection(direction Direction) (err error)
}

// A PullStopper can disable a pull-up or pull-down resistor.
type PullStopper interface {
	StopPulling() (err error)
}

// An UpPuller can enable a pull-up resistor (which, if it is also a
// DownPuller, may implicitly stop pulling down.)
type UpPuller interface {
	PullUp() (err error)
}

// A DownPuller can enable a pull-down resistor (which, if it is also an
// UpPuller, may implicitly stop pulling up.)
type DownPuller interface {
	PullDown() (err error)
}

// A Puller can enable and disable pull-up and pull-down resistors.
type Puller interface {
	UpPuller
	DownPuller
	PullStopper
}

// A GpioPin can do the most commonly-available operations available for GPIO
// pins: set a signal direction, and then set a value or read a value depending
// on the chosen direction.
//
// This aggregate interface is provided to describe the return type of
// functions that instantiate GpioPins, but hardware drivers that consume
// GPIO pins and yet don't need bidirectional access (which is the common case)
// should depend directly on ValueGetter or ValueSetter to illustrate clearly
// to the user how the GPIO pin will be used by the driver, and thus help avoid
// unusual situations like two devices trying to drive the same signal.
type GpioPin interface {
	ValueGetter
	ValueSetter
	DirectionSetter
}

// A EdgeWaiter can block until its value changes.
type EdgeWaiter interface {

	// SetSensitivity specifies what directions of change the waiter will
	// watch for. Set this to something other than NoEdges before calling
	// WaitForEdge, or WaitForEdge will never return.
	SetSensitivity(sensitivity EdgeSensitivity) (err error)

	// WaitForEdge blocks until the selected change (chosen by SetSensitivity)
	// has occured.
	WaitForEdge() (err error)
}

func (value Value) String() string {
	switch value {
	case Low:
		return "Low"
	case High:
		return "High"
	default:
		panic("Cannot String Invalid Value")
	}
}

func (direction Direction) String() string {
	switch direction {
	case In:
		return "In"
	case Out:
		return "Out"
	default:
		panic("Cannot String Invalid Direction")
	}
}

func (sensitivity EdgeSensitivity) String() string {
	switch sensitivity {
	case NoEdges:
		return "NoEdges"
	case BothEdges:
		return "BothEdges"
	case RisingEdge:
		return "RisingEdge"
	case FallingEdge:
		return "FallingEdge"
	default:
		panic("Cannot String Invalid EdgeSensitivity")
	}
}
