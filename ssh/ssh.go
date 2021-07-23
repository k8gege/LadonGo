package ssh
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/dic"
	"github.com/k8gege/LadonGo/logger"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"strings"
	//"log"
)

var SshHelp = func () {
    fmt.Println("Usage: Ladon SshCmd host port user pass cmd")
}

type Node struct {
    User     string
    Password string
}

func NewNode(user, password string) *Node {
    node := new(Node)
    node.User = user
    node.Password = password
    return node
}

func (this *Node) Conn(addr string,port string) (*ssh.Client, error) {

    authMethods := []ssh.AuthMethod{}

    keyboardInteractiveChallenge := func(
        user,
        instruction string,
        questions []string,
        echos []bool,
    ) (answers []string, err error) {
        if len(questions) == 0 {
            return []string{}, nil
        }
        return []string{this.Password}, nil
    }

    authMethods = append(authMethods, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
    authMethods = append(authMethods, ssh.Password(this.Password))

    sshConfig := &ssh.ClientConfig{
        User: this.User,
        Auth: authMethods,
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
		},
    }

    client, err := ssh.Dial("tcp", fmt.Sprintf("%s:"+port, addr), sshConfig)

    if err != nil {
        //return nil, fmt.Errorf(fmt.Sprintf("Connect fail \n(%s)", addr, err.Error()))
    }

    return client, nil
}

func ExecCmd(host string, port string, user string, pass string, cmdline string) {

    node := NewNode(user, pass)
    client, err := node.Conn(host,port)
	if err == nil{
			defer client.Close()
		    session, err := client.NewSession()
			if err != nil {
				fmt.Println("Create ssh session fail!",err)
			}
			defer session.Close()
			combo,err := session.CombinedOutput(cmdline)
			if err != nil {
				fmt.Println("ExecCmd fail!",err)
			}
			fmt.Println(string(combo))
	}
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
	

}


func SshCheck(host string, port string, user string, pass string) ( result bool,err error) {
	result = false

    node := NewNode(user, pass)
    client, err := node.Conn(host,port)
	if err == nil{
		//fmt.Println(host+" "+port+" "+user+" "+pass+" ISOK")
			//defer client.Close()
		    session, err := client.NewSession()
			if err != nil {
				//fmt.Println("Create ssh session fail!",err)
			}
			defer session.Close()
			combo,err := session.CombinedOutput("echo ISOK;")
			//combo,err := session.CombinedOutput(cmdline)
			if err != nil {
				//fmt.Println("ExecCmd fail!",err)
			}
			//log.Println("Result:",string(combo))
			//fmt.Println(string(combo))
			if strings.Contains(string(combo),"ISOK") {
				//fmt.Println("Found: "+host+" "+port+" "+user+" "+pass+" ISOK")
				result = true
			}
	}
    if err != nil {
        //fmt.Println(err)
    }
    defer client.Close()
	
	return result,err
}

func SshAuth(host string, port string, user string, pass string) (result bool,err error) {
	result=false
    authMethods := []ssh.AuthMethod{}

    keyboardInteractiveChallenge := func(
        user,
        instruction string,
        questions []string,
        echos []bool,
    ) (answers []string, err error) {
        if len(questions) == 0 {
            return []string{}, nil
        }
        return []string{pass}, nil
    }

    authMethods = append(authMethods, ssh.KeyboardInteractive(keyboardInteractiveChallenge))
    authMethods = append(authMethods, ssh.Password(pass))

    sshConfig := &ssh.ClientConfig{
        User: user,
        Auth: authMethods,
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
		},
    }

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", host, port), sshConfig)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		errRet := session.Run("echo ISOK")
		if err == nil && errRet == nil {
			defer session.Close()
			result = true
		}
	}
	return result,err
}

func SshScan2(ScanType string,Target string) {

	Loop:
	for _, u := range dic.UserDic() {
		for _, p := range dic.PassDic() {
			fmt.Println("Check... "+Target+" "+u+" "+p)
			res,err := SshAuth(Target, "22", u, p)
			if res==true && err==nil {
				logger.PrintIsok2(ScanType,Target,"22",u, p)
				break Loop
			}
		}
	}

}

func SshScan(ScanType string,Target string) {
	if port.PortCheck(Target,22) {
		if dic.UserPassIsExist() {
			Loop:
			for _, up := range dic.UserPassDic() {
				s :=strings.Split(up, " ")
				u := s[0]
				p := s[1]
				fmt.Println("Check... "+Target+" "+u+" "+p)
				res,err := SshAuth(Target, "22", u, p)
				if res==true && err==nil {
					logger.PrintIsok2(ScanType,Target,"22",u, p)
					break Loop
				}
				
			}
		} else {
			SshScan2(ScanType,Target)	
		}
	}
}
