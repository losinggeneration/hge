package hge

import (
	"math"
)

type Sprite struct {
	Hge *HGE

	Quad                  Quad
	TX, TY, Width, Height float64
	TexWidth, TexHeight   float64
	HotX, HotY            float64
	XFlip, YFlip, HSFlip  bool
}

func NewSprite(texture Texture, texx, texy, w, h float64) Sprite {
	var sprite Sprite

	sprite.Hge = Create(VERSION)

	sprite.TX = texx
	sprite.TY = texy
	sprite.Width = w
	sprite.Height = h

	if texture != 0 {
		sprite.TexWidth = float64(sprite.Hge.Texture_GetWidth(texture))
		sprite.TexHeight = float64(sprite.Hge.Texture_GetHeight(texture))
	} else {
		sprite.TexWidth = 1.0
		sprite.TexHeight = 1.0
	}

	sprite.HotX = 0
	sprite.HotY = 0
	sprite.XFlip = false
	sprite.YFlip = false
	sprite.HSFlip = false
	sprite.Quad.Tex = texture

	texx1 := texx / sprite.TexWidth
	texy1 := texy / sprite.TexHeight
	texx2 := (texx + w) / sprite.TexWidth
	texy2 := (texy + h) / sprite.TexHeight

	sprite.Quad.V[0].TX = float32(texx1)
	sprite.Quad.V[0].TY = float32(texy1)
	sprite.Quad.V[1].TX = float32(texx2)
	sprite.Quad.V[1].TY = float32(texy1)
	sprite.Quad.V[2].TX = float32(texx2)
	sprite.Quad.V[2].TY = float32(texy2)
	sprite.Quad.V[3].TX = float32(texx1)
	sprite.Quad.V[3].TY = float32(texy2)

	sprite.Quad.V[0].Z = 0.5
	sprite.Quad.V[1].Z = 0.5
	sprite.Quad.V[2].Z = 0.5
	sprite.Quad.V[3].Z = 0.5

	sprite.Quad.V[0].Col = 0xffffffff
	sprite.Quad.V[1].Col = 0xffffffff
	sprite.Quad.V[2].Col = 0xffffffff
	sprite.Quad.V[3].Col = 0xffffffff

	sprite.Quad.Blend = BLEND_DEFAULT

	return sprite
}

