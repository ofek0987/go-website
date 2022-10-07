package main

import (
	"fmt"
	"go-ssh/core"
)

func main() {
	peer := "192.168.1.13:22"
	fmt.Println("Starting SSH")
	transport, err := core.NewTransport(peer)
	if err != nil {
		panic(err)
	}
	transport.Close()
}
