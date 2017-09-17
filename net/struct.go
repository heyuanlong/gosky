package net

import (
	"net"

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

type sMsgToOne struct {
	sconn  *SConn
	msgType int
	data []byte
}

type SHandler interface {
	OnConnect( *SConn)()
	OnDisConnect(*SConn)()
	OnTimeOut(*SConn)()
	OnError(*SConn)()
	OnMessage(*SConn,int,[]byte)()
}


func newSConn(conn net.Conn,ch chan *sMsgToWrite) *SConn {
	psc := &SConn{}
	psc.conn = conn
	psc.ch = ch
	psc.writeBuf = make([]byte,0)
	return psc
}
func (p *SConn) Write(msgType int, data []byte)  {
	buf := SetPackage(msgType,data)
	st := &sMsgToWrite{p,buf}
	p.ch <- st
}
func (p *SConn) Close() error {
	return p.conn.Close()
}