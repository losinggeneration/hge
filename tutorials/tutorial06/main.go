package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/helpers/font"
	"github.com/losinggeneration/hge-go/helpers/gui"
	"github.com/losinggeneration/hge-go/helpers/sprite"
	"github.com/losinggeneration/hge-go/hge"
	. "github.com/losinggeneration/hge-go/hge/gfx"
	. "github.com/losinggeneration/hge-go/hge/input"
	. "github.com/losinggeneration/hge-go/hge/sound"
	. "github.com/losinggeneration/hge-go/hge/timer"
	"math"
)

var (
	GUI  gui.GUI
	fnt  *font.Font
	quad Quad
)

var lastId int = 0
var t float64 = 0.0

func frame() int {
	dt := Delta()

	// If ESCAPE was pressed, tell the GUI to finish
	if NewKey(K_ESCAPE).State() {
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
	BeginScene()
	quad.Render()
	GUI.Render()
	fnt.SetColor(0xFFFFFFFF)
	fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d", Delta(), GetFPS())
	EndScene()

	return 0
}

func main() {
	h := hge.New()
	h.SetState(hge.LOGFILE, "tutorial06.log")
	h.SetState(hge.FRAMEFUNC, frame)
	h.SetState(hge.RENDERFUNC, render)
	h.SetState(hge.TITLE, "HGE Tutorial 06 - Creating menus")
	h.SetState(hge.WINDOWED, true)
	h.SetState(hge.SCREENWIDTH, 800)
	h.SetState(hge.SCREENHEIGHT, 600)
	h.SetState(hge.SCREENBPP, 32)

	if err := h.Initiate(); err != nil {
		fmt.Println("Error: ", h.GetErrorMessage())
	} else {
		defer h.Shutdown()

		quad.Texture = LoadTexture("bg.png")

		if quad.Texture == nil {
			fmt.Println("Error loading bg.png")
			return
		}

		snd := NewEffect("menu.ogg")

		if snd == nil {
			fmt.Println("Error loading menu.ogg")
			return
		}

		cursorTex := LoadTexture("cursor.png")

		if cursorTex == nil {
			fmt.Println("Error loading cursor.png")
			return
		}

		// Set up the quad we will use for background animation
		quad.Blend = BLEND_ALPHABLEND | BLEND_COLORMUL | BLEND_NOZWRITE

		for i := 0; i < 4; i++ {
			// Set up z-coordinate of vertices
			quad.V[i].Z = 0.5
			// Set up color. The format of DWORD col is 0xAARRGGBB
			quad.V[i].Color = 0xFFFFFFFF
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

		h.Start()
	}

}
