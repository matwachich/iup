package iupplot

import (
	"fmt"
	"unsafe"

	"github.com/matwachich/iup"
)

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -L${SRCDIR}/../lib/im -L${SRCDIR}/../lib/cd
#cgo LDFLAGS: -liup_plot -lcdcontextplus -lcdgl -liupgl -liupcd -liup
#cgo LDFLAGS: -lftgl
#cgo LDFLAGS: -lstdc++
#cgo LDFLAGS: -lcd -lfreetype6 -lz
#cgo LDFLAGS: -lwinspool -lglu32 -lopengl32 -lgdiplus -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>

#include <iup.h>
#include <iup_plot.h>

void __IupPlotPaintTo(Ihandle *ih, void* cnv) {
	IupPlotPaintTo(ih, (struct _cdCanvas*)cnv);
}
*/
import "C"

//Open initializes IupPlot widget class.
func Open() {
	//void IupPlotOpen(void);
	C.IupPlotOpen()
}

//Plot returns the identifier of the created plot, or NULL if an error occurs.
func Plot() iup.Ihandle {
	//Ihandle* IupPlot(void);
	return mkih(C.IupPlot())
}

/***********************************************/
/*           Additional API                    */

//void IupPlotBegin(Ihandle *ih, int strXdata);
func Begin(ih iup.Ihandle, strXdata int) {
	C.IupPlotBegin(pih(ih), C.int(strXdata))
}

//void IupPlotAdd(Ihandle *ih, double x, double y);
func Add(ih iup.Ihandle, x, y float64) {
	C.IupPlotAdd(pih(ih), C.double(x), C.double(y))
}

//void IupPlotAddStr(Ihandle *ih, const char* x, double y);
func AddStr(ih iup.Ihandle, x string, y float64) {
	c_x := C.CString(x)
	defer C.free(unsafe.Pointer(c_x))

	C.IupPlotAddStr(pih(ih), c_x, C.double(y))
}

//void IupPlotAddSegment(Ihandle *ih, double x, double y);
func AddSegment(ih iup.Ihandle, x, y float64) {
	C.IupPlotAddSegment(pih(ih), C.double(x), C.double(y))
}

//int  IupPlotEnd(Ihandle *ih);
func End(ih iup.Ihandle) int {
	return int(C.IupPlotEnd(pih(ih)))
}

//int  IupPlotLoadData(Ihandle* ih, const char* filename, int strXdata);
func LoadData(ih iup.Ihandle, fileName string, strXdata int) int {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	return int(C.IupPlotLoadData(pih(ih), c_fileName, C.int(strXdata)))
}

/* available only when linking with "iupluaplot" */
//int IupPlotSetFormula(Ihandle* ih, int sample_count, const char* formula, const char* init);

//void IupPlotInsert(Ihandle *ih, int ds_index, int sample_index, double x, double y);
func Insert(ih iup.Ihandle, dsIndex, sampleIndex int, x, y float64) {
	C.IupPlotInsert(pih(ih), C.int(dsIndex), C.int(sampleIndex), C.double(x), C.double(y))
}

//void IupPlotInsertStr(Ihandle *ih, int ds_index, int sample_index, const char* x, double y);
func InsertStr(ih iup.Ihandle, dsIndex, sampleIndex int, x string, y float64) {
	c_x := C.CString(x)
	defer C.free(unsafe.Pointer(c_x))

	C.IupPlotInsertStr(pih(ih), C.int(dsIndex), C.int(sampleIndex), c_x, C.double(y))
}

//void IupPlotInsertSegment(Ihandle *ih, int ds_index, int sample_index, double x, double y);
func InsertSegment(ih iup.Ihandle, dsIndex, sampleIndex int, x, y float64) {
	C.IupPlotInsertSegment(pih(ih), C.int(dsIndex), C.int(sampleIndex), C.double(x), C.double(y))
}

