package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/openpixelsystems/edid-tool/blocks"
)

func main() {
	inFilePtr := flag.String("in", "", "Input file")
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

	offset := 0

	edid := blocks.EDIDV2{}
	edid.FromBytes(data[offset:])
	edid.Print()
}
