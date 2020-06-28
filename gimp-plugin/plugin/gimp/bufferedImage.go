package gimp

import (
	"image"
	"image/color"
)

// ColorNil represents a nil color
var ColorNil color.Color = &color.RGBA{
	R: 0,
	G: 0,
	B: 0,
	A: 0,
}

// BufferedImage represents an image that has fully been loaded into memory
type BufferedImage struct {
	OffsetX int
	OffsetY int
	Width   int
	Height  int
	Data    [][]color.Color
}

// ColorModel returns the Image's color model.
func (b BufferedImage) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
func (b BufferedImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: b.Width - 1,
			Y: b.Height - 1,
		},
	}
}

// At returns the color of the pixel at (x, y).
func (b BufferedImage) At(x, y int) color.Color {
	x -= b.OffsetX
	y -= b.OffsetY
	if x < 0 || x >= len(b.Data) || y < 0 || y >= len(b.Data[x]) {
		return ColorNil
	}
	return b.Data[x][y]
}
