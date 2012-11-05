package hge

import "C"

type StateFunc func() int

var funcCBs []StateFunc = make([]StateFunc, EXITFUNC+1)

//export goFrameFunc
func goFrameFunc() int {
	return funcCBs[FRAMEFUNC]()
}

//export goRenderFunc
func goRenderFunc() int {
	return funcCBs[RENDERFUNC]()
}

//export goFocusLostFunc
func goFocusLostFunc() int {
	return funcCBs[FOCUSLOSTFUNC]()
}

//export goFocusGainFunc
func goFocusGainFunc() int {
	return funcCBs[FOCUSGAINFUNC]()
}

//export goGfxRestoreFunc
func goGfxRestoreFunc() int {
	return funcCBs[GFXRESTOREFUNC]()
}

//export goExitFunc
func goExitFunc() int {
	return funcCBs[EXITFUNC]()
}
