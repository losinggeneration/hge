package hge

import (
	"github.com/losinggeneration/hge-go/hge/gfx"
	"github.com/losinggeneration/hge-go/hge/input"
	"github.com/losinggeneration/hge-go/hge/rand"
	"github.com/losinggeneration/hge-go/hge/sound"
	"log"
	"math"
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

type Hwnd struct {
}

// POWERSTATUS system state special constants
const (
	PWR_AC          = iota
	PWR_UNSUPPORTED = iota
)

// FPS system state special constants
const (
	FPS_UNLIMITED = iota
	FPS_VSYNC     = iota
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

	h.setDefaultStates()

	return h
}

var singleton *HGE = nil

func Shared(a ...interface{}) *HGE {
	if singleton == nil {
		singleton = New(a...)
	}

	return singleton
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
	input.Initialize()

	if err := gfx.Initialize(); err != nil {
		h.Shutdown()
		return err
	}

	if err := sound.Initialize(); err != nil {
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
