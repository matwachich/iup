//Creates four buttons. The first uses images, the second turns the first on and off, the third exits the application and the last does nothing
package main

import (
	"fmt"

	"github.com/matwachich/iup"
)

// Defines released button's image
var pixmap_release = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 4, 4, 4, 4, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 4, 4, 4, 4, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
}

//Defines pressed button's image
var pixmap_press = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 4, 4, 4, 4, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 4, 4, 4, 4, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
}

//Defines inactive button's image
var pixmap_inactive = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 4, 4, 4, 4, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 4, 4, 4, 4, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 4, 4, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
}

func main() {
	iup.Open()
	defer iup.Close()

	iup.SetGlobal("UTF8MODE", "YES")

	//create a Text
	text := iup.Text()

	//turns on read-only mode
	text.SetAttribute("READONLY", "YES")
	text.SetAttribute("EXPAND", "HORIZONTAL")

	//associate text with handle "text"
	iup.SetHandle("text", text)

	// ---
	//defines release button's image size
	img_release := iup.Image(16, 16, pixmap_release)

	//define released button's image colors
	img_release.SetAttributes(map[string]string{
		"1": "215 215 215",
		"2": "40 40 40",
		"3": "30 50 210",
		"4": "240 0 0",
	})

	//associate img_release with handle img_release
	iup.SetHandle("img_release", img_release)
	// ---

	// ---
	//defines press button's image size
	img_press := iup.Image(16, 16, pixmap_press)

	//define released button's image colors
	img_press.SetAttributes(map[string]string{
		"1": "40 40 40",
		"2": "215 215 215",
		"3": "0 20 180",
		"4": "210 0 0",
	})

	//associate img_press with handle img_press
	iup.SetHandle("img_press", img_press)
	// ---

	// ---
	//defines release button's image size
	img_inactive := iup.Image(16, 16, pixmap_inactive)

	//define released button's image colors
	img_inactive.SetAttributes(map[string]string{
		"1": "215 215 215",
		"2": "40 40 40",
		"3": "100 100 100",
		"4": "200 200 200",
	})

	//associate img_release with handle img_release
	iup.SetHandle("img_inactive", img_inactive)
	// ---

	//creates a button
	btn_image := iup.Button("Button with image", "btn_image")

	//set released, pressed an inactive button images
	btn_image.SetAttributes(`IMAGE=img_release, IMPRESS=img_press, IMINACTIVE=img_inactive`)

	//creates a button
	btn_big := iup.Button("Big useless button")

	//sets big button attributes
	btn_big.SetAttribute("SIZE", "EIGHTHxEIGHTH")

	//creates a button entitled Exit
	btn_exit := iup.Button("Exit")

	//creates a button entitled on/off
	btn_on_off := iup.Button("on/off")

	//creates a dialog with the four buttons and the text
	dlg := iup.Dialog(
		iup.Vbox(
			iup.Hbox(
				btn_image, btn_on_off, btn_exit,
			),
			text,
			btn_big,
		),
	)

	//sets dialog attributes
	iup.SetAttributes(dlg, `EXPAND=YES, TITLE="IupButton", RESIZE=NO, MENUBOX=NO, MAXBOX=NO, MINBOX=NO`)

	//register callbacks
	iup.SetCallback(btn_exit, "ACTION", btn_exit_cb)
	iup.SetCallback(btn_on_off, "ACTION", btn_on_off_cb)
	iup.SetCallback(btn_image, "BUTTON_CB", btn_image_cb)
	iup.SetCallback(btn_big, "BUTTON_CB", btn_big_cb)

	//run
	iup.Show(dlg)
	iup.MainLoop()
}

func btn_exit_cb(ih iup.Ihandle) int {
	return iup.CLOSE
}

func btn_on_off_cb(ih iup.Ihandle) int {
	btn_image := iup.GetHandle("btn_image")

	if btn_image.GetInt("ACTIVE") != 0 {
		btn_image.SetAttribute("ACTIVE", "NO")
	} else {
		iup.SetAttribute(btn_image, "ACTIVE", "YES")
	}

	return iup.DEFAULT
}

func btn_image_cb(ih iup.Ihandle, button, pressed int) int {
	if button == iup.BUTTON1 {
		text := iup.GetHandle("text")
		if pressed != 0 {
			text.SetAttribute("VALUE", "Red button pressed")
		} else {
			text.SetAttribute("VALUE", "Red button released")
		}
	}

	return iup.DEFAULT
}

func btn_big_cb(ih iup.Ihandle, button, pressed int) int {
	fmt.Printf("BUTTON_CB(button=%v,pressed=%v)\n", button, pressed)
	return iup.DEFAULT
}
