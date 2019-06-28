package enetgo

/*
#cgo LDFLAGS: -L. -lenet
#cgo CFLAGS: -I/include/ -DENET_NO_PRAGMA_LINK
 #include "enet.h"
*/
import "C"
import "unsafe"

// LibraryInitialize wraps the static inititalize we need to call
func LibraryInitialize() {
	// library init
	C.enet_initialize()
}

// ENetAddress refers to an ip/port address which we will connect to
type ENetAddress struct {
	native C.ENetAddress
}

// SetIPPort sets the IP and port of an existing ENetAddress object
func (address *ENetAddress) SetIPPort(ip string, port uint16) {
	address.native.port = C.ushort(port)
	cstr := C.CString(ip)
	defer C.free(unsafe.Pointer(cstr))
	C.enet_address_set_host_ip(&address.native, cstr)
}

// ENetHost holds a native
type ENetHost struct {
	native *C.ENetHost
}

// CreateHost builds a new ENetHost at a given IP address, with a given configuration
func CreateHost(address *ENetAddress, peerCount uint64, channelLimit uint64, inBandwidth uint32, outBandwith uint32, bufferSize int) (host *ENetHost) {
	newHost := new(ENetHost)
	newHost.native = (*C.ENetHost)(C.enet_host_create(&address.native, C.ulonglong(peerCount), C.ulonglong(channelLimit), C.uint(inBandwidth), C.uint(outBandwith), C.int(bufferSize)))
	return newHost
}

// CheckEvents checks if we have any new events waiting in the queue
func (host *ENetHost) CheckEvents() int {
	netEvent := &C.ENetEvent{}
	numEvents := C.enet_host_check_events(host.native, netEvent)
	return int(numEvents)
}

// Service polls the socket and relays back any events we get while waiting
func (host *ENetHost) Service(milliseconds uint) int {
	netEvent := &C.ENetEvent{}
	numEvents := C.enet_host_service(host.native, netEvent, C.uint(milliseconds))
	return int(numEvents)
}
