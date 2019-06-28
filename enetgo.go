package enetgo

/*
#cgo windows LDFLAGS: -L. -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I/include/ -DENET_NO_PRAGMA_LINK
 #include "enet.h"
*/
import "C"

type ENetPeerState int
type ENetPacketFlag int
type ENetEventType int

type ENetCallbacks C.ENetCallbacks

const (
	ENET_PACKET_FLAG_NONE                  = ENetPacketFlag(C.ENET_PACKET_FLAG_NONE)
	ENET_PACKET_FLAG_RELIABLE              = ENetPacketFlag(C.ENET_PACKET_FLAG_RELIABLE)
	ENET_PACKET_FLAG_UNSEQUENCED           = ENetPacketFlag(C.ENET_PACKET_FLAG_UNSEQUENCED)
	ENET_PACKET_FLAG_NO_ALLOCATE           = ENetPacketFlag(C.ENET_PACKET_FLAG_NO_ALLOCATE)
	ENET_PACKET_FLAG_UNRELIABLE_FRAGMENTED = ENetPacketFlag(C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENTED)
	ENET_PACKET_FLAG_SENT                  = ENetPacketFlag(C.ENET_PACKET_FLAG_SENT)
)

const (
	ENET_PEER_STATE_DISCONNECTED             = ENetPeerState(C.ENET_PEER_STATE_DISCONNECTED)
	ENET_PEER_STATE_CONNECTING               = ENetPeerState(C.ENET_PEER_STATE_CONNECTING)
	ENET_PEER_STATE_ACKNOWLEDGING_CONNECT    = ENetPeerState(C.ENET_PEER_STATE_ACKNOWLEDGING_CONNECT)
	ENET_PEER_STATE_CONNECTION_PENDING       = ENetPeerState(C.ENET_PEER_STATE_CONNECTION_PENDING)
	ENET_PEER_STATE_CONNECTION_SUCCEEDED     = ENetPeerState(C.ENET_PEER_STATE_CONNECTION_SUCCEEDED)
	ENET_PEER_STATE_CONNECTED                = ENetPeerState(C.ENET_PEER_STATE_CONNECTED)
	ENET_PEER_STATE_DISCONNECT_LATER         = ENetPeerState(C.ENET_PEER_STATE_DISCONNECT_LATER)
	ENET_PEER_STATE_DISCONNECTING            = ENetPeerState(C.ENET_PEER_STATE_DISCONNECTING)
	ENET_PEER_STATE_ACKNOWLEDGING_DISCONNECT = ENetPeerState(C.ENET_PEER_STATE_ACKNOWLEDGING_DISCONNECT)
	ENET_PEER_STATE_ZOMBIE                   = ENetPeerState(C.ENET_PEER_STATE_ZOMBIE)
)

const (
	ENET_EVENT_TYPE_NONE               = ENetEventType(C.ENET_EVENT_TYPE_NONE)
	ENET_EVENT_TYPE_CONNECT            = ENetEventType(C.ENET_EVENT_TYPE_CONNECT)
	ENET_EVENT_TYPE_DISCONNECT         = ENetEventType(C.ENET_EVENT_TYPE_DISCONNECT)
	ENET_EVENT_TYPE_RECEIVE            = ENetEventType(C.ENET_EVENT_TYPE_RECEIVE)
	ENET_EVENT_TYPE_DISCONNECT_TIMEOUT = ENetEventType(C.ENET_EVENT_TYPE_DISCONNECT_TIMEOUT)
)

// Initialize wraps the static inititalize we need to call
func Initialize() int {
	return int(C.enet_initialize())
}

// InitializeWithCallbacks allows us to bing callbacks to the setup events
func InitializeWithCallbacks(version int, inits *ENetCallbacks) int {
	return int(C.enet_initialize_with_callbacks(C.ENetVersion(version), (*C.ENetCallbacks)(inits)))
}

// Deinitialize shuts down enet
func Deinitialize() {
	C.enet_deinitialize()
}
