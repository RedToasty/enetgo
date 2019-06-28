package enetgo

/*
#cgo windows LDFLAGS: -L. -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I/include/ -DENET_NO_PRAGMA_LINK
 #include "enet.h"
*/
import "C"

type ENetPacketFlag int

type ENetCallbacks C.ENetCallbacks

const (
	ENET_PACKET_FLAG_NONE                  = ENetPacketFlag(C.ENET_PACKET_FLAG_NONE)
	ENET_PACKET_FLAG_RELIABLE              = ENetPacketFlag(C.ENET_PACKET_FLAG_RELIABLE)
	ENET_PACKET_FLAG_UNSEQUENCED           = ENetPacketFlag(C.ENET_PACKET_FLAG_UNSEQUENCED)
	ENET_PACKET_FLAG_NO_ALLOCATE           = ENetPacketFlag(C.ENET_PACKET_FLAG_NO_ALLOCATE)
	ENET_PACKET_FLAG_UNRELIABLE_FRAGMENTED = ENetPacketFlag(C.ENET_PACKET_FLAG_UNRELIABLE_FRAGMENTED)
	ENET_PACKET_FLAG_SENT                  = ENetPacketFlag(C.ENET_PACKET_FLAG_SENT)
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
