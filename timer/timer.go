package timer

import "time"

type timer struct {
	s      time.Time // the start time
	l      time.Time // the last update time
	f      int       // frames per second
	frames int       // current frame count
}

var t *timer

func Reset() {
	n := time.Now()
	t = &timer{s: n, l: n}
}

// Updates the internal state. Should be called every frame by the main loop
func Update() {
	n := time.Now()

	if t.frames == 0 {
		// Scope the global t for the goroutine so Reset can be called
		// and this will update the old reference. This prevents the goroutine
		// from potentially clobbering the value after a call to Reset has
		// happened.
		update := t
		go func() {
			select {
			case <-time.After(1 * time.Second):
				update.f = update.frames
				update.frames = 0
			}
		}()
		t.frames++
	} else {
		t.frames++
	}

	t.l = n
}

// The time since we've created the timer
func Time() float64 {
	return time.Since(t.s).Seconds()
}

// Time since the last call to Update
func Delta() float64 {
	return time.Since(t.l).Seconds()
}

// The approximate frames per second
func FPS() int {
	return t.f
}
