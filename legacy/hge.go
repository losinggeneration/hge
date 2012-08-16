package legacy

import (
	hge "github.com/losinggeneration/hge-go/hge"
)

// HGE struct
type HGE struct{}

// Creates a new instance of an HGE structure
func Create(ver int) *HGE {
	h := new(HGE)

	return h
}

// Releases the memory the C++ library allocated for the HGE struct
func (h *HGE) Release() {
	hge.Free()
}

// Initializes hardware and software needed to run engine.
func (h *HGE) System_Initiate() bool {
	return hge.Initiate() == nil
}

//  Restores video mode and frees allocated resources.
func (h *HGE) System_Shutdown() {
	hge.Shutdown()
}

// Starts running user defined frame function.
func (h *HGE) System_Start() bool {
	return hge.Start() == nil
}

//  Returns last occured HGE error description.
func (h *HGE) System_GetErrorMessage() string {
	return hge.GetErrorMessage()
}

// Writes a formatted message to the log file.
func (h *HGE) System_Log(format string, v ...interface{}) {
	hge.Log(format, v...)
}

// Launches an URL or external executable/data file.
func (h *HGE) System_Launch(url string) bool {
	return hge.Launch(url)
}

//  Saves current screen snapshot into a file.
func (h *HGE) System_Snapshot(a ...interface{}) {
	hge.Snapshot(a...)
}

// Sets internal system states.
// First param should be one of: BoolState, IntState, StringState, FuncState, HwndState
// Second parameter must be of the matching type, bool, int, string, StateFunc/func() int, *Hwnd
func (h *HGE) System_SetState(a ...interface{}) {
	hge.SetState(a...)
}

// Returns internal system state values.
func (h *HGE) System_GetState(a ...interface{}) interface{} {
	return hge.GetState(a...)
}

// Loads a resource into memory from disk.
func (h *HGE) Resource_Load(filename string) (*hge.Resource, hge.Dword) {
	return hge.NewResource(filename)
}

// Deletes a previously loaded resource from memory.
func (h *HGE) Resource_Free(res hge.Resource) {
	res.Free()
}

// Loads a resource, puts the loaded data into a byte array, and frees the data.
func (h *HGE) ResourceLoadBytes(filename string) []byte {
	return hge.LoadBytes(filename)
}

// Loads a resource, puts the data into a string, and frees the data.
func (h *HGE) ResourceLoadString(filename string) *string {
	return hge.LoadString(filename)
}

// Attaches a resource pack.
func (h *HGE) Resource_AttachPack(filename string, a ...interface{}) bool {
	return hge.Resource(0).AttachPack(filename, a...)
}

// Removes a resource pack.
func (h *HGE) Resource_RemovePack(filename string) {
	hge.Resource(0).RemovePack(filename)
}

// Removes all resource packs previously attached.
func (h *HGE) Resource_RemoveAllPacks() {
	hge.Resource(0).RemoveAllPacks()
}

// Builds absolute file path.
func (h *HGE) Resource_MakePath(a ...interface{}) string {
	return hge.Resource(0).MakePath(a...)
}

// Enumerates files by given wildcard.
func (h *HGE) Resource_EnumFiles(a ...interface{}) string {
	return hge.Resource(0).EnumFiles(a...)
}

// Enumerates folders by given wildcard.
func (h *HGE) Resource_EnumFolders(a ...interface{}) string {
	return hge.Resource(0).EnumFolders(a...)
}

func (h *HGE) Ini_SetInt(section, name string, value int) {
	hge.NewIni(section, name).SetInt(value)
}

func (h *HGE) Ini_GetInt(section, name string, def_val int) int {
	return hge.NewIni(section, name).GetInt(def_val)
}

func (h *HGE) Ini_SetFloat(section, name string, value float64) {
	hge.NewIni(section, name).SetFloat(value)
}

func (h *HGE) Ini_GetFloat(section, name string, def_val float64) float64 {
	return hge.NewIni(section, name).GetFloat(def_val)
}

func (h *HGE) Ini_SetString(section, name, value string) {
	hge.NewIni(section, name).SetString(value)
}

