package net

import (
)

func (p *gnet) oneRun( )  {

	for {
		select {
		case v := <-p.oneChanMsg:
			//v.sconn
			//v.msgType
			//v.data
			p.g_Handler.OnMessage(v.sconn, v.msgType, v.data)
		default:
		}
	}

}