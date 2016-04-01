package iupcontrols

import (
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/cd -L${SRCDIR}/../lib/im
#cgo LDFLAGS: -liupcontrols -liupcd -liup
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lgdi32 -luser32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iupcontrols.h>
*/
import "C"

//int  IupControlsOpen(void);
func Open() int {
	return int(C.IupControlsOpen())
}

//Colorbar creates a color palette to enable a color selection from several samples.
//It can select one or two colors. The primary color is selected with the left mouse button,
//and the secondary color is selected with the right mouse button.
//You can double click a cell to change its color and you can double click the preview
//area to switch between primary and secondary colors.
//
//This is an additional control that depends on the CD library. It is included in the IupControls library.
//
//It inherits from IupCanvas.
func Colorbar() iup.Ihandle {
	//Ihandle* IupColorbar(void);
	return iup.Ihandle(unsafe.Pointer(C.IupColorbar()))
}

//Cells creates a grid widget (set of cells) that enables several application-specific drawing,
//such as: chess tables, tiles editors, degrade scales, drawable spreadsheets and so forth.
//
//This element is mostly based on application callbacks functions that determine the number of cells (rows and columns),
//their appearance and interaction. This mechanism offers full flexibility to applications,
//but requires programmers attention to avoid infinite loops inside this functions.
//Using callbacks, cells can be also grouped to form major or hierarchical elements, such as headers, footers etc.
//This callback approach was intentionally chosen to allow all cells to be dynamically and directly
//changed based on application's data structures. Since the size of each cell is given by the application
//the size of the control also must be given using SIZE or RASTERSIZE attributes.
//
//This is an additional control that depends on the CD library. It is included in the IupControls library.
//
//It inherits from IupCanvas.
func Cells() iup.Ihandle {
	//Ihandle* IupCells(void);
	return iup.Ihandle(unsafe.Pointer(C.IupCells()))
}

//ColorBrowser creates an element for selecting a color. The selection is done using a cylindrical projection of the RGB cube.
//The transformation defines a coordinate color system called HSI, that is still the RGB color space but using cylindrical coordinates.
//
//H is for Hue, and it is the angle around the RGB cube diagonal starting at red (RGB=255 0 0).
//
//S is for Saturation, and it is the normal distance from the color to the diagonal, normalized by its maximum value at the specified Hue.
//This also defines a point at the diagonal used to define I.
//
//I is for Intensity, and it is the distance from the point defined at the diagonal to black (RGB=0 0 0).
//I can also be seen as the projection of the color vector onto the diagonal. But I is not linear, see Notes below.
//
//This is an additional control that depends on the CD library. It is included in the IupControls library.
//
//For a dialog that simply returns the selected color, you can use function IupGetColor or IupColorDlg.
func ColorBrowser() iup.Ihandle {
	//Ihandle *IupColorBrowser(void);
	return iup.Ihandle(unsafe.Pointer(C.IupColorBrowser()))
}

//Gauge creates a Gauge control. Shows a percent value that can be updated to simulate a progression. It inherits from IupCanvas.
//
//This is an additional control that depends on the CD library. It is included in the IupControls library.
//
//It is recommended that new applications use the IupProgressBar control of the main library.
func Gauge() iup.Ihandle {
	//Ihandle *IupGauge(void);
	return iup.Ihandle(unsafe.Pointer(C.IupGauge()))
}

//Dial creates a dial for regulating a given angular variable.
//
//This is an additional control that depends on the CD library. It is included in the IupControls library.
//
//It inherits from IupCanvas.
func Dial(_type string) iup.Ihandle {
	c_type := C.CString(_type)
	defer C.free(unsafe.Pointer(c_type))

	//Ihandle *IupDial(const char* type);
	return iup.Ihandle(unsafe.Pointer(C.IupDial(c_type)))
}

//Matrix creates a matrix of alphanumeric fields. Therefore, all values of the matrix fields are strings.
//The matrix is not a grid container like many systems have. It inherits from IupCanvas.
//
//This is an additional control that depends on the CD library. It is included in the IupControls library.
//
//It has two modes of operation: normal and callback mode.
//In normal mode, string values are stored in attributes for each cell.
//In callback mode these attributes are ignored and the cells are filled with strings returned by the "VALUE_CB" callback.
//So the existence of this callback defines the mode the matrix will operate.
func Matrix(action ...string) iup.Ihandle {
	var c_action *C.char
	if len(action) == 0 {
		c_action = nil
	} else {
		c_action = C.CString(action[0])
		defer C.free(unsafe.Pointer(c_action))
	}

	//Ihandle* IupMatrix(const char *action);
	return iup.Ihandle(unsafe.Pointer(C.IupMatrix(c_action)))
}

//MatrixList creates an interface element that displays a list of items, just like IupList, but internally uses a IupMatrix.
//
//It uses the matrix columns to display labels, color boxes and check boxes in a way that is not possible using IupList.
//But the control mimics the IupList attributes, callbacks and interaction,
//so the usage by the programmer and by the user should be very similar.
//
//This is an additional control that depends on the IupControls library.
func MatrixList() iup.Ihandle {
	//Ihandle* IupMatrixList(void);
	return iup.Ihandle(unsafe.Pointer(C.IupMatrixList()))
}

/* Used by IupColorbar */
const (
	PRIMARY   = -1 // iup.ColorBar parameter
	SECONDARY = -2 // iup.ColorBar parameter
)
