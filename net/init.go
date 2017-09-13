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

var mapHandler map[int] *structHandler
var wChans [g_WRITE_GO_NUMS] chan *SMsgToWrite
var oneChan chan *SMsgToOne

func init() {
	mapHandler = make(map[int] *structHandler,0)
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		wChans[i] = make(chan *SMsgToWrite,g_WCHAN_BUF_LENS)
	}
	oneChan = make(chan *SMsgToOne,g_ONECHAN_BUF_LENS)
}