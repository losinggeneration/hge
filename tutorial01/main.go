package main

import (
	"fmt"
	"hge"
)

var h *hge.HGE

func FrameFunc() int {
	if h.Input_GetKeyState(hge.K_ESCAPE) {
		return 1
	}

	return 0
}

func main() {
	h = hge.Create(hge.VERSION)
	defer h.Release()

	h.System_SetState(hge.FRAMEFUNC, FrameFunc)
	h.System_SetState(hge.TITLE, "HGE Tutorial 01 - Minimal HGE application")
	h.System_SetState(hge.WINDOWED, true)
	h.System_SetState(hge.USESOUND, false)

	h.System_Log("Test")
	h.System_Log("Test vararg: %s %d", "test", 15)

	if h.System_Initiate() {
		h.System_Log("Test")
		h.System_Log("Test vararg: %s %d", "test", 15)
		h.System_Start()
	} else {
		fmt.Println("Error: ", h.System_GetErrorMessage())
	}

	h.System_Log("Test")
	h.System_Shutdown()
}
