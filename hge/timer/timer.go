package timer

import "time"

type timer struct {
	s      time.Time // the start time
	l      time.Time // the last update time
	f      int       // frames per second
	frames int       // current frame count
}

var t timer

func Reset() {
	n := time.Now()
	t = timer{s: n, l: n}
}

func Update() {
	n := time.Now()

	if t.frames == 0 {
		go func() {
			select {
			case <-time.After(1 * time.Second):
				t.f = t.frames
				t.frames = 0
			}
		}()
		t.frames++
	} else {
		t.frames++
	}

	t.l = n
}

func Time() float64 {
	return time.Since(t.s).Seconds()
}

func Delta() float64 {
	return time.Since(t.l).Seconds()
}

func FPS() int {
	return t.f
}
