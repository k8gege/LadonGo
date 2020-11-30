package tcp
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "net"
    "sync"
    "time"
	"strconv"
)

func PortCheck(host string, port int)(bool) {
    p := strconv.Itoa(port)
    conn, err := net.DialTimeout("tcp", host+":"+p, time.Second*5)
    if err != nil {
        //fmt.Println(host,p,"Close")
	return false
    } else {
	fmt.Println(host,p,"Open")
	conn.Close()
	return false
	}
}

// 假定是 SSH 服务。
// 返回 banner 第一行。
func assume_ssh(address string) (string, error) {
    conn, err := net.DialTimeout("tcp", address, time.Second*10)
    if err != nil {
        return "", err
    }
    defer conn.Close()
    tcpconn := conn.(*net.TCPConn)
    tcpconn.SetReadDeadline(time.Now().Add(time.Second * 5))
    reader := bufio.NewReader(conn)
    return reader.ReadString('\n')
}

func split_http_head(data []byte, atEOF bool) (advance int, token []byte, err error) {
    head_end := bytes.Index(data, []byte("\r\n\r\n"))
    if head_end == -1 {
        return 0, nil, nil
    }
    return head_end + 4, data[:head_end+4], nil
}

// 假定是 HTTP 服务。
// 返回 "/" HTTP 返回头。
func assume_http(address string) (string, error) {
    conn, err := net.DialTimeout("tcp", address, time.Second*10)
    if err != nil {
        return "", err
    }
    defer conn.Close()
    tcpconn := conn.(*net.TCPConn)
    tcpconn.SetWriteDeadline(time.Now().Add(time.Second * 5))
    if _, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")); err != nil {
        return "", err
    }
    tcpconn.SetReadDeadline(time.Now().Add(time.Second * 5))
    scanner := bufio.NewScanner(conn)
    scanner.Split(split_http_head)
    if scanner.Scan() {
        return scanner.Text(), nil
    }
    err = scanner.Err()
    if err == nil {
        err = io.EOF
    }
    return "", err
}

func GetBanner(address string,port int) {
	if PortCheck(address,port) {
		p := strconv.Itoa(port)
		result := make(chan string, 2)
		done := make(chan int, 1)
		var g sync.WaitGroup
		g.Add(2)
		go func() {
			if r, e := assume_ssh(address+":"+p); e == nil {
				result <- fmt.Sprintf("SSH: %s", r)
			}
			g.Done()
		}()
		go func() {
			if r, e := assume_http(address+":"+p); e == nil {
				result <- fmt.Sprintf("HTTP: %s", r)
			}
			g.Done()
		}()
		go func() {
			g.Wait()
			done <- 1
		}()
		select {
		case <-done:
			//fmt.Printf("%s\n%s\tnull\n", address,p)			
		case r := <-result:
			fmt.Printf("%s:%s\t%s", address,p, r)
		}
	}
}
