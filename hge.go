// This is mostly a convenience/utility package. It can easily change the
// sub-package states, such as how many frames-per-second should be rendered in
// the gfx package. It also provides some utilities like a main loop and logging
// to a file.
package hge

import (
	"fmt"
	"io"
	"log"
	"math"
	"time"

	"github.com/losinggeneration/hge/gfx"
	"github.com/losinggeneration/hge/input"
	"github.com/losinggeneration/hge/rand"
	"github.com/losinggeneration/hge/sound"
	"github.com/losinggeneration/hge/timer"
)

// The current version of this package
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

// POWERSTATUS system state special constants
const (
	PWR_AC          = -1
	PWR_UNSUPPORTED = -2
)

// FPS system state special constants
const (
	FPS_UNLIMITED = 0
	FPS_VSYNC     = -1
)

type logger struct {
	*log.Logger
	file io.WriteCloser
}

func newLogger(out io.WriteCloser, prefix string, flag int) *logger {
	return &logger{log.New(out, prefix, flag), out}
}

// HGE struct
type HGE struct {
	log        *logger
	last_error error
}

// Creates a new instance of an HGE structure
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

// This will create a shared instance of an HGE structure.
// It's basically a singleton interface.
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
	h.Log(" hge can be found at http://github.com/losinggeneration/hge/")
	h.Log(" This project is based on (mostly the design of) HGE.")
	h.Log(" Please don't bother Relish Games or Icculus about this port.")
	h.Log("-------------------------------------------------------------------")
	h.Log("")

	h.Log("HGE Started...")

	h.Log("hge version: %X.%X.%X", VERSION>>8, (VERSION&0xF0)>>4, VERSION&0xF)

	h.Log("Date: %s", time.Now())

	h.Log("Application: %s", stateStrings[TITLE])

	// Init subsystems
	if err := initNative(h); err != nil {
		h.postError(err)
		h.Shutdown()
		return err
	}

	timer.Reset()
	rand.Seed()
	initPowerStatus()

	if err := input.Initialize(); err != nil {
		h.postError(err)
	}

	// Later on this should be a fatal error
	if err := gfx.Initialize(); err != nil {
		h.postError(err)
		h.Shutdown()
		return err
	}

	// later on this should be a fatal error
	if err := sound.Initialize(); err != nil {
		h.postError(err)
		// h.Shutdown()
		// return err
	}

	h.Log("Init done.\n")
	return nil
}

// Shuts all subsystems down (if needed)
func (h *HGE) Shutdown() {
	shutdownNative()
	h.Log("Finished")
	if h.log != nil {
		h.log.file.Close()
	}
}

// This is the main game loop. It does things like updates the graphics, handles
// user input, and all. The user supplied functions are run if defined as well.
// The frame function must be defined before calling Run.
func (h *HGE) Run() error {
	hwnd := stateHwnds[HWND]
	procFrameFunc := stateFuncs[FRAMEFUNC]

	if hwnd == nil {
		return h.logError("Run: Initiate wasn't called")
	}

	if procFrameFunc == nil {
		return h.logError("Run: No frame function defined")
	}

	active := true
	dt := 0.0

	vsync := stateInts[FPS] == -1

	delta := 1.0 / float64(stateInts[MINDELTATIME])
	timer.Update()

	for {
		dt += timer.Delta()

		input.Process()

		// Check if mouse is over HGE window for Input_IsMouseOver
		input.UpdateMouse()

		// If HGE window is focused or we have the "don't suspend" state - process the main loop
		if active || stateBools[DONTSUSPEND] {
			// If we reached the time for the next frame
			// or we just run in unlimited FPS mode, then
			// do the stuff
			if dt >= delta {
				timer.Update()
				// Do user's stuff
				if procFrameFunc() {
					break
				}
				if stateFuncs[RENDERFUNC] != nil {
					if stateFuncs[RENDERFUNC]() {
						break
					}
				}

				// Clean up input events that were generated by
				// WindowProc and weren't handled by user's code
				input.ClearQueue()
				dt = 0
			}
		}

		if vsync || (!active && !stateBools[DONTSUSPEND]) {
			time.Sleep(1 * time.Millisecond)
		}
	}

	input.ClearQueue()

	return nil
}

// Returns last occurred HGE error description.
func (h *HGE) GetErrorMessage() string {
	msg := fmt.Sprint(h.last_error)
	h.last_error = nil
	return msg
}

// Convenience to pass error on to Logging function
func (h *HGE) postError(e error) error {
	return h.logError(e.Error())
}

// Log the error, set last_error, and return the error to the user
func (h *HGE) logError(format string, v ...interface{}) error {
	h.Log(format, v...)
	h.last_error = fmt.Errorf(format, v...)
	return h.last_error
}

// Writes a formatted message to the log file.
func (h *HGE) Log(format string, v ...interface{}) {
	if h.log != nil {
		h.log.Printf(">> "+format, v...)
	} else {
		// log to stdout if there's no log file set
		fmt.Printf(">> "+format+"\n", v...)
	}
}

// Launches an URL or external executable/data file.
// TODO
func (h *HGE) Launch(url string) bool {
	return true
}

// Saves current screen snapshot into a file.
// TODO
func (h *HGE) Snapshot(a ...interface{}) {
}
