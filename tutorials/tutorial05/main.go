package main

import (
	"fmt"
	"math"

	"github.com/losinggeneration/hge"
	"github.com/losinggeneration/hge/gfx"
	dist "github.com/losinggeneration/hge/helpers/distortionmesh"
	"github.com/losinggeneration/hge/helpers/font"
	"github.com/losinggeneration/hge/input"
	"github.com/losinggeneration/hge/timer"
)

var (
	tex *gfx.Texture
	dis dist.DistortionMesh
	fnt *font.Font
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
	t += timer.Delta()

	// Process keys
	switch input.GetKey() {
	case input.K_ESCAPE:
		return 1

	case input.K_SPACE:
		trans++

		if trans > 2 {
			trans = 0
		}

		dis.Clear(gfx.RGBAToColor(0xFF000000))
	}

	// Calculate new displacements and coloring for one of the three effects
	switch trans {
	case 0:
		for i := 1; i < rows-1; i++ {
			for j := 1; j < cols-1; j++ {
				dis.SetDisplacement(j, i, math.Cos(t*float64(10+(i+j))/2)*5, math.Sin(t*float64(10+(i+j))/2)*5, dist.DISP_NODE)
			}
		}

	case 1:
		for i := 0; i < rows; i++ {
			for j := 1; j < cols-1; j++ {
				dis.SetDisplacement(j, i, math.Cos(t*float64(5+j)/2)*15, 0, dist.DISP_NODE)
				col := uint32((math.Cos(t*float64(5+(i+j))/2) + 1) * 35)
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
				dis.SetDisplacement(int(j), int(i), dx, dy, dist.DISP_CENTER)
				col := uint32((math.Cos(r+t*4) + 1) * 40)
				dis.SetColor(int(j), int(i), 0xFF<<24|col<<16|(col/2)<<8)
			}
		}
	}

	return 0
}

func RenderFunc() int {
	// Render graphics
	gfx.BeginScene()
	gfx.Clear(gfx.RGBAToColor(0))
	dis.Render(meshx, meshy)
	fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d\n\nUse your\nSPACE!", timer.Delta(), timer.FPS())
	gfx.EndScene()

	return 0
}

func main() {
	h := hge.New()

	h.SetState(hge.LOGFILE, "tutorial05.log")
	h.SetState(hge.FRAMEFUNC, FrameFunc)
	h.SetState(hge.RENDERFUNC, RenderFunc)
	h.SetState(hge.TITLE, "HGE Tutorial 05 - Using distortion mesh")
	h.SetState(hge.WINDOWED, true)
	h.SetState(hge.SCREENWIDTH, 800)
	h.SetState(hge.SCREENHEIGHT, 600)
	h.SetState(hge.SCREENBPP, 32)
	h.SetState(hge.USESOUND, false)

	if err := h.Initiate(); err == nil {
		defer h.Shutdown()
		tex, err = gfx.LoadTexture("texture.jpg")
		if tex == nil || err != nil {
			fmt.Println("Error: Can't load texture.jpg")
			return
		}

		dis = dist.New(cols, rows)
		dis.SetTexture(tex)
		dis.SetTextureRect(0, 0, 512, 512)
		dis.SetBlendMode(gfx.BLEND_COLORADD | gfx.BLEND_ALPHABLEND | gfx.BLEND_ZWRITE)
		dis.Clear(gfx.RGBAToColor(0xFF000000))

		// Load a font
		fnt = font.New("font1.fnt")

		if fnt == nil {
			fmt.Println("Error: Can't load font1.fnt or font1.png")
			return
		}

		// Let's rock now!
		h.Start()
	}
}
