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
package oldenutils

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

// Type-safe buffer get
func GetBuffer() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

// Type-safe buffer put
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
