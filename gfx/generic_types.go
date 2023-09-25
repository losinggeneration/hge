package gfx

// HGE Blending constants
const (
	BLEND_COLORADD   = 1
	BLEND_COLORMUL   = 0
	BLEND_ALPHABLEND = 2
	BLEND_ALPHAADD   = 0
	BLEND_ZWRITE     = 4
	BLEND_NOZWRITE   = 0

	BLEND_DEFAULT   = (BLEND_COLORMUL | BLEND_ALPHABLEND | BLEND_NOZWRITE)
	BLEND_DEFAULT_Z = (BLEND_COLORMUL | BLEND_ALPHABLEND | BLEND_ZWRITE)
)

// HGE Primitive type constants
const (
	PrimLines   = iota
	PrimTriples = iota
	PrimQuads   = iota
)

type Color struct {
	R, G, B, A uint32
}

// HGE Vertex structure
type Vertex struct {
	X, Y   float32 // screen position
	Z      float32 // Z-buffer depth 0..1
	Color  Color   // color
	TX, TY float32 // texture coordinates
}

type Line struct {
	X1, Y1, X2, Y2 float64
	Z              float64
	Color          Color
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
