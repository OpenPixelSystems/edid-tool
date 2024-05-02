package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (ec *ExtensionAndChecksum) Type() interfaces.BlockType {
	return BLOCK_TYPE_EXTENSION_CHECKSUM
}

func (ec *ExtensionAndChecksum) Print() {
	fmt.Println("Extension and checksum")
	fmt.Println("\tNumber of extension blocks: ", ec.Extensions)
	fmt.Println("\tChecksum: ", ec.Checksum)
}

func (ec *ExtensionAndChecksum) FromBytes(data []byte) (int, error) {
	offset := 0

	ec.Extensions = data[offset]
	offset++

	ec.Checksum = data[offset]
	offset++

	// create the raw byte slice
	ec.Raw = data[:offset]

	return offset, nil
}

func (ec *ExtensionAndChecksum) ToBytes() []byte {
	return ec.Raw
}

func (ec *ExtensionAndChecksum) Len() int {
	return len(ec.Raw)
}
