# WARNING

This API still needs input. There may still be breakage. A stable legacy API is provided for those wanting to use the C like API.

# Haaf's Game Engine (HGE) in Go!

This is a binding (currently the hge-unix version) and port of HGE to Go.
It's split up into a few parts:
* The binding which provides a binding to the main C++ HGE class via C binding of hge-unix.
* The helper classes available in C++ where ported to Go using just the core binding.
* There is also a native reimplementation of the engine with some breaking changes.
* The legacy API which provides a near one-to-one mapping of the C++ singleton.

As it currently sits, the binding & helpers work pretty well together. There's still a lot of work yet to be done with the native port, but some of the core functionality is already there and working.

## Requirements:
* For the bindings: hge-unix is all that's supported. In addition, the c_api branch must be compiled and built with -DBUILD_C_API=ON from http://code.google.com/p/hge-unix/
** As a note, if it's installed to /usr/local you may need to put /usr/local/lib/pkgconfig in your PKG_CONFIG_PATH shell variable.
* You'll need a working Go 1 compiler.
** It will also need cgo to compile the bindings. (I've only tested with gc Go.)
* Even though it's primarally tested with Linux, I have had success running it with Windows, but you're currently on your own with that.
* The native bindings currently require SDL & OpenGL (banthar/Go-SDL & chsc/gogl)

## Building:
### Native:
The native build uses build tags with the go command to conditionally include files. In most cases, the default go get or build should work, but if you're wanting to use a different graphics, sound, input, etc backend they'll be documented here as they become available.

* You will likely want to use go get to fetch all dependencies automatically.
* go get github.com/losinggeneration/hge-go/hge
* Note that hge-go/hge isn't *all* that useful for inclusion. You'll need to get some of the subdirectories of hge at the very least, and likely some of hge-go/helpers/. One good way to do this is to build the tutorials.
* You can however try running make which tries to run go build for the entire engine.

### Bindings:
If you'd prefer the library that depend upon the hge-unix (and potentially original hge) codebase(s), you'll need to make sure you have hge compiled as well as the hge c api wrapper (which is required for the Go bindings.) If they're installed and pkg-conifg can find them, using go get should work as expected.

* You should use go get: go get github.com/losinggeneration/hge-go/binding/hge
* As noted above, you'll likely need some of the other directories within hge-go/binding. Building hge-go/binding/tutorials is a good way to build & test out the library.
