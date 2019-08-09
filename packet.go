package enetgo

/*
#cgo windows LDFLAGS: -L${SRCDIR}/libs/ -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I${SRCDIR}/include/ -DENET_NO_PRAGMA_LINK -g
 #include "enet.h"
*/
import "C"
import "unsafe"

type ENetPacket C.ENetPacket
type ENetPacketFlag int

const (
	ENET_PACKET_FLAG_NONE                  = ENetPacketFlag(C.ENET_PACKET_FLAG_NONE)
	ENET_PACKET_FLAG_RELIABLE              = ENetPacketFlag(C.ENET_PACKET_FLAG_RELIABLE)
	ENET_PACKET_FLAG_UNSEQUENCED           = ENetPacketFlag(C.ENET_PACKET_FLAG_UNSEQUENCED)
	ENET_PACKET_FLAG_NO_ALLOCATE           = ENetPacketFlag(C.ENET_PACKET_FLAG_NO_ALLOCATE)
	ENET_PACKET_FLAG_UNRELIABLE_FRAGMENTED = ENetPacketFlag(C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENTED)
	ENET_PACKET_FLAG_SENT                  = ENetPacketFlag(C.ENET_PACKET_FLAG_SENT)
	ENET_PACKET_FLAG_INSTANT               = ENetPacketFlag(C.ENET_PACKET_FLAG_INSTANT)
)

func NewPacket(data []byte, dataLength int, flags ENetPacketFlag) *ENetPacket {
	return (*ENetPacket)(C.enet_packet_create(unsafe.Pointer(&data[0]), C.size_t(dataLength), C.enet_uint32(flags)))
}

func (p *ENetPacket) Data() []byte {
	return (*[1 << 30]byte)(unsafe.Pointer(p.data))[0:p.DataLength()]
}

func (p *ENetPacket) DataLength() int {
	return int(p.dataLength)
}

func (p *ENetPacket) Destroy() {
	C.enet_packet_destroy((*C.ENetPacket)(p))
}
