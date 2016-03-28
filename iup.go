package iup

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

/*
#cgo CFLAGS: -m32 -I./include
#cgo LDFLAGS: -m32 -L${SRCDIR}/lib
#cgo LDFLAGS: -liup -liupcd -liupim -liupimglib -lgdi32 -lcomdlg32 -lcomctl32 -luuid -loleaut32 -lole32

#include <stdlib.h>
#include <string.h>
#include <iup.h>

Icallback __IupSetCallback(Ihandle* ih, const char *name, void* func) {
	return IupSetCallback(ih, name, (Icallback)func);
}

Icallback __IupSetFunction(const char *name, void* func) {
	return IupSetFunction(name, (Icallback)func);
}

*/
import "C"

// constants
const (
	NAME           = C.IUP_NAME
	DESCRIPTION    = C.IUP_DESCRIPTION
	COPYRIGHT      = C.IUP_COPYRIGHT
	VERSION        = C.IUP_VERSION /* bug fixes are reported only by IupVersion functions */
	VERSION_NUMBER = C.IUP_VERSION_NUMBER
	VERSION_DATE   = C.IUP_VERSION_DATE /* does not include bug fix releases */
)

//Ihandle type.
type Ihandle uintptr

// callbacks are different, so functions that accept callbacks take it as interface{}
//type Icallback func(Ihandle) int

func (ih Ihandle) ptr() *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(ih))
}
func mkih(p *C.Ihandle) Ihandle {
	return Ihandle(unsafe.Pointer(p))
}

/* ---------------------------------------------------------------------------------------------- */
/*                                          Main API                                              */
/* ---------------------------------------------------------------------------------------------- */

//Open initializes the IUP toolkit.
//Must be called before any other IUP function.
func Open() int {
	//int IupOpen (int *argc, char ***argv);
	//TODO implement custom args passing
	return int(C.IupOpen(&C._argc, &C._argv)) // from stdlib.h
}

//Close ends the IUP toolkit and releases internal memory.
//It will also automatically destroy all dialogs and all elements that have names.
func Close() {
	//void IupClose (void);
	C.IupClose()
}

//ImageLibOpen register the names but do not load the images.
//The images will be loaded only if they are used in a control.
//The loaded images will be automatically released at IupClose.
func ImageLibOpen() {
	//void IupImageLibOpen (void);
	C.IupImageLibOpen()
}

//MainLoop executes the user interaction until a callback returns IUP_CLOSE, IupExitLoop is called, or hiding the last visible dialog.
func MainLoop() int {
	//int IupMainLoop (void);
	return int(C.IupMainLoop())
}

//LoopStep runs one iteration of the message loop.
func LoopStep() int {
	//int IupLoopStep (void);
	return int(C.IupLoopStep())
}

//LoopStepWait runs one iteration of the message loop.
func LoopStepWait() int {
	//int IupLoopStepWait (void);
	return int(C.IupLoopStepWait())
}

//MainLoopLevel returns the current cascade level of IupMainLoop. When no calls were done, return value is 0.
func MainLoopLevel() int {
	//int IupMainLoopLevel (void);
	return int(C.IupMainLoopLevel())
}

//Flush processes all pending messages in the message queue.
func Flush() {
	//void IupFlush (void);
	C.IupFlush()
}

//ExitLoop terminates the current message loop. It has the same effect of a callback returning IUP_CLOSE.
func ExitLoop() {
	//void IupExitLoop (void);
	C.IupExitLoop()
}

//RecordInput records all mouse and keyboard input in a file for later reproduction.
func RecordInput(fileName string, mode int) int {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	//int IupRecordInput(const char* filename, int mode);
	return int(C.IupRecordInput(c_fileName, C.int(mode)))
}

//PlayInput reproduces all mouse and keyboard input from a given file.
func PlayInput(fileName string) int {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	//int IupPlayInput(const char* filename);
	return int(C.IupPlayInput(c_fileName))
}

//Update mark the element or its children to be redraw when the control returns to the system.
func Update(ih Ihandle) {
	//void IupUpdate (Ihandle* ih);
	C.IupUpdate(ih.ptr())
}

//UpdateChildren mark the element or its children to be redraw when the control returns to the system.
func UpdateChildren(ih Ihandle) {
	//void IupUpdateChildren(Ihandle* ih);
	C.IupUpdateChildren(ih.ptr())
}

//Redraw force the element and its children to be redraw immediately.
func Redraw(ih Ihandle, children int) {
	//void IupRedraw (Ihandle* ih, int children);
	C.IupRedraw(ih.ptr(), C.int(children))
}

//Refresh updates the size and layout of all controls in the same dialog.
//To be used after changing size attributes, or attributes that affect the size of the control.
//Can be used for any element inside a dialog, but the layout of the dialog and all controls will be updated.
//It can change the layout of all the controls inside the dialog because of the dynamic layout positioning.
func Refresh(ih Ihandle) {
	//void IupRefresh (Ihandle* ih);
	C.IupRefresh(ih.ptr())
}

//RefreshChildren updates the size and layout of controls after changing size attributes,
//or attributes that affect the size of the control. Can be used for any element inside a dialog,
//only its children will be updated. It can change the layout of all the controls inside
//the given element because of the dynamic layout positioning.
func RefreshChildren(ih Ihandle) {
	//void IupRefreshChildren(Ihandle* ih);
	C.IupRefreshChildren(ih.ptr())
}

//Execute runs the executable with the given parameters.
//It is a non synchronous operation, i.e. the function will return just after
//execute the command and it will not wait for its result.
//In Windows, there is no need to add the ".exe" file extension.
//Used by the IupHelp function.
func Execute(fileName, parameters string) int {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	//int IupExecute(const char *filename, const char* parameters);
	return int(C.IupExecute(c_fileName, C.CString(parameters)))
}

//Help opens the given URL. In UNIX executes Netscape, Safari (MacOS) or Firefox (in Linux) passing the desired URL as a parameter.
//In Windows executes the shell "open" operation on the given URL.
//In UNIX you can change the used browser setting the environment variable IUP_HELPAPP or using the global attribute "HELPAPP".
//It is a non synchronous operation, i.e. the function will return just after execute the command and it will not wait for its result.
//Since IUP 3.17, it will use the IupExecute function.
func Help(url string) int {
	c_url := C.CString(url)
	defer C.free(unsafe.Pointer(c_url))

	//int IupHelp(const char* url);
	return int(C.IupHelp(c_url))
}

//Load compiles a LED specification
func Load(fileName string) error {
	c_fileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_fileName))

	//char* IupLoad (const char *filename);
	ret := C.IupLoad(c_fileName)
	if ret == nil {
		return nil
	} else {
		return fmt.Errorf(C.GoString(ret))
	}
}

//LoadBuffer compiles a LED specification
func LoadBuffer(buffer string) error {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	//char* IupLoadBuffer (const char *buffer);
	ret := C.IupLoadBuffer(c_buffer)
	if ret == nil {
		return nil
	} else {
		return fmt.Errorf(C.GoString(ret))
	}
}

//Version returns a string with the IUP version number
func Version() string {
	//char* IupVersion (void);
	return C.GoString(C.IupVersion())
}

func VersionDate() string {
	//char* IupVersionDate (void);
	return C.GoString(C.IupVersionDate())
}

//VersionNumber returns a string with the IUP version number
func VersionNumber() int {
	//int IupVersionNumber (void);
	return int(C.IupVersionNumber())
}

//SetLanguage sets the language name used by some pre-defined dialogs.
//Can also be changed using the global attribute LANGUAGE.
func SetLanguage(lng string) {
	c_lng := C.CString(lng)
	defer C.free(unsafe.Pointer(c_lng))

	//void IupSetLanguage (const char *lng);
	C.IupSetLanguage(c_lng)
}

//GetLanguage returns the language used by some pre-defined dialogs.
//Returns the same value as the LANGUAGE global attribute.
func GetLanguage() string {
	//char* IupGetLanguage (void);
	return C.GoString(C.IupGetLanguage())
}

//SetLanguageString associates a name with a string as an auxiliary method for Internationalization of applications.
func SetLanguageString(name, str string) {
	c_name, c_str := C.CString(name), C.CString(str)
	defer C.free(unsafe.Pointer(c_name))
	defer C.free(unsafe.Pointer(c_str))

	//void IupSetLanguageString(const char* name, const char* str);
	//void IupStoreLanguageString(const char* name, const char* str);
	C.IupStoreLanguageString(c_name, c_str) //NOTE string always duplicated
}

//GetLanguageString returns a language dependent string.
//The string must have been associated with the name using the IupSetLanguageString or IupSetLanguagePack functions.
func GetLanguageString(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//char* IupGetLanguageString(const char* name);
	return C.GoString(C.IupGetLanguageString(c_name))
}

//SetLanguagePack sets a pack of associations between names and string values.
//It is simply a IupUser element with several attributes set.
//Internally will call IupSetLanguageString for each name in the pack.
func SetLanguagePack(ih Ihandle) {
	//void IupSetLanguagePack(Ihandle* ih);
	C.IupSetLanguagePack(ih.ptr())
}

//Destroy destroys an interface element and all its children.
//Only dialogs, timers, popup menus and images should be normally destroyed, but detached controls can also be destroyed.
func Destroy(ih Ihandle) {
	//void IupDestroy (Ihandle* ih);
	C.IupDestroy(ih.ptr())
}

//Detach detaches an interface element from its parent.
func Detach(child Ihandle) {
	//void IupDetach (Ihandle* child);
	C.IupDetach(child.ptr())
}

//Append inserts an interface element at the end of the container, after the last element of the container.
//Valid for any element that contains other elements like dialog, frame, hbox, vbox, zbox or menu.
func Append(ih, child Ihandle) Ihandle {
	//Ihandle* IupAppend (Ihandle* ih, Ihandle* child);
	return mkih(C.IupAppend(ih.ptr(), child.ptr()))
}

//Insert Inserts an interface element before another child of the container.
//Valid for any element that contains other elements like dialog, frame, hbox, vbox, zbox, menu, etc.
func Insert(ih, refChild, child Ihandle) Ihandle {
	//Ihandle* IupInsert (Ihandle* ih, Ihandle* ref_child, Ihandle* child);
	return mkih(C.IupInsert(ih.ptr(), refChild.ptr(), child.ptr()))
}

//GetChild returns the a child of the control given its position.
func GetChild(ih Ihandle, pos int) Ihandle {
	//Ihandle* IupGetChild (Ihandle* ih, int pos);
	return mkih(C.IupGetChild(ih.ptr(), C.int(pos)))
}

