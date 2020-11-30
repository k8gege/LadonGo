package snmp
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"github.com/alouca/gosnmp"
	"fmt"
)

func SnmpOK(host string){
s, err := gosnmp.NewGoSNMP(host, "public", gosnmp.Version2c, 5)
if err != nil {
	//log.Fatal(err)
}
resp, err := s.Get(".1.3.6.1.2.1.1.1.0")
if err == nil {
	for _, v := range resp.Variables {
		switch v.Type {
		case gosnmp.OctetString:
			fmt.Println("Snmp:  "+host)
		}
	}
}
}

func GetInfo(host string){
s, err := gosnmp.NewGoSNMP(host, "public", gosnmp.Version2c, 5)
if err != nil {
	//log.Fatal(err)
}
resp, err := s.Get(".1.3.6.1.2.1.1.1.0")
if err == nil {
	for _, v := range resp.Variables {
		switch v.Type {
		case gosnmp.OctetString:
			fmt.Println("Snmp: "+host,"\t",HostName(host),"\t", v.Value.(string))
		}
	}
}
}

func HostName(host string)string{
s, err := gosnmp.NewGoSNMP(host, "public", gosnmp.Version2c, 5)
if err != nil {
	//log.Fatal(err)
}
resp, err := s.Get(".1.3.6.1.2.1.1.5.0")
if err == nil {
	for _, v := range resp.Variables {
		switch v.Type {
		case gosnmp.OctetString:
			//fmt.Println("SNMP:",host,"\t", v.Value.(string))
			return  v.Value.(string)
		}
	}
}
return ""
}