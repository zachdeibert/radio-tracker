package plugin

// #cgo pkg-config: gimp-2.0
// #include <stdlib.h>
// #include <libgimp/gimp.h>
//
// #define MAX_RETURN_VALS 4
//
// static GimpParam returnVals[MAX_RETURN_VALS];
//
// static void *allocArgs(int n) {
//     return malloc(sizeof(GimpParamDef) * n);
// }
//
// static void setArg(void *args, int n, GimpPDBArgType type, char *name, char *desc) {
//     GimpParamDef *def = ((GimpParamDef *) args) + n;
//     def->type = type;
//     def->name = name;
//     def->description = desc;
// }
//
// static void freeArgs(void *args, int n) {
//     GimpParamDef *defs = (GimpParamDef *) args;
//     for (int i = 0; i < n; ++i) {
//         free(defs[i].name);
//         free(defs[i].description);
//     }
//     free(defs);
// }
//
// static void install(char *name, char *blurb, char *help, char *author, char *copyright, char *date, char *menuLabel, char *imageTypes, GimpPDBProcType type, int numParams, int numReturnVals, void *params, void *returnVals) {
//     gimp_install_procedure(name, blurb, help, author, copyright, date, menuLabel, imageTypes, type, numParams, numReturnVals, (GimpParamDef *) params, (GimpParamDef *) returnVals);
//     free(name);
//     free(blurb);
//     free(help);
//     free(author);
//     free(copyright);
//     free(date);
//     free(menuLabel);
//     free(imageTypes);
// }
//
// static GimpParam getParam(GimpParam *params, int n) {
//     return params[n];
// }
//
// static int checkReturns(int n, int *out, GimpParam **params) {
//     if (n > MAX_RETURN_VALS) {
//         return MAX_RETURN_VALS;
//     }
//     *out = n;
//     *params = returnVals;
//     return 0;
// }
//
// static void setReturn(int n, GimpParam p) {
//     returnVals[n] = p;
// }
import "C"
import (
	"fmt"
	"unsafe"
)

var GimpPlugin C.GimpPDBProcType = C.GIMP_PLUGIN

// ProcedureRegistration interface
type ProcedureRegistration interface {
	installPlugIn(procedure Procedure)
}

// ProcedureArg is an argument for a Procedure
type ProcedureArg struct {
	Type        ParamType
	Name        string
	Description string
}

// Procedure represents a plugin procedure that can be installed into GIMP
type Procedure struct {
	Name         string
	Blurb        string
	Help         string
	Author       string
	Copyright    string
	Date         string
	MenuLabel    string
	ImageTypes   string
	Type         C.GimpPDBProcType
	Params       []ProcedureArg
	ReturnVals   []ProcedureArg
	Registration []ProcedureRegistration
	Handler      func(proc Procedure, params []Param) []Param
}

var procedures []Procedure = []Procedure{}

// AddProcedure adds a new procedure
func AddProcedure(procedure Procedure) {
	procedures = append(procedures, procedure)
}

func (p Procedure) install() {
	var params unsafe.Pointer = nil
	var returnVals unsafe.Pointer = nil
	if len(p.Params) > 0 {
		params = C.allocArgs(C.int(len(p.Params)))
		for i, param := range p.Params {
			C.setArg(params, C.int(i), C.GimpPDBArgType(param.Type), C.CString(param.Name), C.CString(param.Description))
		}
		defer C.freeArgs(params, C.int(len(p.Params)))
	}
	if len(p.ReturnVals) > 0 {
		returnVals = C.allocArgs(C.int(len(p.ReturnVals)))
		for i, ret := range p.ReturnVals {
			C.setArg(returnVals, C.int(i), C.GimpPDBArgType(ret.Type), C.CString(ret.Name), C.CString(ret.Description))
		}
		defer C.freeArgs(returnVals, C.int(len(p.ReturnVals)))
	}
	C.install(C.CString(p.Name), C.CString(p.Blurb), C.CString(p.Help), C.CString(p.Author), C.CString(p.Copyright), C.CString(p.Date), C.CString(p.MenuLabel), C.CString(p.ImageTypes), p.Type, C.int(len(p.Params)), C.int(len(p.ReturnVals)), params, returnVals)
	for _, reg := range p.Registration {
		reg.installPlugIn(p)
	}
}

func (p Procedure) run(nParams C.gint, param *C.GimpParam, nReturnVals *C.gint, returnVals **C.GimpParam) {
	params := make([]Param, nParams)
	for i := range params {
		params[i].fromC(C.getParam(param, C.int(i)))
	}
	ret := p.Handler(p, params)
	if max := int(C.checkReturns(C.int(len(ret)), nReturnVals, returnVals)); max != 0 {
		panic(fmt.Errorf("Procedure had too many return values;  Got %d values, but can only handle %d", len(ret), max))
	}
	for i, r := range ret {
		C.setReturn(C.int(i), r.toC())
	}
}
