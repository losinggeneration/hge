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

// HGE Handle types
type (
	Texture C.HTEXTURE
	Target  C.HTARGET
	Effect  C.HEFFECT
	Music   C.HMUSIC
	Stream  C.HSTREAM
	Channel C.HCHANNEL
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

// HGE Blending constants
const (
	BLEND_COLORADD   = C.BLEND_COLORADD
	BLEND_COLORMUL   = C.BLEND_COLORMUL
	BLEND_ALPHABLEND = C.BLEND_ALPHABLEND
	BLEND_ALPHAADD   = C.BLEND_ALPHAADD
	BLEND_ZWRITE     = C.BLEND_ZWRITE
	BLEND_NOZWRITE   = C.BLEND_NOZWRITE

	BLEND_DEFAULT   = C.BLEND_DEFAULT
	BLEND_DEFAULT_Z = C.BLEND_DEFAULT_Z
)

// HGE System state constants
var (
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

// HGE_FPS system state special constants
const (
	FPS_UNLIMITED = C.HGE_FPS_UNLIMITED
	FPS_VSYNC     = C.HGE_FPS_VSYNC
)

// HGE_POWERSTATUS system state special constants
const (
	PWR_AC          = C.HGE_PWR_AC
	PWR_UNSUPPORTED = C.HGE_PWR_UNSUPPORTED
)

// HGE Primitive type constants
const (
	PRIM_LINES   = C.HGE_PRIM_LINES
	PRIM_TRIPLES = C.HGE_PRIM_TRIPLES
	PRIM_QUADS   = C.HGE_PRIM_QUADS
)

// HGE Vertex structure
type Vertex struct {
	X, Y   float32 // screen position
	Z      float32 // Z-buffer depth 0..1
	Col    Dword   // color
	TX, TY float32 // texture coordinates
}

// HGE Triple structure
type Triple struct {
	V     [3]Vertex
	Tex   Texture
	Blend int
}

// HGE Quad structure
type Quad struct {
	V     [4]Vertex
	Tex   Texture
	Blend int
}

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

type Resource uintptr

func btoi(b bool) C.BOOL {
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

var h hge

func init() {
	h = newHGE(VERSION)
}

// Creates a new instance of an hge structure
func newHGE(a ...interface{}) hge {
	ver := VERSION

	if len(a) == 1 {
		if v, ok := a[0].(int); ok {
			ver = v
		}
	}

	h = C.HGE_Create(C.int(ver))

	return h
}

// Releases the memory the C++ library allocated for the HGE struct
func Free() {
	C.HGE_Release(h)
}

// Initializes hardware and software needed to run engine.
func Initiate() error {
	if C.HGE_System_Initiate(h) == 0 {
		return &Error{}
	}

	return nil
}

//  Restores video mode and frees allocated resources.
func Shutdown() {
	C.HGE_System_Shutdown(h)
}

// Starts running user defined frame function.
func Start() error {
	if C.HGE_System_Start(h) == 0 {
		return &Error{}
	}

	return nil
}

//  Returns last occured HGE error description.
func GetErrorMessage() string {
	return C.GoString(C.HGE_System_GetErrorMessage(h))
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

	C.goHGE_System_Log(h, fstr)
}

// Launches an URL or external executable/data file.
func Launch(url string) bool {
	urlstr := C.CString(url)
	defer C.free(unsafe.Pointer(urlstr))

	return C.HGE_System_Launch(h, urlstr) == 1
}

//  Saves current screen snapshot into a file.
func Snapshot(a ...interface{}) {
	if len(a) == 1 {
		if filename, ok := a[0].(string); ok {
			fname := C.CString(filename)
			defer C.free(unsafe.Pointer(fname))

			C.HGE_System_Snapshot(h, fname)
			return
		}
	}

	C.HGE_System_Snapshot(h, nil)
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
	C.HGE_System_SetStateBool(h, C.HGE_BoolState_t(state), btoi(value))
}

func setStateFunc(state FuncState, value StateFunc) {
	funcCBs[state] = value
	switch state {
	case FRAMEFUNC:
		C.setFrameFunc(h, C.HGE_FuncState_t(state))
	case RENDERFUNC:
		C.setRenderFunc(h, C.HGE_FuncState_t(state))
	case FOCUSLOSTFUNC:
		C.setFocusLostFunc(h, C.HGE_FuncState_t(state))
	case FOCUSGAINFUNC:
		C.setFocusGainFunc(h, C.HGE_FuncState_t(state))
	case GFXRESTOREFUNC:
		C.setGfxRestoreFunc(h, C.HGE_FuncState_t(state))
	case EXITFUNC:
		C.setExitFunc(h, C.HGE_FuncState_t(state))
	}
}

func setStateHwnd(state HwndState, value *Hwnd) {
	C.HGE_System_SetStateHwnd(h, C.HGE_HwndState_t(state), value.hwnd)
}

func setStateInt(state IntState, value int) {
	C.HGE_System_SetStateInt(h, C.HGE_IntState_t(state), C.int(value))
}

func setStateString(state StringState, value string) {
	val := C.CString(value)
	defer C.free(unsafe.Pointer(val))

	C.HGE_System_SetStateString(h, C.HGE_StringState_t(state), val)
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
	return C.HGE_System_GetStateBool(h, C.HGE_BoolState_t(state)) == 1
}

func getStateFunc(state FuncState) StateFunc {
	// I don't know how to convert the HGE_Callback C function type to a Go
	// function, so we just pass back the Go function
	return funcCBs[state]
}

func getStateHwnd(state HwndState) Hwnd {
	var hwnd Hwnd
	hwnd.hwnd = C.HGE_System_GetStateHwnd(h, C.HGE_HwndState_t(state))
	return hwnd
}

func getStateInt(state IntState) int {
	return int(C.HGE_System_GetStateInt(h, C.HGE_IntState_t(state)))
}

func getStateString(state StringState) string {
	return C.GoString(C.HGE_System_GetStateString(h, C.HGE_StringState_t(state)))
}

// Loads a resource into memory from disk.
func NewResource(filename string) (*Resource, Dword) {
	var s C.DWORD
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	r := Resource(C.HGE_Resource_Load(h, fname, &s))

	return &r, Dword(s)
}

// Deletes a previously loaded resource from memory.
func (r Resource) Free() {
	C.HGE_Resource_Free(h, unsafe.Pointer(r))
}

// Loads a resource, puts the loaded data into a byte array, and frees the data.
func LoadBytes(filename string) []byte {
	r, size := NewResource(filename)

	if r == nil {
		return nil
	}

	b := C.GoBytes(unsafe.Pointer(*r), C.int(size))
	r.Free()

	return b
}

// Loads a resource, puts the data into a string, and frees the data.
func LoadString(filename string) *string {
	r, size := NewResource(filename)

	if r == nil {
		return nil
	}

	s := C.GoStringN((*C.char)(unsafe.Pointer(*r)), C.int(size))
	r.Free()

	return &s
}

// Attaches a resource pack.
func (_ Resource) AttachPack(filename string, a ...interface{}) bool {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	if len(a) == 1 {
		var password *C.char

		switch a[0].(type) {
		case string:
			password = C.CString(a[0].(string))
			defer C.free(unsafe.Pointer(password))
		}

		return C.HGE_Resource_AttachPack(h, fname, password) == 1
	}

	return C.HGE_Resource_AttachPack(h, fname, nil) == 1
}

// Removes a resource pack.
func (_ Resource) RemovePack(filename string) {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	C.HGE_Resource_RemovePack(h, fname)
}

// Removes all resource packs previously attached.
func (_ Resource) RemoveAllPacks() {
	C.HGE_Resource_RemoveAllPacks(h)
}

// Builds absolute file path.
func (_ Resource) MakePath(a ...interface{}) string {
	if len(a) == 1 {
		if filename, ok := a[0].(string); ok {
			fname := C.CString(filename)
			defer C.free(unsafe.Pointer(fname))

			return C.GoString(C.HGE_Resource_MakePath(h, fname))
		}
	}

	return C.GoString(C.HGE_Resource_MakePath(h, nil))
}

// Enumerates files by given wildcard.
func (_ Resource) EnumFiles(a ...interface{}) string {
	if len(a) == 1 {
		if wildcard, ok := a[0].(string); ok {
			wcard := C.CString(wildcard)
			defer C.free(unsafe.Pointer(wcard))

			return C.GoString(C.HGE_Resource_EnumFiles(h, wcard))
		}
	}

	return C.GoString(C.HGE_Resource_EnumFiles(h, nil))
}

// Enumerates folders by given wildcard.
func (_ Resource) EnumFolders(a ...interface{}) string {
	if len(a) == 1 {
		if wildcard, ok := a[0].(string); ok {
			wcard := C.CString(wildcard)
			defer C.free(unsafe.Pointer(wcard))

			return C.GoString(C.HGE_Resource_EnumFolders(h, wcard))
		}
	}

	return C.GoString(C.HGE_Resource_EnumFolders(h, nil))
}

type Ini struct {
	Section, Name string
}

func NewIni(section, name string) Ini {
	return Ini{section, name}
}

func (i Ini) SetInt(value int) {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	C.HGE_Ini_SetInt(h, s, n, C.int(value))
}

func (i Ini) GetInt(def_val int) int {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	return int(C.HGE_Ini_GetInt(h, s, n, C.int(def_val)))
}

func (i Ini) SetFloat(value float64) {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	C.HGE_Ini_SetFloat(h, s, n, C.float(value))
}

func (i Ini) GetFloat(def_val float64) float64 {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	return float64(C.HGE_Ini_GetFloat(h, s, n, C.float(def_val)))
}

func (i Ini) SetString(value string) {
	s, n, v := C.CString(i.Section), C.CString(i.Name), C.CString(value)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(v))

	C.HGE_Ini_SetString(h, s, n, v)
}

func (i Ini) GetString(def_val string) string {
	s, n, df := C.CString(i.Section), C.CString(i.Name), C.CString(def_val)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(df))

	return C.GoString(C.HGE_Ini_GetString(h, s, n, df))
}

func RandomSeed(a ...interface{}) {
	if len(a) == 1 {
		if seed, ok := a[0].(int); ok {
			C.HGE_Random_Seed(h, C.int(seed))
			return
		}
	}

	C.HGE_Random_Seed(h, 0)
}

func RandomInt(min, max int) int {
	return int(C.HGE_Random_Int(h, C.int(min), C.int(max)))
}

func RandomFloat(min, max float64) float64 {
	return float64(C.HGE_Random_Float(h, C.float(min), C.float(max)))
}

type Timer float64

func NewTimer(a ...interface{}) *Timer {
	t := new(Timer)
	*t = 0.0
	if len(a) == 1 {
		switch a[0].(type) {
		case float64:
			*t = Timer(a[0].(float64))
		case float32:
			*t = Timer(a[0].(float32))
		}
	}
	return t
}

func (t *Timer) Time() float64 {
	*t = Timer(C.HGE_Timer_GetTime(h))
	return float64(*t)
}

func (t *Timer) Delta() float64 {
	*t = Timer(C.HGE_Timer_GetDelta(h))
	return float64(*t)
}

func GetFPS() int {
	return int(C.HGE_Timer_GetFPS(h))
}

func NewEffect(filename string, a ...interface{}) Effect {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	if len(a) == 1 {
		if size, ok := a[0].(Dword); ok {
			return Effect(C.HGE_Effect_Load(h, fname, C.DWORD(size)))
		}
	}

	return Effect(C.HGE_Effect_Load(h, fname, 0))
}

func (e Effect) Free() {
	C.HGE_Effect_Free(h, (C.HEFFECT(e)))
}

func (e Effect) Play() Channel {
	return Channel(C.HGE_Effect_Play(h, C.HEFFECT(e)))
}

func (e Effect) PlayEx(a ...interface{}) Channel {
	volume, pan := 100, 0
	pitch := 1.0
	loop := false

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if v, ok := a[i].(int); ok {
				volume = v
			}
		}
		if i == 1 {
			if p, ok := a[i].(int); ok {
				pan = p
			}
		}
		if i == 2 {
			if p, ok := a[i].(float32); ok {
				pitch = float64(p)
			}
			if p, ok := a[i].(float64); ok {
				pitch = p
			}
		}
		if i == 3 {
			if l, ok := a[i].(bool); ok {
				loop = l
			}
		}
	}

	return Channel(C.HGE_Effect_PlayEx(h, C.HEFFECT(e), C.int(volume), C.int(pan), C.float(pitch), btoi(loop)))
}

