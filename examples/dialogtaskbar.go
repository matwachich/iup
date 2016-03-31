// this example must be compiled with the manifest
// only for visual C++ 10 and above
package main

import (
	"github.com/matwachich/iup"
	"github.com/matwachich/iup/iupcontrols"
)

var (
	increment                      float32 = 0.01
	timer, progressbar1, btn_pause iup.Ihandle
)

func main() {
	iup.Open()
	iup.ImageLibOpen()
	iupcontrols.Open()
	defer iup.Close()

	timer = iup.Timer().SetAttribute("TIME", "100").SetCallback("ACTION_CB", time_cb)

	progressbar1 = iup.ProgressBar().SetAttributes(`EXPAND=YES, DASHED=YES`)

	btn_pause = iup.Button("").SetAttributes(`IMAGE="IUP_MediaPause", TIP="Play/Pause"`).SetCallback("ACTION", btn_pause_cb)
	btn_restart := iup.Button("").SetAttributes(`IMAGE="IUP_MediaGotoBegin", TIP="Restart"`).SetCallback("ACTION", btn_restart_cb)
	btn_decelerate := iup.Button("").SetAttributes(`IMAGE="IUP_MediaRewind", TIP="Decelerate"`).SetCallback("ACTION", btn_decelerate_cb)
	btn_accelerate := iup.Button("").SetAttributes(`IMAGE="IUP_MediaForward", TIP="Accelerate"`).SetCallback("ACTION", btn_accelerate_cb)
	btn_noprogress := iup.Button("").SetAttributes(`IMAGE="IUP_MessageInfo", TIP="No Progress"`).SetCallback("ACTION", btn_noprogress_cb)
	btn_indeterminate := iup.Button("").SetAttributes(`IMAGE="IUP_ToolsSettings", TIP="Indeterminate"`).SetCallback("ACTION", btn_indeterminate_cb)
	btn_error := iup.Button("").SetAttributes(`IMAGE="IUP_MessageError", TIP="Error"`).SetCallback("ACTION", btn_error_cb)

	hbox := iup.Hbox(
		iup.Fill(), btn_pause, btn_restart, btn_decelerate, btn_accelerate, btn_noprogress, btn_indeterminate, btn_error, iup.Fill(),
	)

	vbox := iup.Hbox(iup.Vbox(progressbar1, hbox)).SetAttributes(`MARGIN=10x10, GAP=10`)

	dlg := iup.Dialog(vbox).SetHandle("mydialog").SetAttributes(map[string]string{
		"TITLE":                "IupDialog: Progress bar in the Windows 7 Taskbar",
		"RASTERSIZE":           "420x",
		"TASKBARPROGRESS":      "YES",
		"TASKBARPROGRESSSTATE": "NORMAL",
	})

	iup.Show(dlg)
	timer.SetAttribute("RUN", "YES")
	iup.MainLoop()
}

func time_cb() int {
	// update prograssbar
	value := progressbar1.GetFloat("VALUE")
	value += increment
	if value > 1.0 {
		value = 0
	}
	progressbar1.SetAttribute("VALUE", value)

	// update taskbar
	iup.GetHandle("mydialog").SetAttribute("TASKBARPROGRESSVALUE", int(value*100))

	return iup.DEFAULT
}

func btn_pause_cb() int {
	if timer.GetInt("RUN") != 0 {
		timer.SetAttribute("RUN", "NO")
		btn_pause.SetAttribute("IMAGE", "IUP_MediaPlay")
		iup.GetHandle("mydialog").SetAttribute("TASKBARPROGRESSSTATE", "PAUSED")
	} else {
		timer.SetAttribute("RUN", "YES")
		btn_pause.SetAttribute("IMAGE", "IUP_MediaPause")
		iup.GetHandle("mydialog").SetAttribute("TASKBARPROGRESSSTATE", "NORMAL")
	}

	return iup.DEFAULT
}

func btn_restart_cb() int {
	progressbar1.SetAttribute("VALUE", "0")
	iup.GetHandle("mydialog").SetAttribute("TASKBARPROGRESSSTATE", "NOPROGRESS")

	return iup.DEFAULT
}

func btn_accelerate_cb() int {
	increment *= 2
	return iup.DEFAULT
}

func btn_decelerate_cb() int {
	increment /= 2
	return iup.DEFAULT
}

func btn_indeterminate_cb() int {
	timer.SetAttribute("RUN", "NO")
	btn_pause.SetAttribute("IMAGE", "IUP_MediaPlay")
	iup.GetHandle("mydialog").SetAttribute("TASKBARPROGRESSSTATE", "INDETERMINATE")

	return iup.DEFAULT
}

func btn_noprogress_cb() int {
	timer.SetAttribute("RUN", "NO")
	btn_pause.SetAttribute("IMAGE", "IUP_MediaPlay")
	iup.GetHandle("mydialog").SetAttribute("TASKBARPROGRESSSTATE", "NOPROGRESS")

	return iup.DEFAULT
}

func btn_error_cb() int {
	timer.SetAttribute("RUN", "NO")
	btn_pause.SetAttribute("IMAGE", "IUP_MediaPlay")
	iup.GetHandle("mydialog").SetAttribute("TASKBARPROGRESSSTATE", "ERROR")

	return iup.DEFAULT
}
