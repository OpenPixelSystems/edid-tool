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
	copy(edid.Raw[:], data[:])
	copy(edid.RawEDID[:], data[:blocks.EDID_SIZE])
	copy(edid.RawCTA[:], data[blocks.EDID_SIZE:blocks.EXTENDED_EDID_SIZE])

	// parse the fixed header
	fh := blocks.FixedHeader{}
	nOffset, _ := fh.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &fh)

	// parse the basic display parameters
	bd := blocks.BasicDisplayParameters{}
	nOffset, _ = bd.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &bd)

	// parse the chromaticity coordinates
	cc := blocks.ChromaticityCoordinates{}
	nOffset, _ = cc.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &cc)

	// parse the established timings
	et := blocks.EstablishedTimings{}
	nOffset, _ = et.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &et)

	// parse the standard timings
	st := blocks.StandardTimings{}
	nOffset, _ = st.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &st)

	// parse the detailed timing
	dt := blocks.DisplayTimingDescriptor{}
	nOffset, _ = dt.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &dt)

	// parse the extension and checksum
	ec := blocks.ExtensionAndChecksum{}
	nOffset, _ = ec.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &ec)

	crcValid := edid.ChecksumValid()

	for _, block := range edid.Blocks {
		block.Print()
	}
	fmt.Println("Checksum valid: ", crcValid)
}
