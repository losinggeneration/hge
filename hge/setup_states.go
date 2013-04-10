package hge

import (
	"github.com/losinggeneration/hge-go/hge/gfx"
	"log"
	"os"
)

//////// Bool states
func setupWindowed(h *HGE) error {
	return nil
}

func setupZBuffer(h *HGE) error {
	gfx.SetZBuffer(stateBools[ZBUFFER])
	return nil
}

func setupTextureFilter(h *HGE) error {
	gfx.SetTextureFilter(stateBools[TEXTUREFILTER])
	return nil
}

func setupUseSound(h *HGE) error {
	return nil
}

func setupDonetSuspend(h *HGE) error {
	return nil
}

func setupHideMouse(h *HGE) error {
	return nil
}

func setupShowSplash(h *HGE) error {
	return nil
}

//////// Func states: no setup needed

//////// Hwnd states
func setupHwnd(h *HGE) error {
	return nil
}

func setupHwndParent(h *HGE) error {
	return nil
}

//////// Int states

func setupScreenWidth(h *HGE) error {
	gfx.SetWidth(stateInts[SCREENWIDTH])
	return nil
}

func setupScreenHeight(h *HGE) error {
	gfx.SetHeight(stateInts[SCREENHEIGHT])
	return nil
}

func setupScreenBPP(h *HGE) error {
	return nil
}

func setupOrigScreenWidth(h *HGE) error {
	return nil
}

func setupOrigScreenHeight(h *HGE) error {
	return nil
}

func setupFPS(h *HGE) error {
	return nil
}

func setupMinDeltaTime(h *HGE) error {
	if stateInts[MINDELTATIME] < 1 {
		stateInts[MINDELTATIME] = 1000
		return h.logError("Error: State MINDELTATIME must not be less than 1. Setting to default: 1000")
	}
	return nil
}

func setupSampleRate(h *HGE) error {
	return nil
}

func setupFxVolume(h *HGE) error {
	return nil
}

func setupMusVolume(h *HGE) error {
	return nil
}

func setupStreamVolume(h *HGE) error {
	return nil
}

func setupPowerStatus(h *HGE) error {
	return nil
}

//////// String State setups

func setupIcon(h *HGE) error {
	return nil
}

func setupTitle(h *HGE) error {
	setTitle()
	return nil
}

func setupInifile(h *HGE) error {
	return nil
}

// TODO the log file likely needs close called on it at some point
func setupLogfile(h *HGE) error {
	if stateStrings[LOGFILE] == "" {
		if h.log != nil {
			h.log.file.Close()
			h.log = nil
		}
		return nil
	}

	file, err := os.Create(stateStrings[LOGFILE])
	if err != nil {
		return h.postError(err)
	}
	h.log = newLogger(file, "<< ", log.LstdFlags)
	return nil
}
