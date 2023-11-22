package mysql
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"strings"
)

func MysqlAuth(ip string, port string, user string, pass string) ( result bool,err error) {
	result = false
    db, err := sql.Open("mysql", user+":"+pass+"@tcp("+ip+":"+port+")/mysql?charset=utf8")
    if err != nil {
	    return result, err
    }
	if db.Ping()==nil {
		result = true
	}
	return result,err
}

func MysqlScan2(ScanType string,Target string) {

	Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			fmt.Println("Check... "+Target+" "+u+" "+p)
			res,err := MysqlAuth(Target, "3306", u, p)
			if res==true && err==nil {
				logger.PrintIsok2(ScanType,Target,"3306",u, p)
				break Loop
			}
		}
	}

}

func MysqlScan(ScanType string,Target string) {
	if port.PortCheck(Target,3306) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := MysqlAuth(Target, "3306", u, p)
				if res==true && err==nil {
					logger.PrintIsok2(ScanType,Target,"3306",u, p)
					break Loop
				}
				
			}
		} else {
			MysqlScan2(ScanType,Target)	
		}
	}
}
