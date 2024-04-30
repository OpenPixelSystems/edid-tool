package blocks

import (
	"encoding/binary"
	"fmt"
)

func (dtd *DetailedTimingDescriptor) Print() {
	fmt.Println("\tDetailed timing descriptor")
	fmt.Printf("\t\tPixel Clock: %f MHz\n", float64(binary.LittleEndian.Uint16([]byte(dtd.PixelClock[:])))/100.0)
	horizontalActive := (int(dtd.HorizontalMSB&0xf0) << 4) | int(dtd.HorizontalActiveLSB)
	blanking := (int(dtd.HorizontalMSB&0x0f) << 8) | int(dtd.HorizontalBlankingLSB)
	verticalActive := (int(dtd.VerticalMSB&0xf0) << 4) | int(dtd.VerticalActiveLSB)
	verticalBlanking := (int(dtd.VerticalMSB&0x0f) << 8) | int(dtd.VerticalBlankingLSB)
	horizontalFrontPorch := (int(dtd.HorizontalVerticalMSB&0xC0) << 2) | (int(dtd.HorizontalFrontPorchLSB))
	horizontalSyncPulse := (int(dtd.HorizontalVerticalMSB&0x30) << 4) | (int(dtd.HorizontalSyncPulseLSB))
	verticalFrontPorch := (int(dtd.HorizontalVerticalMSB&0x0c) << 2) | (int(dtd.VerticalFrontPorchSyncPulseLSB&0xF0) >> 4)
	verticalSyncPulse := (int(dtd.HorizontalVerticalMSB&0x03) << 4) | (int(dtd.VerticalFrontPorchSyncPulseLSB & 0x0F))
	horizontalImageSize := (int(dtd.SizeMSB&0xF0) << 4) | (int(dtd.HorizontalImageSize))
	verticalImageSize := (int(dtd.SizeMSB&0x0F) << 8) | (int(dtd.VerticalImageSize))

	fmt.Printf("\t\tHorizontal Active: %d\n", horizontalActive)
	fmt.Printf("\t\tVertical Active: %d\n", verticalActive)
	fmt.Printf("\t\tHorizontal Blanking: %d", blanking)
	fmt.Printf(" Front Porch: %d", horizontalFrontPorch)
	fmt.Printf(" Sync Pulse: %d\n", horizontalSyncPulse)
	fmt.Printf("\t\tVertical Blanking: %d", verticalBlanking)
	fmt.Printf(" Front Porch: %d", verticalFrontPorch)
	fmt.Printf(" Sync Pulse: %d\n", verticalSyncPulse)
	fmt.Printf("\t\tImage Size: %dmm x %dmm\n", horizontalImageSize, verticalImageSize)
	fmt.Printf("\t\tHorizontal Border: %d\n", dtd.HorizontalBorder)
	fmt.Printf("\t\tVertical Border: %d\n", dtd.VerticalBorder)
	fmt.Printf("\t\tFeatures:\n")

}

func (dtd *DetailedTimingDescriptor) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the pixel clock
	copy(dtd.PixelClock[:], data[offset:offset+PIXEL_CLOCK_SIZE])
	offset += PIXEL_CLOCK_SIZE

	// parse the horizontal active pixels
	dtd.HorizontalActiveLSB = data[offset]
	offset += HORIZONTAL_ACTIVE_SIZE

	dtd.HorizontalBlankingLSB = data[offset]
	offset += HORIZONTAL_BLANKING_SIZE

	dtd.HorizontalMSB = data[offset]
	offset += HORIZONTAL_MSB_SIZE

	dtd.VerticalActiveLSB = data[offset]
	offset += VERTICAL_ACTIVE_SIZE

	dtd.VerticalBlankingLSB = data[offset]
	offset += VERTICAL_BLANKING_SIZE

	dtd.VerticalMSB = data[offset]
	offset += VERTICAL_MSB_SIZE

	dtd.HorizontalFrontPorchLSB = data[offset]
	offset += HORIZONTAL_FRONT_PORCH_SIZE

	dtd.HorizontalSyncPulseLSB = data[offset]
	offset += HORIZONTAL_SYNC_PULSE_SIZE

	dtd.VerticalFrontPorchSyncPulseLSB = data[offset]
	offset += VERTICAL_FRONT_PORCH_SYNC_PULSE_SIZE

	dtd.HorizontalVerticalMSB = data[offset]
	offset += HORIZONTAL_VERTICAL_MSB_SIZE

	dtd.HorizontalImageSize = data[offset]
	offset += HORIZONTAL_IMAGE_SIZE

	dtd.VerticalImageSize = data[offset]
	offset += VERTICAL_IMAGE_SIZE

	dtd.SizeMSB = data[offset]
	offset += SIZE_MSB_SIZE

	dtd.HorizontalBorder = data[offset]
	offset += HORIZONTAL_BORDER_SIZE

	dtd.VerticalBorder = data[offset]
	offset += VERTICAL_BORDER_SIZE

	dtd.Features = data[offset]
	offset += FEATURES_SIZE

	// create the raw byte slice
	dtd.Raw = data[:offset]

	return offset, nil
}

func (dtd *DetailedTimingDescriptor) ToBytes() []byte {
	return dtd.Raw
}

func (dtd *DetailedTimingDescriptor) Len() int {
	return len(dtd.Raw)
}
