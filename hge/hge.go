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

// Hardware color macros
// #define ARGB(a,r,g,b)	((Dword(a)<<24) + (Dword(r)<<16) + (Dword(g)<<8) + Dword(b))
// #define GETA(col)		((col)>>24)
// #define GETR(col)		(((col)>>16) & 0xFF)
// #define GETG(col)		(((col)>>8) & 0xFF)
// #define GETB(col)		((col) & 0xFF)
// #define SETA(col,a)		(((col) & 0x00FFFFFF) + (Dword(a)<<24))
// #define SETR(col,r)		(((col) & 0xFF00FFFF) + (Dword(r)<<16))
// #define SETG(col,g)		(((col) & 0xFFFF00FF) + (Dword(g)<<8))
// #define SETB(col,b)		(((col) & 0xFFFFFF00) + Dword(b))

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

func BoolToCInt(b bool) C.BOOL {
	if b {
		return 1
	}

	return 0
}

// HGE struct from C
type hge *C.HGE_t

type Error struct{}

func (e *Error) Error() string {
	return GetErrorMessage()
}

var HGE hge

func init() {
	HGE = newHGE(VERSION)
}

// Creates a new instance of an hge structure
func newHGE(a ...interface{}) hge {
	ver := VERSION

	if len(a) == 1 {
		if v, ok := a[0].(int); ok {
			ver = v
		}
	}

	h := C.HGE_Create(C.int(ver))

	return h
}

// Releases the memory the C++ library allocated for the HGE struct
func Free() {
	C.HGE_Release(HGE)
}

// Initializes hardware and software needed to run engine.
func Initiate() error {
	if C.HGE_System_Initiate(HGE) == 0 {
		return &Error{}
	}

	return nil
}

//  Restores video mode and frees allocated resources.
func Shutdown() {
	C.HGE_System_Shutdown(HGE)
}

// Starts running user defined frame function.
func Start() error {
	if C.HGE_System_Start(HGE) == 0 {
		return &Error{}
	}

	return nil
}

//  Returns last occured HGE error description.
func GetErrorMessage() string {
	return C.GoString(C.HGE_System_GetErrorMessage(HGE))
}

// Writes a formatted message to the log file.
func Log(format string, v ...interface{}) {
	var str string

	if v != nil {
		str = fmt.Sprintf(format, v...)
	} else {
		str = format
	}

	fstr := C.CString(str)
	defer C.free(unsafe.Pointer(fstr))

	C.goHGE_System_Log(HGE, fstr)
}

// Launches an URL or external executable/data file.
func Launch(url string) bool {
	urlstr := C.CString(url)
	defer C.free(unsafe.Pointer(urlstr))

	return C.HGE_System_Launch(HGE, urlstr) == 1
}

//  Saves current screen snapshot into a file.
func Snapshot(a ...interface{}) {
	if len(a) == 1 {
		if filename, ok := a[0].(string); ok {
			fname := C.CString(filename)
			defer C.free(unsafe.Pointer(fname))

			C.HGE_System_Snapshot(HGE, fname)
			return
		}
	}

	C.HGE_System_Snapshot(HGE, nil)
}

// Sets internal system states.
// First param should be one of: BoolState, IntState, StringState, FuncState, HwndState
// Second parameter must be of the matching type, bool, int, string, StateFunc/func() int, *Hwnd
func SetState(a ...interface{}) {
	if len(a) == 2 {
		switch a[0].(type) {
		case BoolState:
			if bs, ok := a[1].(bool); ok {
				setStateBool(a[0].(BoolState), bs)
				return
			}

		case IntState:
			if is, ok := a[1].(int); ok {
				setStateInt(a[0].(IntState), is)
				return
			}

		case StringState:
			if ss, ok := a[1].(string); ok {
				setStateString(a[0].(StringState), ss)
				return
			}

		case FuncState:
			switch a[1].(type) {
			case StateFunc:
				setStateFunc(a[0].(FuncState), a[1].(StateFunc))
				return
			case func() int:
				setStateFunc(a[0].(FuncState), a[1].(func() int))
				return
			}

		case HwndState:
			if hs, ok := a[1].(*Hwnd); ok {
				setStateHwnd(a[0].(HwndState), hs)
				return
			}
		}
	}
}

func setStateBool(state BoolState, value bool) {
	C.HGE_System_SetStateBool(HGE, C.HGE_BoolState_t(state), BoolToCInt(value))
}

func setStateFunc(state FuncState, value StateFunc) {
	funcCBs[state] = value
	switch state {
	case FRAMEFUNC:
		C.setFrameFunc(HGE, C.HGE_FuncState_t(state))
	case RENDERFUNC:
		C.setRenderFunc(HGE, C.HGE_FuncState_t(state))
	case FOCUSLOSTFUNC:
		C.setFocusLostFunc(HGE, C.HGE_FuncState_t(state))
	case FOCUSGAINFUNC:
		C.setFocusGainFunc(HGE, C.HGE_FuncState_t(state))
	case GFXRESTOREFUNC:
		C.setGfxRestoreFunc(HGE, C.HGE_FuncState_t(state))
	case EXITFUNC:
		C.setExitFunc(HGE, C.HGE_FuncState_t(state))
	}
}

func setStateHwnd(state HwndState, value *Hwnd) {
	C.HGE_System_SetStateHwnd(HGE, C.HGE_HwndState_t(state), value.hwnd)
}

func setStateInt(state IntState, value int) {
	C.HGE_System_SetStateInt(HGE, C.HGE_IntState_t(state), C.int(value))
}

func setStateString(state StringState, value string) {
	val := C.CString(value)
	defer C.free(unsafe.Pointer(val))

	C.HGE_System_SetStateString(HGE, C.HGE_StringState_t(state), val)
}

// Returns internal system state values.
func GetState(a ...interface{}) interface{} {
	if len(a) == 1 {
		switch a[0].(type) {
		case BoolState:
			return getStateBool(a[0].(BoolState))

		case IntState:
			return getStateInt(a[0].(IntState))

		case StringState:
			return getStateString(a[0].(StringState))

		case FuncState:
			return getStateFunc(a[0].(FuncState))

		case HwndState:
			return getStateHwnd(a[0].(HwndState))
		}
	}

	return nil
}

func getStateBool(state BoolState) bool {
	return C.HGE_System_GetStateBool(HGE, C.HGE_BoolState_t(state)) == 1
}

func getStateFunc(state FuncState) StateFunc {
	// I don't know how to convert the HGE_Callback C function type to a Go
	// function, so we just pass back the Go function
	return funcCBs[state]
}

func getStateHwnd(state HwndState) Hwnd {
	var hwnd Hwnd
	hwnd.hwnd = C.HGE_System_GetStateHwnd(HGE, C.HGE_HwndState_t(state))
	return hwnd
}

func getStateInt(state IntState) int {
	return int(C.HGE_System_GetStateInt(HGE, C.HGE_IntState_t(state)))
}

func getStateString(state StringState) string {
	return C.GoString(C.HGE_System_GetStateString(HGE, C.HGE_StringState_t(state)))
}
