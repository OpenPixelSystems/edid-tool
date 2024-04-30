package blocks

import (
	"fmt"
)

func (st *StandardTimings) Print() {
	fmt.Println("Standard timings")
	for i := 0; i < STANDARD_TIMINGS_COUNT; i++ {
		if st.Timings[i][0] == 0x01 && st.Timings[i][1] == 0x01 {
			fmt.Println("\tStandard Timing ", i, ": Unused")
			continue
		}
		horizontalActive := (int(st.Timings[i][0]) + 31) * 8
		aspectRatio := (st.Timings[i][1] & 0xC0) >> 6
		verticalActive := int((float64(horizontalActive) / AspectRatioToFloat(aspectRatio)))

		verticalFrequency := (st.Timings[i][1] & 0x3F) + 60
		fmt.Printf("\tStandard Timing %d: %d x %d @ %dHz (%s)\n", i, horizontalActive, verticalActive, verticalFrequency, AspectRatioByteToString(aspectRatio))
	}

}

func (st *StandardTimings) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the fixed header
	for i := 0; i < STANDARD_TIMINGS_COUNT; i++ {
		copy(st.Timings[i][:], data[offset:offset+STANDARD_TIMINGS_SIZE])
		offset += STANDARD_TIMINGS_SIZE
	}

	// create the raw byte slice
	st.Raw = data[:offset]

	return offset, nil
}

func (st *StandardTimings) ToBytes() []byte {
	return st.Raw
}

func (st *StandardTimings) Len() int {
	return len(st.Raw)
}
