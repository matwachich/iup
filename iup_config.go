package iup

import (
	"syscall"
	"unsafe"
)

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L${SRCDIR}/lib
#cgo LDFLAGS: -liup -liupcd -liupim -liupimglib
#cgo LDFLAGS: -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>
#include <iup.h>

#include <iup_config.h>

void __IupConfigRecentInit(Ihandle* ih, Ihandle* menu, void* recent_cb, int max_recent) {
	IupConfigRecentInit(ih, menu, (Icallback)recent_cb, max_recent);
}
*/
import "C"

//Config returns a new database where the variables will be stored.
func Config() Ihandle {
	//Ihandle* IupConfig(void);
	return Ihandle(unsafe.Pointer(C.IupConfig()))
}

//ConfigLoad loads the configuration file.
func ConfigLoad(ih Ihandle) int {
	//int IupConfigLoad(Ihandle* ih);
	return int(C.IupConfigLoad(ih.ptr()))
}

//ConfigSave saves the configuration file.
func ConfigSave(ih Ihandle) int {
	//int IupConfigSave(Ihandle* ih);
	return int(C.IupConfigSave(ih.ptr()))
}

/* -------------------------------------------------------------------------- */

func ConfigSetVariableStr(ih Ihandle, group, key string, value string) {
	c_group, c_key, c_value := C.CString(group), C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))
	defer C.free(unsafe.Pointer(c_value))

	//void IupConfigSetVariableStr(Ihandle* ih, const char* group, const char* key, const char* value);
	C.IupConfigSetVariableStr(ih.ptr(), c_group, c_key, c_value)
}

func ConfigSetVariableStrId(ih Ihandle, group, key string, id int, value string) {
	c_group, c_key, c_value := C.CString(group), C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))
	defer C.free(unsafe.Pointer(c_value))

	//void IupConfigSetVariableStrId(Ihandle* ih, const char* group, const char* key, int id, const char* value);
	C.IupConfigSetVariableStrId(ih.ptr(), c_group, c_key, C.int(id), c_value)
}

func ConfigSetVariableInt(ih Ihandle, group, key string, value int) {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//void IupConfigSetVariableInt(Ihandle* ih, const char* group, const char* key, int value);
	C.IupConfigSetVariableInt(ih.ptr(), c_group, c_key, C.int(value))
}

func ConfigSetVariableIntId(ih Ihandle, group, key string, id int, value int) {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//void IupConfigSetVariableIntId(Ihandle* ih, const char* group, const char* key, int id, int value);
	C.IupConfigSetVariableIntId(ih.ptr(), c_group, c_key, C.int(id), C.int(value))
}

func ConfigSetVariableDouble(ih Ihandle, group, key string, value float64) {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//void IupConfigSetVariableDouble(Ihandle* ih, const char* group, const char* key, double value);
	C.IupConfigSetVariableDouble(ih.ptr(), c_group, c_key, C.double(value))
}

func ConfigSetVariableDoubleId(ih Ihandle, group, key string, id int, value float64) {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//void IupConfigSetVariableDoubleId(Ihandle* ih, const char* group, const char* key, int id, double value);
	C.IupConfigSetVariableDoubleId(ih.ptr(), c_group, c_key, C.int(id), C.double(value))
}

func ConfigGetVariableStr(ih Ihandle, group, key string) string {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//const char* IupConfigGetVariableStr(Ihandle* ih, const char* group, const char* key);
	return C.GoString(C.IupConfigGetVariableStr(ih.ptr(), c_group, c_key))
}

func ConfigGetVariableStrId(ih Ihandle, group, key string, id int) string {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//const char* IupConfigGetVariableStrId(Ihandle* ih, const char* group, const char* key, int id);
	return C.GoString(C.IupConfigGetVariableStrId(ih.ptr(), c_group, c_key, C.int(id)))
}

func ConfigGetVariableInt(ih Ihandle, group, key string) int {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//int    IupConfigGetVariableInt(Ihandle* ih, const char* group, const char* key);
	return int(C.IupConfigGetVariableInt(ih.ptr(), c_group, c_key))
}

func ConfigGetVariableIntId(ih Ihandle, group, key string, id int) int {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//int    IupConfigGetVariableIntId(Ihandle* ih, const char* group, const char* key, int id);
	return int(C.IupConfigGetVariableIntId(ih.ptr(), c_group, c_key, C.int(id)))
}

