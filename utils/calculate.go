package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
	"net/url"
	"io/ioutil"
)

func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

func GetRandomString(lens int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//这方式比较特别，按照123456来记忆吧：01月02号 下午3点04分05秒 2006年
func GetTimesString( ) string{
	return time.Now().Format("2006-01-02 15:04:05")
}
func GetTimes() int64  {
	return time.Now().Unix()
}
func GetTimesNano() int64  {
	return time.Now().UnixNano()
}

//urldecode
func Urldecode(s string) (string){
	ss,_ := url.QueryUnescape(s)
	return ss
}
//urlencode
func Urlencode(s string) string{
	return url.QueryEscape(s)
}

func ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
