package blocks

func ManIdByteToChar(b byte) byte {
	return b + 0x40
}

func AspectRatioByteToString(ar byte) string {
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

func AspectRatioToFloat(ar byte) float64 {
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
