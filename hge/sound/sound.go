package sound

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import (
	"github.com/losinggeneration/hge-go/hge"
	"unsafe"
)

var soundHGE *hge.HGE

func init() {
	soundHGE = hge.New()
}

func boolToCInt(b bool) C.BOOL {
	return C.BOOL(hge.BoolToCInt(b))
}

// HGE Handle type
type Effect C.HEFFECT

func NewEffect(filename string, a ...interface{}) Effect {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	if len(a) == 1 {
		if size, ok := a[0].(hge.Dword); ok {
			return Effect(C.HGE_Effect_Load(soundHGE.HGE, fname, C.DWORD(size)))
		}
	}

	return Effect(C.HGE_Effect_Load(soundHGE.HGE, fname, 0))
}

func (e Effect) Free() {
	C.HGE_Effect_Free(soundHGE.HGE, (C.HEFFECT(e)))
}

func (e Effect) Play() Channel {
	return Channel(C.HGE_Effect_Play(soundHGE.HGE, C.HEFFECT(e)))
}

func (e Effect) PlayEx(a ...interface{}) Channel {
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

	return Channel(C.HGE_Effect_PlayEx(soundHGE.HGE, C.HEFFECT(e), C.int(volume), C.int(pan), C.float(pitch), boolToCInt(loop)))
}

// HGE Handle type
type Channel C.HCHANNEL

func (c Channel) SetPanning(pan int) {
	C.HGE_Channel_SetPanning(soundHGE.HGE, C.HCHANNEL(c), C.int(pan))
}

func (c Channel) SetVolume(volume int) {
	C.HGE_Channel_SetVolume(soundHGE.HGE, C.HCHANNEL(c), C.int(volume))
}

func (c Channel) SetPitch(pitch float64) {
	C.HGE_Channel_SetPitch(soundHGE.HGE, C.HCHANNEL(c), C.float(pitch))
}

func (c Channel) Pause() {
	C.HGE_Channel_Pause(soundHGE.HGE, C.HCHANNEL(c))
}

func (c Channel) Resume() {
	C.HGE_Channel_Resume(soundHGE.HGE, C.HCHANNEL(c))
}

func (c Channel) Stop() {
	C.HGE_Channel_Stop(soundHGE.HGE, C.HCHANNEL(c))
}

// Pause all sounds on all channels
func PauseAll() {
	C.HGE_Channel_PauseAll(soundHGE.HGE)
}

// Resume all sounds on all channels
func ResumeAll() {
	C.HGE_Channel_ResumeAll(soundHGE.HGE)
}

// Stop all sounds on all channels
func StopAll() {
	C.HGE_Channel_StopAll(soundHGE.HGE)
}

func (c Channel) IsPlaying() bool {
	return C.HGE_Channel_IsPlaying(soundHGE.HGE, C.HCHANNEL(c)) == 1
}

func (c Channel) Len() float64 {
	return float64(C.HGE_Channel_GetLength(soundHGE.HGE, C.HCHANNEL(c)))
}

func (c Channel) Pos() float64 {
	return float64(C.HGE_Channel_GetPos(soundHGE.HGE, C.HCHANNEL(c)))
}

func (c Channel) SetPos(seconds float64) {
	C.HGE_Channel_SetPos(soundHGE.HGE, C.HCHANNEL(c), C.float(seconds))
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

	C.HGE_Channel_SlideTo(soundHGE.HGE, C.HCHANNEL(c), C.float(time), C.int(volume), C.int(pan), C.float(pitch))
}

func (c Channel) IsSliding() bool {
	return C.HGE_Channel_IsSliding(soundHGE.HGE, C.HCHANNEL(c)) == 1
}

// HGE Handle type
type Music C.HMUSIC

func NewMusic(filename string, size hge.Dword) Music {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	return Music(C.HGE_Music_Load(soundHGE.HGE, fname, C.DWORD(size)))
}

func (m Music) Free() {
	C.HGE_Music_Free(soundHGE.HGE, C.HMUSIC(m))
}

func (m Music) Play(loop bool, a ...interface{}) Channel {
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

	return Channel(C.HGE_Music_Play(soundHGE.HGE, C.HMUSIC(m), boolToCInt(loop), C.int(volume), C.int(order), C.int(row)))
}

func (m Music) SetAmplification(ampl int) {
	C.HGE_Music_SetAmplification(soundHGE.HGE, C.HMUSIC(m), C.int(ampl))
}

func (m Music) Amplification() int {
	return int(C.HGE_Music_GetAmplification(soundHGE.HGE, C.HMUSIC(m)))
}

func (m Music) Len() int {
	return int(C.HGE_Music_GetLength(soundHGE.HGE, C.HMUSIC(m)))
}

func (m Music) SetPos(order, row int) {
	C.HGE_Music_SetPos(soundHGE.HGE, C.HMUSIC(m), C.int(order), C.int(row))
}

func (m Music) Pos() (order, row int, ok bool) {
	var o, r C.int

	ok = C.HGE_Music_GetPos(soundHGE.HGE, C.HMUSIC(m), &o, &r) == 1

	return int(o), int(r), ok
}

func (m Music) SetInstrVolume(instr int, volume int) {
	C.HGE_Music_SetInstrVolume(soundHGE.HGE, C.HMUSIC(m), C.int(instr), C.int(volume))
}

func (m Music) InstrVolume(instr int) int {
	return int(C.HGE_Music_GetInstrVolume(soundHGE.HGE, C.HMUSIC(m), C.int(instr)))
}

func (m Music) SetChannelVolume(channel Channel, volume int) {
	C.HGE_Music_SetChannelVolume(soundHGE.HGE, C.HMUSIC(m), C.int(channel), C.int(volume))
}

func (m Music) ChannelVolume(channel Channel) int {
	return int(C.HGE_Music_GetChannelVolume(soundHGE.HGE, C.HMUSIC(m), C.int(channel)))
}

// HGE Handle type
type Stream C.HSTREAM

func NewStream(filename string, size hge.Dword) Stream {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	return Stream(C.HGE_Stream_Load(soundHGE.HGE, fname, C.DWORD(size)))
}

func (s Stream) Free() {
	C.HGE_Stream_Free(soundHGE.HGE, C.HSTREAM(s))
}

func (s Stream) Play(loop bool, a ...interface{}) Channel {
	volume := 100

	if len(a) == 1 {
		if v, ok := a[0].(int); ok {
			volume = v
		}
	}

	return Channel(C.HGE_Stream_Play(soundHGE.HGE, C.HSTREAM(s), boolToCInt(loop), C.int(volume)))
}