//GetChildPos returns the position of a child of the given control.
func GetChildPos(ih, child Ihandle) int {
	//int IupGetChildPos (Ihandle* ih, Ihandle* child);
	return int(C.IupGetChildPos(ih.ptr(), child.ptr()))
}

//GetChildCount returns the number of children of the given control.
func GetChildCount(ih Ihandle) int {
	//int IupGetChildCount(Ihandle* ih);
	return int(C.IupGetChildCount(ih.ptr()))
}

//GetNextChild returns the a child of the given control given its brother.
func GetNextChild(ih, child Ihandle) Ihandle {
	//Ihandle* IupGetNextChild (Ihandle* ih, Ihandle* child);
	return mkih(C.IupGetNextChild(ih.ptr(), child.ptr()))
}

//GetBrother returns the brother of an element.
func GetBrother(ih Ihandle) Ihandle {
	//Ihandle* IupGetBrother (Ihandle* ih);
	return mkih(C.IupGetBrother(ih.ptr()))
}

//GetParent returns the parent of a control.
func GetParent(ih Ihandle) Ihandle {
	//Ihandle* IupGetParent (Ihandle* ih);
	return mkih(C.IupGetParent(ih.ptr()))
}

//GetDialog returns the handle of the dialog that contains that interface element.
//Works also for children of a menu that is associated with a dialog.
func GetDialog(ih Ihandle) Ihandle {
	//Ihandle* IupGetDialog (Ihandle* ih);
	return mkih(C.IupGetDialog(ih.ptr()))
}

//GetDialogChild returns the identifier of the child element that has the NAME attribute
//equals to the given value on the same dialog hierarchy.
//Works also for children of a menu that is associated with a dialog.
func GetDialogChild(ih Ihandle, name string) Ihandle {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Ihandle* IupGetDialogChild(Ihandle* ih, const char* name);
	return mkih(C.IupGetDialogChild(ih.ptr(), c_name))
}

//Reparent moves an interface element from one position in the hierarchy tree to another.
//Both new_parent and child must be mapped or unmapped at the same time.
//If ref_child is NULL, then it will append the child to the new_parent.
//If ref_child is NOT NULL then it will insert child before ref_child inside the new_parent.
func Reparent(ih, newParent, refChild Ihandle) int {
	//int IupReparent (Ihandle* ih, Ihandle* new_parent, Ihandle* ref_child);
	return int(C.IupReparent(ih.ptr(), newParent.ptr(), refChild.ptr()))
}

//Popup shows a dialog or menu and restricts user interaction only to the specified element.
//It is equivalent of creating a Modal dialog is some toolkits.
//
//If another dialog is shown after IupPopup using IupShow, then its interaction will not be inhibited.
//Every IupPopup call creates a new popup level that inhibits all previous dialogs interactions, but does not disable new ones
//(even if they were disabled by the Popup, calling IupShow will re-enable the dialog because it will change its popup level).
//IMPORTANT: The popup levels must be closed in the reverse order they were created or unpredictable results will occur.
//
//For a dialog this function will only return the control to the application
//after a callback returns IUP_CLOSE, IupExitLoop is called, or when the popup dialog is hidden, for example using IupHide.
//For a menu it returns automatically after a menu item is selected.
//IMPORTANT: If a menu item callback returns IUP_CLOSE, it will also ends the current popup level dialog.
func Popup(ih Ihandle, x, y int) int {
	//int IupPopup (Ihandle* ih, int x, int y);
	return int(C.IupPopup(ih.ptr(), C.int(x), C.int(y)))
}

//Show displays a dialog in the current position, or changes a control VISIBLE attribute.
func Show(ih Ihandle) int {
	//int IupShow (Ihandle* ih);
	return int(C.IupShow(ih.ptr()))
}

//ShowXY displays a dialog in a given position on the screen.
func ShowXY(ih Ihandle, x, y int) int {
	//int IupShowXY (Ihandle* ih, int x, int y);
	return int(C.IupShowXY(ih.ptr(), C.int(x), C.int(y)))
}

//Hide hides an interface element.
//This function has the same effect as attributing value "NO" to the interface element’s VISIBLE attribute.
func Hide(ih Ihandle) int {
	//int IupHide (Ihandle* ih);
	return int(C.IupHide(ih.ptr()))
}

//Map creates (maps) the native interface objects corresponding to the given IUP interface elements.
//
//It will also called recursively to create the native element of all the children in the element's tree.
//
//The element must be already attached to a mapped container, except the dialog.
//A child can only be mapped if its parent is already mapped.
//
//This function is automatically called before the dialog is shown in IupShow, IupShowXY or IupPopup.
func Map(ih Ihandle) int {
	//int IupMap (Ihandle* ih);
	return int(C.IupMap(ih.ptr()))
}

//Unmap unmap the element from the native system. It will also unmap all its children.
//It will NOT detach the element from its parent, and it will NOT destroy the IUP element.
func Unmap(ih Ihandle) {
	//void IupUnmap (Ihandle *ih);
	C.IupUnmap(ih.ptr())
}

//ResetAttributes removes an attribute from the hash table of the element, and its children if the attribute is inheritable.
//It is useful to reset the state of inheritable attributes in a tree of elements.
func ResetAttribute(ih Ihandle, name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupResetAttribute(Ihandle *ih, const char* name);
	C.IupResetAttribute(ih.ptr(), c_name)
}

//GetAllAttributes returns the names of all attributes of an element that are set in its internal hash table only.
func GetAllAttributes(ih Ihandle) (ret []string) {
	n := int(C.IupGetAllAttributes(ih.ptr(), nil, 0))
	if n > 0 {
		ret = make([]string, n)
		pRets := make([]*C.char, n)
		//int IupGetAllAttributes(Ihandle* ih, char** names, int n);
		C.IupGetAllAttributes(ih.ptr(), (**C.char)(unsafe.Pointer(&pRets[0])), C.int(n))
		for i := 0; i < n; i++ {
			ret[i] = C.GoString(pRets[i])
		}
	}
	return
}

//Ihandle* IupSetAtt(const char* handle_name, Ihandle* ih, const char* name, ...);

//SetAttributes sets several attributes of an interface element.
func SetAttributes(ih Ihandle, str string) Ihandle {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	//Ihandle* IupSetAttributes (Ihandle* ih, const char *str);
	return mkih(C.IupSetAttributes(ih.ptr(), c_str))
}

//GetAttributes returns all attributes of a given element that are set in the internal hash table.
//The known attributes that are pointers (not strings) are returned as integers.
//
//The internal attributes are not returned (attributes prefixed with "_IUP").
//
//Before calling this function the application must ensure that there is no pointer attributes
//set for that element, although all known pointers attributes are handled.
//
//This function should be avoided. Use iup.GetAllAttributes instead.
func GetAttributes(ih Ihandle) string {
	//char* IupGetAttributes (Ihandle* ih);
	return C.GoString(C.IupGetAttributes(ih.ptr()))
}

//SetAttribute sets an interface element attribute.
func SetAttribute(ih Ihandle, name string, value interface{}) { //NOTE string attribute is always copied
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetAttribute (Ihandle* ih, const char* name, const char* value);
	//void IupSetStrAttribute(Ihandle* ih, const char* name, const char* value);
	//void IupSetStrf (Ihandle* ih, const char* name, const char* format, ...);
	//void IupSetInt (Ihandle* ih, const char* name, int value);
	//void IupSetFloat (Ihandle* ih, const char* name, float value);
	//void IupSetDouble (Ihandle* ih, const char* name, double value);

	switch value.(type) {
	case nil:
		C.IupSetAttribute(ih.ptr(), c_name, nil) // default value
	case Ihandle:
		C.IupSetAttribute(ih.ptr(), c_name, (*C.char)(unsafe.Pointer(value.(Ihandle)))) // store pointer
	case uintptr:
		C.IupSetAttribute(ih.ptr(), c_name, (*C.char)(unsafe.Pointer(value.(uintptr)))) // store pointer
	case string:
		c_value := C.CString(value.(string))
		defer C.free(unsafe.Pointer(c_value))

		C.IupSetStrAttribute(ih.ptr(), c_name, c_value) // always copy string
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		C.IupSetInt(ih.ptr(), c_name, C.int(reflect.ValueOf(value).Int()))
	case float32:
		C.IupSetFloat(ih.ptr(), c_name, C.float(value.(float32)))
	case float64:
		C.IupSetDouble(ih.ptr(), c_name, C.double(value.(float64)))
	case [3]uint8:
		C.IupSetRGB(ih.ptr(), c_name, C.uchar(value.([3]uint8)[0]), C.uchar(value.([3]uint8)[1]), C.uchar(value.([3]uint8)[2]))
	default:
		panic(fmt.Errorf("bad argument passed to iup.SetAttribute"))
	}
}

//SetRGB sets an interface element attribute.
func SetRGB(ih Ihandle, name string, r, g, b uint8) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetRGB (Ihandle *ih, const char* name, unsigned char r, unsigned char g, unsigned char b);
	C.IupSetRGB(ih.ptr(), c_name, C.uchar(r), C.uchar(g), C.uchar(b))
}

//GetAttribute returns the name of an interface element attribute.
func GetAttribute(ih Ihandle, name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//char* IupGetAttribute(Ihandle* ih, const char* name);
	return C.GoString(C.IupGetAttribute(ih.ptr(), c_name))
}

//GetInt returns the name of an interface element attribute.
func GetInt(ih Ihandle, name string) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//int IupGetInt (Ihandle* ih, const char* name);
	return int(C.IupGetInt(ih.ptr(), c_name))
}

//GetInt2 returns the name of an interface element attribute.
func GetInt2(ih Ihandle, name string) (count, i1, i2 int) { // count = 0, 1 or 2
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//int IupGetInt2 (Ihandle* ih, const char* name);
	//int IupGetIntInt (Ihandle *ih, const char* name, int *i1, int *i2);
	count = int(C.IupGetIntInt(ih.ptr(), c_name, (*C.int)(unsafe.Pointer(&i1)), (*C.int)(unsafe.Pointer(&i2))))
	return
}

//GetFloat returns the name of an interface element attribute.
func GetFloat(ih Ihandle, name string) float32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//float IupGetFloat (Ihandle* ih, const char* name);
	return float32(C.IupGetFloat(ih.ptr(), c_name))
}

//GetDouble returns the name of an interface element attribute.
func GetDouble(ih Ihandle, name string) float64 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//double IupGetDouble (Ihandle* ih, const char* name);
	return float64(C.IupGetDouble(ih.ptr(), c_name))
}

