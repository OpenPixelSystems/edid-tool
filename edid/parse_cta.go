package edid

import (
	"encoding/binary"
	"fmt"
)

func parseTiledDisplayTopology(dd []byte) int {
	fmt.Println("Parsing tiled display topology")
	revision := dd[1]
	fmt.Printf("\tRevision: 0x%02x\n", revision)
	numberOfPayloadBytes := dd[2]
	fmt.Printf("\tNumber of payload bytes: %d\n", numberOfPayloadBytes)

	// display capabilities
	caps := dd[3]
	oneTileBehavior := caps & TILE_ONE_TILE_BEHAVIOR
	nTileBehavior := caps & TILE_N_TILE_BEHAVIOR >> 3
	bezelDescription := caps & TILE_BEZEL_DESCRIPTOR >> 6
	physicalEnclosure := caps & TILE_PHYSICAL_ENCLOSURE >> 7

	switch oneTileBehavior {
	case 0x00:
		fmt.Println("\tOne tile behavior: None of the following")
	case 0x01:
		fmt.Println("\tOne tile behavior: Display on tile", 0)
	case 0x02:
		fmt.Println("\tOne tile behavior: Scale to fit the display")
	case 0x03:
		fmt.Println("\tOne tile behavior: Clone to other tiles")
	}

	switch nTileBehavior {
	case 0x00:
		fmt.Println("\tN tile behavior: None of the following")
	case 0x01:
		fmt.Println("\tN tile behavior: Display on tile", 0)
	}

	if bezelDescription == 0x00 {
		fmt.Println("\tBezel description: None bezel description")
	} else {
		fmt.Println("\tBezel description: ", bezelDescription)
		// bytes 11 - 15
	}

	if physicalEnclosure == 0x00 {
		fmt.Println("\tPhysical bezel: Multiple physical enclosures")
	} else {
		fmt.Println("\tPhysical bezel: Single physical enclosure")
	}

	// Tiled display topology
	nrTilesLSB := dd[4]
	locationTilesLSB := dd[5]
	tilesMSB := dd[6]

	print("\tNumber of tiles: ", nrTilesLSB, " ", tilesMSB, "\n")
	nrVerticalTiles := ((tilesMSB & 0x03) | (nrTilesLSB & 0x0f)) + 1
	nrHorizontalTiles := ((tilesMSB & 0x0c >> 2) | (nrTilesLSB & 0xf0 >> 4)) + 1
	fmt.Printf("\tNumber of horizontal tiles: %d\n", nrHorizontalTiles)
	fmt.Printf("\tNumber of vertical tiles: %d\n", nrVerticalTiles)

	// tiled display location
	verticalTileLocation := ((tilesMSB & 0x30) << 4) | (locationTilesLSB & 0x0f)
	horizontalTileLocation := ((tilesMSB & 0xc0) >> 2) | (locationTilesLSB & 0xf0 >> 4)
	fmt.Printf("\tVertical tile location: %d\n", verticalTileLocation)
	fmt.Printf("\tHorizontal tile location: %d\n", horizontalTileLocation)

	// tile size
	horizontalSizeLSB := dd[7]
	horizontalSizeMSB := dd[8]
	horizontalSize := ((int(horizontalSizeMSB) << 8) | int(horizontalSizeLSB)) + 1

	verticalSizeLSB := dd[9]
	verticalSizeMSB := dd[10]
	verticalSize := ((int(verticalSizeMSB) << 8) | int(verticalSizeLSB)) + 1

	fmt.Printf("\tTile size: %d x %d\n", horizontalSize, verticalSize)

	// Tile pixel multiplier and bezel information
	// 0xB -> 0x0F

	//Tile diplay topology id
	// 0x10 -> 0x18
	var tiledDisplayVendorID [3]byte
	var tiledDisplayProductID [2]byte
	var tiledDisplaySerialNumber [4]byte
	copy(tiledDisplayVendorID[:], dd[16:19])
	copy(tiledDisplayProductID[:], dd[19:21])
	copy(tiledDisplaySerialNumber[:], dd[21:25])
	fmt.Printf("\tTiled display vendor ID: %s\n", tiledDisplayVendorID)
	fmt.Printf("\tTiled display product ID: %d\n", binary.LittleEndian.Uint16(tiledDisplayProductID[:]))
	fmt.Printf("\tTiled display serial number: %d\n", binary.LittleEndian.Uint32(tiledDisplaySerialNumber[:]))

	return 3 + int(numberOfPayloadBytes)
}

