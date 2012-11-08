// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package input

// import "fmt"
import "github.com/banthar/Go-SDL/sdl"

type Type int
type Key int
type Flag sdl.Mod
type Button int

// HGE Input Event structure
type InputEvent struct {
	Type    Type    // event type
	Key     Key     // key code
	Flags   Flag    // event flags
	Chr     uint8   // character code
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
	mb         [4]bool
	mm         Mouse
	event      InputEvent
	lastEvent  InputEvent
)

// Process events
func Process() {
	e := sdl.PollEvent()
	// 	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
	switch e.(type) {
	case *sdl.KeyboardEvent:
		k, _ := e.(*sdl.KeyboardEvent)

		keys[k.Keysym.Sym] = 1 == k.State
		keySym = Key(k.Keysym.Sym)

		event.Type = Type(k.Type)
		event.Key = Key(keySym)
		event.Flags = Flag(k.Keysym.Mod)
		event.Chr = k.Keysym.Scancode

		if event.Key == lastEvent.Key {
			event.Flags |= INP_REPEAT
		}

	case *sdl.MouseMotionEvent:
		m, _ := e.(*sdl.MouseMotionEvent)

		event.Type = Type(m.Type)
		event.X = float64(m.X)
		event.Y = float64(m.Y)

	case *sdl.MouseButtonEvent:
		m, _ := e.(*sdl.MouseButtonEvent)

		event.Type = Type(m.Type)
		if m.Button == sdl.BUTTON_WHEELUP {
			event.Wheel = 1
		} else if m.Button == sdl.BUTTON_WHEELDOWN {
			event.Wheel = -1
		} else {
			event.Button = Button(m.Button)
		}
	}
	// 	}
}

// Clear the event queue
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
func UpdateMouse() {
	if !event.cleared {
		mm.X = event.X
		mm.Y = event.Y
	}

	mm.Wheel = event.Wheel
}

// Initialize the event structure
func Initialize() error {
	ClearQueue()
	return nil
}

type Mouse struct {
	X, Y  float64
	Wheel int
	Over  bool
}

func New() *Mouse {
	return &Mouse{}
}

func (m *Mouse) Pos() (x, y float64) {
	return m.X, m.Y
}

func (m Mouse) SetPos(x, y float64) {
	m.X, m.Y = x, y
	mm.X, mm.Y = x, y
	sdl.WarpMouse(int(x), int(y))
}

func (m *Mouse) WheelMovement() int {
	return m.Wheel
}

func (m *Mouse) IsOver() bool {
	return false
}

func (k Key) Down() bool {
	return keys[k] == true
}

func (k Key) Up() bool {
	return keys[k] == false
}

func (k Key) State() bool {
	ks := sdl.GetKeyState()
	return ks[k] == 1
}

func (k Key) Name() string {
	return sdl.GetKeyName(sdl.Key(k))
}

func GetKey() Key {
	return keySym
}

func GetChar() uint8 {
	return uint8(keySym)
}

func GetEvent() (e InputEvent, b bool) {
	if event.cleared {
		return event, false
	}

	return event, true
}
