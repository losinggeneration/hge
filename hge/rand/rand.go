package rand

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import "github.com/losinggeneration/hge-go/hge"

func Seed(a ...interface{}) {
	seed := 1
	if len(a) == 1 {
		if s, ok := a[0].(int); ok {
			seed = s
		}
		if s, ok := a[0].(int64); ok {
			seed = int(s)
		}
	}

	New(seed).Seed()
}

func Int(min, max int) int {
	return New(0).Int(min, max)
}

func Float32(min, max float32) float32 {
	return New(0).Float32(min, max)
}

func Float64(min, max float64) float64 {
	return New(0).Float64(min, max)
}

type Rand struct {
	seed    int
	randHGE *hge.HGE
}

func New(seed int) *Rand {
	return &Rand{seed, hge.New()}
}

func (r *Rand) Seed() {
	C.HGE_Random_Seed(r.randHGE.HGE, C.int(r.seed))
}

func (r *Rand) Int(min, max int) int {
	r.Seed()
	return int(C.HGE_Random_Int(r.randHGE.HGE, C.int(min), C.int(max)))
}

func (r *Rand) Float32(min, max float32) float32 {
	r.Seed()
	return float32(C.HGE_Random_Float(r.randHGE.HGE, C.float(min), C.float(max)))
}

func (r *Rand) Float64(min, max float64) float64 {
	r.Seed()
	return float64(C.HGE_Random_Float(r.randHGE.HGE, C.float(min), C.float(max)))
}