func parseVTBType1(vtb []byte) int {
	fmt.Println("Parsing VTB type 1")
	revision := vtb[1]
	fmt.Printf("\tRevision: 0x%02x\n", revision)
	numberOfPayloadBytes := vtb[2]
	fmt.Printf("\tNumber of payload bytes: %d\n", numberOfPayloadBytes)
	fmt.Printf("\tNumber of video timing blocks: %d\n", numberOfPayloadBytes/CTA_VTB_TYPE_1_DESCRIPTOR_SIZE)

	offset := 3
	for i := 0; i < int(numberOfPayloadBytes/CTA_VTB_TYPE_1_DESCRIPTOR_SIZE); i++ {
		var vtd VTBDescriptor
		// Video timing block descriptor
		fmt.Printf("Video timing block %d\n", i+1)
		vtd.pixelClockLBits = vtb[offset+0] // Pixel clock low bits
		vtd.pixelClockMBits = vtb[offset+1] // Pixel clock middle bits
		vtd.pixelClockHBits = vtb[offset+2] // Pixel clock high bits
		pixelClock := float64((int(vtd.pixelClockHBits)<<16)|(int(vtd.pixelClockMBits)<<8)|int(vtd.pixelClockLBits)) / 100.0
		fmt.Printf("\tPixel clock: %fMHz\n", pixelClock)

		vtd.timingOptions = vtb[offset+3]
		fmt.Printf("\tTiming options: 0x%02x\n", vtd.timingOptions)

		vtd.hActiveLSB = vtb[offset+4] // Horizontal active low bits
		vtd.hActiveMSB = vtb[offset+5] // Horizontal active high bits
		hActive := (int(vtd.hActiveMSB)<<8 | int(vtd.hActiveLSB)) + 1

		vtd.hBlankingLSB = vtb[offset+6] // Horizontal blanking low bits
		vtd.hBlankingMSB = vtb[offset+7] // Horizontal blanking high bits
		hBlanking := (int(vtd.hBlankingMSB)<<8 | int(vtd.hBlankingLSB)) + 1

		vtd.hFrontPorchLSB = vtb[offset+8]        // Horizontal front porch low bits
		vtd.hFrontPorchMSB = vtb[offset+9] & 0x7F // Horizontal front porch high bits
		hFrontPorch := (int(vtd.hFrontPorchMSB)<<8 | int(vtd.hFrontPorchLSB)) + 1

		hSyncPol := vtb[offset+9] & 0x80 >> 7 // Horizontal sync polarity
		hSyncPolStr := "P"
		if hSyncPol == 0 {
			hSyncPolStr = "N"
		}
		vtd.hSyncWidthLSB = vtb[offset+10] // Horizontal sync width low bits
		vtd.hSyncWidthMSB = vtb[offset+11] // Horizontal sync width high bits
		hSyncWidth := (int(vtd.hSyncWidthMSB)<<8 | int(vtd.hSyncWidthLSB)) + 1

		hBackPorch := hBlanking - hFrontPorch - hSyncWidth

		fmt.Printf("\tha: %d, hbl: %d, hfp: %d, hbp; %d, hsync: %d, Hpol %s\n", hActive, hBlanking, hFrontPorch, hBackPorch, hSyncWidth, hSyncPolStr)

		vtd.vActiveLSB = vtb[offset+12] // Vertical active low bits
		vtd.vActiveMSB = vtb[offset+13] // Vertical active high bits
		vActive := (int(vtd.vActiveMSB)<<8 | int(vtd.vActiveLSB)) + 1

		vtd.vBlankingLSB = vtb[offset+14] // Vertical blanking low bits
		vtd.vBlankingMSB = vtb[offset+15] // Vertical blanking high bits
		vBlanking := (int(vtd.vBlankingMSB)<<8 | int(vtd.vBlankingLSB)) + 1

		vtd.vFrontPorchLSB = vtb[offset+16]        // Vertical front porch low bits
		vtd.vFrontPorchMSB = vtb[offset+17] & 0x7F // Vertical front porch high bits
		vFrontPorch := (int(vtd.vFrontPorchMSB)<<8 | int(vtd.vFrontPorchLSB)) + 1

		vSyncPol := vtb[offset+17] & 0x80 >> 7 // Vertical sync polarity
		vSyncPolStr := "P"
		if vSyncPol == 0 {
			vSyncPolStr = "N"
		}
		vtd.vSyncWidthLSB = vtb[offset+18] // Vertical sync width low bits
		vtd.vSyncWidthMSB = vtb[offset+19] // Vertical sync width high bits
		vSyncWidth := (int(vtd.vSyncWidthMSB)<<8 | int(vtd.vSyncWidthLSB)) + 1

		vBackPorch := vBlanking - vFrontPorch - vSyncWidth

		htotal := hActive + hBlanking
		vtotal := vActive + vBlanking
		refRate := float64(pixelClock*1000*1000) / float64(htotal*vtotal)
		fmt.Printf("\tva: %d, vbl: %d, vfp: %d, vbp; %d, vsync: %d, Vpol %s\n", vActive, vBlanking, vFrontPorch, vBackPorch, vSyncWidth, vSyncPolStr)
		fmt.Printf("\tTotal: %d x %d, refresh rate: %fHz\n", htotal, vtotal, refRate)

		offset += CTA_VTB_TYPE_1_DESCRIPTOR_SIZE
	}

	return 3 + int(numberOfPayloadBytes)
}