//GetRGB returns the name of an interface element attribute.
func GetRGB(ih Ihandle, name string) (r, g, b uint8) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupGetRGB (Ihandle *ih, const char* name, unsigned char *r, unsigned char *g, unsigned char *b);
	C.IupGetRGB(ih.ptr(), c_name, (*C.uchar)(unsafe.Pointer(&r)), (*C.uchar)(unsafe.Pointer(&g)), (*C.uchar)(unsafe.Pointer(&b)))
	return
}

//SetAttributeId sets an interface element attribute.
func SetAttributeId(ih Ihandle, name string, id int, value interface{}) { //NOTE string attribute is always copied
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetAttributeId(Ihandle *ih, const char* name, int id, const char *value);
	//void IupSetStrAttributeId(Ihandle *ih, const char* name, int id, const char *value);
	//void IupSetStrfId(Ihandle *ih, const char* name, int id, const char* format, ...);
	//void IupSetIntId(Ihandle* ih, const char* name, int id, int value);
	//void IupSetFloatId(Ihandle* ih, const char* name, int id, float value);
	//void IupSetDoubleId(Ihandle* ih, const char* name, int id, double value);

	switch value.(type) {
	case nil:
		C.IupSetAttributeId(ih.ptr(), c_name, C.int(id), nil) // default value
	case Ihandle:
		C.IupSetAttributeId(ih.ptr(), c_name, C.int(id), (*C.char)(unsafe.Pointer(value.(Ihandle)))) // store pointer
	case uintptr:
		C.IupSetAttributeId(ih.ptr(), c_name, C.int(id), (*C.char)(unsafe.Pointer(value.(uintptr)))) // store pointer
	case string:
		c_value := C.CString(value.(string))
		defer C.free(unsafe.Pointer(c_value))

		C.IupSetStrAttributeId(ih.ptr(), c_name, C.int(id), c_value) // always copy string
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		C.IupSetIntId(ih.ptr(), c_name, C.int(id), C.int(reflect.ValueOf(value).Int()))
	case float32:
		C.IupSetFloatId(ih.ptr(), c_name, C.int(id), C.float(value.(float32)))
	case float64:
		C.IupSetDoubleId(ih.ptr(), c_name, C.int(id), C.double(value.(float64)))
	case [3]uint8:
		C.IupSetRGBId(ih.ptr(), c_name, C.int(id), C.uchar(value.([3]uint8)[0]), C.uchar(value.([3]uint8)[1]), C.uchar(value.([3]uint8)[2]))
	default:
		panic(fmt.Errorf("bad argument passed to iup.SetAttributeId"))
	}
}

//SetRGBId sets an interface element attribute.
func SetRGBId(ih Ihandle, name string, id int, r, g, b uint8) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetRGBId(Ihandle *ih, const char* name, int id, unsigned char r, unsigned char g, unsigned char b);
	C.IupSetRGBId(ih.ptr(), c_name, C.int(id), C.uchar(r), C.uchar(g), C.uchar(b))
}

//GetAttributeId returns the name of an interface element attribute.
func GetAttributeId(ih Ihandle, name string, id int) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//char* IupGetAttributeId(Ihandle *ih, const char* name, int id);
	return C.GoString(C.IupGetAttributeId(ih.ptr(), c_name, C.int(id)))
}

//GetIntId returns the name of an interface element attribute.
func GetIntId(ih Ihandle, name string, id int) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//int IupGetIntId(Ihandle *ih, const char* name, int id);
	return int(C.IupGetIntId(ih.ptr(), c_name, C.int(id)))
}

//GetFloatId returns the name of an interface element attribute.
func GetFloatId(ih Ihandle, name string, id int) float32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//float IupGetFloatId(Ihandle *ih, const char* name, int id);
	return float32(C.IupGetFloatId(ih.ptr(), c_name, C.int(id)))
}

//GetDoubleId returns the name of an interface element attribute.
func GetDoubleId(ih Ihandle, name string, id int) float64 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//double IupGetDoubleId(Ihandle *ih, const char* name, int id);
	return float64(C.IupGetDoubleId(ih.ptr(), c_name, C.int(id)))
}

//GetRGBId returns the name of an interface element attribute.
func GetRGBId(ih Ihandle, name string, id int) (r, g, b uint8) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupGetRGBId(Ihandle *ih, const char* name, int id, unsigned char *r, unsigned char *g, unsigned char *b);
	C.IupGetRGBId(ih.ptr(), c_name, C.int(id), (*C.uchar)(unsafe.Pointer(&r)), (*C.uchar)(unsafe.Pointer(&g)), (*C.uchar)(unsafe.Pointer(&b)))
	return
}

//SetAttributeId2 sets an interface element attribute.
func SetAttributeId2(ih Ihandle, name string, lin, col int, value interface{}) { //NOTE string attribute is always copied
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetAttributeId2(Ihandle* ih, const char* name, int lin, int col, const char* value);
	//void IupSetStrAttributeId2(Ihandle* ih, const char* name, int lin, int col, const char* value);
	//void IupSetStrfId2(Ihandle* ih, const char* name, int lin, int col, const char* format, ...);
	//void IupSetIntId2(Ihandle* ih, const char* name, int lin, int col, int value);
	//void IupSetFloatId2(Ihandle* ih, const char* name, int lin, int col, float value);
	//void IupSetDoubleId2(Ihandle* ih, const char* name, int lin, int col, double value);

	switch value.(type) {
	case nil:
		C.IupSetAttributeId2(ih.ptr(), c_name, C.int(lin), C.int(col), nil) // default value
	case Ihandle:
		C.IupSetAttributeId2(ih.ptr(), c_name, C.int(lin), C.int(col), (*C.char)(unsafe.Pointer(value.(Ihandle)))) // store pointer
	case uintptr:
		C.IupSetAttributeId2(ih.ptr(), c_name, C.int(lin), C.int(col), (*C.char)(unsafe.Pointer(value.(uintptr)))) // store pointer
	case string:
		c_value := C.CString(value.(string))
		defer C.free(unsafe.Pointer(c_value))

		C.IupSetStrAttributeId2(ih.ptr(), c_name, C.int(lin), C.int(col), c_value) // always copy string
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		C.IupSetIntId2(ih.ptr(), c_name, C.int(lin), C.int(col), C.int(reflect.ValueOf(value).Int()))
	case float32:
		C.IupSetFloatId2(ih.ptr(), c_name, C.int(lin), C.int(col), C.float(value.(float32)))
	case float64:
		C.IupSetDoubleId2(ih.ptr(), c_name, C.int(lin), C.int(col), C.double(value.(float64)))
	case [3]uint8:
		C.IupSetRGBId2(ih.ptr(), c_name, C.int(lin), C.int(col), C.uchar(value.([3]uint8)[0]), C.uchar(value.([3]uint8)[1]), C.uchar(value.([3]uint8)[2]))
	default:
		panic(fmt.Errorf("bad argument passed to iup.SetAttributeId2"))
	}
}

//SetRGBId2 sets an interface element attribute.
func SetRGBId2(ih Ihandle, name string, lin, col int, r, g, b uint8) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetRGBId2(Ihandle *ih, const char* name, int lin, int col, unsigned char r, unsigned char g, unsigned char b);
	C.IupSetRGBId2(ih.ptr(), c_name, C.int(lin), C.int(col), C.uchar(r), C.uchar(g), C.uchar(b))
}

//GetAttributeId2 returns the name of an interface element attribute.
func GetAttributeId2(ih Ihandle, name string, lin, col int) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//char* IupGetAttributeId2(Ihandle* ih, const char* name, int lin, int col);
	return C.GoString(C.IupGetAttributeId2(ih.ptr(), c_name, C.int(lin), C.int(col)))
}

//GetIntId2 returns the name of an interface element attribute.
func GetIntId2(ih Ihandle, name string, lin, col int) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//int IupGetIntId2(Ihandle* ih, const char* name, int lin, int col);
	return int(C.IupGetIntId2(ih.ptr(), c_name, C.int(lin), C.int(col)))
}

//GetRGBId returns the name of an interface element attribute.
func GetFloatId2(ih Ihandle, name string, lin, col int) float32 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//float IupGetFloatId2(Ihandle* ih, const char* name, int lin, int col);
	return float32(C.IupGetFloatId2(ih.ptr(), c_name, C.int(lin), C.int(col)))
}

//GetDoubleId2 returns the name of an interface element attribute.
func GetDoubleId2(ih Ihandle, name string, lin, col int) float64 {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//double IupGetDoubleId2(Ihandle* ih, const char* name, int lin, int col);
	return float64(C.IupGetDoubleId2(ih.ptr(), c_name, C.int(lin), C.int(col)))
}

//GetRGBId2 returns the name of an interface element attribute.
func GetRGBId2(ih Ihandle, name string, lin, col int) (r, g, b uint8) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupGetRGBId2(Ihandle *ih, const char* name, int lin, int col, unsigned char *r, unsigned char *g, unsigned char *b);
	C.IupGetRGBId2(ih.ptr(), c_name, C.int(lin), C.int(col), (*C.uchar)(unsafe.Pointer(&r)), (*C.uchar)(unsafe.Pointer(&g)), (*C.uchar)(unsafe.Pointer(&b)))
	return
}

//SetGlobal sets an attribute in the global environment.
//If the driver process the attribute then it will not be stored internally.
func SetGlobal(name string, value interface{}) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetGlobal (const char* name, const char* value);
	//void IupSetStrGlobal(const char* name, const char* value);

	switch value.(type) { //TODO handle number values?
	case string:
		c_value := C.CString(value.(string))
		defer C.free(unsafe.Pointer(c_value))

		C.IupSetStrGlobal(c_name, c_value) // always copy value
	case uintptr:
		C.IupSetGlobal(c_name, (*C.char)(unsafe.Pointer(value.(uintptr))))
	case Ihandle:
		C.IupSetGlobal(c_name, (*C.char)(unsafe.Pointer(value.(Ihandle))))
	default:
		panic(fmt.Errorf("bad argument passed to iup.SetGlobal"))
	}
}

//GetGlobal returns an attribute value from the global environment.
//The value can be returned from the driver or from the internal storage.
func GetGlobal(name string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//char* IupGetGlobal (const char* name);
	return C.GoString(C.IupGetGlobal(c_name))
}

//SetFocus sets the interface element that will receive the keyboard focus, i.e., the element that will receive keyboard events.
//But this will be processed only after the con
func SetFocus(ih Ihandle) Ihandle {
	//Ihandle* IupSetFocus (Ihandle* ih);
	return mkih(C.IupSetFocus(ih.ptr()))
}

//GetFocus returns the identifier of the interface element that has the keyboard focus, i.e. the element that will receive keyboard events.
func GetFocus() Ihandle {
	//Ihandle* IupGetFocus (void);
	return mkih(C.IupGetFocus())
}

