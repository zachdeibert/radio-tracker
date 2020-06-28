package gimp

// #cgo pkg-config: gimp-2.0
// #include <stdlib.h>
// #include <libgimp/gimp.h>
//
// void getPixelData(gint32 id, guchar **data, gint *x, gint *y, gint *width, gint *height) {
//     static int init = 0;
//     if (!init) {
//         init = 1;
//         gegl_init(NULL, NULL);
//     }
//     gimp_drawable_offsets(id, x, y);
//     *width = gimp_drawable_width(id);
//     *height = gimp_drawable_height(id);
//     if (data) {
//         GeglBuffer *buffer = gimp_drawable_get_buffer(id);
//         GimpImageType type = gimp_drawable_type(id);
//         const Babl *format = babl_format("R'G'B'A u8");
//         *data = (guchar *) malloc(*width * *height * 4);
//         gegl_buffer_get(buffer, GEGL_RECTANGLE(0, 0, *width, *height), 1.0, format, *data, GEGL_AUTO_ROWSTRIDE, GEGL_ABYSS_NONE);
//         g_object_unref(buffer);
//     }
// }
import "C"
import (
	"image"
	"image/color"
	"unsafe"
)

// Layer represents on layer of the image
type Layer int32

// Bounds returns the size of the layer
func (l Layer) Bounds() image.Rectangle {
	var x, y, w, h C.gint
	C.getPixelData(C.gint32(l), nil, &x, &y, &w, &h)
	return image.Rectangle{
		Min: image.Point{
			X: int(x),
			Y: int(y),
		},
		Max: image.Point{
			X: int(x + w),
			Y: int(y + w),
		},
	}
}

// Image gets the Image this Layer is a part of
func (l Layer) Image() Image {
	return Image(C.gimp_item_get_image(C.gint32(l)))
}

// Name get sthe name of this Layer
func (l Layer) Name() string {
	return C.GoString(C.gimp_item_get_name(C.gint32(l)))
}

// Buffer reads the entire layer into memory
func (l Layer) Buffer() BufferedImage {
	size := l.Image().Bounds().Size()
	var ox, oy, w, h C.gint
	var data *C.guchar
	C.getPixelData(C.gint32(l), &data, &ox, &oy, &w, &h)
	defer C.free(unsafe.Pointer(data))
	colors := make([][]color.Color, int(w))
	dataArr := C.GoBytes(unsafe.Pointer(data), w*h*4)
	for x := range colors {
		colors[x] = make([]color.Color, int(h))
		for y := range colors[x] {
			i := 4 * (x + int(w)*y)
			colors[x][y] = &color.RGBA{
				R: dataArr[i],
				G: dataArr[i+1],
				B: dataArr[i+2],
				A: dataArr[i+3],
			}
		}
	}
	return BufferedImage{
		OffsetX: int(ox),
		OffsetY: int(oy),
		Width:   size.X,
		Height:  size.Y,
		Data:    colors,
	}
}
