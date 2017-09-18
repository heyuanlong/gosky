package main

import (
	"github.com/heyuanlong/gosky/log"
	"github.com/heyuanlong/gosky/net"
)


type xHandler struct {

}


func (*xHandler) OnConnect(*net.SConn)(){
	log.Info("OnConnect")
}
func (*xHandler) OnDisConnect(*net.SConn)(){
	log.Info("OnDisConnect")
}
func (*xHandler) OnTimeOut(*net.SConn)(){
	log.Info("OnTimeOut")
}
func (*xHandler) OnError(*net.SConn)(){
	log.Info("OnError")
}
func (*xHandler) OnMessage(p *net.SConn,msgType int,data []byte)(){
	//log.Info("OnMessage:%d %s",msgType, string(data))
	p.Write(11,data)
}

func main() {
	log.Info("start")
	xh :=  new(xHandler)


	gnet := net.NewGnet()
	gnet.RegOne(11)
	gnet.SetHander(xh)
	go gnet.Start(8089)

	gnet2 := net.NewGnet()
	gnet2.RegOne(11)
	gnet2.SetHander(xh)
	gnet2.Start(8088)
}