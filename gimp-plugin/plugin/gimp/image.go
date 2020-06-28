package gimp

// #cgo pkg-config: gimp-2.0
// #include <libgimp/gimp.h>
//
// static gint32 getLayer(int *layers, int n) {
//     return layers[n];
// }
import "C"
import (
	"image"
)

// Image represents an image
type Image int32

// Bounds returns the size of the image
func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(C.gimp_image_width(C.gint32(i))),
			Y: int(C.gimp_image_height(C.gint32(i))),
		},
	}
}

// Layers get a list of all the layers in the image, from top to bottom
func (i Image) Layers() []Layer {
	var numLayers C.gint
	layerPtr := C.gimp_image_get_layers(C.gint32(i), &numLayers)
	layers := make([]Layer, numLayers)
	for i := range layers {
		layers[i] = Layer(C.getLayer(layerPtr, C.int(i)))
	}
	return layers
}
