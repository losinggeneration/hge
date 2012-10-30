// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package input

// HGE Input Event type constants
const (
	INPUT_KEYDOWN     = iota
	INPUT_KEYUP       = iota
	INPUT_MBUTTONDOWN = iota
	INPUT_MBUTTONUP   = iota
	INPUT_MOUSEMOVE   = iota
	INPUT_MOUSEWHEEL  = iota
)

// HGE Input Event flags
const (
	INP_SHIFT      = iota
	INP_CTRL       = iota
	INP_ALT        = iota
	INP_CAPSLOCK   = iota
	INP_SCROLLLOCK = iota
	INP_NUMLOCK    = iota
	INP_REPEAT     = iota
)

// HGE_ Virtual-key codes
const (
	K_LBUTTON = iota
	K_RBUTTON = iota
	K_MBUTTON = iota

	K_ESCAPE    = iota
	K_BACKSPACE = iota
	K_TAB       = iota
	K_ENTER     = iota
	K_SPACE     = iota

	K_SHIFT = iota
	K_CTRL  = iota
	K_ALT   = iota

	K_LWIN = iota
	K_RWIN = iota
	K_APPS = iota

	K_PAUSE      = iota
	K_CAPSLOCK   = iota
	K_NUMLOCK    = iota
	K_SCROLLLOCK = iota

	K_PGUP   = iota
	K_PGDN   = iota
	K_HOME   = iota
	K_END    = iota
	K_INSERT = iota
	K_DELETE = iota

	K_LEFT  = iota
	K_UP    = iota
	K_RIGHT = iota
	K_DOWN  = iota

	K_0 = iota
	K_1 = iota
	K_2 = iota
	K_3 = iota
	K_4 = iota
	K_5 = iota
	K_6 = iota
	K_7 = iota
	K_8 = iota
	K_9 = iota

	K_A = iota
	K_B = iota
	K_C = iota
	K_D = iota
	K_E = iota
	K_F = iota
	K_G = iota
	K_H = iota
	K_I = iota
	K_J = iota
	K_K = iota
	K_L = iota
	K_M = iota
	K_N = iota
	K_O = iota
	K_P = iota
	K_Q = iota
	K_R = iota
	K_S = iota
	K_T = iota
	K_U = iota
	K_V = iota
	K_W = iota
	K_X = iota
	K_Y = iota
	K_Z = iota

	K_GRAVE      = iota
	K_MINUS      = iota
	K_EQUALS     = iota
	K_BACKSLASH  = iota
	K_LBRACKET   = iota
	K_RBRACKET   = iota
	K_SEMICOLON  = iota
	K_APOSTROPHE = iota
	K_COMMA      = iota
	K_PERIOD     = iota
	K_SLASH      = iota

	K_NUMPAD0 = iota
	K_NUMPAD1 = iota
	K_NUMPAD2 = iota
	K_NUMPAD3 = iota
	K_NUMPAD4 = iota
	K_NUMPAD5 = iota
	K_NUMPAD6 = iota
	K_NUMPAD7 = iota
	K_NUMPAD8 = iota
	K_NUMPAD9 = iota

	K_MULTIPLY = iota
	K_DIVIDE   = iota
	K_ADD      = iota
	K_SUBTRACT = iota
	K_DECIMAL  = iota

	K_F1  = iota
	K_F2  = iota
	K_F3  = iota
	K_F4  = iota
	K_F5  = iota
	K_F6  = iota
	K_F7  = iota
	K_F8  = iota
	K_F9  = iota
	K_F10 = iota
	K_F11 = iota
	K_F12 = iota
)
