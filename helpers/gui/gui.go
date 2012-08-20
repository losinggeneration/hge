package gui

import (
	"container/list"
	. "github.com/losinggeneration/hge-go/helpers/rect"
	. "github.com/losinggeneration/hge-go/helpers/sprite"
	"github.com/losinggeneration/hge-go/hge"
	. "github.com/losinggeneration/hge-go/hge/input"
)

const (
	GUI_NONAVKEYS = 0
	GUI_LEFTRIGHT = 1
	GUI_UPDOWN    = 2
	GUI_CYCLED    = 4
)

type GUIObject struct {
	Id                       int
	Static, Visible, Enabled bool
	Rect                     Rect
	Color                    hge.Dword
	*GUI

	Render func()
	Update func(dt float64)

	Enter     func()
	Leave     func()
	Reset     func()
	IsDone    func() bool
	Focus     func(focused bool)
	MouseOver func(over bool)

	MouseMove    func(x, y float64) bool
	MouseLButton func(down bool) bool
	MouseRButton func(down bool) bool
	MouseWheel   func(notches int) bool
	KeyClick     func(key Key, chr int) bool

	SetColor func(color hge.Dword)
}

func (gobj *GUIObject) Initialize() {
	gobj.Render = nil
	gobj.Update = func(dt float64) {}

	gobj.Enter = func() {}
	gobj.Leave = func() {}
	gobj.Reset = func() {}
	gobj.IsDone = func() bool { return true }
	gobj.Focus = func(focused bool) {}
	gobj.MouseOver = func(over bool) {}

	gobj.MouseMove = func(x, y float64) bool { return false }
	gobj.MouseLButton = func(down bool) bool { return false }
	gobj.MouseRButton = func(down bool) bool { return false }
	gobj.MouseWheel = func(notches int) bool { return false }
	gobj.KeyClick = func(key Key, chr int) bool { return false }

	gobj.SetColor = func(color hge.Dword) { gobj.Color = color }
}

type GUI struct {
	ctrls                                    *list.List
	ctrlLock, ctrlFocus, ctrlOver            *GUIObject
	navMode, enterLeave                      int
	cursor                                   *Sprite
	mouse                                    Mouse
	lPressed, lReleased, rPressed, rReleased bool
}

func NewGUI() GUI {
	var g GUI

	g.ctrls = list.New()

	g.navMode = GUI_NONAVKEYS

	return g
}

func elementById(id int, list *list.List) *list.Element {
	for e := list.Front(); e != nil; e = e.Next() {
		ctrl := e.Value.(*GUIObject)
		if ctrl.Id == id {
			return e
		}
	}

	return nil
}

func (g *GUI) AddCtrl(ctrl *GUIObject) {
	e := elementById(ctrl.Id, g.ctrls)

	if e != nil {
		g.DelCtrl(ctrl.Id)
	}

	ctrl.GUI = g
	g.ctrls.PushBack(ctrl)
}

func (g *GUI) DelCtrl(id int) {
	e := elementById(id, g.ctrls)

	if e != nil {
		g.ctrls.Remove(e)
	}
}

func (g GUI) GetCtrl(id int) *GUIObject {
	e := elementById(id, g.ctrls)

	if e == nil {
		hge.New().Log("No such GUI ctrl id (%d)", id)
		return nil
	}

	return e.Value.(*GUIObject)
}

func (g *GUI) MoveCtrl(id int, x, y float64) {
	var ctrl = g.GetCtrl(id)
	if ctrl == nil {
		return
	}

	ctrl.Rect.X2 = x + ctrl.Rect.X2 - ctrl.Rect.X1
	ctrl.Rect.Y2 = y + ctrl.Rect.Y2 - ctrl.Rect.Y1
	ctrl.Rect.X1 = x
	ctrl.Rect.Y1 = y
}

func (g *GUI) ShowCtrl(id int, visible bool) {
	ctrl := g.GetCtrl(id)
	if ctrl == nil {
		return
	}
	ctrl.Visible = visible
}

func (g *GUI) EnableCtrl(id int, enabled bool) {
	ctrl := g.GetCtrl(id)
	if ctrl == nil {
		return
	}
	ctrl.Enabled = enabled
}

