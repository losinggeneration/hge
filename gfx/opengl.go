package gfx

import gl "github.com/go-gl/gl/v2.1/gl"

var (
	width, height      int32
	x, y               int
	zBuffer            bool
	curBlendMode       int = BLEND_DEFAULT
	defaultTextureType uint32
)

func SetWidth(w int) {
	width = int32(w)
	updateSize(int(width), int(height))
}

func SetHeight(h int) {
	height = int32(h)
	updateSize(int(width), int(height))
}

func SetX(i int) {
	x = i
	updatePosition(x, y)
}

func SetY(i int) {
	y = i
	updatePosition(x, y)
}

func SetZBuffer(b bool) {
	zBuffer = b
}

func Initialize() error {
	gl.Init()

	// For now, just TEXTURE_2D
	defaultTextureType = gl.TEXTURE_2D

	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	gl.PixelStorei(gl.PACK_ALIGNMENT, 1)

	gl.Disable(gl.TEXTURE_2D)
	gl.Enable(defaultTextureType)
	gl.Enable(gl.SCISSOR_TEST)
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.LIGHTING)
	gl.DepthFunc(gl.GEQUAL)

	if zBuffer {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Enable(gl.ALPHA_TEST)
	gl.AlphaFunc(gl.GEQUAL, 1.0/255.0)

	gl.TexEnvi(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)

	default_texture_filter()

	gl.Scissor(0, 0, width, height)
	gl.Viewport(0, 0, width, height)

	// make sure the framebuffer is cleared and force to screen
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	setProjectionMatrix()
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	return nil
}

func setProjectionMatrix() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(width), 0, float64(height), 0.0, 1.0)
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
	swapBuffers()
}

func Clear(color Color) {
	gl.ClearColor(float32(color.R), float32(color.G), float32(color.B), float32(color.A))
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
			gl.DepthMask(true)
		} else {
			gl.DepthMask(false)
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
	gl.Color3ub(uint8(l.Color.R), uint8(l.Color.G), uint8(l.Color.B))
	gl.Vertex2d(float64(l.X1), float64(l.Y1))
	gl.Vertex2d(float64(l.X2), float64(l.Y2))
	gl.End()
}

func (t *Triple) Render() {
	if t.Texture != nil {
		t.Texture.bind()
		default_texture_filter()
	}
	setBlendMode(t.Blend)

	gl.Begin(gl.TRIANGLES)
	for _, v := range t.V {
		if t.Texture != nil {
			// The Y axis has to be inverted for OpenGL, we have the top left
			// corner (0,0) and it's the bottom left in OpenGL
			tex_coord(Vertex{X: v.TX, Y: 1 - v.TY})
		}

		v.Render()
	}
	gl.End()
}

func (v *Vertex) Render() {
	gl.Color4ub(uint8(v.Color.R), uint8(v.Color.G), uint8(v.Color.B), uint8(v.Color.A))
	gl.Vertex2d(float64(v.X), float64(v.Y))
}

func (q *Quad) Render() {
	if q.Texture != nil {
		q.Texture.bind()
		default_texture_filter()
	}
	setBlendMode(q.Blend)

	gl.Begin(gl.QUADS)
	for _, v := range q.V {
		if q.Texture != nil {
			// The Y axis has to be inverted for OpenGL, we have the top left
			// corner (0,0) and it's the bottom left in OpenGL
			tex_coord(Vertex{X: v.TX, Y: 1 - v.TY})
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
