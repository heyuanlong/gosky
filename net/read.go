package net

import (
	"net"
	"time"
	"github.com/heyuanlong/gosky/log"
)


func handleClient(conn net.Conn)  {
	log.Debug("HandleClient")
	defer conn.Close()

	sc := newSConn(conn)
	g_Handler.OnConnect(sc)

	var bufBuf = make([]byte,0)
	var msgBuf = make([]byte, g_MSG_SIZE_MAX)
	for  {
		conn.SetReadDeadline(time.Now().Add(time.Duration(g_DeadLineTime) * time.Second))
		n , err := conn.Read(msgBuf)
		if err!= nil{
			if nerr, ok := err.(*net.OpError); ok && nerr.Timeout() {
				//log.Debug("timeout")
				g_Handler.OnTimeOut(sc)
				return
			}else {
				g_Handler.OnError(sc)
				//log.Debug("read close or fail")
				return
			}
		}
		if (len(bufBuf) + n ) >  g_BUF_SIZE_MAX {
			g_Handler.OnError(sc)
			return
		}
		bufBuf = append(bufBuf,msgBuf[0:n]...)
		msgLen,msgType,pBuf  := parsePackage(bufBuf)
		if msgLen == 0 {
			continue
		}

		//log.Debug("msgType:",msgType)

		if _,ok := mapHandler[msgType] ; ok{
			sb := &sMsgToOne{sc,msgType,pBuf}
			oneChanMsg <- sb
		}else{
			g_Handler.OnMessage(sc,msgType,pBuf)
		}
		bufBuf = bufBuf[msgLen:]
	}
}