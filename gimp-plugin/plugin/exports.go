package plugin

// #cgo pkg-config: gimp-2.0
// #include <libgimp/gimp.h>
import "C"

//export pluginInit
func pluginInit() {

}

//export pluginQuit
func pluginQuit() {

}

//export pluginQuery
func pluginQuery() {
	for _, p := range procedures {
		p.install()
	}
}

//export pluginRun
func pluginRun(name *C.gchar, nParams C.gint, param *C.GimpParam, nReturnVals *C.gint, returnVals **C.GimpParam) {
	nameStr := C.GoString(name)
	for _, p := range procedures {
		if p.Name == nameStr {
			p.run(nParams, param, nReturnVals, returnVals)
		}
	}
}
