package sprite

import (
	"github.com/losinggeneration/hge-go/binding/helpers/rect"
	"github.com/losinggeneration/hge-go/binding/hge"
	"github.com/losinggeneration/hge-go/binding/hge/gfx"
	"math"
)

type Sprite struct {
	gfx.Quad
	TX, TY, W, H         float64
	TexW, TexH           float64
	HotX, HotY           float64
	XFlip, YFlip, HSFlip bool
}

func New(texture *gfx.Texture, texx, texy, w, h float64) Sprite {
	var sprite Sprite

	sprite.TX, sprite.TY = texx, texy
	sprite.W, sprite.H = w, h

	if texture != nil {
		sprite.TexW = float64(texture.Width())
		sprite.TexH = float64(texture.Height())
	} else {
		sprite.TexW = 1.0
		sprite.TexH = 1.0
	}

	sprite.Quad.Texture = texture

	texx1 := texx / sprite.TexW
	texy1 := texy / sprite.TexH
	texx2 := (texx + w) / sprite.TexW
	texy2 := (texy + h) / sprite.TexH

	sprite.Quad.V[0].TX, sprite.Quad.V[0].TY = float32(texx1), float32(texy1)
	sprite.Quad.V[1].TX, sprite.Quad.V[1].TY = float32(texx2), float32(texy1)
	sprite.Quad.V[2].TX, sprite.Quad.V[2].TY = float32(texx2), float32(texy2)
	sprite.Quad.V[3].TX, sprite.Quad.V[3].TY = float32(texx1), float32(texy2)

	sprite.Quad.V[0].Z = 0.5
	sprite.Quad.V[1].Z = 0.5
	sprite.Quad.V[2].Z = 0.5
	sprite.Quad.V[3].Z = 0.5

	sprite.Quad.V[0].Color = 0xffffffff
	sprite.Quad.V[1].Color = 0xffffffff
	sprite.Quad.V[2].Color = 0xffffffff
	sprite.Quad.V[3].Color = 0xffffffff

	sprite.Quad.Blend = gfx.BLEND_DEFAULT

	return sprite
}

func (sprite *Sprite) Render(x, y float64) {
	tempx1 := float32(x - sprite.HotX)
	tempy1 := float32(y - sprite.HotY)
	tempx2 := float32(x + sprite.W - sprite.HotX)
	tempy2 := float32(y + sprite.H - sprite.HotY)

	sprite.Quad.V[0].X, sprite.Quad.V[0].Y = tempx1, tempy1
	sprite.Quad.V[1].X, sprite.Quad.V[1].Y = tempx2, tempy1
	sprite.Quad.V[2].X, sprite.Quad.V[2].Y = tempx2, tempy2
	sprite.Quad.V[3].X, sprite.Quad.V[3].Y = tempx1, tempy2

	sprite.Quad.Render()
}

func (sprite *Sprite) RenderEx(x, y float64, rot float64, arg ...interface{}) {
	var tx1, ty1, tx2, ty2 float64
	var sint, cost float64

	hscale, vscale := 1.0, 0.0

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
	tx2 = (sprite.W - sprite.HotX) * hscale
	ty2 = (sprite.H - sprite.HotY) * vscale

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

	sprite.Quad.Render()
}

func (sprite *Sprite) RenderStretch(x1, y1, x2, y2 float64) {
	sprite.Quad.V[0].X, sprite.Quad.V[0].Y = float32(x1), float32(y1)
	sprite.Quad.V[1].X, sprite.Quad.V[1].Y = float32(x2), float32(y1)
	sprite.Quad.V[2].X, sprite.Quad.V[2].Y = float32(x2), float32(y2)
	sprite.Quad.V[3].X, sprite.Quad.V[3].Y = float32(x1), float32(y2)

	sprite.Quad.Render()
}

func (sprite *Sprite) Render4V(x0, y0, x1, y1, x2, y2, x3, y3 float64) {
	sprite.Quad.V[0].X, sprite.Quad.V[0].Y = float32(x0), float32(y0)
	sprite.Quad.V[1].X, sprite.Quad.V[1].Y = float32(x1), float32(y1)
	sprite.Quad.V[2].X, sprite.Quad.V[2].Y = float32(x2), float32(y2)
	sprite.Quad.V[3].X, sprite.Quad.V[3].Y = float32(x3), float32(y3)

	sprite.Quad.Render()
}

