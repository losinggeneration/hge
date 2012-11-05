package animation

import (
	"github.com/losinggeneration/hge-go/binding/helpers/sprite"
	"github.com/losinggeneration/hge-go/binding/hge/gfx"
)

const (
	FWD        = 0
	REV        = 1
	PINGPONG   = 2
	NOPINGPONG = 0
	LOOP       = 4
	NOLOOP     = 0
)

type Animation struct {
	sprite.Sprite
	origWidth                     int
	playing                       bool
	speed, sinceLastFrame         float64
	mode, delta, frames, curFrame int
}

func New(tex *gfx.Texture, frames int, fps, x, y, w, h float64) Animation {
	var a Animation

	a.Sprite = sprite.New(tex, x, y, w, h)

	a.origWidth = tex.Width(true)

	a.sinceLastFrame = -1.0
	a.speed = 1.0 / fps
	a.frames = frames

	a.mode = FWD | LOOP
	a.delta = 1
	a.SetFrame(0)

	return a
}

func (a *Animation) Play() {
	a.playing = true
	a.sinceLastFrame = -1.0
	if a.mode&REV == REV {
		a.delta = -1
		a.SetFrame(a.frames - 1)
	} else {
		a.delta = 1
		a.SetFrame(0)
	}
}

func (a *Animation) Stop() {
	a.playing = false
}

func (a *Animation) Resume() {
	a.playing = true
}

func (a *Animation) Update(deltaTime float64) {
	if !a.playing {
		return
	}

	if a.sinceLastFrame == -1.0 {
		a.sinceLastFrame = 0.0
	} else {
		a.sinceLastFrame += deltaTime
	}

	for a.sinceLastFrame >= a.speed {
		a.sinceLastFrame -= a.speed

		if a.curFrame+a.delta == a.frames {
			switch a.mode {
			case FWD,
				REV | PINGPONG:
				a.playing = false

			case FWD | PINGPONG,
				FWD | PINGPONG | LOOP,
				REV | PINGPONG | LOOP:
				a.delta = -a.delta
			}
		} else if a.curFrame+a.delta < 0 {
			switch a.mode {
			case REV,
				FWD | PINGPONG:
				a.playing = false

			case REV | PINGPONG,
				REV | PINGPONG | LOOP,
				FWD | PINGPONG | LOOP:
				a.delta = -a.delta
			}
		}

		if a.playing {
			a.SetFrame(a.curFrame + a.delta)
		}
	}
}

func (a Animation) IsPlaying() bool {
	return a.playing
}

func (a *Animation) SetTexture(tex *gfx.Texture) {
	a.Sprite.SetTexture(tex)
	a.origWidth = tex.Width(true)
}

func (a *Animation) SetTextureRect(x1, y1, x2, y2 float64) {
	a.Sprite.SetTextureRect(x1, y1, x2, y2)
	a.SetFrame(a.curFrame)
}

func (a *Animation) SetMode(mode int) {
	a.mode = mode

	if mode&REV == REV {
		a.delta = -1
		a.SetFrame(a.frames - 1)
	} else {
		a.delta = 1
		a.SetFrame(0)
	}
}

func (a *Animation) SetSpeed(fps float64) {
	a.speed = 1.0 / fps
}

func (a *Animation) SetFrame(n int) {
	// 	float tx1, ty1, tx2, ty2;
	// 	bool bX, bY, bHS;
	// 	int ncols = int(orig_width) / int(width);

	cols := int(a.origWidth / int(a.W))

	n = n % a.frames
	if n < 0 {
		n = a.frames + n
	}
	a.curFrame = n

	// calculate texture coords for frame n
	ty1 := a.TY
	tx1 := a.TX + float64(n)*a.W

	if tx1 > float64(a.origWidth)-a.W {
		n -= int(float64(a.origWidth) - a.TX/a.W)
		tx1 = a.W * float64(n%cols)
		ty1 += a.H * float64(1+n/cols)
	}

	tx2 := tx1 + a.W
	ty2 := ty1 + a.H

	tx1 /= a.TexW
	ty1 /= a.TexH
	tx2 /= a.TexW
	ty2 /= a.TexH

	a.Quad.V[0].TX, a.Quad.V[0].TY = float32(tx1), float32(ty1)
	a.Quad.V[1].TX, a.Quad.V[1].TY = float32(tx2), float32(ty1)
	a.Quad.V[2].TX, a.Quad.V[2].TY = float32(tx2), float32(ty2)
	a.Quad.V[3].TX, a.Quad.V[3].TY = float32(tx1), float32(ty2)

	bX := a.XFlip
	bY := a.YFlip
	bHS := a.HSFlip
	a.XFlip = false
	a.YFlip = false
	a.SetFlip(bX, bY, bHS)
}

func (a *Animation) SetFrames(n int) {
	a.frames = n
}

func (a Animation) Mode() int {
	return a.mode
}

func (a Animation) Speed() float64 {
	return 1.0 / a.speed
}

func (a Animation) Frame() int {
	return a.curFrame
}

func (a Animation) Frames() int {
	return a.frames
}
