package net

import (
	"net"
	"strconv"

	"github.com/heyuanlong/gosky/log"
)

func SetHander( st SHandler)  {
	g_Handler = st
}

func Start(port int) error {
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		go writeRun(i)
	}
	go oneRun()

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
		go handleClient(new_conn)
	}
}

func RegOne(msgType int)  {
	mapHandler[msgType] = true
}
