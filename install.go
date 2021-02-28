package main
//Usage: go run install.go
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
"fmt"
"runtime"
"os/exec"
"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
fmt.Println("Install to System")
pwd, _ := os.Getwd()
LadonPath:=""
//fmt.Println(pwd)
if runtime.GOOS!="windows" {
LadonPath="/usr/local/bin/Ladon"
cmd :=exec.Command("/bin/sh","-c","go build -o "+LadonPath+" "+pwd+"/Ladon.go")
cmd.Run()
} else if runtime.GOOS=="windows" {
LadonPath="C:\\Windows\\System32\\Ladon.exe"
cmd :=exec.Command("cmd","/c","go build -o "+LadonPath+" "+pwd+"\\Ladon.go")
cmd.Run()
}
exist, _ := PathExists(LadonPath)
if exist{
	fmt.Println("Done!")
}else{
	fmt.Println("Fail!")
	fmt.Println("Admin is needed")
}


}