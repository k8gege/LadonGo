package http
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
import (
	"fmt"
	"net/http"
	//"os"
	"strings"
)

func IsUrl(url string) ( result string)  {
	if !strings.Contains(url, "http") {
		url := "http://"+url
		return url
	}

	return url
}
func HttpBanner(url string) ( result bool,err error)  {

	url2:=IsUrl(url)
	response, err := http.Head(url2) 
	if err != nil {
		//fmt.Println(err.Error())
		//os.Exit(2)
		return false,err
	}
	
	//fmt.Println(response)
	//fmt.Println(response.Status)
	for k, v := range response.Header { 
		//fmt.Println(k+":", v)
		if k=="Server" {
			fmt.Println(url2, v)
			}
	}
	
	return result,err
 
}

// func main() {

// HttpBanner("http://www.baidu.com")
// HttpBanner("192.168.1.1")

// }