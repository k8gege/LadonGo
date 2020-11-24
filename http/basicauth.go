package http
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	//_"compress/gzip"
	//"encoding/base64"
	"fmt"
	"net/http"
	"github.com/k8gege/LadonGo/logger"
	"github.com/k8gege/LadonGo/dic"
	"strings"
)

func IsBasicAuthURL(url string) (result bool) {
	result=false
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("","")
	resp, err := client.Do(req)
	if err == nil {
		resp.Body.Close()
		if resp.StatusCode == 401 {
			fmt.Println(url+" IS 401URL")
			result=true
		}
	}
	return result
}

func BasicAuth(ScanType,url,user,pass string) (result bool,err error) {
	result=false
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user,pass)
	resp, err := client.Do(req)
	if err == nil {
		resp.Body.Close()
		if resp.StatusCode != 401 {
				//logger.PrintIsok(ScanType,url,user, pass)
				return true,err
			} else {
				//fmt.Println("err")
			}
	}
	return result,err
}

func BasicAuthScan2(ScanType string,Target string) {
	if IsBasicAuthURL(Target) {
		Loop:
		for _, u := range dic.UserDic() {
			for _, p := range dic.PassDic() {
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := BasicAuth(ScanType,Target, u, p)
				if res==true && err==nil {
					logger.PrintIsok(ScanType,Target,u, p)
					break Loop
				}
			}
		}
	}
}

func BasicAuthScan(ScanType string,Target string) {
	if IsBasicAuthURL(Target) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := BasicAuth(ScanType,Target, u, p)
				if res==true && err==nil {
					logger.PrintIsok(ScanType,Target,u, p)
					break Loop
				}	
			}
				
		} else {
			BasicAuthScan2(ScanType,Target)	
		}
	
	}
}