//void IupPlotInsertStrSamples(Ihandle* ih, int ds_index, int sample_index, const char** x, double* y, int count);
func InsertStrSamples(ih iup.Ihandle, dsIndex, sampleIndex int, x []string, y []float64) {
	if len(x) != len(y) {
		panic(fmt.Errorf("bad parameter passed to iupplot.InsertStrSamples"))
	}

	c_x, c_y := make([]*C.char, len(x)), make([]C.double, len(y))
	for i := 0; i < len(x); i++ {
		c_x[i], c_y[i] = C.CString(x[i]), C.double(y[i])
	}
	defer func() {
		for i := 0; i < len(x); i++ {
			C.free(unsafe.Pointer(c_x[i]))
		}
	}()

	C.IupPlotInsertStrSamples(pih(ih), C.int(dsIndex), C.int(sampleIndex), &c_x[0], &c_y[0], C.int(len(x)))
}

//void IupPlotInsertSamples(Ihandle* ih, int ds_index, int sample_index, double *x, double *y, int count);
func InsertSamples(ih iup.Ihandle, dsIndex, sampleIndex int, x, y []float64) {
	if len(x) != len(y) {
		panic(fmt.Errorf("bad parameter passed to iupplot.InsertSamples"))
	}

	c_x, c_y := make([]C.double, len(x)), make([]C.double, len(y))
	for i := 0; i < len(x); i++ {
		c_x[i], c_y[i] = C.double(x[i]), C.double(y[i])
	}

	C.IupPlotInsertSamples(pih(ih), C.int(dsIndex), C.int(sampleIndex), &c_x[0], &c_y[0], C.int(len(x)))
}

//void IupPlotAddSamples(Ihandle* ih, int ds_index, double *x, double *y, int count);
func AddSamples(ih iup.Ihandle, dsIndex int, x, y []float64) {
	if len(x) != len(y) {
		panic(fmt.Errorf("bad parameter passed to iupplot.AddSamples"))
	}

	c_x, c_y := make([]C.double, len(x)), make([]C.double, len(y))
	for i := 0; i < len(x); i++ {
		c_x[i], c_y[i] = C.double(x[i]), C.double(y[i])
	}

	C.IupPlotAddSamples(pih(ih), C.int(dsIndex), &c_x[0], &c_y[0], C.int(len(x)))
}

//void IupPlotAddStrSamples(Ihandle* ih, int ds_index, const char** x, double* y, int count);
func AddStrSamples(ih iup.Ihandle, dsIndex int, x []string, y []float64) {
	if len(x) != len(y) {
		panic(fmt.Errorf("bad parameter passed to iupplot.AddStrSamples"))
	}

	c_x, c_y := make([]*C.char, len(x)), make([]C.double, len(y))
	for i := 0; i < len(x); i++ {
		c_x[i], c_y[i] = C.CString(x[i]), C.double(y[i])
	}
	defer func() {
		for i := 0; i < len(x); i++ {
			C.free(unsafe.Pointer(c_x[i]))
		}
	}()

	C.IupPlotAddStrSamples(pih(ih), C.int(dsIndex), &c_x[0], &c_y[0], C.int(len(x)))
}

//void IupPlotGetSample(Ihandle* ih, int ds_index, int sample_index, double *x, double *y);
func GetSample(ih iup.Ihandle, dsIndex, sampleIndex int) (x, y float64) {
	C.IupPlotGetSample(pih(ih), C.int(dsIndex), C.int(sampleIndex), (*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)))
	return
}

//void IupPlotGetSampleStr(Ihandle* ih, int ds_index, int sample_index, const char* *x, double *y);
func GetSampleStr(ih iup.Ihandle, dsIndex, sampleIndex int) (x string, y float64) {
	var c_x *C.char
	C.IupPlotGetSampleStr(pih(ih), C.int(dsIndex), C.int(sampleIndex), &c_x, (*C.double)(unsafe.Pointer(&y)))
	x = C.GoString(c_x)
	return
}

