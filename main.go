package main

/*
#cgo LDFLAGS: -lWS2_32 -lWinMM
#cgo CFLAGS: -DENET_NO_PRAGMA_LINK -g
 #include "enet.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	// library init
	C.enet_initialize()

	// set up ip/port address
	address := &C.ENetAddress{}
	address.port = 1234
	cstr := C.CString("127.0.0.1")
	defer C.free(unsafe.Pointer(cstr))
	C.enet_address_set_host_ip(address, cstr)

	// create host
	host := (*C.ENetHost)(C.enet_host_create(address, 100, 1, 100, 100, 100))

	// poll for events
	netEvent := &C.ENetEvent{}
	for {
		if C.enet_host_check_events(host, netEvent) > 0 || C.enet_host_service(host, netEvent, 1) > 0 {
			fmt.Println("Message")
		}
	}
}
