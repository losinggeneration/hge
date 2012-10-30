// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package input

import "fmt"

type Key int

// HGE Input Event structure
type InputEvent struct {
	Type  int     // event type
	Key   int     // key code
	Flags int     // event flags
	Chr   int     // character code
	Wheel int     // wheel shift
	X     float32 // mouse cursor x-coordinate
	Y     float32 // mouse cursor y-coordinate
}

func Process() {
}

func ClearQueue() {
}

func UpdateMouse() {
}

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

func NewKey(i int) Key {
	return 0
}

func (k Key) Down() bool {
	return false
}

func (k Key) Up() bool {
	return false
}

func (k Key) State() bool {
	return false
}
func (k Key) Name() string {
	return ""
}

func GetKey() Key {
	return 0
}

func GetChar() int {
	return 0
}

func GetEvent() (e *InputEvent, b bool) {
	return nil, false
}
