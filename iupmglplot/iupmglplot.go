package iupmglplot

import (
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -m32 -I./../include
#cgo LDFLAGS: -m32 -L${SRCDIR}/../lib -L${SRCDIR}/../lib/im -L${SRCDIR}/../lib/cd
#cgo LDFLAGS: -liup_mglplot -lcdcontextplus -lcdgl -liupgl -liupcd -liup
#cgo LDFLAGS: -lftgl
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lwinspool -lglu32 -lopengl32 -lgdiplus -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iup_mglplot.h>
*/
import "C"

type Pt1D struct {
	Name string
	Y    float64
}

type Pt2D struct {
	X, Y float64
}

type Pt3D struct {
	X, Y, Z float64
}

/* Initialize IupMglPlot widget class */
//void IupMglPlotOpen(void);
func Open() {
	C.IupMglPlotOpen()
}

/* Create an IupMglPlot widget instance */
//Ihandle* IupMglPlot(void);
func Plot() iup.Ihandle {
	return mkih(C.IupMglPlot())
}

/***********************************************/
/*           Additional API                    */

/* Linear Data Only */
//void IupMglPlotBegin(Ihandle *ih, int dim);
func Begin(ih iup.Ihandle, dim int) {
	C.IupMglPlotBegin(pih(ih), C.int(dim))
}

//void IupMglPlotAdd1D(Ihandle *ih, const char* name, double y);
func Add1D(ih iup.Ihandle, pt Pt1D) {
	c_name := C.CString(pt.Name)
	defer C.free(unsafe.Pointer(c_name))

	C.IupMglPlotAdd1D(pih(ih), c_name, C.double(pt.Y))
}

//void IupMglPlotAdd2D(Ihandle *ih, double x, double y);
func Add2D(ih iup.Ihandle, pt Pt2D) {
	C.IupMglPlotAdd2D(pih(ih), C.double(pt.X), C.double(pt.Y))
}

//void IupMglPlotAdd3D(Ihandle *ih, double x, double y, double z);
func Add3D(ih iup.Ihandle, pt Pt3D) {
	C.IupMglPlotAdd3D(pih(ih), C.double(pt.X), C.double(pt.Y), C.double(pt.Z))
}

//int  IupMglPlotEnd(Ihandle *ih);
func End(ih iup.Ihandle) int {
	return int(C.IupMglPlotEnd(pih(ih)))
}

/* Linear (dim=1,2,3), Planar (dim=1), Volumetric (dim=1) */
//int IupMglPlotNewDataSet(Ihandle *ih, int dim);
func NewDataSet(ih iup.Ihandle, dim int) int {
	return int(C.IupMglPlotNewDataSet(pih(ih), C.int(dim)))
}

/* Linear Data Only */
//void IupMglPlotInsert1D(Ihandle* ih, int ds_index, int sample_index, const char** names, const double* y, int count);
func Insert1D(ih iup.Ihandle, dsIndex, sampleIndex int, pts []Pt1D) {
	c_names, c_y := make([]*C.char, len(pts)), make([]C.double, len(pts))
	for i := 0; i < len(c_names); i++ {
		c_names[i], c_y[i] = C.CString(pts[i].Name), C.double(pts[i].Y)
	}
	defer func() {
		for i := 0; i < len(c_names); i++ {
			C.free(unsafe.Pointer(c_names[i]))
		}
	}()

	C.IupMglPlotInsert1D(pih(ih), C.int(dsIndex), C.int(sampleIndex), &c_names[0], &c_y[0], C.int(len(pts)))
}

//void IupMglPlotInsert2D(Ihandle* ih, int ds_index, int sample_index, const double* x, const double* y, int count);
func Insert2D(ih iup.Ihandle, dsIndex, sampleIndex int, pts []Pt2D) {
	c_x, c_y := make([]C.double, len(pts)), make([]C.double, len(pts))
	for i := 0; i < len(pts); i++ {
		c_x[i], c_y[i] = C.double(pts[i].X), C.double(pts[i].Y)
	}

	C.IupMglPlotInsert2D(pih(ih), C.int(dsIndex), C.int(sampleIndex), &c_x[0], &c_y[0], C.int(len(pts)))
}

//void IupMglPlotInsert3D(Ihandle* ih, int ds_index, int sample_index, const double* x, const double* y, const double* z, int count);
func Insert3D(ih iup.Ihandle, dsIndex, sampleIndex int, pts []Pt3D) {
	c_x, c_y, c_z := make([]C.double, len(pts)), make([]C.double, len(pts)), make([]C.double, len(pts))
	for i := 0; i < len(pts); i++ {
		c_x[i], c_y[i], c_z[i] = C.double(pts[i].X), C.double(pts[i].Y), C.double(pts[i].Z)
	}

	C.IupMglPlotInsert3D(pih(ih), C.int(dsIndex), C.int(sampleIndex), &c_x[0], &c_y[0], &c_z[0], C.int(len(pts)))
}

/* Linear Data Only */
//void IupMglPlotSet1D(Ihandle* ih, int ds_index, const char** names, const double* y, int count);
func Set1D(ih iup.Ihandle, dsIndex int, pts []Pt1D) {

}

//void IupMglPlotSet2D(Ihandle* ih, int ds_index, const double* x, const double* y, int count);
func Set2D(ih iup.Ihandle, dsIndex int, pts []Pt2D) {

}

//void IupMglPlotSet3D(Ihandle* ih, int ds_index, const double* x, const double* y, const double* z, int count);
func Set3D(ih iup.Ihandle, dsIndex int, pts []Pt3D) {

}

//void IupMglPlotSetFormula(Ihandle* ih, int ds_index, const char* formulaX, const char* formulaY, const char* formulaZ, int count);

/* Linear (dim=1), Planar (dim=1), Volumetric (dim=1) */
//void IupMglPlotSetData(Ihandle* ih, int ds_index, const double* data, int count_x, int count_y, int count_z);
func SetData(ih iup.Ihandle, dsIndex int, data []float64, countX, countY, countZ int) {
	c_data := make([]C.double, len(data))
	for i := 0; i < len(data); i++ {
		c_data[i] = C.double(data[i])
	}

	C.IupMglPlotSetData(pih(ih), C.int(dsIndex), &c_data[0], C.int(countX), C.int(countY), C.int(countZ))
}

//void IupMglPlotLoadData(Ihandle* ih, int ds_index, const char* filename, int count_x, int count_y, int count_z);
func LoadData(ih iup.Ihandle, dsIndex int, fileName string, countX, countY, countZ int) {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	C.IupMglPlotLoadData(pih(ih), C.int(dsIndex), c_fileName, C.int(countX), C.int(countY), C.int(countZ))
}

//void IupMglPlotSetFromFormula(Ihandle* ih, int ds_index, const char* formula, int count_x, int count_y, int count_z);

/* Only inside callbacks */
//void IupMglPlotTransform(Ihandle* ih, double x, double y, double z, int *ix, int *iy);
func Transform(ih iup.Ihandle, x, y, z float64) (ix, iy int) {
	C.IupMglPlotTransform(pih(ih), C.double(x), C.double(y), C.double(z), (*C.int)(unsafe.Pointer(&ix)), (*C.int)(unsafe.Pointer(&iy)))
	return
}

//void IupMglPlotTransformTo(Ihandle* ih, int ix, int iy, double *x, double *y, double *z);
func TransformTo(ih iup.Ihandle, ix, iy int) (x, y, z float64) {
	C.IupMglPlotTransformTo(pih(ih), C.int(ix), C.int(iy), (*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)), (*C.double)(unsafe.Pointer(&z)))
	return
}

