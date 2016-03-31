package main

import (
	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	// create background box
	btn := iup.BackgroundBox(
		iup.Frame(
			iup.Vbox(
				iup.Button("This button does nothing"),
				iup.Text(),
			).SetAttributes(`MARGIN=0x0`),
		),
	)
	//btn.SetAttribute("BGCOLOR", "0 128 0")
	//btn.SetAttribute("BORDER", "YES")

	// create dialog
	dlg := iup.Dialog(
		iup.Vbox(btn),
	)
	iup.SetAttributes(dlg, `MARGIN=10x10, GAP=10, TITLE="IupBackgroundBox Example"`)

	iup.ShowXY(dlg, iup.CENTER, iup.CENTER)

	iup.MainLoop()
}
