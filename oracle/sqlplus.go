package oracle
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"os/exec"
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"strings"
)

func SqlplusAuth(host,port,user,pass string) (bool){
cmd := exec.Command("sqlplus", "-s",  user+"/"+pass+"@"+host+":"+port+"/orcl")
output,err := cmd.Output()
if err != nil {
	panic(err)
	return false
}
if string(output)=="" {
	return true
}
return false
}
	
func SqlPlusScan2(ScanType string,Target string) {
	if port.PortCheck(Target,1521) {
		Loop:
		for _, u := range dic.UserDic() {
			for _, p := range dic.PassDic() {
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res := SqlplusAuth(Target, "1521", u, p)
				if res {
					logger.PrintIsok2(ScanType,Target,"1521",u, p)
					break Loop
				}
			}
		}
	}
}

func SqlPlusScan(ScanType string,Target string) {
	if port.PortCheck(Target,1521) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res := SqlplusAuth(Target, "1521", u, p)
				if res {
					logger.PrintIsok2(ScanType,Target,"1521",u, p)
					break Loop
				}
				
			}
		} else {
			SqlPlusScan(ScanType,Target)	
		}
	}
}
