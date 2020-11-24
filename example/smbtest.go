package main
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"log"
	"github.com/k8gege/LadonGo/smb"
)
func main() {
log.Println(smb.SmbAuth("192.168.1.21", "445", "k8gege", "k8gege520"))
}
