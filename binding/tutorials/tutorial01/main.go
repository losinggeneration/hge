package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/binding/hge"
	"github.com/losinggeneration/hge-go/binding/hge/input"
	"os"
)

func FrameFunc() int {
	if input.NewKey(input.K_ESCAPE).State() {
		return 1
	}

	return 0
}

func main() {
	h := hge.New()
	h.SetState(hge.LOGFILE, "tutorial01.log")
	h.SetState(hge.FRAMEFUNC, FrameFunc)
	h.SetState(hge.TITLE, "HGE Tutorial 01 - Minimal HGE application")
	h.SetState(hge.WINDOWED, true)
	h.SetState(hge.USESOUND, false)

	if err := h.Initiate(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	} else {
		defer h.Shutdown()
		h.Start()
	}
}
