package t3
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	//"flag"
	//"github.com/k8gege/LadonGo/worker"
	"fmt"
	"net"
	//"net/url"
	//"os"
	"sync"
	"log"
	"time"
	"strings"
)

type Workdist struct {
	Host	string
}

const (
	taskload		    = 255
	tasknum			= 255
)
var wg sync.WaitGroup

var t3header = []byte{0x74, 0x33, 0x20, 0x31, 0x32, 0x2e, 0x32, 0x2e, 0x31, 0x0a, 0x41, 0x53, 0x3a, 0x32, 0x35, 0x35, 0x0a, 0x48, 0x4c, 0x3a, 0x31, 0x39, 0x0a, 0x4d, 0x53, 0x3a, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x0a, 0x50, 0x55, 0x3a, 0x74, 0x33, 0x3a, 0x2f, 0x2f, 0x75, 0x73, 0x2d, 0x6c, 0x2d, 0x62, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x3a, 0x37, 0x30, 0x30, 0x31, 0x0a, 0x0a}
var payload []byte

// func main() {

	// u := flag.String("u", "", "192.168.1.1:7001")
	// flag.Parse()
	// // if  *u == "" {
		// // fmt.Println(`Using --help`)
		// // os.Exit(1)
	// // }

	// //fmt.Println("u: "+ *u)
	// //fmt.Println("start send payload1.....")

	// err := T3Connect( *u)
	// if err != nil {
		// println(err.Error())
	// }

// }

func T3version(host string) {
	
	// err := T3Connect(host+":7001")
	// if err != nil {
		// //println(err.Error())
	// }
	T3Connect(host+":7001")
	T3Connect(host+":7002")
	T3Connect(host)
	T3Connect(host+":8080")
}

func Int8byte(y int) []byte {
	var x uint16
	x = uint16(y)
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, x)
	return b_buf.Bytes()
}

func Int32byte(x int) []byte {
	var y uint32
	y = uint32(x)
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, y)
	return b_buf.Bytes()
}

func T3Connect(host string) error {

	conn, err := net.DialTimeout("tcp", host, 5*time.Second)
	if err != nil {
		//fmt.Println("Port Close")
		return err
	}
	_, err = conn.Write(t3header)
	if err != nil {
		//fmt.Println("send err")
		return err
	}
	ReadAll(conn,host)
	return nil
}

func Appendbyte(by1 *[]byte, by2 ...[]byte) {
	for _, b := range by2 {
		*by1 = append(*by1, b...)
	}
}

func ReadAll(conn net.Conn,host string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	ch := make(chan struct{})
	b := make([]byte, 1024)
	reader := bufio.NewReader(conn)

	go func(ch chan struct{}) {

		for {
			n, err := reader.Read(b)
			if err != nil {
				break
			}
			//fmt.Printf("%s\n", b[:n])
			// if string(b[:n])!="HELO"{
				// fmt.Printf("%s\n", b[:n])
				// //result := string(b[:n])
			// }
			res := string(b[:n])
			ver := GetBetweenStr(res,":",".false")
			if strings.Contains(res,".false") {
				fmt.Printf(host+"\tWeblogic%s\n", ver)
				
			}
		}
		ch <- struct{}{}
	}(ch)
	select {
	case <-ch:
	case <-ctx.Done():
		break
	}
	cancel()
}

func GetBetweenStr(str, start, end string) string {
 n := strings.Index(str, start)
 if n == -1 {
 n = 0
 }
 str = string([]byte(str)[n:])
 m := strings.Index(str, end)
 if m == -1 {
 m = len(str)
 }
 str = string([]byte(str)[:m])
 return str
}

func TaskT3(ip string,debugLog *log.Logger){
	tasks := make(chan Workdist,taskload)
	wg.Add(tasknum)

	for gr:=1;gr<=tasknum;gr++ {
		go workert3(tasks,debugLog)
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
func workert3(tasks chan Workdist,debugLog *log.Logger){
	defer wg.Done()
	task,ok := <- tasks
	if !ok {
		return
	}
	host := task.Host
	T3version(host)
}