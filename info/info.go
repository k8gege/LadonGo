package info
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"os/exec"
	//"log"
	"fmt"
	"bytes"
	"runtime"
	"strings"
	//"github.com/k8gege/LadonGo/str"
)

func OSver()  {
    sysType := runtime.GOOS
    if sysType == "linux" {
		cmd := exec.Command("/bin/sh","-c","uname -a")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		//result = true
		fmt.Println(out.String())
    } else if sysType == "windows" {
		cmd := exec.Command("cmd","/c","ver")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		fmt.Println(strings.Replace(out.String(), "\n", "", -1))
    }
    
    //return result
}

func PingOK(host string) ( result bool)  {
    sysType := runtime.GOOS
    if sysType == "linux" {
		cmd := exec.Command("/bin/sh","-c","ping -c 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		if strings.Contains(out.String(), "ttl=") {
			//fmt.Println("ISOK")
			result = true
		 } 
    } else if sysType == "windows" {
		cmd := exec.Command("cmd","/c","ping -n 1 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		if strings.Contains(out.String(), "TTL=") {
			//fmt.Println("ISOK")
			result = true
		 } 
    }
    return result
}

