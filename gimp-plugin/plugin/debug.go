package plugin

// #cgo pkg-config: gimp-2.0
// #include <libgimp/gimp.h>
//
// static void showMessage(char *msg) {
//     g_message(msg);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// ShowMessage shows a debugging message
func ShowMessage(format string, params ...interface{}) {
	str := C.CString(fmt.Sprintf(format, params...))
	defer C.free(unsafe.Pointer(str))
	C.showMessage(str)
}
