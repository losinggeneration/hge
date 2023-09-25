// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package hge

import (
	"errors"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/losinggeneration/hge/gfx"
)

const (
	WINDOWPOS_CENTERED = sdl.WINDOWPOS_CENTERED
)

func setTitle() {
	hwnd := stateHwnds[HWND]
	if hwnd != nil {
		window := hwnd.Window
		window.SetTitle(stateStrings[TITLE])
	}
}

func initNative(h *HGE) error {
	// Prevent crashes due to poor SDL & Go thread interactions
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	// Create window
	bpp := 4
	if stateInts[SCREENBPP] >= 32 {
		bpp = 8
	}

	zbuffer := 0
	if stateBools[ZBUFFER] {
		zbuffer = 16
	}

	err := errors.Join(sdl.GLSetAttribute(sdl.GL_RED_SIZE, bpp),
		sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, bpp),
		sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, bpp),
		sdl.GLSetAttribute(sdl.GL_ALPHA_SIZE, bpp),
		sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, zbuffer),
		sdl.GLSetAttribute(sdl.GL_ACCELERATED_VISUAL, 1),
		sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1),
	)
	if err != nil {
		sdl.Quit()
		return err
	}

	flags := uint32(sdl.WINDOW_OPENGL)
	if !stateBools[WINDOWED] {
		flags |= sdl.WINDOW_FULLSCREEN
	}

	title := stateStrings[TITLE]
	x, y := stateInts[SCREENX], stateInts[SCREENY]
	width, height := stateInts[SCREENWIDTH], stateInts[SCREENHEIGHT]

	window, err := sdl.CreateWindow(title, int32(x), int32(y), int32(width), int32(height), flags)
	if err != nil {
		sdl.Quit()
		return err
	}

	context, err := window.GLCreateContext()
	if err != nil {
		sdl.Quit()
		return err
	}

	if err = window.GLMakeCurrent(context); err != nil {
		sdl.Quit()
		return err
	}

	hwnd := &gfx.Hwnd{Window: window}
	if err = h.setStateHwndPrivate(HWND, hwnd); err != nil {
		sdl.Quit()
		return err
	}
	stateHwnds[HWND] = hwnd

	if !stateBools[WINDOWED] {
		// 		bMouseOver = true;
		// 		if !pHGE->bActive {
		// 			pHGE->_FocusChange(true);
		// 		}
	}

	cursor := sdl.ENABLE
	if stateBools[HIDEMOUSE] {
		cursor = sdl.DISABLE
	}

	if _, err = sdl.ShowCursor(cursor); err != nil {
		sdl.Quit()
		return err
	}

	return nil
}

func shutdownNative() {
	sdl.Quit()
}

func initPowerStatus() error {
	return nil
}
