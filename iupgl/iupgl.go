package iupgl

import (
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -m32 -I./../include
#cgo LDFLAGS: -m32 -L${SRCDIR}/../lib
#cgo LDFLAGS: -liupgl -liup
#cgo LDFLAGS: -lopengl32 -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupgl.h>
*/
import "C"

//void IupGLCanvasOpen(void);
func CanvasOpen() {
	C.IupGLCanvasOpen()
}

//Canvas creates an OpenGL canvas (drawing area for OpenGL). It inherits from IupCanvas.
func Canvas(action string) iup.Ihandle {
	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	//Ihandle *IupGLCanvas(const char *action);
	return mkih(C.IupGLCanvas(c_action))
}

//BackgroundBox creates a simple native container with no decorations, but with OpenGL enabled. It inherits from IupGLCanvas.
//
//OBS: this is identical to the IupBackgroundBox element, but with OpenGL enabled.
func BackgroundBox(child iup.Ihandle) iup.Ihandle {
	//Ihandle* IupGLBackgroundBox(Ihandle* child);
	return mkih(C.IupGLBackgroundBox(pih(child)))
}

//MakeCurrent activates the given canvas as the current OpenGL context.
//All subsequent OpenGL commands are directed to such canvas.
//The first call will set the global attributes GL_VERSION, GL_VENDOR and GL_RENDERER (since 3.16).
func MakeCurrent(ih iup.Ihandle) {
	//void IupGLMakeCurrent(Ihandle* ih);
	C.IupGLMakeCurrent(pih(ih))
}

//IsCurrent returns a non zero value if the given canvas is the current OpenGL context.
func IsCurrent(ih iup.Ihandle) bool {
	//int IupGLIsCurrent(Ihandle* ih);
	return int(C.IupGLIsCurrent(pih(ih))) != 0
}

//SwapBuffers makes the BACK buffer visible. This function is necessary when a double buffer is used.
func SwapBuffers(ih iup.Ihandle) {
	//void IupGLSwapBuffers(Ihandle* ih);
	C.IupGLSwapBuffers(pih(ih))
}

//Palette defines a color in the color palette. This function is necessary when INDEX color is used.
func Palette(ih iup.Ihandle, index int, r, g, b float32) {
	//void IupGLPalette(Ihandle* ih, int index, float r, float g, float b);
	C.IupGLPalette(pih(ih), C.int(index), C.float(r), C.float(g), C.float(b))
}

//UseFont creates a bitmap display list from the current FONT attribute.
//See the documentation of the wglUseFontBitmaps and glXUseXFont functions.
func UseFont(ih iup.Ihandle, first, count, listBase int) {
	//void IupGLUseFont(Ihandle* ih, int first, int count, int list_base);
	C.IupGLUseFont(pih(ih), C.int(first), C.int(count), C.int(listBase))
}

//Wait if gl is non zero it will call glFinish or glXWaitGL, else will call GdiFlush or glXWaitX.
func Wait(gl int) {
	//void IupGLWait(int gl);
	C.IupGLWait(C.int(gl))
}

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}
