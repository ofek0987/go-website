package main

import (
	"fmt"

	"github.com/ofek0987/gssh/core"
)

func main() {
	peer := "172.17.0.3:22"
	fmt.Println("Starting SSH")
	transport, err := core.NewClientTransport(peer)
	if err != nil {
		panic(err)
	}
	transport.Close()
}
