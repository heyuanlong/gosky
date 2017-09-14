package net

import (
	"net"
	"math/rand"
)


type SConn struct {
	conn 		net.Conn
	ch 			chan *sMsgToWrite

	writeBuf	[]byte
}
type sMsgToWrite struct {
	sconn  *SConn
	data []byte
}
type sHandleToOne struct {
	sconn  *SConn
	msgType int      // 1:connect 2:disconnect 3:error
}

type sMsgToOne struct {
	sconn  *SConn
	msgType int
	data []byte
}

type SHandler interface {
	Connect( *SConn)()
	DisConnect(*SConn)()
	Error(*SConn)()
	Message(*SConn,int,[]byte)()
}


func newSConn(conn net.Conn) *SConn {
	psc := &SConn{}
	psc.conn = conn
	psc.ch = wChans[rand.Intn(g_WRITE_GO_NUMS)]
	psc.writeBuf = make([]byte,0)
	return psc
}