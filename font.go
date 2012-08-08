package hge

import (
	"C"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

const (
	TEXT_LEFT     = 0
	TEXT_RIGHT    = 1
	TEXT_CENTER   = 2
	TEXT_HORZMASK = 0x03

	TEXT_TOP      = 0
	TEXT_BOTTOM   = 4
	TEXT_MIDDLE   = 8
	TEXT_VERTMASK = 0x0C
)

const (
	fntHEADERTAG = "[HGEFONT]"
	fntBITMAPTAG = "Bitmap"
	fntCHARTAG   = "Char"
)

/*
 * * HGE Font class
 */
type Font struct {
	hge *HGE

	texture    Texture
	letters    [256]*Sprite
	pre        [256]float64
	post       [256]float64
	height     float64
	scale      float64
	proportion float64
	rot        float64
	tracking   float64
	spacing    float64

	col   Dword
	z     float64
	blend int
}

func getLines(file string) []string {
	lines := strings.FieldsFunc(file, func(r rune) bool {
		if r == '\n' || r == '\r' {
			return true
		}
		return false
	})

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}

func tokenizeLine(line string) (string, string, error) {
	if i := strings.Index(line, "="); i != -1 {
		return strings.TrimSpace(line[:i]), strings.TrimSpace(line[i+1:]), nil
	}

	if len(strings.TrimSpace(line)) == 0 {
		return "", "", nil
	}

	return "", "", errors.New("Unable to tokenize line")
}

func tokenizeChar(value string) (chr byte, x, y, w, h, a, c float64) {
	z := strings.Split(value, ",")
	chr = z[0][0]
	x1, _ := strconv.ParseFloat(z[1], 32)
	x = x1
	y1, _ := strconv.ParseFloat(z[2], 32)
	y = y1
	w1, _ := strconv.ParseFloat(z[3], 32)
	w = w1
	h1, _ := strconv.ParseFloat(z[4], 32)
	h = h1
	a1, _ := strconv.ParseFloat(z[5], 32)
	a = a1
	c1, _ := strconv.ParseFloat(z[6], 32)
	c = c1

	return
}

func NewFont(filename string, arg ...interface{}) *Font {
	mipmap := false

	if len(arg) == 1 {
		if m, ok := arg[0].(bool); ok {
			mipmap = m
		}
	}

	f := new(Font)

	f.hge = Create(VERSION)

	f.height = 0.0
	f.scale = 1.0
	f.proportion = 1.0
	f.rot = 0.0
	f.tracking = 0.0
	f.spacing = 1.0
	f.texture = 0

	f.z = 0.5
	f.blend = BLEND_COLORMUL | BLEND_ALPHABLEND | BLEND_NOZWRITE
	f.col = 0xFFFFFFFF

	data, size := f.hge.Resource_Load(filename)
	if data == nil || size == 0 {
		return nil
	}

	desc := C.GoBytes(unsafe.Pointer(data), C.int(size))
	f.hge.Resource_Free(data)

	lines := getLines(string(desc))

	if len(lines) == 0 || lines[0] != fntHEADERTAG {
		f.hge.System_Log("Font %s has incorrect format.", filename)
		return nil
	}

	// parse the font description
	for i := 1; i < len(lines); i++ {
		option, value, err := tokenizeLine(lines[i])

		if err == nil || len(lines[i]) == 0 || len(option) == 0 || len(value) == 0 {
			continue
		}

		if option == fntBITMAPTAG {
			f.texture = f.hge.Texture_Load(value, 0, mipmap)
		} else if option == fntCHARTAG {
			chr, x, y, w, h, a, c := tokenizeChar(value)

			sprt := NewSprite(f.texture, x, y, w, h)

			f.letters[chr] = &sprt
			f.pre[chr] = a
			f.post[chr] = c
		}
	}

	return f
}

func (f *Font) Render(x, y float64, align int, str string) {
	fx := x

	align &= TEXT_HORZMASK
	if align == TEXT_RIGHT {
		fx -= f.GetStringWidth(str, false)
	}
	if align == TEXT_CENTER {
		fx -= f.GetStringWidth(str, false) / 2.0
	}

	for j := 0; j < len(str); j++ {
		if str[j] == '\n' {
			y += f.height * f.scale * f.spacing
			fx = x
			if align == TEXT_RIGHT {
				fx -= f.GetStringWidth(string(str[j+1]), false)
			}
			if align == TEXT_CENTER {
				fx -= f.GetStringWidth(string(str[j+1]), false) / 2.0
			}
		} else {
			i := str[j]
			if f.letters[i] == nil {
				i = '?'
			}
			if f.letters[i] != nil {
				fx += f.pre[i] * f.scale * f.proportion
				f.letters[i].RenderEx(fx, y, f.rot, f.scale*f.proportion, f.scale)
				fx += (f.letters[i].GetWidth() + f.post[i] + f.tracking) * f.scale * f.proportion
			}
		}
	}
}

func (f *Font) Printf(x, y float64, align int, format string, arg ...interface{}) {
	f.Render(x, y, align, fmt.Sprintf(format, arg...))
}

func (f *Font) Printfb(x, y, w, h float64, align int, format string, arg ...interface{}) {
}

func (f *Font) SetColor(col Dword) {
	f.col = col

	for i := 0; i < 256; i++ {
		if f.letters[i] != nil {
			f.letters[i].SetColor(col)
		}
	}
}

func (f *Font) SetZ(z float64) {
	f.z = z

	for i := 0; i < 256; i++ {
		if f.letters[i] != nil {
			f.letters[i].SetZ(z)
		}
	}
}

func (f *Font) SetBlendMode(blend int) {
	f.blend = blend

	for i := 0; i < 256; i++ {
		if f.letters[i] != nil {
			f.letters[i].SetBlendMode(blend)
		}
	}
}

func (f *Font) SetScale(scale float64) {
	f.scale = scale
}

func (f *Font) SetProportion(prop float64) {
	f.proportion = prop
}

func (f *Font) SetRotation(rot float64) {
	f.rot = rot
}

func (f *Font) SetTracking(tracking float64) {
	f.tracking = tracking
}

func (f *Font) SetSpacing(spacing float64) {
	f.spacing = spacing
}

func (f Font) GetColor() Dword {
	return f.col
}

func (f Font) GetZ() float64 {
	return f.z
}

func (f Font) GetBlendMode() int {
	return f.blend
}

func (f Font) GetScale() float64 {
	return f.scale
}

func (f Font) GetProportion() float64 {
	return f.proportion
}

func (f Font) GetRotation() float64 {
	return f.rot
}

func (f Font) GetTracking() float64 {
	return f.tracking
}

func (f Font) GetSpacing() float64 {
	return f.spacing
}

func (f Font) GetSprite(chr byte) *Sprite {
	return f.letters[chr]
}

func (f Font) GetPreWidth(chr byte) float64 {
	return f.pre[chr]
}

func (f Font) GetPostWidth(chr byte) float64 {
	return f.post[chr]
}

func (f Font) GetHeight() float64 {
	return f.height
}

func (f Font) GetStringWidth(str string, arg ...interface{}) float64 {
	multiline := true
	w := 0.0

	if len(arg) == 1 {
		if m, ok := arg[0].(bool); ok {
			multiline = m
		}
	}

	for j := 0; j < len(str); j++ {
		linew := 0.0

		for ; str[j] != '\n'; j++ {
			i := str[j]
			if f.letters[i] == nil {
				i = '?'
			}
			if f.letters[i] != nil {
				linew += f.letters[i].GetWidth() + f.pre[i] + f.post[i] + f.tracking
			}

			j++
		}

		if !multiline {
			return linew * f.scale * f.proportion
		}

		if linew > w {
			w = linew
		}

		for str[j] == '\n' || str[j] == '\r' {
			j++
		}
	}

	return w * f.scale * f.proportion
}
