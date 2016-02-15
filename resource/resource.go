package resource

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"

	hge "github.com/losinggeneration/hge-go"
)

type Pointer uintptr

type Resource struct {
	Pointer
}

var resourceHGE *hge.HGE

func init() {
	resourceHGE = hge.New()
}

// Loads a resource into memory from disk.
func NewResource(filename string) (*Resource, hge.Dword) {
	var s C.DWORD
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	r := new(Resource)

	r.Pointer = Pointer(C.HGE_Resource_Load(resourceHGE.HGE, fname, &s))

	runtime.SetFinalizer(r, func(runtime *Resource) {
		r.Free()
	})

	return r, hge.Dword(s)
}

// Deletes a previously loaded resource from memory.
func (r *Resource) Free() {
	fmt.Println("Resource.Free")
	C.HGE_Resource_Free(resourceHGE.HGE, unsafe.Pointer(r.Pointer))
}

// Loads a resource, puts the loaded data into a byte array, and frees the data.
func LoadBytes(filename string) []byte {
	r, size := NewResource(filename)

	if r == nil {
		return nil
	}

	b := C.GoBytes(unsafe.Pointer(r.Pointer), C.int(size))

	return b
}

// Loads a resource, puts the data into a string, and frees the data.
func LoadString(filename string) *string {
	r, size := NewResource(filename)

	if r == nil {
		return nil
	}

	s := C.GoStringN((*C.char)(unsafe.Pointer(r.Pointer)), C.int(size))

	return &s
}

// Attaches a resource pack.
func AttachPack(filename string, a ...interface{}) bool {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	if len(a) == 1 {
		var password *C.char

		switch a[0].(type) {
		case string:
			password = C.CString(a[0].(string))
			defer C.free(unsafe.Pointer(password))
		}

		return C.HGE_Resource_AttachPack(resourceHGE.HGE, fname, password) == 1
	}

	return C.HGE_Resource_AttachPack(resourceHGE.HGE, fname, nil) == 1
}

// Removes a resource pack.
func RemovePack(filename string) {
	fname := C.CString(filename)
	defer C.free(unsafe.Pointer(fname))

	C.HGE_Resource_RemovePack(resourceHGE.HGE, fname)
}

// Removes all resource packs previously attached.
func RemoveAllPacks() {
	C.HGE_Resource_RemoveAllPacks(resourceHGE.HGE)
}

// Builds absolute file path.
func MakePath(a ...interface{}) string {
	if len(a) == 1 {
		if filename, ok := a[0].(string); ok {
			fname := C.CString(filename)
			defer C.free(unsafe.Pointer(fname))

			return C.GoString(C.HGE_Resource_MakePath(resourceHGE.HGE, fname))
		}
	}

	return C.GoString(C.HGE_Resource_MakePath(resourceHGE.HGE, nil))
}

// Enumerates files by given wildcard.
func EnumFiles(a ...interface{}) string {
	if len(a) == 1 {
		if wildcard, ok := a[0].(string); ok {
			wcard := C.CString(wildcard)
			defer C.free(unsafe.Pointer(wcard))

			return C.GoString(C.HGE_Resource_EnumFiles(resourceHGE.HGE, wcard))
		}
	}

	return C.GoString(C.HGE_Resource_EnumFiles(resourceHGE.HGE, nil))
}

// Enumerates folders by given wildcard.
func EnumFolders(a ...interface{}) string {
	if len(a) == 1 {
		if wildcard, ok := a[0].(string); ok {
			wcard := C.CString(wildcard)
			defer C.free(unsafe.Pointer(wcard))

			return C.GoString(C.HGE_Resource_EnumFolders(resourceHGE.HGE, wcard))
		}
	}

	return C.GoString(C.HGE_Resource_EnumFolders(resourceHGE.HGE, nil))
}