func (g *GUI) SetNavMode(mode int) {
	g.navMode = mode
}

func (g *GUI) SetCursor(spr *Sprite) {
	g.cursor = spr
}

func (g *GUI) SetColor(color hge.Dword) {
	for e := g.ctrls.Front(); e != nil; e = e.Next() {
		e.Value.(*GUIObject).SetColor(color)
	}
}

func (g *GUI) SetFocus(id int) {
	ctrlNewFocus := g.GetCtrl(id)

	if ctrlNewFocus == g.ctrlFocus {
		return
	}

	if ctrlNewFocus == nil {
		if g.ctrlFocus != nil {
			g.ctrlFocus.Focus(false)
			g.ctrlFocus = nil
		}
	} else if !ctrlNewFocus.Static && ctrlNewFocus.Visible && ctrlNewFocus.Enabled {
		if g.ctrlFocus != nil {
			g.ctrlFocus.Focus(false)
		}
		ctrlNewFocus.Focus(true)
		g.ctrlFocus = ctrlNewFocus
	}
}

func (g GUI) GetFocus() int {
	if g.ctrlFocus != nil {
		return g.ctrlFocus.Id
	}

	return 0
}

func (g *GUI) Enter() {
	for e := g.ctrls.Front(); e != nil; e = e.Next() {
		e.Value.(*GUIObject).Enter()
	}

	g.enterLeave = 2
}

func (g *GUI) Leave() {
	for e := g.ctrls.Front(); e != nil; e = e.Next() {
		e.Value.(*GUIObject).Leave()
	}

	g.ctrlFocus, g.ctrlOver, g.ctrlLock = nil, nil, nil
	g.enterLeave = 1
}

func (g *GUI) Reset() {
	for e := g.ctrls.Front(); e != nil; e = e.Next() {
		e.Value.(*GUIObject).Reset()
	}

	g.ctrlLock, g.ctrlOver, g.ctrlFocus = nil, nil, nil
}

func (g *GUI) Move(dx, dy float64) {
	for e := g.ctrls.Front(); e != nil; e = e.Next() {
		ctrl := e.Value.(*GUIObject)
		ctrl.Rect.X1 += dx
		ctrl.Rect.X2 += dx
		ctrl.Rect.Y1 += dy
		ctrl.Rect.Y2 += dy
	}
}

