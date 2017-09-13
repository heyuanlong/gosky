package net

import (

)

func writeRun( i int)  {
	ch :=  wChans[i]
	for {
		v := <-ch;
		v.conn.Write(v.data)
	}
}



func (p *SConn) Write(msgType int, data []byte)  {
	buf := SetPackage(msgType,data)
	st := &SMsgToWrite{p.conn,buf}
	p.ch <- st
}
