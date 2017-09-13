package net

func oneRun( )  {

	for {
		v := <- oneChan;
		//v.sconn
		//v.msgType
		//v.data

		if call,ok := mapHandler[v.msgType] ; ok{
			call.h.Message(v.sconn,v.msgType,v.data)
		}
	}
}