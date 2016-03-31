// Creates three labels, one using all attributes except for image, other with normal text and the last one with an image.
package main

import (
	"github.com/matwachich/iup"
)

var pixmap = []byte{
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 1, 1, 2, 2, 2, 1, 1, 1,
	1, 1, 2, 2, 1, 1, 1, 1, 1, 2, 2, 1, 1,
	1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 2, 2, 1,
	2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
}

func main() {
	iup.Open()
	defer iup.Close()

	img_star := iup.Image(13, 13, pixmap)
	img_star.SetAttribute("1", "0 0 0")
	img_star.SetAttribute("2", "0 198 0")
	iup.SetHandle("img_star", img_star)

	lbl := iup.Label("This label has the following attributes set:\nBGCOLOR = 255 255 0\nFGCOLOR = 0 0 255\nFONT = COURIER_NORMAL_14\nTITLE = All text contained here\nALIGNMENT = ACENTER")
	lbl.SetAttributes(`BGCOLOR="255 255 0",FGCOLOR="0 0 255",FONT=COURIER_NORMAL_14,ALIGNMENT=ACENTER`)

	lbl_explain := iup.Label("The label on the right has the image of a star")

	lbl_star := iup.Label()
	lbl_star.SetAttribute("IMAGE", img_star)

	dlg := iup.Dialog(
		iup.Vbox(
			lbl,
			iup.Hbox(lbl_explain, lbl_star),
		),
	)
	dlg.SetAttribute("TITLE", "IupLabel Example")

	iup.Show(dlg)
	iup.MainLoop()
}
