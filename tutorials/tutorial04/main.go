package main

import (
	"fmt"

	"github.com/losinggeneration/hge"
	"github.com/losinggeneration/hge/gfx"
	"github.com/losinggeneration/hge/helpers/font"
	"github.com/losinggeneration/hge/helpers/particle"
	"github.com/losinggeneration/hge/helpers/sprite"
	"github.com/losinggeneration/hge/input"
	"github.com/losinggeneration/hge/sound"
	"github.com/losinggeneration/hge/timer"
)

var (
	spr, spt, tar sprite.Sprite
	fnt           *font.Font
	par           *particle.ParticleSystem

	tex *gfx.Texture
	snd *sound.Effect

	// HGE render target handle
	target *gfx.Target

	x  = 100.0
	y  = 100.0
	dx = 0.0
	dy = 0.0
)

const (
	speed    = 90
	friction = 0.98
)

func boom() {
	pan := int((x - 256) / 2.56)
	pitch := (dx*dx+dy*dy)*0.0005 + 0.2
	snd.PlayEx(100, pan, pitch)
}

// This function will be called by HGE when
// render targets were lost and have been just created
// again. We use it here to update the render
// target's texture handle that changes during recreation.
func RestoreFunc() int {
	if target != nil {
		tar.SetTexture(target.Texture())
	}

	return 0
}

func FrameFunc() int {
	dt := timer.Delta()

	// Process keys
	if input.K_ESCAPE.State() {
		return 1
	}
	if input.K_LEFT.State() {
		dx -= speed * dt
	}
	if input.K_RIGHT.State() {
		dx += speed * dt
	}
	if input.K_UP.State() {
		dy -= speed * dt
	}
	if input.K_DOWN.State() {
		dy += speed * dt
	}

	// Do some movement calculations and collision detection
	dx *= friction
	dy *= friction
	x += dx
	y += dy
	if x > 496 {
		x = 496 - (x - 496)
		dx = -dx
		boom()
	}
	if x < 16 {
		x = 16 + 16 - x
		dx = -dx
		boom()
	}
	if y > 496 {
		y = 496 - (y - 496)
		dy = -dy
		boom()
	}
	if y < 16 {
		y = 16 + 16 - y
		dy = -dy
		boom()
	}

	// Update particle system
	par.Info.Emission = (int)(dx*dx + dy*dy)
	par.MoveTo(x, y)
	par.Update(dt)

	return 0
}

func RenderFunc() int {
	// Render graphics to the texture
	gfx.BeginScene(target)
	gfx.Clear(gfx.RGBAToColor(0))
	par.Render()
	spr.Render(x, y)
	gfx.EndScene()

	// Now put several instances of the rendered texture to the screen
	gfx.BeginScene()
	gfx.Clear(gfx.RGBAToColor(0))
	for i := 0.0; i < 6.0; i++ {
		tar.SetColor(uint32(0xFFFFFF | ((int)((5-i)*40+55) << 24)))
		tar.RenderEx(i*100.0, i*50.0, i*hge.Pi/8, 1.0-i*0.1)
	}
	fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d (constant)", timer.Delta(), timer.FPS())
	gfx.EndScene()

	return 0
}

func main() {
	h := hge.New()

	h.SetState(hge.LOGFILE, "tutorial04.log")
	h.SetState(hge.FRAMEFUNC, FrameFunc)
	h.SetState(hge.RENDERFUNC, RenderFunc)
	h.SetState(hge.GFXRESTOREFUNC, RestoreFunc)
	h.SetState(hge.TITLE, "HGE Tutorial 04 - Using render targets")
	h.SetState(hge.FPS, 100)
	h.SetState(hge.WINDOWED, true)
	h.SetState(hge.SCREENWIDTH, 800)
	h.SetState(hge.SCREENHEIGHT, 600)
	h.SetState(hge.SCREENBPP, 32)

	if err := h.Initiate(); err == nil {
		defer h.Shutdown()
		snd = sound.NewEffect("menu.ogg")
		tex, err = gfx.LoadTexture("particles.png")
		if snd == nil || tex == nil || err != nil {
			// If one of the data files is not found, display
			// an error message and shutdown.
			fmt.Printf("Error: Can't load one of the following files:\nmenu.ogg, particles.png, font1.fnt, font1.png, trail.psi\n")
			return
		}

		spr = sprite.New(tex, 96, 64, 32, 32)
		spr.SetColor(0xFFFFA000)
		spr.SetHotSpot(16, 16)

		fnt = font.New("font1.fnt")

		if fnt == nil {
			fmt.Println("Error: Can't load font1.fnt or font1.png")
			return
		}

		spt = sprite.New(tex, 32, 32, 32, 32)
		spt.SetBlendMode(gfx.BLEND_COLORMUL | gfx.BLEND_ALPHAADD | gfx.BLEND_NOZWRITE)
		spt.SetHotSpot(16, 16)
		par = particle.New("trail.psi", spt)

		if par == nil {
			fmt.Println("Error: Cannot load trail.psi")
			return
		}
		par.Fire()

		// Create a render target and a sprite for it
		target = gfx.NewTarget(512, 512, false)
		defer target.Free()
		tar = sprite.New(target.Texture(), 0, 0, 512, 512)
		tar.SetBlendMode(gfx.BLEND_COLORMUL | gfx.BLEND_ALPHAADD | gfx.BLEND_NOZWRITE)

		// Let's rock now!
		h.Start()
	}
}
