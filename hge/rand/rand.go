package rand

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
	seed int
}

func New(seed int) *Rand {
	return &Rand{seed}
}

func (r *Rand) Seed() {
}

func (r *Rand) Int(min, max int) int {
	return 0
}

func (r *Rand) Float32(min, max float32) float32 {
	return 0
}

func (r *Rand) Float64(min, max float64) float64 {
	return 0
}
