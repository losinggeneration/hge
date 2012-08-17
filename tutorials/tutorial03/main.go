package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/helpers/font"
	"github.com/losinggeneration/hge-go/helpers/particle"
	"github.com/losinggeneration/hge-go/helpers/sprite"
	HGE "github.com/losinggeneration/hge-go/hge"
)

var (
	spr, spt sprite.Sprite
	fnt      *font.Font
	par      *particle.ParticleSystem

	tex HGE.Texture
	snd HGE.Effect

	x  = 100.0
	y  = 100.0
	dx = 0.0
	dy = 0.0
)

const (
	speed    = 90.0
	friction = 0.98
)

func boom() {
	pan := int((x - 400) / 4)
	pitch := (dx*dx+dy*dy)*0.0005 + 0.2
	snd.PlayEx(100, pan, pitch)
}

func FrameFunc() int {
	dt := float64(HGE.NewTimer().Delta())

	// Process keys
	if HGE.NewKey(HGE.K_ESCAPE).State() {
		return 1
	}
	if HGE.NewKey(HGE.K_LEFT).State() {
		dx -= speed * dt
	}
	if HGE.NewKey(HGE.K_RIGHT).State() {
		dx += speed * dt
	}
	if HGE.NewKey(HGE.K_UP).State() {
		dy -= speed * dt
	}
	if HGE.NewKey(HGE.K_DOWN).State() {
		dy += speed * dt
	}

	// Do some movement calculations and collision detection
	dx *= friction
	dy *= friction
	x += dx
	y += dy
	if x > 784 {
		x = 784 - (x - 784)
		dx = -dx
		boom()
	}
	if x < 16 {
		x = 16 + 16 - x
		dx = -dx
		boom()
	}
	if y > 584 {
		y = 584 - (y - 584)
		dy = -dy
		boom()
	}
	if y < 16 {
		y = 16 + 16 - y
		dy = -dy
		boom()
	}

	// Update particle system
	par.Info.Emission = (int)(dx*dx+dy*dy) * 2
	par.MoveTo(x, y)
	par.Update(dt)

	return 0
}

func RenderFunc() int {
	HGE.GfxBeginScene()
	HGE.GfxClear(0)
	// currently broken
	par.Render()
	spr.Render(x, y)
	fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d (constant)", HGE.NewTimer().Delta(), HGE.GetFPS())
	HGE.GfxEndScene()

	return 0
}

func main() {
	defer HGE.Free()

	HGE.SetState(HGE.LOGFILE, "tutorial03.log")
	HGE.SetState(HGE.FRAMEFUNC, FrameFunc)
	HGE.SetState(HGE.RENDERFUNC, RenderFunc)
	HGE.SetState(HGE.TITLE, "HGE Tutorial 03 - Using helper classes")
	HGE.SetState(HGE.FPS, 100)
	HGE.SetState(HGE.WINDOWED, true)
	HGE.SetState(HGE.SCREENWIDTH, 800)
	HGE.SetState(HGE.SCREENHEIGHT, 600)
	HGE.SetState(HGE.SCREENBPP, 32)

	if err := HGE.Initiate(); err == nil {
		defer HGE.Shutdown()
		snd = HGE.NewEffect("menu.ogg")
		tex = HGE.LoadTexture("particles.png")
		if snd == 0 || tex == 0 {
			fmt.Printf("Error: Can't load one of the following files:\nmenu.ogg, particles.png, font1.fnt, font1.png, trail.psi\n")
			return
		}

		defer snd.Free()
		defer tex.Free()

		spr = sprite.NewSprite(tex, 96, 64, 32, 32)
		spr.SetColor(0xFFFFA000)
		spr.SetHotSpot(16, 16)

		if fnt = font.NewFont("font1.fnt"); fnt == nil {
			fmt.Println("Error loading font1.fnt")
			return
		}

		spt = sprite.NewSprite(tex, 32, 32, 32, 32)
		spt.SetBlendMode(HGE.BLEND_COLORMUL | HGE.BLEND_ALPHAADD | HGE.BLEND_NOZWRITE)
		spt.SetHotSpot(16, 16)

		par = particle.NewParticleSystem("trail.psi", spt)
		if par == nil {
			fmt.Println("Error loading trail.psi")
			return
		}
		par.Fire()

		HGE.Start()
	}
}
