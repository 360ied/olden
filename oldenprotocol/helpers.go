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
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"strings"

	"olden/oldenutils"
)

type readAction func(*bufio.Reader) error

func doRead(r *bufio.Reader, actions ...readAction) error {
	for _, v := range actions {
		if err := v(r); err != nil {
			return err
		}
	}
	return nil
}

func readByte(ptr *byte) readAction {
	return func(r *bufio.Reader) (err error) {
		*ptr, err = r.ReadByte()
		return
	}
}

func readInt8(ptr *int8) readAction {
	return func(r *bufio.Reader) (err error) {
		var b byte
		b, err = r.ReadByte()
		*ptr = int8(b)
		return
	}
}

func readUint16(ptr *uint16) readAction {
	return func(r *bufio.Reader) (err error) {
		buf := make([]byte, 2)
		_, err = io.ReadFull(r, buf)
		*ptr = binary.BigEndian.Uint16(buf)
		return
	}
}

func readClassicString(ptr *[64]byte) readAction {
	return func(r *bufio.Reader) (err error) {
		_, err = io.ReadFull(r, ptr[:])
		return
	}
}

func readClassicByteArray(ptr *[1024]byte) readAction {
	return func(r *bufio.Reader) (err error) {
		_, err = io.ReadFull(r, ptr[:])
		return
	}
}

// pads a string with spaces and trims it a length of 64
func classicString(s string) string {
	return (s + strings.Repeat("\x20", oldenutils.MaxInt(64-len(s), 0)))[:64]
}

func writeUint16(buf *bytes.Buffer, b []byte, v uint16) {
	binary.BigEndian.PutUint16(b, v)
	buf.Write(b)
}
