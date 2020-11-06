package main
import (
	"log"
	"github.com/k8gege/LadonGo/smb"
)
func main() {
log.Println(smb.SmbAuth("192.168.1.21", "445", "k8gege", "k8gege520"))
}
