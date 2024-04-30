package blocks

import (
	"fmt"
)

func (dtd *DisplayTimingDescriptor) Print() {
	fmt.Println("Display timing descriptor")
	for i, descriptor := range dtd.Descriptors {
		fmt.Println("\tDisplay Timing Descriptor ", i)
		descriptor.Print()
	}
}

func (dtd *DisplayTimingDescriptor) FromBytes(data []byte) (int, error) {
	offset := 0

	// parse the display timing descriptor
	for i := 0; i < DISPLAY_TIMING_DESCRIPTOR_COUNT; i++ {
		if data[offset] != 0x00 && data[offset+1] != 0x00 {
			dd := DetailedTimingDescriptor{}
			nOffset, _ := dd.FromBytes(data[offset:])
			offset += nOffset
			dtd.Descriptors = append(dtd.Descriptors, &dd)
		} else {
			md := MonitorDescriptor{}
			nOffset, _ := md.FromBytes(data[offset:])
			offset += nOffset
			dtd.Descriptors = append(dtd.Descriptors, &md)
		}

	}

	// create the raw byte slice
	dtd.Raw = data[:offset]

	return offset, nil
}

func (dtd *DisplayTimingDescriptor) ToBytes() []byte {
	return dtd.Raw
}

func (dtd *DisplayTimingDescriptor) Len() int {
	return len(dtd.Raw)
}
