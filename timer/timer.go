package timer

import "time"

type Timer struct {
	start  time.Time // the start time
	now    time.Time // the time now
	last   time.Time // the last update time
	fps    int       // frames per second
	frames int       // current frame count
}

var t = NewTimer()

// NewTimer creates a new Timer
func NewTimer() *Timer {
	t := Timer{}
	t.Reset()

	return &t
}

func (t *Timer) Reset() {
	n := time.Now()
	*t = Timer{start: n, now: n, last: n}
}

// Updates the internal state. Should be called every frame by the main loop
// You should not call this from more than one go routine at a time.
func (t *Timer) Update() {
	t.last = t.now
	t.now = time.Now()

	if t.frames == 0 {
		// Scope the global t for the goroutine so Reset can be called
		// and this will update the old reference. This prevents the goroutine
		// from potentially clobbering the value after a call to Reset has
		// happened.
		update := t
		go func() {
			select {
			case <-time.After(1 * time.Second):
				update.fps = update.frames
				update.frames = 0
			}
		}()
		t.frames++
	} else {
		t.frames++
	}
}

// The time since we've created the timer
func (t Timer) Time() float64 {
	return time.Since(t.start).Seconds()
}

// Time since the last call to Update
func (t Timer) Delta() float64 {
	return t.now.Sub(t.last).Seconds()
}

// The approximate frames per second
func (t Timer) FPS() int {
	return t.fps
}

func Reset() {
	t.Reset()
}

func Update() {
	t.Update()
}

func Time() float64 {
	return t.Time()
}

func Delta() float64 {
	return t.Delta()
}

func FPS() int {
	return t.FPS()
}