func NewMusic(filename string, size Dword) Music {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	return Music(C.HGE_Music_Load(h, fname, C.DWORD(size)))
}

func (m Music) Free() {
	C.HGE_Music_Free(h, C.HMUSIC(m))
}

func (m Music) Play(loop bool, a ...interface{}) Channel {
	volume, order, row := 100, -1, -1

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if v, ok := a[i].(int); ok {
				volume = v
			}
		}
		if i == 1 {
			if o, ok := a[i].(int); ok {
				order = o
			}
		}
		if i == 2 {
			if r, ok := a[i].(int); ok {
				row = r
			}
		}
	}

	return Channel(C.HGE_Music_Play(h, C.HMUSIC(m), btoi(loop), C.int(volume), C.int(order), C.int(row)))
}

func (m Music) SetAmplification(ampl int) {
	C.HGE_Music_SetAmplification(h, C.HMUSIC(m), C.int(ampl))
}

func (m Music) Amplification() int {
	return int(C.HGE_Music_GetAmplification(h, C.HMUSIC(m)))
}

func (m Music) Len() int {
	return int(C.HGE_Music_GetLength(h, C.HMUSIC(m)))
}

func (m Music) SetPos(order, row int) {
	C.HGE_Music_SetPos(h, C.HMUSIC(m), C.int(order), C.int(row))
}

