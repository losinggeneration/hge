package timer

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import . "github.com/losinggeneration/hge-go/hge"

func Time() float64 {
	return float64(C.HGE_Timer_GetTime(HGE))
}

func Delta() float64 {
	return float64(C.HGE_Timer_GetDelta(HGE))
}

func GetFPS() int {
	return int(C.HGE_Timer_GetFPS(HGE))
}