func (h *HGE) Ini_GetString(section, name, def_val string) string {
	return hge.NewIni(section, name).GetString(def_val)
}

func (h *HGE) Random_Seed(a ...interface{}) {
	hge.RandomSeed(a...)
}

func (h *HGE) Random_Int(min, max int) int {
	return hge.RandomInt(min, max)
}

func (h *HGE) Random_Float(min, max float64) float64 {
	return hge.RandomFloat(min, max)
}

func (h *HGE) Timer_GetTime() float64 {
	return hge.NewTimer().Time()
}

func (h *HGE) Timer_GetDelta() float64 {
	return hge.NewTimer().Delta()
}

func (h *HGE) Timer_GetFPS() int {
	return hge.GetFPS()
}

func (h *HGE) Effect_Load(filename string, a ...interface{}) hge.Effect {
	return hge.NewEffect(filename, a...)
}

func (h *HGE) Effect_Free(eff hge.Effect) {
	eff.Free()
}

func (h *HGE) Effect_Play(eff hge.Effect) hge.Channel {
	return eff.Play()
}

func (h *HGE) Effect_PlayEx(eff hge.Effect, a ...interface{}) hge.Channel {
	return eff.PlayEx(a...)
}

func (h *HGE) Music_Load(filename string, size hge.Dword) hge.Music {
	return hge.NewMusic(filename, size)
}

func (h *HGE) Music_Free(music hge.Music) {
	music.Free()
}

func (h *HGE) Music_Play(music hge.Music, loop bool, a ...interface{}) hge.Channel {
	return music.Play(loop, a...)
}

func (h *HGE) Music_SetAmplification(music hge.Music, ampl int) {
	music.SetAmplification(ampl)
}

func (h *HGE) Music_GetAmplification(music hge.Music) int {
	return music.Amplification()
}

func (h *HGE) Music_GetLength(music hge.Music) int {
	return music.Len()
}

func (h *HGE) Music_SetPos(music hge.Music, order, row int) {
	music.SetPos(order, row)
}

func (h *HGE) Music_GetPos(music hge.Music) (order, row int, ok bool) {
	return music.Pos()
}

func (h *HGE) Music_SetInstrVolume(music hge.Music, instr int, volume int) {
	music.SetInstrVolume(instr, volume)
}

func (h *HGE) Music_GetInstrVolume(music hge.Music, instr int) int {
	return music.InstrVolume(instr)
}

func (h *HGE) Music_SetChannelVolume(music hge.Music, channel, volume int) {
	music.SetChannelVolume(channel, volume)
}

func (h *HGE) Music_GetChannelVolume(music hge.Music, channel int) int {
	return music.ChannelVolume(channel)
}

func (h *HGE) Stream_Load(filename string, size hge.Dword) hge.Stream {
	return hge.NewStream(filename, size)
}

func (h *HGE) Stream_Free(stream hge.Stream) {
	stream.Free()
}

func (h *HGE) Stream_Play(stream hge.Stream, loop bool, a ...interface{}) hge.Channel {
	return stream.Play(loop, a...)
}

func (h *HGE) Channel_SetPanning(chn hge.Channel, pan int) {
	chn.SetPanning(pan)
}

func (h *HGE) Channel_SetVolume(chn hge.Channel, volume int) {
	chn.SetVolume(volume)
}

func (h *HGE) Channel_SetPitch(chn hge.Channel, pitch float64) {
	chn.SetPitch(pitch)
}

func (h *HGE) Channel_Pause(chn hge.Channel) {
	chn.Pause()
}

func (h *HGE) Channel_Resume(chn hge.Channel) {
	chn.Resume()
}

func (h *HGE) Channel_Stop(chn hge.Channel) {
	chn.Stop()
}

func (h *HGE) Channel_PauseAll() {
	hge.Channel(0).PauseAll()
}

func (h *HGE) Channel_ResumeAll() {
	hge.Channel(0).ResumeAll()
}

func (h *HGE) Channel_StopAll() {
	hge.Channel(0).StopAll()
}

func (h *HGE) Channel_IsPlaying(chn hge.Channel) bool {
	return chn.IsPlaying()
}

func (h *HGE) Channel_GetLength(chn hge.Channel) float64 {
	return chn.Len()
}

