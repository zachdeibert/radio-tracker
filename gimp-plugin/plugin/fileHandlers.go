package plugin

// #cgo pkg-config: gimp-2.0
// #include <stdlib.h>
// #include <libgimp/gimp.h>
import "C"
import "unsafe"

// FileMimeHandler registers a MIME type to be handled by a procedure
type FileMimeHandler struct {
	Type string
}

func (h FileMimeHandler) installPlugIn(procedure Procedure) {
	procName := C.CString(procedure.Name)
	defer C.free(unsafe.Pointer(procName))
	mimeName := C.CString(h.Type)
	defer C.free(unsafe.Pointer(mimeName))
	C.gimp_register_file_handler_mime(procName, mimeName)
}

// FileSaveHandler registers a file extension that can be saved by a procedure
type FileSaveHandler struct {
	Extensions string
	Prefixes   string
}

func (h FileSaveHandler) installPlugIn(procedure Procedure) {
	procName := C.CString(procedure.Name)
	defer C.free(unsafe.Pointer(procName))
	extName := C.CString(h.Extensions)
	defer C.free(unsafe.Pointer(extName))
	preName := C.CString(h.Prefixes)
	defer C.free(unsafe.Pointer(preName))
	C.gimp_register_save_handler(procName, extName, preName)
}
