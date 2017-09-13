package net

import (
	"net"
)

type SMsgToWrite struct {
	conn  net.Conn
	data []byte
}

type SMsgToOne struct {
	msgType int
	data []byte
}

type SConn struct {
	conn 		net.Conn
	ch 			chan *SMsgToWrite
}

type SHandler interface {
	Connect( *SConn,net.Addr )()
	DisConnect(*SConn)()
	Error(*SConn)()
	Message(*SConn,msgType int,[]byte)()
}

type structHandler struct {
	h SHandler
	isOneThread bool
}

