package main

import (
	"fmt"
	HGE "hge"
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
	color        HGE.Dword
}

var (
	objs           [MAX_OBJECTS]Obj
	objects, blend int
	hge            *HGE.HGE

	// Resource handles
	tex, bgtex HGE.Texture
	spr, bgspr HGE.Sprite
	fnt        *HGE.Font
)

var (
	sprBlend = [5]int{HGE.BLEND_COLORMUL | HGE.BLEND_ALPHABLEND | HGE.BLEND_NOZWRITE,
		HGE.BLEND_COLORADD | HGE.BLEND_ALPHABLEND | HGE.BLEND_NOZWRITE,
		HGE.BLEND_COLORMUL | HGE.BLEND_ALPHABLEND | HGE.BLEND_NOZWRITE,
		HGE.BLEND_COLORMUL | HGE.BLEND_ALPHAADD | HGE.BLEND_NOZWRITE,
		HGE.BLEND_COLORMUL | HGE.BLEND_ALPHABLEND | HGE.BLEND_NOZWRITE}

	fntColor = [5]HGE.Dword{0xFFFFFFFF, 0xFF000000, 0xFFFFFFFF, 0xFF000000, 0xFFFFFFFF}

	sprColors = [5][5]HGE.Dword{{0xFFFFFFFF, 0xFFFFE080, 0xFF80A0FF, 0xFFA0FF80, 0xFFFF80A0},
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
		objs[i].color = sprColors[newBlend][hge.Random_Int(0, 4)]
	}
}

func frame() int {
	// 	float dt=hge.Timer_GetDelta();
	dt := hge.Timer_GetDelta()
	// 	int i;
	//
	// Process keys
	switch hge.Input_GetKey() {
	case HGE.K_UP:
		if objects < MAX_OBJECTS {
			objects += 100
		}
	case HGE.K_DOWN:
		if objects > MIN_OBJECTS {
			objects -= 100
		}
	case HGE.K_SPACE:
		blend++
		setBlend(blend)
	case HGE.K_ESCAPE:
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
	hge.Gfx_BeginScene()
	bgspr.Render(0, 0)

	for i := 0; i < objects; i++ {
		spr.SetColor(objs[i].color)
		spr.RenderEx(objs[i].x, objs[i].y, objs[i].rot, objs[i].scale)
	}

	fnt.Printf(7, 7, HGE.TEXT_LEFT, "UP and DOWN to adjust number of hares: %d\nSPACE to change blending mode: %d\nFPS: %d", objects, blend, hge.Timer_GetFPS())
	hge.Gfx_EndScene()

	return 0
}

func main() {
	hge = HGE.Create(HGE.VERSION)
	defer hge.Release()

	// Set desired system states and initialize HGE
	hge.System_SetState(HGE.LOGFILE, "tutorial07.log")
	hge.System_SetState(HGE.FRAMEFUNC, frame)
	hge.System_SetState(HGE.RENDERFUNC, render)
	hge.System_SetState(HGE.TITLE, "HGE Tutorial 07 - Thousand of Hares")
	hge.System_SetState(HGE.USESOUND, false)
	hge.System_SetState(HGE.WINDOWED, true)
	hge.System_SetState(HGE.SCREENWIDTH, SCREEN_WIDTH)
	hge.System_SetState(HGE.SCREENHEIGHT, SCREEN_HEIGHT)
	hge.System_SetState(HGE.SCREENBPP, 32)

	if hge.System_Initiate() {
		defer hge.System_Shutdown()
		// Load textures
		bgtex = hge.Texture_Load("bg2.png")
		tex = hge.Texture_Load("zazaka.png")
		if bgtex == 0 || tex == 0 {
			fmt.Println("Error: Can't load bg2.png or zazaka.png\n")
			return
		}
		// Delete created objects and free loaded resources
		defer hge.Texture_Free(tex)
		defer hge.Texture_Free(bgtex)

		// Load font, create sprites
		fnt = HGE.NewFont("font1.fnt")
		spr = HGE.NewSprite(tex, 0, 0, 64, 64)
		spr.SetHotSpot(32, 32)

		bgspr = HGE.NewSprite(bgtex, 0, 0, 800, 600)
		bgspr.SetBlendMode(HGE.BLEND_COLORADD | HGE.BLEND_ALPHABLEND | HGE.BLEND_NOZWRITE)
		bgspr.SetColor(0xFF000000, 0)
		bgspr.SetColor(0xFF000000, 1)
		bgspr.SetColor(0xFF000040, 2)
		bgspr.SetColor(0xFF000040, 3)

		// Initialize objects list
		objects = 1000

		for i := 0; i < MAX_OBJECTS; i++ {
			objs[i].x = hge.Random_Float(0, SCREEN_WIDTH)
			objs[i].y = hge.Random_Float(0, SCREEN_HEIGHT)
			objs[i].dx = hge.Random_Float(-200, 200)
			objs[i].dy = hge.Random_Float(-200, 200)
			objs[i].scale = hge.Random_Float(0.5, 2.0)
			objs[i].dscale = hge.Random_Float(-1.0, 1.0)
			objs[i].rot = hge.Random_Float(0, HGE.Pi*2)
			objs[i].drot = hge.Random_Float(-1.0, 1.0)
		}

		setBlend(0)

		// Let's rock now!
		hge.System_Start()
	}
}
