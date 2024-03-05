package edid

import (
	"fmt"
)

func ReadEDID(data []byte) (EDID, error) {
	var edid EDID
	if len(data) != EDID_SIZE && len(data) != EXTENDED_EDID_SIZE {
		return edid, fmt.Errorf("Invalid EDID size: %d", len(data))
	}
	copy(edid.rawData[:], data)
	copy(edid.edidData[:], data[:EDID_SIZE])
	copy(edid.ctaData[:], data[EDID_SIZE:EDID_SIZE+CTA_SIZE])

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
	copy(edid.chromaticityCoordinates[:], data[offset:offset+CHROMATICITY_COORDINATES_SIZE])
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
