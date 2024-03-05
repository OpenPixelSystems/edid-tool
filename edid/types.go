package edid

type DTD struct {
	pixelClock                     [2]byte // 2 bytes pixel clock
	horizontalActiveLSB            byte    // 1 byte horizontal active LSB
	horizontalBlankingLSB          byte    // 1 byte horizontal blanking LSB
	horizontalMSB                  byte    // 1 byte horizontal MSB
	verticalActiveLSB              byte    // 1 byte vertical active LSB
	verticalBlankingLSB            byte    // 1 byte vertical blanking LSB
	verticalMSB                    byte    // 1 byte vertical MSB
	horizontalFrontPorchLSB        byte    // 1 byte horizontal front porch LSB
	horizontalSyncPulseLSB         byte    // 1 byte horizontal sync pulse LSB
	verticalFrontPorchSyncPulseLSB byte    // 1 byte vertical front porch sync pulse LSB
	horizontalVerticalMSB          byte    // 1 byte horizontal vertical MSB
	horizontalImageSize            byte    // 1 byte horizontal image size
	verticalImageSize              byte    // 1 byte vertical image size
	sizeMSB                        byte    // 1 byte size MSB
	horizontalBorder               byte    // 1 byte horizontal border
	verticalBorder                 byte    // 1 byte vertical border
	features                       byte    // 1 byte features
}

type VTBType1 struct {
	revision            byte   // 1 byte revision
	numberOfPayloadByte byte   // 1 byte number of payload bytes
	vtb                 []byte // variable length video timing block
}

type VTBDescriptor struct {
	pixelClockLBits byte // 1 byte pixel clock low bits
	pixelClockMBits byte // 1 byte pixel clock middle bits
	pixelClockHBits byte // 1 byte pixel clock high bits
	timingOptions   byte // 1 byte timing options
	hActiveLSB      byte // 1 byte horizontal active LSB
	hActiveMSB      byte // 1 byte horizontal active MSB
	hBlankingLSB    byte // 1 byte horizontal blanking LSB
	hBlankingMSB    byte // 1 byte horizontal blanking MSB
	hFrontPorchLSB  byte // 1 byte horizontal front porch LSB
	hFrontPorchMSB  byte // 1 byte horizontal front porch MSB
	hSyncWidthLSB   byte // 1 byte horizontal sync pulse LSB
	hSyncWidthMSB   byte // 1 byte horizontal sync pulse MSB
	vActiveLSB      byte // 1 byte vertical active LSB
	vActiveMSB      byte // 1 byte vertical active MSB
	vBlankingLSB    byte // 1 byte vertical blanking LSB
	vBlankingMSB    byte // 1 byte vertical blanking MSB
	vFrontPorchLSB  byte // 1 byte vertical front porch LSB
	vFrontPorchMSB  byte // 1 byte vertical front porch MSB
	vSyncWidthLSB   byte // 1 byte vertical sync pulse LSB
	vSyncWidthMSB   byte // 1 byte vertical sync pulse MSB
}

type EDID struct {
	edidData                [EDID_SIZE]byte                                         // 128 bytes edid
	ctaData                 [CTA_SIZE]byte                                          // 1 byte CTA extension tag
	rawData                 [EXTENDED_EDID_SIZE]byte                                // 256 bytes edid
	fixedHeader             [FIXED_HEADER_SIZE]byte                                 // 8 bytes fixed edid header
	manufacturerId          [MANUFACTURER_ID_SIZE]byte                              // 2 bytes manufacturer id
	productCode             [PRODUCT_CODE_SIZE]byte                                 // 2 bytes product code
	serialNumber            [SERIAL_NUMBER_SIZE]byte                                // 4 bytes serial number
	weekOfManufacture       byte                                                    // 1 byte week of manufacture
	yearOfManufacture       byte                                                    // 1 byte year of manufacture
	edidVersion             byte                                                    // 1 byte edid version
	edidRevision            byte                                                    // 1 byte edid revision
	basicDisplayParameters  [BASIC_DISPLAY_PARAMETERS_SIZE]byte                     // 5 bytes basic display parameters
	chromaticityCoordinates [CHROMATICITY_COORDINATES_SIZE]byte                     // 10 bytes chromaticity coordinates
	establishedTimings      [ESTABLISHED_TIMINGS_SIZE]byte                          // 3 bytes established timings
	standardTimings         [STANDARD_TIMINGS_COUNT][STANDARD_TIMINGS_SIZE]byte     // 8 standard timings
	displayDescriptor       [DISPLAY_DESCRIPTOR_COUNT][DISPLAY_DESCRIPTOR_SIZE]byte // 4 display descriptors
	extensionFlag           byte                                                    // 1 byte extension flag
	checksum                byte                                                    // 1 byte checksum
}
