package guictrls

import (
	"container/list"
	"fmt"
	. "github.com/losinggeneration/hge-go/helpers/font"
	. "github.com/losinggeneration/hge-go/helpers/gui"
	. "github.com/losinggeneration/hge-go/helpers/sprite"
	. "github.com/losinggeneration/hge-go/hge"
	. "github.com/losinggeneration/hge-go/hge/gfx"
	. "github.com/losinggeneration/hge-go/hge/input"
)

type GUIText struct {
	GUIObject

	font   *Font
	tx, ty float64
	align  int
	text   string
}

func NewGUIText(id int, x, y, w, h float64, fnt *Font) *GUIText {
	t := new(GUIText)

	t.GUIObject.Initialize()

	t.GUIObject.Id = id
	t.GUIObject.Static = true
	t.GUIObject.Visible = true
	t.GUIObject.Enabled = true
	t.GUIObject.Rect.Set(x, y, x+w, y+h)

	t.font = fnt
	t.tx = x
	t.ty = y + (h-t.font.GetHeight())/2.0

	t.GUIObject.Render = func() {
		t.font.SetColor(t.GUIObject.Color)
		t.font.Render(t.tx, t.ty, t.align, t.text)
	}

	return t
}

func (t *GUIText) SetMode(align int) {
	t.align = align
}

func (t *GUIText) SetText(text string) {
	t.text = text
}

func (t *GUIText) Printf(format string, a ...interface{}) {
	t.text = fmt.Sprintf(format, a...)
}

type GUIButton struct {
	GUIObject

	trigger, pressed, oldState bool
	up, down                   Sprite
}

func NewGUIButton(id int, x, y, w, h float64, tex *Texture, tx, ty float64) *GUIButton {
	b := new(GUIButton)

	b.GUIObject.Initialize()

	b.GUIObject.Id = id
	b.GUIObject.Visible = true
	b.GUIObject.Enabled = true
	b.GUIObject.Rect.Set(x, y, x+w, y+h)

	b.up = NewSprite(tex, tx, ty, w, h)
	b.down = NewSprite(tex, tx+w, ty, w, h)

	b.GUIObject.Render = func() {
		if b.pressed {
			b.down.Render(b.GUIObject.Rect.X1, b.GUIObject.Rect.Y1)
		} else {
			b.up.Render(b.GUIObject.Rect.X1, b.GUIObject.Rect.Y1)
		}
	}

	b.GUIObject.MouseLButton = func(down bool) bool {
		if down {
			b.oldState = b.pressed
			b.pressed = true
			return false
		}

		if b.trigger {
			b.pressed = !b.oldState
		} else {
			b.pressed = false
		}

		return true
	}

	return b
}

func (b *GUIButton) SetMode(trigger bool) {
	b.trigger = trigger
}

func (b *GUIButton) SetState(pressed bool) {
	b.pressed = pressed
}

func (b GUIButton) State() bool {
	return b.pressed
}

const (
	SLIDER_BAR         = 0
	SLIDER_BARRELATIVE = 1
	SLIDER_SLIDER      = 2
)

type GUISlider struct {
	GUIObject
	pressed, vertical bool
	mode              int
	min, max, val     float64
	sl_w, sl_h        float64
	Sprite
}

