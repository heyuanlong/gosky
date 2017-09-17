package net

import (
	"net"
	"strconv"

	"github.com/heyuanlong/gosky/log"
)

type gnet struct{
	g_DeadLineTime int
	g_Handler SHandler
	mapHandler map[int] bool

	wChans [g_WRITE_GO_NUMS] chan *sMsgToWrite
	oneChanMsg 		chan *sMsgToOne
}

func NewGnet() *gnet {
	net := new(gnet)
	net.g_DeadLineTime = 60
	net.mapHandler = make(map[int] bool,0)
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		net.wChans[i] = make(chan *sMsgToWrite,g_WCHAN_BUF_LENS)
	}
	net.oneChanMsg = make(chan *sMsgToOne,g_ONECHAN_BUF_LENS)
	return net
}
func (p *gnet) SetHander( st SHandler)  {
	p.g_Handler = st
}

func (p *gnet) Start(port int) error {
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		go p.writeRun(i)
	}
	go p.oneRun()

	serverPort :=  strconv.Itoa(port)
	listen_sock,err := net.Listen("tcp",":"+serverPort)
	if err != nil{
		log.Error(err.Error())
		return err
	}
	defer listen_sock.Close()
	for{
		new_conn,err := listen_sock.Accept()
		if err != nil {
			log.Error("listen_sock.Accept error:%s",err.Error())
			continue
		}
		go p.handleClient(new_conn)
	}
}

func (p *gnet) RegOne(msgType int)  {
	p.mapHandler[msgType] = true
}
