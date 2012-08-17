package rand

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import . "github.com/losinggeneration/hge-go/hge"

func Seed(a ...interface{}) {
	if len(a) == 1 {
		if seed, ok := a[0].(int); ok {
			C.HGE_Random_Seed(HGE, C.int(seed))
			return
		}
	}

	C.HGE_Random_Seed(HGE, 0)
}

func Int(min, max int) int {
	return int(C.HGE_Random_Int(HGE, C.int(min), C.int(max)))
}

func Float32(min, max float32) float32 {
	return float32(C.HGE_Random_Float(HGE, C.float(min), C.float(max)))
}

func Float64(min, max float64) float64 {
	return float64(C.HGE_Random_Float(HGE, C.float(min), C.float(max)))
}

type Rand struct {
	seed int
}

func New(seed int) *Rand {
	return &Rand{seed}
}

func (r *Rand) Seed() {
	C.HGE_Random_Seed(HGE, C.int(r.seed))
}

func (r *Rand) Int(min, max int) int {
	r.Seed()
	return int(C.HGE_Random_Int(HGE, C.int(min), C.int(max)))
}

func (r *Rand) Float32(min, max float32) float32 {
	r.Seed()
	return float32(C.HGE_Random_Float(HGE, C.float(min), C.float(max)))
}

func (r *Rand) Float64(min, max float64) float64 {
	r.Seed()
	return float64(C.HGE_Random_Float(HGE, C.float(min), C.float(max)))
}
