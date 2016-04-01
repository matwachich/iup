package main

import (
	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	iup.Menu(
		iup.Submenu("File",
			iup.Menu(
				iup.Item("Save\tCtrl+S").SetCallback("ACTION", item_save_cb),
				iup.Item("&Auto Save").SetHandle("item_autosave").SetAttributes(`VALUE=ON`).SetCallback("ACTION", item_autosave_cb),
				iup.Item("Exit").SetAttributes(`KEY="x"`).SetCallback("ACTION", item_exit_cb),
			),
		),
	).SetHandle("menu")

	dlg := iup.Dialog(
		iup.Vbox(
			iup.Text().SetAttributes(`VALUE="This is an empty text", EXPAND=HORIZONTAL`),
			iup.Button("Test").SetAttributes(`EXPAND=HORIZONTAL`),
		),
	)

	dlg.SetAttributes(map[string]interface{}{
		"TITLE":  "IupItem",
		"SIZE":   "120x",
		"MARGIN": "10x10",
		"GAP":    "10",
		"MENU":   "menu",
	})

	dlg.SetCallback("K_cX", item_exit_cb)
	dlg.SetCallback("K_cA", item_autosave_cb)
	dlg.SetCallback("K_cS", item_save_cb)

	iup.Show(dlg)
	iup.MainLoop()
}

func item_save_cb() int {
	iup.Message("IupItem", "Saved!")
	return iup.DEFAULT
}

func item_autosave_cb() int {
	item_autosave := iup.GetHandle("item_autosave")
	if item_autosave.GetInt("VALUE") != 0 {
		item_autosave.SetAttribute("VALUE", "OFF")
		iup.Message("IupItem", "AutoSave OFF")
	} else {
		item_autosave.SetAttribute("VALUE", "ON")
		iup.Message("IupItem", "AutoSave ON")
	}
	return iup.DEFAULT
}

func item_exit_cb() int {
	return iup.CLOSE
}
