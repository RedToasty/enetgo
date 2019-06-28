package enetgo

/*
#cgo windows LDFLAGS: -L. -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I/include/ -DENET_NO_PRAGMA_LINK
 #include "enet.h"
*/
import "C"

type ENetHost C.ENetHost
type ENetEvent C.ENetEvent

// CreateHost builds a new ENetHost at a given IP address, with a given configuration
func CreateHost(address *ENetAddress, peerCount uint64, channelLimit uint64, inBandwidth uint32, outBandwith uint32, bufferSize int) (host *ENetHost) {
	return (*ENetHost)(C.enet_host_create((*C.ENetAddress)(address), C.size_t(peerCount), C.size_t(channelLimit), C.uint(inBandwidth), C.uint(outBandwith), C.int(bufferSize)))
}

func (host *ENetHost) Destroy() {
	C.enet_host_destroy((*C.ENetHost)(host))
}

func (host *ENetHost) Connect(address *ENetAddress, channelCount int, data int) *ENetPeer {
	return (*ENetPeer)(C.enet_host_connect((*C.ENetHost)(host), (*C.ENetAddress)(address), C.size_t(channelCount), C.enet_uint32(data)))
}

// CheckEvents checks if we have any new events waiting in the queue
func (host *ENetHost) CheckEvents(event *ENetEvent) int {
	return int(C.enet_host_check_events((*C.ENetHost)(host), (*C.ENetEvent)(event)))
}

// Service polls the socket and relays back any events we get while waiting
func (host *ENetHost) Service(event *ENetEvent, milliseconds uint) int {
	return int(C.enet_host_service((*C.ENetHost)(host), (*C.ENetEvent)(event), C.uint(milliseconds)))
}
