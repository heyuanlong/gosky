package net

import (
	"net"
	"strconv"
)


func Start(port int) error {
	for i:=0; i < g_WRITE_GO_NUMS ; i++ {
		go writeRun(i)
	}
	go oneRun()

	serverPort :=  strconv.Itoa(port)
	listen_sock,err := net.Listen("tcp",":"+serverPort)
	if err != nil{
		//klog.Klog.Fatalln(err)
		return err
	}
	defer listen_sock.Close()
	for{
		new_conn,err := listen_sock.Accept()
		if err != nil {
			//klog.Klog.Println("listen_sock.Accept error:",err)
			continue
		}
		go handleClient(new_conn)
	}
}

func Reg(msgType int, h SHandler,isOneThread bool)  {
	st := &structHandler{h ,isOneThread}
	mapHandler[msgType] = st
}
