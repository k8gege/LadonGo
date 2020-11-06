package main
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org
//Github: https://github.com/k8gege
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
	"strings"
	"log"
	"os"
	"runtime"
	"net"
	"sync"
)

func help() {
	Detection()
	BruteFor()
	Example()
}

func Example() {
	if runtime.GOOS=="windows" {fmt.Println("\nExample:")
	} else{fmt.Println("\033[32m\nExample:\033[0m")}
	if runtime.GOOS=="windows" {fmt.Println("Ladon 192.168.1.8/24 MS17010")
	} else{fmt.Println("./Ladon 192.168.1.8/24 MS17010")}
	fmt.Println("")
}

func Detection() {
	if runtime.GOOS=="windows" {fmt.Println("\nDetection:")
	} else{fmt.Println("\033[33m\nDetection:\033[0m")}
	fmt.Println("PingScan\t(Using system ping to detect Online hosts)")
	fmt.Println("IcmpScan\t(Using ICMP Protocol to detect Online hosts)")
	fmt.Println("BannerScan\t(Using HTTP Protocol to detect Banner hosts)")
	fmt.Println("WeblogicScan\t(Using T3 Protocol to detect Weblogic hosts)")
	fmt.Println("PortScan\t(Scan hosts open ports using TCP protocol)")
	fmt.Println("MS17010 \t(Using SMB Protocol to detect MS17010 hosts))")
	fmt.Println("SmbGhost\t(Using SMB Protocol to detect SmbGhost hosts))")
	
}

func BruteFor() {
	if runtime.GOOS=="windows" {fmt.Println("\nBrute-Force:")
	} else{fmt.Println("\033[35m\nBruteForce:\033[0m")}
	fmt.Println("SmbScan \t(Using SMB Protocol to Brute-For 445 Port))")
	fmt.Println("SshScan \t(Using SSH Protocol to Brute-For 22 Port))")
	fmt.Println("FtpScan \t(Using FTP Protocol to Brute-For 21 Port))")
	fmt.Println("MysqlScan \t(Using Mysql Protocol to Brute-For 3306 Port))")
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
	fmt.Println("LadonGo 1.0 by k8gege")
	fmt.Println("Arch: "+runtime.GOARCH+" OS: "+runtime.GOOS)
	ParLen := len(os.Args)
	if ParLen==1 {
		help()
		os.Exit(0)
	}
	EndPar := os.Args[ParLen-1]
	Target := os.Args[ParLen-2]
	if strings.ToUpper(EndPar)=="HELP"||strings.ToUpper(EndPar)=="/HELP"||strings.ToUpper(EndPar)=="-H"||strings.ToUpper(EndPar)=="-H" {
		help()
		os.Exit(0)
	}
	if strings.ToUpper(EndPar)=="BRUTEFOR"||strings.ToUpper(EndPar)=="BRUTE"||strings.ToUpper(EndPar)=="BRUTEFORCE" {
		BruteFor()
		os.Exit(0)
	}
	if strings.ToUpper(EndPar)=="DETECTION" {
		Detection()
		os.Exit(0)
	}
	if strings.ToUpper(EndPar)=="EXAMPLE"||strings.ToUpper(EndPar)=="USAGE" {
		Example()
		os.Exit(0)
	}
	fmt.Println("Targe: "+Target)
	log.Println("Start...")
	fmt.Println("Load "+EndPar)
	ScanType := strings.ToUpper(EndPar)
	if strings.Contains(Target, "/") {
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
				LadonScan(ScanType,ip);
			}(ip.String())
		}
		wg.Wait()
	}
	} else {
		LadonScan(ScanType,Target);
	}
	log.Println("Finished")	
}
func End(){
log.Println("Finished")
os.Exit(0)
}
func LadonScan(ScanType string,Target string) {
		if ScanType == "PINGSCAN" ||ScanType == "PING" {
			res,err := ping.CmdPing(Target)
			if err==nil && res==true {
				fmt.Println("PING: "+Target)
			}
		} else if ScanType == "ICMPSCAN" ||ScanType == "ICMP" {
			icmp.Icmp(Target,debugLog)
		} else if ScanType == "PORTSCAN" || ScanType == "SCANPORT" {
			port.ScanPort(Target)
		} else if ScanType == "BANNERSCAN" {
			http.HttpBanner(Target)
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
		}
}