package edid

import (
	"fmt"
	"encoding/binary"
	"strings"
)

const (
	FIXED_HEADER_SIZE = 8 // 8 bytes fixed edid header
	MANUFACTURER_ID_SIZE = 2 // 2 bytes manufacturer id
	PRODUCT_CODE_SIZE = 2 // 2 bytes product code
	SERIAL_NUMBER_SIZE = 4 // 4 bytes serial number
	WEEK_OF_MANUFACTURE_SIZE = 1 // 1 byte week of manufacture
	YEAR_OF_MANUFACTURE_SIZE = 1 // 1 byte year of manufacture
	EDID_VERSION_SIZE = 1 // 1 byte edid version
	EDID_REVISION_SIZE = 1 // 1 byte edid revision
	BASIC_DISPLAY_PARAMETERS_SIZE = 5 // 5 bytes basic display parameters
	VIDEO_INPUT_PARAMETERS_SIZE = 1 // 1 byte video input parameters
	HORIZONTAL_SIZE_SIZE = 1 // 1 byte horizontal size
	VERTICAL_SIZE_SIZE = 1 // 1 byte vertical size
	DISPLAY_GAMMA_SIZE = 1 // 1 byte display gamma
	SUPPORTED_FEATURES_SIZE = 1 // 1 byte supported features
	CHROMATICITY_COORDINATES_SIZE = 10 // 10 bytes chromaticity coordinates
	ESTABLISHED_TIMINGS_SIZE = 3 // 3 bytes established timings
	STANDARD_TIMINGS_SIZE = 2 // 2 bytes standard timings
	STANDARD_TIMINGS_COUNT = 8 // 8 standard timings
	DISPLAY_DESCRIPTOR_SIZE = 18 // 18 bytes display descriptor
	DISPLAY_DESCRIPTOR_COUNT = 4 // 4 display descriptors
	EXTENSION_FLAG_SIZE = 1 // 1 byte extension flag
	CHECKSUM_SIZE = 1 // 1 byte checksum
	EDID_SIZE = 128 // 128 bytes edid
	EXTENDED_EDID_SIZE = 256 // 256 bytes edid

	CTA_EXT_TAG_SIZE = 1 // 1 byte CTA extension tag
	CTA_EXT_REVISION_SIZE = 1 // 1 byte CTA extension revision
	CTA_EXT_DTD_START_SIZE = 1 // 1 byte CTA extension DTD start
	CTA_EXT_NR_OF_DTDS_SIZE = 1 // 1 byte CTA extension number of DTDs

	ESTABLISHED_TIMINGS_720x400_70Hz = 0x80
	ESTABLISHED_TIMINGS_720x400_88Hz = 0x40
	ESTABLISHED_TIMINGS_640x480_60Hz = 0x20
	ESTABLISHED_TIMINGS_640x480_67Hz = 0x10
	ESTABLISHED_TIMINGS_640x480_72Hz = 0x08
	ESTABLISHED_TIMINGS_640x480_75Hz = 0x04
	ESTABLISHED_TIMINGS_800x600_56Hz = 0x02
	ESTABLISHED_TIMINGS_800x600_60Hz = 0x01
	ESTABLISHED_TIMINGS_800x600_72Hz = 0x80
	ESTABLISHED_TIMINGS_800x600_75Hz = 0x40
	ESTABLISHED_TIMINGS_832x624_75Hz = 0x20
	ESTABLISHED_TIMINGS_1024x768_87Hz = 0x10
	ESTABLISHED_TIMINGS_1024x768_60Hz = 0x08
	ESTABLISHED_TIMINGS_1024x768_70Hz = 0x04
	ESTABLISHED_TIMINGS_1024x768_75Hz = 0x02
	ESTABLISHED_TIMINGS_1280x1024_75Hz = 0x01
	ESTABLISHED_TIMINGS_1152x870_75Hz = 0x80

	STD_TIMING_ASPECT_RATIO_16_10 = 0x00
	STD_TIMING_ASPECT_RATIO_4_3 = 0x01
	STD_TIMING_ASPECT_RATIO_5_4 = 0x02
	STD_TIMING_ASPECT_RATIO_16_9 = 0x03

	STD_TIMING_VERTICAL_FREQUENCY = 0x31

	BDP_DIGITAL_INPUT = 0x80
	BDP_BIT_DEPTH = 0x70
	BDP_VIDEO_INTERFACE = 0x0F
	BDP_ANALOG_INPUT = 0x00
	BDP_VIDEO_WHITE_AND_SYNC_LEVELS = 0x60
	BDP_BLANK_TO_BLACK_SETUP = 0x10
	BDP_SYNC_SIGNAL_LEVELS = 0x08
	BDP_COMPOSITE_SYNC = 0x04
	BDP_SYNC_ON_GREEN = 0x02
	BDP_VSYNC_SERRATED = 0x01

	FD_INTERLACED = 0x80
	FD_STEREO = 0x60
	FD_DIGITAL_ANALOG_SYNC = 0x10
	FD_ANALOG_SYNC = 0x08
	FD_ANALOG_SERRATED_VSYNC = 0x04
	FD_ANALOG_SYNC_ON_GREEN = 0x02
	FD_DIGITAL_COMPOSITE_SYNC = 0x08
	FD_DIGITAL_SERRATION = 0x04
	FD_DIGITAL_VSYNC_POLARITY = 0x04
	FD_DIGITAL_HSYNC_POLARITY = 0x02
	FD_STEREO_MODE = 0x01

	DTD_TYPE_MANUFACTURER_SPECIFIC = 0x0f
	DTD_TYPE_MONITOR_SERIAL_NUMBER = 0xFF
	DTD_TYPE_UNSPECIFIED = 0xFE
	DTD_TYPE_RANGE_LIMITS = 0xFD
	DTD_TYPE_MONITOR_NAME = 0xFC
	DTD_TYPE_WHITE_POINT_DATA = 0xFB
	DTD_TYPE_STANDARD_TIMING_IDENTIFICATION = 0xFA
	DTD_TYPE_COLOR_POINT_DATA = 0xF9
	DTD_TYPE_CVT_3_BYTE_CODE = 0xF8
	DTD_TYPE_ADDITIONAL_STANDARD_TIMING = 0xF7
	DTD_TYPE_DUMMY = 0x10

)


