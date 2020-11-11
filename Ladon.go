package main
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"fmt"
	//"github.com/k8gege/LadonGo/worker"
	//"github.com/k8gege/LadonGo/color" //Only Windows
	"github.com/k8gege/LadonGo/t3"
	"github.com/k8gege/LadonGo/icmp"
	"github.com/k8gege/LadonGo/ping"
	"github.com/k8gege/LadonGo/port"
	"github.com/k8gege/LadonGo/http"
	"github.com/k8gege/LadonGo/smb"
	"github.com/k8gege/LadonGo/ftp"
	"github.com/k8gege/LadonGo/ssh"
	"github.com/k8gege/LadonGo/mysql"
	"github.com/k8gege/LadonGo/winrm"
	"github.com/k8gege/LadonGo/rexec"
	"strings"
	"log"
	"time"
	"os"
	"runtime"
	"net"
	"sync"
)

func help() {
	if runtime.GOOS=="windows" {fmt.Println("\nHelp:")
	} else{fmt.Println("\033[32m\nHelp:\033[0m")}
	if runtime.GOOS=="windows" {
		//fmt.Println("Ladon Help")
		fmt.Println("Ladon FuncList")
		fmt.Println("Ladon Detection")
		fmt.Println("Ladon VulDetection")
		fmt.Println("Ladon BruteFor")
		fmt.Println("Ladon RemoteExec")
		fmt.Println("Ladon Example")
	} else{
		//fmt.Println("./Ladon Help")
		fmt.Println("./Ladon FuncList")
		fmt.Println("./Ladon Detection")
		fmt.Println("./Ladon VulDetection")
		fmt.Println("./Ladon BruteFor")
		fmt.Println("./Ladon RemoteExec")
		fmt.Println("./Ladon Example")
	}
}

func FuncList() {
	//help()
	Detection()
	VulDetection()
	BruteFor()
	RemoteExec()
	Example()
}

func Example() {
	if runtime.GOOS=="windows" {fmt.Println("\nExample:")
	} else{fmt.Println("\033[32m\nExample:\033[0m")}
	if runtime.GOOS=="windows" {
		fmt.Println("Ladon 192.168.1.8/24 MS17010")
		fmt.Println("Ladon 192.168.1/c MS17010")
		fmt.Println("Ladon 192.168/b MS17010")
		fmt.Println("Ladon 192/a MS17010")
	} else{
		fmt.Println("./Ladon 192.168.1.8/24 MS17010")
		fmt.Println("./Ladon 192.168.1/c MS17010")
		fmt.Println("./Ladon 192.168/b MS17010")
		fmt.Println("./Ladon 192/a MS17010")
	}
	fmt.Println("")
}

func Detection() {
	if runtime.GOOS=="windows" {fmt.Println("\nDetection:")
	} else{fmt.Println("\033[33m\nDetection:\033[0m")}
	fmt.Println("PingScan\t(Using system ping to detect Online hosts)")
	fmt.Println("IcmpScan\t(Using ICMP Protocol to detect Online hosts)")
	fmt.Println("HttpBanner\t(Using HTTP Protocol Scan Web Banner)")
	fmt.Println("HttpTitle\t(Using HTTP protocol Scan Web titles)")
	fmt.Println("T3Scan  \t(Using T3 Protocol Scan Weblogic hosts)")
	fmt.Println("PortScan\t(Scan hosts open ports using TCP protocol)")	
}

func VulDetection() {
	if runtime.GOOS=="windows" {fmt.Println("\nVulDetection:")
	} else{fmt.Println("\033[33m\nVulDetection:\033[0m")}
	fmt.Println("MS17010 \t(Using SMB Protocol to detect MS17010 hosts))")
	fmt.Println("SmbGhost\t(Using SMB Protocol to detect SmbGhost hosts))")
	
}

func BruteFor() {
	if runtime.GOOS=="windows" {fmt.Println("\nBruteForce:")
	} else{fmt.Println("\033[35m\nBruteForce:\033[0m")}
	fmt.Println("SmbScan \t(Using SMB Protocol to Brute-For 445 Port))")
	fmt.Println("SshScan \t(Using SSH Protocol to Brute-For 22 Port))")
	fmt.Println("FtpScan \t(Using FTP Protocol to Brute-For 21 Port))")
	fmt.Println("MysqlScan \t(Using Mysql Protocol to Brute-For 3306 Port))")
	fmt.Println("WinrmScan \t(Using Winrm Protocol to Brute-For 5985 Port))")
}

func RemoteExec() {
	if runtime.GOOS=="windows" {fmt.Println("\nRemoteExec:")
	} else{fmt.Println("\033[35m\nRemoteExec:\033[0m")}
	fmt.Println("SshCmd   \t(SSH Remote command execution Default 22 Port))")
	fmt.Println("WinrmCmd \t(Winrm Remote command execution Default 5985 Port))")
}


