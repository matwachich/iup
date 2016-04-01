// Creates a predefined color selection dialog. The user receives the color in the RGB format.
package main

import (
	"fmt"

	"github.com/matwachich/iup"
)

func main() {
	iup.Open()
	defer iup.Close()

	if ret, r, g, b := iup.GetColor(100, 100); ret != 0 {
		iup.Message("Color", fmt.Sprintf("RGB = %v %v %v", r, g, b))
	}
}