var (
	FIXED_HEADER_PATTERN = []byte{0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00} // 8 bytes fixed edid header
)

type CTA_EXT_BLOCK struct {
	extensionTag byte
	revision byte
	dtdStart byte
}

type DTD struct {
	pixelClock [2]byte
	horizontalActiveLSB byte
	horizontalBlankingLSB byte
	horizontalMSB byte
	verticalActiveLSB byte
	verticalBlankingLSB byte
	verticalMSB byte
	horizontalFrontPorchLSB byte
	horizontalSyncPulseLSB byte
	verticalFrontPorchSyncPulseLSB byte
	horizontalVerticalMSB byte
	horizontalImageSize byte
	verticalImageSize byte
	sizeMSB byte
	horizontalBorder byte
	verticalBorder byte
	features byte
}


type EDID struct {
	rawData [EXTENDED_EDID_SIZE]byte
	fixedHeader [FIXED_HEADER_SIZE]byte
	manufacturerId [MANUFACTURER_ID_SIZE]byte
	productCode [PRODUCT_CODE_SIZE]byte
	serialNumber [SERIAL_NUMBER_SIZE]byte
	weekOfManufacture byte
	yearOfManufacture byte
	edidVersion byte
	edidRevision byte
	basicDisplayParameters [BASIC_DISPLAY_PARAMETERS_SIZE]byte
	chromaticityCoordinates [CHROMATICITY_COORDINATES_SIZE]byte
	establishedTimings [ESTABLISHED_TIMINGS_SIZE]byte
	standardTimings [STANDARD_TIMINGS_COUNT][STANDARD_TIMINGS_SIZE]byte
	displayDescriptor [DISPLAY_DESCRIPTOR_COUNT][DISPLAY_DESCRIPTOR_SIZE]byte
	extensionFlag byte
	checksum byte
}

