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
)

type (
	ServerIdentification struct {
		ProtocolVersion byte
		ServerName      [64]byte
		ServerMOTD      [64]byte
		UserType        byte
	}
	Ping            struct{}
	LevelInitialize struct{}
	LevelDataChunk  struct {
		Length          uint16
		Data            [1024]byte
		PercentComplete byte
	}
	LevelFinalize struct {
		XSize uint16
		YSize uint16
		ZSize uint16
	}
	SetBlock struct {
		X     uint16
		Y     uint16
		Z     uint16
		Block byte
	}
	SpawnPlayer struct {
		PlayerID   byte
		PlayerName [64]byte
		X          uint16
		Y          uint16
		Z          uint16
		Yaw        byte
		Pitch      byte
	}
	PositionAndOrientation struct {
		PlayerID byte
		X        uint16
		Y        uint16
		Z        uint16
		Yaw      byte
		Pitch    byte
	}
	PositionAndOrientationUpdate struct {
		PlayerID byte
		XDelta   int8
		YDelta   int8
		ZDelta   int8
		Yaw      byte
		Pitch    byte
	}
	PositionUpdate struct {
		PlayerID byte
		XDelta   int8
		YDelta   int8
		ZDelta   int8
	}
	OrientationUpdate struct {
		PlayerID byte
		Yaw      byte
		Pitch    byte
	}
	DespawnPlayer struct {
		PlayerID byte
	}
	Message struct {
		PlayerID byte
		Message  [64]byte
	}
	DisconnectPlayer struct {
		Reason [64]byte
	}
	UpdateUserType struct {
		Type byte
	}
)

const (
	ServerIdentificationID         byte = 0x00
	PingID                         byte = 0x01
	LevelInitializeID              byte = 0x02
	LevelDataChunkID               byte = 0x03
	LevelFinalizeID                byte = 0x04
	SetBlockID                     byte = 0x06
	SpawnPlayerID                  byte = 0x07
	PositionAndOrientationID       byte = 0x08
	PositionAndOrientationUpdateID byte = 0x09
	PositionUpdateID               byte = 0x0a
	OrientationUpdateID            byte = 0x0b
	DespawnPlayerID                byte = 0x0c
	MessageID                      byte = 0x0d
	DisconnectPlayerID             byte = 0x0e
	UpdateUserTypeID               byte = 0x0f
)

// ReadIncoming always returns a non-nil error
func ReadIncoming(r *bufio.Reader, callback func(packet interface{}), unknownPacketIDCallback func(r *bufio.Reader, packetID byte) error) error {
	for {
		packetID, packetIDErr := r.ReadByte()
		if packetIDErr != nil {
			return packetIDErr
		}
		switch packetID {
		case ServerIdentificationID:
			p := ServerIdentification{}
			if err := doRead(r,
				readByte(&p.ProtocolVersion),
				readClassicString(&p.ServerName),
				readClassicString(&p.ServerMOTD),
				readByte(&p.UserType),
			); err != nil {
				return err
			}
			callback(p)
		case PingID:
			callback(Ping{})
		case LevelInitializeID:
			callback(LevelInitialize{})
		case LevelDataChunkID:
			p := LevelDataChunk{}
			if err := doRead(r,
				readUint16(&p.Length),
				readClassicByteArray(&p.Data),
				readByte(&p.PercentComplete),
			); err != nil {
				return err
			}
			callback(p)
		case LevelFinalizeID:
			p := LevelFinalize{}
			if err := doRead(r,
				readUint16(&p.XSize),
				readUint16(&p.YSize),
				readUint16(&p.ZSize),
			); err != nil {
				return err
			}
			callback(p)
		case SetBlockID:
			p := SetBlock{}
			if err := doRead(r,
				readUint16(&p.X),
				readUint16(&p.Y),
				readUint16(&p.Z),
				readByte(&p.Block),
			); err != nil {
				return err
			}
			callback(p)
		case SpawnPlayerID:
			p := SpawnPlayer{}
			if err := doRead(r,
				readByte(&p.PlayerID),
				readClassicString(&p.PlayerName),
				readUint16(&p.X),
				readUint16(&p.Y),
				readUint16(&p.Z),
				readByte(&p.Yaw),
				readByte(&p.Pitch),
			); err != nil {
				return err
			}
			callback(p)
		case PositionAndOrientationID:
			p := PositionAndOrientation{}
			if err := doRead(r,
				readByte(&p.PlayerID),
				readUint16(&p.X),
				readUint16(&p.Y),
				readUint16(&p.Z),
				readByte(&p.Yaw),
				readByte(&p.Pitch),
			); err != nil {
				return err
			}
			callback(p)
		case PositionAndOrientationUpdateID:
			p := PositionAndOrientationUpdate{}
			if err := doRead(r,
				readByte(&p.PlayerID),
				readInt8(&p.XDelta),
				readInt8(&p.YDelta),
				readInt8(&p.ZDelta),
				readByte(&p.Yaw),
				readByte(&p.Pitch),
			); err != nil {
				return err
			}
			callback(p)
		case PositionUpdateID:
			p := PositionUpdate{}
			if err := doRead(r,
				readByte(&p.PlayerID),
				readInt8(&p.XDelta),
				readInt8(&p.YDelta),
				readInt8(&p.ZDelta),
			); err != nil {
				return err
			}
			callback(p)
		case OrientationUpdateID:
			p := OrientationUpdate{}
			if err := doRead(r,
				readByte(&p.PlayerID),
				readByte(&p.Yaw),
				readByte(&p.Pitch),
			); err != nil {
				return err
			}
			callback(p)
		case DespawnPlayerID:
			p := DespawnPlayer{}
			if err := doRead(r,
				readByte(&p.PlayerID),
			); err != nil {
				return err
			}
			callback(p)
		case MessageID:
			p := Message{}
			if err := doRead(r,
				readByte(&p.PlayerID),
				readClassicString(&p.Message),
			); err != nil {
				return err
			}
			callback(p)
		case DisconnectPlayerID:
			p := DisconnectPlayer{}
			if err := doRead(r,
				readClassicString(&p.Reason),
			); err != nil {
				return err
			}
			callback(p)
		case UpdateUserTypeID:
			p := UpdateUserType{}
			if err := doRead(r,
				readByte(&p.Type),
			); err != nil {
				return err
			}
			callback(p)
		default:
			if err := unknownPacketIDCallback(r, packetID); err != nil {
				return err
			}
		}
	}
}
