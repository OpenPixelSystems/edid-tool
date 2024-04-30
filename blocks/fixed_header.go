package blocks

import (
	"encoding/binary"
	"fmt"
)

func (fh *FixedHeader) Print() {
	fmt.Println("Fixed Header")
	fmt.Println("\tHeader: ", fh.Header)
	fmt.Println("\tManufacturer", string(fh.ManufacturerId[:]))
	fmt.Printf("\tProduct Code: %x\n", fh.ProductCode)
	fmt.Println("\tSerial Number: ", fh.SerialNumber)
	fmt.Println("\tWeek of Manufacture: ", int(fh.WeekOfManufacture))
	fmt.Println("\tYear of Manufacture: ", int(fh.YearOfManufacture)+1990)
	fmt.Printf("\tEDID Version: %d.%d\n", fh.EdidVersion, fh.EdidRevision)
}

func (fh *FixedHeader) FromBytes(data []byte) (int, error) {
	offset := 0
	// parse the fixed header
	copy(fh.Header[:], data[offset:offset+FIXED_HEADER_SIZE])
	offset += FIXED_HEADER_SIZE

	// parse the manufacturer id
	fh.ManufacturerId[0] = ManIdByteToChar((data[offset] >> 2) & 0x1F)
	fh.ManufacturerId[1] = ManIdByteToChar(((data[offset] & 0x3) << 3) | ((data[offset+1] & 0xE0) >> 5))
	fh.ManufacturerId[2] = ManIdByteToChar(data[offset+1] & 0x1F)
	offset += MANUFACTURER_ID_SIZE

	// parse the product code
	fh.ProductCode = binary.LittleEndian.Uint16(data[offset : offset+PRODUCT_CODE_SIZE])
	offset += PRODUCT_CODE_SIZE

	// parse the serial number
	fh.SerialNumber = binary.LittleEndian.Uint32(data[offset : offset+SERIAL_NUMBER_SIZE])
	offset += SERIAL_NUMBER_SIZE

	// parse the week of manufacture
	fh.WeekOfManufacture = data[offset]
	offset += WEEK_OF_MANUFACTURE_SIZE

	// parse the year of manufacture
	fh.YearOfManufacture = data[offset]
	offset += YEAR_OF_MANUFACTURE_SIZE

	// parse the edid version
	fh.EdidVersion = data[offset]
	offset += EDID_VERSION_SIZE

	// parse the edid revision
	fh.EdidRevision = data[offset]
	offset += EDID_REVISION_SIZE

	// create the raw byte slice
	fh.Raw = data[:offset]

	return offset, nil
}

func (fh *FixedHeader) ToBytes() []byte {
	return fh.Raw
}

func (fh *FixedHeader) Len() int {
	return len(fh.Raw)
}
