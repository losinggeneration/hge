package hge

import (
	"log"
	"os"
)

//////// Bool states
func setupWindowed(h *HGE) error {
	return nil
}

func setupZBuffer(h *HGE) error {
	return nil
}

func setupTextureFilter(h *HGE) error {
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
func setupHwndParent(h *HGE) error {
	return nil
}

//////// Int states

func setupScreenWidth(h *HGE) error {
	return nil
}

func setupScreenHeight(h *HGE) error {
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
	file, err := os.Create(stateStrings[LOGFILE])
	if err != nil {
		return h.postError(err)
	}
	h.log = log.New(file, "<< ", log.LstdFlags)
	return nil
}
