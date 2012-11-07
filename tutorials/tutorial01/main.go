package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/hge"
	"github.com/losinggeneration/hge-go/hge/input"
	"os"
	"time"
)

func FrameFunc() bool {
	if input.K_ESCAPE.Down() {
		return true
	}

	// 	if input.K_ESCAPE.State() {
	// 		return true
	// 	}

	return false
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

	fmt.Println("Finished!")
	time.Sleep(1 * time.Second)
}
