package logger
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	//"fmt"
	"log"
	"os"
	//"runtime"
	"github.com/fatih/color"
)

func PrintIsok(ScanType,h ,u ,p string){
		//if runtime.GOOS=="windows" {
			//fmt.Println("Found: "+h+" "+u+" "+p+" ISOK")
		//} else
		//{fmt.Println("\033[35mFound: "+h+" "+u+" "+p+" ISOK\033[0m")}
		color.Magenta("Found: "+h+" "+u+" "+p+" ISOK")
		logFile, err := os.OpenFile(ScanType+".Log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		log.SetOutput(logFile)
		log.Println("Found: "+h+" "+u+" "+p+" ISOK")
}

func PrintIsok2(ScanType,h ,port,u ,p string){
	//if runtime.GOOS=="windows" {
		//fmt.Println("Found: "+h+" "+port+" "+u+" "+p+" ISOK")
	//} else
	//{fmt.Println("\033[35mFound: "+h+" "+port+" "+u+" "+p+" ISOK\033[0m")}
	color.Magenta("Found: "+h+" "+u+" "+p+" ISOK")
		logFile, err := os.OpenFile(ScanType+".Log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		log.SetOutput(logFile)
		log.Println("Found: "+h+" "+port+" "+u+" "+p+" ISOK")
}