package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (cta *CTAExtension) Type() interfaces.BlockType {
	return BLOCK_TYPE_CTA
}

func (cta *CTAExtension) Print() {
	fmt.Println("CTA Extension")
	fmt.Printf("Tag: 0x%x\n", cta.ExtensionTag)
	for _, block := range cta.Blocks {
		block.Print()
	}
	fmt.Printf("Checksum: 0x%x\n", cta.Checksum)
}

func (cta *CTAExtension) FromBytes(data []byte) (int, error) {
	offset := 0
	nOffset := 0

	// parse the CTAExtension block
	// cta tag
	cta.ExtensionTag = data[offset]
	offset++

	// parse the extension blocks
	switch cta.ExtensionTag {
	// Make const for this
	case 0x02:
		cta861 := CTA861Extension{}
		nOffset, _ = cta861.FromBytes(data[offset:])
		offset += nOffset
		cta.Blocks = append(cta.Blocks, &cta861)
	case 0x70:
		dID := DisplayIDExtension{}
		nOffset, _ = dID.FromBytes(data[offset:])
		offset += nOffset
		cta.Blocks = append(cta.Blocks, &dID)
	}

	cta.Checksum = data[offset]
	cta.Raw = data[:offset]

	return offset, nil
}

func (cta *CTAExtension) ToBytes() []byte {
	return cta.Raw
}

func (cta *CTAExtension) Len() int {
	return len(cta.Raw)
}
