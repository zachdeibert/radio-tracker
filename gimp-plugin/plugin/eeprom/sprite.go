package eeprom

import (
	"image"
	"strings"
)

type sprite struct {
	Blocks     []block
	Name       string
	DefineName string
	CamelName  string
	PascalName string
}

func (s *sprite) decode(img image.Image, name string) {
	s.decodeNames(name)
	bounds := getMinimumBounds(img)
	s.Blocks = []block{{}}
	s.Blocks[0].decode(img, bounds)
}

func getMinimumBounds(img image.Image) image.Rectangle {
	bounds := img.Bounds()
	for shrunk := true; shrunk; {
		shrunk = false
		min := true
		max := true
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if min && getPixel(img.At(bounds.Min.X, y)) != pixelClear {
				min = false
				if !max {
					break
				}
			}
			if max && getPixel(img.At(bounds.Max.X, y)) != pixelClear {
				max = false
				if !min {
					break
				}
			}
		}
		if min {
			bounds.Min.X++
			shrunk = true
		}
		if max {
			bounds.Max.X--
			shrunk = true
		}
		min = true
		max = true
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if min && getPixel(img.At(x, bounds.Min.Y)) != pixelClear {
				min = false
				if !max {
					break
				}
			}
			if max && getPixel(img.At(x, bounds.Max.Y)) != pixelClear {
				max = false
				if !min {
					break
				}
			}
		}
		if min {
			bounds.Min.Y++
			shrunk = true
		}
		if max {
			bounds.Max.Y--
			shrunk = true
		}
	}
	return bounds
}

func (s *sprite) decodeNames(name string) {
	s.Name = name
	var defineName strings.Builder
	var camelName strings.Builder
	var pascalName strings.Builder
	newWord := false
	for i, c := range name {
		if c >= 'a' && c <= 'z' {
			defineName.WriteRune(c + 'A' - 'a')
			if newWord {
				camelName.WriteRune(c + 'A' - 'a')
				pascalName.WriteRune(c + 'A' - 'a')
				newWord = false
			} else {
				camelName.WriteRune(c)
				if i == 0 {
					pascalName.WriteRune(c + 'A' - 'a')
				} else {
					pascalName.WriteRune(c)
				}
			}
		} else if (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			defineName.WriteRune(c)
			camelName.WriteRune(c)
			pascalName.WriteRune(c)
			newWord = false
		} else if !newWord {
			defineName.WriteRune('_')
			newWord = true
		}
	}
	s.DefineName = defineName.String()
	s.CamelName = camelName.String()
	s.PascalName = pascalName.String()
}

func (s sprite) encodedSize() int {
	size := 1
	for _, b := range s.Blocks {
		size += b.encodedSize()
	}
	return size
}

func (s sprite) encode() []byte {
	enc := []byte{byte(len(s.Blocks))}
	for _, b := range s.Blocks {
		enc = append(enc, b.encode()...)
	}
	return enc
}
