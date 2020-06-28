package eeprom

import (
	"fmt"
	"io"
)

func writeHeader(writer io.Writer, sprites []sprite) {
	fmt.Fprintln(writer, "#ifndef RADIO_TRACKER_GENERATED_SPRITES_H")
	fmt.Fprintln(writer, "#define RADIO_TRACKER_GENERATED_SPRITES_H")
	fmt.Fprintln(writer, "")
	offset := 0
	for _, s := range sprites {
		fmt.Fprintf(writer, "#define RADIO_TRACKER_SPRITE_%s_OFFSET %d\n", s.DefineName, offset)
		fmt.Fprintln(writer, "")
		offset += s.encodedSize()
	}
	fmt.Fprintln(writer, "#endif")
}
