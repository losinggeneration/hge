// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
// I doubt there will ever bee the need for anything like: +build sdl,opengl
// or: +build sdl,software
// but it's an option

// +build opengl,1_4

package gfx

// import "fmt"
import (
	"github.com/banthar/Go-SDL/sdl"
	gl "github.com/chsc/gogl/gl21"

// 	"runtime"
)

var (
	width, height gl.Sizei
	zBuffer       bool
	curBlendMode  int = BLEND_DEFAULT
)

func SetWidth(w int) {
	width = gl.Sizei(w)
}

func SetHeight(h int) {
	height = gl.Sizei(h)
}

func SetZBuffer(b bool) {
	zBuffer = b
}

func Initialize() error {
	if err := gl.InitVersion10(); err != nil {
		return err
	}
	if err := gl.InitVersion10Deprecated(); err != nil {
		return err
	}
	if err := gl.InitVersion11(); err != nil {
		return err
	}
	if err := gl.InitVersion11Deprecated(); err != nil {
		return err
	}
	if err := gl.InitVersion12(); err != nil {
		return err
	}
	if err := gl.InitVersion12Deprecated(); err != nil {
		return err
	}
	if err := gl.InitVersion13(); err != nil {
		return err
	}
	if err := gl.InitVersion13Deprecated(); err != nil {
		return err
	}
	if err := gl.InitVersion14(); err != nil {
		return err
	}
	if err := gl.InitVersion14Deprecated(); err != nil {
		return err
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	return nil
}

func setProjectionMatrix() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, gl.Double(width), 0, gl.Double(height), 0.0, 1.0)
}

func BeginScene(a ...interface{}) bool {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	if zBuffer {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}

	gl.Scissor(0, 0, width, height)
	gl.Viewport(0, 0, width, height)
	setProjectionMatrix()

	return true
}

func EndScene() {
	gl.Finish()
	sdl.GL_SwapBuffers()
}

func Clear(color Color) {
	gl.ClearColor(gl.Float(color.R), gl.Float(color.G), gl.Float(color.B), gl.Float(color.A))
}

func setBlendMode(blend int) {
	if (blend & BLEND_ALPHABLEND) != (curBlendMode & BLEND_ALPHABLEND) {
		if blend&BLEND_ALPHABLEND == BLEND_ALPHABLEND {
			gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		} else {
			gl.BlendFunc(gl.SRC_ALPHA, gl.ONE)
		}
	}

	if (blend & BLEND_ZWRITE) != (curBlendMode & BLEND_ZWRITE) {
		if blend&BLEND_ZWRITE == BLEND_ZWRITE {
			gl.DepthMask(gl.TRUE)
		} else {
			gl.DepthMask(gl.FALSE)
		}
	}

	if (blend & BLEND_COLORADD) != (curBlendMode & BLEND_COLORADD) {
		if blend&BLEND_COLORADD == BLEND_COLORADD {
			gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.ADD)
		} else {
			gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
		}
	}

	curBlendMode = blend
}

// Converts a RGBA uint32 to a Color structure
func RGBAToColor(c uint32) Color {
	return Color{R: c >> 24, G: (c >> 16) & 0xFF, B: (c >> 8) & 0xFF, A: c & 0xFF}
}

// Converts an ARGB uint32 to a Color structure
func ARGBToColor(c uint32) Color {
	return Color{A: c >> 24, R: (c >> 16) & 0xFF, G: (c >> 8) & 0xFF, B: c & 0xFF}
}

func NewLine(x1, y1, x2, y2 float64, a ...interface{}) Line {
	color := ARGBToColor(0xFFFFFFFF)
	z := 0.5

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if c, ok := a[i].(uint); ok {
				color = ARGBToColor(uint32(c))
			}
			if c, ok := a[i].(uint32); ok {
				color = ARGBToColor(c)
			}
			if c, ok := a[i].(Color); ok {
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
	gl.Begin(gl.LINES)
	gl.Color3ub(gl.Ubyte(l.Color.R), gl.Ubyte(l.Color.G), gl.Ubyte(l.Color.B))
	gl.Vertex2d(gl.Double(l.X1), gl.Double(l.Y1))
	gl.Vertex2d(gl.Double(l.X2), gl.Double(l.Y2))
	gl.End()
}

func (t *Triple) Render() {
}

func (v *Vertex) Render() {
	gl.Color4ub(gl.Ubyte(v.Color.R), gl.Ubyte(v.Color.G), gl.Ubyte(v.Color.B), gl.Ubyte(v.Color.A))
	gl.Vertex2d(gl.Double(v.X), gl.Double(v.Y))
}

func (q *Quad) Render() {
	if q.Texture != nil {
		q.Texture.bind()
	}
	setBlendMode(q.Blend)

	gl.Begin(gl.QUADS)
	for _, v := range q.V {
		if q.Texture != nil {
			q.Texture.coord(Vertex{X: v.TX, Y: 1 - v.TY})
		}
		v.Render()
	}
	gl.End()
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
