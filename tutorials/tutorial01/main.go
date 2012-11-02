package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/hge"
	"github.com/losinggeneration/hge-go/hge/input"
	"os"
	"time"
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

	h.Log("Test")
	h.Log("Test vararg: %s %d", "test", 15)

	if err := h.Initiate(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	} else {
		defer h.Shutdown()
		h.Log("Test")
		h.Log("Test vararg: %s %d", "test", 15)
		h.Start()
	}

	h.Log("Test")
	time.Sleep(1 * time.Second)
}
