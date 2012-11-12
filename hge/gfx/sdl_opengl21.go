// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
// I doubt there will ever bee the need for anything like: +build sdl,opengl
// or: +build sdl,software
// but it's an option

// +build opengl,2_1

package gfx

import "fmt"
import (
	"github.com/banthar/Go-SDL/sdl"
	gl "github.com/chsc/gogl/gl21"
	"runtime"
)

var (
	width, height gl.Sizei
)

func SetWidth(w int) {
	width = gl.Sizei(w)
}

func SetHeight(h int) {
	height = gl.Sizei(h)
}

func Initialize() error {
	err := gl.Init()
	if err != nil {
		return err
	}

	gl.Viewport(0, 0, width, height)

	return fmt.Errorf("Gfx initialize not implemented")
}
func BeginScene(a ...interface{}) bool {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	return false
}

func EndScene() {
	sdl.GL_SwapBuffers()
}

func Clear(color uint32) {
}

func NewLine(x1, y1, x2, y2 float64, a ...interface{}) Line {
	color := uint32(0xFFFFFFFF)
	z := 0.5

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if c, ok := a[i].(uint); ok {
				color = uint32(c)
			}
			if c, ok := a[i].(uint32); ok {
				color = c
			}
		}
		if i == 1 {
			if z1, ok := a[i].(float32); ok {
				z = float64(z1)
			}
			if z1, ok := a[i].(float64); ok {
				z = z1
			}
		}
	}

	return Line{x1, y1, x2, y2, z, color}
}

func (l Line) Render() {
}

func (t *Triple) Render() {
}

func (q *Quad) Render() {
}

func StartBatch(prim_type int, tex *Texture, blend int) (ver *Vertex, max_prim int, ok bool) {
	return nil, 0, false
}

func FinishBatch(prim int) {
}

func SetClipping(a ...interface{}) {
}

func SetTransform(a ...interface{}) {
}

// HGE Handle type
type Target struct {
	target interface{}
}

func NewTarget(width, height int, zbuffer bool) *Target {
	return nil
}

func (t *Target) Free() {
}

func (t *Target) Texture() *Texture {
	return nil
}

// HGE Handle type
type Texture struct {
	tex gl.Uint
}

func NewTexture(width, height int) *Texture {
	var t gl.Uint
	gl.GenTextures(1, &t)
	tex := Texture{tex: t}

	runtime.SetFinalizer(&tex, func(texture *Texture) {
		texture.free()
	})

	return &tex
}

func LoadTexture(filename string, a ...interface{}) (*Texture, error) {
	return NewTexture(0, 0), nil
	return &Texture{}, fmt.Errorf("Unable to load texture")
}

func (t *Texture) free() {
	gl.DeleteTextures(1, &t.tex)
}

func (t *Texture) Width(a ...interface{}) int {
	return 0
}

func (t *Texture) Height(a ...interface{}) int {
	return 0
}

func (t *Texture) Lock(a ...interface{}) *uint32 {
	return nil
}

func (t *Texture) Unlock() {
}
