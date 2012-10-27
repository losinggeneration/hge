package hge

import (
	"fmt"
	"github.com/losinggeneration/hge-go/hge/rand"
	"log"
	"math"
	"os"
	"time"
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

	boolstate BoolState = iota
)

const (
	FRAMEFUNC      FuncState = iota // func() bool frame function (default: nil) (you MUST set this)
	RENDERFUNC     FuncState = iota // func() bool render function (default: nil)
	FOCUSLOSTFUNC  FuncState = iota // func() bool focus lost function (default: nil)
	FOCUSGAINFUNC  FuncState = iota // func() bool focus gain function (default: nil)
	GFXRESTOREFUNC FuncState = iota // func() bool gfx restore function (default: nil)
	EXITFUNC       FuncState = iota // func() bool exit function (default: nil)

	funcstate FuncState = iota
)

const (
	HWND       HwndState = iota // int		window handle: read only
	HWNDPARENT HwndState = iota // int		parent win handle	(default: 0)

	hwndstate HwndState = iota
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

	intstate IntState = iota
)

const (
	ICON  StringState = iota // string icon resource (default: nil)
	TITLE StringState = iota // string window title (default: "HGE")

	INIFILE StringState = iota // string ini file (default: nil) (meaning no file)
	LOGFILE StringState = iota // string log file (default: nil) (meaning no file)

	stringstate StringState = iota
)

type (
	BoolState   int
	FuncState   int
	HwndState   int
	IntState    int
	StringState int
)

type StateFunc func() bool

var (
	stateBools   = new([boolstate]bool)
	stateFuncs   = new([funcstate]StateFunc)
	stateHwnds   = new([hwndstate]*Hwnd)
	stateInts    = new([intstate]int)
	stateStrings = new([stringstate]string)
)

// HGE_POWERSTATUS system state special constants
const (
	PWR_AC          = iota
	PWR_UNSUPPORTED = iota
)

// HGE struct
type HGE struct {
	log *log.Logger
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

	return h
}

// Initializes hardware and software needed to run engine.
func (h *HGE) Initialize() error {
	h.Log("")
	h.Log("-------------------------------------------------------------------")
	h.Log(" hge-go can be found at http://github.com/losinggeneration/hge-go/")
	h.Log("  Please don't bother Relish Games about the Go port of HGE.")
	h.Log(" They are responsible for the Windows C++ version, not this build.")
	h.Log("-------------------------------------------------------------------")
	h.Log("")

	h.Log("HGE Started...")

	h.Log("hge-unix version: %X.%X.%X", VERSION>>8, (VERSION&0xF0)>>4, VERSION&0xF)

	h.Log("Date: %s", time.Now())

	h.Log("Application: %s", stateStrings[TITLE])

	// Init subsystems
	if err := initNative(); err != nil {
		h.Shutdown()
		return err
	}
	rand.Seed()
	initPowerStatus()
	initInput()
	if err := initGfx(); err != nil {
		h.Shutdown()
		return err
	}
	if err := initSound(); err != nil {
		h.Shutdown()
		return err
	}

	h.Log("Init done.\n")
	return nil
}

//  Restores video mode and frees allocated resources.
func (h *HGE) Shutdown() {
}

// Starts running user defined frame function.
func (h *HGE) Run() error {
	return nil
}

//  Returns last occured HGE error description.
func (h *HGE) GetErrorMessage() string {
	return ""
}

// Writes a formatted message to the log file.
func (h *HGE) Log(format string, v ...interface{}) {
	if h.log != nil {
		h.log.Printf(">> "+format, v...)
	}
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
func (h *HGE) SetState(a ...interface{}) error {
	if len(a) == 2 {
		switch a[0].(type) {
		case BoolState:
			if bs, ok := a[1].(bool); ok {
				return h.setStateBool(a[0].(BoolState), bs)
			}

		case IntState:
			if is, ok := a[1].(int); ok {
				return h.setStateInt(a[0].(IntState), is)
			}

		case StringState:
			if ss, ok := a[1].(string); ok {
				return h.setStateString(a[0].(StringState), ss)
			}

		case FuncState:
			switch a[1].(type) {
			case StateFunc:
				return h.setStateFunc(a[0].(FuncState), a[1].(StateFunc))
			case func() bool:
				return h.setStateFunc(a[0].(FuncState), a[1].(func() bool))
			}

		case HwndState:
			if hs, ok := a[1].(*Hwnd); ok {
				return h.setStateHwnd(a[0].(HwndState), hs)
			}
		}
	}

	return fmt.Errorf("Invalid arguments passed to SetState")
}

func (h *HGE) setStateBool(state BoolState, value bool) error {
	if state >= boolstate || state < 0 {
		h.Log("Invalid bool state")
		return fmt.Errorf("Invald bool state: %d %s", state, value)
	}

	stateBools[state] = value

	return nil
}

func (h *HGE) setStateFunc(state FuncState, value StateFunc) error {
	if state >= funcstate || state < 0 {
		h.Log("Invalid function state")
		return fmt.Errorf("Invald function state: %d %s", state, value)
	}

	stateFuncs[state] = value

	return nil
}

func (h *HGE) setStateHwnd(state HwndState, value *Hwnd) error {
	if state >= hwndstate || state < 0 {
		h.Log("Invalid hwnd state")
		return fmt.Errorf("Invald hwnd state: %d %s", state, value)
	}

	stateHwnds[state] = value

	return nil
}

func (h *HGE) setStateInt(state IntState, value int) error {
	if state >= intstate || state < 0 {
		h.Log("Invalid int state")
		return fmt.Errorf("Invald int state: %d %s", state, value)
	}

	stateInts[state] = value

	return nil
}

func (h *HGE) setStateString(state StringState, value string) error {
	if state >= stringstate || state < 0 {
		h.Log("Invalid string state")
		return fmt.Errorf("Invald string state: %d %s", state, value)
	}

	stateStrings[state] = value

	switch state {
	case LOGFILE:
		l, e := setupLogfile()
		h.log = l
		return e
	}

	return nil
}

// TODO the log file likely needs close called on it at some point
func setupLogfile() (*log.Logger, error) {
	file, err := os.Create(stateStrings[LOGFILE])
	if err != nil {
		return nil, err
	}
	return log.New(file, "<< ", log.LstdFlags), nil
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
	if state >= boolstate || state < 0 {
		return false
	}

	return stateBools[state]
}

func (h *HGE) getStateFunc(state FuncState) StateFunc {
	if state >= funcstate || state < 0 {
		return nil
	}

	return stateFuncs[state]
}

func (h *HGE) getStateHwnd(state HwndState) Hwnd {
	return Hwnd{}
}

func (h *HGE) getStateInt(state IntState) int {
	if state >= intstate || state < 0 {
		return 0
	}

	return stateInts[state]
}

func (h *HGE) getStateString(state StringState) string {
	if state >= stringstate || state < 0 {
		return ""
	}

	return stateStrings[state]
}
