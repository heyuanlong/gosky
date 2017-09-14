package conf

import (
	"github.com/robfig/config"
)
var logfile string

func init() {
	logfile = "conf/conf.cnf"
}

func SetFile(f string)  {
	logfile = f
}

func GetConf() (*config.Config,error) {
	c, err := config.ReadDefault(logfile)
	if err != nil {
		return nil,err
	}
	return c,nil
}

func GetString(head string ,key string) (value string, err error) {
	c , err := GetConf()
	if err != nil {
		return "", err
	}
	return c.String(head,key)
}

func GetInt(head string ,key string) (value int, err error) {
	c , err := GetConf()
	if err != nil {
		return 0, err
	}
	return c.Int(head,key)
}
func GetBool(head string ,key string) (value bool, err error) {
	c , err := GetConf()
	if err != nil {
		return false, err
	}
	return c.Bool(head,key)
}
func GetFloat(head string ,key string) (value float64, err error) {
	c , err := GetConf()
	if err != nil {
		return 0, err
	}
	return c.Float(head,key)
}


