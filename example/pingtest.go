package main
import (
	"fmt"
	"flag"
	"os"
	"github.com/k8gege/LadonGo/ping"
	)
//go run pingtest.go -host k8gege.org
var (
	host string
)

func init() {
	flag.StringVar(&host, "host", "", "IP/Host/Domain")
}

func main() {
flag.Parse()
if host == "" {
	println("Please " + os.Args[0] + " -h")
	os.Exit(0)
}

res,err := ping.CmdPing(host)
//res,err := CmdPing("k8gege.org")
if err==nil && res==true {
fmt.Println(host+" IsOnline")
}

}