func (sprite *Sprite) Render(x, y float64) {
	tempx1 := x - sprite.HotX
	tempy1 := y - sprite.HotY
	tempx2 := x + sprite.Width - sprite.HotX
	tempy2 := y + sprite.Height - sprite.HotY

	sprite.Quad.V[0].X = float32(tempx1)
	sprite.Quad.V[0].Y = float32(tempy1)
	sprite.Quad.V[1].X = float32(tempx2)
	sprite.Quad.V[1].Y = float32(tempy1)
	sprite.Quad.V[2].X = float32(tempx2)
	sprite.Quad.V[2].Y = float32(tempy2)
	sprite.Quad.V[3].X = float32(tempx1)
	sprite.Quad.V[3].Y = float32(tempy2)

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) RenderEx(x, y float64, rot float64, arg ...interface{}) {
	var tx1, ty1, tx2, ty2 float64
	var sint, cost float64

	hscale := 1.0
	vscale := 0.0

	for i := 0; i < len(arg); i++ {
		if i == 0 {
			if h, ok := arg[i].(float64); ok {
				hscale = h
			}
		}
		if i == 1 {
			if v, ok := arg[i].(float64); ok {
				vscale = v
			}
		}
	}

	if vscale == 0 {
		vscale = hscale
	}

	tx1 = -sprite.HotX * hscale
	ty1 = -sprite.HotY * vscale
	tx2 = (sprite.Width - sprite.HotX) * hscale
	ty2 = (sprite.Height - sprite.HotY) * vscale

	if rot != 0.0 {
		cost = math.Cos(rot)
		sint = math.Sin(rot)

		sprite.Quad.V[0].X = float32(tx1*cost - ty1*sint + x)
		sprite.Quad.V[0].Y = float32(tx1*sint + ty1*cost + y)

		sprite.Quad.V[1].X = float32(tx2*cost - ty1*sint + x)
		sprite.Quad.V[1].Y = float32(tx2*sint + ty1*cost + y)

		sprite.Quad.V[2].X = float32(tx2*cost - ty2*sint + x)
		sprite.Quad.V[2].Y = float32(tx2*sint + ty2*cost + y)

		sprite.Quad.V[3].X = float32(tx1*cost - ty2*sint + x)
		sprite.Quad.V[3].Y = float32(tx1*sint + ty2*cost + y)
	} else {
		sprite.Quad.V[0].X = float32(tx1 + x)
		sprite.Quad.V[0].Y = float32(ty1 + y)
		sprite.Quad.V[1].X = float32(tx2 + x)
		sprite.Quad.V[1].Y = float32(ty1 + y)
		sprite.Quad.V[2].X = float32(tx2 + x)
		sprite.Quad.V[2].Y = float32(ty2 + y)
		sprite.Quad.V[3].X = float32(tx1 + x)
		sprite.Quad.V[3].Y = float32(ty2 + y)
	}

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) RenderStretch(x1, y1, x2, y2 float64) {
	sprite.Quad.V[0].X = float32(x1)
	sprite.Quad.V[0].Y = float32(y1)
	sprite.Quad.V[1].X = float32(x2)
	sprite.Quad.V[1].Y = float32(y1)
	sprite.Quad.V[2].X = float32(x2)
	sprite.Quad.V[2].Y = float32(y2)
	sprite.Quad.V[3].X = float32(x1)
	sprite.Quad.V[3].Y = float32(y2)

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) Render4V(x0, y0, x1, y1, x2, y2, x3, y3 float64) {
	sprite.Quad.V[0].X = float32(x0)
	sprite.Quad.V[0].Y = float32(y0)
	sprite.Quad.V[1].X = float32(x1)
	sprite.Quad.V[1].Y = float32(y1)
	sprite.Quad.V[2].X = float32(x2)
	sprite.Quad.V[2].Y = float32(y2)
	sprite.Quad.V[3].X = float32(x3)
	sprite.Quad.V[3].Y = float32(y3)

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) SetTexture(tex Texture) {
	var tx1, ty1, tx2, ty2 float64
	var tw, th float64

	sprite.Quad.Tex = tex

	if tex != 0 {
		tw = float64(sprite.Hge.Texture_GetWidth(tex))
		th = float64(sprite.Hge.Texture_GetHeight(tex))
	} else {
		tw = 1.0
		th = 1.0
	}

	if tw != sprite.TexWidth || th != sprite.TexHeight {
		tx1 = float64(sprite.Quad.V[0].TX) * sprite.TexWidth
		ty1 = float64(sprite.Quad.V[0].TY) * sprite.TexHeight
		tx2 = float64(sprite.Quad.V[2].TX) * sprite.TexWidth
		ty2 = float64(sprite.Quad.V[2].TY) * sprite.TexHeight

		sprite.TexWidth = tw
		sprite.TexHeight = th

		tx1 /= tw
		ty1 /= th
		tx2 /= tw
		ty2 /= th

		sprite.Quad.V[0].TX = float32(tx1)
		sprite.Quad.V[0].TY = float32(ty1)
		sprite.Quad.V[1].TX = float32(tx2)
		sprite.Quad.V[1].TY = float32(ty1)
		sprite.Quad.V[2].TX = float32(tx2)
		sprite.Quad.V[2].TY = float32(ty2)
		sprite.Quad.V[3].TX = float32(tx1)
		sprite.Quad.V[3].TY = float32(ty2)
	}
}

func (sprite *Sprite) SetTextureRect(x, y, w, h float64, adjSize bool) {
	sprite.TX = x
	sprite.TY = y

	if adjSize {
		sprite.Width = w
		sprite.Height = h
	}

	tx1 := sprite.TX / sprite.TexWidth
	ty1 := sprite.TY / sprite.TexHeight
	tx2 := (sprite.TX + w) / sprite.TexWidth
	ty2 := (sprite.TY + h) / sprite.TexHeight

	sprite.Quad.V[0].TX = float32(tx1)
	sprite.Quad.V[0].TY = float32(ty1)
	sprite.Quad.V[1].TX = float32(tx2)
	sprite.Quad.V[1].TY = float32(ty1)
	sprite.Quad.V[2].TX = float32(tx2)
	sprite.Quad.V[2].TY = float32(ty2)
	sprite.Quad.V[3].TX = float32(tx1)
	sprite.Quad.V[3].TY = float32(ty2)

	bX := sprite.XFlip
	bY := sprite.YFlip
	bHS := sprite.HSFlip
	sprite.XFlip = false
	sprite.YFlip = false
	sprite.SetFlip(bX, bY, bHS)
}

func (sprite *Sprite) SetColor(col Dword, arg ...interface{}) {
	i := -1

	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	if i != -1 {
		sprite.Quad.V[i].Col = col
	} else {
		sprite.Quad.V[0].Col = col
		sprite.Quad.V[1].Col = col
		sprite.Quad.V[2].Col = col
		sprite.Quad.V[3].Col = col
	}
}

func (sprite *Sprite) SetZ(z float64, arg ...interface{}) {
	i := -1

	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	if i != -1 {
		sprite.Quad.V[i].Z = float32(z)
	} else {
		sprite.Quad.V[0].Z = float32(z)
		sprite.Quad.V[1].Z = float32(z)
		sprite.Quad.V[2].Z = float32(z)
		sprite.Quad.V[3].Z = float32(z)
	}
}

func (sprite *Sprite) SetBlendMode(blend int) {
	sprite.Quad.Blend = blend
}

func (sprite *Sprite) SetHotSpot(x, y float64) {
	sprite.HotX = x
	sprite.HotY = y
}