func ReadEDID(data []byte) (EDID, error) {
	var edid EDID
	if len(data) != EDID_SIZE && len(data) != EXTENDED_EDID_SIZE{
		return edid, fmt.Errorf("Invalid EDID size: %d", len(data))
	}
	copy(edid.rawData[:], data)

	offset := 0
	copy(edid.fixedHeader[:], data[offset:offset+FIXED_HEADER_SIZE])
	offset += FIXED_HEADER_SIZE
	copy(edid.manufacturerId[:], data[offset:offset+MANUFACTURER_ID_SIZE])
	offset += MANUFACTURER_ID_SIZE
	copy(edid.productCode[:], data[offset:offset+PRODUCT_CODE_SIZE])
	offset += PRODUCT_CODE_SIZE
	copy(edid.serialNumber[:], data[offset:offset+SERIAL_NUMBER_SIZE])
	offset += SERIAL_NUMBER_SIZE
	edid.weekOfManufacture = data[offset]
	offset += WEEK_OF_MANUFACTURE_SIZE
	edid.yearOfManufacture = data[offset]
	offset += YEAR_OF_MANUFACTURE_SIZE
	edid.edidVersion = data[offset]
	offset += EDID_VERSION_SIZE
	edid.edidRevision = data[offset]
	offset += EDID_REVISION_SIZE
	copy(edid.basicDisplayParameters[:], data[offset:offset+BASIC_DISPLAY_PARAMETERS_SIZE])
	offset += BASIC_DISPLAY_PARAMETERS_SIZE
	copy(edid.chromaticityCoordinates[:],data[offset:offset+CHROMATICITY_COORDINATES_SIZE])
	offset += CHROMATICITY_COORDINATES_SIZE
	copy(edid.establishedTimings[:], data[offset:offset+ESTABLISHED_TIMINGS_SIZE])
	offset += ESTABLISHED_TIMINGS_SIZE

	for i := 0; i < STANDARD_TIMINGS_COUNT; i++ {
		copy(edid.standardTimings[i][:], data[offset:offset+STANDARD_TIMINGS_SIZE])
		offset += STANDARD_TIMINGS_SIZE
	}

	for i := 0; i < DISPLAY_DESCRIPTOR_COUNT; i++ {
		copy(edid.displayDescriptor[i][:], data[offset:offset+DISPLAY_DESCRIPTOR_SIZE])
		offset += DISPLAY_DESCRIPTOR_SIZE
	}


	edid.extensionFlag = data[offset]
	offset += EXTENSION_FLAG_SIZE
	edid.checksum = data[offset]

	return edid, nil
}

func manIdByteToChar(b byte) byte {
	return b + 0x40
}

