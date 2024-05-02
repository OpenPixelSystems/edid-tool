package blocks

import "fmt"

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
