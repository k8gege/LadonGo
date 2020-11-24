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

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func TxtRead(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open ",filename,"error, %v", err)
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
	file, err := os.Open("user.txt")
	if err != nil {
		fmt.Println("Open user.txt error, %v", err)
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
	file, err := os.Open("pass.txt")
	if err != nil {
		fmt.Println("Open pass.txt error, %v", err)
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
	file, err := os.Open("userpass.txt")
	if err != nil {
		fmt.Println("Open userpass.txt error, %v", err)
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