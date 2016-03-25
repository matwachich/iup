package iupweb

import (
	"unsafe"

	"github.com/matwachich/iup"
)

//broken! cannot use with GCC (visual C only)

/*
#cgo CFLAGS: -m32 -I./../include
#cgo LDFLAGS: -m32 -L${SRCDIR}/../lib
#cgo LDFLAGS: -liupweb -liupcontrols -liupole -liup
#cgo LDFLAGS: -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupweb.h>
*/
import "C"

//int IupWebBrowserOpen(void);
func Open() int {
	return int(C.IupWebBrowserOpen())
}

//Ihandle *IupWebBrowser(void);
func WebBrowser() iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(C.IupWebBrowser()))
}
