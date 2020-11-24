package main
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"log"
	"github.com/k8gege/LadonGo/ssh"
)
func main() {
	log.Println(ssh.SshAuth("192.168.1.19", "22", "root", "k8team"))
	log.Println(ssh.SshAuth("192.168.1.19", "22", "root", "k8gege"))
}
