package net

import (
	"net"
	"time"
	"math/rand"
	"github.com/heyuanlong/gosky/log"
)


func (p *gnet) handleClient(conn net.Conn)  {
	log.Debug("HandleClient")
	defer conn.Close()

	sc := newSConn(conn,p.wChans[rand.Intn(g_WRITE_GO_NUMS)])
	p.g_Handler.OnConnect(sc)

	bufBuf := make([]byte,0,g_BUF_SIZE_MAX)
	msgBuf := make([]byte, g_MSG_SIZE_MAX)
	for  {
		
		conn.SetReadDeadline(time.Now().Add(time.Duration(p.g_DeadLineTime) * time.Second))
		n , err := conn.Read(msgBuf)
		if err!= nil{
			if nerr, ok := err.(*net.OpError); ok && nerr.Timeout() {
				p.g_Handler.OnTimeOut(sc)
				return
			}else {
				p.g_Handler.OnError(sc)
				return
			}
		}
		if (len(bufBuf) + n ) >  g_BUF_SIZE_MAX {
			p.g_Handler.OnError(sc)
			return
		}
		bufBuf = append(bufBuf,msgBuf[0:n]...)
		msgLen,msgType,pBuf  := ParsePackage(bufBuf)
		if msgLen == 0 {
			continue
		}

		if _,ok := p.mapHandler[msgType] ; ok{
			sb := &sMsgToOne{sc,msgType,pBuf}
			p.oneChanMsg <- sb
		}else{
			p.g_Handler.OnMessage(sc,msgType,pBuf)
		}
		bufBuf = bufBuf[msgLen:]
	}
}