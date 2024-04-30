package blocks

import (
	"fmt"
)

func (et *EstablishedTimings) Print() {
	fmt.Println("Established Timings")
	if et.Timing1&ESTABLISHED_TIMINGS_720x400_70Hz != 0 {
		fmt.Println("\t720x400 @ 70Hz")
	}
	if et.Timing1&ESTABLISHED_TIMINGS_720x400_88Hz != 0 {
		fmt.Println("\t720x400 @ 88Hz")
	}
	if et.Timing1&ESTABLISHED_TIMINGS_640x480_60Hz != 0 {
		fmt.Println("\t640x480 @ 60Hz")
	}
	if et.Timing1&ESTABLISHED_TIMINGS_640x480_67Hz != 0 {
		fmt.Println("\t640x480 @ 67Hz")
	}
	if et.Timing1&ESTABLISHED_TIMINGS_640x480_72Hz != 0 {
		fmt.Println("\t640x480 @ 72Hz")
	}
	if et.Timing1&ESTABLISHED_TIMINGS_640x480_75Hz != 0 {
		fmt.Println("\t640x480 @ 75Hz")
	}
	if et.Timing1&ESTABLISHED_TIMINGS_800x600_56Hz != 0 {
		fmt.Println("\t800x600 @ 56Hz")
	}
	if et.Timing1&ESTABLISHED_TIMINGS_800x600_60Hz != 0 {
		fmt.Println("\t800x600 @ 60Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_800x600_72Hz != 0 {
		fmt.Println("\t800x600 @ 72Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_800x600_75Hz != 0 {
		fmt.Println("\t800x600 @ 75Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_832x624_75Hz != 0 {
		fmt.Println("\t832x624 @ 75Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_1024x768_87Hz != 0 {
		fmt.Println("\t1024x768 @ 87Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_1024x768_60Hz != 0 {
		fmt.Println("\t1024x768 @ 60Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_1024x768_70Hz != 0 {
		fmt.Println("\t1024x768 @ 70Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_1024x768_75Hz != 0 {
		fmt.Println("\t1024x768 @ 75Hz")
	}
	if et.Timing2&ESTABLISHED_TIMINGS_1280x1024_75Hz != 0 {
		fmt.Println("\t1280x1024 @ 75Hz")
	}
	if et.Timing3&ESTABLISHED_TIMINGS_1152x870_75Hz != 0 {
		fmt.Println("\t1152x870 @ 75Hz")
	}

}

func (et *EstablishedTimings) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the established timings
	et.Timing1 = data[offset]
	offset += 1

	et.Timing2 = data[offset]
	offset += 1

	et.Timing3 = data[offset]
	offset += 1

	// create the raw byte slice
	et.Raw = data[:offset]

	return offset, nil
}

func (et *EstablishedTimings) ToBytes() []byte {
	return et.Raw
}

func (et *EstablishedTimings) Len() int {
	return len(et.Raw)
}
