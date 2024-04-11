package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/openpixelsystems/edid-tool/edid"
)

func main() {
	inFilePtr := flag.String("in", "", "Input file")
	outFilePtr := flag.String("out", "", "Output file")
	displayNamePtr := flag.String("name", "", "Display name")
	serialNumberPtr := flag.Uint("serial", 0, "Serial number")

	flag.Parse()

	if *inFilePtr == "" {
		fmt.Println("Input file is required")
		os.Exit(1)
	}

	f, err := os.Open(*inFilePtr)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	edidObj, err := edid.ReadEDID(data)
	if err != nil {
		fmt.Println(err)
	}

	warnings, _ := edidObj.Parse()
	fmt.Printf("\nParsing finished with %d warnings:\n", len(warnings))
	for _, warning := range warnings {
		fmt.Println("\t", warning)
	}

	//edidObj.ModifyManufacturerId([3]byte{'O', 'P', 'S'})
	nameDescriptor := edid.GenerateMonitorNameDescriptor(*displayNamePtr)
	if *displayNamePtr != "" {
		edidObj.ModifyDisplayDescriptor(2, nameDescriptor)
	}

	if *serialNumberPtr != 0 {
		edidObj.ModifySerialNumber(uint32(*serialNumberPtr))
	}

	//edidObj.Parse()

	edidData := edid.GenerateEDID(&edidObj)
	f, err = os.Create(*outFilePtr)
	f.Write(edidData)
}
