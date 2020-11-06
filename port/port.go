package port
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
import (
	"net"
	"fmt"
	"log"
	"sync"
)

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
	//	fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
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
	//	fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
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
	//	fmt.Println(tcpAddr.IP,tcpAddr.Port,"Close")
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
	//fmt.Println("Ping: "+host)
	
	//Default
	ScanPort(host)
	
	
	// res,err := CheckPort(host)
	// if err==nil && res==true {
		// fmt.Println("PING: "+host)
	// }
}

func ScanPort(host string){

	//Default
	CheckPort(net.ParseIP(host),21)
	CheckPort(net.ParseIP(host),22)
	CheckPort(net.ParseIP(host),135)
	CheckPort(net.ParseIP(host),139)
	CheckPort(net.ParseIP(host),445)
	CheckPort(net.ParseIP(host),1433)
	CheckPort(net.ParseIP(host),3306)
	CheckPort(net.ParseIP(host),1521)
	CheckPort(net.ParseIP(host),6379)
	CheckPort(net.ParseIP(host),5800)
	CheckPort(net.ParseIP(host),5900)
	CheckPort(net.ParseIP(host),3389)
	CheckPort(net.ParseIP(host),5985)
	
	CheckPort(net.ParseIP(host),80)
	CheckPort(net.ParseIP(host),81)
	CheckPort(net.ParseIP(host),443)
	CheckPort(net.ParseIP(host),7001)
	CheckPort(net.ParseIP(host),7002)
	CheckPort(net.ParseIP(host),8080)
	CheckPort(net.ParseIP(host),8081)
	CheckPort(net.ParseIP(host),8089)
	CheckPort(net.ParseIP(host),8443)
	CheckPort(net.ParseIP(host),10000)
	
	
	// res,err := CheckPort(host)
	// if err==nil && res==true {
		// fmt.Println("PING: "+host)
	// }
}