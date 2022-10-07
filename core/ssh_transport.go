package core

import (
	"errors"
	"go-ssh/common"
	"log"
	"net"
	"strings"
)

type sshTransport struct {
	conn net.Conn
}

func NewTransport(peerAddress string) (sshTransport, error) {

	conn, err := net.Dial("tcp", peerAddress)
	if err != nil {
		return sshTransport{}, err
	}
	transport := sshTransport{
		conn: conn,
	}
	if err = transport.initialHandshake(); err != nil {
		return transport, err
	}
	return transport, nil

}

func (this sshTransport) Close() {
	this.conn.Close()
}

func (this sshTransport) initialHandshake() error {

	if _, err := this.conn.Write([]byte(common.LOCAL_VERSION)); err != nil {
		return err
	}
	reply := make([]byte, common.MAX_FULL_PACKET_SIZE)
	if _, err := this.conn.Read(reply); err != nil {
		return err
	}
	reply_as_str := string(reply)
	if !strings.HasPrefix(reply_as_str, common.VERSION_PREFIX) {
		return errors.New(common.PEER_VERSION_ERROR)
	}
	log.Println("Peer version: " + reply_as_str)
	return nil
}
