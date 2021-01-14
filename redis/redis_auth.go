package redis
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
    //"fmt"
 	"strconv"
    "github.com/monnand/goredis"
    "github.com/k8gege/LadonGo/port"
    "github.com/k8gege/LadonGo/logger"
)
 
func RedisNullAuth(host string,iport int) (result bool) {
     result = false
if port.PortCheck(host,iport) {
    var client goredis.Client
    port:=strconv.Itoa(iport)
    client.Addr = host+":"+port
    err := client.Set("test", []byte("ISOK"))
    if err != nil {
        //panic(err)
    }
 
    res, _ := client.Get("test")
	if string(res)=="ISOK" {	
		result = true
	}

    client.Set("test", []byte("test"))
}
return result
}

func RedisNullScan(ScanType,host string) {
if RedisNullAuth(host,6379) {
	logger.PrintIsok0(ScanType,host,"6379")
}

}