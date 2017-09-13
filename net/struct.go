package net

import (
	"net"
	"math/rand"
)


type SConn struct {
	conn 		net.Conn
	ch 			chan *SMsgToWrite

	writeBuf	[]byte
}
type SMsgToWrite struct {
	sconn  *SConn
	data []byte
}
type SMsgToOne struct {
	sconn  *SConn
	msgType int
	data []byte
}


type SHandler interface {
	Connect( *SConn,net.Addr )()
	DisConnect(*SConn)()
	Error(*SConn)()
	Message(*SConn,int,[]byte)()
}

type structHandler struct {
	h SHandler
	isOneThread bool
}





func newSConn(conn net.Conn) *SConn {
	psc := &SConn{}
	psc.conn = conn
	psc.ch = wChans[rand.Intn(g_WRITE_GO_NUMS)]
	psc.writeBuf = make([]byte,0)
	return psc
}