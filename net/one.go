package net

import (
	"github.com/heyuanlong/gosky/log"
)

func oneRun( )  {

	for {
		select {
		case v := <-oneChanMsg:
			//v.sconn
			//v.msgType
			//v.data

			g_Handler.Message(v.sconn, v.msgType, v.data)


		case v := <-oneChanHandle:
			log.Info("msgType:%d", v.msgType)
			switch v.msgType {
			case 1:
				g_Handler.Connect(v.sconn)
			case 2:
				g_Handler.DisConnect(v.sconn)
			case 3:
				g_Handler.Error(v.sconn)
			default:
			}

		default:
		}
	}

}