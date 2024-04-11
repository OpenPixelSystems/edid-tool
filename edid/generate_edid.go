package edid

import (
	"fmt"
)

func generateChecksum(data []byte) byte {
	var sum byte
	for _, b := range data {
		sum += b
	}
	return byte(0x100 - int(sum))
}

func charToManIDByte(b byte) byte {
	fmt.Println(b - 0x40)
	return b - 0x40
}

func GenerateMonitorNameDescriptor(name string) [DISPLAY_DESCRIPTOR_SIZE]byte {
	var descriptor [DISPLAY_DESCRIPTOR_SIZE]byte
	descriptor[3] = DTD_TYPE_MONITOR_NAME
	copy(descriptor[5:], name)
	descriptor[DISPLAY_DESCRIPTOR_SIZE-1] = '\n'
	return descriptor
}

func (edid *EDID) ModifyDisplayDescriptor(descriptorId int, descriptor [DISPLAY_DESCRIPTOR_SIZE]byte) {
	fmt.Println(edid.displayDescriptor[descriptorId][:])
	fmt.Println(descriptor)
	copy(edid.displayDescriptor[descriptorId][:], descriptor[:])
}

func (edid *EDID) ModifyManufacturerId(manufacturerId [3]byte) {
	var manufacturerIdBytes [MANUFACTURER_ID_SIZE]byte
	manufacturerIdBytes[0] = (charToManIDByte(manufacturerId[0]) << 2) | (charToManIDByte(manufacturerId[1]) >> 3)
	manufacturerIdBytes[1] = ((charToManIDByte(manufacturerId[1])) << 5) | (charToManIDByte(manufacturerId[2]) & 0x1F)
	fmt.Printf("%x %x\n", manufacturerIdBytes[0], manufacturerIdBytes[1])
	copy(edid.manufacturerId[:], manufacturerIdBytes[:])
}

func (edid *EDID) ModifySerialNumber(serialNumber uint32) {
	edid.serialNumber[3] = byte(serialNumber >> 24)
	edid.serialNumber[2] = byte(serialNumber >> 16)
	edid.serialNumber[1] = byte(serialNumber >> 8)
	edid.serialNumber[0] = byte(serialNumber)
}

func GenerateEDID(reference *EDID) []byte {
	var data []byte
	data = append(data, reference.fixedHeader[:]...)
	data = append(data, reference.manufacturerId[:]...)
	data = append(data, reference.productCode[:]...)
	data = append(data, reference.serialNumber[:]...)
	data = append(data, reference.weekOfManufacture)
	data = append(data, reference.yearOfManufacture)
	data = append(data, reference.edidVersion)
	data = append(data, reference.edidRevision)
	data = append(data, reference.basicDisplayParameters[:]...)
	data = append(data, reference.chromaticityCoordinates[:]...)
	data = append(data, reference.establishedTimings[:]...)
	for i := 0; i < STANDARD_TIMINGS_COUNT; i++ {
		data = append(data, reference.standardTimings[i][:]...)
	}
	for i := 0; i < DISPLAY_DESCRIPTOR_COUNT; i++ {
		data = append(data, reference.displayDescriptor[i][:]...)
	}
	data = append(data, reference.extensionFlag)
	data = append(data, generateChecksum(data))
	return data
}
