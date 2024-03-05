package edid

const (
	FIXED_HEADER_SIZE             = 8   // 8 bytes fixed edid header
	MANUFACTURER_ID_SIZE          = 2   // 2 bytes manufacturer id
	PRODUCT_CODE_SIZE             = 2   // 2 bytes product code
	SERIAL_NUMBER_SIZE            = 4   // 4 bytes serial number
	WEEK_OF_MANUFACTURE_SIZE      = 1   // 1 byte week of manufacture
	YEAR_OF_MANUFACTURE_SIZE      = 1   // 1 byte year of manufacture
	EDID_VERSION_SIZE             = 1   // 1 byte edid version
	EDID_REVISION_SIZE            = 1   // 1 byte edid revision
	BASIC_DISPLAY_PARAMETERS_SIZE = 5   // 5 bytes basic display parameters
	VIDEO_INPUT_PARAMETERS_SIZE   = 1   // 1 byte video input parameters
	HORIZONTAL_SIZE_SIZE          = 1   // 1 byte horizontal size
	VERTICAL_SIZE_SIZE            = 1   // 1 byte vertical size
	DISPLAY_GAMMA_SIZE            = 1   // 1 byte display gamma
	SUPPORTED_FEATURES_SIZE       = 1   // 1 byte supported features
	CHROMATICITY_COORDINATES_SIZE = 10  // 10 bytes chromaticity coordinates
	ESTABLISHED_TIMINGS_SIZE      = 3   // 3 bytes established timings
	STANDARD_TIMINGS_SIZE         = 2   // 2 bytes standard timings
	STANDARD_TIMINGS_COUNT        = 8   // 8 standard timings
	DISPLAY_DESCRIPTOR_SIZE       = 18  // 18 bytes display descriptor
	DISPLAY_DESCRIPTOR_COUNT      = 4   // 4 display descriptors
	EXTENSION_FLAG_SIZE           = 1   // 1 byte extension flag
	CHECKSUM_SIZE                 = 1   // 1 byte checksum
	EDID_SIZE                     = 128 // 128 bytes edid
	CTA_SIZE                      = 128 // 128 bytes CTA extension
	EXTENDED_EDID_SIZE            = 256 // 256 bytes edid

	CTA_EXT_TAG_SIZE        = 1 // 1 byte CTA extension tag
	CTA_EXT_REVISION_SIZE   = 1 // 1 byte CTA extension revision
	CTA_EXT_DTD_START_SIZE  = 1 // 1 byte CTA extension DTD start
	CTA_EXT_NR_OF_DTDS_SIZE = 1 // 1 byte CTA extension number of DTDs

	ESTABLISHED_TIMINGS_720x400_70Hz   = 0x80 // 720x400 @ 70Hz
	ESTABLISHED_TIMINGS_720x400_88Hz   = 0x40 // 720x400 @ 88Hz
	ESTABLISHED_TIMINGS_640x480_60Hz   = 0x20 // 640x480 @ 60Hz
	ESTABLISHED_TIMINGS_640x480_67Hz   = 0x10 // 640x480 @ 67Hz
	ESTABLISHED_TIMINGS_640x480_72Hz   = 0x08 // 640x480 @ 72Hz
	ESTABLISHED_TIMINGS_640x480_75Hz   = 0x04 // 640x480 @ 75Hz
	ESTABLISHED_TIMINGS_800x600_56Hz   = 0x02 // 800x600 @ 56Hz
	ESTABLISHED_TIMINGS_800x600_60Hz   = 0x01 // 800x600 @ 60Hz
	ESTABLISHED_TIMINGS_800x600_72Hz   = 0x80 // 800x600 @ 72Hz
	ESTABLISHED_TIMINGS_800x600_75Hz   = 0x40 // 800x600 @ 75Hz
	ESTABLISHED_TIMINGS_832x624_75Hz   = 0x20 // 832x624 @ 75Hz
	ESTABLISHED_TIMINGS_1024x768_87Hz  = 0x10 // 1024x768 @ 87Hz
	ESTABLISHED_TIMINGS_1024x768_60Hz  = 0x08 // 1024x768 @ 60Hz
	ESTABLISHED_TIMINGS_1024x768_70Hz  = 0x04 // 1024x768 @ 70Hz
	ESTABLISHED_TIMINGS_1024x768_75Hz  = 0x02 // 1024x768 @ 75Hz
	ESTABLISHED_TIMINGS_1280x1024_75Hz = 0x01 // 1280x1024 @ 75Hz
	ESTABLISHED_TIMINGS_1152x870_75Hz  = 0x80 // 1152x870 @ 75Hz

	STD_TIMING_ASPECT_RATIO_16_10 = 0x00 // 16:10
	STD_TIMING_ASPECT_RATIO_4_3   = 0x01 // 4:3
	STD_TIMING_ASPECT_RATIO_5_4   = 0x02 // 5:4
	STD_TIMING_ASPECT_RATIO_16_9  = 0x03 // 16:9

	STD_TIMING_VERTICAL_FREQUENCY = 0x31 // 60Hz

	BDP_DIGITAL_INPUT               = 0x80 // Digital input
	BDP_BIT_DEPTH                   = 0x70 // Bit depth
	BDP_VIDEO_INTERFACE             = 0x0F // Video interface
	BDP_ANALOG_INPUT                = 0x00 // Analog input
	BDP_VIDEO_WHITE_AND_SYNC_LEVELS = 0x60 // Video white and sync levels
	BDP_BLANK_TO_BLACK_SETUP        = 0x10 // Blank to black setup
	BDP_SYNC_SIGNAL_LEVELS          = 0x08 // Sync signal levels
	BDP_COMPOSITE_SYNC              = 0x04 // Composite sync
	BDP_SYNC_ON_GREEN               = 0x02 // Sync on green
	BDP_VSYNC_SERRATED              = 0x01 // Vsync serrated

	FD_INTERLACED             = 0x80 // Interlaced
	FD_STEREO                 = 0x60 // Stereo
	FD_DIGITAL_ANALOG_SYNC    = 0x10 // Digital/Analog sync
	FD_ANALOG_SYNC            = 0x08 // Analog sync
	FD_ANALOG_SERRATED_VSYNC  = 0x04 // Analog serrated vsync
	FD_ANALOG_SYNC_ON_GREEN   = 0x02 // Analog sync on green
	FD_DIGITAL_COMPOSITE_SYNC = 0x08 // Digital composite sync
	FD_DIGITAL_SERRATION      = 0x04 // Digital serration
	FD_DIGITAL_VSYNC_POLARITY = 0x04 // Digital VSync polarity
	FD_DIGITAL_HSYNC_POLARITY = 0x02 // Digital HSync polarity
	FD_STEREO_MODE            = 0x01 // Stereo mode

	DTD_TYPE_MANUFACTURER_SPECIFIC          = 0x0F // Manufacturer specific
	DTD_TYPE_MONITOR_SERIAL_NUMBER          = 0xFF // Monitor serial number
	DTD_TYPE_UNSPECIFIED                    = 0xFE // Unspecified
	DTD_TYPE_RANGE_LIMITS                   = 0xFD // Range limits
	DTD_TYPE_MONITOR_NAME                   = 0xFC // Monitor name
	DTD_TYPE_WHITE_POINT_DATA               = 0xFB // White point data
	DTD_TYPE_STANDARD_TIMING_IDENTIFICATION = 0xFA // Standard timing identification
	DTD_TYPE_COLOR_POINT_DATA               = 0xF9 // Color point data
	DTD_TYPE_CVT_3_BYTE_CODE                = 0xF8 // CVT 3-byte code
	DTD_TYPE_ADDITIONAL_STANDARD_TIMING     = 0xF7 // Additional standard timing
	DTD_TYPE_DUMMY                          = 0x10 // Dummy

	CTA_EXT_TAG_AUDIO_DATA_BLOCK              = 0x01 // Audio data block
	CTA_EXT_TAG_VIDEO_DATA_BLOCK              = 0x02 // Video data block
	CTA_EXT_TAG_VENDOR_SPECIFIC_DATA_BLOCK    = 0x03 // Vendor-specific data block
	CTA_EXT_TAG_SPEAKER_ALLOCATION_DATA_BLOCK = 0x04 // Speaker allocation data block
	CTA_EXT_TAG_VESA_DTC_DATA_BLOCK           = 0x05 // VESA DTC data block
	CTA_EXT_TAG_VIDEO_FORMAT_DATA_BLOCK       = 0x06 // Video format data block
	CTA_EXT_TAG_USE_EXTENDED_TAG              = 0x07 // Use extended tag

	CTA_BLOCK_TILED_DISPLAY_LEGACY = 0x12 // Tiled display legacy
	CTA_BLOCK_TILED_DISPLAY        = 0x28 // Tiled display
	CTA_BLOCK_TILED_SIZE           = 25   // Tiled size
	CTA_BLOCK_VTB_TYPE_1           = 0x03 // VTB type 1
	CTA_VTB_TYPE_1_DESCRIPTOR_SIZE = 20   // VTB type 1 descriptor size

	TILE_ONE_TILE_BEHAVIOR  = 0x07 // One tile behavior
	TILE_N_TILE_BEHAVIOR    = 0x18 // N tile behavior
	TILE_BEZEL_DESCRIPTOR   = 0x40 // Bezel descriptor
	TILE_PHYSICAL_ENCLOSURE = 0x80 // Physical enclosure

)

var (
	FIXED_HEADER_PATTERN = []byte{0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x00} // 8 bytes fixed edid header
)