func (m Music) Pos() (order, row int, ok bool) {
	var o, r C.int

	ok = C.HGE_Music_GetPos(h, C.HMUSIC(m), &o, &r) == 1

	return int(o), int(r), ok
}

func (m Music) SetInstrVolume(instr int, volume int) {
	C.HGE_Music_SetInstrVolume(h, C.HMUSIC(m), C.int(instr), C.int(volume))
}

func (m Music) InstrVolume(instr int) int {
	return int(C.HGE_Music_GetInstrVolume(h, C.HMUSIC(m), C.int(instr)))
}

func (m Music) SetChannelVolume(channel, volume int) {
	C.HGE_Music_SetChannelVolume(h, C.HMUSIC(m), C.int(channel), C.int(volume))
}

func (m Music) ChannelVolume(channel int) int {
	return int(C.HGE_Music_GetChannelVolume(h, C.HMUSIC(m), C.int(channel)))
}

func NewStream(filename string, size Dword) Stream {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	return Stream(C.HGE_Stream_Load(h, fname, C.DWORD(size)))
}

func (s Stream) Free() {
	C.HGE_Stream_Free(h, C.HSTREAM(s))
}

func (s Stream) Play(loop bool, a ...interface{}) Channel {
	volume := 100

	if len(a) == 1 {
		if v, ok := a[0].(int); ok {
			volume = v
		}
	}

	return Channel(C.HGE_Stream_Play(h, C.HSTREAM(s), btoi(loop), C.int(volume)))
}

