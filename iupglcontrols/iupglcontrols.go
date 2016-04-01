package iupglcontrols

import (
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/im -L${SRCDIR}/../lib/cd
#cgo LDFLAGS: -liupglcontrols -liupgl -liup
#cgo LDFLAGS: -lftgl
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lglu32 -lopengl32 -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupglcontrols.h>
*/
import "C"

//int  IupGLControlsOpen(void);
func Open() int {
	return int(C.IupGLControlsOpen())
}

//CanvasBox creates an OpenGL canvas container. It inherits from IupGLCanvas.
//
//This is an additional control that depends on the OpenGL library.
//It is included in the IupGLControls library.
//
//To use the controls available in the IupGLControls library inside your OpenGL canvas
//you must replace the IupGLCanvas by the IupGLCanvasBox element.
//
//It can have any number of children. Controls from the IupGLControls library can be used
//as children along with the void containers such as IupHbox, IupVbox, and so on, including IupFill.
//Native elements can also be placed on top although they will not be clipped by IupGLFrame and other IupGlControls containers.
//
//The elements that are a direct child of the box can be positioned using the VERTICALALIGN and/or HORIZONTALALIGN attributes,
//or using a position in pixels relative to the top left corner of the box by setting the attribute POSITION.
//
//Each direct children will be sized to its natural size by default, except if EXPANDHORIZONTAL or EXPANDVERTICAL are set.
//
//The box can be created with no elements and be dynamic filled using IupAppend or IupInsert.
func CanvasBox(children ...iup.Ihandle) iup.Ihandle {
	children = append(children, 0)

	//Ihandle* IupGLCanvasBoxv(Ihandle** children);
	//Ihandle* IupGLCanvasBox(Ihandle* child, ...);
	return mkih(C.IupGLCanvasBoxv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//SubCanvas creates an embedded OpenGL sub-canvas. It exists only inside an IupGLCanvasBox.
//
//This is an additional control that depends on the OpenGL library.
//It is included in the IupGLControls library.
//
//It is a void element that does not map to a native canvas.
//It mimics an IupCanvas with several attributes and callbacks in common,
//but everything is done inside a region of the IupGLCanvasBox.
//
//The element does not to be a direct child of the IupGLCanvasBox.
//It can be place inside void containers like IupHbox, IupVbox, IupGridBox and so on.
//
//This control is used as the base control for all visual elements of the IupGLControls library.
func SubCanvas() iup.Ihandle {
	//Ihandle* IupGLSubCanvas(void);
	return mkih(C.IupGLSubCanvas())
}

//Label creates an embedded OpenGL label interface element, which displays a text and/or an image.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func Label(title string) iup.Ihandle {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupGLLabel(const char* title);
	return mkih(C.IupGLLabel(c_title))
}

//Separator creates an embedded OpenGL separator interface element, which displays a vertical or horizontal line.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func Separator() iup.Ihandle {
	//Ihandle* IupGLSeparator(void);
	return mkih(C.IupGLSeparator())
}

//Button creates an embedded OpenGL interface element that is a button.
//When selected, this element activates a function in the application.
//Its visual presentation can contain a text and/or an image.
//It inherits from IupGLLabel. It exists only inside an IupGLCanvasBox.
func Button(title string) iup.Ihandle {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupGLButton(const char* title);
	return mkih(C.IupGLButton(c_title))
}

//Toggle creates an embedded OpenGL toggle interface element.
//It is a two-state (on/off) button that, when selected, generates
//an action that activates a function in the associated application.
//Its visual representation can contain a text and/or an image.
//It inherits from IupGLButton. It exists only inside an IupGLCanvasBox.
func Toggle(title string) iup.Ihandle {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupGLToggle(const char* title);
	return mkih(C.IupGLToggle(c_title))
}

//Link Creates an embedded OpenGL label that displays an underlined clickable text.
//It inherits from IupGLLabel. It exists only inside an IupGLCanvasBox.
func Link(url, title string) iup.Ihandle {
	c_url, c_title := C.CString(url), C.CString(title)
	defer C.free(unsafe.Pointer(c_url))
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupGLLink(const char *url, const char * title);
	return mkih(C.IupGLLink(c_url, c_title))
}

//ProgressBar creates an embedded OpenGL progress bar control.
//Shows a percent value that can be updated to simulate a progression.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func ProgressBar() iup.Ihandle {
	//Ihandle* IupGLProgressBar(void);
	return mkih(C.IupGLProgressBar())
}

//Val creates an embedded OpenGL Valuator control. Selects a value in a limited interval.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func Val() iup.Ihandle {
	//Ihandle* IupGLVal(void);
	return mkih(C.IupGLVal())
}

//Frame creates an embedded OpenGL container, which draws a frame with a title around its child.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func Frame(child iup.Ihandle) iup.Ihandle {
	//Ihandle* IupGLFrame(Ihandle* child);
	return mkih(C.IupGLFrame(pih(child)))
}

//Expander creates an embedded OpenGL container that can interactively show or hide its child.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func Expander(child iup.Ihandle) iup.Ihandle {
	//Ihandle* IupGLExpander(Ihandle* child);
	return mkih(C.IupGLExpander(pih(child)))
}

//ScrollBox Creates an embedded OpenGL container that allows its child to be scrolled.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func ScrollBox(child iup.Ihandle) iup.Ihandle {
	//Ihandle* IupGLScrollBox(Ihandle* child);
	return mkih(C.IupGLScrollBox(pih(child)))
}

//SizeBox creates a void container that allows its child to be resized.
//Allows expanding and contracting the child size in one or two directions.
//It inherits from IupGLSubCanvas. It exists only inside an IupGLCanvasBox.
func SizeBox(child iup.Ihandle) iup.Ihandle {
	//Ihandle* IupGLSizeBox(Ihandle* child);
	return mkih(C.IupGLSizeBox(pih(child)))
}

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}
