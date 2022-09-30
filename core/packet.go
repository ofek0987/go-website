package core

import (
	"encoding/binary"
)

type Packet struct {
	PacketLen   uint32
	PanndingLen uint8
	Payload     []byte
	Padding     []byte
	Mac         []byte
}

func (p Packet) ToBytes() []byte {
	payloadLenInBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(payloadLenInBytes, p.PacketLen)
	paddingLenInBytes := byte(p.PanndingLen)
	result := payloadLenInBytes
	result = append(result, paddingLenInBytes)
	result = append(result, p.Payload...)
	result = append(result, p.Padding...)
	result = append(result, p.Mac...)
	return result
}

func ParseSSHPacket(raw []byte) Packet {
	PacketLenInBytes := raw[:4]
	raw = raw[4:]
	PacketLen := binary.BigEndian.Uint32(PacketLenInBytes)
	paddingLen := uint8(raw[0])
	raw = raw[1:]
	payloadLen := PacketLen - uint32(paddingLen) - 1
	payload := raw[:payloadLen]
	raw = raw[payloadLen:]
	padding := raw[:paddingLen]
	mac := raw[paddingLen:]
	return Packet{
		PacketLen:   PacketLen,
		PanndingLen: paddingLen,
		Payload:     payload,
		Padding:     padding,
		Mac:         mac,
	}
}
