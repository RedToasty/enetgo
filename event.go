package enetgo

/*
#cgo windows LDFLAGS: -L. -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I/include/ -DENET_NO_PRAGMA_LINK
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

func (event *ENetEvent) GetType() ENetEventType {
	return ENetEventType((*C.ENetEvent)(event).eType)
}