func NewGUISlider(id int, x, y, w, h float64, tex *Texture, tx, ty, sw, sh float64, a ...interface{}) *GUISlider {
	s := new(GUISlider)

	if len(a) == 1 {
		if b, ok := a[0].(bool); ok {
			s.vertical = b
		}
	}

	s.GUIObject.Initialize()

	s.GUIObject.Id = id
	s.GUIObject.Visible = true
	s.GUIObject.Enabled = true
	s.GUIObject.Rect.Set(x, y, x+w, y+h)

	s.mode = SLIDER_BAR
	s.max, s.val = 100, 50
	s.sl_w, s.sl_h = sw, sh

	s.Sprite = NewSprite(tex, tx, ty, sw, sh)

	s.GUIObject.Render = func() {
		var x1, y1, x2, y2 float64

		xx := s.GUIObject.Rect.X1 + (s.GUIObject.Rect.X2-s.GUIObject.Rect.X1)*(s.val-s.min)/(s.max-s.min)
		yy := s.GUIObject.Rect.Y1 + (s.GUIObject.Rect.Y2-s.GUIObject.Rect.Y1)*(s.val-s.min)/(s.max-s.min)

		if s.vertical {
			switch s.mode {
			case SLIDER_BAR:
				x1 = s.GUIObject.Rect.X1
				y1 = s.GUIObject.Rect.Y1
				x2 = s.GUIObject.Rect.X2
				y2 = y

			case SLIDER_BARRELATIVE:
				x1 = s.GUIObject.Rect.X1
				y1 = (s.GUIObject.Rect.Y1 + s.GUIObject.Rect.Y2) / 2
				x2 = s.GUIObject.Rect.X2
				y2 = y

			case SLIDER_SLIDER:
				x1 = (s.GUIObject.Rect.X1 + s.GUIObject.Rect.X2 - s.sl_w) / 2
				y1 = yy - s.sl_h/2
				x2 = (s.GUIObject.Rect.X1 + s.GUIObject.Rect.X2 + s.sl_w) / 2
				y2 = yy + s.sl_h/2

			}
		} else {
			switch s.mode {
			case SLIDER_BAR:
				x1 = s.GUIObject.Rect.X1
				y1 = s.GUIObject.Rect.Y1
				x2 = xx
				y2 = s.GUIObject.Rect.Y2

			case SLIDER_BARRELATIVE:
				x1 = (s.GUIObject.Rect.X1 + s.GUIObject.Rect.X2) / 2
				y1 = s.GUIObject.Rect.Y1
				x2 = xx
				y2 = s.GUIObject.Rect.Y2

			case SLIDER_SLIDER:
				x1 = xx - s.sl_w/2
				y1 = (s.GUIObject.Rect.Y1 + s.GUIObject.Rect.Y2 - s.sl_h) / 2
				x2 = xx + s.sl_w/2
				y2 = (s.GUIObject.Rect.Y1 + s.GUIObject.Rect.Y2 + s.sl_h) / 2
			}
		}

		s.Sprite.RenderStretch(x1, y1, x2, y2)
	}

	s.GUIObject.MouseMove = func(x, y float64) bool {
		if s.pressed {
			if s.vertical {
				if y > s.GUIObject.Rect.Y2-s.GUIObject.Rect.Y1 {
					y = s.GUIObject.Rect.Y2 - s.GUIObject.Rect.Y1
				}
				if y < 0 {
					y = 0
				}
				s.val = s.min + (s.max-s.min)*y/(s.GUIObject.Rect.Y2-s.GUIObject.Rect.Y1)
			} else {
				if x > s.GUIObject.Rect.X2-s.GUIObject.Rect.X1 {
					x = s.GUIObject.Rect.X2 - s.GUIObject.Rect.X1
				}
				if x < 0 {
					x = 0
				}
				s.val = s.min + (s.max-s.min)*x/(s.GUIObject.Rect.X2-s.GUIObject.Rect.X1)
			}
			return true
		}

		return false
	}

	s.GUIObject.MouseLButton = func(down bool) bool {
		s.pressed = down
		return false
	}

	return s
}

func (s *GUISlider) SetMode(min, max float64, mode int) {
	s.min, s.max, s.mode = min, max, mode
}

func (s *GUISlider) SetValue(val float64) {
	if val < s.min {
		s.val = s.min
	} else if val > s.max {
		s.val = s.max
	} else {
		s.val = val
	}
}

func (s *GUISlider) Value() float64 {
	return s.val
}

type guiListboxItem struct {
	text string
}

type GUIListBox struct {
	GUIObject

	highlight Sprite
	*Font
	color, highlightColor        Dword
	items, selectedItem, topItem int
	mx, my                       float64
	*list.List
}

