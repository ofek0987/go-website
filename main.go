package main

import (
	"errors"
	"fmt"
	"go-ssh/core"
	"log"
	"net"
	"os"
	"strings"
)

const (
	LOCAL_VERSION        = "SSH-2.0-gossh_0.1\r\n"
	MAX_FULL_PACKET_SIZE = 35000
)

type SSHTrasport struct {
	conn net.Conn
}

func checkFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

func initial_handshake(conn net.Conn) error {
	_, err := conn.Write([]byte(LOCAL_VERSION))
	checkFatalErr(err)
	reply := make([]byte, MAX_FULL_PACKET_SIZE)
	_, err = conn.Read(reply)
	checkFatalErr(err)
	reply_as_str := string(reply)
	if !strings.HasPrefix(reply_as_str, "SSH-") {
		return errors.New("Bad peer vesrion")
	}
	log.Println("Peer version: " + reply_as_str)
	return nil
}
func handleKexInit(conn net.Conn) {
	reply := make([]byte, MAX_FULL_PACKET_SIZE)
	len, err := conn.Read(reply)
	checkFatalErr(err)
	fmt.Println(len)
	packet := core.ParseSSHPacket(reply[:len])
	fmt.Println(packet.Mac)
}
func main() {
	peer := "192.168.1.13:22"
	fmt.Println("Starting SSH")
	conn, err := net.Dial("tcp", peer)
	checkFatalErr(err)
	defer conn.Close()
	initial_handshake(conn)
	handleKexInit(conn)
}
