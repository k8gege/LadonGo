package ftp
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
import (
	"github.com/k8gege/LadonGo/goftp"
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"fmt"
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

func FtpScan(ScanType string,Target string) {
	if port.PortCheck(Target,21) {
		Loop:
		for _, u := range dic.UserDic() {
			for _, p := range dic.PassDic() {
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := FtpAuth(Target, "21", u, p)
				if res==true && err==nil {
					logger.PrintIsok(ScanType,Target,u, p)
					break Loop
				}
			}
		}
	}
}