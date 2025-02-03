package blocks

import "github.com/openpixelsystems/edid-tool/interfaces"

const (
	DESCRIPTOR_TYPE_DETAILED_TIMING                interfaces.DescriptorType = 0x00 // Detailed timing
	DESCRIPTOR_TYPE_MANUFACTURER_SPECIFIC          interfaces.DescriptorType = 0x0F // Manufacturer specific
	DESCRIPTOR_TYPE_MONITOR_SERIAL_NUMBER          interfaces.DescriptorType = 0xFF // Monitor serial number
	DESCRIPTOR_TYPE_UNSPECIFIED                    interfaces.DescriptorType = 0xFE // Unspecified
	DESCRIPTOR_TYPE_RANGE_LIMITS                   interfaces.DescriptorType = 0xFD // Range limits
	DESCRIPTOR_TYPE_MONITOR_NAME                   interfaces.DescriptorType = 0xFC // Monitor name
	DESCRIPTOR_TYPE_WHITE_POINT_DATA               interfaces.DescriptorType = 0xFB // White point data
	DESCRIPTOR_TYPE_STANDARD_TIMING_IDENTIFICATION interfaces.DescriptorType = 0xFA // Standard timing identification
	DESCRIPTOR_TYPE_COLOR_POINT_DATA               interfaces.DescriptorType = 0xF9 // Color point data
	DESCRIPTOR_TYPE_CVT_3_BYTE_CODE                interfaces.DescriptorType = 0xF8 // CVT 3-byte code
	DESCRIPTOR_TYPE_ADDITIONAL_STANDARD_TIMING     interfaces.DescriptorType = 0xF7 // Additional standard timing
	DESCRIPTOR_TYPE_DUMMY                          interfaces.DescriptorType = 0x10 // Dummy

)
