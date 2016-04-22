// +build !sdl

// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
// I doubt there will ever bee the need for anything like: +build sdl,openal
// but it's an option
package sound

import (
	"fmt"
	"log"
	"math"
	"runtime"

	"github.com/losinggeneration/openal"
	"github.com/losinggeneration/sndfile"

	"golang.org/x/mobile/exp/audio/al"
)

var (
	logger  *log.Logger
	device  *openal.Device
	context *openal.Context
	silent  = false

	formats = map[uint]map[uint]uint32{
		8: map[uint]uint32{
			1: openal.FORMAT_MONO8,
			2: openal.FORMAT_STEREO8,
		},
		16: map[uint]uint32{
			1: openal.FORMAT_MONO16,
			2: openal.FORMAT_STEREO16,
		},
	}
)

// HGE Handle type
type Effect struct {
	channel Channel
	buffer  al.Buffer
}

type buffer struct {
	data      []byte
	channels  uint
	frequency uint
	bits      uint
}

// HGE Handle type
type Channel struct {
	source al.Source
}

func Initialize(log *log.Logger) error {
	logger = log

	device = openal.OpenDevice(nil)
	if device == nil {
		logger.Println("openal.OpenDevice(nil) failed, using no sound")
		silent = true
		return nil
	}

	context = openal.NewContext(device, openal.FREQUENCY, 44100)
	if context == nil {
		device.Close()
		logger.Println("openal.NewContext() failed, using no sound")
		silent = true
		return nil
	}

	context.MakeCurrent()
	context.Process()

	return al.OpenDevice()
}

func loadFile(filename string) (*buffer, error) {
	s, err := sndfile.Open(filename)
	if err != nil {
		return nil, err
	}

	b, err := s.ReadFrames(s.Info.Frames)
	if err != nil {
		return nil, err
	}

	return &buffer{
		frequency: uint(s.Info.SampleRate),
		channels:  uint(s.Info.Channels),
		bits:      16,
		data:      sndfile.Int16ToByte(b),
	}, nil
}

func NewEffect(filename string) (*Effect, error) {
	e := Effect{}

	buff, err := loadFile(filename)
	if err != nil {
		return nil, err
	}

	f := formats[buff.bits]
	if f == nil {
		return nil, fmt.Errorf("unable to get effect's format")
	}

	selectedFormat := f[buff.channels]
	if selectedFormat == 0 {
		return nil, fmt.Errorf("unabe to get effect's channels")
	}

	s := al.GenSources(1)

	e.channel.source = s[0]

	b := al.GenBuffers(1)
	e.buffer = b[0]

	e.buffer.BufferData(selectedFormat, buff.data, int32(buff.frequency))

	runtime.SetFinalizer(&e, func(e *Effect) {
		e.Free()
	})

	e.channel.setBuffer(e.buffer)

	return &e, nil
}

func (e *Effect) Free() {
	al.DeleteBuffers(e.buffer)
	e.channel.Free()
}

func (e *Effect) Play() Channel {
	return e.PlayEx(100, 0, 1.0, false)
}

// volume = 100, pan = 0, pitch = 1.0, loop = false
func (e *Effect) PlayEx(a ...interface{}) Channel {
	if silent {
		return Channel{}
	}

	volume := 100
	pan := 0
	pitch := 1.0
	loop := false
	ok := false

	if len(a) > 0 {
		volume, ok = a[0].(int)
		if !ok {
			volume = 100
		}
	}
	if len(a) > 1 {
		pan, ok = a[1].(int)
		if !ok {
			pan = 0
		}
	}
	if len(a) > 2 {
		pitch, ok = a[2].(float64)
		if !ok {
			pitch = 1.0
		}
	}
	if len(a) > 3 {
		loop, ok = a[3].(bool)
		if !ok {
			loop = false
		}
	}

	e.channel.Stop()
	e.channel.SetVolume(volume)
	e.channel.SetPitch(pitch)
	e.channel.SetPanning(pan)
	e.channel.setLoop(loop)
	e.channel.play()

	return e.channel
}

func clip(i, min, max int) int {
	return int(math.Max(float64(min), math.Min(float64(i), float64(max))))
}

func (c Channel) setLoop(loop bool) {
	looping := openal.FALSE
	if loop {
		looping = openal.TRUE
	}

	c.source.Seti(openal.LOOPING, int32(looping))
}

func (c Channel) setBuffer(b al.Buffer) {
	if silent {
		return
	}
	c.source.QueueBuffers(b)
}

func (c Channel) play() {
	if silent {
		return
	}
	al.PlaySources(c.source)
}

func (c Channel) Free() {
	al.DeleteSources(c.source)
}

func (c Channel) SetPanning(pan int) {
	if silent {
		return
	}
	c.source.Setfv(openal.POSITION, []float32{float32(clip(pan, -100, 100)) / 100.0, 0.0, 0.0})
}

func (c Channel) SetVolume(volume int) {
	if silent {
		return
	}
	c.source.Setfv(openal.GAIN, []float32{float32(clip(volume, 0, 100)) / 100.0, 0.0, 0.0})
}

func (c Channel) SetPitch(pitch float64) {
	if silent {
		return
	}
	c.source.Setf(openal.PITCH, float32(pitch))
}

func (c Channel) Pause() {
	if silent {
		return
	}
	al.PauseSources(c.source)
}

func (c Channel) Resume() {
	if silent {
		return
	}
	al.PlaySources(c.source)
}

func (c Channel) Stop() {
	if silent {
		return
	}
	al.StopSources(c.source)
}

// Pause all sounds on all channels
func PauseAll() {
	if silent {
		return
	}
	context.Suspend()
}

// Resume all sounds on all channels
func ResumeAll() {
	if silent {
		return
	}
	context.Process()
}

// Stop all sounds on all channels
func StopAll() {
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
