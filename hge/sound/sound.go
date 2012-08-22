package sound

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import (
	"fmt"
	"github.com/losinggeneration/hge-go/hge"
	"runtime"
	"unsafe"
)

func boolToCInt(b bool) C.BOOL {
	return C.BOOL(hge.BoolToCInt(b))
}

// HGE Handle type
type Effect struct {
	effect   C.HEFFECT
	soundHGE *hge.HGE
}

func NewEffect(filename string, a ...interface{}) *Effect {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	size := hge.Dword(0)

	if len(a) == 1 {
		if s, ok := a[0].(hge.Dword); ok {
			// 			return Effect(C.HGE_Effect_Load(soundHGE.HGE, fname, C.DWORD(size)))
			size = s
		}
	}

	e := new(Effect)
	e.soundHGE = hge.New()
	e.effect = C.HGE_Effect_Load(e.soundHGE.HGE, fname, C.DWORD(size))

	runtime.SetFinalizer(e, func(effect *Effect) {
		effect.free()
	})

	return e
}

func (e *Effect) free() {
	fmt.Println("Effect.Free")
	C.HGE_Effect_Free(e.soundHGE.HGE, (e.effect))
}

func (e *Effect) Play() Channel {
	return Channel{C.HGE_Effect_Play(e.soundHGE.HGE, e.effect), e.soundHGE}
}

func (e *Effect) PlayEx(a ...interface{}) Channel {
	volume, pan := 100, 0
	pitch := 1.0
	loop := false

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if v, ok := a[i].(int); ok {
				volume = v
			}
		}
		if i == 1 {
			if p, ok := a[i].(int); ok {
				pan = p
			}
		}
		if i == 2 {
			if p, ok := a[i].(float32); ok {
				pitch = float64(p)
			}
			if p, ok := a[i].(float64); ok {
				pitch = p
			}
		}
		if i == 3 {
			if l, ok := a[i].(bool); ok {
				loop = l
			}
		}
	}

	return Channel{C.HGE_Effect_PlayEx(e.soundHGE.HGE, e.effect, C.int(volume), C.int(pan), C.float(pitch), boolToCInt(loop)), e.soundHGE}
}

// HGE Handle type
type Channel struct {
	channel  C.HCHANNEL
	soundHGE *hge.HGE
}

func (c Channel) SetPanning(pan int) {
	C.HGE_Channel_SetPanning(c.soundHGE.HGE, c.channel, C.int(pan))
}

func (c Channel) SetVolume(volume int) {
	C.HGE_Channel_SetVolume(c.soundHGE.HGE, c.channel, C.int(volume))
}

func (c Channel) SetPitch(pitch float64) {
	C.HGE_Channel_SetPitch(c.soundHGE.HGE, c.channel, C.float(pitch))
}

func (c Channel) Pause() {
	C.HGE_Channel_Pause(c.soundHGE.HGE, c.channel)
}

func (c Channel) Resume() {
	C.HGE_Channel_Resume(c.soundHGE.HGE, c.channel)
}

func (c Channel) Stop() {
	C.HGE_Channel_Stop(c.soundHGE.HGE, c.channel)
}

// Pause all sounds on all channels
func PauseAll() {
	C.HGE_Channel_PauseAll(hge.New().HGE)
}

// Resume all sounds on all channels
func ResumeAll() {
	C.HGE_Channel_ResumeAll(hge.New().HGE)
}

// Stop all sounds on all channels
func StopAll() {
	C.HGE_Channel_StopAll(hge.New().HGE)
}

func (c Channel) IsPlaying() bool {
	return C.HGE_Channel_IsPlaying(c.soundHGE.HGE, c.channel) == 1
}

func (c Channel) Len() float64 {
	return float64(C.HGE_Channel_GetLength(c.soundHGE.HGE, c.channel))
}

func (c Channel) Pos() float64 {
	return float64(C.HGE_Channel_GetPos(c.soundHGE.HGE, c.channel))
}

func (c Channel) SetPos(seconds float64) {
	C.HGE_Channel_SetPos(c.soundHGE.HGE, c.channel, C.float(seconds))
}

