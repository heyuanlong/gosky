package log

import (
	"github.com/heyuanlong/gosky/conf"
	syslog "log"
	"os"
)

type defaultLog struct {
	level int
	levelTags [5]string
	log *syslog.Logger
}

func newDefaultLog() *defaultLog {
	d := new(defaultLog)
	d.level = debugLevel
	d.levelTags = [5]string{
		"[debug] ",
		"[info ] ",
		"[warn ] ",
		"[error] ",
		"[fatal] "}


	logfile,err := conf.GetString("server","logfile")
	if err != nil {
		d.log = syslog.New(os.Stdout,"",syslog.LstdFlags)
	}else {
		logfd,err  := os.Create(logfile)
		if err != nil {
			d.log = syslog.New(os.Stdout,"",syslog.LstdFlags)
		}
		d.log = syslog.New(logfd,"",syslog.LstdFlags)
	}
	return d
}



func (p *defaultLog) Write(level int,format string, a ...interface{})  {
	if level < p.level {
		return
	}
	p.log.Printf(p.levelTags[level] + format,a...)
	if level == fatalLevel {
		os.Exit(1)
	}
}


