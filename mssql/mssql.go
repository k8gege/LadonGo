package mssql
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
	_"github.com/denisenkom/go-mssqldb"
	"strings"
)

func MssqlAuth(ip ,port,user ,pass string) ( result bool,err error) {
	result = false
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;encrypt=disable", ip, user, pass, port)
	db, err := sql.Open("mssql", connString)
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result = true
		 }
	}
	
	return result,err
}

func MssqlScan2(ScanType string,Target string) {
	if port.PortCheck(Target,1433) {
		Loop:
		for _, u := range dic.UserDic() {
			for _, p := range dic.PassDic() {
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := MssqlAuth(Target, "1433", u, p)
				if res==true && err==nil {
					logger.PrintIsok2(ScanType,Target,"1433",u, p)
					break Loop
				}
			}
		}
	}
}

func MssqlScan(ScanType string,Target string) {
	if port.PortCheck(Target,1433) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := MssqlAuth(Target, "1433", u, p)
				if res==true && err==nil {
					logger.PrintIsok2(ScanType,Target,"1433",u, p)
					break Loop
				}
				
			}
		} else {
			MssqlScan2(ScanType,Target)	
		}
	}
}
