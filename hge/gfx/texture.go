// +build: opengl

package gfx

// import "fmt"	
import (
	"errors"
	gl "github.com/chsc/gogl/gl21"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"runtime"
)

type img struct {
	bytes []byte
	w, h  gl.Sizei
}

func decodeImage(r io.Reader) (*img, error) {
	i, _, e := image.Decode(r)
	if e != nil {
		return nil, e
	}

	rgba, ok := i.(*image.NRGBA)
	if !ok {
		return nil, errors.New("Texuter must be an NRGBA image.")
	}

	is := new(img)
	is.w, is.h = gl.Sizei(i.Bounds().Dx()), gl.Sizei(i.Bounds().Dy())
	is.bytes = make([]byte, is.w*is.h*4)
	lw := is.w * 4
	d := gl.Sizei(len(is.bytes)) - lw

	for s := 0; s < len(rgba.Pix); s += rgba.Stride {
		copy(is.bytes[d:d+lw], rgba.Pix[s:s+rgba.Stride])
		d -= lw
	}

	return is, nil
}

// HGE Handle type
type Texture struct {
	tex  gl.Uint
	w, h gl.Sizei
}

func NewTexture(width, height int) *Texture {
	var t gl.Uint
	gl.GenTextures(1, &t)
	tex := Texture{tex: t}

	runtime.SetFinalizer(&tex, func(texture *Texture) {
		texture.free()
	})

	return &tex
}

func LoadTexture(filename string, a ...interface{}) (*Texture, error) {
	t := NewTexture(0, 0)

	f, e := os.Open(filename)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	b, e := decodeImage(f)
	if e != nil {
		return nil, e
	}

	t.w, t.h = b.w, b.h

	gl.GenTextures(1, &t.tex)
	gl.BindTexture(gl.TEXTURE_2D, t.tex)
	gl.TexImage2D(gl.TEXTURE_2D, 0, 4, b.w, b.h, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&b.bytes[0]))

	return t, nil
}

func (t *Texture) free() {
	gl.DeleteTextures(1, &t.tex)
}

func (t *Texture) Width(a ...interface{}) int {
	return int(t.w)
}

func (t *Texture) Height(a ...interface{}) int {
	return int(t.h)
}

func (t *Texture) Lock(a ...interface{}) *uint32 {
	return nil
}

func (t *Texture) Unlock() {
}
