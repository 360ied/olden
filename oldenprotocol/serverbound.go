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

// no-io protocol implementation
// all byte slices given to callback should not be retained
package oldenprotocol

import (
	"olden/oldenutils"
)

type SendCallback func([]byte)

func WithSendPlayerIdentification(username string, verificationKey []byte, callback SendCallback) {
	buf := oldenutils.GetBuffer()
	defer oldenutils.PutBuffer(buf)

	buf.WriteByte(0x00)                         // packet id
	buf.WriteByte(0x07)                         // protocol version 7
	buf.WriteString(classicString(username))    // username
	buf.Write(classicStrBytes(verificationKey)) // verification key
	buf.WriteByte(0x00)                         // unused. note: this is used by classic protocol extension

	callback(buf.Bytes())
}

func WithSendSetBlock(x, y, z uint16, block byte, callback SendCallback) {
	buf := oldenutils.GetBuffer()
	defer oldenutils.PutBuffer(buf)

	uint16Buf := make([]byte, 2)

	buf.WriteByte(0x05)            // packet id
	writeUint16(buf, uint16Buf, x) // x
	writeUint16(buf, uint16Buf, y) // y
	writeUint16(buf, uint16Buf, z) // z
	if block == 0x00 {             // if block is air, set mode to 0x00, because you can't place air
		buf.WriteByte(0x00) // mode
	} else { // block != 0x00
		buf.WriteByte(0x01) // mode
	}
	buf.WriteByte(block) // block

	callback(buf.Bytes())
}

func WithSendPositionAndOrientation(playerID byte, x, y, z uint16, yaw, pitch uint8, callback SendCallback) {
	buf := oldenutils.GetBuffer()
	defer oldenutils.PutBuffer(buf)

	uint16Buf := make([]byte, 2)

	buf.WriteByte(0x08)            // packet id
	buf.WriteByte(playerID)        // player id
	writeUint16(buf, uint16Buf, x) // x
	writeUint16(buf, uint16Buf, y) // y
	writeUint16(buf, uint16Buf, z) // z
	buf.WriteByte(yaw)             // yaw
	buf.WriteByte(pitch)           // pitch

	callback(buf.Bytes())
}

func WithSendMessage(message string, callback SendCallback) {
	buf := oldenutils.GetBuffer()
	defer oldenutils.PutBuffer(buf)

	buf.WriteByte(0x0d) // packet id
	buf.WriteByte(0xff) // unused, always 0xff
	buf.WriteString(classicString(message))

	callback(buf.Bytes())
}
