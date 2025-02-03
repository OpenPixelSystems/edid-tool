package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (t *DataBlockTiledDisplay) Type() interfaces.BlockType {
	return DATABLOCK_TYPE_TILED_DISPLAY
}

func (t *DataBlockTiledDisplay) Print() {
	fmt.Println("DataBlockTiled Display")
}

func (t *DataBlockTiledDisplay) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the tiled display block
	t.Revision = data[offset]
	offset++

	// parse the tiled display block
	t.NumberOfPayloadBytes = data[offset]
	offset++

	// parse the tiled display block
	t.TileCaps = data[offset]
	offset++

	copy(t.TileTopologyAndLocation[:], data[offset:offset+3])
	offset += 3

	copy(t.TileSize[:], data[offset:offset+2])
	offset += 2

	copy(t.TileSize[:], data[offset:offset+5])
	offset += 5

	copy(t.TileDisplayTopologyID[:], data[offset:offset+9])
	offset += 9

	t.Raw = data[:offset]

	return offset, nil
}

func (t *DataBlockTiledDisplay) ToBytes() []byte {
	return t.Raw
}

func (t *DataBlockTiledDisplay) Len() int {
	return len(t.Raw)
}
