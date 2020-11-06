package http
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
import (
	"sync"
	"fmt"
	"log"

)

type Workdist struct {
	Host	string
}

const (
	taskload		    = 255
	tasknum			= 255
)
var wg sync.WaitGroup

func TaskHttp(ip string,debugLog *log.Logger){
	tasks := make(chan Workdist,taskload)
	wg.Add(tasknum)

	for gr:=1;gr<=tasknum;gr++ {
		go workerhttp(tasks,debugLog)
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
func workerhttp(tasks chan Workdist,debugLog *log.Logger){
	defer wg.Done()
	task,ok := <- tasks
	if !ok {
		return
	}
	host := task.Host
	HttpBanner(host)
}