func (h *HGE) Channel_GetPos(chn hge.Channel) float64 {
	return chn.Pos()
}

func (h *HGE) Channel_SetPos(chn hge.Channel, seconds float64) {
	chn.SetPos(seconds)
}

func (h *HGE) Channel_SlideTo(chn hge.Channel, time float64, a ...interface{}) {
	chn.SlideTo(time, a...)
}

func (h *HGE) Channel_IsSliding(chn hge.Channel) bool {
	return chn.IsSliding()
}

func (h *HGE) Input_GetMousePos() (x, y float64) {
	return hge.MousePos()
}

func (h *HGE) Input_SetMousePos(x, y float64) {
	hge.SetMousePos(x, y)
}

func (h *HGE) Input_GetMouseWheel() int {
	return hge.MouseWheel()
}

func (h *HGE) Input_IsMouseOver() bool {
	return hge.IsMouseOver()
}

func (h *HGE) Input_KeyDown(key int) bool {
	return hge.NewKey(key).Down()
}

func (h *HGE) Input_KeyUp(key int) bool {
	return hge.NewKey(key).Up()
}

func (h *HGE) Input_GetKeyState(key int) bool {
	return hge.NewKey(key).State()
}

func (h *HGE) Input_GetKeyName(key int) string {
	return hge.NewKey(key).Name()
}

func (h *HGE) Input_GetKey() int {
	return int(hge.GetKey())
}

func (h *HGE) Input_GetChar() int {
	return hge.GetChar()
}

func (h *HGE) Input_GetEvent(event *hge.InputEvent) bool {
	event, b := hge.GetEvent()
	return b
}

func (h *HGE) Gfx_BeginScene(a ...interface{}) bool {
	return hge.GfxBeginScene(a)
}

func (h *HGE) Gfx_EndScene() {
	hge.GfxEndScene()
}

func (h *HGE) Gfx_Clear(color hge.Dword) {
	hge.GfxClear(color)
}

func (h *HGE) Gfx_RenderLine(x1, y1, x2, y2 float64, a ...interface{}) {
	hge.GfxRenderLine(x1, y1, x2, y2, a...)
}

func (h *HGE) Gfx_RenderTriple(triple *hge.Triple) {
	triple.Render()
}

func (h *HGE) Gfx_RenderQuad(quad *hge.Quad) {
	quad.Render()
}

func (h *HGE) Gfx_StartBatch(prim_type int, tex hge.Texture, blend int) (ver *hge.Vertex, max_prim int, ok bool) {
	return hge.GfxStartBatch(prim_type, tex, blend)
}

func (h *HGE) Gfx_FinishBatch(prim int) {
	hge.GfxFinishBatch(prim)
}

func (h *HGE) Gfx_SetClipping(a ...interface{}) {
	hge.GfxSetClipping(a...)
}

func (h *HGE) Gfx_SetTransform(a ...interface{}) {
	hge.GfxSetTransform(a...)
}

func (h *HGE) Target_Create(width, height int, zbuffer bool) hge.Target {
	return hge.NewTarget(width, height, zbuffer)
}

func (h *HGE) Target_Free(target hge.Target) {
	target.Free()
}

func (h *HGE) Target_GetTexture(target hge.Target) hge.Texture {
	return target.Texture()
}

func (h *HGE) Texture_Create(width, height int) hge.Texture {
	return hge.NewTexture(width, height)
}

func (h *HGE) Texture_Load(filename string, a ...interface{}) hge.Texture {
	return hge.LoadTexture(filename, a...)
}

func (h *HGE) Texture_Free(tex hge.Texture) {
	tex.Free()
}

func (h *HGE) Texture_GetWidth(tex hge.Texture, a ...interface{}) int {
	return tex.Width(a...)
}

func (h *HGE) Texture_GetHeight(tex hge.Texture, a ...interface{}) int {
	return tex.Height(a...)
}

func (h *HGE) Texture_Lock(tex hge.Texture, a ...interface{}) *hge.Dword {
	return tex.Lock(a...)
}

func (h *HGE) Texture_Unlock(tex hge.Texture) {
	tex.Unlock()
}
