package timer

import (
	"testing"
	"time"
)

func TestReset(te *testing.T) {
	Reset()

	if !t.s.Equal(t.l) {
		te.Error("Expected start time and last update time to be equal after reset")
	}

	last := t

	time.Sleep(5 * time.Microsecond)

	Reset()
	if last.s.Equal(t.s) {
		te.Error("Did not expect start times to be equal after Reset")
	}

	if last.l.Equal(t.l) {
		te.Error("Did not expect last update times to be equal after Reset")
	}
}

func TestUpdate(te *testing.T) {
	Reset()
	start := t.s
	last := t.l

	time.Sleep(5 * time.Millisecond)
	Update()

	if !start.Equal(t.s) {
		te.Error("Did not expect start time to change after call to Update")
	}

	if last.Equal(t.l) {
		te.Error("Expected last update time to change after call to Update")
	}

	if start.Equal(t.l) {
		te.Error("Did not expect start & last update time to be equal after call to Update")
	}

	if t.frames != 1 {
		te.Error("Did not expect frames to be 0 after call to Update")
	}

	last = t.l

	time.Sleep(5 * time.Millisecond)
	Update()

	if !start.Equal(t.s) {
		te.Error("Did not expect start time to change after call to Update")
	}

	if last.Equal(t.l) {
		te.Error("Expected last update time to change after call to Update")
	}

	if start.Equal(t.l) {
		te.Error("Did not expect start & last update time to be equal after call to Update")
	}

	if t.frames != 2 {
		te.Error("Did not expect frames to be 0 after call to Update")
	}
}

func TestTime(t *testing.T) {
	Reset()

	time.Sleep(5 * time.Millisecond)
	if now := Time(); now <= .001 {
		t.Error("Not eneough time has passed when Time was called")
	}

	time.Sleep(5 * time.Millisecond)
	if now := Time(); now <= .001 {
		t.Error("Not eneough time has passed when Time was called")
	}
}

func TestDelta(t *testing.T) {
	Reset()

	time.Sleep(5 * time.Millisecond)
	if now := Delta(); now <= .001 {
		t.Error("Not eneough time has passed when Delta was called")
	}

	time.Sleep(5 * time.Millisecond)
	if now := Delta(); now <= .001 {
		t.Error("Not eneough time has passed when Delta was called")
	}
}

func TestFPS(te *testing.T) {
	Reset()
	expected := 10

	if f := FPS(); f != 0 {
		te.Errorf("Expected FPS to return 0 but got %v", f)
	}

	for i := 0; i < expected; i++ {
		Update()
	}

	time.Sleep(1015 * time.Millisecond)

	if f := FPS(); f != expected {
		te.Errorf("Expected FPS to return %v but got %v", expected, f)
	}
}
