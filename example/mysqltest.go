package main
import (
	"github.com/k8gege/LadonGo/mysql"
	"fmt"
)
func main() {
fmt.Println(mysql.MysqlAuth("192.168.1.21","3306","root","k8gege"))
}
