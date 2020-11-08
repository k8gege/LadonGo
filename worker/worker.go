package worker
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"sync"
	"fmt"
	"log"
	"github.com/k8gege/LadonGo/icmp"
	"github.com/k8gege/LadonGo/ping"
)

type Workdist struct {
	Host	string
}

const (
	taskload		    = 255
	tasknum			= 255
)
var wg sync.WaitGroup

func TaskPing(ip string,debugLog *log.Logger){
	tasks := make(chan Workdist,taskload)
	wg.Add(tasknum)

	for gr:=1;gr<=tasknum;gr++ {
		go workerPing(tasks,debugLog)
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

func workerPing(tasks chan Workdist,debugLog *log.Logger){
	defer wg.Done()
	task,ok := <- tasks
	if !ok {
		return
	}
	host := task.Host
	//fmt.Println("Ping: "+host)
	res,err := ping.CmdPing(host)
	if err==nil && res==true {
		fmt.Println("PING: "+host)
	}
}

func TaskIcmp(ip string,debugLog *log.Logger){
	tasks := make(chan Workdist,taskload)
	wg.Add(tasknum)

	for gr:=1;gr<=tasknum;gr++ {
		go workericmp(tasks,debugLog)
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
func workericmp(tasks chan Workdist,debugLog *log.Logger){
	defer wg.Done()
	task,ok := <- tasks
	if !ok {
		return
	}
	host := task.Host
	icmp.Icmp(host,debugLog)
}

