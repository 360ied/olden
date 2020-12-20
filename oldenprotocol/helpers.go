package oldenprotocol

import (
	"bytes"
	"encoding/binary"
	"strings"

	"olden/oldenutils"
)

// pads a string with spaces and trims it a length of 64
func classicString(s string) string {
	return (s + strings.Repeat("\x20", oldenutils.MaxInt(64-len(s), 0)))[:64]
}

func classicStrBytes(s []byte) []byte {
	return append(s, bytes.Repeat([]byte{0x20}, oldenutils.MaxInt(64-len(s), 0))...)[:64]
}

func writeUint16(buf *bytes.Buffer, b []byte, v uint16) {
	binary.BigEndian.PutUint16(b, v)
	buf.Write(b)
}
