package core

import (
	"encoding/binary"
	"errors"
)

const (
	BAD_PACKET_ERROR       = "Bad packet"
	PACKET_LEN_BITS_COUNT  = 4
	PADDING_LEN_BITS_COUNT = 1
)

type Packet struct {
	PacketLen   uint32
	PanndingLen uint8
	Payload     []byte
	Padding     []byte
	Mac         []byte
}

func (p Packet) ToBytes() []byte {
	payloadLenInBytes := make([]byte, PACKET_LEN_BITS_COUNT)
	binary.BigEndian.PutUint32(payloadLenInBytes, p.PacketLen)
	paddingLenInBytes := byte(p.PanndingLen)
	result := payloadLenInBytes
	result = append(result, paddingLenInBytes)
	result = append(result, p.Payload...)
	result = append(result, p.Padding...)
	result = append(result, p.Mac...)
	return result
}

func ParseSSHPacket(raw []byte) (Packet, error) {
	rawLen := len(raw)
	if rawLen <= PACKET_LEN_BITS_COUNT {
		return Packet{}, errors.New(BAD_PACKET_ERROR)
	}
	PacketLenInBytes := raw[:PACKET_LEN_BITS_COUNT]
	PacketLen := binary.BigEndian.Uint32(PacketLenInBytes)
	if uint32(rawLen) < PacketLen {
		return Packet{}, errors.New(BAD_PACKET_ERROR)
	}
	raw = raw[PACKET_LEN_BITS_COUNT:]
	paddingLen := uint8(raw[0])
	raw = raw[PADDING_LEN_BITS_COUNT:]
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
	}, nil
}
