package main

import (
	"github.com/matwachich/iup"
	"github.com/matwachich/iup/iupmglplot"
)

func main() {
	iup.Open()
	iupmglplot.Open()
	defer iup.Close()

	lbl := iupmglplot.Label("\\int \\alpha \\sqrt{sin(\\pi x)^2 + \\gamma_{i_k}} dx").SetAttributes(`RASTERSIZE=400x80, LABELFONTSIZE=10`)

	dlg := iup.Dialog(iup.Vbox(lbl)).SetAttributes(`MARGIN=10x10, TITLE="IupMglLabel Example"`)

	iup.Show(dlg)
	iup.MainLoop()
}
