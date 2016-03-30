# IUP
GoLang wrapper for IUP GUI library (3.18)

This is a work in progress, but since it's a simple wrapper, I think it is ready for reasonably serious work.

It is only tested and developped for windows (actually on a 64bits Windows 7) since I have absolutly no experience with linux.

# Installation
To install and use this wrapper, you first need to download it

```
go get github.com/matwachich/iup
```

Then, you must create a directory inside the iup folder called "lib" and put in it the static IUP libraries. You also can if you want/need to use IM or CD create the folders "im" and "cd" inside the "lib" folder and put in it the IM and CD static libraries.

Then you can test this hello world example:

```go
package main

import "github.com/matwachich/iup"

func main() {
    iup.Open()
    defer iup.Close()
    
    dlg := iup.Dialog(iup.Vbox(
        iup.Label("Hello, world!"),
        iup.Button("Click me").SetCallback("ACTION", func(ih iup.Ihandle) int {
            return iup.CLOSE
        }),
    ).SetAttributes(`MARGIN=10x10,GAP=10`))
    
    iup.Show(dlg)
    iup.MainLoop()
}
```

# Additional libraries
The sub folders in the IUP folder are additional libraries of IUP. They all depend on the main iup lib.

They must be subfolders in order to be able to find the static libs.

# Known issues
iupweb doesn't work because it needs MS Visual C++ compiler (OLE...)

# Todo
Find a way to compile the IUP lib from source (as I said, absolutly no experience with unix tools (make, configure...)) to include directly the source code of IUP to this wrapper.
