package main
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/k8gege/LadonGo/routeros"
	"fmt"
)
func main() {
fmt.Println(routeros.RouterOSAuth("192.168.250.110","8728","admin","admin"))
}
