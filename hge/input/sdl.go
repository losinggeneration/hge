// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package input

import "fmt"
import "github.com/banthar/Go-SDL/sdl"

type Type int
type Key int
type Flag int

// HGE Input Event structure
type InputEvent struct {
	Type    Type    // event type
	Key     Key     // key code
	Flags   Flag    // event flags
	Chr     int     // character code
	Wheel   int     // wheel shift
	X       float32 // mouse cursor x-coordinate
	Y       float32 // mouse cursor y-coordinate
	cleared bool    // true if there is no useful data here
}

var (
	keys      [last_key]bool
	keySym    Key
	lastEvent InputEvent
)

// Process events
func Process() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.KeyboardEvent:
			k, _ := e.(*sdl.KeyboardEvent)
			keys[k.Keysym.Sym] = 1 == k.State
			keySym = Key(k.Keysym.Sym)
			lastEvent.Key = Key(keySym)
		}
	}
}

// Clear the event queue
func ClearQueue() {
	keySym = 0
	for i := 0; i < last_key; i++ {
		keys[i] = false
	}

	lastEvent = InputEvent{cleared: true}
}

// Update the internal mouse structure
func UpdateMouse() {
}

// Initialize the event structure
func Initialize() error {
	return fmt.Errorf("Input Initialize not implemented")
}

type Mouse struct {
	X, Y  float64
	Wheel int
	Over  bool
}

func NewMouse(x, y float64) *Mouse {
	return &Mouse{X: x, Y: y}
}

func (m *Mouse) Pos() (x, y float64) {
	return 0, 0
}

func (m Mouse) SetPos(a ...interface{}) {
}

func (m *Mouse) WheelMovement() int {
	return 0
}

func (m *Mouse) IsOver() bool {
	return false
}

func (k Key) Down() bool {
	return keys[k]
}

func (k Key) Up() bool {
	return keys[k]
}

func (k Key) State() bool {
	ks := sdl.GetKeyState()
	return ks[k] == 1
}
func (k Key) Name() string {
	return ""
}

func GetKey() Key {
	return keySym
}

func GetChar() uint8 {
	return uint8(keySym)
}

func GetEvent() (e *InputEvent, b bool) {
	if lastEvent.cleared {
		return nil, false
	}

	return &lastEvent, true
}
