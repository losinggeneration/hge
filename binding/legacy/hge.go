package legacy

import (
	"github.com/losinggeneration/hge/binding/hge"
	"github.com/losinggeneration/hge/binding/hge/gfx"
	"github.com/losinggeneration/hge/binding/hge/ini"
	"github.com/losinggeneration/hge/binding/hge/input"
	"github.com/losinggeneration/hge/binding/hge/rand"
	"github.com/losinggeneration/hge/binding/hge/resource"
	"github.com/losinggeneration/hge/binding/hge/sound"
	"github.com/losinggeneration/hge/binding/hge/timer"
)

// HGE struct
type HGE struct {
	h *hge.HGE
}

// Creates a new instance of an HGE structure
func Create(ver int) *HGE {
	return &HGE{hge.New()}
}

// Releases the memory the C++ library allocated for the HGE struct
func (h *HGE) Release() {
	h.h.Free()
}

// Initializes hardware and software needed to run engine.
func (h *HGE) System_Initiate() bool {
	return h.h.Initiate() == nil
}

//  Restores video mode and frees allocated resources.
func (h *HGE) System_Shutdown() {
	h.h.Shutdown()
}

// Starts running user defined frame function.
func (h *HGE) System_Start() bool {
	return h.h.Start() == nil
}

//  Returns last occured HGE error description.
func (h *HGE) System_GetErrorMessage() string {
	return h.h.GetErrorMessage()
}

// Writes a formatted message to the log file.
func (h *HGE) System_Log(format string, v ...interface{}) {
	h.h.Log(format, v...)
}

// Launches an URL or external executable/data file.
func (h *HGE) System_Launch(url string) bool {
	return h.h.Launch(url)
}

//  Saves current screen snapshot into a file.
func (h *HGE) System_Snapshot(a ...interface{}) {
	h.h.Snapshot(a...)
}

// Sets internal system states.
// First param should be one of: BoolState, IntState, StringState, FuncState, HwndState
// Second parameter must be of the matching type, bool, int, string, StateFunc/func() int, *Hwnd
func (h *HGE) System_SetState(a ...interface{}) {
	h.h.SetState(a...)
}

// Returns internal system state values.
func (h *HGE) System_GetState(a ...interface{}) interface{} {
	return h.h.GetState(a...)
}

// Loads a resource into memory from disk.
func (h *HGE) Resource_Load(filename string) (*resource.Resource, hge.Dword) {
	return resource.NewResource(filename)
}

// Deletes a previously loaded resource from memory.
func (h *HGE) Resource_Free(res resource.Resource) {
	res.Free()
}

// Loads a resource, puts the loaded data into a byte array, and frees the data.
func (h *HGE) ResourceLoadBytes(filename string) []byte {
	return resource.LoadBytes(filename)
}

// Loads a resource, puts the data into a string, and frees the data.
func (h *HGE) ResourceLoadString(filename string) *string {
	return resource.LoadString(filename)
}

// Attaches a resource pack.
func (h *HGE) Resource_AttachPack(filename string, a ...interface{}) bool {
	return resource.AttachPack(filename, a...)
}

// Removes a resource pack.
func (h *HGE) Resource_RemovePack(filename string) {
	resource.RemovePack(filename)
}

// Removes all resource packs previously attached.
func (h *HGE) Resource_RemoveAllPacks() {
	resource.RemoveAllPacks()
}

// Builds absolute file path.
func (h *HGE) Resource_MakePath(a ...interface{}) string {
	return resource.MakePath(a...)
}

// Enumerates files by given wildcard.
func (h *HGE) Resource_EnumFiles(a ...interface{}) string {
	return resource.EnumFiles(a...)
}

// Enumerates folders by given wildcard.
func (h *HGE) Resource_EnumFolders(a ...interface{}) string {
	return resource.EnumFolders(a...)
}

func (h *HGE) Ini_SetInt(section, name string, value int) {
	ini.NewIni(section, name).SetInt(value)
}

func (h *HGE) Ini_GetInt(section, name string, def_val int) int {
	return ini.NewIni(section, name).GetInt(def_val)
}

func (h *HGE) Ini_SetFloat(section, name string, value float64) {
	ini.NewIni(section, name).SetFloat(value)
}

func (h *HGE) Ini_GetFloat(section, name string, def_val float64) float64 {
	return ini.NewIni(section, name).GetFloat(def_val)
}

