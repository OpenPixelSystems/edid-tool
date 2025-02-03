package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (cta *CTA861Extension) Type() interfaces.BlockType {
	return BLOCK_TYPE_CTA861
}

func (cta *CTA861Extension) Print() {
	fmt.Println("CTA861 Extension")
	fmt.Printf("\tRevision: 0x%x\n", cta.Revision)
	fmt.Println("\tDTD Offset: ", cta.DTDOffset)
	fmt.Println("\tFormat flags: ", cta.FormatAndLength>>4)
	fmt.Println("\tNumber of DTD blocks: ", cta.FormatAndLength&0x0F)
}

func (cta *CTA861Extension) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the CTA861Extension block

	// cta revision
	cta.Revision = data[offset]
	offset++

	// DTD offset
	// If no None DTD data is present this should be set to 0x04
	// If no DTD data is present this should be set to 0x00
	cta.DTDOffset = data[offset]
	offset++

	// Format info and number of DTD blocks
	cta.FormatAndLength = data[offset]
	offset++

	cta.Raw = data[:offset]

	return offset, nil
}

func (cta *CTA861Extension) ToBytes() []byte {
	return cta.Raw
}

func (cta *CTA861Extension) Len() int {
	return len(cta.Raw)
}
