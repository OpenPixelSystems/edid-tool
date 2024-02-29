package edid

import (
	"fmt"
	"encoding/binary"
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

)

const (
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
)


var (
	FIXED_HEADER_PATTERN = []byte{0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00} // 8 bytes fixed edid header
)

type CTA_EXT_BLOCK struct {
	extensionTag byte
	revision byte
	dtdStart byte
}

type EDID struct {
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

	return nil
}
