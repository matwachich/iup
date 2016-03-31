package main

import (
	"fmt"

	"github.com/matwachich/iup"
	"github.com/matwachich/iup/iupcontrols"
)

var text_red, text_green, text_blue iup.Ihandle

func main() {
	iup.Open()
	iupcontrols.Open()
	defer iup.Close()

	text_red, text_green, text_blue = iup.Text(), iup.Text(), iup.Text()

	colorbrowser := iupcontrols.ColorBrowser()

	//ISSUE the callback receives unsigned char, but it only works as this (byte doesn't work)
	colorbrowser.SetCallback("CHANGE_CB", func(self iup.Ihandle, r, g, b uintptr) int {
		drag_cb(self, r, g, b)
		return iup.DEFAULT
	})
	colorbrowser.SetCallback("DRAG_CB", drag_cb)

	vbox := iup.Vbox(
		iup.Fill(),
		text_red,
		iup.Fill(),
		text_green,
		iup.Fill(),
		text_blue,
		iup.Fill(),
	)

	hbox := iup.Hbox(
		colorbrowser, iup.Fill(), vbox,
	)

	dlg := iup.Dialog(hbox).SetAttributes(`TITLE="Color Browser"`)
	defer dlg.Destroy()

	iup.Show(dlg)
	iup.MainLoop()
}

func drag_cb(self iup.Ihandle, r, g, b uintptr) int {
	iup.SetAttribute(text_red, "VALUE", fmt.Sprintf("%d", r))
	iup.SetAttribute(text_green, "VALUE", fmt.Sprintf("%d", g))
	iup.SetAttribute(text_blue, "VALUE", fmt.Sprintf("%d", b))
	iup.LoopStep()
	return iup.DEFAULT
}
