package main

import (
	"github.com/zachdeibert/radio-tracker/gimp-plugin/plugin"
	"github.com/zachdeibert/radio-tracker/gimp-plugin/plugin/eeprom"
)

func main() {
	eeprom.Register()
	plugin.Main()
}
