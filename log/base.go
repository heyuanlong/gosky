package log

import (
)
const (
	debugLevel  	 	= 0
	infoLevel 			= 1
	warningLevel		= 2
	errorLevel   		= 3
	fatalLevel  		= 4
)
type interfaceLog interface {
	Write(level int,format string, a ...interface{})
}
var glog interfaceLog

func init() {
	SetLog( newDefaultLog() )
}

func SetLog(lg interfaceLog)  {
	glog = lg
}
func Debug(format string, a ...interface{}) {
	glog.Write(debugLevel,format,a...)
}
func Info(format string, a ...interface{}) {
	glog.Write(infoLevel,format,a...)
}
func Warning(format string, a ...interface{}) {
	glog.Write(warningLevel,format,a...)
}
func Error(format string, a ...interface{}) {
	glog.Write(errorLevel,format,a...)
}
func Fatal(format string, a ...interface{}) {
	glog.Write(fatalLevel,format,a...)
}