func (h *HGE) Ini_SetString(section, name, value string) {
	ini.NewIni(section, name).SetString(value)
}

func (h *HGE) Ini_GetString(section, name, def_val string) string {
	return ini.NewIni(section, name).GetString(def_val)
}

var random = rand.New(0)

func (h *HGE) Random_Seed(a ...interface{}) {
	seed := 1
	if len(a) == 1 {
		if s, ok := a[0].(int); ok {
			seed = s
		}
		if s, ok := a[0].(int64); ok {
			seed = int(s)
		}
	}

	random = rand.New(seed)
	random.Seed()
}

func (h *HGE) Random_Int(min, max int) int {
	return random.Int(min, max)
}

func (h *HGE) Random_Float(min, max float64) float64 {
	return random.Float64(min, max)
}

func (h *HGE) Timer_GetTime() float64 {
	return timer.Time()
}

func (h *HGE) Timer_GetDelta() float64 {
	return timer.Delta()
}

func (h *HGE) Timer_GetFPS() int {
	return timer.GetFPS()
}

func (h *HGE) Effect_Load(filename string, a ...interface{}) *sound.Effect {
	return sound.NewEffect(filename, a...)
}

func (h *HGE) Effect_Free(eff *sound.Effect) {
	eff.Free()
}

func (h *HGE) Effect_Play(eff *sound.Effect) sound.Channel {
	return eff.Play()
}

func (h *HGE) Effect_PlayEx(eff *sound.Effect, a ...interface{}) sound.Channel {
	return eff.PlayEx(a...)
}

func (h *HGE) Music_Load(filename string, size hge.Dword) *sound.Music {
	return sound.NewMusic(filename, size)
}

func (h *HGE) Music_Free(music *sound.Music) {
	music.Free()
}

func (h *HGE) Music_Play(music *sound.Music, loop bool, a ...interface{}) sound.Channel {
	return music.Play(loop, a...)
}

func (h *HGE) Music_SetAmplification(music *sound.Music, ampl int) {
	music.SetAmplification(ampl)
}

func (h *HGE) Music_GetAmplification(music *sound.Music) int {
	return music.Amplification()
}

func (h *HGE) Music_GetLength(music *sound.Music) int {
	return music.Len()
}

func (h *HGE) Music_SetPos(music *sound.Music, order, row int) {
	music.SetPos(order, row)
}

func (h *HGE) Music_GetPos(music *sound.Music) (order, row int, ok bool) {
	return music.Pos()
}

func (h *HGE) Music_SetInstrVolume(music *sound.Music, instr int, volume int) {
	music.SetInstrVolume(instr, volume)
}

func (h *HGE) Music_GetInstrVolume(music *sound.Music, instr int) int {
	return music.InstrVolume(instr)
}

func (h *HGE) Music_SetChannelVolume(music *sound.Music, channel, volume int) {
	music.SetChannelVolume(channel, volume)
}

func (h *HGE) Music_GetChannelVolume(music *sound.Music, channel int) int {
	return music.ChannelVolume(channel)
}

func (h *HGE) Stream_Load(filename string, size hge.Dword) *sound.Stream {
	return sound.NewStream(filename, size)
}

func (h *HGE) Stream_Free(stream *sound.Stream) {
	stream.Free()
}

func (h *HGE) Stream_Play(stream *sound.Stream, loop bool, a ...interface{}) sound.Channel {
	return stream.Play(loop, a...)
}

func (h *HGE) Channel_SetPanning(chn sound.Channel, pan int) {
	chn.SetPanning(pan)
}

func (h *HGE) Channel_SetVolume(chn sound.Channel, volume int) {
	chn.SetVolume(volume)
}

func (h *HGE) Channel_SetPitch(chn sound.Channel, pitch float64) {
	chn.SetPitch(pitch)
}

func (h *HGE) Channel_Pause(chn sound.Channel) {
	chn.Pause()
}

func (h *HGE) Channel_Resume(chn sound.Channel) {
	chn.Resume()
}

func (h *HGE) Channel_Stop(chn sound.Channel) {
	chn.Stop()
}

func (h *HGE) Channel_PauseAll() {
	sound.PauseAll()
}

func (h *HGE) Channel_ResumeAll() {
	sound.ResumeAll()
}

func (h *HGE) Channel_StopAll() {
	sound.StopAll()
}

