package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/helpers/font"
	"github.com/losinggeneration/hge-go/helpers/sprite"
	. "github.com/losinggeneration/hge-go/hge"
	. "github.com/losinggeneration/hge-go/hge/gfx"
	. "github.com/losinggeneration/hge-go/hge/input"
	hge "github.com/losinggeneration/hge-go/legacy"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600

	MIN_OBJECTS = 100
	MAX_OBJECTS = 2000
)

type Obj struct {
	x, y         float64
	dx, dy       float64
	scale, rot   float64
	dscale, drot float64
	color        Dword
}

var (
	objs           [MAX_OBJECTS]Obj
	objects, blend int
	h              *hge.HGE

	// Resource handles
	tex, bgtex *Texture
	spr, bgspr sprite.Sprite
	fnt        *font.Font
)

var (
	sprBlend = [5]int{BLEND_COLORMUL | BLEND_ALPHABLEND | BLEND_NOZWRITE,
		BLEND_COLORADD | BLEND_ALPHABLEND | BLEND_NOZWRITE,
		BLEND_COLORMUL | BLEND_ALPHABLEND | BLEND_NOZWRITE,
		BLEND_COLORMUL | BLEND_ALPHAADD | BLEND_NOZWRITE,
		BLEND_COLORMUL | BLEND_ALPHABLEND | BLEND_NOZWRITE}

	fntColor = [5]Dword{0xFFFFFFFF, 0xFF000000, 0xFFFFFFFF, 0xFF000000, 0xFFFFFFFF}

	sprColors = [5][5]Dword{{0xFFFFFFFF, 0xFFFFE080, 0xFF80A0FF, 0xFFA0FF80, 0xFFFF80A0},
		{0xFF000000, 0xFF303000, 0xFF000060, 0xFF006000, 0xFF600000},
		{0x80FFFFFF, 0x80FFE080, 0x8080A0FF, 0x80A0FF80, 0x80FF80A0},
		{0x80FFFFFF, 0x80FFE080, 0x8080A0FF, 0x80A0FF80, 0x80FF80A0},
		{0x40202020, 0x40302010, 0x40102030, 0x40203010, 0x40102030}}
)

// Set up blending mode for the scene
func setBlend(newBlend int) {
	if newBlend > 4 {
		newBlend = 0
	}
	blend = newBlend

	spr.SetBlendMode(sprBlend[newBlend])
	fnt.SetColor(fntColor[newBlend])
	for i := 0; i < MAX_OBJECTS; i++ {
		objs[i].color = sprColors[newBlend][h.Random_Int(0, 4)]
	}
}

func frame() int {
	// 	float dt=h.Timer_GetDelta();
	dt := h.Timer_GetDelta()
	// 	int i;
	//
	// Process keys
	switch h.Input_GetKey() {
	case K_UP:
		if objects < MAX_OBJECTS {
			objects += 100
		}
	case K_DOWN:
		if objects > MIN_OBJECTS {
			objects -= 100
		}
	case K_SPACE:
		blend++
		setBlend(blend)
	case K_ESCAPE:
		return 1
	}

	// Update the scene
	for i := 0; i < objects; i++ {
		objs[i].x += objs[i].dx * dt
		if objs[i].x > SCREEN_WIDTH || objs[i].x < 0 {
			objs[i].dx = -objs[i].dx
		}
		objs[i].y += objs[i].dy * dt
		if objs[i].y > SCREEN_HEIGHT || objs[i].y < 0 {
			objs[i].dy = -objs[i].dy
		}
		objs[i].scale += objs[i].dscale * dt
		if objs[i].scale > 2 || objs[i].scale < 0.5 {
			objs[i].dscale = -objs[i].dscale
		}
		objs[i].rot += objs[i].drot * dt
	}

	return 0
}

func render() int {
	// Render the scene
	h.Gfx_BeginScene()
	bgspr.Render(0, 0)

	for i := 0; i < objects; i++ {
		spr.SetColor(objs[i].color)
		spr.RenderEx(objs[i].x, objs[i].y, objs[i].rot, objs[i].scale)
	}

	fnt.Printf(7, 7, font.TEXT_LEFT, "UP and DOWN to adjust number of hares: %d\nSPACE to change blending mode: %d\nFPS: %d", objects, blend, h.Timer_GetFPS())
	h.Gfx_EndScene()

	return 0
}

func main() {
	h = hge.Create(VERSION)
	defer h.Release()

	// Set desired system states and initialize HGE
	h.System_SetState(LOGFILE, "tutorial07.log")
	h.System_SetState(FRAMEFUNC, frame)
	h.System_SetState(RENDERFUNC, render)
	h.System_SetState(TITLE, "HGE Tutorial 07 - Thousand of Hares")
	h.System_SetState(USESOUND, false)
	h.System_SetState(WINDOWED, true)
	h.System_SetState(SCREENWIDTH, SCREEN_WIDTH)
	h.System_SetState(SCREENHEIGHT, SCREEN_HEIGHT)
	h.System_SetState(SCREENBPP, 32)

	if h.System_Initiate() {
		defer h.System_Shutdown()
		// Load textures
		bgtex = h.Texture_Load("bg2.png")
		tex = h.Texture_Load("zazaka.png")
		if bgtex == nil || tex == nil {
			fmt.Println("Error: Can't load bg2.png or zazaka.png\n")
			return
		}
		// Delete created objects and free loaded resources
		defer h.Texture_Free(tex)
		defer h.Texture_Free(bgtex)

		// Load font, create sprites
		fnt = font.NewFont("font1.fnt")
		spr = sprite.NewSprite(tex, 0, 0, 64, 64)
		spr.SetHotSpot(32, 32)

		bgspr = sprite.NewSprite(bgtex, 0, 0, 800, 600)
		bgspr.SetBlendMode(BLEND_COLORADD | BLEND_ALPHABLEND | BLEND_NOZWRITE)
		bgspr.SetColor(0xFF000000, 0)
		bgspr.SetColor(0xFF000000, 1)
		bgspr.SetColor(0xFF000040, 2)
		bgspr.SetColor(0xFF000040, 3)

		// Initialize objects list
		objects = 1000

		for i := 0; i < MAX_OBJECTS; i++ {
			objs[i].x = h.Random_Float(0, SCREEN_WIDTH)
			objs[i].y = h.Random_Float(0, SCREEN_HEIGHT)
			objs[i].dx = h.Random_Float(-200, 200)
			objs[i].dy = h.Random_Float(-200, 200)
			objs[i].scale = h.Random_Float(0.5, 2.0)
			objs[i].dscale = h.Random_Float(-1.0, 1.0)
			objs[i].rot = h.Random_Float(0, Pi*2)
			objs[i].drot = h.Random_Float(-1.0, 1.0)
		}

		setBlend(0)

		// Let's rock now!
		h.System_Start()
	}
}
