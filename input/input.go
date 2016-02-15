package input

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import (
	"unsafe"

	"github.com/losinggeneration/hge"
)

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

// HGE Input Event type constants
const (
	INPUT_KEYDOWN     = C.HGE_INPUT_KEYDOWN
	INPUT_KEYUP       = C.HGE_INPUT_KEYUP
	INPUT_MBUTTONDOWN = C.HGE_INPUT_MBUTTONDOWN
	INPUT_MBUTTONUP   = C.HGE_INPUT_MBUTTONUP
	INPUT_MOUSEMOVE   = C.HGE_INPUT_MOUSEMOVE
	INPUT_MOUSEWHEEL  = C.HGE_INPUT_MOUSEWHEEL
)

// HGE Input Event flags
const (
	INP_SHIFT      = C.HGE_INP_SHIFT
	INP_CTRL       = C.HGE_INP_CTRL
	INP_ALT        = C.HGE_INP_ALT
	INP_CAPSLOCK   = C.HGE_INP_CAPSLOCK
	INP_SCROLLLOCK = C.HGE_INP_SCROLLLOCK
	INP_NUMLOCK    = C.HGE_INP_NUMLOCK
	INP_REPEAT     = C.HGE_INP_REPEAT
)

// HGE_ Virtual-key codes
const (
	K_LBUTTON = C.HGE_K_LBUTTON
	K_RBUTTON = C.HGE_K_RBUTTON
	K_MBUTTON = C.HGE_K_MBUTTON

	K_ESCAPE    = C.HGE_K_ESCAPE
	K_BACKSPACE = C.HGE_K_BACKSPACE
	K_TAB       = C.HGE_K_TAB
	K_ENTER     = C.HGE_K_ENTER
	K_SPACE     = C.HGE_K_SPACE

	K_SHIFT = C.HGE_K_SHIFT
	K_CTRL  = C.HGE_K_CTRL
	K_ALT   = C.HGE_K_ALT

	K_LWIN = C.HGE_K_LWIN
	K_RWIN = C.HGE_K_RWIN
	K_APPS = C.HGE_K_APPS

	K_PAUSE      = C.HGE_K_PAUSE
	K_CAPSLOCK   = C.HGE_K_CAPSLOCK
	K_NUMLOCK    = C.HGE_K_NUMLOCK
	K_SCROLLLOCK = C.HGE_K_SCROLLLOCK

	K_PGUP   = C.HGE_K_PGUP
	K_PGDN   = C.HGE_K_PGDN
	K_HOME   = C.HGE_K_HOME
	K_END    = C.HGE_K_END
	K_INSERT = C.HGE_K_INSERT
	K_DELETE = C.HGE_K_DELETE

	K_LEFT  = C.HGE_K_LEFT
	K_UP    = C.HGE_K_UP
	K_RIGHT = C.HGE_K_RIGHT
	K_DOWN  = C.HGE_K_DOWN

	K_0 = C.HGE_K_0
	K_1 = C.HGE_K_1
	K_2 = C.HGE_K_2
	K_3 = C.HGE_K_3
	K_4 = C.HGE_K_4
	K_5 = C.HGE_K_5
	K_6 = C.HGE_K_6
	K_7 = C.HGE_K_7
	K_8 = C.HGE_K_8
	K_9 = C.HGE_K_9

	K_A = C.HGE_K_A
	K_B = C.HGE_K_B
	K_C = C.HGE_K_C
	K_D = C.HGE_K_D
	K_E = C.HGE_K_E
	K_F = C.HGE_K_F
	K_G = C.HGE_K_G
	K_H = C.HGE_K_H
	K_I = C.HGE_K_I
	K_J = C.HGE_K_J
	K_K = C.HGE_K_K
	K_L = C.HGE_K_L
	K_M = C.HGE_K_M
	K_N = C.HGE_K_N
	K_O = C.HGE_K_O
	K_P = C.HGE_K_P
	K_Q = C.HGE_K_Q
	K_R = C.HGE_K_R
	K_S = C.HGE_K_S
	K_T = C.HGE_K_T
	K_U = C.HGE_K_U
	K_V = C.HGE_K_V
	K_W = C.HGE_K_W
	K_X = C.HGE_K_X
	K_Y = C.HGE_K_Y
	K_Z = C.HGE_K_Z

	K_GRAVE      = C.HGE_K_GRAVE
	K_MINUS      = C.HGE_K_MINUS
	K_EQUALS     = C.HGE_K_EQUALS
	K_BACKSLASH  = C.HGE_K_BACKSLASH
	K_LBRACKET   = C.HGE_K_LBRACKET
	K_RBRACKET   = C.HGE_K_RBRACKET
	K_SEMICOLON  = C.HGE_K_SEMICOLON
	K_APOSTROPHE = C.HGE_K_APOSTROPHE
	K_COMMA      = C.HGE_K_COMMA
	K_PERIOD     = C.HGE_K_PERIOD
	K_SLASH      = C.HGE_K_SLASH

	K_NUMPAD0 = C.HGE_K_NUMPAD0
	K_NUMPAD1 = C.HGE_K_NUMPAD1
	K_NUMPAD2 = C.HGE_K_NUMPAD2
	K_NUMPAD3 = C.HGE_K_NUMPAD3
	K_NUMPAD4 = C.HGE_K_NUMPAD4
	K_NUMPAD5 = C.HGE_K_NUMPAD5
	K_NUMPAD6 = C.HGE_K_NUMPAD6
	K_NUMPAD7 = C.HGE_K_NUMPAD7
	K_NUMPAD8 = C.HGE_K_NUMPAD8
	K_NUMPAD9 = C.HGE_K_NUMPAD9

	K_MULTIPLY = C.HGE_K_MULTIPLY
	K_DIVIDE   = C.HGE_K_DIVIDE
	K_ADD      = C.HGE_K_ADD
	K_SUBTRACT = C.HGE_K_SUBTRACT
	K_DECIMAL  = C.HGE_K_DECIMAL

	K_F1  = C.HGE_K_F1
	K_F2  = C.HGE_K_F2
	K_F3  = C.HGE_K_F3
	K_F4  = C.HGE_K_F4
	K_F5  = C.HGE_K_F5
	K_F6  = C.HGE_K_F6
	K_F7  = C.HGE_K_F7
	K_F8  = C.HGE_K_F8
	K_F9  = C.HGE_K_F9
	K_F10 = C.HGE_K_F10
	K_F11 = C.HGE_K_F11
	K_F12 = C.HGE_K_F12
)