func (c Channel) SlideTo(time float64, a ...interface{}) {
	volume, pan := 100, 0
	pitch := 1.0

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if v, ok := a[i].(int); ok {
				volume = v
			}
		}
		if i == 1 {
			if p, ok := a[i].(int); ok {
				pan = p
			}
		}
		if i == 2 {
			if p, ok := a[i].(float32); ok {
				pitch = float64(p)
			}
			if p, ok := a[i].(float64); ok {
				pitch = p
			}
		}
	}

	C.HGE_Channel_SlideTo(c.soundHGE.HGE, c.channel, C.float(time), C.int(volume), C.int(pan), C.float(pitch))
}

func (c Channel) IsSliding() bool {
	return C.HGE_Channel_IsSliding(c.soundHGE.HGE, c.channel) == 1
}

// HGE Handle type
type Music struct {
	music    C.HMUSIC
	soundHGE *hge.HGE
}

func NewMusic(filename string, size hge.Dword) *Music {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	m := new(Music)
	m.soundHGE = hge.New()
	m.music = C.HGE_Music_Load(m.soundHGE.HGE, fname, C.DWORD(size))

	runtime.SetFinalizer(m, func(music *Music) {
		music.free()
	})

	return m
}

func (m *Music) free() {
	fmt.Println("Music.Free")
	C.HGE_Music_Free(m.soundHGE.HGE, m.music)
}

func (m *Music) Play(loop bool, a ...interface{}) Channel {
	volume, order, row := 100, -1, -1

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if v, ok := a[i].(int); ok {
				volume = v
			}
		}
		if i == 1 {
			if o, ok := a[i].(int); ok {
				order = o
			}
		}
		if i == 2 {
			if r, ok := a[i].(int); ok {
				row = r
			}
		}
	}

	return Channel{C.HGE_Music_Play(m.soundHGE.HGE, m.music, boolToCInt(loop), C.int(volume), C.int(order), C.int(row)), m.soundHGE}
}

func (m *Music) SetAmplification(ampl int) {
	C.HGE_Music_SetAmplification(m.soundHGE.HGE, m.music, C.int(ampl))
}

func (m *Music) Amplification() int {
	return int(C.HGE_Music_GetAmplification(m.soundHGE.HGE, m.music))
}

func (m *Music) Len() int {
	return int(C.HGE_Music_GetLength(m.soundHGE.HGE, m.music))
}

func (m *Music) SetPos(order, row int) {
	C.HGE_Music_SetPos(m.soundHGE.HGE, m.music, C.int(order), C.int(row))
}

func (m *Music) Pos() (order, row int, ok bool) {
	var o, r C.int

	ok = C.HGE_Music_GetPos(m.soundHGE.HGE, m.music, &o, &r) == 1

	return int(o), int(r), ok
}

func (m *Music) SetInstrVolume(instr int, volume int) {
	C.HGE_Music_SetInstrVolume(m.soundHGE.HGE, m.music, C.int(instr), C.int(volume))
}

func (m *Music) InstrVolume(instr int) int {
	return int(C.HGE_Music_GetInstrVolume(m.soundHGE.HGE, m.music, C.int(instr)))
}

func (m *Music) SetChannelVolume(channel Channel, volume int) {
	C.HGE_Music_SetChannelVolume(m.soundHGE.HGE, m.music, C.int(channel.channel), C.int(volume))
}

func (m *Music) ChannelVolume(channel Channel) int {
	return int(C.HGE_Music_GetChannelVolume(m.soundHGE.HGE, m.music, C.int(channel.channel)))
}

// HGE Handle type
type Stream struct {
	stream   C.HSTREAM
	soundHGE *hge.HGE
}

func NewStream(filename string, size hge.Dword) *Stream {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	s := new(Stream)
	s.soundHGE = hge.New()
	s.stream = C.HGE_Stream_Load(s.soundHGE.HGE, fname, C.DWORD(size))

	runtime.SetFinalizer(s, func(stream *Stream) {
		stream.free()
	})

	return s
}

func (s *Stream) free() {
	fmt.Println("Stream.free")
	C.HGE_Stream_Free(s.soundHGE.HGE, s.stream)
}

func (s *Stream) Play(loop bool, a ...interface{}) Channel {
	volume := 100

	if len(a) == 1 {
		if v, ok := a[0].(int); ok {
			volume = v
		}
	}

	return Channel{C.HGE_Stream_Play(s.soundHGE.HGE, s.stream, boolToCInt(loop), C.int(volume)), s.soundHGE}
}
