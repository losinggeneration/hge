# WARNING

This API still needs input. There may still be breakage. A stable legacy API is provided for those wanting to use the C like API.

# Haaf's Game Engine (HGE) in Go!

======================

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
* The native bindings currently require SDL

## Building:
* You should use go get: go get github.com/losinggeneration/hge-go/hge
* Additionally, you can: import "github.com/losinggeneration/hge-go/hge" and it should work as expected.
* The bindings are in "github.com/losinggeneration/hge-go/binding/hge"
