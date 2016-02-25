// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package input

import "github.com/veandco/go-sdl2/sdl"

type Type int        // A HGE Input Event type constants
type Key int         // A HGE Virtual-key code
type Flag sdl.Keymod // HGE Input Event flags (multiple ones may be OR'd)
type Button int      // A HGE Input Mouse button

// HGE Input Event structure
type InputEvent struct {
	Type    Type    // event type
	Key     Key     // key code
	Flags   Flag    // event flags
	Chr     uint32  // character code
	Button  Button  // Mouse Button
	Wheel   int     // wheel shift
	X       float64 // mouse cursor x-coordinate
	Y       float64 // mouse cursor y-coordinate
	cleared bool    // true if there is no useful data here
}

var (
	keys       [last_key]bool
	keySym     Key
	lastKeySym Key
	mb         [3]bool
	mm         Mouse
	event      InputEvent
	lastEvent  InputEvent
)

// Process events
// Called automatically by hge.Run()
func Process() {
	e := sdl.PollEvent()
	// 	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
	switch e := e.(type) {
	case *sdl.KeyDownEvent:
		keys[e.Keysym.Sym] = 1 == e.State
		keySym = Key(e.Keysym.Sym)

		event.Type = Type(e.Type)
		event.Key = Key(keySym)
		event.Flags = Flag(e.Keysym.Mod)
		event.Chr = uint32(e.Keysym.Scancode)

		if event.Key == lastEvent.Key {
			event.Flags |= INP_REPEAT
		}

	case *sdl.MouseMotionEvent:
		event.Type = Type(e.Type)
		event.X = float64(e.X)
		event.Y = float64(e.Y)

	case *sdl.MouseButtonEvent:
		event.Type = Type(e.Type)
		event.X = float64(e.X)
		event.Y = float64(e.Y)
		if e.Button == sdl.BUTTON_LEFT {
			mb[0] = e.State == 1
			event.Button = Button(e.Button)
		} else if e.Button == sdl.BUTTON_MIDDLE {
			mb[1] = e.State == 1
			event.Button = Button(e.Button)
		} else if e.Button == sdl.BUTTON_RIGHT {
			mb[2] = e.State == 1
			event.Button = Button(e.Button)
		}
	case *sdl.MouseWheelEvent:
		event.Type = Type(e.Type)
		event.Wheel = int(e.Y)
	}
	// 	}
}

// Clear the event queue
// Called automatically by hge.Run()
func ClearQueue() {
	for i := 0; i < last_key; i++ {
		keys[i] = false
	}

	lastEvent = event
	lastKeySym = keySym
	event = InputEvent{cleared: true}
	keySym = 0
}

// Update the internal mouse structure
// Called automatically by hge.Run()
func UpdateMouse() {
	if !event.cleared {
		mm.X = event.X
		mm.Y = event.Y
	}

	mm.Wheel = event.Wheel
}

// Initialize the event structure
// Called automatically in hge.Initialize()
func Initialize() error {
	ClearQueue()
	return nil
}

type Mouse struct {
	X, Y  float64 // The position of the mouse relative to the top left corner
	Wheel int     // The scroll wheel movement: negative down/positive up
	Over  bool    // If the mouse is over the window
}

// Returns the last known position of the mouse relative to the top-left corner
func (m *Mouse) Pos() (x, y float64) {
	return m.X, m.Y
}

// Set the mouse position to (x,y)
func (m Mouse) SetPos(x, y float64) {
	m.X, m.Y = x, y
	mm.X, mm.Y = x, y
	// There's WrapMouseInWindow, but it's a method off of Window
	//sdl.WarpMouse(int(x), int(y))
}

// Returns the movement if there's been any mouse movement since the last time
// the events were updated.
// Returns a positive value is returned if the wheel is scrolled up
// Returns a negative value if it's scrolled down
// Returns zero on no movement
func (m *Mouse) WheelMovement() int {
	return m.Wheel
}

// Returns if the mouse is currently over the window
// Always true when in fullscreen mode
// TODO Currently unimplemented
func (m *Mouse) IsOver() bool {
	return false
}

func (b Button) Down() bool {
	return mb[b]
}

func (b Button) Up() bool {
	return mb[b]
}

// Returns true if there's been a key pressed down since the last time the
// events were updated
func (k Key) Down() bool {
	return keys[k] == true
}

// Returns true if there's been a key pressed down since the last time the
// events were updated
func (k Key) Up() bool {
	return keys[k] == false
}

// The current state of the button
// Returns true if the key is currently pressed, false otherwise
func (k Key) State() bool {
	ks := sdl.GetKeyboardState()
	return ks[k] == 1
}

// Return the key name
func (k Key) Name() string {
	return sdl.GetKeyName(sdl.Keycode(k))
}

// Get the last polled key
func GetKey() Key {
	return keySym
}

// Get the key character code
func GetChar() uint8 {
	return uint8(keySym)
}

// Get the event structure from the last time the events where updated
func GetEvent() (e InputEvent, b bool) {
	if event.cleared {
		return event, false
	}

	return event, true
}
