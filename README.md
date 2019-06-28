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
				fmt.Println("message")
			} else {
				fmt.Println("No message")
			}
		}
	}
