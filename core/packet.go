package core

import (
	"encoding/binary"
	"errors"
	"go-ssh/common"
)

type Packet struct {
	PacketLen   uint32
	PanndingLen uint8
	Payload     []byte
	Padding     []byte
	Mac         []byte
}

func (this Packet) ToBytes() []byte {
	packetLenInBytes := make([]byte, common.PACKET_LEN_BYTES_COUNT)
	binary.BigEndian.PutUint32(packetLenInBytes, this.PacketLen)
	paddingLenInBytes := byte(this.PanndingLen)
	result := append(packetLenInBytes, paddingLenInBytes)
	result = append(result, this.Payload...)
	result = append(result, this.Padding...)
	result = append(result, this.Mac...)
	return result
}

func ParseSSHPacket(raw []byte) (Packet, error) {
	rawLen := len(raw)
	if rawLen <= common.PACKET_LEN_BYTES_COUNT {
		return Packet{}, errors.New(common.BAD_PACKET_ERROR)
	}
	PacketLenInBytes := raw[:common.PACKET_LEN_BYTES_COUNT]
	PacketLen := binary.BigEndian.Uint32(PacketLenInBytes)
	if uint32(rawLen) < PacketLen {
		return Packet{}, errors.New(common.BAD_PACKET_ERROR)
	}
	raw = raw[common.PACKET_LEN_BYTES_COUNT:]
	paddingLen := uint8(raw[0])
	raw = raw[common.PADDING_LEN_BYTES_COUNT:]
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