//PreviousField shifts the focus to the previous element that can have the focus.
//It is relative to the given element and does not depend on the element currently with the focus.
func PreviousField(ih Ihandle) Ihandle {
	//Ihandle* IupPreviousField(Ihandle* ih);
	return mkih(C.IupPreviousField(ih.ptr()))
}

//Shifts the focus to the next element that can have the focus.
//It is relative to the given element and does not depend on the element currently with the focus.
//
//It will search for the next element first in the children, then in the brothers,
//then in the uncles and their children, and so on.
//
//This sequence is not the same sequence used by the Tab key, which is dependent on the native system.
func NextField(ih Ihandle) Ihandle {
	//Ihandle* IupNextField (Ihandle* ih);
	return mkih(C.IupNextField(ih.ptr()))
}

//GetCallback returns the callback associated to an event.
//
//Go specific note: The value returned is not a Go function, it is a pointer to a C function.
func GetCallback(ih Ihandle, name string) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Icallback IupGetCallback (Ihandle* ih, const char *name);
	return uintptr(unsafe.Pointer(C.IupGetCallback(ih.ptr(), c_name)))
}

//SetCallback associates a callback to an event.
//
//Go specific note: You can pass a Go function or a pointer to a C function (as returned by iup.GetCallback for example)
func SetCallback(ih Ihandle, name string, fn interface{}) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Icallback IupSetCallback (Ihandle* ih, const char *name, Icallback func);
	//Ihandle* IupSetCallbacks(Ihandle* ih, const char *name, Icallback func, ...);

	switch fn.(type) {
	case uintptr:
		return uintptr(unsafe.Pointer(C.__IupSetCallback(ih.ptr(), c_name, unsafe.Pointer(fn.(uintptr)))))
	default:
		return uintptr(unsafe.Pointer(C.__IupSetCallback(ih.ptr(), c_name, unsafe.Pointer(syscall.NewCallbackCDecl(fn)))))
	}
}

//GetFunction returns the function associated to an action only when they were set by IupSetFunction.
//It will not work if IupSetCallback were used.
//
//Go specific note: The value returned is not a Go function, it is a pointer to a C function.
func GetFunction(name string) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Icallback IupGetFunction(const char *name);
	return uintptr(unsafe.Pointer(C.IupGetFunction(c_name)))
}

//SetFunction associates a function to an action as a global callback.
//
//This function should not be used by new applications, use it only for global callbacks.
//For regular elements use IupSetCallback instead.
//
//Notice that the application or libraries may set the same name for two different functions by mistake.
//IupSetCallback does not depends on global names.
//
//Go specific note: You can pass a Go function or a pointer to a C function (as returned by iup.GetCallback for example)
func SetFunction(name string, fn interface{}) uintptr {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Icallback IupSetFunction(const char *name, Icallback func);

	switch fn.(type) {
	case uintptr:
		return uintptr(unsafe.Pointer(C.__IupSetFunction(c_name, unsafe.Pointer(fn.(uintptr)))))
	default:
		return uintptr(unsafe.Pointer(C.__IupSetFunction(c_name, unsafe.Pointer(syscall.NewCallbackCDecl(fn)))))
	}
}

//GetHandle returns the identifier of an interface element that has an associated name using IupSetHandle or using LED.
func GetHandle(name string) Ihandle {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Ihandle* IupGetHandle (const char *name);
	return mkih(C.IupGetHandle(c_name))
}

//SetHandle associates a name with an interface element.
func SetHandle(name string, ih Ihandle) Ihandle {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Ihandle* IupSetHandle (const char *name, Ihandle* ih);
	return mkih(C.IupSetHandle(c_name, ih.ptr()))
}

//GetAllNames returns the names of all interface elements that have an associated name using IupSetHandle or using LED.
func GetAllNames() (names []string) {
	//int IupGetAllNames (char** names, int n);
	n := int(C.IupGetAllNames(nil, 0))
	if n > 0 {
		names = make([]string, n)
		pNames := make([]*C.char, n)
		C.IupGetAllNames((**C.char)(unsafe.Pointer(&pNames[0])), C.int(n))
		for i := 0; i < n; i++ {
			names[i] = C.GoString(pNames[i])
		}
	}
	return
}

//GetAllDialogs returns the names of all dialogs that have an associated name using IupSetHandle or using LED.
//Other dialogs will not be returned.
func GetAllDialogs() (names []string) {
	//int IupGetAllDialogs(char** names, int n);
	n := int(C.IupGetAllDialogs(nil, 0))
	if n > 0 {
		names = make([]string, n)
		pNames := make([]*C.char, n)
		C.IupGetAllDialogs((**C.char)(unsafe.Pointer(&pNames[0])), C.int(n))
		for i := 0; i < n; i++ {
			names[i] = C.GoString(pNames[i])
		}
	}
	return
}

//GetName Returns a name of an interface element, if the element has an associated name using IupSetHandle or using LED (which calls IupSetHandle when parsed).
//
//Notice that a handle can have many names. IupGetName will return the last name set.
func GetName(ih Ihandle) string {
	//char* IupGetName (Ihandle* ih);
	return C.GoString(C.IupGetName(ih.ptr()))
}

//SetAttributeHandle instead of using IupSetHandle and IupSetAttribute with a new creative name,
//this function automatically creates a non conflict name and associates the name with the attribute.
//
//It is very useful for associating images and menus.
func SetAttributeHandle(ih Ihandle, name string, ihNamed Ihandle) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupSetAttributeHandle(Ihandle* ih, const char* name, Ihandle* ih_named);
	C.IupSetAttributeHandle(ih.ptr(), c_name, ihNamed.ptr())
}

//GetAttributeHandle instead of using IupGetAttribute and IupGetHandle, this function directly returns the associated handle.
func GetAttributeHandle(ih Ihandle, name string) Ihandle {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//Ihandle* IupGetAttributeHandle(Ihandle* ih, const char* name);
	return mkih(C.IupGetAttributeHandle(ih.ptr(), c_name))
}

//GetClassName returns the name of the class of an interface element.
func GetClassName(ih Ihandle) string {
	//char* IupGetClassName(Ihandle* ih);
	return C.GoString(C.IupGetClassName(ih.ptr()))
}

//GetClassType returns the name of the native type of an interface element.
func GetClassType(ih Ihandle) string {
	//char* IupGetClassType(Ihandle* ih);
	return C.GoString(C.IupGetClassType(ih.ptr()))
}

//GetAllClasses returns the names of all registered classes.
func GetAllClasses() (names []string) {
	//int IupGetAllClasses(char** names, int n);
	n := int(C.IupGetAllClasses(nil, 0))
	if n > 0 {
		names = make([]string, n)
		pNames := make([]*C.char, n)
		C.IupGetAllClasses((**C.char)(unsafe.Pointer(&pNames[0])), C.int(n))
		for i := 0; i < n; i++ {
			names[i] = C.GoString(pNames[i])
		}
	}
	return
}

//GetClassAttributes returns the names of all registered attributes of a class.
func GetClassAttributes(className string) (names []string) {
	c_className := C.CString(className)
	defer C.free(unsafe.Pointer(c_className))

	//int IupGetClassAttributes(const char* classname, char** names, int n);
	n := int(C.IupGetClassAttributes(c_className, nil, 0))
	if n > 0 {
		names = make([]string, n)
		pNames := make([]*C.char, n)
		C.IupGetClassAttributes(c_className, (**C.char)(unsafe.Pointer(&pNames[0])), C.int(n))
		for i := 0; i < n; i++ {
			names[i] = C.GoString(pNames[i])
		}
	}
	return
}

//GetClassCallbacks returns the names of all registered callbacks of a class.
func GetClassCallbacks(className string) (names []string) {
	c_className := C.CString(className)
	defer C.free(unsafe.Pointer(c_className))

	//int IupGetClassCallbacks(const char* classname, char** names, int n);
	n := int(C.IupGetClassCallbacks(c_className, nil, 0))
	if n > 0 {
		names = make([]string, n)
		pNames := make([]*C.char, n)
		C.IupGetClassCallbacks(c_className, (**C.char)(unsafe.Pointer(&pNames[0])), C.int(n))
		for i := 0; i < n; i++ {
			names[i] = C.GoString(pNames[i])
		}
	}
	return
}

//SaveClassAttributes saves all registered attributes on the internal hash table.
func SaveClassAttributes(ih Ihandle) {
	//void IupSaveClassAttributes(Ihandle* ih);
	C.IupSaveClassAttributes(ih.ptr())
}

//CopyClassAttributes copies all registered attributes from one element to another.
//Both elements must be of the same class.
func CopyClassAttributes(srcIh, dstIh Ihandle) {
	//void IupCopyClassAttributes(Ihandle* src_ih, Ihandle* dst_ih);
	C.IupCopyClassAttributes(srcIh.ptr(), dstIh.ptr())
}

//SetClassDefaultAttributes changes the default value of an attribute for a class.
//It can be any attribute, i.e. registered attributes or user custom attributes.
func SetClassDefaultAttributes(className, name, value string) {
	c_className, c_name, c_value := C.CString(className), C.CString(name), C.CString(value)
	defer C.free(unsafe.Pointer(c_className))
	defer C.free(unsafe.Pointer(c_name))
	defer C.free(unsafe.Pointer(c_value))

	//void IupSetClassDefaultAttribute(const char* classname, const char *name, const char* value);
	C.IupSetClassDefaultAttribute(c_className, c_name, c_value)
}

//ClassMatch checks if the give class name matches the class name of the given interface element.
func ClassMatch(ih Ihandle, className string) bool {
	c_className := C.CString(className)
	defer C.free(unsafe.Pointer(c_className))

	//int IupClassMatch(Ihandle* ih, const char* classname);
	return C.IupClassMatch(ih.ptr(), c_className) != C.int(0)
}

//Create creates an interface element given its class name and parameters.
//This function is called from all constructors like IupDialog(...), IupLabel(...), and so on.
//
//After creation the element still needs to be attached to a container and mapped to the native system so it can be visible.
func Create(className string) Ihandle {
	c_className := C.CString(className)
	defer C.free(unsafe.Pointer(c_className))

	//Ihandle* IupCreate (const char *classname);
	return mkih(C.IupCreate(c_className))
}

//TODO implement?
//Ihandle* IupCreatev(const char *classname, void* *params);
//Ihandle* IupCreatep(const char *classname, void *first, ...);

/* ---------------------------------------------------------------------------------------------- */
/*                                          Elements                                              */
/* ---------------------------------------------------------------------------------------------- */

