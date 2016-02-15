package gfx

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import (
	"fmt"
	"github.com/losinggeneration/hge-go/hge"
	"runtime"
	"unsafe"
)

// HGE Blending constants
const (
	BLEND_COLORADD   = C.BLEND_COLORADD
	BLEND_COLORMUL   = C.BLEND_COLORMUL
	BLEND_ALPHABLEND = C.BLEND_ALPHABLEND
	BLEND_ALPHAADD   = C.BLEND_ALPHAADD
	BLEND_ZWRITE     = C.BLEND_ZWRITE
	BLEND_NOZWRITE   = C.BLEND_NOZWRITE

	BLEND_DEFAULT   = C.BLEND_DEFAULT
	BLEND_DEFAULT_Z = C.BLEND_DEFAULT_Z
)

// HGE_FPS system state special constants
const (
	FPS_UNLIMITED = C.HGE_FPS_UNLIMITED
	FPS_VSYNC     = C.HGE_FPS_VSYNC
)

// HGE Primitive type constants
const (
	PRIM_LINES   = C.HGE_PRIM_LINES
	PRIM_TRIPLES = C.HGE_PRIM_TRIPLES
	PRIM_QUADS   = C.HGE_PRIM_QUADS
)

// HGE Vertex structure
type Vertex struct {
	X, Y   float32   // screen position
	Z      float32   // Z-buffer depth 0..1
	Color  hge.Dword // color
	TX, TY float32   // texture coordinates
}

type Line struct {
	X1, Y1, X2, Y2 float64
	Z              float64
	Color          hge.Dword
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
	tex   C.HTEXTURE
	Blend int
}

type cQuad struct {
	V     [4]Vertex
	tex   C.HTEXTURE
	Blend int
}

func boolToCInt(b bool) C.BOOL {
	return C.BOOL(hge.BoolToCInt(b))
}

var gfxHGE *hge.HGE

func init() {
	gfxHGE = hge.New()
}

func BeginScene(a ...interface{}) bool {
	if len(a) == 1 {
		if target, ok := a[0].(Target); ok {
			return C.HGE_Gfx_BeginScene(gfxHGE.HGE, target.target) == 1
		}
		if target, ok := a[0].(*Target); ok {
			return C.HGE_Gfx_BeginScene(gfxHGE.HGE, target.target) == 1
		}
	}

	return C.HGE_Gfx_BeginScene(gfxHGE.HGE, 0) == 1
}

func EndScene() {
	C.HGE_Gfx_EndScene(gfxHGE.HGE)
}

func Clear(color hge.Dword) {
	C.HGE_Gfx_Clear(gfxHGE.HGE, C.DWORD(color))
}

func NewLine(x1, y1, x2, y2 float64, a ...interface{}) Line {
	color := hge.Dword(0xFFFFFFFF)
	z := 0.5

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if c, ok := a[i].(uint); ok {
				color = hge.Dword(c)
			}
			if c, ok := a[i].(hge.Dword); ok {
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
	C.HGE_Gfx_RenderLine(gfxHGE.HGE, C.float(l.X1), C.float(l.Y1), C.float(l.X2), C.float(l.Y2), C.DWORD(l.Color), C.float(l.Z))
}

func (t *Triple) Render() {
	var ct *cTriple

	if t.Texture != nil {
		ct = &cTriple{t.V, t.Texture.texture, t.Blend}
	} else {
		ct = &cTriple{t.V, 0, t.Blend}
	}
	C.HGE_Gfx_RenderTriple(gfxHGE.HGE, (*C.HGE_Triple_t)(unsafe.Pointer(ct)))
}

func (q *Quad) Render() {
	var cq *cQuad

	if q.Texture != nil {
		cq = &cQuad{q.V, q.Texture.texture, q.Blend}
	} else {
		cq = &cQuad{q.V, 0, q.Blend}
	}
	C.HGE_Gfx_RenderQuad(gfxHGE.HGE, (*C.HGE_Quad_t)(unsafe.Pointer(cq)))
}

func StartBatch(prim_type int, tex *Texture, blend int) (ver *Vertex, max_prim int, ok bool) {
	mp := C.int(0)

	v := C.HGE_Gfx_StartBatch(gfxHGE.HGE, C.int(prim_type), tex.texture, C.int(blend), &mp)

	if v == nil {
		return nil, 0, false
	}

	return (*Vertex)(unsafe.Pointer(v)), int(mp), true
}

func FinishBatch(prim int) {
	C.HGE_Gfx_FinishBatch(gfxHGE.HGE, C.int(prim))
}

func SetClipping(a ...interface{}) {
	var x, y, w, hi int

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if x1, ok := a[i].(int); ok {
				x = x1
			}
		}
		if i == 1 {
			if y1, ok := a[i].(int); ok {
				y = y1
			}
		}
		if i == 2 {
			if w1, ok := a[i].(int); ok {
				w = w1
			}
		}
		if i == 3 {
			if h1, ok := a[i].(int); ok {
				hi = h1
			}
		}
	}

	C.HGE_Gfx_SetClipping(gfxHGE.HGE, C.int(x), C.int(y), C.int(w), C.int(hi))
}

