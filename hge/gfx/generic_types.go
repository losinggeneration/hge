package gfx

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
