// Shows a dialog with a multiline, a text, a list and some buttons. You can test the multiline attributes by clicking on the buttons. Each button is related to an attribute. Select if you want to set or get an attribute using the dropdown list. The value in the text will be used as value when a button is pressed.
package main

import (
	"fmt"

	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	multi := iup.MultiLine().SetHandle("multi").SetAttributes(`EXPAND=YES`)
	text := iup.MultiLine().SetHandle("text").SetAttributes(`EXPAND=HORIZONTAL`)
	list := iup.List().SetHandle("list").SetAttributes(`DROPDOWN=YES, 1=SET, 2=GET`)

	btn_append := iup.Button("Append").SetCallback("ACTION", btn_append_cb)
	btn_insert := iup.Button("Insert").SetCallback("ACTION", btn_insert_cb)
	btn_border := iup.Button("Border").SetCallback("ACTION", btn_border_cb)
	btn_caret := iup.Button("Caret").SetCallback("ACTION", btn_caret_cb)
	btn_readonly := iup.Button("Read only").SetCallback("ACTION", btn_readonly_cb)
	btn_selection := iup.Button("Selection").SetCallback("ACTION", btn_selection_cb)
	btn_selectedtext := iup.Button("Selected Text").SetCallback("ACTION", btn_selectedtext_cb)
	btn_nc := iup.Button("Number of characters").SetCallback("ACTION", btn_nc_cb)
	btn_value := iup.Button("Value").SetCallback("ACTION", btn_value_cb)

	dlg := iup.Dialog(
		iup.Vbox(
			multi,
			iup.Hbox(text, list),
			iup.Hbox(btn_append, btn_insert, btn_border, btn_caret, btn_readonly, btn_selection),
			iup.Hbox(btn_selectedtext, btn_nc, btn_value),
		),
	).SetAttributes(`TITLE="IupMultiLine Example", SIZE=HALFxQUARTER`)

	iup.Show(dlg)
	iup.MainLoop()
}

func btn_append_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("APPEND")
	} else {
		getAttrib("APPEND")
	}
	return iup.DEFAULT
}

func btn_insert_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("INSERT")
	} else {
		getAttrib("INSERT")
	}
	return iup.DEFAULT
}

func btn_border_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("BORDER")
	} else {
		getAttrib("BORDER")
	}
	return iup.DEFAULT
}

func btn_caret_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("CARET")
	} else {
		getAttrib("CARET")
	}
	return iup.DEFAULT
}

func btn_readonly_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("READONLY")
	} else {
		getAttrib("READONLY")
	}
	return iup.DEFAULT
}

func btn_selection_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("SELECTION")
	} else {
		getAttrib("SELECTION")
	}
	return iup.DEFAULT
}

func btn_selectedtext_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("SELECTEDTEXT")
	} else {
		getAttrib("SELECTEDTEXT")
	}
	return iup.DEFAULT
}

func btn_nc_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("NC")
	} else {
		getAttrib("NC")
	}
	return iup.DEFAULT
}

func btn_value_cb() int {
	if iup.GetHandle("list").GetInt("VALUE") == 1 {
		setAttrib("VALUE")
	} else {
		getAttrib("VALUE")
	}
	return iup.DEFAULT
}

func setAttrib(attrib string) {
	multi, text := iup.GetHandle("multi"), iup.GetHandle("text")
	multi.SetAttribute(attrib, text.GetAttribute("VALUE"))

	iup.Message("Set Attribute", fmt.Sprintf("Attribute %q set with value %q", attrib, text.GetAttribute("VALUE")))
}

func getAttrib(attrib string) {
	multi, text := iup.GetHandle("multi"), iup.GetHandle("text")
	text.SetAttribute("VALUE", multi.GetAttribute(attrib))
}
