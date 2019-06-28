Implementation example, drop the dll in the %PATH% and run.

	package main

	import (
		"fmt"
		"github.com/redtoasty/enetgo"
	)

	func main() {

		enetgo.Initialize()

		address := &enetgo.ENetAddress{}
		address.SetHostIP("127.0.0.1", 1234)

		host := enetgo.CreateHost(address, 100, 4, 100, 100, 8*1024)
		netEvent := &enetgo.ENetEvent{}

		fmt.Println("Starting loop")
		for {
			if host.CheckEvents(netEvent) > 0 || host.Service(netEvent, 1) > 0 {
				eventType := netEvent.GetType()
				switch eventType {
				case enetgo.ENET_EVENT_TYPE_CONNECT:
					fmt.Println("Connection event")
					break
				case enetgo.ENET_EVENT_TYPE_DISCONNECT:
					fmt.Println("Disconnection event")
					break
				case enetgo.ENET_EVENT_TYPE_RECEIVE:
					fmt.Println("Receive message event")
					break
				}
			}
		}
	}

