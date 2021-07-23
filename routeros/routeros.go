package routeros
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"strings"
	"github.com/go-routeros/routeros"
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
)

func RouterOSAuth(ip string, port string, user string, pass string) ( result bool,err error) {
	result = false
    _, err = routeros.Dial(ip+":"+port,user,pass)
    if err == nil {
	result = true
    }

	//defer c.Close()

	return result,err
}

func RouterOSScan2(ScanType string,Target string) {
	Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			fmt.Println("Check... "+Target+" "+u+" "+p)
			res,err := RouterOSAuth(Target, "8728", u, p)
			if res==true && err==nil {
				logger.PrintIsok2(ScanType,Target,"8728",u, p)
				break Loop
			}
		}
	}
}

func RouterOSScan(ScanType string,Target string) {
	if port.PortCheck(Target,8728) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := RouterOSAuth(Target, "8728", u, p)
				if res==true && err==nil {
					logger.PrintIsok2(ScanType,Target,"8728",u, p)
					break Loop
				}
				
			}
		} else {
			RouterOSScan2(ScanType,Target)	
		}
	}
}