func (sprite *Sprite) SetTexture(tex *gfx.Texture) {
	var tw, th float64

	sprite.Quad.Texture = tex

	if tex != nil {
		tw = float64(tex.Width())
		th = float64(tex.Height())
	} else {
		tw, th = 1.0, 1.0
	}

	if tw != sprite.TexW || th != sprite.TexH {
		tx1 := float64(sprite.Quad.V[0].TX) * sprite.TexW
		ty1 := float64(sprite.Quad.V[0].TY) * sprite.TexH
		tx2 := float64(sprite.Quad.V[2].TX) * sprite.TexW
		ty2 := float64(sprite.Quad.V[2].TY) * sprite.TexH

		sprite.TexW, sprite.TexH = tw, th

		tx1 /= tw
		ty1 /= th
		tx2 /= tw
		ty2 /= th

		sprite.Quad.V[0].TX, sprite.Quad.V[0].TY = float32(tx1), float32(ty1)
		sprite.Quad.V[1].TX, sprite.Quad.V[1].TY = float32(tx2), float32(ty1)
		sprite.Quad.V[2].TX, sprite.Quad.V[2].TY = float32(tx2), float32(ty2)
		sprite.Quad.V[3].TX, sprite.Quad.V[3].TY = float32(tx1), float32(ty2)
	}
}

func (sprite *Sprite) SetTextureRect(x, y, w, h float64, a ...interface{}) {
	adjSize := true

	if len(a) == 1 {
		if b, ok := a[0].(bool); ok {
			adjSize = b
		}
	}

	sprite.TX, sprite.TY = x, y

	if adjSize {
		sprite.W, sprite.H = w, h
	}

	tx1 := sprite.TX / sprite.TexW
	ty1 := sprite.TY / sprite.TexH
	tx2 := (sprite.TX + w) / sprite.TexW
	ty2 := (sprite.TY + h) / sprite.TexH

	sprite.Quad.V[0].TX, sprite.Quad.V[0].TY = float32(tx1), float32(ty1)
	sprite.Quad.V[1].TX, sprite.Quad.V[1].TY = float32(tx2), float32(ty1)
	sprite.Quad.V[2].TX, sprite.Quad.V[2].TY = float32(tx2), float32(ty2)
	sprite.Quad.V[3].TX, sprite.Quad.V[3].TY = float32(tx1), float32(ty2)

	bX, bY, bHS := sprite.XFlip, sprite.YFlip, sprite.HSFlip
	sprite.XFlip, sprite.YFlip = false, false

	sprite.SetFlip(bX, bY, bHS)
}

func (sprite *Sprite) SetColor(col hge.Dword, arg ...interface{}) {
	i := -1

	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	if i != -1 {
		sprite.Quad.V[i].Color = col
	} else {
		sprite.Quad.V[0].Color = col
		sprite.Quad.V[1].Color = col
		sprite.Quad.V[2].Color = col
		sprite.Quad.V[3].Color = col
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
	sprite.HotX, sprite.HotY = x, y
}

func (sprite *Sprite) SetFlip(x, y, hotSpot bool) {
	var tx, ty float64

	if sprite.HSFlip && sprite.XFlip {
		sprite.HotX = sprite.W - sprite.HotX
	}
	if sprite.HSFlip && sprite.YFlip {
		sprite.HotY = sprite.H - sprite.HotY
	}

	sprite.HSFlip = hotSpot

	if sprite.HSFlip && sprite.XFlip {
		sprite.HotX = sprite.W - sprite.HotX
	}
	if sprite.HSFlip && sprite.YFlip {
		sprite.HotY = sprite.H - sprite.HotY
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

func (sprite *Sprite) Texture() *gfx.Texture {
	return sprite.Quad.Texture
}

func (sprite *Sprite) TextureRect() (x, y, w, h float64) {
	return sprite.TX, sprite.TY, sprite.W, sprite.H
}

func (sprite *Sprite) Color(arg ...interface{}) hge.Dword {
	i := 0
	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	return sprite.Quad.V[i].Color
}

func (sprite *Sprite) Z(arg ...interface{}) float64 {
	i := 0
	if len(arg) == 1 {
		if ni, ok := arg[0].(int); ok {
			i = ni
		}
	}

	return float64(sprite.Quad.V[i].Z)
}

func (sprite *Sprite) BlendMode() int {
	return sprite.Quad.Blend
}

func (sprite *Sprite) HotSpot() (x, y float64) {
	x, y = sprite.HotX, sprite.HotY
	return
}

func (sprite *Sprite) Flip() (x, y bool) {
	x, y = sprite.XFlip, sprite.YFlip
	return
}

func (sprite *Sprite) Width() float64 {
	return sprite.W
}

func (sprite *Sprite) Height() float64 {
	return sprite.H
}

func (sprite *Sprite) BoundingBox(x, y float64) *rect.Rect {
	return rect.New(x-sprite.HotX, y-sprite.HotY, x-sprite.HotX+sprite.W, y-sprite.HotY+sprite.H)
}

func (sprite *Sprite) BoundingBoxEx(x, y, rot, hscale, vscale float64) *rect.Rect {
	var tx1, ty1, tx2, ty2 float64
	var sint, cost float64

	rect := new(rect.Rect)

	tx1 = -sprite.HotX * hscale
	ty1 = -sprite.HotY * vscale
	tx2 = (sprite.W - sprite.HotX) * hscale
	ty2 = (sprite.H - sprite.HotY) * vscale

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
