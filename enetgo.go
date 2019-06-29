package enetgo

/*
#cgo windows LDFLAGS: -L${SRCDIR}/libs/ -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I${SRCDIR}/include/ -DENET_NO_PRAGMA_LINK -g
 #include "enet.h"
*/
import "C"

type ENetCallbacks C.ENetCallbacks

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
