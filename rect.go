package hge

import (
	"math"
)

type Rect struct {
	X1, Y1, X2, Y2 float32
	clean          bool
}

func NewRect(x1 float32, y1 float32, x2 float32, y2 float32) Rect {
	var r Rect

	r.X1 = x1
	r.Y1 = y1
	r.X2 = x2
	r.Y2 = y2
	r.clean = false

	return r
}

func NewCleanRect() Rect {
	var r Rect

	r.clean = true

	return r
}

func (rect *Rect) Clear() {
	rect.clean = true
}

func (rect *Rect) IsClean() bool {
	return rect.clean
}

func (rect *Rect) Set(x1 float32, y1 float32, x2 float32, y2 float32) {
	rect.X1 = x1
	rect.X2 = x2
	rect.Y1 = y1
	rect.Y2 = y2
	rect.clean = false
}

func (rect *Rect) SetRadius(x float32, y float32, r float32) {
	rect.X1 = x - r
	rect.X2 = x + r
	rect.Y1 = y - r
	rect.Y2 = y + r
	rect.clean = false
}

func (rect *Rect) Encapsulate(x float32, y float32) {
	if rect.clean {
		rect.X1 = x
		rect.X2 = x
		rect.Y1 = y
		rect.Y2 = y
		rect.clean = false
	} else {
		if x < rect.X1 {
			rect.X1 = x
		}
		if x > rect.X2 {
			rect.X2 = x
		}
		if y < rect.Y1 {
			rect.Y1 = y
		}
		if y > rect.Y2 {
			rect.Y2 = y
		}
	}
}

func (rect *Rect) TestPoint(x float32, y float32) bool {
	if x >= rect.X1 && x < rect.X2 && y >= rect.Y1 && y < rect.Y2 {
		return true
	}

	return false
}

func (rect *Rect) Intersect(r *Rect) bool {
	if math.Abs(float64(rect.X1+rect.X2-r.X1-r.X2)) < float64(rect.X2-rect.X1+r.X2-r.X1) {
		if math.Abs(float64(rect.Y1+rect.Y2-r.Y1-r.Y2)) < float64(rect.Y2-rect.Y1+r.Y2-r.Y1) {
			return true
		}
	}

	return false
}
