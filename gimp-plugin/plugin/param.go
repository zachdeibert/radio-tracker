package plugin

// #cgo pkg-config: gimp-2.0
// #include <stdint.h>
// #include <libgimp/gimp.h>
//
// static GimpParam toCint32(GimpPDBArgType type, int32_t intVal) {
//     GimpParam c;
//     c.type = type;
//     c.data.d_int32 = intVal;
//     return c;
// }
//
// static void fromC(GimpParam c, GimpPDBArgType *type, int32_t *intVal, char **outStr) {
//     *type = c.type;
//     *intVal = c.data.d_int32;
//     *outStr = c.data.d_string;
// }
import "C"
import "fmt"

// ParamType is the type of a Param
type ParamType int

const (
	// ParamTypeInt32 type
	ParamTypeInt32 = ParamType(C.GIMP_PDB_INT32)
	// ParamTypeString type
	ParamTypeString = ParamType(C.GIMP_PDB_STRING)
	// ParamTypeImage type
	ParamTypeImage = ParamType(C.GIMP_PDB_IMAGE)
	// ParamTypeDrawable type
	ParamTypeDrawable = ParamType(C.GIMP_PDB_DRAWABLE)
	// ParamTypeStatus type
	ParamTypeStatus = ParamType(C.GIMP_PDB_STATUS)
)

const (
	// StatusSuccess type
	StatusSuccess = int32(C.GIMP_PDB_SUCCESS)
)

// Param represents a parameter or return value
type Param struct {
	Type      ParamType
	IntVal    int32
	StringVal string
}

func (p Param) toC() C.GimpParam {
	switch p.Type {
	case ParamTypeInt32:
	case ParamTypeImage:
	case ParamTypeDrawable:
	case ParamTypeStatus:
		return C.toCint32(C.GimpPDBArgType(p.Type), C.int32_t(p.IntVal))
	}
	panic(fmt.Errorf("Unable to marshall parameter of type %d to the C structure", p.Type))
}

func (p *Param) fromC(c C.GimpParam) {
	var t C.GimpPDBArgType
	var intVal C.int32_t
	var strVal *C.char
	C.fromC(c, &t, &intVal, &strVal)
	p.Type = ParamType(t)
	p.IntVal = int32(intVal)
	if p.Type == ParamTypeString {
		p.StringVal = C.GoString(strVal)
	}
}
