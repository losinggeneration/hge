package main

import (
	"fmt"
	"github.com/losinggeneration/hge-go/binding/helpers/font"
	"github.com/losinggeneration/hge-go/binding/helpers/particle"
	"github.com/losinggeneration/hge-go/binding/helpers/sprite"
	. "github.com/losinggeneration/hge-go/binding/hge"
	. "github.com/losinggeneration/hge-go/binding/hge/gfx"
	. "github.com/losinggeneration/hge-go/binding/hge/input"
	. "github.com/losinggeneration/hge-go/binding/hge/sound"
	. "github.com/losinggeneration/hge-go/binding/hge/timer"
	"runtime"
	"time"
)

// An example of using closures
func main() {
	func() {
		const (
			speed    = 90.0
			friction = 0.98
		)

		var (
			spr sprite.Sprite
			fnt *font.Font
			par *particle.ParticleSystem

			snd *Effect

			x  = 100.0
			y  = 100.0
			dx = 0.0
			dy = 0.0
		)

		hge := New()

		hge.SetState(LOGFILE, "tutorial03.log")
		hge.SetState(TITLE, "HGE Tutorial 03 - Using helper classes")
		hge.SetState(FPS, 100)
		hge.SetState(WINDOWED, true)
		hge.SetState(SCREENWIDTH, 800)
		hge.SetState(SCREENHEIGHT, 600)
		hge.SetState(SCREENBPP, 32)
		hge.SetState(FRAMEFUNC, func() int {
			dt := float64(Delta())

			boom := func() {
				pan := int((x - 400) / 4)
				pitch := (dx*dx+dy*dy)*0.0005 + 0.2
				snd.PlayEx(100, pan, pitch)
			}

			// Process keys
			if NewKey(K_ESCAPE).State() {
				return 1
			}
			if NewKey(K_LEFT).State() {
				dx -= speed * dt
			}
			if NewKey(K_RIGHT).State() {
				dx += speed * dt
			}
			if NewKey(K_UP).State() {
				dy -= speed * dt
			}
			if NewKey(K_DOWN).State() {
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

			// Update particle system
			par.Info.Emission = (int)(dx*dx+dy*dy) * 2
			par.MoveTo(x, y)
			par.Update(dt)

			return 0
		})

		hge.SetState(RENDERFUNC, func() int {
			BeginScene()
			Clear(0)
			par.Render()
			spr.Render(x, y)
			fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d (constant)", Delta(), GetFPS())
			EndScene()

			return 0
		})

		if err := hge.Initiate(); err == nil {
			defer hge.Shutdown()

			snd = NewEffect("menu.ogg")
			tex := LoadTexture("particles.png")
			if snd == nil || tex == nil {
				fmt.Printf("Error: Can't load one of the following files:\nmenu.ogg, particles.png, font1.fnt, font1.png, trail.psi\n")
				return
			}

			spr = sprite.New(tex, 96, 64, 32, 32)
			spr.SetColor(0xFFFFA000)
			spr.SetHotSpot(16, 16)

			if fnt = font.New("font1.fnt"); fnt == nil {
				fmt.Println("Error loading font1.fnt")
				return
			}

			spt := sprite.New(tex, 32, 32, 32, 32)
			spt.SetBlendMode(BLEND_COLORMUL | BLEND_ALPHAADD | BLEND_NOZWRITE)
			spt.SetHotSpot(16, 16)

			par = particle.New("trail.psi", spt)
			if par == nil {
				fmt.Println("Error loading trail.psi")
				return
			}
			par.Fire()

			hge.Start()
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Sleeping")
		time.Sleep(1 * time.Second)
		runtime.GC()
	}
}
