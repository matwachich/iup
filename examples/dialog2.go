package main

import (
	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	dlg := iup.Dialog(
		iup.Vbox(
			iup.Label("Very Long Text Label").SetAttributes(`EXPAND=YES, ALIGNMENT=ACENTER`),
			iup.Button("Quit").SetHandle("quit_bt_name").SetCallback("ACTION", func(ih iup.Ihandle) int {
				return iup.CLOSE
			}),
		).SetAttributes(`ALIGNMENT=ACENTER, MARGIN=10x10, GAP=10`),
	).SetAttributes(`TITLE="Dialog Title", DEFAULTESC=quit_bt_name`)

	iup.Show(dlg)
	iup.MainLoop()
}
