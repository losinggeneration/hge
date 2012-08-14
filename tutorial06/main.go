package main

import (
	"fmt"
	HGE "hge"
	"math"
)

var (
	hge  *HGE.HGE = HGE.Create(HGE.VERSION)
	gui  HGE.GUI
	fnt  *HGE.Font
	quad HGE.Quad
)

var lastId int = 0
var t float64 = 0.0

func frame() int {
	// 	float dt=hge.Timer_GetDelta();
	// 	static float t=0.0f;
	// 	float tx,ty;
	// 	int id;
	// 	static int lastid=0;

	dt := hge.Timer_GetDelta()

	// If ESCAPE was pressed, tell the GUI to finish
	if hge.Input_GetKeyState(HGE.K_ESCAPE) {
		lastId = 5
		gui.Leave()
	}

	// We update the GUI and take an action if
	// one of the menu items was selected
	id := gui.Update(dt)
	if id == -1 {
		switch lastId {
		case 1, 2, 3, 4:
			gui.SetFocus(1)
			gui.Enter()

		case 5:
			return 1
		}
	} else if id > 0 {
		lastId = id
		gui.Leave()
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
	hge.Gfx_BeginScene()
	hge.Gfx_RenderQuad(&quad)
	gui.Render()
	fnt.SetColor(0xFFFFFFFF)
	fnt.Printf(5, 5, HGE.TEXT_LEFT, "dt:%.3f\nFPS:%d", hge.Timer_GetDelta(), hge.Timer_GetFPS())
	hge.Gfx_EndScene()

	return 0
}

func main() {
	defer hge.Release()

	hge.System_SetState(HGE.LOGFILE, "tutorial06.log")
	hge.System_SetState(HGE.FRAMEFUNC, frame)
	hge.System_SetState(HGE.RENDERFUNC, render)
	hge.System_SetState(HGE.TITLE, "HGE Tutorial 06 - Creating menus")
	hge.System_SetState(HGE.WINDOWED, true)
	hge.System_SetState(HGE.SCREENWIDTH, 800)
	hge.System_SetState(HGE.SCREENHEIGHT, 600)
	hge.System_SetState(HGE.SCREENBPP, 32)

	if hge.System_Initiate() {
		defer hge.System_Shutdown()

		quad.Tex = hge.Texture_Load("bg.png")

		if quad.Tex == 0 {
			fmt.Println("Error loading bg.png")
			return
		}
		defer hge.Texture_Free(quad.Tex)

		snd := hge.Effect_Load("menu.ogg")

		if snd == 0 {
			fmt.Println("Error loading menu.ogg")
			return
		}
		defer hge.Effect_Free(snd)

		cursorTex := hge.Texture_Load("cursor.png")

		if cursorTex == 0 {
			fmt.Println("Error loading cursor.png")
			return
		}
		defer hge.Texture_Free(cursorTex)

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

		fnt = HGE.NewFont("font1.fnt")

		if fnt == nil {
			fmt.Println("Error loading font1.fnt")
			return
		}

		cursor := HGE.NewSprite(cursorTex, 0, 0, 32, 32)
		gui = HGE.NewGUI()

		gui.AddCtrl(NewGUIMenuItem(1, fnt, snd, 400, 200, 0.0, "Play"))
		gui.AddCtrl(NewGUIMenuItem(2, fnt, snd, 400, 240, 0.1, "Options"))
		gui.AddCtrl(NewGUIMenuItem(3, fnt, snd, 400, 280, 0.2, "Instructions"))
		gui.AddCtrl(NewGUIMenuItem(4, fnt, snd, 400, 320, 0.3, "Credits"))
		gui.AddCtrl(NewGUIMenuItem(5, fnt, snd, 400, 360, 0.4, "Exit"))

		gui.SetNavMode(HGE.GUI_UPDOWN | HGE.GUI_CYCLED)
		gui.SetCursor(&cursor)
		gui.SetFocus(1)
		gui.Enter()

		hge.System_Start()
	} else {
		fmt.Println("Error: ", hge.System_GetErrorMessage())
	}

}
