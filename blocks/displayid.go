package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (did *DisplayIDExtension) Type() interfaces.BlockType {
	return BLOCK_TYPE_DISPLAYID
}

func (did *DisplayIDExtension) Print() {
	fmt.Println("\tDisplay ID")
	fmt.Printf("\t\tRevision: 0x%x\n", did.Revision)
	fmt.Printf("\t\tBytes in Section: 0x%x\n", did.BytesInSection)
	fmt.Printf("\t\tProduct Identifier: 0x%x\n", did.ProductIdentifier)
	fmt.Printf("\t\tExtension Count: 0x%x\n", did.ExtensionCount)

	for _, block := range did.DataBlocks {
		block.Print()
	}

	fmt.Printf("\t\tChecksum: 0x%x\n", did.Checksum)
}

func (did *DisplayIDExtension) FromBytes(data []byte) (int, error) {
	offset := 0

	did.Revision = data[offset]
	offset++

	did.BytesInSection = data[offset]
	offset++

	did.ProductIdentifier = data[offset]
	offset++

	did.ExtensionCount = data[offset]
	offset++

	//for i := 0; i < int(did.ExtensionCount); i++ {
	// parse the extension block
	DataBlock := DisplayIDDataBlock{}
	DataBlock.FromBytes(data[offset:])
	//nOffset, _ := DataBlock.FromBytes(data[offset:])
	//offset += nOffset
	did.DataBlocks = append(did.DataBlocks, &DataBlock)
	//}

	// override offset until done
	offset += 121

	did.Checksum = data[offset]
	offset++

	did.Raw = data[:offset]

	return offset, nil
}

func (did *DisplayIDExtension) ToBytes() []byte {
	return did.Raw
}

func (did *DisplayIDExtension) Len() int {
	return len(did.Raw)
}
