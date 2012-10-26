package timer

import (
	"testing"
	"time"
)

func TestReset(t *testing.T) {
	tm := time.Now()
	Reset()
	since := time.Since(tm).Seconds()
	if Time() >= since+1e-5 || Delta() >= since+1e-5 || FPS() != 0 {
		t.Errorf("Timer did not reset correctly")
	}

	tm = time.Now()
	Reset()
	select {
	case <-time.After(2 * time.Second):
		since = time.Since(tm).Seconds()
		if Time() >= since+1e-5 || Delta() >= since+1e-5 || FPS() != 0 {
			t.Errorf("Timer did not reset correctly after two seconds")
		}
	}
}