//Fill creates void element, which dynamically occupies empty spaces always trying to expand itself.
//Its parent should be an IupHbox, an IupVbox or a IupGridBox, or else this type of expansion will not work.
//If an EXPAND is set on at least one of the other children of the box, then the fill expansion is ignored.
//
//It does not have a native representation.
func Fill() Ihandle {
	//Ihandle* IupFill (void);
	return mkih(C.IupFill())
}

//Radio creates a void container for grouping mutual exclusive toggles.
//Only one of its descendent toggles will be active at a time.
//The toggles can be at any composition.
//
//It does not have a native representation.
func Radio(child Ihandle) Ihandle {
	//Ihandle* IupRadio (Ihandle* child);
	return mkih(C.IupRadio(child.ptr()))
}

//Vbox creates a void container for composing elements vertically.
//It is a box that arranges the elements it contains from top to bottom.
//
//It does not have a native representation.
func Vbox(children ...Ihandle) Ihandle {
	children = append(children, Ihandle(0))

	//Ihandle* IupVbox (Ihandle* child, ...);
	//Ihandle* IupVboxv (Ihandle* *children);
	return mkih(C.IupVboxv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//Zbox Creates a void container for composing elements in hidden layers with only one layer visible.
//It is a box that piles up the children it contains, only the one child is visible.
//
//It does not have a native representation.
//
//Zbox works by changing the VISIBLE attribute of its children, so if any of the grand children
//has its VISIBLE attribute directly defined then Zbox will NOT change its state.
func Zbox(children ...Ihandle) Ihandle {
	children = append(children, Ihandle(0))

	//Ihandle* IupZbox (Ihandle* child, ...);
	//Ihandle* IupZboxv (Ihandle* *children);
	return mkih(C.IupZboxv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//Hbox creates a void container for composing elements horizontally.
//It is a box that arranges the elements it contains from left to right.
//
//It does not have a native representation.
func Hbox(children ...Ihandle) Ihandle {
	children = append(children, Ihandle(0))

	//Ihandle* IupHbox (Ihandle* child,...);
	//Ihandle* IupHboxv (Ihandle* *children);
	return mkih(C.IupHboxv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//Normalizer creates a void container that does not affect the dialog layout.
//It acts by normalizing all the controls in a list so their natural size becomes the biggest natural size amongst them.
//All natural widths will be set to the biggest width, and all natural heights will be set to the biggest height.
//The controls of the list must be inside a valid container in the dialog.
func Normalizer(ihList ...Ihandle) Ihandle {
	ihList = append(ihList, Ihandle(0))

	//Ihandle* IupNormalizer (Ihandle* ih_first, ...);
	//Ihandle* IupNormalizerv(Ihandle* *ih_list);
	return mkih(C.IupNormalizerv((**C.Ihandle)(unsafe.Pointer(&(ihList[0])))))
}

//Cbox creates a void container for position elements in absolute coordinates.
//It is a concrete layout container.
//
//It does not have a native representation.
//
//The IupCbox is equivalent of a IupVbox or IupHbox where all the children have the FLOATING attribute set to YES,
//but children must use CX and CY attributes instead of the POSITION attribute.
func Cbox(children ...Ihandle) Ihandle {
	children = append(children, Ihandle(0))

	//Ihandle* IupCbox (Ihandle* child, ...);
	//Ihandle* IupCboxv (Ihandle* *children);
	return mkih(C.IupCboxv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//Sbox creates a void container that allows its child to be resized.
//Allows expanding and contracting the child size in one direction.
//
//It does not have a native representation but it contains also a IupCanvas to implement the bar handler.
func Sbox(child Ihandle) Ihandle {
	//Ihandle* IupSbox (Ihandle *child);
	return mkih(C.IupSbox(child.ptr()))
}

//Split creates a void container that split its client area in two.
//Allows the provided controls to be enclosed in a box that allows expanding
//and contracting the element size in one direction, but when one is expanded the other is contracted.
//
//It does not have a native representation, but it contains also a IupCanvas to implement the bar handler.
func Split(child1, child2 Ihandle) Ihandle {
	//Ihandle* IupSplit (Ihandle* child1, Ihandle* child2);
	return mkih(C.IupSplit(child1.ptr(), child2.ptr()))
}

//ScrollBox creates a native container that allows its child to be scrolled. It inherits from IupCanvas.
func ScrollBox(child Ihandle) Ihandle {
	//Ihandle* IupScrollBox (Ihandle* child);
	return mkih(C.IupScrollBox(child.ptr()))
}

//GridBox creates a void container for composing elements in a regular grid.
//It is a box that arranges the elements it contains from top to bottom and from left to right,
//but can distribute the elements in lines or in columns.
//
//The child elements are added to the control just like a vbox and hbox, sequentially.
//Then they are distributed accordingly the attributes ORIENTATION and NUMDIV.
//When ORIENTATION=HORIZONTAL children are distributed from left to right on the first line until NUMDIV, then on the second line, and so on.
//When ORIENTATION=VERTICAL children are distributed from top to bottom on the first column until NUMDIV, then on the second column, and so on.
//The number of lines and the number of columns can be easily obtained from the combination of these attributes:
//
//  if (orientation==IGBOX_HORIZONTAL)
//    num_lin = child_count / num_div
//    num_col = num_div
//  else
//    num_lin = num_div
//    num_col = child_count / num_div
//
//Notice that the total number of spaces can be larger than the number of actual children,
//the last line or column can be incomplete.
//
//The column sizes will be based only on the width of the children of the reference line, usually line 0.
//The line sizes will be based only on the height of the children of the reference column, usually column 0.
//
//It does not have a native representation.
func GridBox(children ...Ihandle) Ihandle {
	children = append(children, Ihandle(0))

	//Ihandle* IupGridBox (Ihandle* child, ...);
	//Ihandle* IupGridBoxv (Ihandle **children);
	return mkih(C.IupGridBoxv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//Expander creates a void container that can interactively show or hide its child.
//
//It does not have a native representation, but it contains also several elements to implement the bar handler.
func Expander(child Ihandle) Ihandle {
	//Ihandle* IupExpander (Ihandle *child);
	return mkih(C.IupExpander(child.ptr()))
}

//DetachBox creates a detachable void container.
//
//Dragging and dropping this element, it creates a new dialog composed by its child
//or elements arranged in it (for example, a child like IupVbox or IupHbox).
//During the drag, the ESC key can be pressed to cancel the action.
//
//It does not have a native representation, but it contains also a IupCanvas to implement the bar handler.
func DetachBox(child Ihandle) Ihandle {
	//Ihandle* IupDetachBox (Ihandle *child);
	return mkih(C.IupDetachBox(child.ptr()))
}

//BackgroundBox creates a simple native container with no decorations.
//Useful for controlling children visibility for IupZbox or IupExpander. It inherits from IupCanvas.
func BackgroundBox(child Ihandle) Ihandle {
	//Ihandle* IupBackgroundBox(Ihandle *child);
	return mkih(C.IupBackgroundBox(child.ptr()))
}

//Frame creates a native container, which draws a frame with a title around its child.
func Frame(child Ihandle) Ihandle {
	//Ihandle* IupFrame (Ihandle* child);
	return mkih(C.IupFrame(child.ptr()))
}

//Image creates an image to be shown on a label, button, toggle, or as a cursor.
func Image(width, height int, pixMap []byte) Ihandle {
	//Ihandle* IupImage (int width, int height, const unsigned char *pixmap);
	return mkih(C.IupImage(C.int(width), C.int(height), (*C.uchar)(unsafe.Pointer(&pixMap[0]))))
}

//ImageRGB creates an image to be shown on a label, button, toggle, or as a cursor.
func ImageRGB(width, height int, pixMap []byte) Ihandle {
	//Ihandle* IupImageRGB (int width, int height, const unsigned char *pixmap);
	return mkih(C.IupImageRGB(C.int(width), C.int(height), (*C.uchar)(unsafe.Pointer(&pixMap[0]))))
}

//ImageRGBA creates an image to be shown on a label, button, toggle, or as a cursor.
func ImageRGBA(width, height int, pixMap []byte) Ihandle {
	//Ihandle* IupImageRGBA (int width, int height, const unsigned char *pixmap);
	return mkih(C.IupImageRGBA(C.int(width), C.int(height), (*C.uchar)(unsafe.Pointer(&pixMap[0]))))
}

//Item creates an item of the menu interface element. When selected, it generates an action.
func Item(title string, action ...string) Ihandle {
	c_title, c_action := C.CString(title), optionalAction(action)
	defer C.free(unsafe.Pointer(c_title))
	defer cStrFree(c_action)

	//Ihandle* IupItem (const char* title, const char* action);
	return mkih(C.IupItem(c_title, c_action))
}

//Submenu creates a menu item that, when selected, opens another menu.
func Submenu(title string, child Ihandle) Ihandle {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupSubmenu (const char* title, Ihandle* child);
	return mkih(C.IupSubmenu(c_title, child.ptr()))
}

//Separator creates the separator interface element. It shows a line between two menu items.
func Separator() Ihandle {
	//Ihandle* IupSeparator (void);
	return mkih(C.IupSeparator())
}

//Menu Creates a menu element, which groups 3 types of interface elements: item, submenu and separator.
//Any other interface element defined inside a menu will be an error.
func Menu(children ...Ihandle) Ihandle {
	children = append(children, Ihandle(0))

	//Ihandle* IupMenu (Ihandle* child,...);
	//Ihandle* IupMenuv (Ihandle* *children);
	return mkih(C.IupMenuv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//Button creates an interface element that is a button.
//When selected, this element activates a function in the application.
//Its visual presentation can contain a text and/or an image.
func Button(title string, action ...string) Ihandle {
	c_title, c_action := C.CString(title), optionalAction(action)
	defer C.free(unsafe.Pointer(c_title))
	defer cStrFree(c_action)

	//Ihandle* IupButton (const char* title, const char* action);
	return mkih(C.IupButton(c_title, c_action))
}

//Canvas creates an interface element that is a canvas - a working area for your application.
func Canvas(action ...string) Ihandle {
	c_action := optionalAction(action)
	defer cStrFree(c_action)

	//Ihandle* IupCanvas (const char* action);
	return mkih(C.IupCanvas(c_action))
}

//Dialog creates a dialog element. It manages user interaction with the interface elements.
//For any interface element to be shown, it must be encapsulated in a dialog.
func Dialog(child Ihandle) Ihandle {
	//Ihandle* IupDialog (Ihandle* child);
	return mkih(C.IupDialog(child.ptr()))
}

//User creates a user element in IUP, which is not associated to any interface element.
//It is used to map an external element to a IUP element.
//Its use is usually for additional elements, but you can use it to create an Ihandle* to store private attributes.
//
//It is also a void container. Its children can be dynamically added using IupAppend or IupInsert.
func User() Ihandle {
	//Ihandle* IupUser (void);
	return mkih(C.IupUser())
}

//Label creates a label interface element, which displays a separator, a text or an image.
func Label(title string) Ihandle {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupLabel (const char* title);
	return mkih(C.IupLabel(c_title))
}

//List Creates an interface element that displays a list of items.
//The list can be visible or can be dropped down. It also can have an edit box for text input. So it is a 4 in 1 element.
//In native systems the dropped down case is called Combo Box.
func List(action ...string) Ihandle {
	c_action := optionalAction(action)
	defer cStrFree(c_action)

	//Ihandle* IupList (const char* action);
	return mkih(C.IupList(c_action))
}

//Text creates an editable text field.
func Text(action ...string) Ihandle {
	c_action := optionalAction(action)
	defer cStrFree(c_action)

	//Ihandle* IupText (const char* action);
	return mkih(C.IupText(c_action))
}

//MultiLine creates an editable field with one or more lines.
//
//Since IUP 3.0, IupText has support for multiple lines when the MULTILINE attribute is set to YES.
//Now when a IupMultiline element is created in fact a IupText element with MULTILINE=YES is created.
func MultiLine(action ...string) Ihandle {
	c_action := optionalAction(action)
	defer cStrFree(c_action)

	//Ihandle* IupMultiLine (const char* action);
	return mkih(C.IupMultiLine(c_action))
}

//Toggle creates the toggle interface element.
//It is a two-state (on/off) button that, when selected, generates an action that activates a function in the associated application.
//Its visual representation can contain a text or an image.
func Toggle(title string, action ...string) Ihandle {
	c_title, c_action := C.CString(title), optionalAction(action)
	defer C.free(unsafe.Pointer(c_title))
	defer cStrFree(c_action)

	//Ihandle* IupToggle (const char* title, const char* action);
	return mkih(C.IupToggle(c_title, c_action))
}

//Timer creates a timer which periodically invokes a callback when the time is up.
//Each timer should be destroyed using IupDestroy.
func Timer() Ihandle {
	//Ihandle* IupTimer (void);
	return mkih(C.IupTimer())
}

//Clipboard creates an element that allows access to the clipboard.
//Each clipboard should be destroyed using IupDestroy,
//but you can use only one for the entire application because it does not store any data inside.
//Or you can simply create and destroy every time you need to copy or paste.
func Clipboard() Ihandle {
	//Ihandle* IupClipboard (void);
	return mkih(C.IupClipboard())
}

//ProgressBar creates a progress bar control. Shows a percent value that can be updated to simulate a progression.
//
//It is similar of IupGauge, but uses native controls internally. Also does not have support for text inside the bar.
func ProgressBar() Ihandle {
	//Ihandle* IupProgressBar(void);
	return mkih(C.IupProgressBar())
}

//Val creates a Valuator control. Selects a value in a limited interval.
//Also known as Scale or Trackbar in native systems.
func Val(_type ...string) Ihandle {
	c_type := optionalAction(_type)
	defer cStrFree(c_type)

	//Ihandle* IupVal (const char *type);
	return mkih(C.IupVal(c_type))
}

//Tabs creates a native container for composing elements in hidden layers with only one layer visible (just like IupZbox),
//but its visibility can be interactively controlled.
//The interaction is done in a line of tabs with titles and arranged according to the tab type.
//Also known as Notebook in native systems.
func Tabs(children ...Ihandle) Ihandle {
	children = append(children, Ihandle(0))

	//Ihandle* IupTabs (Ihandle* child, ...);
	//Ihandle* IupTabsv (Ihandle* *children);
	return mkih(C.IupTabsv((**C.Ihandle)(unsafe.Pointer(&(children[0])))))
}

//Tree creates a tree containing nodes of branches or leaves. Both branches and leaves can have an associated text and image.
//
//The branches can be expanded or collapsed. When a branch is expanded,
//its immediate children are visible, and when it is collapsed they are hidden.
//
//The leaves can generate an "executed" or "renamed" actions, branches can only generate a "renamed" action.
//
//The focus node is the node with the focus rectangle, marked nodes have their background inverted.
func Tree() Ihandle {
	//Ihandle* IupTree (void);
	return mkih(C.IupTree())
}

//Link creates a label that displays an underlined clickable text. It inherits from IupLabel.
func Link(url, title string) Ihandle {
	c_url, c_title := C.CString(url), C.CString(title)
	defer C.free(unsafe.Pointer(c_url))
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupLink (const char* url, const char* title);
	return mkih(C.IupLink(c_url, c_title))
}

//FlatButton creates an interface element that is a button, but it does not have native decorations.
//When selected, this element activates a function in the application.
//Its visual presentation can contain a text and/or an image.
//
//It behaves just like an IupButton, but since it is not a native control it has more flexibility for additional options.
//It can also behave like an IupToggle (without the checkmark).
//
//It inherits from IupCanvas.
func FlatButton(title string) Ihandle {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	//Ihandle* IupFlatButton (const char* title);
	return mkih(C.IupFlatButton(c_title))
}

//AnimatedLabel creates an animated label interface element, which displays an image that is changed periodically.
//
//It uses an animation that is simply an IupUser with several IupImage as children.
//
//It inherits from IupLabel.
func AnimatedLabel(animation Ihandle) Ihandle {
	//Ihandle* IupAnimatedLabel(Ihandle* animation);
	return mkih(C.IupAnimatedLabel(animation.ptr()))
}

//DatePick creates a date editing interface element, which can displays a calendar for selecting a date.
//
//In Windows is a native element. In GTK and Motif is a custom element. In Motif is not capable of displaying the calendar.
func DatePick() Ihandle {
	//Ihandle* IupDatePick (void);
	return mkih(C.IupDatePick())
}

//Calendar creates a month calendar interface element, where the user can select a date.
//
//GTK and Windows only. NOT available in Motif.
func Calendar() Ihandle {
	//Ihandle* IupCalendar (void);
	return mkih(C.IupCalendar())
}

/* Old controls, use SPIN attribute of IupText */
//Ihandle* IupSpin (void);
//Ihandle* IupSpinbox (Ihandle* child);

/* ---------------------------------------------------------------------------------------------- */
/*                                          Utilities                                             */
/* ---------------------------------------------------------------------------------------------- */

/* String compare utility */

//StringCompare is an utility function to compare strings lexicographically.
//Used internally in IupMatrixEx when sorting, but available in the main library.
//
//This means that numbers and text in the string are sorted separately (for ex: A1 A2 A11 A30 B1).
//Also natural alphabetic order is used: 123...aAáÁ...bBcC...
//The comparison will work only for Latin-1 characters, even if UTF8MODE is Yes.
func StringCompare(str1, str2 string, caseSensitive, lexicographic bool) int {
	c_str1, c_str2 := C.CString(str1), C.CString(str2)
	defer C.free(unsafe.Pointer(c_str1))
	defer C.free(unsafe.Pointer(c_str2))

	//int IupStringCompare(const char* str1, const char* str2, int casesensitive, int lexicographic);
	return int(C.IupStringCompare(c_str1, c_str2, C.int(bool2int(caseSensitive)), C.int(bool2int(lexicographic))))
}

/* IupImage utility */

//SaveImageAsText saves the IupImage as a text file to be reused in other programs.
//
//It does NOT depends on the IM library.
func SaveImageAsText(ih Ihandle, fileName, format, name string) bool {
	c_fileName, c_format, c_name := C.CString(fileName), C.CString(format), C.CString(name)
	defer C.free(unsafe.Pointer(c_fileName))
	defer C.free(unsafe.Pointer(c_format))
	defer C.free(unsafe.Pointer(c_name))

	//int IupSaveImageAsText(Ihandle* ih, const char* file_name, const char* format, const char* name);
	return int(C.IupSaveImageAsText(ih.ptr(), c_fileName, c_format, c_name)) != 0
}

/* IupText and IupScintilla utilities */

//TextConvertLinColToPos converts a (lin, col) character positioning into an absolute position.
//lin and col starts at 1, pos starts at 0. For single line controls pos is always "col-1".
func TextConvertLinColToPos(ih Ihandle, lin, col int) (pos int) {
	//void IupTextConvertLinColToPos(Ihandle* ih, int lin, int col, int *pos);
	C.IupTextConvertLinColToPos(ih.ptr(), C.int(lin), C.int(col), (*C.int)(unsafe.Pointer(&pos)))
	return
}

//TextConvertPosToLinCol Converts an absolute position into a (lin, col) character positioning.
//lin and col starts at 1, pos starts at 0. For single line controls lin is always 1, and col is always "pos+1".
func TextConvertPosToLinCol(ih Ihandle, pos int) (lin, col int) {
	//void IupTextConvertPosToLinCol(Ihandle* ih, int pos, int *lin, int *col);
	C.IupTextConvertPosToLinCol(ih.ptr(), C.int(pos), (*C.int)(unsafe.Pointer(&lin)), (*C.int)(unsafe.Pointer(&col)))
	return
}

/* IupText, IupList, IupTree, IupMatrix and IupScintilla utility */

//ConvertXYToPos converts a (x,y) coordinate in an item position.
//
//It can be used for IupText and IupScintilla (returns a position in the string), IupList (returns an item), IupTree (returns a node identifier) or IupMatrix (returns a cell position, where pos=lin*numcol + col).
func ConvertXYToPos(ih Ihandle, x, y int) int {
	//int IupConvertXYToPos(Ihandle* ih, int x, int y);
	return int(C.IupConvertXYToPos(ih.ptr(), C.int(x), C.int(y)))
}

//NOTE no need to implement. All strings are already copied (and not referenced)
/* OLD names, kept for backward compatibility, will never be removed. */
//void IupStoreGlobal(const char* name, const char* value);
//void IupStoreAttribute(Ihandle* ih, const char* name, const char* value);
//void IupSetfAttribute(Ihandle* ih, const char* name, const char* format, ...);
//void IupStoreAttributeId(Ihandle *ih, const char* name, int id, const char *value);
//void IupSetfAttributeId(Ihandle *ih, const char* name, int id, const char* f, ...);
//void IupStoreAttributeId2(Ihandle* ih, const char* name, int lin, int col, const char* value);
//void IupSetfAttributeId2(Ihandle* ih, const char* name, int lin, int col, const char* format, ...);

/* IupTree utilities */

//TreeSetUserId associates an userid with a given id. If the id of the node is changed, the userid remains the same.
func TreeSetUserId(ih Ihandle, id int, userId uintptr) bool {
	//int IupTreeSetUserId(Ihandle* ih, int id, void* userid);
	return int(C.IupTreeSetUserId(ih.ptr(), C.int(id), unsafe.Pointer(userId))) != 0
}

//TreeGetUserId Returns the pointer or Lua table associated to the node or NULL if none was associated.
//SetUserId must have been called for the node with the given id.
func TreeGetUserId(ih Ihandle, id int) uintptr {
	//void* IupTreeGetUserId(Ihandle* ih, int id);
	return uintptr(C.IupTreeGetUserId(ih.ptr(), C.int(id)))
}

//TreeGetId returns the id of the node that has the userid on success or -1 (nil) if not found.
//SetUserId must have been called with the same userid.
func TreeGetId(ih Ihandle, userId uintptr) int {
	//int IupTreeGetId(Ihandle* ih, void *userid);
	return int(C.IupTreeGetId(ih.ptr(), unsafe.Pointer(userId)))
}

//TreeSetAttributeHandle
func TreeSetAttributeHandle(ih Ihandle, name string, id int, ihNamed Ihandle) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	//void IupTreeSetAttributeHandle(Ihandle* ih, const char* name, int id, Ihandle* ih_named);
	C.IupTreeSetAttributeHandle(ih.ptr(), c_name, C.int(id), ihNamed.ptr())
}

/* DEPRECATED IupTree utilities, use Iup*AttributeId functions. It will be removed in a future version. */
//void IupTreeSetAttribute (Ihandle* ih, const char* name, int id, const char* value);
//void IupTreeStoreAttribute(Ihandle* ih, const char* name, int id, const char* value);
//char* IupTreeGetAttribute (Ihandle* ih, const char* name, int id);
//int IupTreeGetInt (Ihandle* ih, const char* name, int id);
//float IupTreeGetFloat (Ihandle* ih, const char* name, int id);
//void IupTreeSetfAttribute (Ihandle* ih, const char* name, int id, const char* format, ...);

/* DEPRECATED callback management. It will be removed in a future version. */
//const char* IupGetActionName(void);

/* DEPRECATED font names. It will be removed in a future version. */
//char* IupMapFont (const char *iupfont);
//char* IupUnMapFont (const char *driverfont);

/* ---------------------------------------------------------------------------------------------- */
/*                                     Pre-defined dialogs                                        */
/* ---------------------------------------------------------------------------------------------- */

//FileDlg creates the File Dialog element. It is a predefined dialog for selecting files or a directory.
//The dialog can be shown with the IupPopup function only.
func FileDlg() Ihandle {
	//Ihandle* IupFileDlg(void);
	return mkih(C.IupFileDlg())
}

//MessageDlg creates the Message Dialog element. It is a predefined dialog for displaying a message.
//The dialog can be shown with the IupPopup function only.
func MessageDlg() Ihandle {
	//Ihandle* IupMessageDlg(void);
	return mkih(C.IupMessageDlg())
}

//ColorDlg creates the Color Dialog element. It is a predefined dialog for selecting a color.
//
//There are 3 versions of the dialog. One for Windows only, one for GTK only and one for all systems,
//but it is based on the IupColorBrowser control that depends on the CD library.
//
//The Windows and GTK dialogs can be shown only with the IupPopup function.
//The IupColorBrowser based dialog is a IupDialog that can be shown as any regular IupDialog.
//
//IMPORTANT: The IupColorBrowser based dialog is included in the Controls Library.
//When the Controls Library is initialized the Windows and GTK dialogs are not available anymore,
//i.e. before the Controls Library initialization only the Windows and GTK dialogs are available,
//after only the IupColorBrowser based dialog is available.
func ColorDlg() Ihandle {
	//Ihandle* IupColorDlg(void);
	return mkih(C.IupColorDlg())
}

//FontDlg creates the Font Dialog element. It is a predefined dialog for selecting a font.
//The dialog can be shown with the IupPopup function only.
func FontDlg() Ihandle {
	//Ihandle* IupFontDlg(void);
	return mkih(C.IupFontDlg())
}

//ProgressDlg creates a progress dialog element. It is a predefined dialog for displaying the progress of an operation.
//The dialog is meant to be shown with the show functions IupShow or IupShowXY.
func ProgressDlg() Ihandle {
	//Ihandle* IupProgressDlg(void);
	return mkih(C.IupProgressDlg())
}

//GetFile shows a modal dialog of the native interface system to select a filename. Uses the IupFileDlg element.
func GetFile(path string) (sel string, ret int) {
	if len(path) > 4095 {
		panic(fmt.Errorf("iup.GetFile: path is too long (maximum is 4095)"))
	}
	buf := bytes.NewBuffer([]byte(path))
	buf.Grow(4096 - len(path))
	byt := buf.Bytes()

	//int IupGetFile(char *arq);
	ret = int(C.IupGetFile((*C.char)((unsafe.Pointer)(&byt[0]))))

	sel = string(byt[:int(C.strlen((*C.char)((unsafe.Pointer)(&byt[0]))))])
	return
}

//Message shows a modal dialog containing a message. It simply creates and popup a IupMessageDlg.
func Message(title, msg string) {
	c_title, c_msg := C.CString(title), C.CString(msg)
	defer C.free(unsafe.Pointer(c_title))
	defer C.free(unsafe.Pointer(c_msg))

	//void IupMessage(const char *title, const char *msg);
	//void IupMessagef(const char *title, const char *format, ...);
	C.IupMessage(c_title, c_msg)
}

//Alarm shows a modal dialog containing a message and up to three buttons.
func Alarm(title, msg, b1, b2, b3 string) int {
	c_title, c_msg, c_b1, c_b2, c_b3 := C.CString(title), C.CString(msg), C.CString(b1), cStrOrNull(b2), cStrOrNull(b3)
	defer C.free(unsafe.Pointer(c_title))
	defer C.free(unsafe.Pointer(c_msg))
	defer C.free(unsafe.Pointer(c_b1))
	defer cStrFree(c_b2)
	defer cStrFree(c_b3)

	//int IupAlarm(const char *title, const char *msg, const char *b1, const char *b2, const char *b3);
	return int(C.IupAlarm(c_title, c_msg, c_b1, c_b2, c_b3))
}

//int IupScanf(const char *format, ...);

//ListDialog shows a modal dialog to select items from a simple or multiple selection list.
func ListDialog(_type int, title string, list []string, op, maxCol, maxLin int) (ret int, marks []bool) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	pList := make([]*C.char, len(list))
	for i := 0; i < len(list); i++ {
		pList[i] = C.CString(list[i])
	}
	defer func() {
		for i := 0; i < len(list); i++ {
			C.free(unsafe.Pointer(pList[i]))
		}
	}()

	pMark := make([]C.int, len(list))
	defer func() {
		marks = make([]bool, len(list))
		for i := 0; i < len(list); i++ {
			marks[i] = pMark[i] != C.int(0)
		}
	}()

	//int IupListDialog(int type, const char *title, int size, const char** list, int op, int max_col, int max_lin, int* marks);
	ret = int(C.IupListDialog(C.int(_type), c_title, C.int(len(list)), (**C.char)(unsafe.Pointer(&(pList[0]))), C.int(op), C.int(maxCol), C.int(maxLin), (*C.int)(unsafe.Pointer(&pMark[0]))))
	return
}

//GetText shows a modal dialog to edit a multiline text.
//
//Go specific notes: There is no more text length limit.
func GetText(title, text string) string {
	multiText := MultiLine().SetAttributes(map[string]string{
		"EXPAND":         "YES",
		"VALUE":          text,
		"FONT":           "Courier, 12",
		"VISIBLELINES":   "10",
		"VISIBLECOLUMNS": "50",
	})

	ok := Button("_@IUP_OK").SetAttribute("PADDING", GetGlobal("DEFAULTBUTTONPADDING"))
	ok.SetCallback("ACTION", func(ih Ihandle) int {
		GetDialog(ih).SetAttribute("STATUS", 1)
		return CLOSE
	})

	cancel := Button("_@IUP_CANCEL").SetAttribute("PADDING", GetGlobal("DEFAULTBUTTONPADDING"))
	cancel.SetCallback("ACTION", func(ih Ihandle) int {
		GetDialog(ih).SetAttribute("STATUS", -1)
		return CLOSE
	})

	button_box := Hbox(Fill(), ok, cancel).SetAttributes(`MARGIN=0x0,NORMALIZESIZE=HORIZONTAL`)
	dlg_box := Vbox(multiText, button_box).SetAttributes(`MARGIN=10x10,GAP=10`)

	dlg := Dialog(dlg_box).SetAttributes(map[string]interface{}{
		"TITLE":        title,
		"MINBOX":       "NO",
		"MAXBOX":       "NO",
		"DEFAULTENTER": ok,
		"DEFAULTESC":   cancel,
		"PARENTDIALOG": GetGlobal("PARENTDIALOG"),
		"ICON":         GetGlobal("ICON"),
	})
	defer dlg.Destroy()

	Map(dlg)

	multiText.SetAttribute("VISIBLELINES", nil).SetAttribute("VISIBLECOLUMNS", nil)

	Popup(dlg, CENTERPARENT, CENTERPARENT)

	if GetInt(dlg, "STATUS") == 1 {
		return GetAttribute(multiText, "VALUE")
	} else {
		return ""
	}
}

//GetColor shows a modal dialog which allows the user to select a color. Based on IupColorDlg.
func GetColor(x, y int) (ret int, r, g, b uint8) {
	//int IupGetColor(int x, int y, unsigned char* r, unsigned char* g, unsigned char* b);
	ret = int(C.IupGetColor(C.int(x), C.int(y), (*C.uchar)(unsafe.Pointer(&r)), (*C.uchar)(unsafe.Pointer(&g)), (*C.uchar)(unsafe.Pointer(&b))))
	return
}

//typedef int (*Iparamcb)(Ihandle* dialog, int param_index, void* user_data);
type Iparamcb func(dialog Ihandle, paramIndex int, userData uintptr) int

//GetParam shows a modal dialog for capturing parameter values using several types of controls.
func GetParam(title string, action Iparamcb, userData uintptr, format string, args ...interface{}) bool {
	//int IupGetParam(const char* title, Iparamcb action, void* user_data, const char* format,...);
	//int IupGetParamv(const char* title, Iparamcb action, void* user_data, const char* format, int param_count, int param_extra, void** param_data);

	splitFmt := strings.Split(strings.TrimSuffix(strings.TrimPrefix(format, "\n"), "\n"), "\n")
	handles := make([]Ihandle, len(splitFmt))

	p := 0
	for i := 0; i < len(splitFmt); i++ {
		handles[i] = Param(splitFmt[i])
		if handles[i] == 0 {
			return false
		}

		switch strings.ToUpper(GetAttribute(handles[i], "DATATYPE")) {
		case "FLOAT":
			SetAttribute(handles[i], "VALUE", *(args[p].(*float32)))
			p++
		case "DOUBLE":
			SetAttribute(handles[i], "VALUE", *(args[p].(*float64)))
			p++
		case "INT":
			SetAttribute(handles[i], "VALUE", *(args[p].(*int)))
			p++
		case "STRING":
			SetAttribute(handles[i], "VALUE", *(args[p].(*string)))
			p++
		case "HANDLE":
			SetAttribute(handles[i], "VALUE", args[p].(Ihandle))
			p++
		}
	}

	dlg := ParamBox(0, handles)
	defer Destroy(dlg)

	dlg.SetAttribute("PARENTDIALOG", GetGlobalIh("PARENTDIALOG"))
	dlg.SetAttribute("ICON", GetGlobalIh("ICON"))
	dlg.SetAttribute("TITLE", title)
	dlg.SetAttribute("USERDATA", userData)
	dlg.SetCallback("PARAM_CB", action)
	dlg.SetCallback("CLOSE_CB", func(ih Ihandle) int {
		if action != nil && action(dlg, GETPARAM_CLOSE, userData) == 0 {
			return IGNORE
		} else {
			return CLOSE
		}
	})

	if action != nil {
		Map(dlg)
		action(dlg, GETPARAM_INIT, userData)
	}

	Popup(dlg, CENTERPARENT, CENTERPARENT)

	if dlg.GetInt("STATUS") == 0 {
		return false
	}

	// else
	p = 0
	for i := 0; i < len(splitFmt); i++ {
		switch strings.ToUpper(GetAttribute(handles[i], "DATATYPE")) {
		case "FLOAT":
			*args[p].(*float32) = handles[i].GetFloat("VALUE")
			p++
		case "DOUBLE":
			*args[p].(*float64) = handles[i].GetDouble("VALUE")
			p++
		case "INT":
			*args[p].(*int) = handles[i].GetInt("VALUE")
			p++
		case "STRING":
			*args[p].(*string) = handles[i].GetAttribute("VALUE")
			p++
		}
	}
	return true
}

//Param creates an IupUser element to be used in the IupParamBox.
//Each parameter format follows the same specifications as the IupGetParam function, including the line feed.
func Param(format string) Ihandle {
	c_format := C.CString(format)
	defer C.free(unsafe.Pointer(c_format))

	//Ihandle* IupParamf(const char* format);
	return mkih(C.IupParamf(c_format))
}

//ParamBox creates the IupGetParam dialog contents with the array of parameters. This includes the button box at the bottom.
func ParamBox(parent Ihandle, params []Ihandle) Ihandle {
	//Ihandle* IupParamBox(Ihandle* parent, Ihandle** params, int count);
	return mkih(C.IupParamBox(parent.ptr(), (**C.struct_Ihandle_)(unsafe.Pointer(&params[0])), C.int(len(params))))
}

//LayoutDialog creates a Layout Dialog. It is a predefined dialog to visually edit the layout of another dialog in run time.
//It is a standard IupDialog constructed with other IUP elements.
//The dialog can be shown with any of the show functions IupShow, IupShowXY or IupPopup.
//
//Any existent dialog can be selected. It does not need to be mapped on the native system nor visible.
//It could have been created in C, LED or Lua.
//
//The layout dialog is composed by two areas: one showing the given dialog children hierarchy tree, and one displaying its layout.
//
//This is a dialog intended for developers, so they can see and inspect their dialogs in other ways.
func LayoutDialog(dialog Ihandle) Ihandle {
	//Ihandle* IupLayoutDialog(Ihandle* dialog);
	return mkih(C.IupLayoutDialog(dialog.ptr()))
}

//ElementPropertiesDialog creates an Element Properties Dialog.
//It is a predefined dialog to edit the properties of an element in run time.
//It is a standard IupDialog constructed with other IUP elements.
//The dialog can be shown with any of the show functions IupShow, IupShowXY or IupPopup.
//
//Any existent element can be edited. It does not need to be mapped on the native system nor visible.
//It could have been created in C, LED or Lua.
//
//This is a dialog intended for developers, so they can see and inspect their elements in other ways.
//
//It contains 3 Tab sections: one for the registered attributes of the element, one for custom attributes set at the hash table, and one for the callbacks.
//The callbacks are just for inspection, and custom attribute should be handled carefully because they may be not strings.
//Registered attributes values are shown in red when they were changed by the application.
func ElementPropertiesDialog(elem Ihandle) Ihandle {
	//Ihandle* IupElementPropertiesDialog(Ihandle* elem);
	return mkih(C.IupElementPropertiesDialog(elem.ptr()))
}

/* ---------------------------------------------------------------------------------------------- */

/* ------------------------------ */
/* Common flags and return values */
/* ------------------------------ */

const (
	ERROR      = 1   // common flag and return value
	NOERROR    = 0   // common flag and return value
	OPENED     = -1  // common flag and return value
	INVALID    = -1  // common flag and return value
	INVALID_ID = -10 // common flag and return value
)

/* ---------------------- */
/* Callback return values */
/* ---------------------- */

const (
	IGNORE   = -1 // callback return value
	DEFAULT  = -2 // callback return value
	CLOSE    = -3 // callback return value
	CONTINUE = -4 // callback return value
)

/* --------------------------------------- */
/* IupPopup and IupShowXY parameter values */
/* --------------------------------------- */

const (
	CENTER       = 0xFFFF // iup.Popup and iup.ShowXY parameter value
	LEFT         = 0xFFFE // iup.Popup and iup.ShowXY parameter value
	RIGHT        = 0xFFFD // iup.Popup and iup.ShowXY parameter value
	MOUSEPOS     = 0xFFFC // iup.Popup and iup.ShowXY parameter value
	CURRENT      = 0xFFFB // iup.Popup and iup.ShowXY parameter value
	CENTERPARENT = 0xFFFA // iup.Popup and iup.ShowXY parameter value
	TOP          = LEFT   // iup.Popup and iup.ShowXY parameter value
	BOTTOM       = RIGHT  // iup.Popup and iup.ShowXY parameter value
)

/* ----------------------- */
/* SHOW_CB callback values */
/* ----------------------- */

const (
	SHOW     = iota // SHOW_CB callback value
	RESTORE         // SHOW_CB callback value
	MINIMIZE        // SHOW_CB callback value
	MAXIMIZE        // SHOW_CB callback value
	HIDE            // SHOW_CB callback value
)

/* ------------------------- */
/* SCROLL_CB callback values */
/* ------------------------- */

const (
	SBUP      = iota // SCROLL_CB callback value
	SBDN             // SCROLL_CB callback value
	SBPGUP           // SCROLL_CB callback value
	SBPGDN           // SCROLL_CB callback value
	SBPOSV           // SCROLL_CB callback value
	SBDRAGV          // SCROLL_CB callback value
	SBLEFT           // SCROLL_CB callback value
	SBRIGHT          // SCROLL_CB callback value
	SBPGLEFT         // SCROLL_CB callback value
	SBPGRIGHT        // SCROLL_CB callback value
	SBPOSH           // SCROLL_CB callback value
	SBDRAGH          // SCROLL_CB callback value
)

/* ------------------------------ */
/* Mouse button values and macros */
/* ------------------------------ */

const (
	BUTTON1 = '1' // mouse button value
	BUTTON2 = '2' // mouse button value
	BUTTON3 = '3' // mouse button value
	BUTTON4 = '4' // mouse button value
	BUTTON5 = '5' // mouse button value
)

// mouse button macro
func IsShift(s uintptr) bool {
	return CStrToString(s)[0] == 'S'
}

// mouse button macro
func IsControl(s uintptr) bool {
	return CStrToString(s)[1] == 'C'
}

// mouse button macro
func IsButton1(s uintptr) bool {
	return CStrToString(s)[2] == '1'
}

// mouse button macro
func IsButton2(s uintptr) bool {
	return CStrToString(s)[3] == '2'
}

// mouse button macro
func IsButton3(s uintptr) bool {
	return CStrToString(s)[4] == '3'
}

// mouse button macro
func IsDouble(s uintptr) bool {
	return CStrToString(s)[5] == 'D'
}

// mouse button macro
func IsAlt(s uintptr) bool {
	return CStrToString(s)[6] == 'A'
}

// mouse button macro
func IsSys(s uintptr) bool {
	return CStrToString(s)[7] == 'Y'
}

// mouse button macro
func IsButton4(s uintptr) bool {
	return CStrToString(s)[8] == '4'
}

// mouse button macro
func IsButton5(s uintptr) bool {
	return CStrToString(s)[9] == '5'
}

/* ----------------- */
/* Pre-defined masks */
/* ----------------- */

const (
	MASK_FLOAT       = "[+/-]?(/d+/.?/d*|/./d+)"                 // pre-defined masks
	MASK_UFLOAT      = "(/d+/.?/d*|/./d+)"                       // pre-defined masks
	MASK_EFLOAT      = "[+/-]?(/d+/.?/d*|/./d+)([eE][+/-]?/d+)?" // pre-defined masks
	MASK_FLOATCOMMA  = "[+/-]?(/d+/,?/d*|/,/d+)"                 // pre-defined masks
	MASK_UFLOATCOMMA = "(/d+/,?/d*|/,/d+)"                       // pre-defined masks
	MASK_INT         = "[+/-]?/d+"                               // pre-defined masks
	MASK_UINT        = "/d+"                                     // pre-defined masks
)

/* ------------------------------- */
/* IupGetParam callback situations */
/* ------------------------------- */

const (
	GETPARAM_BUTTON1 = -1               // iup.GetParam callback situation
	GETPARAM_INIT    = -2               // iup.GetParam callback situation
	GETPARAM_BUTTON2 = -3               // iup.GetParam callback situation
	GETPARAM_BUTTON3 = -4               // iup.GetParam callback situation
	GETPARAM_CLOSE   = -5               // iup.GetParam callback situation
	GETPARAM_OK      = GETPARAM_BUTTON1 // iup.GetParam callback situation
	GETPARAM_CANCEL  = GETPARAM_BUTTON2 // iup.GetParam callback situation
	GETPARAM_HELP    = GETPARAM_BUTTON3 // iup.GetParam callback situation
)

/* ------------------ */
/* Record input modes */
/* ------------------ */

const (
	RECBINARY = iota // record input mode
	RECTEXT          // record input mode
)
