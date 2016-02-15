// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
// I doubt there will ever bee the need for anything like: +build sdl,openal
// but it's an option
package sound

import "fmt"

func Initialize() error {
	return fmt.Errorf("Sound Initialize not implemented")
}

// HGE Handle type
type Effect struct {
	effect interface{}
}

func NewEffect(filename string, a ...interface{}) *Effect {
	return &Effect{}
}

func (e *Effect) Free() {
}

func (e *Effect) Play() Channel {
	return e.PlayEx(100, 0, 1.0, false)
}

func (e *Effect) PlayEx(a ...interface{}) Channel {
	return Channel{}
}

// HGE Handle type
type Channel struct {
	channel interface{}
}

func (c Channel) SetPanning(pan int) {
}

func (c Channel) SetVolume(volume int) {
}

func (c Channel) SetPitch(pitch float64) {
}

func (c Channel) Pause() {
}

func (c Channel) Resume() {
}

func (c Channel) Stop() {
}

// Pause all sounds on all channels
func PauseAll() {
}

// Resume all sounds on all channels
func ResumeAll() {
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
