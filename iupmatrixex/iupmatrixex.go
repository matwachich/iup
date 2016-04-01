//An extension library for IupMatrix.
//This library adds new features to the IupMatrix in order to extend the current features.
//Adds support for Import/Export, Clipboard, Undo/Redo, Find, Sort, Column Visibility,
//Numeric Columns, Numeric Units, Context Menu and others.
//
//The extension can be used in two different ways: using a new control called IupMatrixEx,
//or by calling IupMatrixExInit for an existing IupMatrix control.
//
//It can be used in callback mode or in standard more.
//
//Based on the DMatrix library created by Bruno Kassar and Luiz Cristóvão Gomez Coelho.
package iupmatrixex

import (
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/cd
#cgo LDFLAGS: -liupmatrixex -liupcontrols -liupcd -liup
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupmatrixex.h>
*/
import "C"

//Open
func Open() {
	//void IupMatrixExOpen(void);
	C.IupMatrixExOpen()
}

//MatrixEx returns the identifier of the created editing component, or NULL if an error occurs.
func MatrixEx() iup.Ihandle {
	//Ihandle* IupMatrixEx(void);
	return mkih(C.IupMatrixEx())
}

//MatrixExInit register the new attributes and callbacks in a regular IupMatrix control.
func MatrixExInit(ih iup.Ihandle) {
	//void IupMatrixExInit(Ihandle* ih);
	C.IupMatrixExInit(pih(ih))
}

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}