func (h *HGE) Channel_IsPlaying(chn sound.Channel) bool {
	return chn.IsPlaying()
}

func (h *HGE) Channel_GetLength(chn sound.Channel) float64 {
	return chn.Len()
}

func (h *HGE) Channel_GetPos(chn sound.Channel) float64 {
	return chn.Pos()
}

func (h *HGE) Channel_SetPos(chn sound.Channel, seconds float64) {
	chn.SetPos(seconds)
}

func (h *HGE) Channel_SlideTo(chn sound.Channel, time float64, a ...interface{}) {
	chn.SlideTo(time, a...)
}

func (h *HGE) Channel_IsSliding(chn sound.Channel) bool {
	return chn.IsSliding()
}

func (h *HGE) Input_GetMousePos() (x, y float64) {
	return input.NewMouse(0, 0).Pos()
}

func (h *HGE) Input_SetMousePos(x, y float64) {
	input.Mouse{}.SetPos(x, y)
}

func (h *HGE) Input_GetMouseWheel() int {
	return input.NewMouse(0, 0).WheelMovement()
}

func (h *HGE) Input_IsMouseOver() bool {
	return input.NewMouse(0, 0).IsOver()
}

func (h *HGE) Input_KeyDown(key int) bool {
	return input.NewKey(key).Down()
}

func (h *HGE) Input_KeyUp(key int) bool {
	return input.NewKey(key).Up()
}

func (h *HGE) Input_GetKeyState(key int) bool {
	return input.NewKey(key).State()
}

func (h *HGE) Input_GetKeyName(key int) string {
	return input.NewKey(key).Name()
}

func (h *HGE) Input_GetKey() int {
	return int(input.GetKey())
}

func (h *HGE) Input_GetChar() int {
	return input.GetChar()
}

func (h *HGE) Input_GetEvent(event *input.InputEvent) bool {
	event, b := input.GetEvent()
	return b
}

func (h *HGE) Gfx_BeginScene(a ...interface{}) bool {
	return gfx.BeginScene(a...)
}

func (h *HGE) Gfx_EndScene() {
	gfx.EndScene()
}

func (h *HGE) Gfx_Clear(color hge.Dword) {
	gfx.Clear(color)
}

func (h *HGE) Gfx_RenderLine(x1, y1, x2, y2 float64, a ...interface{}) {
	gfx.NewLine(x1, y1, x2, y2, a...).Render()
}

func (h *HGE) Gfx_RenderTriple(triple *gfx.Triple) {
	triple.Render()
}

func (h *HGE) Gfx_RenderQuad(quad *gfx.Quad) {
	quad.Render()
}

func (h *HGE) Gfx_StartBatch(prim_type int, tex *gfx.Texture, blend int) (ver *gfx.Vertex, max_prim int, ok bool) {
	return gfx.StartBatch(prim_type, tex, blend)
}

func (h *HGE) Gfx_FinishBatch(prim int) {
	gfx.FinishBatch(prim)
}

func (h *HGE) Gfx_SetClipping(a ...interface{}) {
	gfx.SetClipping(a...)
}

func (h *HGE) Gfx_SetTransform(a ...interface{}) {
	gfx.SetTransform(a...)
}

func (h *HGE) Target_Create(width, height int, zbuffer bool) *gfx.Target {
	return gfx.NewTarget(width, height, zbuffer)
}

func (h *HGE) Target_Free(target *gfx.Target) {
	target.Free()
}

func (h *HGE) Target_GetTexture(target gfx.Target) *gfx.Texture {
	return target.Texture()
}

func (h *HGE) Texture_Create(width, height int) *gfx.Texture {
	return gfx.NewTexture(width, height)
}

func (h *HGE) Texture_Load(filename string, a ...interface{}) *gfx.Texture {
	return gfx.LoadTexture(filename, a...)
}

func (h *HGE) Texture_Free(tex *gfx.Texture) {
	tex.Free()
}

func (h *HGE) Texture_GetWidth(tex gfx.Texture, a ...interface{}) int {
	return tex.Width(a...)
}

func (h *HGE) Texture_GetHeight(tex gfx.Texture, a ...interface{}) int {
	return tex.Height(a...)
}

func (h *HGE) Texture_Lock(tex gfx.Texture, a ...interface{}) *hge.Dword {
	return tex.Lock(a...)
}

func (h *HGE) Texture_Unlock(tex gfx.Texture) {
	tex.Unlock()
}
