package port
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"net"
	"fmt"
	"log"
	"sync"
	"time"
	//"io/ioutil"
	"bufio"
	"github.com/k8gege/LadonGo/tcp"
)
func IsBanner(address string)(string, error) {

    conn, err := net.DialTimeout("tcp", address, time.Second*10)
    if err != nil {
        return "",err
    }
    defer conn.Close()
    tcpconn := conn.(*net.TCPConn)
    tcpconn.SetReadDeadline(time.Now().Add(time.Second * 5))
    reader := bufio.NewReader(conn)
    return reader.ReadString('\n')
}

func CheckPort(ip net.IP, port int) {
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if conn !=nil{
		fmt.Println(tcpAddr.IP,tcpAddr.Port,"Open")
		conn.Close()
	}
	if err != nil {
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
	//	fmt.Println(err)
	}
}

func PortCheck(host string, port int)(result bool) {
	result = false
	ip := net.ParseIP(host)
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if conn !=nil{
		fmt.Println(tcpAddr.IP,tcpAddr.Port,"Open")
		conn.Close()
		result = true
	}
	if err != nil {
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
	//	fmt.Println(err)
	}
	return result
}

func PortIsOpen(ip net.IP, port int) ( result bool,err error) {
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if conn !=nil{
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Open")
		conn.Close()
		result = true
	}
	if err != nil {
		//fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
	//	fmt.Println(err)
	}
	return result,err
}

type Workdist struct {
	Host	string
}

const (
	taskload		    = 255
	tasknum			= 255
)
var wg sync.WaitGroup

func TaskPort(ip string,debugLog *log.Logger){
	tasks := make(chan Workdist,taskload)
	wg.Add(tasknum)

	for gr:=1;gr<=tasknum;gr++ {
		go workerPort(tasks,debugLog)
	}

	for i:=1;i<256;i++ {
		host:=fmt.Sprintf("%s.%d",ip,i)
		task := Workdist{
			Host:host,
		}
		tasks <- task
	}
	close(tasks)
	wg.Wait()
}

func workerPort(tasks chan Workdist,debugLog *log.Logger){
	defer wg.Done()
	task,ok := <- tasks
	if !ok {
		return
	}
	host := task.Host

	//Default
	ScanPort(host)

}

var DefaultPorts = []int{21,22,23,25,80,443,8080,110,135,139,445,389,489,587,1433,1434,1521,1522,1723,2121,3306,3389,4899,5631,5632,5800,5900,7071,43958,65500,4444,8888,6789,4848,5985,5986,8081,8089,8443,10000,6379,7001,7002}
      
func ScanPort(host string){
var wg sync.WaitGroup
for _, p:= range DefaultPorts {
wg.Add(1)
//CheckPort(net.ParseIP(host),p)
tcp.PortCheck(host,p)
defer wg.Done()
}
wg.Wait()
}

func ScanPortBanner(host string){	
for _, p:= range DefaultPorts {
tcp.GetBanner(host,p)
}
}