/*
 * Haaf's Game Engine 1.8
 * Copyright (C) 2003-2007, Relish Games
 * hge.relishgames.com
 *
 * hge_tut02 - Using input, sound and rendering
 */

package main

import (
	"fmt"
	"hge"
)

var (
	h *hge.HGE

	quad hge.Quad

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

	// Set up quad's screen coordinates
	quad.V[0].X = x - 16
	quad.V[0].Y = y - 16
	quad.V[1].X = x + 16
	quad.V[1].Y = y - 16
	quad.V[2].X = x + 16
	quad.V[2].Y = y + 16
	quad.V[3].X = x - 16
	quad.V[3].Y = y + 16

	// Continue execution
	return 0
}

// This function will be called by HGE when
// the application window should be redrawn.
// Put your rendering code here.
func RenderFunc() int {
	// Begin rendering quads.
	// This function must be called
	// before any actual rendering.
	h.Gfx_BeginScene()

	// Clear screen with black color
	h.Gfx_Clear(0)

	// Render quads here. This time just
	// one of them will serve our needs.
	h.Gfx_RenderQuad(&quad)

	// End rendering and update the screen
	h.Gfx_EndScene()

	// RenderFunc should always return false
	return 0
}

func main() {
	// Get HGE interface
	h = hge.Create(hge.VERSION)

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
		// Load sound and texture
		snd = h.Effect_Load("menu.ogg")
		quad.Tex = h.Texture_Load("particles.png")
		if snd == 0 || quad.Tex == 0 {
			// If one of the data files is not found, display
			// an error message and shutdown.
			fmt.Println("Error: Can't load menu.ogg or particles.png")
			h.System_Shutdown()
			h.Release()
			return
		}

		// Set up quad which we will use for rendering sprite
		quad.Blend = hge.BLEND_ALPHAADD | hge.BLEND_COLORMUL | hge.BLEND_ZWRITE

		for i := 0; i < 4; i++ {
			// Set up z-coordinate of vertices
			quad.V[i].Z = 0.5
			// Set up color. The format of DWORD col is 0xAARRGGBB
			quad.V[i].Col = 0xFFFFA000
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

		// Let's rock now!
		h.System_Start()

		// Free loaded texture and sound
		h.Texture_Free(quad.Tex)
		h.Effect_Free(snd)
	} else {
		fmt.Println("Error: %s\n", h.System_GetErrorMessage())
	}

	// Clean up and shutdown
	h.System_Shutdown()
	h.Release()
}
