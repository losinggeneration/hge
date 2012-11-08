// Functions to give  random numbers over exclusive ranges.
package rand

import (
	"math"
	"math/rand"
)

var rnd *Rand

// Seed uses the provided seed value to initialize the generator to a
// deterministic state. If Seed is not called, the generator behaves as if
// seeded by Seed(1).
func Seed(a ...interface{}) {
	seed := int64(1)
	if len(a) == 1 {
		if s, ok := a[0].(int); ok {
			seed = int64(s)
		}
		if s, ok := a[0].(int64); ok {
			seed = s
		}
	}

	rnd := New(seed)
	rnd.Seed()
}

// Returns a pseudo-random number in (min, max). It panics if max-min <= 0.
func Int(min, max int) int {
	return rnd.Int(min, max)
}

// Returns a pseudo-random number in (min, max). It panics if max-min <= 0.
func Float32(min, max float32) float32 {
	return rnd.Float32(min, max)
}

// Returns a pseudo-random number in (min, max). It panics if max-min <= 0.
func Float64(min, max float64) float64 {
	return rnd.Float64(min, max)
}

// A Rand is a source of random numbers.
type Rand struct {
	seed int64
	r    *rand.Rand
}

// New returns a new Rand that uses random values from the seeded number to
// generate other random values.
func New(seed int64) *Rand {
	return &Rand{seed: seed,
		r: rand.New(rand.NewSource(seed)),
	}
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
func (r *Rand) Seed() {
	r.r.Seed(r.seed)
}

// Returns a pseudo-random number in (min, max). It panics if max-min <= 0.
func (r *Rand) Int(min, max int) int {
	return r.r.Intn(max-min) + min
}

// Returns a pseudo-random number in (min, max). It panics if max-min <= 0.
func (r *Rand) Float32(min, max float32) float32 {
	return float32(r.Float64(float64(min), float64(max)))
}

// Returns a pseudo-random number in (min, max). It panics if max-min <= 0.
func (r *Rand) Float64(min, max float64) float64 {
	return math.Mod(r.r.ExpFloat64(), max-min) + min
}
