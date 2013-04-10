package hge

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
#include "callback.h"
#include <stdio.h>
void goHGE_System_Log(HGE_t *h, const char *str) {
	HGE_System_Log(h, str);
}
*/
import "C"

import (
	"fmt"
	"math"
	"runtime"
	"unsafe"
)

const (
	VERSION = C.HGE_VERSION
)

type Dword uint32

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
	WINDOWED      BoolState = C.HGE_C_WINDOWED      // bool run in window? (default: false)
	ZBUFFER       BoolState = C.HGE_C_ZBUFFER       // bool use z-buffer? (default: false)
	TEXTUREFILTER BoolState = C.HGE_C_TEXTUREFILTER // bool texture filtering? (default: true)

	USESOUND BoolState = C.HGE_C_USESOUND // bool use sound? (default: true)

	DONTSUSPEND BoolState = C.HGE_C_DONTSUSPEND // bool focus lost:suspend? (default: false)
	HIDEMOUSE   BoolState = C.HGE_C_HIDEMOUSE   // bool hide system cursor? (default: true)

	SHOWSPLASH BoolState = C.HGE_C_SHOWSPLASH // bool show splash? (default: true)

	BOOLSTATE_FORCE_DWORD BoolState = C.HGE_C_BOOLSTATE_FORCE_DWORD
)

const (
	FRAMEFUNC      FuncState = C.HGE_C_FRAMEFUNC      // func() bool frame function (default: nil) (you MUST set this)
	RENDERFUNC     FuncState = C.HGE_C_RENDERFUNC     // func() bool render function (default: nil)
	FOCUSLOSTFUNC  FuncState = C.HGE_C_FOCUSLOSTFUNC  // func() bool focus lost function (default: nil)
	FOCUSGAINFUNC  FuncState = C.HGE_C_FOCUSGAINFUNC  // func() bool focus gain function (default: nil)
	GFXRESTOREFUNC FuncState = C.HGE_C_GFXRESTOREFUNC // func() bool gfx restore function (default: nil)
	EXITFUNC       FuncState = C.HGE_C_EXITFUNC       // func() bool exit function (default: nil)

	FUNCSTATE_FORCE_DWORD FuncState = C.HGE_C_FUNCSTATE_FORCE_DWORD
)

const (
	HWND       HwndState = C.HGE_C_HWND       // int		window handle: read only
	HWNDPARENT HwndState = C.HGE_C_HWNDPARENT // int		parent win handle	(default: 0)

	HWNDSTATE_FORCE_DWORD HwndState = C.HGE_C_HWNDSTATE_FORCE_DWORD
)

type Hwnd struct {
	hwnd C.HWND
}

const (
	SCREENWIDTH  IntState = C.HGE_C_SCREENWIDTH  // int screen width (default: 800)
	SCREENHEIGHT IntState = C.HGE_C_SCREENHEIGHT // int screen height (default: 600)
	SCREENBPP    IntState = C.HGE_C_SCREENBPP    // int screen bitdepth (default: 32) (desktop bpp in windowed mode)

	SAMPLERATE   IntState = C.HGE_C_SAMPLERATE   // int sample rate (default: 44100)
	FXVOLUME     IntState = C.HGE_C_FXVOLUME     // int global fx volume (default: 100)
	MUSVOLUME    IntState = C.HGE_C_MUSVOLUME    // int global music volume (default: 100)
	STREAMVOLUME IntState = C.HGE_C_STREAMVOLUME // int stream music volume (default: 100)

	FPS IntState = C.HGE_C_FPS // int fixed fps (default: hge.FPS_UNLIMITED)

	POWERSTATUS IntState = C.HGE_C_POWERSTATUS // int battery life percent + status

	ORIGSCREENWIDTH  IntState = C.HGE_C_ORIGSCREENWIDTH  // int original screen width (default: 800 ... not valid until hge.System_Initiate()!)
	ORIGSCREENHEIGHT IntState = C.HGE_C_ORIGSCREENHEIGHT // int original screen height (default: 600 ... not valid until hge.System_Initiate()!))

	INTSTATE_FORCE_DWORD IntState = C.HGE_C_INTSTATE_FORCE_DWORD
)

const (
	ICON  StringState = C.HGE_C_ICON  // string icon resource (default: nil)
	TITLE StringState = C.HGE_C_TITLE // string window title (default: "HGE")

	INIFILE StringState = C.HGE_C_INIFILE // string ini file (default: nil) (meaning no file)
	LOGFILE StringState = C.HGE_C_LOGFILE // string log file (default: nil) (meaning no file)

	STRINGSTATE_FORCE_DWORD StringState = C.HGE_C_STRINGSTATE_FORCE_DWORD
)

type (
	BoolState   int
	FuncState   int
	HwndState   int
	IntState    int
	StringState int
)

// HGE_POWERSTATUS system state special constants
const (
	PWR_AC          = C.HGE_PWR_AC
	PWR_UNSUPPORTED = C.HGE_PWR_UNSUPPORTED
)

var BoolToCInt = map[bool]C.BOOL{
	false: 0,
	true:  1,
}

// HGE struct from C
type HGE struct {
	HGE *C.HGE_t
}

type Error struct {
	*HGE
}

