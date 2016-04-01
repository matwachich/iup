//Creates a multiline source code text editor that uses the Scintilla library.
//
//Scintilla is a free library that provides text-editing functions,
//with an emphasis on advanced features for source code editing.
//It comes with complete source code and a license that permits
//use in any free project or commercial product, and it is available on http://www.scintilla.org/.
//
//IupScintilla library includes the Scintilla 3.6.2 source code, so no external references are needed.
//
//Supported in Windows and in the systems the GTK driver is supported.
package iupscintilla

import (
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/cd -L${SRCDIR}/../lib/im
#cgo LDFLAGS: -liup_scintilla -liupcontrols -liupcd -liup
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32 -limm32

#include <stdlib.h>
#include <string.h>

#include <iup.h>
#include <iup_scintilla.h>

void* IupScintillaSendMessage(Ihandle* ih, unsigned int iMessage, void* wParam, void* lParam);
*/
import "C"

//Open
func Open() {
	//void IupScintillaOpen(void);
	C.IupScintillaOpen()
}

//Scintilla returns the identifier of the created editing component, or NULL if an error occurs.
func Scintilla() iup.Ihandle {
	//Ihandle *IupScintilla(void);
	return iup.Ihandle(unsafe.Pointer(C.IupScintilla()))
}

//SendMessage sends a message to the Scintilla control in any platform.
func SendMessage(ih iup.Ihandle, msg uint, wParam, lParam uintptr) uintptr {
	//sptr_t IupScintillaSendMessage(Ihandle* ih, unsigned int iMessage, uptr_t wParam, sptr_t lParam);
	return uintptr(C.IupScintillaSendMessage(pih(ih), C.uint(msg), unsafe.Pointer(wParam), unsafe.Pointer(lParam)))
}

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}