func NewGUIListBox(id int, x, y, w, h float64, font *Font, color, highlightColor, spriteHighlightColor Dword) *GUIListBox {
	l := new(GUIListBox)

	l.GUIObject.Initialize()

	l.GUIObject.Id = id
	l.GUIObject.Visible = true
	l.GUIObject.Enabled = true
	l.GUIObject.Rect.Set(x, y, x+w, y+h)
	l.Font = font
	l.highlight = NewSprite(nil, 0, 0, w, font.GetHeight())
	l.highlight.SetColor(highlightColor)
	l.color = color
	l.highlightColor = highlightColor
	l.List = list.New()

	l.GUIObject.Render = func() {
		item := l.List.Front()

		for i := 0; i < l.topItem; i++ {
			if item == nil {
				return
			}

			item = item.Next()
		}

		for i := 0; i < l.NumRows(); i++ {
			if i >= l.items {
				return
			}
			if l.topItem+i == l.selectedItem {
				l.highlight.Render(l.GUIObject.Rect.X1, l.GUIObject.Rect.Y1+float64(i)*l.Font.GetHeight())
				l.Font.SetColor(l.highlightColor)
			} else {
				l.Font.SetColor(l.color)
			}

			l.Font.Render(l.GUIObject.Rect.X1+3, l.GUIObject.Rect.Y1+float64(i)*l.Font.GetHeight(), TEXT_LEFT, item.Value.(*guiListboxItem).text)

			item = item.Next()
		}
	}

	l.GUIObject.MouseMove = func(x, y float64) bool {
		l.mx, l.my = x, y
		return false
	}

	l.GUIObject.MouseLButton = func(down bool) bool {
		if down {
			item := l.topItem + int(l.my)/int(l.Font.GetHeight())
			if item < l.Len() {
				l.selectedItem = item
				return true
			}
		}
		return false
	}

	l.GUIObject.MouseWheel = func(notches int) bool {
		l.topItem -= notches
		if l.topItem < 0 {
			l.topItem = 0
		}
		if l.topItem > l.Len()-l.NumRows() {
			l.topItem = l.Len() - l.NumRows()
		}

		return true
	}

	l.GUIObject.KeyClick = func(key Key, chr int) bool {
		switch key {
		case K_DOWN:
			if l.selectedItem < l.Len()-1 {
				l.selectedItem++
				if l.selectedItem > l.topItem+l.NumRows()-1 {
					l.topItem = l.selectedItem - l.NumRows() + 1
				}
				return true
			}

		case K_UP:
			if l.selectedItem > 0 {
				l.selectedItem--
				if l.selectedItem < l.topItem {
					l.topItem = l.selectedItem
				}
				return true
			}

		}
		return false
	}

	return l
}

func (l *GUIListBox) Add(item string) int {
	newItem := new(guiListboxItem)
	newItem.text = item

	l.List.PushBack(newItem)
	l.items++

	return l.items - 1
}

func (l *GUIListBox) Remove(n int) {
	if n > l.List.Len() {
		return
	}

	item := l.List.Front()

	for i := 0; i < l.items; i++ {
		item = item.Next()
	}

	l.List.Remove(item)
	l.items--
}

func (l *GUIListBox) Selected() int {
	return l.selectedItem
}

func (l *GUIListBox) SetSelected(n int) {
	if n >= 0 && n < l.Len() {
		l.selectedItem = n
	}
}

func (l *GUIListBox) Top() int {
	return l.topItem
}

func (l *GUIListBox) SetTop(n int) {
	if n >= 0 && n <= l.Len()-l.NumRows() {
		l.topItem = n
	}
}

func (l *GUIListBox) Text(n int) string {
	if n > l.List.Len() || n < 0 {
		return "+++ERROR+++"
	}

	item := l.List.Front()

	for i := 0; i < n; i++ {
		item = item.Next()
	}

	return item.Value.(*guiListboxItem).text
}

func (l *GUIListBox) NumItems() int {
	return l.items
}

func (l *GUIListBox) NumRows() int {
	rect := l.GUIObject.Rect
	return int((rect.Y2 - rect.Y1) / l.Font.GetHeight())
}

func (l *GUIListBox) Clear() {
	l.List.Init()
	l.items = 0
}
