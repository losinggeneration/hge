// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package input

import "github.com/banthar/Go-SDL/sdl"

// HGE Input Event type constants
const (
	INPUT_KEYDOWN     Type = iota
	INPUT_KEYUP       Type = iota
	INPUT_MBUTTONDOWN Type = iota
	INPUT_MBUTTONUP   Type = iota
	INPUT_MOUSEMOVE   Type = iota
	INPUT_MOUSEWHEEL  Type = iota
)

// HGE Input Event flags
const (
	INP_SHIFT      Flag = iota
	INP_CTRL       Flag = iota
	INP_ALT        Flag = iota
	INP_CAPSLOCK   Flag = iota
	INP_SCROLLLOCK Flag = iota
	INP_NUMLOCK    Flag = iota
	INP_REPEAT     Flag = iota
)

// This is the offset so it doesn't clash with any SDL key defines
const key_offset = 0x160

// HGE_ Virtual-key codes
const (
	K_LBUTTON Key = iota + key_offset
	K_RBUTTON Key = iota + key_offset
	K_MBUTTON Key = iota + key_offset

	K_ESCAPE    Key = sdl.K_ESCAPE
	K_BACKSPACE Key = sdl.K_BACKSPACE
	K_TAB       Key = sdl.K_TAB
	K_ENTER     Key = sdl.K_RETURN
	K_SPACE     Key = sdl.K_SPACE

	K_SHIFT Key = sdl.K_RSHIFT | sdl.K_LSHIFT
	K_CTRL  Key = sdl.K_RCTRL | sdl.K_LCTRL
	K_ALT   Key = sdl.K_RALT | sdl.K_LALT

	K_LWIN Key = sdl.K_LSUPER
	K_RWIN Key = sdl.K_RSUPER
	K_APPS Key = iota + key_offset

	K_PAUSE      Key = sdl.K_PAUSE
	K_CAPSLOCK   Key = sdl.K_CAPSLOCK
	K_NUMLOCK    Key = sdl.K_NUMLOCK
	K_SCROLLLOCK Key = sdl.K_SCROLLOCK

	K_PGUP   Key = sdl.K_PAGEUP
	K_PGDN   Key = sdl.K_PAGEDOWN
	K_HOME   Key = sdl.K_HOME
	K_END    Key = sdl.K_END
	K_INSERT Key = sdl.K_INSERT
	K_DELETE Key = sdl.K_DELETE

	K_LEFT  Key = sdl.K_LEFT
	K_UP    Key = sdl.K_UP
	K_RIGHT Key = sdl.K_RIGHT
	K_DOWN  Key = sdl.K_DOWN

	K_0 Key = sdl.K_0
	K_1 Key = sdl.K_1
	K_2 Key = sdl.K_2
	K_3 Key = sdl.K_3
	K_4 Key = sdl.K_4
	K_5 Key = sdl.K_5
	K_6 Key = sdl.K_6
	K_7 Key = sdl.K_7
	K_8 Key = sdl.K_8
	K_9 Key = sdl.K_9

	K_A Key = sdl.K_a
	K_B Key = sdl.K_b
	K_C Key = sdl.K_c
	K_D Key = sdl.K_d
	K_E Key = sdl.K_e
	K_F Key = sdl.K_f
	K_G Key = sdl.K_g
	K_H Key = sdl.K_h
	K_I Key = sdl.K_i
	K_J Key = sdl.K_j
	K_K Key = sdl.K_k
	K_L Key = sdl.K_l
	K_M Key = sdl.K_m
	K_N Key = sdl.K_n
	K_O Key = sdl.K_o
	K_P Key = sdl.K_p
	K_Q Key = sdl.K_q
	K_R Key = sdl.K_r
	K_S Key = sdl.K_s
	K_T Key = sdl.K_t
	K_U Key = sdl.K_u
	K_V Key = sdl.K_v
	K_W Key = sdl.K_w
	K_X Key = sdl.K_x
	K_Y Key = sdl.K_y
	K_Z Key = sdl.K_z

	K_GRAVE      Key = sdl.K_BACKQUOTE
	K_MINUS      Key = sdl.K_MINUS
	K_EQUALS     Key = sdl.K_EQUALS
	K_BACKSLASH  Key = sdl.K_BACKSLASH
	K_LBRACKET   Key = sdl.K_LEFTBRACKET
	K_RBRACKET   Key = sdl.K_RIGHTBRACKET
	K_SEMICOLON  Key = sdl.K_SEMICOLON
	K_APOSTROPHE Key = sdl.K_QUOTE
	K_COMMA      Key = sdl.K_COMMA
	K_PERIOD     Key = sdl.K_PERIOD
	K_SLASH      Key = sdl.K_SLASH

	K_NUMPAD0 Key = sdl.K_KP0
	K_NUMPAD1 Key = sdl.K_KP1
	K_NUMPAD2 Key = sdl.K_KP2
	K_NUMPAD3 Key = sdl.K_KP3
	K_NUMPAD4 Key = sdl.K_KP4
	K_NUMPAD5 Key = sdl.K_KP5
	K_NUMPAD6 Key = sdl.K_KP6
	K_NUMPAD7 Key = sdl.K_KP7
	K_NUMPAD8 Key = sdl.K_KP8
	K_NUMPAD9 Key = sdl.K_KP9

	K_MULTIPLY Key = sdl.K_KP_MULTIPLY
	K_DIVIDE   Key = sdl.K_KP_DIVIDE
	K_ADD      Key = sdl.K_KP_PLUS
	K_SUBTRACT Key = sdl.K_KP_MINUS
	K_DECIMAL  Key = sdl.K_KP_PERIOD

	K_F1  Key = sdl.K_F1
	K_F2  Key = sdl.K_F2
	K_F3  Key = sdl.K_F3
	K_F4  Key = sdl.K_F4
	K_F5  Key = sdl.K_F5
	K_F6  Key = sdl.K_F6
	K_F7  Key = sdl.K_F7
	K_F8  Key = sdl.K_F8
	K_F9  Key = sdl.K_F9
	K_F10 Key = sdl.K_F10
	K_F11 Key = sdl.K_F11
	K_F12 Key = sdl.K_F12
)
