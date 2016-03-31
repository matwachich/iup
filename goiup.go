package iup

import (
	"fmt"
	"unsafe"
)

/*
#include <stdlib.h>
#include <iup.h>
*/
import "C"

// global attributes

//char* IupGetGlobal (const char* name);
func GetGlobalPtr(name string) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return uintptr(unsafe.Pointer(C.IupGetGlobal(c_name)))
}

//char* IupGetGlobal (const char* name);
func GetGlobalIh(name string) Ihandle {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return Ihandle(unsafe.Pointer(C.IupGetGlobal(c_name)))
}

// additional functions (mainly methods on Ihandle) for golang

//Go specific note: doesn't exists in C IUP.
func (ih Ihandle) GetIhandle(pih *Ihandle) Ihandle {
	if pih != nil {
		*pih = ih
	}
	return ih
}

func (ih Ihandle) SetHandle(name string) Ihandle {
	SetHandle(name, ih)
	return ih
}

func (ih Ihandle) Destroy() {
	Destroy(ih)
}

func (ih Ihandle) ResetAttribute(name string) Ihandle {
	ResetAttribute(ih, name)
	return ih
}

func (ih Ihandle) GetAllAttributes() []string {
	return GetAllAttributes(ih)
}

func (ih Ihandle) SetAttributes(params ...interface{}) Ihandle {
	for _, param := range params {
		switch param.(type) {
		case string:
			SetAttributes(ih, param.(string))
		case map[string]string:
			for key, val := range param.(map[string]string) {
				SetAttribute(ih, key, val)
			}
		case map[string]interface{}:
			for key, val := range param.(map[string]interface{}) {
				SetAttribute(ih, key, val)
			}
		}
	}
	return ih
}

func (ih Ihandle) GetAttributes() string {
	return GetAttributes(ih)
}

func (ih Ihandle) SetAttribute(name string, value ...interface{}) Ihandle {
	switch len(value) {
	case 1:
		SetAttribute(ih, name, value[0])
	case 2:
		SetAttribute(ih, name, fmt.Sprintf("%vx%v", value[0], value[1]))
	case 3:
		SetAttribute(ih, name, [3]byte{value[0].(byte), value[1].(byte), value[2].(byte)})
	default:
		panic(fmt.Errorf("bad argument passed to iup.Ihandle.SetAttribute"))
	}
	return ih
}

func (ih Ihandle) SetAttributeId(name string, id int, value ...interface{}) Ihandle {
	switch len(value) {
	case 1:
		SetAttributeId(ih, name, id, value[0])
	case 2:
		SetAttributeId(ih, name, id, fmt.Sprintf("%vx%v", value[0], value[1]))
	case 3:
		SetAttributeId(ih, name, id, [3]byte{value[0].(byte), value[1].(byte), value[2].(byte)})
	default:
		panic(fmt.Errorf("bad argument passed to iup.Ihandle.SetAttributeId"))
	}
	return ih
}

func (ih Ihandle) SetAttributeId2(name string, lin, col int, value ...interface{}) Ihandle {
	switch len(value) {
	case 1:
		SetAttributeId2(ih, name, lin, col, value[0])
	case 2:
		SetAttributeId2(ih, name, lin, col, fmt.Sprintf("%vx%v", value[0], value[1]))
	case 3:
		SetAttributeId2(ih, name, lin, col, [3]byte{value[0].(byte), value[1].(byte), value[2].(byte)})
	default:
		panic(fmt.Errorf("bad argument passed to iup.Ihandle.SetAttributeId2"))
	}
	return ih
}

func (ih Ihandle) GetAttribute(name string, ids ...interface{}) string {
	switch len(ids) {
	case 0:
		return GetAttribute(ih, name)
	case 1:
		return GetAttributeId(ih, name, ids[0].(int))
	case 2:
		return GetAttributeId2(ih, name, ids[0].(int), ids[1].(int))
	default:
		panic(fmt.Errorf("bad arguments passed to iup.Ihandle.GetAttribute"))
	}
}

