package net

import (

)

func writeRun( i int)  {
	ch :=  wChans[i]
	for {
		v := <-ch;
		//v.writeBuf
		v.sconn.conn.Write(v.data)
	}
}



func (p *SConn) Write(msgType int, data []byte)  {
	buf := SetPackage(msgType,data)
	st := &SMsgToWrite{p,buf}
	p.ch <- st
}
