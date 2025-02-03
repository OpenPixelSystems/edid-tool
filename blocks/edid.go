package blocks

import (
	"fmt"
)

func (edid *EDIDV2) GetChecksum() (byte, error) {
	for _, block := range edid.Blocks {
		if block.Type() == BLOCK_TYPE_EXTENSION_CHECKSUM {
			return block.(*ExtensionAndChecksum).Checksum, nil
		}
	}
	return 0, fmt.Errorf("Checksum not found")
}

func (edid *EDIDV2) ChecksumValid() bool {
	var sum byte
	for _, b := range edid.Raw[:EDID_SIZE-1] {
		sum += b
	}
	checksum, err := edid.GetChecksum()
	if err != nil {
		return false
	}

	return (int(sum)+int(checksum) == 256)
}

func (edid *EDIDV2) FromBytes(data []byte) (int, error) {
	offset := 0

	copy(edid.Raw[:], data[:])
	copy(edid.RawEDID[:], data[:EDID_SIZE])
	copy(edid.RawCTA[:], data[EDID_SIZE:EXTENDED_EDID_SIZE])

	// parse the fixed header
	fh := FixedHeader{}
	nOffset, _ := fh.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &fh)

	// parse the basic display parameters
	bd := BasicDisplayParameters{}
	nOffset, _ = bd.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &bd)

	// parse the chromaticity coordinates
	cc := ChromaticityCoordinates{}
	nOffset, _ = cc.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &cc)

	// parse the established timings
	et := EstablishedTimings{}
	nOffset, _ = et.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &et)

	// parse the standard timings
	st := StandardTimings{}
	nOffset, _ = st.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &st)

	// parse the detailed timing
	dt := DisplayTimingDescriptor{}
	nOffset, _ = dt.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &dt)

	// parse the extension and checksum
	ec := ExtensionAndChecksum{}
	nOffset, _ = ec.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &ec)

	cta := CTAExtension{}
	nOffset, _ = cta.FromBytes(data[offset:])
	offset += nOffset
	edid.Blocks = append(edid.Blocks, &cta)

	return offset, nil
}

func (edid *EDIDV2) Print() {
	for _, block := range edid.Blocks {
		block.Print()
	}

	crcValid := edid.ChecksumValid()
	fmt.Println("Checksum valid: ", crcValid)
}
