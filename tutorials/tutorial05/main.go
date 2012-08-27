package main

import (
	"fmt"
	dist "github.com/losinggeneration/hge-go/helpers/distortionmesh"
	"github.com/losinggeneration/hge-go/helpers/font"
	. "github.com/losinggeneration/hge-go/hge"
	. "github.com/losinggeneration/hge-go/hge/gfx"
	. "github.com/losinggeneration/hge-go/hge/input"
	. "github.com/losinggeneration/hge-go/hge/timer"
	"math"
)

var (
	tex *Texture
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
	t += Delta()

	// Process keys
	switch GetKey() {
	case K_ESCAPE:
		return 1

	case K_SPACE:
		trans++

		if trans > 2 {
			trans = 0
		}

		dis.Clear(Dword(0xFF000000))
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
				col := Dword((math.Cos(t*float64(5+(i+j))/2) + 1) * 35)
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
				col := Dword((math.Cos(r+t*4) + 1) * 40)
				dis.SetColor(int(j), int(i), 0xFF<<24|col<<16|(col/2)<<8)
			}
		}
	}

	return 0
}

func RenderFunc() int {
	// Render graphics
	BeginScene()
	Clear(0)
	dis.Render(meshx, meshy)
	fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d\n\nUse your\nSPACE!", Delta(), GetFPS())
	EndScene()

	return 0
}

func main() {
	h := New()
	defer h.Free()

	h.SetState(LOGFILE, "tutorial05.log")
	h.SetState(FRAMEFUNC, FrameFunc)
	h.SetState(RENDERFUNC, RenderFunc)
	h.SetState(TITLE, "HGE Tutorial 05 - Using distortion mesh")
	h.SetState(WINDOWED, true)
	h.SetState(SCREENWIDTH, 800)
	h.SetState(SCREENHEIGHT, 600)
	h.SetState(SCREENBPP, 32)
	h.SetState(USESOUND, false)

	if err := h.Initiate(); err == nil {
		defer h.Shutdown()
		tex = LoadTexture("texture.jpg")
		if tex == nil {
			fmt.Println("Error: Can't load texture.jpg")
			return
		}
		defer tex.Free()

		dis = dist.New(cols, rows)
		dis.SetTexture(tex)
		dis.SetTextureRect(0, 0, 512, 512)
		dis.SetBlendMode(BLEND_COLORADD | BLEND_ALPHABLEND | BLEND_ZWRITE)
		dis.Clear(Dword(0xFF000000))

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