func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
var debugLog *log.Logger
func main() {
	fmt.Println("LadonGo 2.0 by k8gege")
	fmt.Println("Arch: "+runtime.GOARCH+" OS: "+runtime.GOOS)
	fmt.Println("")
	ParLen := len(os.Args)
	if ParLen==1 {
		help()
		os.Exit(0)
	}

	if ParLen==2 {
		SecPar := strings.ToUpper(os.Args[1])
		//fmt.Println(SecPar)
		if SecPar=="HELP"||SecPar=="/HELP"||SecPar=="-H"||SecPar=="-H" {
			help()
			os.Exit(0)
		}
		if SecPar=="HELPLIST"||SecPar=="FUNCLIST" {
			FuncList()
			os.Exit(0)
		}
		if SecPar=="BRUTEFOR"||SecPar=="BRUTE"||SecPar=="BRUTEFORCE" {
			BruteFor()
			os.Exit(0)
		}
		if SecPar=="DETECTION" {
			Detection()
			os.Exit(0)
		}
		if SecPar=="VULDETECTION" {
			VulDetection()
			os.Exit(0)
		}
		if SecPar=="EXAMPLE"||SecPar=="USAGE" {
			Example()
			os.Exit(0)
		}
		if SecPar == "WINRMCMD" || SecPar == "WINRMEXEC"|| SecPar == "SSHSHELL" {
			rexec.WinrmHelp()
			os.Exit(0)
		}
		if SecPar == "SSHCMD" || SecPar == "SSHEXEC" || SecPar == "SSHSHELL" {
			ssh.SshHelp()
			os.Exit(0)
		}
	}
	
	if ParLen>4 {
		SecPar := strings.ToUpper(os.Args[1])
		if SecPar == "WINRMCMD" || SecPar == "WINRMEXEC" || SecPar == "WINRMSHELL"{
			rexec.WinrmCmd(os.Args[2],os.Args[3],os.Args[4],os.Args[5],os.Args[6])
			os.Exit(0)
		}
		if SecPar == "SSHCMD" || SecPar == "SSHEXEC" || SecPar == "SSHSHELL" {
			ssh.ExecCmd(os.Args[2],os.Args[3],os.Args[4],os.Args[5],os.Args[6])
			os.Exit(0)
		}
	}
	
	EndPar := os.Args[ParLen-1]
	Target := os.Args[ParLen-2]

	fmt.Println("Targe: "+Target)
	fmt.Println("Load "+EndPar)
	//log.Println("Start...")
	fmt.Println("\nScanStart: "+time.Now().Format("2006-01-02 03:04:05"))	
	ScanType := strings.ToUpper(EndPar)
	if strings.Contains(Target, "/c")||strings.Contains(Target, "/C") {
		CScan(ScanType,Target)
	} else if strings.Contains(Target, "/b")||strings.Contains(Target, "/B") {
		BScan(ScanType,Target)
	} else if strings.Contains(Target, "/a")||strings.Contains(Target, "/A") {
		AScan(ScanType,Target)
	} else if strings.Contains(Target, "/") {
		if Target != ""  {
		ip, ipNet, err := net.ParseCIDR(Target)
		if err != nil {
			fmt.Println(Target +" invalid CIDR")
			return
		}
		var wg sync.WaitGroup
		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				LadonScan(ScanType,ip)
			}(ip.String())
		}
		wg.Wait()
	}
	} else {
		LadonScan(ScanType,Target)
	}
	//log.Println("Finished")	
	fmt.Println(" Finished: "+time.Now().Format("2006-01-02 03:04:05"))
}
func End(){
fmt.Println(" Finished: "+time.Now().Format("2006-01-02 03:04:05"))
os.Exit(0)
}
func CScan(ScanType string,Target string){
	ip:=strings.Replace(Target, "/c", "", -1)
	ip = strings.Replace(ip, "/C", "", -1)
	var wg sync.WaitGroup
	for i:=1;i<256;i++ {
		ip:=fmt.Sprintf("%s.%d",ip,i)
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				//fmt.Println("c: "+ip)
				LadonScan(ScanType,ip);
			}(ip)
	}
	wg.Wait()
}
func BScan(ScanType string,Target string){
	ip:=strings.Replace(Target, "/b", "", -1)
	ip = strings.Replace(ip, "/B", "", -1)
	for i:=1;i<256;i++ {
		ip:=fmt.Sprintf("%s.%d",ip,i)
		CScan(ScanType,ip)
	}
}
func AScan(ScanType string,Target string){
	ip:=strings.Replace(Target, "/a", "", -1)
	ip = strings.Replace(ip, "/A", "", -1)
	for i:=1;i<256;i++ {
		ip:=fmt.Sprintf("%s.%d",ip,i)
		BScan(ScanType,ip)
	}
}
func LadonScan(ScanType string,Target string) {
	if ScanType == "PINGSCAN" ||ScanType == "PING" {
		ping.PingName(Target)
	} else if ScanType == "ICMPSCAN" ||ScanType == "ICMP" {
		icmp.Icmp(Target,debugLog)
	} else if ScanType == "PORTSCAN" || ScanType == "SCANPORT" {
		port.ScanPort(Target)
	} else if ScanType == "HTTPBANNER" {
		http.HttpBanner(Target)
	} else if ScanType == "HTTPTITLE" || ScanType == "WEBTITLE" {
		http.ScanTitle(Target)
	} else if ScanType == "T3SCAN" || ScanType=="WEBLOGICSCAN" {
		t3.T3version(Target)
	} else if ScanType == "MS17010" {
		smb.MS17010(Target,3)
	} else if ScanType == "SMBSCAN" {
		smb.SmbScan(ScanType,Target)
	} else if ScanType == "FTPSCAN" {
		ftp.FtpScan(ScanType,Target)
	} else if ScanType == "SMBGHOST"||ScanType == "CVE-2020-0796" {
		smb.SmbGhost(Target,445)
	} else if ScanType == "SSHSCAN" {
		ssh.SshScan(ScanType,Target)
	} else if ScanType == "MYSQLSCAN" {
		mysql.MysqlScan(ScanType,Target)
	} else if ScanType == "WINRMSCAN" {
		winrm.WinrmScan(ScanType,Target)
	}
}