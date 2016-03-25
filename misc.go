package iup

import (
	"unsafe"
)

/*
#include <stdlib.h>
*/
import "C"

func bool2int(v bool) int {
	if v {
		return 1
	} else {
		return 0
	}
}

func optionalAction(action []string) *C.char {
	if len(action) == 0 {
		return nil
	} else {
		return cStrOrNull(action[0])
	}
}

func cStrOrNull(s string) *C.char {
	if len(s) == 0 {
		return nil
	}
	return C.CString(s)
}

func cStrFree(p *C.char) {
	if p != nil {
		C.free(unsafe.Pointer(p))
	}
}

// Convert string pointers (char*) passed by callbacks to GoString
func CStrToString(pstr uintptr) string {
	return C.GoString((*C.char)(unsafe.Pointer(pstr)))
}
