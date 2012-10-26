package gfx

import (
	"github.com/losinggeneration/hge-go/hge"
)

// HGE Blending constants
const (
	BLEND_COLORADD   = iota
	BLEND_COLORMUL   = iota
	BLEND_ALPHABLEND = iota
	BLEND_ALPHAADD   = iota
	BLEND_ZWRITE     = iota
	BLEND_NOZWRITE   = iota

	BLEND_DEFAULT   = iota
	BLEND_DEFAULT_Z = iota
)

// HGE_FPS system state special constants
const (
	FPS_UNLIMITED = iota
	FPS_VSYNC     = iota
)

// HGE Primitive type constants
const (
	PRIM_LINES   = iota
	PRIM_TRIPLES = iota
	PRIM_QUADS   = iota
)

// HGE Vertex structure
type Vertex struct {
	X, Y   float32 // screen position
	Z      float32 // Z-buffer depth 0..1
	Color  uint32  // color
	TX, TY float32 // texture coordinates
}

type Line struct {
	X1, Y1, X2, Y2 float64
	Z              float64
	Color          uint32
}

// HGE Triple structure
type Triple struct {
	V [3]Vertex
	*Texture
	Blend int
}

// HGE Quad structure
type Quad struct {
	V [4]Vertex
	*Texture
	Blend int
}

type cTriple struct {
	v     [3]Vertex
	tex   interface{}
	Blend int
}

type cQuad struct {
	V     [4]Vertex
	tex   interface{}
	Blend int
}

var gfxHGE *hge.HGE

func init() {
	gfxHGE = hge.New()
}

func BeginScene(a ...interface{}) bool {
	return false
}

func EndScene() {
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
	texture interface{}
}

func NewTexture(width, height int) *Texture {
	return nil
}

func LoadTexture(filename string, a ...interface{}) *Texture {
	return nil
}

func (t *Texture) Free() {
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
