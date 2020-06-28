package eeprom

import "image"

type block struct {
	Width   uint8
	Height  uint8
	XOffset uint8
	YOffset uint8
	Pixels  pixelBuf
}

func (b *block) decode(img image.Image, rect image.Rectangle) {
	b.Width = uint8(rect.Dx())
	b.Height = uint8(rect.Dy())
	b.XOffset = uint8(rect.Min.X)
	b.YOffset = uint8(rect.Min.Y)
	b.Pixels.decode(img, rect)
}

func (b block) encodedSize() int {
	return 4 + b.Pixels.encodedSize()
}

func (b block) encode() []byte {
	return append([]byte{
		byte(b.Width),
		byte(b.Height),
		byte(b.XOffset),
		byte(b.YOffset),
	}, b.Pixels.encode()...)
}
