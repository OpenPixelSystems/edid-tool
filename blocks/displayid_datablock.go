package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (d *DisplayIDDataBlock) Type() interfaces.DataBlockType {
	return DATABLOCK_TYPE_DISPLAYID_DATABLOCK
}

func (d *DisplayIDDataBlock) Print() {
	fmt.Println("Display ID Data Block")
	fmt.Printf("\tTag: 0x%x\n", d.BlockType)
}

func (d *DisplayIDDataBlock) FromBytes(data []byte) (int, error) {
	offset := 0

	d.BlockType = data[offset]
	offset++

	switch d.BlockType {
	case DATABLOCK_TYPE_TILED_DISPLAY, DATABLOCK_TYPE_TILED_DISPLAY_LEGACY:
		fmt.Println("Tiled Display")
	case DATABLOCK_TYPE_VTB_TYPE_1:
		fmt.Println("VTB Type 1")
	default:
		fmt.Println("Unknown Display ID Data Block")
	}

	d.Raw = data[:offset]

	return offset, nil
}

func (d *DisplayIDDataBlock) ToBytes() []byte {
	return d.Raw
}

func (d *DisplayIDDataBlock) Len() int {
	return len(d.Raw)
}
