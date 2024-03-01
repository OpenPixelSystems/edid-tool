package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"

	"github.com/openpixelsystems/edid-tool/edid"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(unsafe.Sizeof(edid.EDID{}))

	f, err := os.Open("/home/bravl/tmp/edid.hex")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	edidObj, err := edid.ReadEDID(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(edidObj)
	edidObj.Parse()
}
