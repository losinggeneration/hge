package hge

import (
	"math"
	"runtime"
)

const (
	VERSION = 0x181
)

// Common math constants
const (
	Pi     = math.Pi
	Pi_2   = math.Pi / 2
	Pi_4   = math.Pi / 4
	One_Pi = 1 / math.Pi
	Two_Pi = 2 / math.Pi
)

// HGE System state constants
const (
	WINDOWED      BoolState = iota // bool run in window? (default: false)
	ZBUFFER       BoolState = iota // bool use z-buffer? (default: false)
	TEXTUREFILTER BoolState = iota // bool texture filtering? (default: true)

	USESOUND BoolState = iota // bool use sound? (default: true)

	DONTSUSPEND BoolState = iota // bool focus lost:suspend? (default: false)
	HIDEMOUSE   BoolState = iota // bool hide system cursor? (default: true)

	SHOWSPLASH BoolState = iota // bool show splash? (default: true)

// 	BOOLSTATE_FORCE_DWORD BoolState = C.HGE_C_BOOLSTATE_FORCE_DWORD
)

const (
	FRAMEFUNC      FuncState = iota // func() bool frame function (default: nil) (you MUST set this)
	RENDERFUNC     FuncState = iota // func() bool render function (default: nil)
	FOCUSLOSTFUNC  FuncState = iota // func() bool focus lost function (default: nil)
	FOCUSGAINFUNC  FuncState = iota // func() bool focus gain function (default: nil)
	GFXRESTOREFUNC FuncState = iota // func() bool gfx restore function (default: nil)
	EXITFUNC       FuncState = iota // func() bool exit function (default: nil)

// 	FUNCSTATE_FORCE_DWORD FuncState = C.HGE_C_FUNCSTATE_FORCE_DWORD
)

const (
	HWND       HwndState = iota // int		window handle: read only
	HWNDPARENT HwndState = iota // int		parent win handle	(default: 0)

// 	HWNDSTATE_FORCE_DWORD HwndState = C.HGE_C_HWNDSTATE_FORCE_DWORD
)

type Hwnd struct {
}

const (
	SCREENWIDTH  IntState = iota // int screen width (default: 800)
	SCREENHEIGHT IntState = iota // int screen height (default: 600)
	SCREENBPP    IntState = iota // int screen bitdepth (default: 32) (desktop bpp in windowed mode)

	SAMPLERATE   IntState = iota // int sample rate (default: 44100)
	FXVOLUME     IntState = iota // int global fx volume (default: 100)
	MUSVOLUME    IntState = iota // int global music volume (default: 100)
	STREAMVOLUME IntState = iota // int stream music volume (default: 100)

	FPS IntState = iota // int fixed fps (default: hge.FPS_UNLIMITED)

	POWERSTATUS IntState = iota // int battery life percent + status

	ORIGSCREENWIDTH  IntState = iota // int original screen width (default: 800 ... not valid until hge.System_Initiate()!)
	ORIGSCREENHEIGHT IntState = iota // int original screen height (default: 600 ... not valid until hge.System_Initiate()!))

// 	INTSTATE_FORCE_DWORD IntState = C.HGE_C_INTSTATE_FORCE_DWORD
)

const (
	ICON  StringState = iota // string icon resource (default: nil)
	TITLE StringState = iota // string window title (default: "HGE")

	INIFILE StringState = iota // string ini file (default: nil) (meaning no file)
	LOGFILE StringState = iota // string log file (default: nil) (meaning no file)

// 	STRINGSTATE_FORCE_DWORD StringState = C.HGE_C_STRINGSTATE_FORCE_DWORD
)

type (
	BoolState   int
	FuncState   int
	HwndState   int
	IntState    int
	StringState int
)

type StateFunc func() int

// HGE_POWERSTATUS system state special constants
const (
	PWR_AC          = iota
	PWR_UNSUPPORTED = iota
)

// HGE struct from C
type HGE struct {
}

type Error struct {
	*HGE
}

func (e *Error) Error() string {
	return e.GetErrorMessage()
}

// Creates a new instance of an hge structure
func New(a ...interface{}) *HGE {
	if len(a) == 1 {
		if v, ok := a[0].(int); ok {
			if VERSION != v {
				return nil
			}
		}
	}

	h := new(HGE)
	runtime.SetFinalizer(h, func(hge *HGE) {
		hge.Free()
	})

	return h
}

// Releases the memory the C++ library allocated for the HGE struct
func (h *HGE) Free() {
}

// Initializes hardware and software needed to run engine.
func (h *HGE) Initiate() error {
	return nil
}

//  Restores video mode and frees allocated resources.
func (h *HGE) Shutdown() {
}

// Starts running user defined frame func (h *HGE)tion.
func (h *HGE) Start() error {
	return nil
}

//  Returns last occured HGE error description.
func (h *HGE) GetErrorMessage() string {
	return ""
}

// Writes a formatted message to the log file.
func (h *HGE) Log(format string, v ...interface{}) {
}

// Launches an URL or external executable/data file.
func (h *HGE) Launch(url string) bool {
	return true
}

//  Saves current screen snapshot into a file.
func (h *HGE) Snapshot(a ...interface{}) {
}

// Sets internal system states.
// First param should be one of: BoolState, IntState, StringState, FuncState, HwndState
// Second parameter must be of the matching type, bool, int, string, StateFunc/func (h *HGE)() int, *Hwnd
func (h *HGE) SetState(a ...interface{}) {
	if len(a) == 2 {
		switch a[0].(type) {
		case BoolState:
			if bs, ok := a[1].(bool); ok {
				h.setStateBool(a[0].(BoolState), bs)
				return
			}

		case IntState:
			if is, ok := a[1].(int); ok {
				h.setStateInt(a[0].(IntState), is)
				return
			}

		case StringState:
			if ss, ok := a[1].(string); ok {
				h.setStateString(a[0].(StringState), ss)
				return
			}

		case FuncState:
			switch a[1].(type) {
			case StateFunc:
				h.setStateFunc(a[0].(FuncState), a[1].(StateFunc))
				return
			case func() int:
				h.setStateFunc(a[0].(FuncState), a[1].(func() int))
				return
			}

		case HwndState:
			if hs, ok := a[1].(*Hwnd); ok {
				h.setStateHwnd(a[0].(HwndState), hs)
				return
			}
		}
	}
}

func (h *HGE) setStateBool(state BoolState, value bool) {
}

func (h *HGE) setStateFunc(state FuncState, value StateFunc) {
}

func (h *HGE) setStateHwnd(state HwndState, value *Hwnd) {
}

func (h *HGE) setStateInt(state IntState, value int) {
}

func (h *HGE) setStateString(state StringState, value string) {
}

// Returns internal system state values.
func (h *HGE) GetState(a ...interface{}) interface{} {
	if len(a) == 1 {
		switch a[0].(type) {
		case BoolState:
			return h.getStateBool(a[0].(BoolState))

		case IntState:
			return h.getStateInt(a[0].(IntState))

		case StringState:
			return h.getStateString(a[0].(StringState))

		case FuncState:
			return h.getStateFunc(a[0].(FuncState))

		case HwndState:
			return h.getStateHwnd(a[0].(HwndState))
		}
	}

	return nil
}

func (h *HGE) getStateBool(state BoolState) bool {
	return true
}

func (h *HGE) getStateFunc(state FuncState) StateFunc {
	return nil
}

func (h *HGE) getStateHwnd(state HwndState) Hwnd {
	return Hwnd{}
}

func (h *HGE) getStateInt(state IntState) int {
	return 0
}

func (h *HGE) getStateString(state StringState) string {
	return ""
}
