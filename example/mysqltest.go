package main
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/k8gege/LadonGo/mysql"
	"fmt"
)
func main() {
fmt.Println(mysql.MysqlAuth("192.168.1.21","3306","root","k8gege"))
}
