package timer

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import "github.com/losinggeneration/hge/binding/hge"

var timerHGE *hge.HGE

func init() {
	timerHGE = hge.New()
}

func Time() float64 {
	return float64(C.HGE_Timer_GetTime(timerHGE.HGE))
}

func Delta() float64 {
	return float64(C.HGE_Timer_GetDelta(timerHGE.HGE))
}

func GetFPS() int {
	return int(C.HGE_Timer_GetFPS(timerHGE.HGE))
}
