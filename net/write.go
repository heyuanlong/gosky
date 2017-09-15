package net

import (
	"github.com/heyuanlong/gosky/log"
)

func writeRun( i int)  {
	ch :=  wChans[i]
	for {
		v := <-ch;
			dealWrite(v)
	}
}

func dealWrite(v *sMsgToWrite)  {
	if len(v.sconn.writeBuf) > 0 {
		v.data = append(v.sconn.writeBuf,v.data...)
	}
	dataLen := len(v.data)
	n,err := v.sconn.conn.Write(v.data)
	if err != nil {
		v.sconn.Close()
	}
	if dataLen > n {
		log.Info("datalen:%d n:%d",dataLen,n)
		v.sconn.writeBuf = v.data[n:]
	}
}


func (p *SConn) Write(msgType int, data []byte)  {
	buf := setPackage(msgType,data)
	st := &sMsgToWrite{p,buf}
	p.ch <- st
}
