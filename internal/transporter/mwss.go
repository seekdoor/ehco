package transporter

import (
	"net"

	"github.com/Ehco1996/ehco/internal/logger"
	"github.com/Ehco1996/ehco/internal/web"
)

type Mwss struct {
	raw *Raw
	mtp *mwssTransporter
}

func (s *Mwss) GetOrCreateBufferCh(uaddr *net.UDPAddr) *BufferCh {
	return s.raw.GetOrCreateBufferCh(uaddr)
}

func (s *Mwss) HandleUDPConn(uaddr *net.UDPAddr, local *net.UDPConn) {
	s.raw.HandleUDPConn(uaddr, local)
}

func (s *Mwss) HandleTCPConn(c *net.TCPConn) error {
	defer c.Close()
	remote := s.raw.TCPRemotes.Next()
	web.CurConnectionCount.WithLabelValues(remote.Label, web.METRIC_CONN_TCP).Inc()
	defer web.CurConnectionCount.WithLabelValues(remote.Label, web.METRIC_CONN_TCP).Dec()

	muxwsc, err := s.mtp.Dial(remote.Address + "/mwss/")
	if err != nil {
		return err
	}
	defer muxwsc.Close()
	logger.Infof("[mwss] HandleTCPConn from:%s to:%s", c.RemoteAddr(), remote.Label)
	return transport(muxwsc, c, remote.Label)
}