func (e *Error) Error() string {
	return e.HGE.GetErrorMessage()
}

// Creates a new instance of an hge structure
func New(a ...interface{}) *HGE {
	ver := VERSION

	if len(a) == 1 {
		if v, ok := a[0].(int); ok {
			ver = v
		}
	}

	h := new(HGE)
	h.HGE = C.HGE_Create(C.int(ver))
	fmt.Println("Created HGE", *h)
	runtime.SetFinalizer(h, func(hge *HGE) {
		hge.Free()
	})

	return h
}

// Releases the memory the C++ library allocated for the HGE struct
func (h *HGE) Free() {
	fmt.Println("Freeing HGE", *h)
	C.HGE_Release(h.HGE)
}

// Initializes hardware and software needed to run engine.
func (h *HGE) Initiate() error {
	if C.HGE_System_Initiate(h.HGE) == 0 {
		return &Error{h}
	}

	return nil
}

//  Restores video mode and frees allocated resources.
func (h *HGE) Shutdown() {
	C.HGE_System_Shutdown(h.HGE)
}

// Starts running user defined frame func (h *HGE)tion.
func (h *HGE) Start() error {
	if C.HGE_System_Start(h.HGE) == 0 {
		return &Error{h}
	}

	return nil
}

//  Returns last occured HGE error description.
func (h *HGE) GetErrorMessage() string {
	return C.GoString(C.HGE_System_GetErrorMessage(h.HGE))
}

// Writes a formatted message to the log file.
func (h *HGE) Log(format string, v ...interface{}) {
	var str string

	if v != nil {
		str = fmt.Sprintf(format, v...)
	} else {
		str = format
	}

	fstr := C.CString(str)
	defer C.free(unsafe.Pointer(fstr))

	C.goHGE_System_Log(h.HGE, fstr)
}

// Launches an URL or external executable/data file.
func (h *HGE) Launch(url string) bool {
	urlstr := C.CString(url)
	defer C.free(unsafe.Pointer(urlstr))

	return C.HGE_System_Launch(h.HGE, urlstr) == 1
}

//  Saves current screen snapshot into a file.
func (h *HGE) Snapshot(a ...interface{}) {
	if len(a) == 1 {
		if filename, ok := a[0].(string); ok {
			fname := C.CString(filename)
			defer C.free(unsafe.Pointer(fname))

			C.HGE_System_Snapshot(h.HGE, fname)
			return
		}
	}

	C.HGE_System_Snapshot(h.HGE, nil)
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
			case func() bool:
				h.setStateFunc(a[0].(FuncState), func() int {
					if a[1].(func() bool)() {
						return 1
					}

					return 0
				})
				return
			default:
				h.setStateFunc(a[0].(FuncState), nil)
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
	C.HGE_System_SetStateBool(h.HGE, C.HGE_BoolState_t(state), BoolToCInt[value])
}

func (h *HGE) setStateFunc(state FuncState, value StateFunc) {
	funcCBs[state] = value
	switch state {
	case FRAMEFUNC:
		C.setFrameFunc(h.HGE, C.HGE_FuncState_t(state))
	case RENDERFUNC:
		C.setRenderFunc(h.HGE, C.HGE_FuncState_t(state))
	case FOCUSLOSTFUNC:
		C.setFocusLostFunc(h.HGE, C.HGE_FuncState_t(state))
	case FOCUSGAINFUNC:
		C.setFocusGainFunc(h.HGE, C.HGE_FuncState_t(state))
	case GFXRESTOREFUNC:
		C.setGfxRestoreFunc(h.HGE, C.HGE_FuncState_t(state))
	case EXITFUNC:
		C.setExitFunc(h.HGE, C.HGE_FuncState_t(state))
	}
}

func (h *HGE) setStateHwnd(state HwndState, value *Hwnd) {
	C.HGE_System_SetStateHwnd(h.HGE, C.HGE_HwndState_t(state), value.hwnd)
}

func (h *HGE) setStateInt(state IntState, value int) {
	C.HGE_System_SetStateInt(h.HGE, C.HGE_IntState_t(state), C.int(value))
}

func (h *HGE) setStateString(state StringState, value string) {
	val := C.CString(value)
	defer C.free(unsafe.Pointer(val))

	C.HGE_System_SetStateString(h.HGE, C.HGE_StringState_t(state), val)
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
	return C.HGE_System_GetStateBool(h.HGE, C.HGE_BoolState_t(state)) == 1
}

func (h *HGE) getStateFunc(state FuncState) StateFunc {
	// I don't know how to convert the HGE_Callback C func (h *HGE)tion type to a Go
	// func (h *HGE)tion, so we just pass back the Go func (h *HGE)tion
	return funcCBs[state]
}

func (h *HGE) getStateHwnd(state HwndState) Hwnd {
	var hwnd Hwnd
	hwnd.hwnd = C.HGE_System_GetStateHwnd(h.HGE, C.HGE_HwndState_t(state))
	return hwnd
}

func (h *HGE) getStateInt(state IntState) int {
	return int(C.HGE_System_GetStateInt(h.HGE, C.HGE_IntState_t(state)))
}

func (h *HGE) getStateString(state StringState) string {
	return C.GoString(C.HGE_System_GetStateString(h.HGE, C.HGE_StringState_t(state)))
}
