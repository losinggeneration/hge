package legacy

import (
	"github.com/losinggeneration/hge-go/hge"
	"github.com/losinggeneration/hge-go/hge/gfx"
	"github.com/losinggeneration/hge-go/hge/input"
)

// The current version of this package
const (
	VERSION = hge.VERSION
)

// Common math constants
const (
	Pi     = hge.Pi
	Pi_2   = hge.Pi_2
	Pi_4   = hge.Pi_4
	One_Pi = hge.One_Pi
	Two_Pi = hge.Two_Pi
)

// POWERSTATUS system state special constants
const (
	PWR_AC          = hge.PWR_AC
	PWR_UNSUPPORTED = hge.PWR_UNSUPPORTED
)

// HGE System state constants
const (
	WINDOWED      hge.BoolState = hge.BoolState(hge.WINDOWED)
	ZBUFFER       hge.BoolState = hge.BoolState(hge.ZBUFFER)
	TEXTUREFILTER hge.BoolState = hge.BoolState(hge.TEXTUREFILTER)

	USESOUND hge.BoolState = hge.BoolState(hge.USESOUND)

	DONTSUSPEND hge.BoolState = hge.BoolState(hge.DONTSUSPEND)
	HIDEMOUSE   hge.BoolState = hge.BoolState(hge.HIDEMOUSE)

	SHOWSPLASH hge.BoolState = hge.BoolState(hge.SHOWSPLASH)
)

// When any of these return true, it indicates to stop the main loop from
// continuing to run.
const (
	FRAMEFUNC      hge.FuncState = hge.FuncState(hge.FRAMEFUNC)
	RENDERFUNC     hge.FuncState = hge.FuncState(hge.RENDERFUNC)
	FOCUSLOSTFUNC  hge.FuncState = hge.FuncState(hge.FOCUSLOSTFUNC)
	FOCUSGAINFUNC  hge.FuncState = hge.FuncState(hge.FOCUSGAINFUNC)
	GFXRESTOREFUNC hge.FuncState = hge.FuncState(hge.GFXRESTOREFUNC)
	EXITFUNC       hge.FuncState = hge.FuncState(hge.EXITFUNC)
)

const (
	HWND       hge.HwndState = hge.HwndState(hge.HWND)
	HWNDPARENT hge.HwndState = hge.HwndState(hge.HWNDPARENT)
)

const (
	SCREENWIDTH      hge.IntState = hge.IntState(hge.SCREENWIDTH)
	SCREENHEIGHT     hge.IntState = hge.IntState(hge.SCREENHEIGHT)
	SCREENBPP        hge.IntState = hge.IntState(hge.SCREENBPP)
	ORIGSCREENWIDTH  hge.IntState = hge.IntState(hge.ORIGSCREENWIDTH)
	ORIGSCREENHEIGHT hge.IntState = hge.IntState(hge.ORIGSCREENHEIGHT)
	FPS              hge.IntState = hge.IntState(hge.FPS)

	SAMPLERATE   hge.IntState = hge.IntState(hge.SAMPLERATE)
	FXVOLUME     hge.IntState = hge.IntState(hge.FXVOLUME)
	MUSVOLUME    hge.IntState = hge.IntState(hge.MUSVOLUME)
	STREAMVOLUME hge.IntState = hge.IntState(hge.STREAMVOLUME)

	POWERSTATUS hge.IntState = hge.IntState(hge.POWERSTATUS)
)

const (
	ICON  hge.StringState = hge.StringState(hge.ICON)
	TITLE hge.StringState = hge.StringState(hge.TITLE)

	INIFILE hge.StringState = hge.StringState(hge.INIFILE)
	LOGFILE hge.StringState = hge.StringState(hge.LOGFILE)
)

// HGE Blending constants
const (
	BLEND_COLORADD   = gfx.BLEND_COLORADD
	BLEND_COLORMUL   = gfx.BLEND_COLORMUL
	BLEND_ALPHABLEND = gfx.BLEND_ALPHABLEND
	BLEND_ALPHAADD   = gfx.BLEND_ALPHAADD
	BLEND_ZWRITE     = gfx.BLEND_ZWRITE
	BLEND_NOZWRITE   = gfx.BLEND_NOZWRITE

	BLEND_DEFAULT   = gfx.BLEND_DEFAULT
	BLEND_DEFAULT_Z = gfx.BLEND_DEFAULT_Z
)

// HGE_FPS system state special constants
const (
	FPS_UNLIMITED = gfx.FPS_UNLIMITED
	FPS_VSYNC     = gfx.FPS_VSYNC
)

// HGE Primitive type constants
const (
	PRIM_LINES   = gfx.PRIM_LINES
	PRIM_TRIPLES = gfx.PRIM_TRIPLES
	PRIM_QUADS   = gfx.PRIM_QUADS
)

// HGE Input Event type constants
const (
	INPUT_KEYDOWN     = (input.INPUT_KEYDOWN)
	INPUT_KEYUP       = (input.INPUT_KEYUP)
	INPUT_MBUTTONDOWN = (input.INPUT_MBUTTONDOWN)
	INPUT_MBUTTONUP   = (input.INPUT_MBUTTONUP)
	INPUT_MOUSEMOVE   = (input.INPUT_MOUSEMOVE)
	INPUT_MOUSEWHEEL  = (input.INPUT_MOUSEWHEEL)
)

