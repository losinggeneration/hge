package ini

/*
#cgo pkg-config: hge-unix-c
#include "hge_c.h"
*/
import "C"

import (
	"github.com/losinggeneration/hge-go/binding/hge"
	"unsafe"
)

type Ini struct {
	Section, Name string
	iniHGE        *hge.HGE
}

func NewIni(section, name string) Ini {
	return Ini{section, name, hge.New()}
}

func (i Ini) SetInt(value int) {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	C.HGE_Ini_SetInt(i.iniHGE.HGE, s, n, C.int(value))
}

func (i Ini) GetInt(def_val int) int {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	return int(C.HGE_Ini_GetInt(i.iniHGE.HGE, s, n, C.int(def_val)))
}

func (i Ini) SetFloat(value float64) {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	C.HGE_Ini_SetFloat(i.iniHGE.HGE, s, n, C.float(value))
}

func (i Ini) GetFloat(def_val float64) float64 {
	s, n := C.CString(i.Section), C.CString(i.Name)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))

	return float64(C.HGE_Ini_GetFloat(i.iniHGE.HGE, s, n, C.float(def_val)))
}

func (i Ini) SetString(value string) {
	s, n, v := C.CString(i.Section), C.CString(i.Name), C.CString(value)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(v))

	C.HGE_Ini_SetString(i.iniHGE.HGE, s, n, v)
}

func (i Ini) GetString(def_val string) string {
	s, n, df := C.CString(i.Section), C.CString(i.Name), C.CString(def_val)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(df))

	return C.GoString(C.HGE_Ini_GetString(i.iniHGE.HGE, s, n, df))
}
