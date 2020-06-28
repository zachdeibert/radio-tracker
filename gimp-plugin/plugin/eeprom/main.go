package eeprom

import (
	"image/png"
	"os"

	"github.com/zachdeibert/radio-tracker/gimp-plugin/plugin"
	"github.com/zachdeibert/radio-tracker/gimp-plugin/plugin/gimp"
)

func saveFile(proc plugin.Procedure, params []plugin.Param) []plugin.Param {
	plugin.ShowMessage("Filename: %s", params[3].StringVal)
	im := gimp.Image(params[1].IntVal)
	for i, l := range im.Layers() {
		if i == 0 {
			img := l.Buffer()
			f, err := os.Create(params[3].StringVal)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			if err = png.Encode(f, img); err != nil {
				panic(err)
			}
		}
	}
	return []plugin.Param{
		{
			Type:   plugin.ParamTypeStatus,
			IntVal: plugin.StatusSuccess,
		},
	}
}
