package core

import (
	"errors"
	"log"
	"net"
	"strings"

	"github.com/ofek0987/gssh/common"
)

type sshTransport struct {
	conn net.Conn
}

func NewClientTransport(peerAddress string) (sshTransport, error) {

	conn, err := net.Dial("tcp", peerAddress)
	if err != nil {
		return sshTransport{}, err
	}
	transport := sshTransport{
		conn: conn,
	}
	if err = transport.initialHandshakeClient(); err != nil {
		return transport, err
	}
	return transport, nil

}

func (this sshTransport) Close() {
	this.conn.Close()
}

func (this sshTransport) initialHandshakeClient() error {

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
