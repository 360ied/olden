//    Olden, a Minecraft classic client
//    Copyright (C) 2020  360ied
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <https://www.gnu.org/licenses/>.
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
