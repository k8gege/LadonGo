package nbt
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"os"
	"strings"
)

func PrintVersion(app string) {
	var version = "master"
	fmt.Fprintf(os.Stderr, "%s v%s\n", app, version)
}

func TrimName(name string) string {
	return strings.TrimSpace(strings.Replace(name, "\x00", "", -1))
}