func SetTransform(a ...interface{}) {
	var (
		x, y, dx, dy        float64
		rot, hscale, vscale float64
	)

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if x1, ok := a[i].(float32); ok {
				x = float64(x1)
			}
			if x1, ok := a[i].(float64); ok {
				x = x1
			}
		}
		if i == 1 {
			if y1, ok := a[i].(float32); ok {
				y = float64(y1)
			}
			if y1, ok := a[i].(float64); ok {
				y = y1
			}
		}
		if i == 2 {
			if dx1, ok := a[i].(float32); ok {
				dx = float64(dx1)
			}
			if dx1, ok := a[i].(float64); ok {
				dx = dx1
			}
		}
		if i == 3 {
			if dy1, ok := a[i].(float32); ok {
				dy = float64(dy1)
			}
			if dy1, ok := a[i].(float64); ok {
				dy = dy1
			}
		}
		if i == 4 {
			if rot1, ok := a[i].(float32); ok {
				rot = float64(rot1)
			}
			if rot1, ok := a[i].(float64); ok {
				rot = rot1
			}
		}
		if i == 5 {
			if hscale1, ok := a[i].(float32); ok {
				hscale = float64(hscale1)
			}
			if hscale1, ok := a[i].(float64); ok {
				hscale = hscale1
			}
		}
		if i == 6 {
			if vscale1, ok := a[i].(float32); ok {
				vscale = float64(vscale1)
			}
			if vscale1, ok := a[i].(float64); ok {
				vscale = vscale1
			}
		}
	}

	C.HGE_Gfx_SetTransform(gfxHGE.HGE, C.float(x), C.float(y), C.float(dx), C.float(dy), C.float(rot), C.float(hscale), C.float(vscale))
}

// HGE Handle type
type Target struct {
	target C.HTARGET
}

func NewTarget(width, height int, zbuffer bool) *Target {
	t := new(Target)
	t.target = C.HGE_Target_Create(gfxHGE.HGE, C.int(width), C.int(height), boolToCInt(zbuffer))

	if t.target == 0 {
		return nil
	}

	runtime.SetFinalizer(t, func(target *Target) {
		target.Free()
	})

	return t
}

func (t *Target) Free() {
	fmt.Println("Freeing Target", t)
	C.HGE_Target_Free(gfxHGE.HGE, t.target)
}

func (t *Target) Texture() *Texture {
	return &Texture{C.HGE_Target_GetTexture(gfxHGE.HGE, t.target)}
}

// HGE Handle type
type Texture struct {
	texture C.HTEXTURE
}

func NewTexture(width, height int) *Texture {
	t := new(Texture)
	t.texture = C.HGE_Texture_Create(gfxHGE.HGE, C.int(width), C.int(height))

	if t.texture == 0 {
		return nil
	}

	runtime.SetFinalizer(t, func(texture *Texture) {
		texture.Free()
	})

	return t
}

func LoadTexture(filename string, a ...interface{}) *Texture {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	size := hge.Dword(0)
	mipmap := false

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if s, ok := a[i].(hge.Dword); ok {
				size = s
			}
		}
		if i == 1 {
			if m, ok := a[i].(bool); ok {
				mipmap = m
			}
		}
	}

	t := new(Texture)
	t.texture = C.HGE_Texture_Load(gfxHGE.HGE, fname, C.DWORD(size), boolToCInt(mipmap))
	if t.texture == 0 {
		return nil
	}

	runtime.SetFinalizer(t, func(texture *Texture) {
		texture.Free()
	})

	return t
}

func (t *Texture) Free() {
	fmt.Println("Freeing texture", t)
	C.HGE_Texture_Free(gfxHGE.HGE, t.texture)
}

func (t *Texture) Width(a ...interface{}) int {
	original := false
	if len(a) == 1 {
		if o, ok := a[0].(bool); ok {
			original = o
		}
	}

	return int(C.HGE_Texture_GetWidth(gfxHGE.HGE, t.texture, boolToCInt(original)))
}

func (t *Texture) Height(a ...interface{}) int {
	original := false
	if len(a) == 1 {
		if o, ok := a[0].(bool); ok {
			original = o
		}
	}

	return int(C.HGE_Texture_GetHeight(gfxHGE.HGE, t.texture, boolToCInt(original)))
}

func (t *Texture) Lock(a ...interface{}) *hge.Dword {
	readonly := true
	var left, top, width, height int

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if r, ok := a[i].(bool); ok {
				readonly = r
			}
		}
		if i == 1 {
			if l, ok := a[i].(int); ok {
				left = l
			}
		}
		if i == 2 {
			if t, ok := a[i].(int); ok {
				top = t
			}
		}
		if i == 3 {
			if w, ok := a[i].(int); ok {
				width = w
			}
		}
		if i == 4 {
			if h, ok := a[i].(int); ok {
				height = h
			}
		}
	}

	d := C.HGE_Texture_Lock(gfxHGE.HGE, t.texture, boolToCInt(readonly), C.int(left), C.int(top), C.int(width), C.int(height))
	return (*hge.Dword)(d)
}

func (t *Texture) Unlock() {
	C.HGE_Texture_Unlock(gfxHGE.HGE, t.texture)
}
