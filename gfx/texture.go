package gfx

import (
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"runtime"
	"unsafe"

	gl "github.com/go-gl/gl/v2.1/gl"
)

var (
	texFilter bool
)

func SetTextureFilter(filter bool) {
	texFilter = filter
}

type img struct {
	bytes []byte
	w, h  uint32
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
	is.w, is.h = uint32(i.Bounds().Dx()), uint32(i.Bounds().Dy())
	// 4 is one byte for each R-G-B-A
	is.bytes = make([]byte, is.w*is.h*4)
	lw := is.w * 4                  // line width
	d := uint32(len(is.bytes)) - lw // The bottom line

	// Reverse the image lines
	for s := 0; s < len(rgba.Pix); s += rgba.Stride {
		copy(is.bytes[d:d+lw], rgba.Pix[s:s+rgba.Stride])
		d -= lw
	}

	return is, nil
}

func isPowerOfTwo(x uint32) bool {
	return ((x & (x - 1)) == 0)
}

func nextPowerOfTwo(x uint32) uint32 {
	x--
	for i := uint32(1); i < 32; i *= 2 {
		x |= x >> i
	}
	return x + 1
}

// HGE Handle type
type Texture struct {
	tex     uint32
	texType uint32
	w, h    uint32
}

func NewTexture(width, height int) *Texture {
	var t uint32
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
		gl.TexImage2D(t.texType, 0, gl.RGBA, int32(b.w), int32(b.h), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&b.bytes[0]))
	} else {
		potw, poth := nextPowerOfTwo(b.w), nextPowerOfTwo(b.h)
		gl.TexImage2D(t.texType, 0, gl.RGBA, int32(potw), int32(poth), 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
		gl.TexSubImage2D(t.texType, 0, 0, 0, int32(b.w), int32(b.h), gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&b.bytes[0]))
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

func default_texture_filter() {
	var filter int32
	if texFilter {
		filter = gl.LINEAR
	} else {
		filter = gl.NEAREST
	}

	gl.TexParameteri(defaultTextureType, gl.TEXTURE_MIN_FILTER, filter)
	gl.TexParameteri(defaultTextureType, gl.TEXTURE_MAG_FILTER, filter)
}

func tex_coord(v Vertex) {
	gl.TexCoord2d(float64(v.X), float64(v.Y))
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
