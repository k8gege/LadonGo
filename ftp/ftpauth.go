package ftp
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/k8gege/LadonGo/goftp"
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"fmt"
	"strings"
)

func FtpAuth(ip string, port string, user string, pass string) ( result bool,err error) {
	result = false

//var err error
var Lftp *goftp.FTP

if Lftp, err = goftp.Connect(ip+":"+port); err != nil {
    //fmt.Println(err)
}

defer Lftp.Close()

if err = Lftp.Login(user,pass); err == nil {
	result = true
}
	return result,err
}

func FtpScan2(ScanType string,Target string) {
	Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			fmt.Println("Check... "+Target+" "+u+" "+p)
			res,err := FtpAuth(Target, "21", u, p)
			if res==true && err==nil {
				logger.PrintIsok2(ScanType,Target,"21",u, p)
				break Loop
			}
		}
	}
}

func FtpScan(ScanType string,Target string) {
	if port.PortCheck(Target,21) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := FtpAuth(Target, "21", u, p)
				if res==true && err==nil {
					logger.PrintIsok2(ScanType,Target,"21",u, p)
					break Loop
				}
				
			}
		} else {
			FtpScan2(ScanType,Target)	
		}
	}
}