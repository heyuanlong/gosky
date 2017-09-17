package main

import(
	"fmt"
	"net"
	"time"
	"os"

	gnet "github.com/heyuanlong/gosky/net"
)

/*


*/

func main() {
	var buf[512]byte

	service := "127.0.0.1:8089"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",service)
	checkErr(err)
	conn,err := net.DialTCP("tcp",nil ,tcpAddr)
	defer conn.Close()
	checkErr(err)

{
	t1 := time.Now() // get current time
	for i := 0; i < 4000; i++ {
		data := gnet.SetPackage(10,[]byte("Hello servie111111111111111111111111111111111111111111"))
		_ ,err := conn.Write(data)
		checkErr(err)

		n, err :=conn.Read(buf[0:])
		gnet.ParsePackage(buf[:n])
		//fmt.Print(string(msg))
	}

  	elapsed := time.Since(t1)
    fmt.Println("App elapsed: ", elapsed)
}
{
	t1 := time.Now() // get current time
	for i := 0; i < 4000; i++ {
		data := gnet.SetPackage(11,[]byte("Hello servie111111111111111111111111111111111111111111"))
		_ ,err := conn.Write(data)
		checkErr(err)

		n, err :=conn.Read(buf[0:])
		gnet.ParsePackage(buf[:n])
		//fmt.Print(string(msg))
	}

  	elapsed := time.Since(t1)
    fmt.Println("App elapsed: ", elapsed)
}

}

func checkErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
