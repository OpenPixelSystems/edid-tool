package blocks

import "github.com/openpixelsystems/edid-tool/interfaces"

type EDIDV2 struct {
	raw     [EXTENDED_EDID_SIZE]byte
	rawEDID [EDID_SIZE]byte
	rawCTA  [CTA_SIZE]byte
	Blocks  []interfaces.EDIDBlock
}

const (
	FH_MANUFACTURER_ID_SIZE = 3
)

type FixedHeader struct {
	Header            [FIXED_HEADER_SIZE]byte
	ManufacturerId    [FH_MANUFACTURER_ID_SIZE]byte
	ProductCode       uint16
	SerialNumber      uint32
	WeekOfManufacture byte
	YearOfManufacture byte
	EdidVersion       byte
	EdidRevision      byte
	Raw               []byte
}

type BasicDisplayParameters struct {
	VideoInputBitmap byte
	HorizontalSize   byte
	VerticalSize     byte
	DisplayGamma     byte
	FeaturesBitmap   byte
	Raw              []byte
}

type ChromaticityCoordinates struct {
	RedGreenLSB  byte
	BlueWhiteLSB byte
	RedXMSB      byte
	RedYMSB      byte
	GreenXMSB    byte
	GreenYMSB    byte
	BlueXMSB     byte
	BlueYMSB     byte
	WhiteXMSB    byte
	WhiteYMSB    byte
	Raw          []byte
}

type EstablishedTimings struct {
	Timing1 byte
	Timing2 byte
	Timing3 byte
	Raw     []byte
}

type StandardTimings struct {
	Timings [STANDARD_TIMINGS_COUNT][STANDARD_TIMINGS_SIZE]byte
	Raw     []byte
}

type MonitorDescriptor struct {
	reserved0 [3]byte
	Type      byte
	reserved1 byte
	Data      [DISPLAY_DESCRIPTOR_DATA_SIZE]byte
	Raw       []byte
}

type DetailedTimingDescriptor struct {
	PixelClock                     [2]byte // 2 bytes pixel clock
	HorizontalActiveLSB            byte    // 1 byte horizontal active LSB
	HorizontalBlankingLSB          byte    // 1 byte horizontal blanking LSB
	HorizontalMSB                  byte    // 1 byte horizontal MSB
	VerticalActiveLSB              byte    // 1 byte vertical active LSB
	VerticalBlankingLSB            byte    // 1 byte vertical blanking LSB
	VerticalMSB                    byte    // 1 byte vertical MSB
	HorizontalFrontPorchLSB        byte    // 1 byte horizontal front porch LSB
	HorizontalSyncPulseLSB         byte    // 1 byte horizontal sync pulse LSB
	VerticalFrontPorchSyncPulseLSB byte    // 1 byte vertical front porch sync pulse LSB
	HorizontalVerticalMSB          byte    // 1 byte horizontal vertical MSB
	HorizontalImageSize            byte    // 1 byte horizontal image size
	VerticalImageSize              byte    // 1 byte vertical image size
	SizeMSB                        byte    // 1 byte size MSB
	HorizontalBorder               byte    // 1 byte horizontal border
	VerticalBorder                 byte    // 1 byte vertical border
	Features                       byte    // 1 byte features
	Raw                            []byte
}

type DisplayTimingDescriptor struct {
	Descriptors []interfaces.DisplayDescriptor
	Raw         []byte
}