// HGE Input Event flags
const (
	INP_SHIFT      = (input.INP_SHIFT)
	INP_CTRL       = (input.INP_CTRL)
	INP_ALT        = (input.INP_ALT)
	INP_CAPSLOCK   = (input.INP_CAPSLOCK)
	INP_SCROLLLOCK = (input.INP_SCROLLLOCK)
	INP_NUMLOCK    = (input.INP_NUMLOCK)
	INP_REPEAT     = (input.INP_REPEAT)
)

// HGE Virtual-key codes
const (
	// HGE Input Mouse button
	K_LBUTTON = (input.K_LBUTTON)
	K_RBUTTON = (input.K_RBUTTON)
	K_MBUTTON = (input.K_MBUTTON)

	// HGE Input key buttons
	K_ESCAPE    = input.K_ESCAPE
	K_BACKSPACE = input.K_BACKSPACE
	K_TAB       = input.K_TAB
	K_ENTER     = input.K_ENTER
	K_SPACE     = input.K_SPACE

	K_SHIFT = input.K_SHIFT
	K_CTRL  = input.K_CTRL
	K_ALT   = input.K_ALT

	K_LWIN = input.K_LWIN
	K_RWIN = input.K_RWIN
	K_APPS = input.K_APPS

	K_PAUSE      = input.K_PAUSE
	K_CAPSLOCK   = input.K_CAPSLOCK
	K_NUMLOCK    = input.K_NUMLOCK
	K_SCROLLLOCK = input.K_SCROLLLOCK

	K_PGUP   = input.K_PGUP
	K_PGDN   = input.K_PGDN
	K_HOME   = input.K_HOME
	K_END    = input.K_END
	K_INSERT = input.K_INSERT
	K_DELETE = input.K_DELETE

	K_LEFT  = input.K_LEFT
	K_UP    = input.K_UP
	K_RIGHT = input.K_RIGHT
	K_DOWN  = input.K_DOWN

	K_0 = input.K_0
	K_1 = input.K_1
	K_2 = input.K_2
	K_3 = input.K_3
	K_4 = input.K_4
	K_5 = input.K_5
	K_6 = input.K_6
	K_7 = input.K_7
	K_8 = input.K_8
	K_9 = input.K_9

	K_A = input.K_A
	K_B = input.K_B
	K_C = input.K_C
	K_D = input.K_D
	K_E = input.K_E
	K_F = input.K_F
	K_G = input.K_G
	K_H = input.K_H
	K_I = input.K_I
	K_J = input.K_J
	K_K = input.K_K
	K_L = input.K_L
	K_M = input.K_M
	K_N = input.K_N
	K_O = input.K_O
	K_P = input.K_P
	K_Q = input.K_Q
	K_R = input.K_R
	K_S = input.K_S
	K_T = input.K_T
	K_U = input.K_U
	K_V = input.K_V
	K_W = input.K_W
	K_X = input.K_X
	K_Y = input.K_Y
	K_Z = input.K_Z

	K_GRAVE      = input.K_GRAVE
	K_MINUS      = input.K_MINUS
	K_EQUALS     = input.K_EQUALS
	K_BACKSLASH  = input.K_BACKSLASH
	K_LBRACKET   = input.K_LBRACKET
	K_RBRACKET   = input.K_RBRACKET
	K_SEMICOLON  = input.K_SEMICOLON
	K_APOSTROPHE = input.K_APOSTROPHE
	K_COMMA      = input.K_COMMA
	K_PERIOD     = input.K_PERIOD
	K_SLASH      = input.K_SLASH

	K_NUMPAD0 = input.K_NUMPAD0
	K_NUMPAD1 = input.K_NUMPAD1
	K_NUMPAD2 = input.K_NUMPAD2
	K_NUMPAD3 = input.K_NUMPAD3
	K_NUMPAD4 = input.K_NUMPAD4
	K_NUMPAD5 = input.K_NUMPAD5
	K_NUMPAD6 = input.K_NUMPAD6
	K_NUMPAD7 = input.K_NUMPAD7
	K_NUMPAD8 = input.K_NUMPAD8
	K_NUMPAD9 = input.K_NUMPAD9

	K_MULTIPLY = input.K_MULTIPLY
	K_DIVIDE   = input.K_DIVIDE
	K_ADD      = input.K_ADD
	K_SUBTRACT = input.K_SUBTRACT
	K_DECIMAL  = input.K_DECIMAL

	K_F1  = input.K_F1
	K_F2  = input.K_F2
	K_F3  = input.K_F3
	K_F4  = input.K_F4
	K_F5  = input.K_F5
	K_F6  = input.K_F6
	K_F7  = input.K_F7
	K_F8  = input.K_F8
	K_F9  = input.K_F9
	K_F10 = input.K_F10
	K_F11 = input.K_F11
	K_F12 = input.K_F12
)
