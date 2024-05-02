package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (cc *ChromaticityCoordinates) Type() interfaces.BlockType {
	return BLOCK_TYPE_CHROMATICITY_COORDINATES
}

func (cc *ChromaticityCoordinates) Print() {
	fmt.Println("Fixed Header")
	redX := float64(((int((cc.RedGreenLSB & 0xc0) >> 6)) + int((cc.RedXMSB))<<2)) / 1024.0
	redY := float64(((int((cc.RedGreenLSB & 0x30) >> 4)) + int((cc.RedYMSB))<<2)) / 1024.0
	greenX := float64(((int((cc.RedGreenLSB & 0x0c) >> 2)) + int((cc.GreenXMSB))<<2)) / 1024.0
	greenY := float64(((int((cc.RedGreenLSB & 0x03))) + int((cc.GreenYMSB))<<2)) / 1024.0
	blueX := float64(((int((cc.BlueWhiteLSB & 0xc0) >> 6)) + int((cc.BlueXMSB))<<2)) / 1024.0
	blueY := float64(((int((cc.BlueWhiteLSB & 0x30) >> 4)) + int((cc.BlueYMSB))<<2)) / 1024.0
	whiteX := float64(((int((cc.BlueWhiteLSB & 0x0c) >> 2)) + int((cc.WhiteXMSB))<<2)) / 1024.0
	whiteY := float64(((int((cc.BlueWhiteLSB & 0x03))) + int((cc.WhiteYMSB))<<2)) / 1024.0
	fmt.Println("\tRed X:", redX, " Red Y:", redY)
	fmt.Println("\tGreen X:", greenX, " Green Y:", greenY)
	fmt.Println("\tBlue X:", blueX, " Blue Y:", blueY)
	fmt.Println("\tWhite X:", whiteX, " White Y:", whiteY)

}

func (cc *ChromaticityCoordinates) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the chromaticity coordinates
	cc.RedGreenLSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.BlueWhiteLSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.RedXMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.RedYMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.GreenXMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.GreenYMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.BlueXMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.BlueYMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.WhiteXMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	cc.WhiteYMSB = data[offset]
	offset += CHROMATICITY_BYTE_SIZE

	// create the raw byte slice
	cc.Raw = data[:offset]

	return offset, nil
}

func (cc *ChromaticityCoordinates) ToBytes() []byte {
	return cc.Raw
}

func (cc *ChromaticityCoordinates) Len() int {
	return len(cc.Raw)
}