var inputHGE *hge.HGE

func init() {
	inputHGE = hge.New()
}

type Mouse struct {
	X, Y  float64
	Wheel int
	Over  bool
}

func NewMouse(x, y float64) *Mouse {
	return &Mouse{x, y, 0, false}
}

func (m *Mouse) Pos() (x, y float64) {
	var nx, ny C.float

	C.HGE_Input_GetMousePos(inputHGE.HGE, &nx, &ny)
	m.X, m.Y = float64(nx), float64(ny)

	return float64(nx), float64(ny)
}

func (m Mouse) SetPos(a ...interface{}) {
	x, y := m.X, m.Y

	if len(a) > 0 {
		if nx, ok := a[0].(float64); ok {
			x = nx
		}
		if len(a) > 1 {
			if ny, ok := a[1].(float64); ok {
				y = ny
			}
		}
	}
	C.HGE_Input_SetMousePos(inputHGE.HGE, C.float(x), C.float(y))
}

func (m *Mouse) WheelMovement() int {
	m.Wheel = int(C.HGE_Input_GetMouseWheel(inputHGE.HGE))
	return m.Wheel
}

func (m *Mouse) IsOver() bool {
	m.Over = C.HGE_Input_IsMouseOver(inputHGE.HGE) == 1
	return m.Over
}

func NewKey(i int) Key {
	return Key(i)
}

func (k Key) Down() bool {
	return C.HGE_Input_KeyDown(inputHGE.HGE, C.int(k)) == 1
}

func (k Key) Up() bool {
	return C.HGE_Input_KeyUp(inputHGE.HGE, C.int(k)) == 1
}

func (k Key) State() bool {
	return C.HGE_Input_GetKeyState(inputHGE.HGE, C.int(k)) == 1
}
func (k Key) Name() string {
	return C.GoString(C.HGE_Input_GetKeyName(inputHGE.HGE, C.int(k)))
}

func GetKey() Key {
	return Key(C.HGE_Input_GetKey(inputHGE.HGE))
}

func GetChar() int {
	return int(C.HGE_Input_GetChar(inputHGE.HGE))
}

func GetEvent() (e *InputEvent, b bool) {
	e = new(InputEvent)
	b = C.HGE_Input_GetEvent(inputHGE.HGE, (*C.HGE_InputEvent_t)(unsafe.Pointer(e))) == 1
	return e, b
}