func (c Channel) SetPanning(pan int) {
	C.HGE_Channel_SetPanning(h, C.HCHANNEL(c), C.int(pan))
}

func (c Channel) SetVolume(volume int) {
	C.HGE_Channel_SetVolume(h, C.HCHANNEL(c), C.int(volume))
}

func (c Channel) SetPitch(pitch float64) {
	C.HGE_Channel_SetPitch(h, C.HCHANNEL(c), C.float(pitch))
}

func (c Channel) Pause() {
	C.HGE_Channel_Pause(h, C.HCHANNEL(c))
}

func (c Channel) Resume() {
	C.HGE_Channel_Resume(h, C.HCHANNEL(c))
}

func (c Channel) Stop() {
	C.HGE_Channel_Stop(h, C.HCHANNEL(c))
}

func (c Channel) PauseAll() {
	C.HGE_Channel_PauseAll(h)
}

func (c Channel) ResumeAll() {
	C.HGE_Channel_ResumeAll(h)
}

func (c Channel) StopAll() {
	C.HGE_Channel_StopAll(h)
}

func (c Channel) IsPlaying() bool {
	return C.HGE_Channel_IsPlaying(h, C.HCHANNEL(c)) == 1
}

func (c Channel) Len() float64 {
	return float64(C.HGE_Channel_GetLength(h, C.HCHANNEL(c)))
}

func (c Channel) Pos() float64 {
	return float64(C.HGE_Channel_GetPos(h, C.HCHANNEL(c)))
}

func (c Channel) SetPos(seconds float64) {
	C.HGE_Channel_SetPos(h, C.HCHANNEL(c), C.float(seconds))
}

