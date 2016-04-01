// Shows a dialog with all the possible fields.
package main

import (
	"fmt"

	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	var (
		_bool  int     = 1
		_int   int     = 3456
		_real  float32 = 3.453
		_int2  int     = 192
		_real2 float32 = 0.5
		_angle float32 = 90
		_str   string  = "string text"
		_color string  = "255 0 128"
		_list  int     = 2
		_str2  string  = "second text\nsecond line"
		_file  string  = "test.jpg"
	)

	iup.GetParam("Title", action_cb, 0,
		`Boolean: %b[No,Yes]{Boolean Tip}\n
Integer: %i{Integer Tip}\n
Real 1: %r{Real Tip}\n
Sep1 %t\n
Integer: %i[0,255]{Integer Tip 2}\n
Real 2: %r[-1.5,1.5]{Real Tip 2}\n
Sep2 %t\n
Angle: %a[0,360]{Angle Tip}\n
String: %s{String Tip}\n
List: %l|item1|item2|item3|{List Tip}\n
File: %f[OPEN|*.bmp;*.jpg|CURRENT|NO|NO]{File Tip}\n
Color: %c{Color Tip}\n
Sep3 %t\n
Multiline: %m{Multiline Tip}\n`, &_bool, &_int, &_real, &_int2, &_real2, &_angle, &_str, &_list, &_file, &_color, &_str2,
	)

	iup.Message("IupGetParam",
		fmt.Sprintf(`Boolean Value: %d\n
Integer: %v\n
Real 1: %v\n
Integer: %v\n
Real 2: %v\n
Angle: %v\n
String: %v\n
List Index: %v\n
FileName: %v\n
Color: %v\n
Multiline: %v`, _bool, _int, _real, _int2, _real2, _angle, _str, _list, _file, _color, _str2,
		),
	)
}

func action_cb(dlg iup.Ihandle, index int, userData uintptr) int {
	switch index {
	case iup.GETPARAM_INIT:
		fmt.Println("Init")
	case iup.GETPARAM_BUTTON1:
		fmt.Println("Button 1 (OK)")
	case iup.GETPARAM_BUTTON2:
		fmt.Println("Button 2 (Cancel)")
	case iup.GETPARAM_BUTTON3:
		fmt.Println("Button 3 (Help)")
	default:
		hParam := iup.Ihandle(dlg.GetPtr("PARAM", index))
		fmt.Printf("PARAM %d = %v\n", index, hParam.GetAttribute("VALUE"))
	}
	return 1
}
