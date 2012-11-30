/*
 * Haaf's Game Engine 1.8
 * Copyright (C) 2003-2007, Relish Games
 * hge.relishgames.com
 *
 * hge_tut02 - Using input, sound and rendering using the legacy API
 */

package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/binding/hge/gfx"
	"github.com/losinggeneration/hge-go/binding/hge/sound"
	hge "github.com/losinggeneration/hge-go/binding/legacy"

// 	"github.com/losinggeneration/hge-go/binding/hge/input"
)

var (
	h *hge.HGE

	quad gfx.Quad
	line gfx.Line

	snd *sound.Effect

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
	h.Effect_PlayEx(snd, 100, pan, pitch)
}

func FrameFunc() bool {
	dt := float64(h.Timer_GetDelta())

	// Process keys
	if h.Input_GetKeyState(hge.K_ESCAPE) {
		return true
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

	// Set up quad's screen coordinates
	quad.V[0].X = float32(x - 16)
	quad.V[0].Y = float32(y - 16)
	quad.V[1].X = float32(x + 16)
	quad.V[1].Y = float32(y - 16)
	quad.V[2].X = float32(x + 16)
	quad.V[2].Y = float32(y + 16)
	quad.V[3].X = float32(x - 16)
	quad.V[3].Y = float32(y + 16)

	// Continue execution
	return false
}

// This function will be called by HGE when
// the application window should be redrawn.
// Put your rendering code here.
func RenderFunc() bool {
	// Begin rendering quads.
	// This function must be called
	// before any actual rendering.
	h.Gfx_BeginScene()

	// Clear screen with black color
	h.Gfx_Clear(0)

	// Render quads here. This time just
	// one of them will serve our needs.
	h.Gfx_RenderQuad(&quad)
	line.Render()
	h.Gfx_RenderLine(100, 100, 150, 150, uint32(0xFFFFFFFF))

	// End rendering and update the screen
	h.Gfx_EndScene()

	// RenderFunc should always return false
	return false
}

func main() {
	// Get HGE interface
	h = hge.Create(hge.VERSION)
	defer h.Release()

	// Set up log file, frame function, render function and window title
	h.System_SetState(hge.LOGFILE, "tutorial02.log")
	h.System_SetState(hge.FRAMEFUNC, FrameFunc)
	h.System_SetState(hge.RENDERFUNC, RenderFunc)
	h.System_SetState(hge.TITLE, "HGE Tutorial 02 - Using input, sound and rendering")

	// Set up video mode
	h.System_SetState(hge.WINDOWED, true)
	h.System_SetState(hge.SCREENWIDTH, 800)
	h.System_SetState(hge.SCREENHEIGHT, 600)
	h.System_SetState(hge.SCREENBPP, 32)

	if h.System_Initiate() {
		defer h.System_Shutdown()

		// Load sound and texture
		snd = h.Effect_Load("menu.ogg")
		quad.Texture = h.Texture_Load("particles.png")
		if /*snd == nil ||*/ quad.Texture == nil {
			// If one of the data files is not found, display
			// an error message and shutdown.
			fmt.Println("Error: Can't load menu.ogg or particles.png")
			return
		}
		defer h.Effect_Free(snd)
		defer h.Texture_Free(quad.Texture)

		// Set up quad which we will use for rendering sprite
		quad.Blend = hge.BLEND_ALPHAADD | hge.BLEND_COLORMUL | hge.BLEND_ZWRITE

		for i := 0; i < 4; i++ {
			// Set up z-coordinate of vertices
			quad.V[i].Z = 0.5
			// Set up color. The format of DWORD col is 0xAARRGGBB
			quad.V[i].Color = 0xFFFFA000
		}

		// Set up quad's texture coordinates.
		// 0,0 means top left corner and 1,1 -
		// bottom right corner of the texture.
		quad.V[0].TX = 96.0 / 128.0
		quad.V[0].TY = 64.0 / 128.0
		quad.V[1].TX = 128.0 / 128.0
		quad.V[1].TY = 64.0 / 128.0
		quad.V[2].TX = 128.0 / 128.0
		quad.V[2].TY = 96.0 / 128.0
		quad.V[3].TX = 96.0 / 128.0
		quad.V[3].TY = 96.0 / 128.0

		line.X1 = 300
		line.Y1 = 300
		line.X2 = 400
		line.Y2 = 300
		line.Color = 0xFFFFA000

		// Let's rock now!
		h.System_Start()
	} else {
		fmt.Println("Error: %s\n", h.System_GetErrorMessage())
	}
}
