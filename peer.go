package enetgo

/*
#cgo windows LDFLAGS: -L. -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I/include/ -DENET_NO_PRAGMA_LINK
 #include "enet.h"
*/
import "C"

type ENetPeer C.ENetPeer
