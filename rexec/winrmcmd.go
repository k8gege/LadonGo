package rexec
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
  "github.com/masterzen/winrm"
  "fmt"
  "os"
  "strconv"
)

var WinrmHelp = func () {
    fmt.Println("Usage: Ladon WinrmCmd host port user pass cmd")
}

func WinrmCmd(host,port,user,pass ,cmd string){
	iport,err := strconv.Atoi(port)
	endpoint := winrm.NewEndpoint(host, iport, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, user, pass)
	if err != nil {
		panic(err)
	}
	client.Run(cmd, os.Stdout, os.Stderr)
}

// func main() {
	// args := os.Args
    // if len(args) < 5 || args == nil {
		// help()
        // return
    // }
	// host := args[1]
	// port,err := strconv.Atoi(args[2])
	// user := args[3]
	// pass := args[4]
	// cmd := args[5]
	
	// //WinrmAuth(host,user,pass,port)
	
	// //endpoint := winrm.NewEndpoint("192.168.1.116", 5985, false, false, nil, nil, nil, 0)
	// //client, err := winrm.NewClient(endpoint, "k8gege", "k8test520!@#")
	// endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)
	// client, err := winrm.NewClient(endpoint, user, pass)
	// if err != nil {
		// panic(err)
	// }
	// client.Run(cmd, os.Stdout, os.Stderr)
// }