/* Only inside callbacks */
//void IupMglPlotDrawMark(Ihandle* ih, double x, double y, double z);
func DrawMark(ih iup.Ihandle, x, y, z float64) {
	C.IupMglPlotDrawMark(pih(ih), C.double(x), C.double(y), C.double(z))
}

//void IupMglPlotDrawLine(Ihandle* ih, double x1, double y1, double z1, double x2, double y2, double z2);
func DrawLine(ih iup.Ihandle, x1, y1, z1, x2, y2, z2 float64) {
	C.IupMglPlotDrawLine(pih(ih), C.double(x1), C.double(y1), C.double(z1), C.double(x2), C.double(y2), C.double(z2))
}

//void IupMglPlotDrawText(Ihandle* ih, const char* text, double x, double y, double z);
func DrawText(ih iup.Ihandle, text string, x, y, z float64) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.IupMglPlotDrawText(pih(ih), c_text, C.double(x), C.double(y), C.double(z))
}

//void IupMglPlotPaintTo(Ihandle *ih, const char* format, int w, int h, double dpi, void *data);
func PaintTo(ih iup.Ihandle, format string, w, h int, dpi float64, data uintptr) {
	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	C.IupMglPlotPaintTo(pih(ih), c_format, C.int(w), C.int(h), C.double(dpi), unsafe.Pointer(data))
}

/***********************************************/

/* Utility label for showing TeX labels */
//Ihandle* IupMglLabel(const char* title);
func Label(title string) iup.Ihandle {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	return mkih(C.IupMglLabel(c_title))
}

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}
