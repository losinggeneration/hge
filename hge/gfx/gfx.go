package gfx

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import (
	. "github.com/losinggeneration/hge-go/hge"
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
	X, Y   float32 // screen position
	Z      float32 // Z-buffer depth 0..1
	Col    Dword   // color
	TX, TY float32 // texture coordinates
}

// HGE Triple structure
type Triple struct {
	V     [3]Vertex
	Tex   Texture
	Blend int
}

// HGE Quad structure
type Quad struct {
	V     [4]Vertex
	Tex   Texture
	Blend int
}

func boolToCInt(b bool) C.BOOL {
	return C.BOOL(BoolToCInt(b))
}

func BeginScene(a ...interface{}) bool {
	if len(a) == 1 {
		if target, ok := a[0].(Target); ok {
			return C.HGE_Gfx_BeginScene(HGE, C.HTARGET(target)) == 1
		}
	}

	return C.HGE_Gfx_BeginScene(HGE, 0) == 1
}

func EndScene() {
	C.HGE_Gfx_EndScene(HGE)
}

func Clear(color Dword) {
	C.HGE_Gfx_Clear(HGE, C.DWORD(color))
}

func RenderLine(x1, y1, x2, y2 float64, a ...interface{}) {
	color := uint(0xFFFFFFFF)
	z := 0.5

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if c, ok := a[i].(uint); ok {
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

	C.HGE_Gfx_RenderLine(HGE, C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.DWORD(color), C.float(z))
}

func (t *Triple) Render() {
	C.HGE_Gfx_RenderTriple(HGE, (*C.HGE_Triple_t)(unsafe.Pointer(t)))
}

func (q *Quad) Render() {
	C.HGE_Gfx_RenderQuad(HGE, (*C.HGE_Quad_t)(unsafe.Pointer(q)))
}

func StartBatch(prim_type int, tex Texture, blend int) (ver *Vertex, max_prim int, ok bool) {
	mp := C.int(0)

	v := C.HGE_Gfx_StartBatch(HGE, C.int(prim_type), C.HTEXTURE(tex), C.int(blend), &mp)

	if v == nil {
		return nil, 0, false
	}

	return (*Vertex)(unsafe.Pointer(v)), int(mp), true
}

func FinishBatch(prim int) {
	C.HGE_Gfx_FinishBatch(HGE, C.int(prim))
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

	C.HGE_Gfx_SetClipping(HGE, C.int(x), C.int(y), C.int(w), C.int(hi))
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

	C.HGE_Gfx_SetTransform(HGE, C.float(x), C.float(y), C.float(dx), C.float(dy), C.float(rot), C.float(hscale), C.float(vscale))
}

// HGE Handle type
type Target C.HTARGET

func NewTarget(width, height int, zbuffer bool) Target {
	return Target(C.HGE_Target_Create(HGE, C.int(width), C.int(height), boolToCInt(zbuffer)))
}

func (t Target) Free() {
	C.HGE_Target_Free(HGE, C.HTARGET(t))
}

func (t Target) Texture() Texture {
	return Texture(C.HGE_Target_GetTexture(HGE, C.HTARGET(t)))
}

// HGE Handle type
type Texture C.HTEXTURE

func NewTexture(width, height int) Texture {
	return Texture(C.HGE_Texture_Create(HGE, C.int(width), C.int(height)))
}

func LoadTexture(filename string, a ...interface{}) Texture {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	size := Dword(0)
	mipmap := false

	for i := 0; i < len(a); i++ {
		if i == 0 {
			if s, ok := a[i].(Dword); ok {
				size = s
			}
		}
		if i == 1 {
			if m, ok := a[i].(bool); ok {
				mipmap = m
			}
		}
	}

	return Texture(C.HGE_Texture_Load(HGE, fname, C.DWORD(size), boolToCInt(mipmap)))
}

func (t Texture) Free() {
	C.HGE_Texture_Free(HGE, C.HTEXTURE(t))
}

func (t Texture) Width(a ...interface{}) int {
	if len(a) == 1 {
		if original, ok := a[0].(bool); ok {
			return int(C.HGE_Texture_GetWidth(HGE, C.HTEXTURE(t), boolToCInt(original)))
		}
	}

	return int(C.HGE_Texture_GetWidth(HGE, C.HTEXTURE(t), boolToCInt(false)))
}

func (t Texture) Height(a ...interface{}) int {
	if len(a) == 1 {
		if original, ok := a[0].(bool); ok {
			return int(C.HGE_Texture_GetWidth(HGE, C.HTEXTURE(t), boolToCInt(original)))
		}
	}

	return int(C.HGE_Texture_GetHeight(HGE, C.HTEXTURE(t), boolToCInt(false)))
}

func (t Texture) Lock(a ...interface{}) *Dword {
	readonly := true
	left, top, width, height := 0, 0, 0, 0

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

	d := C.HGE_Texture_Lock(HGE, C.HTEXTURE(t), boolToCInt(readonly), C.int(left), C.int(top), C.int(width), C.int(height))
	return (*Dword)(d)
}

func (t Texture) Unlock() {
	C.HGE_Texture_Unlock(HGE, C.HTEXTURE(t))
}
