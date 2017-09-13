package net

import (
	"net"
	"time"
)


func handleClient(conn net.Conn)  {
	//klog.Klog.Println("HandleClient")
	defer conn.Close()

	sc := newSConn(conn)

	var bufBuf = make([]byte,0)
	var msgBuf = make([]byte, g_MSG_SIZE_MAX)
	for  {
		conn.SetReadDeadline(time.Now().Add(time.Duration(g_DeadLineTime) * time.Second))
		n , err := conn.Read(msgBuf)
		if err!= nil{
			if nerr, ok := err.(*net.OpError); ok && nerr.Timeout() {
				//klog.Klog.Println("timeout")
				return
			}else {
				//klog.Klog.Println("read close or fail")
				return
			}
		}
		if (len(bufBuf) + n ) >  g_BUF_SIZE_MAX {
			//klog.Klog.Println("buf too big")
			return
		}
		bufBuf = append(bufBuf,msgBuf[0:n]...)
		msgLen,msgType,pBuf  := ParsePackage(bufBuf)
		if msgLen == 0 {
			continue
		}

		//klog.Klog.Println("msgType:",msgType)
		if v,ok := mapHandler[msgType] ; ok{
			if v.isOneThread {
				sb := &SMsgToOne{sc,msgType,pBuf}
				oneChan <- sb
			}else {
				v.h.Message(sc,msgType,pBuf)
			}
		}else{
			//klog.Klog.Println("not reg this msgType:",msgType)
		}
		bufBuf = bufBuf[msgLen:]
	}
}