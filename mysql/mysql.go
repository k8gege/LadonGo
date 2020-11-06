package mysql
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
import (
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"fmt"
	"database/sql"
	_"github.com/Go-SQL-Driver/MySQL"
)

func MysqlAuth(ip string, port string, user string, pass string) ( result bool,err error) {
	result = false
    db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+":"+port+")/mysql?charset=utf8")
    if err != nil {
    }
	if db.Ping()==nil {
		result = true
	}
	return result,err
}

func MysqlScan(ScanType string,Target string) {
	if port.PortCheck(Target,3306) {
		Loop:
		for _, u := range dic.UserDic() {
			for _, p := range dic.PassDic() {
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := MysqlAuth(Target, "3306", u, p)
				if res==true && err==nil {
					logger.PrintIsok(ScanType,Target,u, p)
					break Loop
				}
			}
		}
	}
}