package hge

const (
	ANIM_FWD        = 0
	ANIM_REV        = 1
	ANIM_PINGPONG   = 2
	ANIM_NOPINGPONG = 0
	ANIM_LOOP       = 4
	ANIM_NOLOOP     = 0
)

type Animation struct {
	sprite                        Sprite
	origWidth                     int
	playing                       bool
	speed, sinceLastFrame         float64
	mode, delta, frames, curFrame int
}

func NewAnimation(tex Texture, frames int, fps, x, y, w, h float64) Animation {
	var a Animation

	a.sprite = NewSprite(tex, x, y, w, h)

	a.origWidth = a.sprite.Hge.Texture_GetWidth(tex, true)

	a.sinceLastFrame = -1.0
	a.speed = 1.0 / fps
	a.playing = false
	a.frames = frames

	a.mode = ANIM_FWD | ANIM_LOOP
	a.delta = 1
	a.SetFrame(0)

	return a
}

func (a *Animation) Play() {
	a.playing = true
	a.sinceLastFrame = -1.0
	if a.mode&ANIM_REV == ANIM_REV {
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
			case ANIM_FWD,
				ANIM_REV | ANIM_PINGPONG:
				a.playing = false

			case ANIM_FWD | ANIM_PINGPONG,
				ANIM_FWD | ANIM_PINGPONG | ANIM_LOOP,
				ANIM_REV | ANIM_PINGPONG | ANIM_LOOP:
				a.delta = -a.delta
			}
		} else if a.curFrame+a.delta < 0 {
			switch a.mode {
			case ANIM_REV,
				ANIM_FWD | ANIM_PINGPONG:
				a.playing = false

			case ANIM_REV | ANIM_PINGPONG,
				ANIM_REV | ANIM_PINGPONG | ANIM_LOOP,
				ANIM_FWD | ANIM_PINGPONG | ANIM_LOOP:
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

func (a *Animation) SetTexture(tex Texture) {
	a.sprite.SetTexture(tex)
	a.origWidth = a.sprite.Hge.Texture_GetWidth(tex, true)
}

func (a *Animation) SetTextureRect(x1, y1, x2, y2 float64) {
	a.sprite.SetTextureRect(x1, y1, x2, y2)
	a.SetFrame(a.curFrame)
}

func (a *Animation) SetMode(mode int) {
	a.mode = mode

	if mode&ANIM_REV == ANIM_REV {
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

	cols := int(a.origWidth / int(a.sprite.Width))

	n = n % a.frames
	if n < 0 {
		n = a.frames + n
	}
	a.curFrame = n

	// calculate texture coords for frame n
	ty1 := a.sprite.TY
	tx1 := a.sprite.TX + float64(n)*a.sprite.Width

	if tx1 > float64(a.origWidth)-a.sprite.Width {
		n -= int(float64(a.origWidth) - a.sprite.TX/a.sprite.Width)
		tx1 = a.sprite.Width * float64(n%cols)
		ty1 += a.sprite.Height * float64(1+n/cols)
	}

	tx2 := tx1 + a.sprite.Width
	ty2 := ty1 + a.sprite.Height

	tx1 /= a.sprite.TexWidth
	ty1 /= a.sprite.TexHeight
	tx2 /= a.sprite.TexWidth
	ty2 /= a.sprite.TexHeight

	a.sprite.Quad.V[0].TX, a.sprite.Quad.V[0].TY = float32(tx1), float32(ty1)
	a.sprite.Quad.V[1].TX, a.sprite.Quad.V[1].TY = float32(tx2), float32(ty1)
	a.sprite.Quad.V[2].TX, a.sprite.Quad.V[2].TY = float32(tx2), float32(ty2)
	a.sprite.Quad.V[3].TX, a.sprite.Quad.V[3].TY = float32(tx1), float32(ty2)

	bX := a.sprite.XFlip
	bY := a.sprite.YFlip
	bHS := a.sprite.HSFlip
	a.sprite.XFlip = false
	a.sprite.YFlip = false
	a.sprite.SetFlip(bX, bY, bHS)
}

func (a *Animation) SetFrames(n int) {
	a.frames = n
}

func (a Animation) GetMode() int {
	return a.mode
}

func (a Animation) GetSpeed() float64 {
	return 1.0 / a.speed
}

func (a Animation) GetFrame() int {
	return a.curFrame
}

func (a Animation) GetFrames() int {
	return a.frames
}
