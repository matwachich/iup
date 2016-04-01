package iupim

import (
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/im
#cgo LDFLAGS: -liupim -lim -liup
#cgo LDFLAGS: -lz
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupim.h>

void* IupGetNativeHandleImage(void* handle);
void* IupGetImageNativeHandle(const void* image);

Ihandle* IupImageFromImImage(const void* image);
*/
import "C"

//LoadImage loads an IupImage from a file.
//When failed the global attribute "IUPIM_LASTERROR" is set with a message describing the error.
func LoadImage(fileName string) iup.Ihandle {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	//Ihandle* IupLoadImage(const char* file_name);
	return mkih(C.IupLoadImage(c_fileName))
}

//SaveImage saves an IupImage to a file.
//When failed the global attribute "IUPIM_LASTERROR" is set with a message describing the error.
func SaveImage(ih iup.Ihandle, fileName, format string) bool {
	c_fileName, c_format := C.CString(fileName), C.CString(format)
	defer C.free(unsafe.Pointer(c_fileName))
	defer C.free(unsafe.Pointer(c_format))

	//int IupSaveImage(Ihandle* ih, const char* file_name, const char* format);
	return int(C.IupSaveImage(pih(ih), c_fileName, c_format)) != 0
}

//LoadAnimation
func LoadAnimation(fileName string) iup.Ihandle {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	//Ihandle* IupLoadAnimation(const char* file_name);
	return mkih(C.IupLoadAnimation(c_fileName))
}

//LoadAnimationFrames
func LoadAnimationFrames(fileNames []string) iup.Ihandle {
	if len(fileNames) == 0 {
		return iup.Ihandle(0)
	}

	pFileNames := make([]*C.char, len(fileNames))
	for i := 0; i < len(fileNames); i++ {
		pFileNames[i] = C.CString(fileNames[i])
	}
	defer func() {
		for i := 0; i < len(pFileNames); i++ {
			C.free(unsafe.Pointer(pFileNames[i]))
		}
	}()

	//Ihandle* IupLoadAnimationFrames(const char** file_name_list, int file_count);
	return mkih(C.IupLoadAnimationFrames((**C.char)(unsafe.Pointer(&pFileNames[0])), C.int(len(pFileNames))))
}

//#ifdef __IM_IMAGE_H
//imImage* IupGetNativeHandleImage(void* handle);
//void* IupGetImageNativeHandle(const imImage* image);

//Ihandle* IupImageFromImImage(const imImage* image);
//#endif

type ImImage uintptr

//GetNativeHandleImage native Handle to imImage
func GetNativeHandleImage(handle uintptr) ImImage {
	return ImImage(C.IupGetNativeHandleImage(unsafe.Pointer(handle)))
}

//GetImageNativeHandle imImage to native handle
func GetImageNativeHandle(image ImImage) uintptr {
	return uintptr(C.IupGetImageNativeHandle(unsafe.Pointer(image)))
}

//ImageFromImImage imImage to IupImage
func ImageFromImImage(image ImImage) iup.Ihandle {
	return mkih(C.IupImageFromImImage(unsafe.Pointer(image)))
}

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}
