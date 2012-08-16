package main

import (
	"fmt"
	. "github.com/losinggeneration/hge-go/helpers/font"
	. "github.com/losinggeneration/hge-go/helpers/particle"
	. "github.com/losinggeneration/hge-go/helpers/sprite"
	hge "github.com/losinggeneration/hge-go/hge"
	. "github.com/losinggeneration/hge-go/legacy"
)

var (
	h *HGE

	spr, spt, tar Sprite
	fnt           *Font
	par           *ParticleSystem

	tex hge.Texture
	snd hge.Effect

	// HGE render target handle
	target hge.Target

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
	h.Effect_PlayEx(snd, 100, pan, pitch)
}

// This function will be called by HGE when
// render targets were lost and have been just created
// again. We use it here to update the render
// target's texture handle that changes during recreation.
func GfxRestoreFunc() int {
	if target != 0 {
		tar.SetTexture(h.Target_GetTexture(target))
	}

	return 0
}

func FrameFunc() int {
	dt := h.Timer_GetDelta()

	// Process keys
	if h.Input_GetKeyState(hge.K_ESCAPE) {
		return 1
	}
	if h.Input_GetKeyState(hge.K_LEFT) {
		dx -= speed * dt
	}
	if h.Input_GetKeyState(hge.K_RIGHT) {
		dx += speed * dt
	}
	if h.Input_GetKeyState(hge.K_UP) {
		dy -= speed * dt
	}
	if h.Input_GetKeyState(hge.K_DOWN) {
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
	h.Gfx_BeginScene(target)
	h.Gfx_Clear(0)
	par.Render()
	spr.Render(x, y)
	h.Gfx_EndScene()

	// Now put several instances of the rendered texture to the screen
	h.Gfx_BeginScene()
	h.Gfx_Clear(0)
	for i := 0.0; i < 6.0; i++ {
		tar.SetColor(hge.Dword(0xFFFFFF | ((int)((5-i)*40+55) << 24)))
		tar.RenderEx(i*100.0, i*50.0, i*hge.Pi/8, 1.0-i*0.1)
	}
	fnt.Printf(5, 5, TEXT_LEFT, "dt:%.3f\nFPS:%d (constant)", h.Timer_GetDelta(), h.Timer_GetFPS())
	h.Gfx_EndScene()

	return 0
}

func main() {
	h = Create(hge.VERSION)
	defer h.Release()

	h.System_SetState(hge.LOGFILE, "tutorial04.log")
	h.System_SetState(hge.FRAMEFUNC, FrameFunc)
	h.System_SetState(hge.RENDERFUNC, RenderFunc)
	h.System_SetState(hge.GFXRESTOREFUNC, GfxRestoreFunc)
	h.System_SetState(hge.TITLE, "HGE Tutorial 04 - Using render targets")
	h.System_SetState(hge.FPS, 100)
	h.System_SetState(hge.WINDOWED, true)
	h.System_SetState(hge.SCREENWIDTH, 800)
	h.System_SetState(hge.SCREENHEIGHT, 600)
	h.System_SetState(hge.SCREENBPP, 32)

	target = 0

	if h.System_Initiate() {
		defer h.System_Shutdown()
		snd = h.Effect_Load("menu.ogg")
		tex = h.Texture_Load("particles.png")
		if snd == 0 || tex == 0 {
			// If one of the data files is not found, display
			// an error message and shutdown.
			fmt.Printf("Error: Can't load one of the following files:\nmenu.ogg, particles.png, font1.fnt, font1.png, trail.psi\n")
			return
		}

		// Delete created objects and free loaded resources
		defer h.Effect_Free(snd)
		defer h.Texture_Free(tex)

		spr = NewSprite(tex, 96, 64, 32, 32)
		spr.SetColor(0xFFFFA000)
		spr.SetHotSpot(16, 16)

		fnt = NewFont("font1.fnt")

		if fnt == nil {
			fmt.Println("Error: Can't load font1.fnt or font1.png")
			return
		}

		spt = NewSprite(tex, 32, 32, 32, 32)
		spt.SetBlendMode(hge.BLEND_COLORMUL | hge.BLEND_ALPHAADD | hge.BLEND_NOZWRITE)
		spt.SetHotSpot(16, 16)
		par = NewParticleSystem("trail.psi", spt)

		if par == nil {
			fmt.Println("Error: Cannot load trail.psi")
			return
		}
		par.Fire()

		// Create a render target and a sprite for it
		target = h.Target_Create(512, 512, false)
		defer h.Target_Free(target)
		tar = NewSprite(h.Target_GetTexture(target), 0, 0, 512, 512)
		tar.SetBlendMode(hge.BLEND_COLORMUL | hge.BLEND_ALPHAADD | hge.BLEND_NOZWRITE)

		// Let's rock now!
		h.System_Start()
	}
}
