package eeprom

import (
	"image"
	"image/color"
)

type pixel uint8

const (
	pixelOff   pixel = 0
	pixelOn    pixel = 1
	pixelClear pixel = 2
	pixelEnd   pixel = 3
)

type pixelBuf struct {
	Pixels []pixel
}

func getPixel(c color.Color) pixel {
	r, g, b, a := c.RGBA()
	if a < 0x7FFF {
		return pixelClear
	} else if b > r+g {
		return pixelOn
	}
	return pixelOff
}

func (p *pixelBuf) decode(img image.Image, rect image.Rectangle) {
	size := rect.Size()
	p.Pixels = make([]pixel, size.X*size.Y)
	i := 0
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			p.Pixels[i] = getPixel(img.At(x, y))
			i++
		}
	}
}

func (p pixelBuf) encodedSize() int {
	if len(p.Pixels)%4 == 0 {
		return len(p.Pixels)/4 + 1
	}
	return (len(p.Pixels) + 3) / 4
}

func (p pixelBuf) encode() []byte {
	res := make([]byte, p.encodedSize())
	res[len(res)-1] = byte((pixelEnd << 6) | (pixelEnd << 4) | (pixelEnd << 2) | pixelEnd)
	for i, p := range p.Pixels {
		b := 2 * (i % 4)
		res[i/4] = ((res[i/4] & ^byte(3<<b)) | byte(p<<b))
	}
	return res
}