func ConfigGetVariableDouble(ih Ihandle, group, key string) float64 {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//double IupConfigGetVariableDouble(Ihandle* ih, const char* group, const char* key);
	return float64(C.IupConfigGetVariableDouble(ih.ptr(), c_group, c_key))
}

func ConfigGetVariableDoubleId(ih Ihandle, group, key string, id int) float64 {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//double IupConfigGetVariableDoubleId(Ihandle* ih, const char* group, const char* key, int id);
	return float64(C.IupConfigGetVariableDoubleId(ih.ptr(), c_group, c_key, C.int(id)))
}

func ConfigGetVariableStrDef(ih Ihandle, group, key string, def string) string {
	c_group, c_key, c_def := C.CString(group), C.CString(key), C.CString(def)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))
	defer C.free(unsafe.Pointer(c_def))

	//const char* IupConfigGetVariableStrDef(Ihandle* ih, const char* group, const char* key, const char* def);
	return C.GoString(C.IupConfigGetVariableStrDef(ih.ptr(), c_group, c_key, c_def))
}

func ConfigGetVariableStrIdDef(ih Ihandle, group, key string, id int, def string) string {
	c_group, c_key, c_def := C.CString(group), C.CString(key), C.CString(def)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))
	defer C.free(unsafe.Pointer(c_def))

	//const char* IupConfigGetVariableStrIdDef(Ihandle* ih, const char* group, const char* key, int id, const char* def);
	return C.GoString(C.IupConfigGetVariableStrIdDef(ih.ptr(), c_group, c_key, C.int(id), c_def))
}

func ConfigGetVariableIntDef(ih Ihandle, group, key string, def int) int {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//int    IupConfigGetVariableIntDef(Ihandle* ih, const char* group, const char* key, int def);
	return int(C.IupConfigGetVariableIntDef(ih.ptr(), c_group, c_key, C.int(def)))
}

func ConfigGetVariableIntIdDef(ih Ihandle, group, key string, id int, def int) int {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//int    IupConfigGetVariableIntIdDef(Ihandle* ih, const char* group, const char* key, int id, int def);
	return int(C.IupConfigGetVariableIntIdDef(ih.ptr(), c_group, c_key, C.int(id), C.int(def)))
}

func ConfigGetVariableDoubleDef(ih Ihandle, group, key string, def float64) float64 {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//double IupConfigGetVariableDoubleDef(Ihandle* ih, const char* group, const char* key, double def);
	return float64(C.IupConfigGetVariableDoubleDef(ih.ptr(), c_group, c_key, C.double(def)))
}

func ConfigGetVariableDoubleIdDef(ih Ihandle, group, key string, id int, def float64) float64 {
	c_group, c_key := C.CString(group), C.CString(key)
	defer C.free(unsafe.Pointer(c_group))
	defer C.free(unsafe.Pointer(c_key))

	//double IupConfigGetVariableDoubleIdDef(Ihandle* ih, const char* group, const char* key, int id, double def);
	return float64(C.IupConfigGetVariableDoubleIdDef(ih.ptr(), c_group, c_key, C.int(id), C.double(def)))
}

/* -------------------------------------------------------------------------- */

func ConfigRecentInit(ih, menu Ihandle, recentCb interface{}, maxRecent int) {
	//void IupConfigRecentInit(Ihandle* ih, Ihandle* menu, Icallback recent_cb, int max_recent);
	C.__IupConfigRecentInit(ih.ptr(), menu.ptr(), unsafe.Pointer(syscall.NewCallbackCDecl(recentCb)), C.int(maxRecent))
}

func ConfigRecentUpdate(ih Ihandle, fileName string) {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	//void IupConfigRecentUpdate(Ihandle* ih, const char* filename);
	C.IupConfigRecentUpdate(ih.ptr(), c_fileName)
}

func ConfigDialogShow(ih, dialog Ihandle, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupConfigDialogShow(Ihandle* ih, Ihandle* dialog, const char* name);
	C.IupConfigDialogShow(ih.ptr(), dialog.ptr(), c_name)
}

func ConfigDialogClosed(ih, dialog Ihandle, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupConfigDialogClosed(Ihandle* ih, Ihandle* dialog, const char* name);
	C.IupConfigDialogClosed(ih.ptr(), dialog.ptr(), c_name)
}
