package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/hge"
	"os"
)

func FrameFunc() int {
	if hge.NewKey(hge.K_ESCAPE).State() {
		return 1
	}

	return 0
}

func main() {
	defer hge.Free()

	hge.SetState(hge.LOGFILE, "tutorial01.log")
	hge.SetState(hge.FRAMEFUNC, FrameFunc)
	hge.SetState(hge.TITLE, "HGE Tutorial 01 - Minimal HGE application")
	hge.SetState(hge.WINDOWED, true)
	hge.SetState(hge.USESOUND, false)

	hge.Log("Test")
	hge.Log("Test vararg: %s %d", "test", 15)

	if err := hge.Initiate(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	} else {
		defer hge.Shutdown()
		hge.Log("Test")
		hge.Log("Test vararg: %s %d", "test", 15)
		hge.Start()
	}

	hge.Log("Test")
}
