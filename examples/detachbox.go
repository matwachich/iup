package main

import (
	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	bt := iup.Button("Detach Me!").SetHandle("detach").SetCallback("ACTION", btn_detach_cb)

	ml := iup.MultiLine().SetAttributes(`EXPAND=YES, VISIBLELINES=5`)

	hbox := iup.Hbox(bt, ml).SetAttributes(`MARGIN=10x0`)

	dbox := iup.DetachBox(hbox).SetAttributes(map[string]string{
		"ORIENTATION": "VERTICAL",
		/*"SHOWGRIP":    "NO",
		"BARSIZE":     "0",
		"COLOR":       "255 0 0",*/
	}).SetHandle("dbox").SetCallback("DETACHED_CB", detached_cb)

	lbl := iup.Label("Label").SetAttributes(`EXPAND=VERTICAL`)

	bt2 := iup.Button("Restore Me!").SetAttributes(`EXPAND=YES, ACTIVE=NO`).SetHandle("restore").SetCallback("ACTION", btn_restore_cb)

	txt := iup.Text().SetAttributes(`EXPAND=HORIZONTAL`)

	dlg := iup.Dialog(iup.Vbox(
		dbox, lbl, bt2, txt,
	)).SetAttributes(map[string]string{
		"TITLE":      "IupDetachBox Exampel",
		"MARGIN":     "10x10",
		"GAP":        "10",
		"RASTERSIZE": "300x300",
	})

	iup.Show(dlg)
	iup.MainLoop()
}

func detached_cb(ih, newParent iup.Ihandle, x, y int) int {
	newParent.SetAttribute("TITLE", "New Dialog")

	iup.GetHandle("restore").SetAttribute("ACTIVE", "YES")
	iup.GetHandle("detach").SetAttribute("ACTIVE", "NO")

	return iup.DEFAULT
}

func btn_restore_cb(bt iup.Ihandle) int {
	iup.GetHandle("dbox").SetAttribute("RESTORE", nil) // default value

	bt.SetAttribute("ACTIVE", "NO")
	iup.GetHandle("detach").SetAttribute("ACTIVE", "YES")

	return iup.DEFAULT
}

func btn_detach_cb(bt iup.Ihandle) int {
	iup.GetHandle("dbox").SetAttribute("DETACH", nil) // default value

	bt.SetAttribute("ACTIVE", "NO")
	iup.GetHandle("restore").SetAttribute("ACTIVE", "YES")

	return iup.DEFAULT
}
