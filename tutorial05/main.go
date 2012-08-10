package main

import (
	"fmt"
	"hge"
	"math"
)

var (
	h   *hge.HGE
	tex hge.Texture
	dis hge.DistortionMesh
	fnt *hge.Font
)

const (
	rows  = 16
	cols  = 16
	cellw = 512.0 / (cols - 1)
	cellh = 512.0 / (rows - 1)

	meshx = 144.0
	meshy = 44.0
)

var trans = 0
var t = 0.0

func FrameFunc() int {
	t += h.Timer_GetDelta()

	// Process keys
	switch h.Input_GetKey() {
	case hge.K_ESCAPE:
		return 1

	case hge.K_SPACE:
		trans++

		if trans > 2 {
			trans = 0
		}

		dis.Clear(hge.Dword(0xFF000000))
	}

	// Calculate new displacements and coloring for one of the three effects
	switch trans {
	case 0:
		for i := 1; i < rows-1; i++ {
			for j := 1; j < cols-1; j++ {
				dis.SetDisplacement(j, i, math.Cos(t*float64(10+(i+j))/2)*5, math.Sin(t*float64(10+(i+j))/2)*5, hge.DISP_NODE)
			}
		}

	case 1:
		for i := 0; i < rows; i++ {
			for j := 1; j < cols-1; j++ {
				dis.SetDisplacement(j, i, math.Cos(t*float64(5+j)/2)*15, 0, hge.DISP_NODE)
				col := hge.Dword((math.Cos(t*float64(5+(i+j))/2) + 1) * 35)
				dis.SetColor(j, i, 0xFF<<24|col<<16|col<<8|col)
			}
		}

	case 2:
		for i := 0.0; i < rows; i++ {
			for j := 0.0; j < cols; j++ {
				r := math.Sqrt(math.Pow(j-float64(cols)/2, 2) + math.Pow(i-float64(rows)/2, 2))
				a := r * math.Cos(t*2) * 0.1
				dx := math.Sin(a)*(i*cellh-256) + math.Cos(a)*(j*cellw-256)
				dy := math.Cos(a)*(i*cellh-256) - math.Sin(a)*(j*cellw-256)
				dis.SetDisplacement(int(j), int(i), dx, dy, hge.DISP_CENTER)
				col := hge.Dword((math.Cos(r+t*4) + 1) * 40)
				dis.SetColor(int(j), int(i), 0xFF<<24|col<<16|(col/2)<<8)
			}
		}
	}

	return 0
}

func RenderFunc() int {
	// Render graphics
	h.Gfx_BeginScene()
	h.Gfx_Clear(0)
	dis.Render(meshx, meshy)
	fnt.Printf(5, 5, hge.TEXT_LEFT, "dt:%.3f\nFPS:%d\n\nUse your\nSPACE!", h.Timer_GetDelta(), h.Timer_GetFPS())
	h.Gfx_EndScene()

	return 0
}

func main() {
	h = hge.Create(hge.VERSION)
	defer h.Release()

	h.System_SetState(hge.LOGFILE, "tutorial05.log")
	h.System_SetState(hge.FRAMEFUNC, FrameFunc)
	h.System_SetState(hge.RENDERFUNC, RenderFunc)
	h.System_SetState(hge.TITLE, "HGE Tutorial 05 - Using distortion mesh")
	h.System_SetState(hge.WINDOWED, true)
	h.System_SetState(hge.SCREENWIDTH, 800)
	h.System_SetState(hge.SCREENHEIGHT, 600)
	h.System_SetState(hge.SCREENBPP, 32)
	h.System_SetState(hge.USESOUND, false)

	if h.System_Initiate() {
		defer h.System_Shutdown()
		tex = h.Texture_Load("texture.jpg")
		if tex == 0 {
			fmt.Println("Error: Can't load texture.jpg")
			return
		}
		defer h.Texture_Free(tex)

		dis = hge.NewDistortionMesh(cols, rows)
		dis.SetTexture(tex)
		dis.SetTextureRect(0, 0, 512, 512)
		dis.SetBlendMode(hge.BLEND_COLORADD | hge.BLEND_ALPHABLEND | hge.BLEND_ZWRITE)
		dis.Clear(hge.Dword(0xFF000000))

		// Load a font
		fnt = hge.NewFont("font1.fnt")

		if fnt == nil {
			fmt.Println("Error: Can't load font1.fnt or font1.png")
			return
		}

		// Let's rock now!
		h.System_Start()
	}
}
