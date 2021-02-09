package transporter

import (
	"net"

	"github.com/Ehco1996/ehco/internal/constant"
	"github.com/Ehco1996/ehco/internal/lb"
)

// RelayTransporter
type RelayTransporter interface {
	GetOrCreateBufferCh(uaddr *net.UDPAddr) *BufferCh

	HandleTCPConn(c *net.TCPConn) error
	HandleUDPConn(uaddr *net.UDPAddr, local *net.UDPConn)
}

func PickTransporter(transType string, tcpLBNodes, udpLBNodes *lb.LBNodes) RelayTransporter {
	raw := Raw{
		TCPNodes:       tcpLBNodes,
		UDPNodes:       udpLBNodes,
		UDPBufferChMap: make(map[string]*BufferCh),
	}
	switch transType {
	case constant.Transport_RAW:
		return &raw
	}
	return nil
}
