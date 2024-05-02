package blocks

import (
	"fmt"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (bd *BasicDisplayParameters) Type() interfaces.BlockType {
	return BLOCK_TYPE_BASIC_DISPLAY_PARAMETERS
}

func (bd *BasicDisplayParameters) Print() {
	fmt.Println("Basic Display Parameters")
	displayType := bd.VideoInputBitmap & BDP_DIGITAL_INPUT
	if displayType != 0 {
		fmt.Println("\tDigital Input")
		bitDepth := (bd.VideoInputBitmap & BDP_BIT_DEPTH) >> 4
		switch bitDepth {
		case 0x00:
			fmt.Println("\tUndefined")
		case 0x01:
			fmt.Println("\t6 bits per color")
		case 0x02:
			fmt.Println("\t8 bits per color")
		case 0x03:
			fmt.Println("\t10 bits per color")
		case 0x04:
			fmt.Println("\t12 bits per color")
		case 0x05:
			fmt.Println("\t14 bits per color")
		case 0x06:
			fmt.Println("\t16 bits per color")
		default:
			fmt.Println("\tReserved")
		}
		videoInterface := bd.VideoInputBitmap & BDP_VIDEO_INTERFACE
		switch videoInterface {
		case 0x00:
			fmt.Println("\tUndefined")
		case 0x01:
			fmt.Println("\tDVI")
		case 0x02:
			fmt.Println("\tHDMI-a")
		case 0x03:
			fmt.Println("\tHDMI-b")
		case 0x04:
			fmt.Println("\tMDDI")
		case 0x05:
			fmt.Println("\tDisplayPort")
		default:
			fmt.Println("\tReserved")
		}
	} else {
		fmt.Println("\tAnalog Input")
		videoWhiteAndSyncLevels := (bd.VideoInputBitmap & BDP_VIDEO_WHITE_AND_SYNC_LEVELS) >> 5
		switch videoWhiteAndSyncLevels {
		case 0x00:
			fmt.Println("\t0.7/0.3 V")
		case 0x01:
			fmt.Println("\t0.714/0.286 V")
		case 0x02:
			fmt.Println("\t1.0/0.4 V")
		case 0x03:
			fmt.Println("\t0.7/0.0 V")
		default:
			fmt.Println("\tReserved")

		}
		blankToBlackSetup := (bd.VideoInputBitmap & BDP_BLANK_TO_BLACK_SETUP) >> 4
		if blankToBlackSetup == 0x01 {
			fmt.Println("\tBlank to black setup (pedestal) expected")
		} else {
			fmt.Println("\tBlank to black setup (pedestal) not expected")
		}
		sepSyncLevels := (bd.VideoInputBitmap & BDP_SYNC_SIGNAL_LEVELS) >> 3
		if sepSyncLevels == 0x01 {
			fmt.Println("\tSeparate sync levels supported")
		} else {
			fmt.Println("\tSeparate sync levels not supported")
		}
		compositeSync := (bd.VideoInputBitmap & BDP_COMPOSITE_SYNC) >> 2
		if compositeSync == 0x01 {
			fmt.Println("\tComposite sync supported")
		} else {
			fmt.Println("\tComposite sync not supported")
		}
		syncOnGreen := (bd.VideoInputBitmap & BDP_SYNC_ON_GREEN) >> 1
		if syncOnGreen == 0x01 {
			fmt.Println("\tSync on green supported")
		} else {
			fmt.Println("\tSync on green not supported")
		}
		vsyncSerrated := bd.VideoInputBitmap & BDP_VSYNC_SERRATED
		if vsyncSerrated == 0x01 {
			fmt.Println("\tVsync serrated")
		} else {
			fmt.Println("\tVsync not serrated")
		}
	}
	fmt.Println("\tMaximum Image Size: ", bd.HorizontalSize, "cm x ", bd.VerticalSize, "cm")
	displayGamma := bd.DisplayGamma + 100
	fmt.Println("\tDisplay Gamma: ", float32(displayGamma)/100.0)
	supportedFeatures := bd.FeaturesBitmap
	if supportedFeatures&0x80 != 0 {
		fmt.Println("\tDPMS standby supported")
	}
	if supportedFeatures&0x40 != 0 {
		fmt.Println("\tDPMS suspend supported")
	}
	if supportedFeatures&0x20 != 0 {
		fmt.Println("\tDPMS active-off supported")
	}

	dt := ((supportedFeatures & 0x18) >> 3)
	if displayType != 0x00 {
		switch dt {
		case 0x00:
			fmt.Println("\tDisplay type: RGB 4:4:4")
		case 0x01:
			fmt.Println("\tDisplay type: RGB 4:4:4 + YCrCb 4:4:4")
		case 0x02:
			fmt.Println("\tDisplay type: RGB 4:4:4 + YCrCb 4:2:2")
		case 0x03:
			fmt.Println("\tDisplay type: RGB 4:4:4 + YCrCb 4:4:4 + YCrCb 4:2:2")
		}
	} else {
		switch dt {
		case 0x00:
			fmt.Println("\tDisplay type: Monochrome/Grayscale")
		case 0x01:
			fmt.Println("\tDisplay type: RGB color")
		case 0x02:
			fmt.Println("\tDisplay type: Non-RGB color")
		case 0x03:
			fmt.Println("\tDisplay type: Undefined")
		}
	}
}

func (bd *BasicDisplayParameters) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the video input parameters
	bd.VideoInputBitmap = data[0]
	offset += VIDEO_INPUT_PARAMETERS_SIZE

	// parse the max horizontal image size
	bd.HorizontalSize = data[offset]
	offset += HORIZONTAL_SIZE_SIZE

	// parse the max vertical image size
	bd.VerticalSize = data[offset]
	offset += VERTICAL_SIZE_SIZE

	// parse the display gamma
	bd.DisplayGamma = data[offset]
	offset += DISPLAY_GAMMA_SIZE

	// parse the supported features
	bd.FeaturesBitmap = data[offset]
	offset += SUPPORTED_FEATURES_SIZE

	// create the raw byte slice
	bd.Raw = data[:offset]

	return offset, nil
}

func (bd *BasicDisplayParameters) ToBytes() []byte {
	return bd.Raw
}

func (bd *BasicDisplayParameters) Len() int {
	return len(bd.Raw)
}
