// Shows a color-selection dialog.
package main

import (
	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	list := []string{
		"Blue",
		"Red",
		"Green",
		"Yellow",
		"Black",
		"White",
		"Gray",
		"Brown",
	}
	marks := []bool{false, false, false, false, true, true, false, false}

	if iup.ListDialog(2, "Color Selection", list, 0, 16, 5, &marks) == -1 {
		iup.Message("IupListDialog Example", "Operation canceled")
	} else {
		selection := ""
		for i, mark := range marks {
			if mark {
				selection += list[i] + " "
			}
		}

		if len(selection) == 0 {
			iup.Message("IupListDialog Example", "No option selected")
		} else {
			iup.Message("IupListDialog Example", "Selected colors: "+selection)
		}
	}
}