//int  IupPlotGetSampleSelection(Ihandle* ih, int ds_index, int sample_index);
func GetSampleSelection(ih iup.Ihandle, dsIndex, sampleIndex int) int {
	return int(C.IupPlotGetSampleSelection(pih(ih), C.int(dsIndex), C.int(sampleIndex)))
}

//double IupPlotGetSampleExtra(Ihandle* ih, int ds_index, int sample_index);
func GetSampleExtra(ih iup.Ihandle, dsIndex, sampleIndex int) float64 {
	return float64(C.IupPlotGetSampleExtra(pih(ih), C.int(dsIndex), C.int(sampleIndex)))
}

//void IupPlotSetSample(Ihandle* ih, int ds_index, int sample_index, double x, double y);
func SetSample(ih iup.Ihandle, dsIndex, sampleIndex int, x, y float64) {
	C.IupPlotSetSample(pih(ih), C.int(dsIndex), C.int(sampleIndex), C.double(x), C.double(y))
}

//void IupPlotSetSampleStr(Ihandle* ih, int ds_index, int sample_index, const char* x, double y);
func SetSampleStr(ih iup.Ihandle, dsIndex, sampleIndex int, x string, y float64) {
	c_x := C.CString(x)
	defer C.free(unsafe.Pointer(c_x))

	C.IupPlotSetSampleStr(pih(ih), C.int(dsIndex), C.int(sampleIndex), c_x, C.double(y))
}

//void IupPlotSetSampleSelection(Ihandle* ih, int ds_index, int sample_index, int selected);
func SetSampleSelection(ih iup.Ihandle, dsIndex, sampleIndex, selected int) {
	C.IupPlotSetSampleSelection(pih(ih), C.int(dsIndex), C.int(sampleIndex), C.int(selected))
}

//void IupPlotSetSampleExtra(Ihandle* ih, int ds_index, int sample_index, double extra);
func SetSampleExtra(ih iup.Ihandle, dsIndex, sampleIndex int, extra float64) {
	C.IupPlotSetSampleExtra(pih(ih), C.int(dsIndex), C.int(sampleIndex), C.double(extra))
}

//void IupPlotTransform(Ihandle* ih, double x, double y, double *cnv_x, double *cnv_y);
func Transform(ih iup.Ihandle, x, y float64) (cnv_x, cnv_y float64) {
	C.IupPlotTransform(pih(ih), C.double(x), C.double(y), (*C.double)(unsafe.Pointer(&cnv_x)), (*C.double)(unsafe.Pointer(&cnv_y)))
	return
}

//void IupPlotTransformTo(Ihandle* ih, double cnv_x, double cnv_y, double *x, double *y);
func TransformTo(ih iup.Ihandle, cnv_x, cnv_y float64) (x, y float64) {
	C.IupPlotTransformTo(pih(ih), C.double(cnv_x), C.double(cnv_y), (*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)))
	return
}

//int  IupPlotFindSample(Ihandle* ih, double cnv_x, double cnv_y, int *ds_index, int *sample_index);
func FindSample(ih iup.Ihandle, cnv_x, cnv_y float64) (dsIndex, sampleIndex int) {
	C.IupPlotFindSample(pih(ih), C.double(cnv_x), C.double(cnv_y), (*C.int)(unsafe.Pointer(&dsIndex)), (*C.int)(unsafe.Pointer(&sampleIndex)))
	return
}

//struct _cdCanvas;

//void IupPlotPaintTo(Ihandle *ih, struct _cdCanvas* cnv);
func PaintTo(ih iup.Ihandle, cnv uintptr) {
	C.__IupPlotPaintTo(pih(ih), unsafe.Pointer(cnv))
}

/***********************************************/

/* -------------------------------------------------------------------------- */

func pih(ih iup.Ihandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}

func mkih(p *C.Ihandle) iup.Ihandle {
	return iup.Ihandle(unsafe.Pointer(p))
}
