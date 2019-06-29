package enetgo

/*
#cgo windows LDFLAGS: -L${SRCDIR}/libs/ -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I${SRCDIR}/include/ -DENET_NO_PRAGMA_LINK -g
 #include "enet.h"
*/
import "C"
import "unsafe"

// ENetAddress refers to an ip/port address which we will connect to
type ENetAddress C.ENetAddress

// SetHostIP sets the IP and port of an existing ENetAddress object
func (address *ENetAddress) SetHostIP(ip string, port uint16) {
	address.port = C.ushort(port)
	cstr := C.CString(ip)
	defer C.free(unsafe.Pointer(cstr))
	C.enet_address_set_host_ip((*C.ENetAddress)(address), cstr)
}
