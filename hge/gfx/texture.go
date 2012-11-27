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

var (
	texFilter bool
)

func SetTextureFilter(filter bool) {
	texFilter = filter
}

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
		return nil, errors.New("Texture must be an NRGBA image.")
	}

	is := new(img)
	is.w, is.h = gl.Sizei(i.Bounds().Dx()), gl.Sizei(i.Bounds().Dy())
	// 4 is one byte for each R-G-B-A
	is.bytes = make([]byte, is.w*is.h*4)
	lw := is.w * 4                    // line width
	d := gl.Sizei(len(is.bytes)) - lw // The bottom line

	// Reverse the image lines
	for s := 0; s < len(rgba.Pix); s += rgba.Stride {
		copy(is.bytes[d:d+lw], rgba.Pix[s:s+rgba.Stride])
		d -= lw
	}

	return is, nil
}

func isPowerOfTwo(x gl.Sizei) bool {
	return ((x & (x - 1)) == 0)
}

func nextPowerOfTwo(x gl.Sizei) gl.Sizei {
	x--
	for i := uint(1); i < 32; i *= 2 {
		x |= x >> i
	}
	return x + 1
}

// HGE Handle type
type Texture struct {
	tex     gl.Uint
	texType gl.Enum
	w, h    gl.Sizei
}

func NewTexture(width, height int) *Texture {
	var t gl.Uint
	gl.GenTextures(1, &t)
	tex := Texture{tex: t, texType: gl.TEXTURE_2D}

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
	t.bind()

	gl.TexParameterf(t.texType, gl.TEXTURE_MIN_LOD, 0.0)
	gl.TexParameterf(t.texType, gl.TEXTURE_MAX_LOD, 0.0)
	gl.TexParameteri(t.texType, gl.TEXTURE_BASE_LEVEL, 0)
	gl.TexParameteri(t.texType, gl.TEXTURE_MAX_LEVEL, 0)

	// Textures are required to be a power of two unless there's
	// GL_ARB_texture_non_power_of_two (an extention for 1.4 and in 2.0 core)
	if isPowerOfTwo(b.w) && isPowerOfTwo(b.h) {
		gl.TexImage2D(t.texType, 0, gl.RGBA, b.w, b.h, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&b.bytes[0]))
	} else {
		potw, poth := nextPowerOfTwo(b.w), nextPowerOfTwo(b.h)
		gl.TexImage2D(t.texType, 0, gl.RGBA, potw, poth, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
		gl.TexSubImage2D(t.texType, 0, 0, 0, b.w, b.h, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&b.bytes[0]))
	}

	// Unbind this texture now
	gl.BindTexture(t.texType, 0)
	return t, nil
}

func (t *Texture) free() {
	gl.DeleteTextures(1, &t.tex)
}

func (t *Texture) bind() {
	gl.BindTexture(t.texType, t.tex)
}

func (t *Texture) filter() {
	var filter gl.Int
	if texFilter {
		filter = gl.LINEAR
	} else {
		filter = gl.NEAREST
	}

	gl.TexParameteri(t.texType, gl.TEXTURE_MIN_FILTER, filter)
	gl.TexParameteri(t.texType, gl.TEXTURE_MAG_FILTER, filter)
}

func tex_coord(v Vertex) {
	gl.TexCoord2d(gl.Double(v.X), gl.Double(v.Y))
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
