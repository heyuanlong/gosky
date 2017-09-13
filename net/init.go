package net

import (
	"net"
)

const (
	g_MSG_SIZE_MAX    	= 65536
	g_BUF_SIZE_MAX    	= 655360

	g_WRITE_GO_NUMS		= 8
	g_CHAN_BUF_LENS		= 256
)


var g_DeadLineTime	  =	60

var mapHandler map[int] *structHandler
var wChans [g_WRITE_GO_NUMS] chan *SMsgToWrite
var oneChan chan *SMsgToOne

func init() {
	mapHandler = make(map[int] Handler,0)
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		chans[i] = make(chan *SBytes,g_CHAN_BUF_LENS)
	}
}