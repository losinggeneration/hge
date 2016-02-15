package gfx

func (c Color) ToARGB() uint32 {
	return (c.A << 24) | (c.R << 16) | (c.G << 8) | c.B
}

func (c Color) ToRGBA() uint32 {
	return (c.R << 24) | (c.G << 16) | (c.B << 8) | c.A
}

// Converts a RGBA uint32 to a Color structure
func RGBAToColor(c uint32) Color {
	return Color{R: c >> 24, G: (c >> 16) & 0xFF, B: (c >> 8) & 0xFF, A: c & 0xFF}
}

// Converts an ARGB uint32 to a Color structure
func ARGBToColor(c uint32) Color {
	return Color{A: c >> 24, R: (c >> 16) & 0xFF, G: (c >> 8) & 0xFF, B: c & 0xFF}
}
