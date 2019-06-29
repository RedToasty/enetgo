package enetgo

/*
#cgo windows LDFLAGS: -L${SRCDIR}/libs/ -lenet
#cgo linux LDFLAGS: -L/usr/local/lib -lenet
#cgo CFLAGS: -I${SRCDIR}/include/ -DENET_NO_PRAGMA_LINK -g
 #include "enet.h"
*/
import "C"
import "unsafe"

type ENetPeer C.ENetPeer
type ENetPeerState int

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

func (p *ENetPeer) ConnectID() int {
	return int(p.connectID)
}

// SetData can set a reference to any arbitrary Go data.
func (p *ENetPeer) SetData(data interface{}) {
	p.data = unsafe.Pointer(&data)
}

// Data returns referenced Go data.  Must be dereferenced as a pointer.
//
//   peer.Data().(*MyType)
func (p *ENetPeer) Data() interface{} {
	return unsafe.Pointer(p.data)
}

// Public

func (peer *ENetPeer) State() ENetPeerState {
	return ENetPeerState((*C.ENetPeer)(peer).state)
}

func (p *ENetPeer) Send(channelID uint8, packet *ENetPacket) int {
	return int(C.enet_peer_send((*C.ENetPeer)(p), C.enet_uint8(channelID), (*C.ENetPacket)(packet)))
}

func (p *ENetPeer) Receive(channelID *C.enet_uint8) *ENetPacket {
	return (*ENetPacket)(C.enet_peer_receive((*C.ENetPeer)(p), channelID))
}

func (p *ENetPeer) Ping() {
	C.enet_peer_ping((*C.ENetPeer)(p))
}

func (p *ENetPeer) PingInterval(pingInterval uint32) {
	C.enet_peer_ping_interval((*C.ENetPeer)(p), C.enet_uint32(pingInterval))
}

func (p *ENetPeer) Timeout(timeoutLimit uint32, timeoutMinimum uint32, timeoutMaximum uint32) {
	C.enet_peer_timeout((*C.ENetPeer)(p), C.enet_uint32(timeoutLimit), C.enet_uint32(timeoutMinimum), C.enet_uint32(timeoutMaximum))
}

func (p *ENetPeer) Reset() {
	C.enet_peer_reset((*C.ENetPeer)(p))
}

func (p *ENetPeer) Disconnect(data uint32) {
	C.enet_peer_disconnect((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) DisconnectNow(data uint32) {
	C.enet_peer_disconnect_now((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) DisconnectLater(data uint32) {
	C.enet_peer_disconnect_later((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) ThrottleConfigure(interval, acceleration uint32, deceleration uint32, threshold uint32) {
	C.enet_peer_throttle_configure((*C.ENetPeer)(p), C.enet_uint32(interval), C.enet_uint32(acceleration), C.enet_uint32(deceleration), C.enet_uint32(threshold))
}
