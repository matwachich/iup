// Draws a frame around a button. Note that "FGCOLOR" is added to the frame but it is inherited by the button.
package main

import (
	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	frame := iup.Frame(
		iup.Hbox(
			iup.Fill(),
			iup.Label("IupFrame Attributes:\nFGCOLOR = \"255 0 0\"\nSIZE = \"EIGHTHxEIGHTH\"\nTITLE = \"This is the frame\"\nMARGIN = \"10x10\""),
			iup.Fill(),
		),
	).SetAttributes("FGCOLOR=\"255 0 0\", SIZE=EIGHTHxEIGHTH, TITLE=\"This is the frame\", MARGIN=10x10")

	dlg := iup.Dialog(frame).SetAttribute("TITLE", "IupFrame")
	iup.Show(dlg)
	iup.MainLoop()
}
