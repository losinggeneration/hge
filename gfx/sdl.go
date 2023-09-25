// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
// I doubt there will ever be the need for anything like: +build sdl,opengl
// or: +build sdl,software
// but it's an option

package gfx

import "github.com/veandco/go-sdl2/sdl"

type Hwnd struct {
	*sdl.Window
}

var hwnd *Hwnd

// States
func SetHwnd(h *Hwnd) {
	hwnd = h
}

func updateSize(width, height int) {
	if hwnd != nil {
		hwnd.SetSize(int32(width), int32(height))
	}
}

func updatePosition(x, y int) {
	if hwnd != nil {
		hwnd.SetPosition(int32(x), int32(y))
	}
}

func swapBuffers() {
	if hwnd != nil {
		hwnd.GLSwap()
	}
}
