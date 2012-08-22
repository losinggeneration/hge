package main

import (
	Color "github.com/losinggeneration/hge-go/helpers/color"
	Font "github.com/losinggeneration/hge-go/helpers/font"
	"github.com/losinggeneration/hge-go/helpers/gui"
	. "github.com/losinggeneration/hge-go/hge/input"
	. "github.com/losinggeneration/hge-go/hge/sound"
)

type GUIMenuItem struct {
	gui.GUIObject

	font                                                              *Font.Font
	snd                                                               *Effect
	delay                                                             float64
	title                                                             string
	scolor, dcolor, scolor2, dcolor2, sshadow, dshadow, color, shadow Color.ColorRGB
	soffset, doffset, offset                                          float64
	timer, timer2                                                     float64
}

func NewGUIMenuItem(id int, font *Font.Font, snd *Effect, x, y, delay float64, title string) *gui.GUIObject {
	mi := new(GUIMenuItem)

	mi.GUIObject.Initialize()

	mi.GUIObject.Id = id
	mi.font = font
	mi.snd = snd
	mi.delay = delay
	mi.title = title

	mi.color.SetHWColor(0xFFFFE060)
	mi.shadow.SetHWColor(0x30000000)
	mi.offset = 0.0
	mi.timer = -1.0
	mi.timer2 = -1.0

	mi.GUIObject.Static = false
	mi.GUIObject.Visible = true
	mi.GUIObject.Enabled = true

	w := font.GetStringWidth(title)
	mi.GUIObject.Rect.Set(x-w/2, y, x+w/2, y+mi.font.GetHeight())

	mi.GUIObject.Render = func() {
		mi.font.SetColor(mi.shadow.HWColor())
		mi.font.Render(mi.GUIObject.Rect.X1+mi.offset+3, mi.GUIObject.Rect.Y1+3, Font.TEXT_LEFT, mi.title)
		mi.font.SetColor(mi.color.HWColor())
		mi.font.Render(mi.GUIObject.Rect.X1-mi.offset, mi.GUIObject.Rect.Y1-mi.offset, Font.TEXT_LEFT, mi.title)
	}

	mi.GUIObject.Update = func(dt float64) {
		if mi.timer2 != -1.0 {
			mi.timer2 += dt
			if mi.timer2 >= mi.delay+0.1 {
				mi.color = mi.scolor2.Add(mi.dcolor2)
				mi.shadow = mi.sshadow.Add(mi.dshadow)
				mi.offset = 0.0
				mi.timer2 = -1.0
			} else {
				if mi.timer2 < mi.delay {
					mi.color = mi.scolor2
					mi.shadow = mi.sshadow
				} else {
					mi.color = mi.scolor2.Add(mi.dcolor2.MulScalar((mi.timer2 - mi.delay) * 10))
					mi.shadow = mi.sshadow.Add(mi.dshadow.MulScalar((mi.timer2 - mi.delay) * 10))
				}
			}
		} else if mi.timer != -1.0 {
			mi.timer += dt
			if mi.timer >= 0.2 {
				mi.color = mi.scolor.Add(mi.dcolor)
				mi.offset = mi.soffset + mi.doffset
				mi.timer = -1.0
			} else {
				mi.color = mi.scolor.Add(mi.dcolor.MulScalar(mi.timer * 5))
				mi.offset = mi.soffset + mi.doffset*mi.timer*5
			}
		}
	}

	mi.GUIObject.Enter = func() {
		var tcolor2 Color.ColorRGB

		mi.scolor2.SetHWColor(0x00FFE060)
		tcolor2.SetHWColor(0xFFFFE060)
		mi.dcolor2 = tcolor2.Sub(mi.scolor2)

		mi.sshadow.SetHWColor(0x00000000)
		tcolor2.SetHWColor(0x30000000)
		mi.dshadow = tcolor2.Sub(mi.sshadow)

		mi.timer2 = 0.0
	}

	mi.GUIObject.Leave = func() {
		var tcolor2 Color.ColorRGB

		mi.scolor2.SetHWColor(0xFFFFE060)
		tcolor2.SetHWColor(0x00FFE060)
		mi.dcolor2 = tcolor2.Sub(mi.scolor2)

		mi.sshadow.SetHWColor(0x30000000)
		tcolor2.SetHWColor(0x00000000)
		mi.dshadow = tcolor2.Sub(mi.sshadow)

		mi.timer2 = 0.0
	}

	mi.GUIObject.IsDone = func() bool {
		if mi.timer2 == -1.0 {
			return true
		}

		return false
	}

	mi.GUIObject.Focus = func(focused bool) {
		var tcolor Color.ColorRGB

		if focused {
			snd.Play()
			mi.scolor.SetHWColor(0xFFFFE060)
			tcolor.SetHWColor(0xFFFFFFFF)
			mi.soffset = 0
			mi.doffset = 4
		} else {
			mi.scolor.SetHWColor(0xFFFFFFFF)
			tcolor.SetHWColor(0xFFFFE060)
			mi.soffset = 4
			mi.doffset = -4
		}

		mi.dcolor = tcolor.Sub(mi.scolor)
		mi.timer = 0.0
	}

	mi.GUIObject.MouseOver = func(over bool) {
		if over {
			mi.GUIObject.GUI.SetFocus(mi.GUIObject.Id)
		}
	}

	mi.GUIObject.MouseLButton = func(down bool) bool {
		if !down {
			mi.offset = 4
			return true
		}

		snd.Play()
		mi.offset = 0
		return false
	}

	mi.GUIObject.KeyClick = func(key Key, chr int) bool {
		if key == K_ENTER || key == K_SPACE {
			mi.GUIObject.MouseLButton(true)
			return mi.GUIObject.MouseLButton(false)
		}

		return false
	}

	return &mi.GUIObject
}
