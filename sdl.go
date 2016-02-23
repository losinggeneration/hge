// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package hge

import (
	"fmt"

	"github.com/banthar/Go-SDL/sdl"
)

type Hwnd sdl.Surface

func setTitle() {
	sdl.WM_SetCaption(stateStrings[TITLE], stateStrings[TITLE])
}

func initNative() error {
	if sdl.Init(sdl.INIT_EVERYTHING) == -1 {
		return fmt.Errorf(sdl.GetError())
	}

	// 	if sdl.GL_LoadLibrary() == -1 {
	// 		sdl.Quit()
	// 		return fmt.Errorf(sdl.GetError())
	// 	}

	vidinfo := sdl.GetVideoInfo()
	origScreenWidth := vidinfo.Current_w
	origScreenHeight := vidinfo.Current_h

	// Create window
	bpp := 4
	if stateInts[SCREENBPP] >= 32 {
		bpp = 8
	}

	zbuffer := 0
	if stateBools[ZBUFFER] {
		zbuffer = 16
	}

	setTitle()

	sdl.GL_SetAttribute(sdl.GL_RED_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_GREEN_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_BLUE_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_ALPHA_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_DEPTH_SIZE, zbuffer)
	sdl.GL_SetAttribute(sdl.GL_ACCELERATED_VISUAL, 1)
	sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	// 	sdl.GL_SetAttribute(sdl.GL_SWAP_CONTROL, vsync ? 1 : 0);

	flags := uint32(sdl.OPENGL)
	if !stateBools[WINDOWED] {
		flags |= sdl.FULLSCREEN
	}

	width, height := stateInts[SCREENWIDTH], stateInts[SCREENHEIGHT]
	if width > int(origScreenWidth) {
		width = int(origScreenWidth)
	}

	if height > int(origScreenHeight) {
		height = int(origScreenHeight)
	}

	fmt.Printf("Screen: %dx%d\n", width, height)

	hwnd := sdl.SetVideoMode(width, height, stateInts[SCREENBPP], flags)
	if hwnd == nil {
		sdl.Quit()
		return fmt.Errorf(sdl.GetError())
	}

	stateHwnds[HWND] = (*Hwnd)(hwnd)

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

	sdl.ShowCursor(cursor)

	return nil
}

func shutdownNative() {
	sdl.Quit()
}

func initPowerStatus() error {
	return nil
}
