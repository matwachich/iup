// Shows a multiline that ignores the treatment of the DEL key, canceling its effect.
package main

import (
	"github.com/matwachich/iup"
)

func ml_action(self iup.Ihandle, c int, after uintptr) int {
	if c == iup.K_g {
		return iup.IGNORE
	}
	return iup.DEFAULT
}

func main() {
	iup.Open()
	defer iup.Close()

	iup.Show(
		iup.Dialog(
			iup.MultiLine().SetCallback("ACTION", ml_action).SetAttributes(map[string]string{
				"EXPAND": "YES",
				"BORDER": "YES",
				"VALUE":  "I ignore the \"g\" key!",
			}),
		).SetAttributes(`TITLE=IupMultiline, SIZE=QUARTERxQUARTER`),
	)

	iup.MainLoop()
}