func (sprite *Sprite) SetFlip(x, y, hotSpot bool) {
	var tx, ty float64

	if sprite.HSFlip && sprite.XFlip {
		sprite.HotX = sprite.Width - sprite.HotX
	}
	if sprite.HSFlip && sprite.YFlip {
		sprite.HotY = sprite.Height - sprite.HotY
	}

	sprite.HSFlip = hotSpot

	if sprite.HSFlip && sprite.XFlip {
		sprite.HotX = sprite.Width - sprite.HotX
	}
	if sprite.HSFlip && sprite.YFlip {
		sprite.HotY = sprite.Height - sprite.HotY
	}

	if x != sprite.XFlip {
		tx = float64(sprite.Quad.V[0].TX)
		sprite.Quad.V[0].TX = sprite.Quad.V[1].TX
		sprite.Quad.V[1].TX = float32(tx)
		ty = float64(sprite.Quad.V[0].TY)
		sprite.Quad.V[0].TY = sprite.Quad.V[1].TY
		sprite.Quad.V[1].TY = float32(ty)
		tx = float64(sprite.Quad.V[3].TX)
		sprite.Quad.V[3].TX = sprite.Quad.V[2].TX
		sprite.Quad.V[2].TX = float32(tx)
		ty = float64(sprite.Quad.V[3].TY)
		sprite.Quad.V[3].TY = sprite.Quad.V[2].TY
		sprite.Quad.V[2].TY = float32(ty)

		sprite.XFlip = !sprite.XFlip
	}

	if y != sprite.YFlip {
		tx = float64(sprite.Quad.V[0].TX)
		sprite.Quad.V[0].TX = sprite.Quad.V[3].TX
		sprite.Quad.V[3].TX = float32(tx)
		ty = float64(sprite.Quad.V[0].TY)
		sprite.Quad.V[0].TY = sprite.Quad.V[3].TY
		sprite.Quad.V[3].TY = float32(ty)
		tx = float64(sprite.Quad.V[1].TX)
		sprite.Quad.V[1].TX = sprite.Quad.V[2].TX
		sprite.Quad.V[2].TX = float32(tx)
		ty = float64(sprite.Quad.V[1].TY)
		sprite.Quad.V[1].TY = sprite.Quad.V[2].TY
		sprite.Quad.V[2].TY = float32(ty)

		sprite.YFlip = !sprite.YFlip
	}
}

func (sprite *Sprite) GetTexture() Texture {
	return sprite.Quad.Tex
}

func (sprite *Sprite) GetTextureRect() (x, y, w, h float64) {
	return sprite.TX, sprite.TY, sprite.Width, sprite.Height
}

func (sprite *Sprite) GetColor(arg ...interface{}) Dword {
	i := 0
	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	return sprite.Quad.V[i].Col
}

func (sprite *Sprite) GetZ(arg ...interface{}) float64 {
	i := 0
	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	return float64(sprite.Quad.V[i].Z)
}

func (sprite *Sprite) GetBlendMode() int {
	return sprite.Quad.Blend
}

func (sprite *Sprite) GetHotSpot() (x, y float64) {
	x = sprite.HotX
	y = sprite.HotY
	return
}

func (sprite *Sprite) GetFlip() (x, y bool) {
	x = sprite.XFlip
	y = sprite.YFlip
	return
}

func (sprite *Sprite) GetWidth() float64 {
	return sprite.Width
}

func (sprite *Sprite) GetHeight() float64 {
	return sprite.Height
}

func (sprite *Sprite) GetBoundingBox(x, y float64, rect *Rect) *Rect {
	rect.Set(x-sprite.HotX, y-sprite.HotY, x-sprite.HotX+sprite.Width, y-sprite.HotY+sprite.Height)
	return rect
}

func (sprite *Sprite) GetBoundingBoxEx(x, y, rot, hscale, vscale float64, rect *Rect) *Rect {
	var tx1, ty1, tx2, ty2 float64
	var sint, cost float64

	rect.Clear()

	tx1 = -sprite.HotX * hscale
	ty1 = -sprite.HotY * vscale
	tx2 = (sprite.Width - sprite.HotX) * hscale
	ty2 = (sprite.Height - sprite.HotY) * vscale

	if rot != 0.0 {
		cost = math.Cos(rot)
		sint = math.Sin(rot)

		rect.Encapsulate(tx1*cost-ty1*sint+x, tx1*sint+ty1*cost+y)
		rect.Encapsulate(tx2*cost-ty1*sint+x, tx2*sint+ty1*cost+y)
		rect.Encapsulate(tx2*cost-ty2*sint+x, tx2*sint+ty2*cost+y)
		rect.Encapsulate(tx1*cost-ty2*sint+x, tx1*sint+ty2*cost+y)
	} else {
		rect.Encapsulate(tx1+x, ty1+y)
		rect.Encapsulate(tx2+x, ty1+y)
		rect.Encapsulate(tx2+x, ty2+y)
		rect.Encapsulate(tx1+x, ty2+y)
	}

	return rect
}