func (c Channel) SlideTo(time float64, a ...interface{}) {
	volume, pan := 100, 0
	pitch := 1.0

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if v, ok := a[i].(int); ok {
				volume = v
			}
		}
		if i == 1 {
			if p, ok := a[i].(int); ok {
				pan = p
			}
		}
		if i == 2 {
			if p, ok := a[i].(float32); ok {
				pitch = float64(p)
			}
			if p, ok := a[i].(float64); ok {
				pitch = p
			}
		}
	}

	C.HGE_Channel_SlideTo(h, C.HCHANNEL(c), C.float(time), C.int(volume), C.int(pan), C.float(pitch))
}

func (c Channel) IsSliding() bool {
	return C.HGE_Channel_IsSliding(h, C.HCHANNEL(c)) == 1
}

func MousePos() (x, y float64) {
	var nx, ny C.float

	C.HGE_Input_GetMousePos(h, &nx, &ny)

	return float64(nx), float64(ny)
}

func SetMousePos(x, y float64) {
	C.HGE_Input_SetMousePos(h, C.float(x), C.float(y))
}

func MouseWheel() int {
	return int(C.HGE_Input_GetMouseWheel(h))
}

func IsMouseOver() bool {
	return C.HGE_Input_IsMouseOver(h) == 1
}

func NewKey(i int) Key {
	return Key(i)
}

func (k Key) Down() bool {
	return C.HGE_Input_KeyDown(h, C.int(k)) == 1
}

func (k Key) Up() bool {
	return C.HGE_Input_KeyUp(h, C.int(k)) == 1
}

func (k Key) State() bool {
	return C.HGE_Input_GetKeyState(h, C.int(k)) == 1
}
func (k Key) Name() string {
	return C.GoString(C.HGE_Input_GetKeyName(h, C.int(k)))
}

func GetKey() Key {
	return Key(C.HGE_Input_GetKey(h))
}

func GetChar() int {
	return int(C.HGE_Input_GetChar(h))
}

func GetEvent() (e *InputEvent, b bool) {
	b = C.HGE_Input_GetEvent(h, (*C.HGE_InputEvent_t)(unsafe.Pointer(e))) == 1
	return e, b
}

func GfxBeginScene(a ...interface{}) bool {
	if len(a) == 1 {
		if target, ok := a[0].(Target); ok {
			return C.HGE_Gfx_BeginScene(h, C.HTARGET(target)) == 1
		}
	}

	return C.HGE_Gfx_BeginScene(h, 0) == 1
}

func GfxEndScene() {
	C.HGE_Gfx_EndScene(h)
}

func GfxClear(color Dword) {
	C.HGE_Gfx_Clear(h, C.DWORD(color))
}

func GfxRenderLine(x1, y1, x2, y2 float64, a ...interface{}) {
	color := uint(0xFFFFFFFF)
	z := 0.5

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if c, ok := a[i].(uint); ok {
				color = c
			}
		}
		if i == 1 {
			if z1, ok := a[i].(float32); ok {
				z = float64(z1)
			}
			if z1, ok := a[i].(float64); ok {
				z = z1
			}
		}
	}

	C.HGE_Gfx_RenderLine(h, C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.DWORD(color), C.float(z))
}

func (t *Triple) Render() {
	C.HGE_Gfx_RenderTriple(h, (*C.HGE_Triple_t)(unsafe.Pointer(t)))
}

func (q *Quad) Render() {
	C.HGE_Gfx_RenderQuad(h, (*C.HGE_Quad_t)(unsafe.Pointer(q)))
}

func GfxStartBatch(prim_type int, tex Texture, blend int) (ver *Vertex, max_prim int, ok bool) {
	mp := C.int(0)

	v := C.HGE_Gfx_StartBatch(h, C.int(prim_type), C.HTEXTURE(tex), C.int(blend), &mp)

	if v == nil {
		return nil, 0, false
	}

	return (*Vertex)(unsafe.Pointer(v)), int(mp), true
}

func GfxFinishBatch(prim int) {
	C.HGE_Gfx_FinishBatch(h, C.int(prim))
}

func GfxSetClipping(a ...interface{}) {
	var x, y, w, hi int

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if x1, ok := a[i].(int); ok {
				x = x1
			}
		}
		if i == 1 {
			if y1, ok := a[i].(int); ok {
				y = y1
			}
		}
		if i == 2 {
			if w1, ok := a[i].(int); ok {
				w = w1
			}
		}
		if i == 3 {
			if h1, ok := a[i].(int); ok {
				hi = h1
			}
		}
	}

	C.HGE_Gfx_SetClipping(h, C.int(x), C.int(y), C.int(w), C.int(hi))
}

