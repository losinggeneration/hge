package color

import (
	"math"

	. "github.com/losinggeneration/hge/binding/hge"
)

func colorClamp(x *float64) {
	if *x < 0.0 {
		*x = 0.0
	}
	if *x > 1.0 {
		*x = 1.0
	}
}

type ColorRGB struct {
	R, G, B, A float64
}

type Color interface {
	Clamp()
	SetHWColor(col Dword)
	GetHWColor() Dword
}

func NewRGB(r, g, b, a float64) ColorRGB {
	return ColorRGB{r, g, b, a}
}

func NewRGBCol(col Dword) ColorRGB {
	var c ColorRGB

	c.SetHWColor(col)

	return c
}

func (c ColorRGB) Sub(c2 ColorRGB) ColorRGB {
	return NewRGB(c.R-c2.R, c.G-c2.G, c.B-c2.B, c.A-c2.A)
}

func (c ColorRGB) Add(c2 ColorRGB) ColorRGB {
	return NewRGB(c.R+c2.R, c.G+c2.G, c.B+c2.B, c.A+c2.A)
}

func (c ColorRGB) Mul(c2 ColorRGB) ColorRGB {
	return NewRGB(c.R*c2.R, c.G*c2.G, c.B*c2.B, c.A*c2.A)
}

func (c *ColorRGB) SubEqual(c2 ColorRGB) *ColorRGB {
	c.R -= c2.R
	c.G -= c2.G
	c.B -= c2.B
	c.A -= c2.A

	return c
}

func (c *ColorRGB) AddEqual(c2 ColorRGB) *ColorRGB {
	c.R += c2.R
	c.G += c2.G
	c.B += c2.B
	c.A += c2.A

	return c
}

func (c ColorRGB) Eq(c2 ColorRGB) bool {
	return c.R == c2.R && c.G == c2.G && c.B == c2.B && c.A == c2.A
}

func (c ColorRGB) DivScalar(scalar float64) ColorRGB {
	return NewRGB(c.R/scalar, c.G/scalar, c.B/scalar, c.A/scalar)
}

func (c ColorRGB) MulScalar(scalar float64) ColorRGB {
	return NewRGB(c.R*scalar, c.G*scalar, c.B*scalar, c.A*scalar)
}

func (c *ColorRGB) MulScalarEqual(scalar float64) *ColorRGB {
	c.R *= scalar
	c.G *= scalar
	c.B *= scalar
	c.A *= scalar

	return c
}

func (c *ColorRGB) Clamp() {
	colorClamp(&c.R)
	colorClamp(&c.G)
	colorClamp(&c.B)
	colorClamp(&c.A)
}

func (c *ColorRGB) SetHWColor(col Dword) {
	c.A = float64(col>>24) / 255.0
	c.R = float64((col>>16)&0xFF) / 255.0
	c.G = float64((col>>8)&0xFF) / 255.0
	c.B = float64(col&0xFF) / 255.0
}

func (c ColorRGB) HWColor() Dword {
	return (Dword(c.A*255.0) << 24) + (Dword(c.R*255.0) << 16) + (Dword(c.G*255.0) << 8) + Dword(c.B*255.0)
}

type ColorHSV struct {
	H, S, V, A float64
}

func NewHSV(h, s, v, a float64) ColorHSV {
	return ColorHSV{h, s, v, a}
}

func NewHSVCol(col Dword) ColorHSV {
	var c ColorHSV

	c.SetHWColor(col)

	return c
}

func (c ColorHSV) Sub(c2 ColorHSV) ColorHSV {
	return NewHSV(c.H-c2.H, c.S-c2.S, c.V-c2.V, c.A-c2.A)
}

func (c ColorHSV) Add(c2 ColorHSV) ColorHSV {
	return NewHSV(c.H+c2.H, c.S+c2.S, c.V+c2.V, c.A+c2.A)
}

func (c ColorHSV) Mul(c2 ColorHSV) ColorHSV {
	return NewHSV(c.H*c2.H, c.S*c2.S, c.V*c2.V, c.A*c2.A)
}

func (c *ColorHSV) SubEqual(c2 ColorHSV) *ColorHSV {
	c.H -= c2.H
	c.S -= c2.S
	c.V -= c2.V
	c.A -= c2.A

	return c
}

func (c *ColorHSV) AddEqual(c2 ColorHSV) *ColorHSV {
	c.H += c2.H
	c.S += c2.S
	c.V += c2.V
	c.A += c2.A

	return c
}

func (c ColorHSV) Eq(c2 ColorHSV) bool {
	return c.H == c2.H && c.S == c2.S && c.V == c2.V && c.A == c2.A
}

func (c ColorHSV) DivScalar(scalar float64) ColorHSV {
	return NewHSV(c.H/scalar, c.S/scalar, c.V/scalar, c.A/scalar)
}

func (c ColorHSV) MulScalar(scalar float64) ColorHSV {
	return NewHSV(c.H*scalar, c.S*scalar, c.V*scalar, c.A*scalar)
}

func (c *ColorHSV) MulScalarEqual(scalar float64) *ColorHSV {
	c.H *= scalar
	c.S *= scalar
	c.V *= scalar
	c.A *= scalar

	return c
}

func (c *ColorHSV) Clamp() {
	colorClamp(&c.H)
	colorClamp(&c.S)
	colorClamp(&c.V)
	colorClamp(&c.A)
}

func (c *ColorHSV) SetHWColor(col Dword) {
	c.A = float64(col>>24) / 255.0
	r := float64((col>>16)&0xFF) / 255.0
	g := float64((col>>8)&0xFF) / 255.0
	b := float64(col&0xFF) / 255.0

	minv := math.Min(math.Min(r, g), b)
	maxv := math.Max(math.Max(r, g), b)
	delta := maxv - minv

	c.V = maxv

	if delta == 0 {
		c.H, c.S = 0, 0
	} else {
		c.S = delta / maxv
		del_R := (((maxv - r) / 6) + (delta / 2)) / delta
		del_G := (((maxv - g) / 6) + (delta / 2)) / delta
		del_B := (((maxv - b) / 6) + (delta / 2)) / delta

		if r == maxv {
			c.H = del_B - del_G
		} else if g == maxv {
			c.H = (1.0 / 3.0) + del_R - del_B
		} else if b == maxv {
			c.H = (2.0 / 3.0) + del_G - del_R
		}

		if c.H < 0 {
			c.H += 1
		}
		if c.H > 1 {
			c.H -= 1
		}
	}
}

func (c ColorHSV) HWColor() Dword {
	var r, g, b float64
	if c.S == 0 {
		r = c.V
		g = c.V
		b = c.V
	} else {
		xh := c.H * 6
		if xh == 6 {
			xh = 0
		}
		i := math.Floor(xh)
		p1 := c.V * (1 - c.S)
		p2 := c.V * (1 - c.S*(xh-i))
		p3 := c.V * (1 - c.S*(1-(xh-i)))

		if i == 0 {
			r = c.V
			g = p3
			b = p1
		} else if i == 1 {
			r = p2
			g = c.V
			b = p1
		} else if i == 2 {
			r = p1
			g = c.V
			b = p3
		} else if i == 3 {
			r = p1
			g = p2
			b = c.V
		} else if i == 4 {
			r = p3
			g = p1
			b = c.V
		} else {
			r = c.V
			g = p1
			b = p2
		}
	}

	return (Dword(c.A*255.0) << 24) + (Dword(r*255.0) << 16) + (Dword(g*255.0) << 8) + Dword(b*255.0)
}