func (g *GUI) Update(dt float64) int {
	// Update the mouse variables
	g.mouse.Pos()
	g.lPressed = NewKey(K_LBUTTON).Down()
	g.lReleased = NewKey(K_LBUTTON).Up()
	g.rPressed = NewKey(K_RBUTTON).Down()
	g.rReleased = NewKey(K_RBUTTON).Up()
	g.mouse.WheelMovement()

	// Update all controls
	for e := g.ctrls.Front(); e != nil; e = e.Next() {
		e.Value.(*GUIObject).Update(dt)
	}

	// Handle Enter/Leave
	if g.enterLeave > 0 {
		done := true
		for e := g.ctrls.Front(); e != nil; e = e.Next() {
			if !e.Value.(*GUIObject).IsDone() {
				done = false
				break
			}
		}
		if !done {
			return 0
		} else {
			if g.enterLeave == 1 {
				return -1
			} else {
				g.enterLeave = 0
			}
		}
	}

	// Handle keys
	key := GetKey()
	if ((g.navMode&GUI_LEFTRIGHT) == GUI_LEFTRIGHT && key == K_LEFT) ||
		((g.navMode&GUI_UPDOWN) == GUI_UPDOWN && key == K_UP) {
		ctrl := g.ctrlFocus
		if ctrl == nil {
			e := g.ctrls.Front()
			if e == nil {
				return 0
			}

			ctrl = e.Value.(*GUIObject)
			if ctrl == nil {
				return 0
			}
		}

		for e := elementById(ctrl.Id, g.ctrls).Prev(); ; e = e.Prev() {
			if e == nil && (g.navMode&GUI_CYCLED) == GUI_CYCLED || g.ctrlFocus == nil {
				ctrl = g.ctrls.Back().Value.(*GUIObject)
			} else {
				ctrl = e.Value.(*GUIObject)
			}

			if ctrl == g.ctrlFocus {
				break
			}

			if ctrl.Static == false || ctrl.Visible == true || ctrl.Enabled == true {
				break
			}
		}

		if ctrl != g.ctrlFocus {
			if g.ctrlFocus != nil {
				g.ctrlFocus.Focus(false)
			}
			if ctrl != nil {
				ctrl.Focus(true)
			}
			g.ctrlFocus = ctrl
		}
	} else if ((g.navMode&GUI_LEFTRIGHT) == GUI_LEFTRIGHT && key == K_RIGHT) ||
		((g.navMode&GUI_UPDOWN) == GUI_UPDOWN && key == K_DOWN) {
		ctrl := g.ctrlFocus
		if ctrl == nil {
			e := g.ctrls.Back()
			if e == nil {
				return 0
			}

			ctrl = e.Value.(*GUIObject)
			if ctrl == nil {
				return 0
			}
		}

		for e := elementById(ctrl.Id, g.ctrls).Next(); ; e = e.Next() {
			if e == nil && (g.navMode&GUI_CYCLED) == GUI_CYCLED || g.ctrlFocus == nil {
				ctrl = g.ctrls.Front().Value.(*GUIObject)
			} else {
				ctrl = e.Value.(*GUIObject)
			}

			if ctrl == g.ctrlFocus {
				break
			}

			if ctrl.Static == false || ctrl.Visible == true || ctrl.Enabled == true {
				break
			}
		}

		if ctrl != g.ctrlFocus {
			if g.ctrlFocus != nil {
				g.ctrlFocus.Focus(false)
			}
			if ctrl != nil {
				ctrl.Focus(true)
			}
			g.ctrlFocus = ctrl
		}
	} else if g.ctrlFocus != nil && key > 0 && key != K_LBUTTON && key != K_RBUTTON {
		if g.ctrlFocus.KeyClick(key, GetChar()) {
			return g.ctrlFocus.Id
		}
	}

	// Handle mouse
	lDown := NewKey(K_LBUTTON).State()
	rDown := NewKey(K_RBUTTON).State()

	if g.ctrlLock != nil {
		ctrl := g.ctrlLock
		if !lDown && !rDown {
			g.ctrlLock = nil
		}
		if g.process(ctrl) {
			return ctrl.Id
		}
	} else {
		for e := g.ctrls.Front(); e != nil; e = e.Next() {
			ctrl := e.Value.(*GUIObject)
			if ctrl.Rect.TestPoint(g.mouse.X, g.mouse.Y) && ctrl.Enabled {
				if g.ctrlOver != ctrl {
					if g.ctrlOver != nil {
						g.ctrlOver.MouseOver(false)
					}

					ctrl.MouseOver(true)
					g.ctrlOver = ctrl
				}

				if g.process(ctrl) {
					return ctrl.Id
				} else {
					return 0
				}
			}
		}

		if g.ctrlOver != nil {
			g.ctrlOver.MouseOver(false)
			g.ctrlOver = nil
		}

	}

	return 0
}

func (g *GUI) Render() {
	for e := g.ctrls.Front(); e != nil; e = e.Next() {
		ctrl := e.Value.(*GUIObject)
		if ctrl.Visible {
			ctrl.Render()
		}
	}

	if g.mouse.IsOver() && g.cursor != nil {
		g.cursor.Render(g.mouse.X, g.mouse.Y)
	}
}

func (g *GUI) process(ctrl *GUIObject) bool {
	result := false

	if g.lPressed {
		g.ctrlLock = ctrl
		g.SetFocus(ctrl.Id)
		result = result || ctrl.MouseLButton(true)
	}
	if g.rPressed {
		g.ctrlLock = ctrl
		g.SetFocus(ctrl.Id)
		result = result || ctrl.MouseRButton(true)
	}
	if g.lReleased {
		result = result || ctrl.MouseLButton(false)
	}
	if g.rReleased {
		result = result || ctrl.MouseRButton(false)
	}
	if g.mouse.Wheel > 0 {
		result = result || ctrl.MouseWheel(g.mouse.Wheel)
	}
	result = result || ctrl.MouseMove(g.mouse.X-ctrl.Rect.X1, g.mouse.Y-ctrl.Rect.Y1)

	return result
}
