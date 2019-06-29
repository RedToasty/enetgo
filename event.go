package enetgo

/*
#cgo windows LDFLAGS: -L${SRCDIR}/libs/ -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I${SRCDIR}/include/ -DENET_NO_PRAGMA_LINK -g
 #include "enet.h"
*/
import "C"

type ENetEvent C.ENetEvent
type ENetEventType int

const (
	ENET_EVENT_TYPE_NONE               = ENetEventType(C.ENET_EVENT_TYPE_NONE)
	ENET_EVENT_TYPE_CONNECT            = ENetEventType(C.ENET_EVENT_TYPE_CONNECT)
	ENET_EVENT_TYPE_DISCONNECT         = ENetEventType(C.ENET_EVENT_TYPE_DISCONNECT)
	ENET_EVENT_TYPE_RECEIVE            = ENetEventType(C.ENET_EVENT_TYPE_RECEIVE)
	ENET_EVENT_TYPE_DISCONNECT_TIMEOUT = ENetEventType(C.ENET_EVENT_TYPE_DISCONNECT_TIMEOUT)
)

func (event *ENetEvent) EventType() ENetEventType {
	return ENetEventType((*C.ENetEvent)(event).eventType)
}

func (event *ENetEvent) Peer() *ENetPeer {
	return (*ENetPeer)((*C.ENetEvent)(event).peer)
}

func (event *ENetEvent) Packet() *ENetPacket {
	return (*ENetPacket)(event.packet)
}

func (event *ENetEvent) ChannelID() int {
	return int(event.channelID)
}

func (event *ENetEvent) DataLength() int {
	return int(event.packet.dataLength)
}
