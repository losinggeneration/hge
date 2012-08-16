package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/helpers/font"
	"github.com/losinggeneration/hge-go/helpers/gui"
	"github.com/losinggeneration/hge-go/helpers/sprite"
	HGE "github.com/losinggeneration/hge-go/hge"
	hge "github.com/losinggeneration/hge-go/legacy"
	"math"
)

var (
	h    *hge.HGE = hge.Create(HGE.VERSION)
	GUI  gui.GUI
	fnt  *font.Font
	quad HGE.Quad
)

var lastId int = 0
var t float64 = 0.0

func frame() int {
	dt := h.Timer_GetDelta()

	// If ESCAPE was pressed, tell the GUI to finish
	if h.Input_GetKeyState(HGE.K_ESCAPE) {
		lastId = 5
		GUI.Leave()
	}

	// We update the GUI and take an action if
	// one of the menu items was selected
	id := GUI.Update(dt)
	if id == -1 {
		switch lastId {
		case 1, 2, 3, 4:
			GUI.SetFocus(1)
			GUI.Enter()

		case 5:
			return 1
		}
	} else if id > 0 {
		lastId = id
		GUI.Leave()
	}

	// Here we update our background animation
	t += dt
	tx := 50 * math.Cos(t/60)
	ty := 50 * math.Sin(t/60)

	quad.V[0].TX, quad.V[0].TY = float32(tx), float32(ty)
	quad.V[1].TX, quad.V[1].TY = float32(tx+800/64), float32(ty)
	quad.V[2].TX, quad.V[2].TY = float32(tx+800/64), float32(ty+600/64)
	quad.V[3].TX, quad.V[3].TY = float32(tx), float32(ty+600/64)

	return 0
}

func render() int {
	// Render graphics
	h.Gfx_BeginScene()
	h.Gfx_RenderQuad(&quad)
	GUI.Render()
	fnt.SetColor(0xFFFFFFFF)
	fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d", h.Timer_GetDelta(), h.Timer_GetFPS())
	h.Gfx_EndScene()

	return 0
}

func main() {
	defer h.Release()

	h.System_SetState(HGE.LOGFILE, "tutorial06.log")
	h.System_SetState(HGE.FRAMEFUNC, frame)
	h.System_SetState(HGE.RENDERFUNC, render)
	h.System_SetState(HGE.TITLE, "HGE Tutorial 06 - Creating menus")
	h.System_SetState(HGE.WINDOWED, true)
	h.System_SetState(HGE.SCREENWIDTH, 800)
	h.System_SetState(HGE.SCREENHEIGHT, 600)
	h.System_SetState(HGE.SCREENBPP, 32)

	if h.System_Initiate() {
		defer h.System_Shutdown()

		quad.Tex = h.Texture_Load("bg.png")

		if quad.Tex == 0 {
			fmt.Println("Error loading bg.png")
			return
		}
		defer h.Texture_Free(quad.Tex)

		snd := h.Effect_Load("menu.ogg")

		if snd == 0 {
			fmt.Println("Error loading menu.ogg")
			return
		}
		defer h.Effect_Free(snd)

		cursorTex := h.Texture_Load("cursor.png")

		if cursorTex == 0 {
			fmt.Println("Error loading cursor.png")
			return
		}
		defer h.Texture_Free(cursorTex)

		// Set up the quad we will use for background animation
		quad.Blend = HGE.BLEND_ALPHABLEND | HGE.BLEND_COLORMUL | HGE.BLEND_NOZWRITE

		for i := 0; i < 4; i++ {
			// Set up z-coordinate of vertices
			quad.V[i].Z = 0.5
			// Set up color. The format of DWORD col is 0xAARRGGBB
			quad.V[i].Col = 0xFFFFFFFF
		}

		quad.V[0].X, quad.V[0].Y = 0, 0
		quad.V[1].X, quad.V[1].Y = 800, 0
		quad.V[2].X, quad.V[2].Y = 800, 600
		quad.V[3].X, quad.V[3].Y = 0, 600

		fnt = font.NewFont("font1.fnt")

		if fnt == nil {
			fmt.Println("Error loading font1.fnt")
			return
		}

		cursor := sprite.NewSprite(cursorTex, 0, 0, 32, 32)
		GUI = gui.NewGUI()

		GUI.AddCtrl(NewGUIMenuItem(1, fnt, snd, 400, 200, 0.0, "Play"))
		GUI.AddCtrl(NewGUIMenuItem(2, fnt, snd, 400, 240, 0.1, "Options"))
		GUI.AddCtrl(NewGUIMenuItem(3, fnt, snd, 400, 280, 0.2, "Instructions"))
		GUI.AddCtrl(NewGUIMenuItem(4, fnt, snd, 400, 320, 0.3, "Credits"))
		GUI.AddCtrl(NewGUIMenuItem(5, fnt, snd, 400, 360, 0.4, "Exit"))

		GUI.SetNavMode(gui.GUI_UPDOWN | gui.GUI_CYCLED)
		GUI.SetCursor(&cursor)
		GUI.SetFocus(1)
		GUI.Enter()

		h.System_Start()
	} else {
		fmt.Println("Error: ", h.System_GetErrorMessage())
	}

}
