package main

import (
	"fmt"
	"hge"
)

var (
	h *hge.HGE

	spr, spt hge.Sprite
	fnt      *hge.Font
	// par *ParticleSystem

	tex hge.Texture
	snd hge.Effect

	x  = float32(100.0)
	y  = float32(100.0)
	dx = float32(0.0)
	dy = float32(0.0)
)

const (
	speed    = 90.0
	friction = 0.98
)

func boom() {
	pan := int((x - 400) / 4)
	pitch := (dx*dx+dy*dy)*0.0005 + 0.2
	h.Effect_PlayEx(snd, 100, pan, pitch)
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
	// par->info.nEmission=(int)(dx*dx+dy*dy)*2;
	// par->MoveTo(x,y);
	// par->Update(dt);

	return 0
}

func RenderFunc() int {
	h.Gfx_BeginScene()
	h.Gfx_Clear(0)
	// par->Render();
	spr.Render(x, y)
	fnt.Printf(5, 5, hge.TEXT_LEFT, "dt:%.3f\nFPS:%d (constant)", h.Timer_GetDelta(), h.Timer_GetFPS())
	h.Gfx_EndScene()

	return 0
}

func main() {
	h = hge.Create(hge.VERSION)

	h.System_SetState(hge.LOGFILE, "tutorial03.log")
	h.System_SetState(hge.FRAMEFUNC, FrameFunc)
	h.System_SetState(hge.RENDERFUNC, RenderFunc)
	h.System_SetState(hge.TITLE, "HGE Tutorial 03 - Using helper classes")
	h.System_SetState(hge.FPS, 100)
	h.System_SetState(hge.WINDOWED, true)
	h.System_SetState(hge.SCREENWIDTH, 800)
	h.System_SetState(hge.SCREENHEIGHT, 600)
	h.System_SetState(hge.SCREENBPP, 32)

	if h.System_Initiate() {
		snd = h.Effect_Load("menu.ogg")
		tex = h.Texture_Load("particles.png")
		if snd == 0 || tex == 0 {
			fmt.Printf("Error: Can't load one of the following files:\nmenu.ogg, particles.png, font1.fnt, font1.png, trail.psi\n")
			h.System_Shutdown()
			h.Release()
			return
		}

		spr = hge.NewSprite(tex, 96, 64, 32, 32)
		spr.SetColor(0xFFFFA000)
		spr.SetHotSpot(16, 16)

		if fnt = hge.NewFont("font1.fnt"); fnt == nil {
			fmt.Println("Error loading font1.fnt")
			return
		}

		spt = hge.NewSprite(tex, 32, 32, 32, 32)
		spt.SetBlendMode(hge.BLEND_COLORMUL | hge.BLEND_ALPHAADD | hge.BLEND_NOZWRITE)
		spt.SetHotSpot(16, 16)
		// par=new hgeParticleSystem("trail.psi",spt);
		// par->Fire();

		h.System_Start()

		h.Texture_Free(tex)
		h.Effect_Free(snd)
	}

	h.System_Shutdown()
	h.Release()
}
