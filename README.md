Implementation example, drop the dll in the %PATH% and run.

	package main
	
	import (
		"fmt"
		"github.com/redtoasty/enetgo"
	)
	
	func main() {
	
		enetgo.LibraryInitialize()
		
		address := &enetgo.ENetAddress{}
		address.SetIPPort("127.0.0.1", 1234)

		host := enetgo.CreateHost(address, 100, 4, 100, 100, 8*1024)
		for {
			if host.CheckEvents() > 0 || host.Service(1) > 0 {
				fmt.Println("message")
			}
		}
	}
