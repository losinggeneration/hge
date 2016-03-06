package main

import (
	"fmt"
	"os"

	"github.com/losinggeneration/hge"
	"github.com/losinggeneration/hge/gfx"
	"github.com/losinggeneration/hge/input"
)

func Frame() bool {
	key := input.GetKey()
	if key.Down() {
		fmt.Println("Key:", key.Name())
	}

	if input.K_ESCAPE.Down() {
		return true
	}

	return false
}

func Render() bool {
	gfx.BeginScene()
	gfx.EndScene()

	return false
}

func main() {
	h := hge.New()
	h.SetState(hge.LOGFILE, "tutorial01.log")
	h.SetState(hge.FRAMEFUNC, Frame)
	h.SetState(hge.RENDERFUNC, Render)
	h.SetState(hge.TITLE, "HGE Tutorial 01 - Minimal HGE application")
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
}
