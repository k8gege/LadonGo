package dic
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	"bufio"
	"strings"
	//"log"
	"os"
)

func UserPassIsExist() bool{
if IsExist("userpass.txt") {
	return true
}
return false
}

func PwdIsExist() bool{
if IsExist("userpass.txt") {
	return true
}
if IsExist("user.txt") {
	return true
}
if IsExist("pass.txt") {
	return true
}
return false
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func TxtRead(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open ",filename,"error, %v", err)
	}
	fi,_:=os.Stat(filename)
	if fi.Size() ==0 {
	fmt.Println("Error: "+filename+" is null!")
	os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip != "" {
			lines = append(lines, ip)
		}
	}
	return lines
}
func UserDic() (users []string) {
	dicname:="user.txt"
	file, err := os.Open(dicname)
	if err != nil {
		fmt.Println("Open "+dicname+" error, %v", err)
	}
	fi,_:=os.Stat(dicname)
	if fi.Size() ==0 {
	fmt.Println("Error: "+dicname+" is null!")
	os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	return users
}

func PassDic() (password []string) {
	dicname:="pass.txt"
	file, err := os.Open(dicname)
	if err != nil {
		fmt.Println("Open "+dicname+" error, %v", err)
	}
	fi,_:=os.Stat(dicname)
	if fi.Size() ==0 {
	fmt.Println("Error: "+dicname+" is null!")
	os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			password = append(password, passwd)
		}
	}
	return password
}

func UserPassDic() (userpass []string) {
	dicname:="userpass.txt"
	file, err := os.Open(dicname)
	if err != nil {
		fmt.Println("Open "+dicname+" error, %v", err)
	}
	fi,_:=os.Stat(dicname)
	if fi.Size() ==0 {
	fmt.Println("Error: "+dicname+" is null!")
	os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			userpass = append(userpass, passwd)
		}
	}
	return userpass
}