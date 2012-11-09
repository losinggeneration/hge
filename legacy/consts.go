package legacy

import (
	"github.com/losinggeneration/hge-go/hge"
	// 	"github.com/losinggeneration/hge-go/hge/gfx"
	// 	"github.com/losinggeneration/hge-go/hge/ini"
	"github.com/losinggeneration/hge-go/hge/input"

// 	"github.com/losinggeneration/hge-go/hge/rand"
// 	"github.com/losinggeneration/hge-go/hge/resource"
// 	"github.com/losinggeneration/hge-go/hge/sound"
// 	"github.com/losinggeneration/hge-go/hge/timer"
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

// FPS system state special constants
const (
	FPS_UNLIMITED = hge.FPS_UNLIMITED
	FPS_VSYNC     = hge.FPS_VSYNC
)

// HGE System state constants
const (
	WINDOWED      BoolState = BoolState(hge.WINDOWED)
	ZBUFFER       BoolState = BoolState(hge.ZBUFFER)
	TEXTUREFILTER BoolState = BoolState(hge.TEXTUREFILTER)

	USESOUND BoolState = BoolState(hge.USESOUND)

	DONTSUSPEND BoolState = BoolState(hge.DONTSUSPEND)
	HIDEMOUSE   BoolState = BoolState(hge.HIDEMOUSE)

	SHOWSPLASH BoolState = BoolState(hge.SHOWSPLASH)
)

// When any of these return true, it indicates to stop the main loop from
// continuing to run.
const (
	FRAMEFUNC      FuncState = FuncState(hge.FRAMEFUNC)
	RENDERFUNC     FuncState = FuncState(hge.RENDERFUNC)
	FOCUSLOSTFUNC  FuncState = FuncState(hge.FOCUSLOSTFUNC)
	FOCUSGAINFUNC  FuncState = FuncState(hge.FOCUSGAINFUNC)
	GFXRESTOREFUNC FuncState = FuncState(hge.GFXRESTOREFUNC)
	EXITFUNC       FuncState = FuncState(hge.EXITFUNC)
)

const (
	HWND       HwndState = HwndState(hge.HWND)
	HWNDPARENT HwndState = HwndState(hge.HWNDPARENT)
)

const (
	SCREENWIDTH      IntState = IntState(hge.SCREENWIDTH)
	SCREENHEIGHT     IntState = IntState(hge.SCREENHEIGHT)
	SCREENBPP        IntState = IntState(hge.SCREENBPP)
	ORIGSCREENWIDTH  IntState = IntState(hge.ORIGSCREENWIDTH)
	ORIGSCREENHEIGHT IntState = IntState(hge.ORIGSCREENHEIGHT)
	FPS              IntState = IntState(hge.FPS)

	SAMPLERATE   IntState = IntState(hge.SAMPLERATE)
	FXVOLUME     IntState = IntState(hge.FXVOLUME)
	MUSVOLUME    IntState = IntState(hge.MUSVOLUME)
	STREAMVOLUME IntState = IntState(hge.STREAMVOLUME)

	POWERSTATUS IntState = IntState(hge.POWERSTATUS)
)

const (
	ICON  StringState = StringState(hge.ICON)
	TITLE StringState = StringState(hge.TITLE)

	INIFILE StringState = StringState(hge.INIFILE)
	LOGFILE StringState = StringState(hge.LOGFILE)
)

// HGE Input Event type constants
const (
	INPUT_KEYDOWN     Type = Type(input.INPUT_KEYDOWN)
	INPUT_KEYUP       Type = Type(input.INPUT_KEYUP)
	INPUT_MBUTTONDOWN Type = Type(input.INPUT_MBUTTONDOWN)
	INPUT_MBUTTONUP   Type = Type(input.INPUT_MBUTTONUP)
	INPUT_MOUSEMOVE   Type = Type(input.INPUT_MOUSEMOVE)
	INPUT_MOUSEWHEEL  Type = Type(input.INPUT_MOUSEWHEEL)
)

// HGE Input Event flags
const (
	INP_SHIFT      Flag = Flag(input.INP_SHIFT)
	INP_CTRL       Flag = Flag(input.INP_CTRL)
	INP_ALT        Flag = Flag(input.INP_ALT)
	INP_CAPSLOCK   Flag = Flag(input.INP_CAPSLOCK)
	INP_SCROLLLOCK Flag = Flag(input.INP_SCROLLLOCK)
	INP_NUMLOCK    Flag = Flag(input.INP_NUMLOCK)
	INP_REPEAT     Flag = Flag(input.INP_REPEAT)
)

// HGE Input Mouse button
const (
	M_LBUTTON Button = Button(input.M_LBUTTON)
	M_RBUTTON Button = Button(input.M_RBUTTON)
	M_MBUTTON Button = Button(input.M_MBUTTON)
)

// HGE Virtual-key codes
const (
	K_ESCAPE    Key = Key(input.K_ESCAPE)
	K_BACKSPACE Key = Key(input.K_BACKSPACE)
	K_TAB       Key = Key(input.K_TAB)
	K_ENTER     Key = Key(input.K_ENTER)
	K_SPACE     Key = Key(input.K_SPACE)

	K_SHIFT Key = Key(input.K_SHIFT)
	K_CTRL  Key = Key(input.K_CTRL)
	K_ALT   Key = Key(input.K_ALT)

	K_LWIN Key = Key(input.K_LWIN)
	K_RWIN Key = Key(input.K_RWIN)
	K_APPS Key = Key(input.K_APPS)

	K_PAUSE      Key = Key(input.K_PAUSE)
	K_CAPSLOCK   Key = Key(input.K_CAPSLOCK)
	K_NUMLOCK    Key = Key(input.K_NUMLOCK)
	K_SCROLLLOCK Key = Key(input.K_SCROLLLOCK)

	K_PGUP   Key = Key(input.K_PGUP)
	K_PGDN   Key = Key(input.K_PGDN)
	K_HOME   Key = Key(input.K_HOME)
	K_END    Key = Key(input.K_END)
	K_INSERT Key = Key(input.K_INSERT)
	K_DELETE Key = Key(input.K_DELETE)

	K_LEFT  Key = Key(input.K_LEFT)
	K_UP    Key = Key(input.K_UP)
	K_RIGHT Key = Key(input.K_RIGHT)
	K_DOWN  Key = Key(input.K_DOWN)

	K_0 Key = Key(input.K_0)
	K_1 Key = Key(input.K_1)
	K_2 Key = Key(input.K_2)
	K_3 Key = Key(input.K_3)
	K_4 Key = Key(input.K_4)
	K_5 Key = Key(input.K_5)
	K_6 Key = Key(input.K_6)
	K_7 Key = Key(input.K_7)
	K_8 Key = Key(input.K_8)
	K_9 Key = Key(input.K_9)

	K_A Key = Key(input.K_A)
	K_B Key = Key(input.K_B)
	K_C Key = Key(input.K_C)
	K_D Key = Key(input.K_D)
	K_E Key = Key(input.K_E)
	K_F Key = Key(input.K_F)
	K_G Key = Key(input.K_G)
	K_H Key = Key(input.K_H)
	K_I Key = Key(input.K_I)
	K_J Key = Key(input.K_J)
	K_K Key = Key(input.K_K)
	K_L Key = Key(input.K_L)
	K_M Key = Key(input.K_M)
	K_N Key = Key(input.K_N)
	K_O Key = Key(input.K_O)
	K_P Key = Key(input.K_P)
	K_Q Key = Key(input.K_Q)
	K_R Key = Key(input.K_R)
	K_S Key = Key(input.K_S)
	K_T Key = Key(input.K_T)
	K_U Key = Key(input.K_U)
	K_V Key = Key(input.K_V)
	K_W Key = Key(input.K_W)
	K_X Key = Key(input.K_X)
	K_Y Key = Key(input.K_Y)
	K_Z Key = Key(input.K_Z)

	K_GRAVE      Key = Key(input.K_GRAVE)
	K_MINUS      Key = Key(input.K_MINUS)
	K_EQUALS     Key = Key(input.K_EQUALS)
	K_BACKSLASH  Key = Key(input.K_BACKSLASH)
	K_LBRACKET   Key = Key(input.K_LBRACKET)
	K_RBRACKET   Key = Key(input.K_RBRACKET)
	K_SEMICOLON  Key = Key(input.K_SEMICOLON)
	K_APOSTROPHE Key = Key(input.K_APOSTROPHE)
	K_COMMA      Key = Key(input.K_COMMA)
	K_PERIOD     Key = Key(input.K_PERIOD)
	K_SLASH      Key = Key(input.K_SLASH)

	K_NUMPAD0 Key = Key(input.K_NUMPAD0)
	K_NUMPAD1 Key = Key(input.K_NUMPAD1)
	K_NUMPAD2 Key = Key(input.K_NUMPAD2)
	K_NUMPAD3 Key = Key(input.K_NUMPAD3)
	K_NUMPAD4 Key = Key(input.K_NUMPAD4)
	K_NUMPAD5 Key = Key(input.K_NUMPAD5)
	K_NUMPAD6 Key = Key(input.K_NUMPAD6)
	K_NUMPAD7 Key = Key(input.K_NUMPAD7)
	K_NUMPAD8 Key = Key(input.K_NUMPAD8)
	K_NUMPAD9 Key = Key(input.K_NUMPAD9)

	K_MULTIPLY Key = Key(input.K_MULTIPLY)
	K_DIVIDE   Key = Key(input.K_DIVIDE)
	K_ADD      Key = Key(input.K_ADD)
	K_SUBTRACT Key = Key(input.K_SUBTRACT)
	K_DECIMAL  Key = Key(input.K_DECIMAL)

	K_F1  Key = Key(input.K_F1)
	K_F2  Key = Key(input.K_F2)
	K_F3  Key = Key(input.K_F3)
	K_F4  Key = Key(input.K_F4)
	K_F5  Key = Key(input.K_F5)
	K_F6  Key = Key(input.K_F6)
	K_F7  Key = Key(input.K_F7)
	K_F8  Key = Key(input.K_F8)
	K_F9  Key = Key(input.K_F9)
	K_F10 Key = Key(input.K_F10)
	K_F11 Key = Key(input.K_F11)
	K_F12 Key = Key(input.K_F12)
)
