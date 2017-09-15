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

var g_DeadLineTime	  =	60

var g_Handler SHandler
var mapHandler map[int] bool

var wChans [g_WRITE_GO_NUMS] chan *sMsgToWrite
var oneChanMsg 		chan *sMsgToOne

func init() {
	mapHandler = make(map[int] bool,0)
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		wChans[i] = make(chan *sMsgToWrite,g_WCHAN_BUF_LENS)
	}
	oneChanMsg = make(chan *sMsgToOne,g_ONECHAN_BUF_LENS)
}