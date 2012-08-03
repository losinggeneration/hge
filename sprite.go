package hge

import (
	"math"
)

type Sprite struct {
	Hge *HGE

	Quad                  Quad
	TX, TY, Width, Height float32
	TexWidth, TexHeight   float32
	HotX, HotY            float32
	XFlip, YFlip, HSFlip  bool
}

func NewSprite(texture Texture, texx, texy, w, h float32) Sprite {
	var texx1, texy1, texx2, texy2 float32
	var sprite Sprite

	sprite.Hge = Create(VERSION)

	sprite.TX = texx
	sprite.TY = texy
	sprite.Width = w
	sprite.Height = h

	if texture != 0 {
		sprite.TexWidth = float32(sprite.Hge.Texture_GetWidth(texture))
		sprite.TexHeight = float32(sprite.Hge.Texture_GetHeight(texture))
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

	texx1 = texx / sprite.TexWidth
	texy1 = texy / sprite.TexHeight
	texx2 = (texx + w) / sprite.TexWidth
	texy2 = (texy + h) / sprite.TexHeight

	sprite.Quad.V[0].TX = texx1
	sprite.Quad.V[0].TY = texy1
	sprite.Quad.V[1].TX = texx2
	sprite.Quad.V[1].TY = texy1
	sprite.Quad.V[2].TX = texx2
	sprite.Quad.V[2].TY = texy2
	sprite.Quad.V[3].TX = texx1
	sprite.Quad.V[3].TY = texy2

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

func (sprite *Sprite) Render(x, y float32) {
	tempx1 := x - sprite.HotX
	tempy1 := y - sprite.HotY
	tempx2 := x + sprite.Width - sprite.HotX
	tempy2 := y + sprite.Height - sprite.HotY

	sprite.Quad.V[0].X = tempx1
	sprite.Quad.V[0].Y = tempy1
	sprite.Quad.V[1].X = tempx2
	sprite.Quad.V[1].Y = tempy1
	sprite.Quad.V[2].X = tempx2
	sprite.Quad.V[2].Y = tempy2
	sprite.Quad.V[3].X = tempx1
	sprite.Quad.V[3].Y = tempy2

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) RenderEx(x, y float32, rot float64, arg ...interface{}) {
	var tx1, ty1, tx2, ty2 float32
	var sint, cost float32

	hscale := float32(1.0)
	vscale := float32(0.0)

	for i := 0; i < len(arg); i++ {
		if i == 0 {
			if h, ok := arg[i].(float32); ok {
				hscale = h
			}
		}
		if i == 1 {
			if v, ok := arg[i].(float32); ok {
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
		cost = float32(math.Cos(rot))
		sint = float32(math.Sin(rot))

		sprite.Quad.V[0].X = tx1*cost - ty1*sint + x
		sprite.Quad.V[0].Y = tx1*sint + ty1*cost + y

		sprite.Quad.V[1].X = tx2*cost - ty1*sint + x
		sprite.Quad.V[1].Y = tx2*sint + ty1*cost + y

		sprite.Quad.V[2].X = tx2*cost - ty2*sint + x
		sprite.Quad.V[2].Y = tx2*sint + ty2*cost + y

		sprite.Quad.V[3].X = tx1*cost - ty2*sint + x
		sprite.Quad.V[3].Y = tx1*sint + ty2*cost + y
	} else {
		sprite.Quad.V[0].X = tx1 + x
		sprite.Quad.V[0].Y = ty1 + y
		sprite.Quad.V[1].X = tx2 + x
		sprite.Quad.V[1].Y = ty1 + y
		sprite.Quad.V[2].X = tx2 + x
		sprite.Quad.V[2].Y = ty2 + y
		sprite.Quad.V[3].X = tx1 + x
		sprite.Quad.V[3].Y = ty2 + y
	}

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) RenderStretch(x1, y1, x2, y2 float32) {
	sprite.Quad.V[0].X = x1
	sprite.Quad.V[0].Y = y1
	sprite.Quad.V[1].X = x2
	sprite.Quad.V[1].Y = y1
	sprite.Quad.V[2].X = x2
	sprite.Quad.V[2].Y = y2
	sprite.Quad.V[3].X = x1
	sprite.Quad.V[3].Y = y2

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) Render4V(x0, y0, x1, y1, x2, y2, x3, y3 float32) {
	sprite.Quad.V[0].X = x0
	sprite.Quad.V[0].Y = y0
	sprite.Quad.V[1].X = x1
	sprite.Quad.V[1].Y = y1
	sprite.Quad.V[2].X = x2
	sprite.Quad.V[2].Y = y2
	sprite.Quad.V[3].X = x3
	sprite.Quad.V[3].Y = y3

	sprite.Hge.Gfx_RenderQuad(&sprite.Quad)
}

func (sprite *Sprite) SetTexture(tex Texture) {
	var tx1, ty1, tx2, ty2 float32
	var tw, th float32

	sprite.Quad.Tex = tex

	if tex != 0 {
		tw = float32(sprite.Hge.Texture_GetWidth(tex))
		th = float32(sprite.Hge.Texture_GetHeight(tex))
	} else {
		tw = 1.0
		th = 1.0
	}

	if tw != sprite.TexWidth || th != sprite.TexHeight {
		tx1 = sprite.Quad.V[0].TX * sprite.TexWidth
		ty1 = sprite.Quad.V[0].TY * sprite.TexHeight
		tx2 = sprite.Quad.V[2].TX * sprite.TexWidth
		ty2 = sprite.Quad.V[2].TY * sprite.TexHeight

		sprite.TexWidth = tw
		sprite.TexHeight = th

		tx1 /= tw
		ty1 /= th
		tx2 /= tw
		ty2 /= th

		sprite.Quad.V[0].TX = tx1
		sprite.Quad.V[0].TY = ty1
		sprite.Quad.V[1].TX = tx2
		sprite.Quad.V[1].TY = ty1
		sprite.Quad.V[2].TX = tx2
		sprite.Quad.V[2].TY = ty2
		sprite.Quad.V[3].TX = tx1
		sprite.Quad.V[3].TY = ty2
	}
}

func (sprite *Sprite) SetTextureRect(x, y, w, h float32, adjSize bool) {
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

	sprite.Quad.V[0].TX = tx1
	sprite.Quad.V[0].TY = ty1
	sprite.Quad.V[1].TX = tx2
	sprite.Quad.V[1].TY = ty1
	sprite.Quad.V[2].TX = tx2
	sprite.Quad.V[2].TY = ty2
	sprite.Quad.V[3].TX = tx1
	sprite.Quad.V[3].TY = ty2

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

func (sprite *Sprite) SetZ(z float32, arg ...interface{}) {
	i := -1

	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	if i != -1 {
		sprite.Quad.V[i].Z = z
	} else {
		sprite.Quad.V[0].Z = z
		sprite.Quad.V[1].Z = z
		sprite.Quad.V[2].Z = z
		sprite.Quad.V[3].Z = z
	}
}

func (sprite *Sprite) SetBlendMode(blend int) {
	sprite.Quad.Blend = blend
}

func (sprite *Sprite) SetHotSpot(x, y float32) {
	sprite.HotX = x
	sprite.HotY = y
}

func (sprite *Sprite) SetFlip(x, y, hotSpot bool) {
	var tx, ty float32

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
		tx = sprite.Quad.V[0].TX
		sprite.Quad.V[0].TX = sprite.Quad.V[1].TX
		sprite.Quad.V[1].TX = tx
		ty = sprite.Quad.V[0].TY
		sprite.Quad.V[0].TY = sprite.Quad.V[1].TY
		sprite.Quad.V[1].TY = ty
		tx = sprite.Quad.V[3].TX
		sprite.Quad.V[3].TX = sprite.Quad.V[2].TX
		sprite.Quad.V[2].TX = tx
		ty = sprite.Quad.V[3].TY
		sprite.Quad.V[3].TY = sprite.Quad.V[2].TY
		sprite.Quad.V[2].TY = ty

		sprite.XFlip = !sprite.XFlip
	}

	if y != sprite.YFlip {
		tx = sprite.Quad.V[0].TX
		sprite.Quad.V[0].TX = sprite.Quad.V[3].TX
		sprite.Quad.V[3].TX = tx
		ty = sprite.Quad.V[0].TY
		sprite.Quad.V[0].TY = sprite.Quad.V[3].TY
		sprite.Quad.V[3].TY = ty
		tx = sprite.Quad.V[1].TX
		sprite.Quad.V[1].TX = sprite.Quad.V[2].TX
		sprite.Quad.V[2].TX = tx
		ty = sprite.Quad.V[1].TY
		sprite.Quad.V[1].TY = sprite.Quad.V[2].TY
		sprite.Quad.V[2].TY = ty

		sprite.YFlip = !sprite.YFlip
	}
}

func (sprite *Sprite) GetTexture() Texture {
	return sprite.Quad.Tex
}

func (sprite *Sprite) GetTextureRect() (x, y, w, h float32) {
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

func (sprite *Sprite) GetZ(arg ...interface{}) float32 {
	i := 0
	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	return sprite.Quad.V[i].Z
}

func (sprite *Sprite) GetBlendMode() int {
	return sprite.Quad.Blend
}

func (sprite *Sprite) GetHotSpot() (x, y float32) {
	x = sprite.HotX
	y = sprite.HotY
	return
}

func (sprite *Sprite) GetFlip() (x, y bool) {
	x = sprite.XFlip
	y = sprite.YFlip
	return
}

func (sprite *Sprite) GetWidth() float32 {
	return sprite.Width
}

func (sprite *Sprite) GetHeight() float32 {
	return sprite.Height
}

func (sprite *Sprite) GetBoundingBox(x, y float32, rect *Rect) *Rect {
	rect.Set(x-sprite.HotX, y-sprite.HotY, x-sprite.HotX+sprite.Width, y-sprite.HotY+sprite.Height)
	return rect
}

func (sprite *Sprite) GetBoundingBoxEx(x, y, rot, hscale, vscale float32, rect *Rect) *Rect {
	var tx1, ty1, tx2, ty2 float32
	var sint, cost float32

	rect.Clear()

	tx1 = -sprite.HotX * hscale
	ty1 = -sprite.HotY * vscale
	tx2 = (sprite.Width - sprite.HotX) * hscale
	ty2 = (sprite.Height - sprite.HotY) * vscale

	if rot != 0.0 {
		cost = float32(math.Cos(float64(rot)))
		sint = float32(math.Sin(float64(rot)))

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
