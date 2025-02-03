package blocks

import (
	"fmt"
	"strings"

	"github.com/openpixelsystems/edid-tool/interfaces"
)

func (md *MonitorDescriptor) Type() interfaces.DescriptorType {
	return md.DescriptorType
}

func (md *MonitorDescriptor) PrintAsRangeDescriptor() {
	horizontRateOffset := (md.Data[0] & 0x0C) >> 2
	verticalRateOffset := (md.Data[0] & 0x03)
	verticalFieldRateMin := int(md.Data[1])
	verticalFieldRateMax := int(md.Data[2])
	if verticalRateOffset == 0x3 {
		verticalFieldRateMin += 255
		verticalFieldRateMax += 255
	} else if verticalRateOffset == 0x2 {
		verticalFieldRateMax += 255
	}
	horizontalLineRateMin := int(md.Data[3])
	horizontalLineRateMax := int(md.Data[4])
	if horizontRateOffset == 0x3 {
		horizontalLineRateMin += 255
		horizontalLineRateMax += 255
	} else if horizontRateOffset == 0x2 {
		horizontalLineRateMax += 255
	}
	maxPixelClock := int(md.Data[5]) * 10
	extendedTimingType := md.Data[6]

	videoTimingParameters := md.Data[7:]

	fmt.Printf("\t\tVertical Field Rate: %d - %d Hz\n", verticalFieldRateMin, verticalFieldRateMax)
	fmt.Printf("\t\tHorizontal Line Rate: %d - %d kHz\n", horizontalLineRateMin, horizontalLineRateMax)
	fmt.Printf("\t\tMax Pixel Clock: %d MHz\n", maxPixelClock)
	switch extendedTimingType {
	case 0x00:
		fmt.Println("\t\tDefault GTF")
	case 0x01:
		fmt.Println("\t\tNo timing information")
	case 0x02:
		fmt.Println("\t\tSecondary GTF supported")
	case 0x03:
		fmt.Println("\t\tCVT supported")
	default:
		fmt.Println("\t\tReserved")
	}
	fmt.Println("\t\tVideo Timing Parameters: ", videoTimingParameters)

}

func (md *MonitorDescriptor) Print() {
	switch md.DescriptorType {
	case DESCRIPTOR_TYPE_MANUFACTURER_SPECIFIC:
		fmt.Println("\tManufacturer specific: ", strings.Replace(string(md.Data[:]), "\n", "", -1))
	case DESCRIPTOR_TYPE_MONITOR_SERIAL_NUMBER:
		fmt.Println("\tMonitor serial number: ", strings.Replace(string(md.Data[:]), "\n", "", -1))
	case DESCRIPTOR_TYPE_UNSPECIFIED:
		fmt.Println("\tUnspecified")
	case DESCRIPTOR_TYPE_RANGE_LIMITS:
		fmt.Println("\tRange limits")
		md.PrintAsRangeDescriptor()
	case DESCRIPTOR_TYPE_MONITOR_NAME:
		fmt.Println("\tMonitor name: ", strings.Replace(string(md.Data[:]), "\n", "", -1))
	case DESCRIPTOR_TYPE_WHITE_POINT_DATA:
		fmt.Println("\tWhite point data")
	case DESCRIPTOR_TYPE_STANDARD_TIMING_IDENTIFICATION:
		fmt.Println("\tStandard timing identification")
	case DESCRIPTOR_TYPE_COLOR_POINT_DATA:
		fmt.Println("\tColor point data")
	case DESCRIPTOR_TYPE_CVT_3_BYTE_CODE:
		fmt.Println("\tCVT 3-byte code")
	case DESCRIPTOR_TYPE_ADDITIONAL_STANDARD_TIMING:
		fmt.Println("\tAdditional standard timing")
	case DESCRIPTOR_TYPE_DUMMY:
		fmt.Println("\tDummy")
	default:
		fmt.Println("\tReserved")

	}
}

func (md *MonitorDescriptor) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the monitor descriptor
	offset += 3 // skip 3 reserved bytes

	md.DescriptorType = interfaces.DescriptorType(data[offset])
	offset += MONITOR_DESCRIPTOR_TYPE_SIZE

	offset += 1 // skip reserved byte

	// -1 for go array addressing.
	fmt.Println("To Copy:   ", data[offset:offset+DISPLAY_DESCRIPTOR_DATA_SIZE])
	copy(md.Data[:], data[offset:offset+DISPLAY_DESCRIPTOR_DATA_SIZE+1])
	offset += DISPLAY_DESCRIPTOR_DATA_SIZE

	// create the raw byte slice
	md.Raw = data[:offset]

	return offset, nil
}

func (md *MonitorDescriptor) ToBytes() []byte {
	offset := 0
	data := make([]byte, DISPLAY_DESCRIPTOR_SIZE)
	// byte 0 and 0 are 0 because it's not a detailed timing descriptor
	data[offset] = 0x00
	offset++

	data[offset] = 0x00
	offset++

	// 1 byte reserved
	data[offset] = 0x00
	offset++

	// 1 byte type
	data[offset] = byte(md.DescriptorType)
	offset++

	// 1 byte reserved
	if md.DescriptorType == DESCRIPTOR_TYPE_RANGE_LIMITS {
		data[offset] = md.Raw[offset]
	} else {
		data[offset] = 0x00
	}
	offset++

	// copy the data
	copy(data[offset:], md.Data[:])

	return data
}

func (md *MonitorDescriptor) Len() int {
	return len(md.Raw)
}