func (edid *EDID) ParseCTA() error {
	fmt.Println("Parsing CTA extension")

	blockType := edid.ctaData[0] >> 4
	totalLength := edid.ctaData[0] & 0x0f
	fmt.Printf("Block type: 0x%02x\n", blockType)
	fmt.Printf("Total length: %d\n", totalLength)
	switch blockType {
	case CTA_EXT_TAG_AUDIO_DATA_BLOCK:
		fmt.Println("Audio data block")
	case CTA_EXT_TAG_VIDEO_DATA_BLOCK:
		fmt.Println("Video data block")
	case CTA_EXT_TAG_VENDOR_SPECIFIC_DATA_BLOCK:
		fmt.Println("Vendor-specific data block")
	case CTA_EXT_TAG_SPEAKER_ALLOCATION_DATA_BLOCK:
		fmt.Println("Speaker allocation data block")
	case CTA_EXT_TAG_VESA_DTC_DATA_BLOCK:
		fmt.Println("VESA DTC data block")
	case CTA_EXT_TAG_VIDEO_FORMAT_DATA_BLOCK:
		fmt.Println("Video format data block")
	case CTA_EXT_TAG_USE_EXTENDED_TAG:
		fmt.Println("Use extended tag")
	}

	dpidRev := edid.ctaData[1]
	fmt.Printf("DisplayID revision: 0x%02x\n", dpidRev)
	dpidVariableLength := edid.ctaData[2]
	fmt.Printf("DisplayID variable length: 0x%02x\n", dpidVariableLength)
	primaryUseCase := edid.ctaData[3]
	fmt.Printf("Primary use case: 0x%02x\n", primaryUseCase)
	extCount := edid.ctaData[4]
	fmt.Printf("Extension count: 0x%02x\n", extCount)

	offset := 5
	done := false
	for !done {
		blockTypeTag := edid.ctaData[offset]
		fmt.Printf("Block type tag @ 0x%02x: 0x%02x\n", offset, blockTypeTag)
		switch blockTypeTag {
		case CTA_BLOCK_TILED_DISPLAY, CTA_BLOCK_TILED_DISPLAY_LEGACY:
			displayData := edid.ctaData[offset : offset+CTA_BLOCK_TILED_SIZE]
			payloadSize := parseTiledDisplayTopology(displayData)
			offset += payloadSize
		case CTA_BLOCK_VTB_TYPE_1:
			fmt.Println("VTB type 1")
			vtdData := edid.ctaData[offset:]
			payloadSize := parseVTBType1(vtdData)
			offset += payloadSize
		case 0x00:
			fmt.Println("End of CTA extension")
			done = true
		default:
			fmt.Println("Unknown block type")
		}
	}
	return nil
}
