// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
// I doubt there will ever bee the need for anything like: +build sdl,openal
// but it's an option
package sound

import (
	"log"
	"runtime"

	mix "github.com/veandco/go-sdl2/sdl_mixer"
)

var (
	channels    = mix.DEFAULT_CHANNELS
	logger      *log.Logger
	nextChannel = 0
)

func Initialize(log *log.Logger) error {
	logger = log
	err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, channels, mix.DEFAULT_CHUNKSIZE)
	if err != nil {
		return err
	}

	return mix.Init(mix.INIT_FLAC | mix.INIT_OGG | mix.INIT_MP3)
}

// HGE Handle type
type Effect struct {
	chunk   *mix.Chunk
	channel Channel
}

func NewEffect(filename string, a ...interface{}) (*Effect, error) {
	e := Effect{}
	if chunk, err := mix.LoadWAV(filename); err != nil {
		return nil, err
	} else {
		e.chunk = chunk
	}

	runtime.SetFinalizer(&e, func(e *Effect) {
		e.chunk.Free()
	})

	e.channel.channel = nextChannel % mix.AllocateChannels(-1)

	return &e, nil
}

func (e *Effect) Free() {
	if e.chunk != nil {
		e.chunk.Free()
	}
}

func (e *Effect) Play() Channel {
	return e.PlayEx(100, 0, 1.0, false)
}

// volume = 100, pan = 0, pitch = 1.0, loop = false
func (e *Effect) PlayEx(a ...interface{}) Channel {
	volume := 100
	pan := 0
	pitch := 1.0
	loop := false

	if len(a) > 0 {
		volume = a[0].(int)
	}
	if len(a) > 1 {
		pan = a[1].(int)
	}
	if len(a) > 2 {
		pitch = a[2].(float64)
	}
	if len(a) > 3 {
		loop = a[3].(bool)
	}

	looping := 0
	if loop {
		looping = -1
	}

	e.channel.SetVolume(volume)
	e.channel.SetPanning(pan)
	e.channel.SetPitch(pitch)

	_, err := e.chunk.Play(e.channel.channel, looping)
	if err != nil {
		logger.Println(err)
	}

	return e.channel
}

// HGE Handle type
type Channel struct {
	channel int
}

func (c Channel) SetPanning(pan int) {
	left := uint8(127)
	right := uint8(127)
	if pan < 0 {
		right = uint8(127 - 127*float64(pan*-1)/100)
	}
	if pan > 0 {
		left = uint8(127 - 127*float64(pan)/100)
	}

	mix.SetPanning(c.channel, left, right)
}

func (c Channel) SetVolume(volume int) {
	p := volume / 100
	mix.Volume(c.channel, int(mix.MAX_VOLUME*p))
}

func (c Channel) SetPitch(pitch float64) {
	mix.RegisterEffect(c.channel, func(channel int, stream []byte) {
		for i, v := range stream {
			if i%2 == 0 {
				stream[i] = byte(float64(v) * pitch)
			}
		}
	}, func(int) {})
}

func (c Channel) Pause() {
	mix.Pause(c.channel)
}

func (c Channel) Resume() {
	mix.Resume(c.channel)
}

func (c Channel) Stop() {
	mix.HaltChannel(c.channel)
}

// Pause all sounds on all channels
func PauseAll() {
	mix.Pause(-1)
}

// Resume all sounds on all channels
func ResumeAll() {
	mix.Resume(-1)
}

// Stop all sounds on all channels
func StopAll() {
	mix.HaltChannel(-1)
}

func (c Channel) IsPlaying() bool {
	return false
}

func (c Channel) Len() float64 {
	return 0
}

func (c Channel) Pos() float64 {
	return 0
}

func (c Channel) SetPos(seconds float64) {
}

func (c Channel) SlideTo(time float64, a ...interface{}) {
}

func (c Channel) IsSliding() bool {
	return true
}

// HGE Handle type
type Music struct {
	music interface{}
}

func NewMusic(filename string, size uint32) *Music {
	return nil
}

func (m *Music) Free() {
}

func (m *Music) Play(loop bool, a ...interface{}) Channel {
	return Channel{}
}

func (m *Music) SetAmplification(ampl int) {
}

func (m *Music) Amplification() int {
	return 0
}

func (m *Music) Len() int {
	return 0
}

func (m *Music) SetPos(order, row int) {
}

func (m *Music) Pos() (order, row int, ok bool) {
	return 0, 0, false
}

func (m *Music) SetInstrVolume(instr int, volume int) {
}

func (m *Music) InstrVolume(instr int) int {
	return 0
}

func (m *Music) SetChannelVolume(channel, volume int) {
}

func (m *Music) ChannelVolume(channel int) int {
	return 0
}

// HGE Handle type
type Stream struct {
	stream interface{}
}

func NewStream(filename string, size uint32) *Stream {
	return nil
}

func (s *Stream) Free() {
}

func (s *Stream) Play(loop bool, a ...interface{}) Channel {
	return Channel{}
}
