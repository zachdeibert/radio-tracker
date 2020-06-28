package eeprom

import (
	"fmt"
	"os"

	"github.com/zachdeibert/radio-tracker/gimp-plugin/plugin"
	"github.com/zachdeibert/radio-tracker/gimp-plugin/plugin/gimp"
)

func writeBinaryFile(filename string, sprites []sprite) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, s := range sprites {
		f.Write(s.encode())
	}
}

func writeHeaderFile(filename string, sprites []sprite) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writeHeader(f, sprites)
}

func saveFile(proc plugin.Procedure, params []plugin.Param) []plugin.Param {
	im := gimp.Image(params[1].IntVal)
	layers := im.Layers()
	sprites := make([]sprite, len(layers))
	for i, l := range layers {
		sprites[i].decode(l.Buffer(), l.Name())
	}
	writeBinaryFile(params[3].StringVal, sprites)
	writeHeaderFile(fmt.Sprintf("%s.h", params[3].StringVal), sprites)
	return []plugin.Param{
		{
			Type:   plugin.ParamTypeStatus,
			IntVal: plugin.StatusSuccess,
		},
	}
}
