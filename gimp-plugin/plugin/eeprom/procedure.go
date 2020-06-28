package eeprom

import "github.com/zachdeibert/radio-tracker/gimp-plugin/plugin"

var procedure = plugin.Procedure{
	Name:       "radio-tracker-eeprom",
	Blurb:      "Radio Tracker EEPROM Exporter",
	Help:       "Writes an EEPROM image for the radio tracker",
	Author:     "Zach Deibert",
	Copyright:  "Copyright (c) 2020 Zach Deibert",
	Date:       "July 2020",
	MenuLabel:  "Radio Tracker EEPROM",
	ImageTypes: "RGBA",
	Type:       plugin.GimpPlugin,
	Params: []plugin.ProcedureArg{
		{
			Type:        plugin.ParamTypeInt32,
			Name:        "run-mode",
			Description: "The run mode { RUN-INTERACTIVE (0), RUN-NONINTERACTIVE (1) }",
		},
		{
			Type:        plugin.ParamTypeImage,
			Name:        "image",
			Description: "Input image",
		},
		{
			Type:        plugin.ParamTypeDrawable,
			Name:        "drawable",
			Description: "Input drawable",
		},
		{
			Type:        plugin.ParamTypeString,
			Name:        "filename",
			Description: "The name of the file",
		},
		{
			Type:        plugin.ParamTypeString,
			Name:        "raw-filename",
			Description: "The name of the file",
		},
	},
	ReturnVals: []plugin.ProcedureArg{},
	Registration: []plugin.ProcedureRegistration{
		&plugin.FileMimeHandler{
			Type: "application/radio-tracker-eeprom",
		},
		&plugin.FileSaveHandler{
			Extensions: "bin",
		},
	},
	Handler: saveFile,
}

// Register adds this plugin to GIMP
func Register() {
	plugin.AddProcedure(procedure)
}