func parseBDP(bdp []byte) {
	displayType := bdp[0] & BDP_DIGITAL_INPUT
	if displayType != 0 {
		fmt.Println("\tDigital Input")
		bitDepth := (bdp[0] & BDP_BIT_DEPTH) >> 4
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
		videoInterface := bdp[0] & BDP_VIDEO_INTERFACE
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
		videoWhiteAndSyncLevels := (bdp[0] & BDP_VIDEO_WHITE_AND_SYNC_LEVELS) >> 5
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
		blankToBlackSetup := (bdp[0] & BDP_BLANK_TO_BLACK_SETUP) >> 4
		if blankToBlackSetup == 0x01 {
			fmt.Println("\tBlank to black setup (pedestal) expected")
		} else {
			fmt.Println("\tBlank to black setup (pedestal) not expected")
		}
		sepSyncLevels := (bdp[0] & BDP_SYNC_SIGNAL_LEVELS) >> 3
		if sepSyncLevels == 0x01 {
			fmt.Println("\tSeparate sync levels supported")
		} else {
			fmt.Println("\tSeparate sync levels not supported")
		}
		compositeSync := (bdp[0] & BDP_COMPOSITE_SYNC) >> 2
		if compositeSync == 0x01 {
			fmt.Println("\tComposite sync supported")
		} else {
			fmt.Println("\tComposite sync not supported")
		}
		syncOnGreen := (bdp[0] & BDP_SYNC_ON_GREEN) >> 1
		if syncOnGreen == 0x01 {
			fmt.Println("\tSync on green supported")
		} else {
			fmt.Println("\tSync on green not supported")
		}
		vsyncSerrated := bdp[0] & BDP_VSYNC_SERRATED
		if vsyncSerrated == 0x01 {
			fmt.Println("\tVsync serrated")
		} else {
			fmt.Println("\tVsync not serrated")
		}
	}
	fmt.Println("\tMaximum Image Size: ", bdp[1], "cm x ", bdp[2], "cm")
	displayGamma := bdp[3] + 100
	fmt.Println("\tDisplay Gamma: ", float32(displayGamma) / 100.0)
	supportedFeatures := bdp[4]
	if supportedFeatures & 0x80 != 0 {
		fmt.Println("\tDPMS standby supported")
	}
	if supportedFeatures & 0x40 != 0 {
		fmt.Println("\tDPMS suspend supported")
	}
	if supportedFeatures & 0x20 != 0 {
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

func parseChromaticityCoordinates(cc []byte) {
	redX := float64(((int((cc[0] & 0xc0) >> 6)) + int((cc[2]))<< 2)) / 1024.0
	redY := float64(((int((cc[0] & 0x30) >> 4)) + int((cc[3]))<< 2)) / 1024.0
	greenX := float64(((int((cc[0] & 0x0c) >> 2)) + int((cc[4]))<< 2)) / 1024.0
	greenY := float64(((int((cc[0] & 0x03))) + int((cc[5]))<< 2)) / 1024.0
	blueX := float64(((int((cc[1] & 0xc0) >> 6)) + int((cc[6]))<< 2)) / 1024.0
	blueY := float64(((int((cc[1] & 0x30) >> 4)) + int((cc[7]))<< 2)) / 1024.0
	whiteX := float64(((int((cc[1] & 0x0c) >> 2)) + int((cc[8]))<< 2)) / 1024.0
	whiteY := float64(((int((cc[1] & 0x03))) + int((cc[9]))<< 2)) / 1024.0
	fmt.Println("\tRed X: ", redX)
	fmt.Println("\tRed Y: ", redY)
	fmt.Println("\tGreen X: ", greenX)
	fmt.Println("\tGreen Y: ", greenY)
	fmt.Println("\tBlue X: ", blueX)
	fmt.Println("\tBlue Y: ", blueY)
	fmt.Println("\tWhite X: ", whiteX)
	fmt.Println("\tWhite Y: ", whiteY)
}

func parseEstablishedTimings(et []byte) {
	if et[0] & ESTABLISHED_TIMINGS_720x400_70Hz != 0 {
		fmt.Println("\t720x400 @ 70Hz")
	}
	if et[0] & ESTABLISHED_TIMINGS_720x400_88Hz != 0 {
		fmt.Println("\t720x400 @ 88Hz")
	}
	if et[0] & ESTABLISHED_TIMINGS_640x480_60Hz != 0 {
		fmt.Println("\t640x480 @ 60Hz")
	}
	if et[0] & ESTABLISHED_TIMINGS_640x480_67Hz != 0 {
		fmt.Println("\t640x480 @ 67Hz")
	}
	if et[0] & ESTABLISHED_TIMINGS_640x480_72Hz != 0 {
		fmt.Println("\t640x480 @ 72Hz")
	}
	if et[0] & ESTABLISHED_TIMINGS_640x480_75Hz != 0 {
		fmt.Println("\t640x480 @ 75Hz")
	}
	if et[0] & ESTABLISHED_TIMINGS_800x600_56Hz != 0 {
		fmt.Println("\t800x600 @ 56Hz")
	}
	if et[0] & ESTABLISHED_TIMINGS_800x600_60Hz != 0 {
		fmt.Println("\t800x600 @ 60Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_800x600_72Hz != 0 {
		fmt.Println("\t800x600 @ 72Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_800x600_75Hz != 0 {
		fmt.Println("\t800x600 @ 75Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_832x624_75Hz != 0 {
		fmt.Println("\t832x624 @ 75Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_1024x768_87Hz != 0 {
		fmt.Println("\t1024x768 @ 87Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_1024x768_60Hz != 0 {
		fmt.Println("\t1024x768 @ 60Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_1024x768_70Hz != 0 {
		fmt.Println("\t1024x768 @ 70Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_1024x768_75Hz != 0 {
		fmt.Println("\t1024x768 @ 75Hz")
	}
	if et[1] & ESTABLISHED_TIMINGS_1280x1024_75Hz != 0 {
		fmt.Println("\t1280x1024 @ 75Hz")
	}
	if et[2] & ESTABLISHED_TIMINGS_1152x870_75Hz != 0 {
		fmt.Println("\t1152x870 @ 75Hz")
	}
}

func aspectRatioByteToString(ar byte) string {
	switch ar {
	case STD_TIMING_ASPECT_RATIO_16_10:
		return "16:10"
	case STD_TIMING_ASPECT_RATIO_4_3:
		return "4:3"
	case STD_TIMING_ASPECT_RATIO_5_4:
		return "5:4"
	case STD_TIMING_ASPECT_RATIO_16_9:
		return "16:9"
	default:
		return "Reserved"
	}
}

func aspectRatioToFloat(ar byte) float64 {
	switch ar {
	case STD_TIMING_ASPECT_RATIO_16_10:
		return 16.0 / 10.0
	case STD_TIMING_ASPECT_RATIO_4_3:
		return 4.0 / 3.0
	case STD_TIMING_ASPECT_RATIO_5_4:
		return 5.0 / 4.0
	case STD_TIMING_ASPECT_RATIO_16_9:
		return 16.0 / 9.0
	default:
		return 0.0
	}
}

func parseStandardTimings(st [STANDARD_TIMINGS_COUNT][STANDARD_TIMINGS_SIZE]byte) {
	for i := 0; i < STANDARD_TIMINGS_COUNT; i++ {
		if st[i][0] == 0x01 && st[i][1] == 0x01 {
			fmt.Println("\tStandard Timing ", i, ": Unused")
			continue
		}
		horizontalActive := (int(st[i][0]) + 31) * 8
		aspectRatio := (st[i][1] & 0xC0) >> 6
		verticalActive := int((float64(horizontalActive) / aspectRatioToFloat(aspectRatio)))

		verticalFrequency := (st[i][1] & 0x3F) + 60
		fmt.Printf("\tStandard Timing %d: %d x %d @ %dHz (%s)\n", i, horizontalActive, verticalActive, verticalFrequency, aspectRatioByteToString(aspectRatio))
	}
}

func parseDisplayDescriptorFeatures(fd byte) {
	interlaced := (fd & FD_INTERLACED) >> 7
	if interlaced == 0x01 {
		fmt.Println("\t\t\tSignal type: Interlaced")
	} else {
		fmt.Println("\t\t\tSignal type: Progressive")
	}
	stereoMode := (fd & FD_STEREO) >> 4 | (fd & FD_STEREO_MODE)
	switch stereoMode {
	case 0x00:
	case 0x01:
		fmt.Println("\t\t\tNo stereo")
	case 0x02:
		fmt.Println("\t\t\tField sequential stereo, right image when stereo sync signal is high")
	case 0x03:
		fmt.Println("\t\t\tTwo way interleaved stereo, right image on even lines")
	case 0x04:
		fmt.Println("\t\t\tField sequential stereo, left image when stereo sync signal is high")
	case 0x05:
		fmt.Println("\t\t\tTwo way interleaved stereo, left image on even lines")
	case 0x06:
		fmt.Println("\t\t\tFour way interleaved stereo")
	case 0x07:
		fmt.Println("\t\t\tSide by side interleaved stereo")
	default:
		fmt.Println("\t\t\tReserved")
	}
	digitalSync := (fd & FD_DIGITAL_ANALOG_SYNC) >> 4
	if digitalSync == 0x01 {
		fmt.Println("\t\t\tDigital sync")

		digitalCompositeSync := (fd & FD_DIGITAL_COMPOSITE_SYNC) >> 3
		if digitalCompositeSync == 0x00 {
			fmt.Println("\t\t\tDigital composite sync")
			digitalSerrated := (fd & FD_DIGITAL_SERRATION) >> 2
			if digitalSerrated == 0x01 {
				fmt.Println("\t\t\tDigital serrated vsync")
			} else {
				fmt.Println("\t\t\tDigital vsync not serrated")
			}
			digitalHSyncPolarity := (fd & FD_DIGITAL_HSYNC_POLARITY) >> 1
			if digitalHSyncPolarity == 0x01 {
				fmt.Println("\t\t\tDigital HSync positive")
			} else {
				fmt.Println("\t\t\tDigital HSync negative")
			}
		} else {
			fmt.Println("\t\t\tDigital separate sync")
			digitalVsyncPolarity := (fd & FD_DIGITAL_VSYNC_POLARITY) >> 2
			if digitalVsyncPolarity == 0x01 {
				fmt.Println("\t\t\tDigital VSync positive")
			} else {
				fmt.Println("\t\t\tDigital VSync negative")
			}
			digitalHSyncPolarity := (fd & FD_DIGITAL_HSYNC_POLARITY) >> 1
			if digitalHSyncPolarity == 0x01 {
				fmt.Println("\t\t\tDigital HSync positive")
			} else {
				fmt.Println("\t\t\tDigital HSync negative")
			}
		}
	} else {
		fmt.Println("\t\t\tAnalog sync")
		analogSync := (fd & FD_ANALOG_SYNC) >> 3
		if analogSync == 0x01 {
			fmt.Println("\t\t\tAnalog bipolar composite sync ")
		} else {
			fmt.Println("\t\t\tAnalog composite sync")
		}
		analogSerratedVsync := (fd & FD_ANALOG_SERRATED_VSYNC) >> 2
		if analogSerratedVsync == 0x01 {
			fmt.Println("\t\t\tAnalog serrated vsync")
		} else {
			fmt.Println("\t\t\tAnalog vsync not serrated")
		}
		analogSyncOnGreen := (fd & FD_ANALOG_SYNC_ON_GREEN) >> 1
		if analogSyncOnGreen == 0x01 {
			fmt.Println("\t\t\tAnalog sync on all RGB signals")
		} else {
			fmt.Println("\t\t\tAnalog sync on green")
		}
	}
}

func parseDisplayTimingDescriptor(dd [DISPLAY_DESCRIPTOR_SIZE]byte) {
	var dtd DTD
	copy(dtd.pixelClock[:], dd[0:2])
	dtd.horizontalActiveLSB = dd[2]
	dtd.horizontalBlankingLSB = dd[3]
	dtd.horizontalMSB = dd[4]
	dtd.verticalActiveLSB = dd[5]
	dtd.verticalBlankingLSB = dd[6]
	dtd.verticalMSB = dd[7]
	dtd.horizontalFrontPorchLSB = dd[8]
	dtd.horizontalSyncPulseLSB = dd[9]
	dtd.verticalFrontPorchSyncPulseLSB = dd[10]
	dtd.horizontalVerticalMSB = dd[11]
	dtd.horizontalImageSize = dd[12]
	dtd.verticalImageSize = dd[13]
	dtd.sizeMSB = dd[14]
	dtd.horizontalBorder = dd[15]
	dtd.verticalBorder = dd[16]
	dtd.features = dd[17]
	fmt.Printf("\t\tPixel Clock: %f MHz\n", float64(binary.LittleEndian.Uint16([]byte(dtd.pixelClock[:])))/ 100.0)
	horizontalActive := (int(dtd.horizontalMSB & 0xf0 ) << 4) | int(dtd.horizontalActiveLSB)
	blanking := (int(dtd.horizontalMSB & 0x0f) << 8) | int(dtd.horizontalBlankingLSB)
	verticalActive := (int(dtd.verticalMSB & 0xf0) << 4) | int(dtd.verticalActiveLSB)
	verticalBlanking := (int(dtd.verticalMSB & 0x0f) << 8) | int(dtd.verticalBlankingLSB)
	horizontalFrontPorch := (int(dtd.horizontalVerticalMSB & 0xC0) << 2) | (int(dtd.horizontalFrontPorchLSB))
	horizontalSyncPulse := (int(dtd.horizontalVerticalMSB & 0x30) << 4) | (int(dtd.horizontalSyncPulseLSB))
	verticalFrontPorch := (int(dtd.horizontalVerticalMSB & 0x0c) << 2) |(int(dtd.verticalFrontPorchSyncPulseLSB & 0xF0) >> 4)
	verticalSyncPulse :=  (int(dtd.horizontalVerticalMSB & 0x03) << 4) | (int(dtd.verticalFrontPorchSyncPulseLSB & 0x0F))
	horizontalImageSize := (int(dtd.sizeMSB & 0xF0) << 4) | (int(dtd.horizontalImageSize))
	verticalImageSize := (int(dtd.sizeMSB & 0x0F) << 8) | (int(dtd.verticalImageSize))

	fmt.Printf("\t\tHorizontal Active: %d\n", horizontalActive)
	fmt.Printf("\t\tHorizontal Blanking: %d\n", blanking)
	fmt.Printf("\t\tVertical Active: %d\n", verticalActive)
	fmt.Printf("\t\tVertical Blanking: %d\n", verticalBlanking)
	fmt.Printf("\t\tHorizontal Front Porch: %d\n", horizontalFrontPorch)
	fmt.Printf("\t\tHorizontal Sync Pulse: %d\n", horizontalSyncPulse)
	fmt.Printf("\t\tVertical Front Porch: %d\n", verticalFrontPorch)
	fmt.Printf("\t\tVertical Sync Pulse: %d\n", verticalSyncPulse)
	fmt.Printf("\t\tImage Size: %dmm x %dmm\n", horizontalImageSize, verticalImageSize)
	fmt.Printf("\t\tHorizontal Border: %d\n", dtd.horizontalBorder)
	fmt.Printf("\t\tVertical Border: %d\n", dtd.verticalBorder)
	fmt.Printf("\t\tFeatures:\n")
	parseDisplayDescriptorFeatures(dtd.features)

}

func parseDisplayRangeLimitDescriptor(drd [DISPLAY_DESCRIPTOR_SIZE]byte) {
	horizontRateOffset := (drd[4] & 0x0C) >> 2
	verticalRateOffset := (drd[4] & 0x03)
	verticalFieldRateMin := int(drd[5])
	verticalFieldRateMax := int(drd[6])
	if verticalRateOffset == 0x3 {
		verticalFieldRateMin += 255
		verticalFieldRateMax += 255
	} else if verticalRateOffset == 0x2 {
		verticalFieldRateMax += 255
	}
	horizontalLineRateMin := int(drd[7])
	horizontalLineRateMax := int(drd[8])
	if horizontRateOffset == 0x3 {
		horizontalLineRateMin += 255
		horizontalLineRateMax += 255
	} else if horizontRateOffset == 0x2 {
		horizontalLineRateMax += 255
	}
	maxPixelClock := int(drd[9]) * 10
	// extendedTimingType := drd[10]
	// videoTimingParameters := drd[11:18]

	fmt.Printf("\t\tVertical Field Rate: %d - %d Hz\n", verticalFieldRateMin, verticalFieldRateMax)
	fmt.Printf("\t\tHorizontal Line Rate: %d - %d kHz\n", horizontalLineRateMin, horizontalLineRateMax)
	fmt.Printf("\t\tMax Pixel Clock: %d MHz\n", maxPixelClock)
}

func parseDisplayDescriptor(dd [DISPLAY_DESCRIPTOR_COUNT][DISPLAY_DESCRIPTOR_SIZE]byte) {
	for i := 0; i < DISPLAY_DESCRIPTOR_COUNT; i++ {
		if dd[i][0] != 0x00 && dd[i][1] != 0x00 {
			fmt.Printf("\tDisplay Descriptor %d\n", i)
			parseDisplayTimingDescriptor(dd[i])
		} else {
			descriptorType := dd[i][3]
			switch descriptorType {
			case DTD_TYPE_MANUFACTURER_SPECIFIC:
				fmt.Println("\tDisplay Descriptor ", i, ": Manufacturer specific: ", strings.Replace(string(dd[i][5:]), "\n","",-1))
			case DTD_TYPE_MONITOR_SERIAL_NUMBER:
				fmt.Println("\tDisplay Descriptor ", i, ": Monitor serial number: ", strings.Replace(string(dd[i][5:]), "\n","",-1))
			case DTD_TYPE_UNSPECIFIED:
				fmt.Println("\tDisplay Descriptor ", i, ": Unspecified")
			case DTD_TYPE_RANGE_LIMITS:
				fmt.Println("\tDisplay Descriptor ", i, ": Range limits")
				parseDisplayRangeLimitDescriptor(dd[i])
			case DTD_TYPE_MONITOR_NAME:
				fmt.Println("\tDisplay Descriptor ", i, ": Monitor name: ", strings.Replace(string(dd[i][5:]), "\n","",-1))
			case DTD_TYPE_WHITE_POINT_DATA:
				fmt.Println("\tDisplay Descriptor ", i, ": White point data")
			case DTD_TYPE_STANDARD_TIMING_IDENTIFICATION:
				fmt.Println("\tDisplay Descriptor ", i, ": Standard timing identification")
			case DTD_TYPE_COLOR_POINT_DATA:
				fmt.Println("\tDisplay Descriptor ", i, ": Color point data")
			case DTD_TYPE_CVT_3_BYTE_CODE:
				fmt.Println("\tDisplay Descriptor ", i, ": CVT 3-byte code")
			case DTD_TYPE_ADDITIONAL_STANDARD_TIMING:
				fmt.Println("\tDisplay Descriptor ", i, ": Additional standard timing")
			case DTD_TYPE_DUMMY:
				fmt.Println("\tDisplay Descriptor ", i, ": Dummy")
			default:
				fmt.Println("\tDisplay Descriptor ", i, ": Reserved")
			}
		}
	}
}

func (edid EDID) Checksum() bool {
	var sum byte
	for _, b := range edid.rawData[:EDID_SIZE-1] {
		sum += b
	}
	return (int(sum) + int(edid.checksum) == 256)
}

func (edid EDID) Parse() error {
	var manId [3]byte
	manId[0] = manIdByteToChar((edid.manufacturerId[0] >> 2) & 0x1F)
	manId[1] = manIdByteToChar(((edid.manufacturerId[0] & 0x3) << 3) | ((edid.manufacturerId[1] & 0xE0) >> 5))
	manId[2] = manIdByteToChar(edid.manufacturerId[1] & 0x1F)

	fmt.Println("Manufacturer ID: ", string(manId[:]))
	fmt.Printf("Product Code: %d\n", binary.LittleEndian.Uint16([]byte(edid.productCode[:])))
	fmt.Printf("Serial Number: %d\n", binary.LittleEndian.Uint32([]byte(edid.serialNumber[:])))
	fmt.Printf("Week of Manufacture: %d\n", edid.weekOfManufacture)
	fmt.Printf("Year of Manufacture: %d\n", int(edid.yearOfManufacture) + 1990)
	fmt.Printf("EDID Version: %d.%d\n", edid.edidVersion, edid.edidRevision)

	fmt.Printf("Basic Display Parameters:\n")
	parseBDP(edid.basicDisplayParameters[:])

	fmt.Printf("Chromaticity Coordinates:\n")
	parseChromaticityCoordinates(edid.chromaticityCoordinates[:])

	fmt.Printf("Established Timings:\n")
	parseEstablishedTimings(edid.establishedTimings[:])

	fmt.Printf("Standard Timings:\n")
	parseStandardTimings(edid.standardTimings)

	fmt.Printf("Display Timing Descriptor:\n")
	parseDisplayDescriptor(edid.displayDescriptor)

	fmt.Printf("Extension Flag: 0x%02X\n", edid.extensionFlag)
	fmt.Printf("Checksum: 0x%02X\n", edid.checksum)
	fmt.Printf("Checksum Valid: %t\n", edid.Checksum())
	return nil
}
