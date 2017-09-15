package net

import (
)

func oneRun( )  {

	for {
		select {
		case v := <-oneChanMsg:
			//v.sconn
			//v.msgType
			//v.data

			g_Handler.OnMessage(v.sconn, v.msgType, v.data)
		default:
		}
	}

}