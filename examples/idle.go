// Creating a program that computes a number's factorial using the idle function.
package main

import (
	"fmt"

	"github.com/matwachich/iup"
)

var (
	step int     // iteration step
	fact float64 // last computed value
)

func main() {
	iup.Open()
	defer iup.Close()

	dlg := iup.Dialog(
		iup.Vbox(
			iup.Text().SetHandle("mens").SetAttributes(`SIZE=300x`).SetCallback("K_CR", func() int { return iup.CLOSE }),
			iup.Button("Calculate").SetCallback("ACTION", action_cb),
		),
	).SetAttribute("TITLE", "Idle Example")

	iup.Show(dlg)
	iup.MainLoop()
}

func action_cb() int {
	step, fact = 0, 1.0

	iup.SetFunction("IDLE_ACTION", idle_function)
	iup.SetAttribute(iup.GetHandle("mens"), "VALUE", "Computing...")

	return iup.DEFAULT
}

func idle_function() int {
	step++
	fact *= float64(step)

	iup.SetAttribute(iup.GetHandle("mens"), "VALUE", fmt.Sprintf("%d -> %10.4f", step, fact))

	if step == 100 {
		iup.SetFunction("IDLE_ACTION", nil)
	}

	return iup.DEFAULT
}
