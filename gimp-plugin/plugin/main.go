package plugin

// #cgo pkg-config: gimp-2.0
// #include <stdlib.h>
// #include <libgimp/gimp.h>
//
// extern void pluginInit(void);
// extern void pluginQuit(void);
// extern void pluginQuery(void);
// extern void pluginRun(gchar *name, gint nParams, GimpParam *param, gint *nReturnVals, GimpParam **returnVals);
// static void pluginRunWrapper(const gchar *name, gint nParams, const GimpParam *param, gint *nReturnVals, GimpParam **returnVals);
//
// GimpPlugInInfo PLUG_IN_INFO = {
//     .init_proc = pluginInit,
//     .quit_proc = pluginQuit,
//     .query_proc = pluginQuery,
//     .run_proc = pluginRunWrapper
// };
//
// static void pluginRunWrapper(const gchar *name, gint nParams, const GimpParam *param, gint *nReturnVals, GimpParam **returnVals) {
//     pluginRun((gchar *) name, nParams, (GimpParam *) param, nReturnVals, returnVals);
// }
//
// static void *allocArgv(int argc) {
//     return malloc(sizeof(char *) * argc);
// }
//
// static void argvSet(void *argv, int i, char *arg) {
//     ((char **) argv)[i] = arg;
// }
//
// static void freeArgv(int argc, void *argv) {
//     for (int i = 0; i < argc; ++i) {
//         free(((char **) argv)[i]);
//     }
//     free(argv);
// }
//
// static int callMain(int argc, void *argv) {
//     return gimp_main(&PLUG_IN_INFO, argc, (char **) argv);
// }
import "C"
import "os"

// Main runs the plugin
func Main() {
	argv := C.allocArgv(C.int(len(os.Args)))
	for i, arg := range os.Args {
		C.argvSet(argv, C.int(i), C.CString(arg))
	}
	defer C.freeArgv(C.int(len(os.Args)), argv)
	os.Exit(int(C.callMain(C.int(len(os.Args)), argv)))
}