func GfxSetTransform(a ...interface{}) {
	var (
		x, y, dx, dy        float64
		rot, hscale, vscale float64
	)

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if x1, ok := a[i].(float32); ok {
				x = float64(x1)
			}
			if x1, ok := a[i].(float64); ok {
				x = x1
			}
		}
		if i == 1 {
			if y1, ok := a[i].(float32); ok {
				y = float64(y1)
			}
			if y1, ok := a[i].(float64); ok {
				y = y1
			}
		}
		if i == 2 {
			if dx1, ok := a[i].(float32); ok {
				dx = float64(dx1)
			}
			if dx1, ok := a[i].(float64); ok {
				dx = dx1
			}
		}
		if i == 3 {
			if dy1, ok := a[i].(float32); ok {
				dy = float64(dy1)
			}
			if dy1, ok := a[i].(float64); ok {
				dy = dy1
			}
		}
		if i == 4 {
			if rot1, ok := a[i].(float32); ok {
				rot = float64(rot1)
			}
			if rot1, ok := a[i].(float64); ok {
				rot = rot1
			}
		}
		if i == 5 {
			if hscale1, ok := a[i].(float32); ok {
				hscale = float64(hscale1)
			}
			if hscale1, ok := a[i].(float64); ok {
				hscale = hscale1
			}
		}
		if i == 6 {
			if vscale1, ok := a[i].(float32); ok {
				vscale = float64(vscale1)
			}
			if vscale1, ok := a[i].(float64); ok {
				vscale = vscale1
			}
		}
	}

	C.HGE_Gfx_SetTransform(h, C.float(x), C.float(y), C.float(dx), C.float(dy), C.float(rot), C.float(hscale), C.float(vscale))
}

func NewTarget(width, height int, zbuffer bool) Target {
	return Target(C.HGE_Target_Create(h, C.int(width), C.int(height), btoi(zbuffer)))
}

func (t Target) Free() {
	C.HGE_Target_Free(h, C.HTARGET(t))
}

func (t Target) Texture() Texture {
	return Texture(C.HGE_Target_GetTexture(h, C.HTARGET(t)))
}

func NewTexture(width, height int) Texture {
	return Texture(C.HGE_Texture_Create(h, C.int(width), C.int(height)))
}

func LoadTexture(filename string, a ...interface{}) Texture {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	size := Dword(0)
	mipmap := false

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if s, ok := a[i].(Dword); ok {
				size = s
			}
		}
		if i == 1 {
			if m, ok := a[i].(bool); ok {
				mipmap = m
			}
		}
	}

	return Texture(C.HGE_Texture_Load(h, fname, C.DWORD(size), btoi(mipmap)))
}

func (t Texture) Free() {
	C.HGE_Texture_Free(h, C.HTEXTURE(t))
}

func (t Texture) Width(a ...interface{}) int {
	if len(a) == 1 {
		if original, ok := a[0].(bool); ok {
			return int(C.HGE_Texture_GetWidth(h, C.HTEXTURE(t), btoi(original)))
		}
	}

	return int(C.HGE_Texture_GetWidth(h, C.HTEXTURE(t), btoi(false)))
}

func (t Texture) Height(a ...interface{}) int {
	if len(a) == 1 {
		if original, ok := a[0].(bool); ok {
			return int(C.HGE_Texture_GetWidth(h, C.HTEXTURE(t), btoi(original)))
		}
	}

	return int(C.HGE_Texture_GetHeight(h, C.HTEXTURE(t), btoi(false)))
}

func (t Texture) Lock(a ...interface{}) *Dword {
	readonly := true
	left, top, width, height := 0, 0, 0, 0

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if r, ok := a[i].(bool); ok {
				readonly = r
			}
		}
		if i == 1 {
			if l, ok := a[i].(int); ok {
				left = l
			}
		}
		if i == 2 {
			if t, ok := a[i].(int); ok {
				top = t
			}
		}
		if i == 3 {
			if w, ok := a[i].(int); ok {
				width = w
			}
		}
		if i == 4 {
			if h, ok := a[i].(int); ok {
				height = h
			}
		}
	}

	d := C.HGE_Texture_Lock(h, C.HTEXTURE(t), btoi(readonly), C.int(left), C.int(top), C.int(width), C.int(height))
	return (*Dword)(d)
}

func (t Texture) Unlock() {
	C.HGE_Texture_Unlock(h, C.HTEXTURE(t))
}

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
