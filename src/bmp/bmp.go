package bmp

import (
	"encoding/binary"
	"errors"
	"os"
	"reflect"
)

type BMPHeader struct {
	HeaderField    uint16
	BPMSize        uint32
	Reserved       [2]uint16
	PixelArrOffset uint32
	HeaderSize     uint32
	Width          uint16
	Height         uint16
	PlaneNumber    uint16
	BitsPerPixel   uint16
}

func ToBMP(color_array []RGB, width, height uint16, path string) error {

	if len(color_array) != int(width)*int(height) {
		return errors.New("the length of the color array doesn't match the informed dimensions")
	}

	var SIZEOF_COLOR = uint16(reflect.TypeOf(RGB{}).Size())
	var ALIGNMENT = (SIZEOF_COLOR*width + 3) & ^uint16(3)

	bmp := BMPHeader{
		HeaderField:    'B' | 'M'<<8,
		BPMSize:        uint32(0x1A + ALIGNMENT*height),
		Reserved:       [2]uint16{0, 0},
		PixelArrOffset: 0x1A,
		HeaderSize:     0x0C,
		Width:          width,
		Height:         height,
		PlaneNumber:    1,
		BitsPerPixel:   8 * SIZEOF_COLOR,
	}

	buf := make([]byte, bmp.BPMSize)

	offset, err := binary.Encode(buf, binary.LittleEndian, bmp)
	if err != nil {
		return err
	}

	for y := height - 1; true; y-- {
		row_offset := offset
		for x := range width {
			buf[row_offset+2] = color_array[y*width+x].R
			buf[row_offset+1] = color_array[y*width+x].G
			buf[row_offset+0] = color_array[y*width+x].B
			row_offset += 3
		}

		offset += int(ALIGNMENT)

		if y == 0 {
			break
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		return err
	}

	return nil
}