func (ih Ihandle) GetInt(name string, ids ...interface{}) int {
	switch len(ids) {
	case 0:
		return GetInt(ih, name)
	case 1:
		return GetIntId(ih, name, ids[0].(int))
	case 2:
		return GetIntId2(ih, name, ids[0].(int), ids[1].(int))
	default:
		panic(fmt.Errorf("bad arguments passed to iup.Ihandle.GetInt"))
	}
}

func (ih Ihandle) GetInt2(name string) (count, i1, i2 int) { // count = 0, 1 or 2
	return GetInt2(ih, name)
}

func (ih Ihandle) GetFloat(name string, ids ...interface{}) float32 {
	switch len(ids) {
	case 0:
		return GetFloat(ih, name)
	case 1:
		return GetFloatId(ih, name, ids[0].(int))
	case 2:
		return GetFloatId2(ih, name, ids[0].(int), ids[1].(int))
	default:
		panic(fmt.Errorf("bad arguments passed to iup.Ihandle.GetFloat"))
	}
}

func (ih Ihandle) GetDouble(name string, ids ...interface{}) float64 {
	switch len(ids) {
	case 0:
		return GetDouble(ih, name)
	case 1:
		return GetDoubleId(ih, name, ids[0].(int))
	case 2:
		return GetDoubleId2(ih, name, ids[0].(int), ids[1].(int))
	default:
		panic(fmt.Errorf("bad arguments passed to iup.Ihandle.GetDouble"))
	}
}

func (ih Ihandle) GetRGB(name string, ids ...interface{}) (r, g, b uint8) {
	switch len(ids) {
	case 0:
		return GetRGB(ih, name)
	case 1:
		return GetRGBId(ih, name, ids[0].(int))
	case 2:
		return GetRGBId2(ih, name, ids[0].(int), ids[1].(int))
	default:
		panic(fmt.Errorf("bad arguments passed to iup.Ihandle.GetRGB"))
	}
}

func (ih Ihandle) GetPtr(name string, ids ...interface{}) uintptr {
	switch len(ids) {
	case 0:
		return GetPtr(ih, name)
	case 1:
		return GetPtrId(ih, name, ids[0].(int))
	case 2:
		return GetPtrId2(ih, name, ids[0].(int), ids[1].(int))
	default:
		panic(fmt.Errorf("bad arguments passed to iup.Ihandle.GetPtr"))
	}
}

func GetPtr(ih Ihandle, name string) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return uintptr(unsafe.Pointer(C.IupGetAttribute(ih.ptr(), c_name)))
}

func GetPtrId(ih Ihandle, name string, id int) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return uintptr(unsafe.Pointer(C.IupGetAttributeId(ih.ptr(), c_name, C.int(id))))
}

func GetPtrId2(ih Ihandle, name string, lin, col int) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return uintptr(unsafe.Pointer(C.IupGetAttributeId2(ih.ptr(), c_name, C.int(lin), C.int(col))))
}

func (ih Ihandle) SetCallback(name string, fn interface{}) Ihandle {
	SetCallback(ih, name, fn)
	return ih
}

func (ih Ihandle) GetCallback(name string) uintptr {
	return GetCallback(ih, name)
}

/* BUTTON_CB's status */

func ButtonCBStatus(p uintptr) (ret string) { // just for debugging
	s := CStrToString(p)
	if s[1] == 'C' {
		ret += "Ctrl+"
	}
	if s[0] == 'S' {
		ret += "Shift+"
	}
	if s[6] == 'A' {
		ret += "Alt+"
	}
	if s[7] == 'Y' {
		ret += "Sys+"
	}
	if s[2] == '1' {
		ret += "BUTTON_1"
	}
	if s[3] == '2' {
		ret += "BUTTON_2"
	}
	if s[4] == '3' {
		ret += "BUTTON_3"
	}
	if s[8] == '4' {
		ret += "BUTTON_4"
	}
	if s[9] == '5' {
		ret += "BUTTON_5"
	}
	if s[5] == 'D' {
		ret += " (Double)"
	}
	return
}
