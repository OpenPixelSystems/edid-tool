package main

import (
	"fmt"
	"flag"
	"io"
	"os"

	"github.com/openpixelsystems/edid-tool/edid"
)

func main() {
	inFilePtr := flag.String("in", "", "Input file")
	//outFilePtr := flag.String("out", "", "Output file")

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

	edidObj.Parse()
}
