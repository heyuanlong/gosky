package net

import (

)

const (
	g_MSG_SIZE_MAX    		= 65536
	g_BUF_SIZE_MAX    		= 655360

	g_WRITE_GO_NUMS			= 8
	g_WCHAN_BUF_LENS		= 256
	g_ONECHAN_BUF_LENS		= 256
)


type gnet struct{
	g_DeadLineTime int
	g_Handler SHandler
	mapHandler map[int] bool

	wChans [g_WRITE_GO_NUMS] chan *sMsgToWrite
	oneChanMsg 		chan *sMsgToOne
}

func NewGnet() {
	net := new(gnet)
	net.g_DeadLineTime = 60
	net.mapHandler = make(map[int] bool,0)
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		net.wChans[i] = make(chan *sMsgToWrite,g_WCHAN_BUF_LENS)
	}
	net.oneChanMsg = make(chan *sMsgToOne,g_ONECHAN_BUF_LENS)
}