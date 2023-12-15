
# Ladon Scanner For Golang
### Wiki
http://k8gege.org/Ladon/LadonGo.html<br>

[![Author](https://img.shields.io/badge/Author-k8gege-blueviolet)](https://github.com/k8gege) 
[![Ladon](https://img.shields.io/badge/LadonGo-5.2-yellowgreen)](https://github.com/k8gege/LadonGo) 
[![Bin](https://img.shields.io/badge/LadonGo-Bin-ff69b4)](https://github.com/k8gege/LadonGo/releases) 
[![GitHub issues](https://img.shields.io/github/issues/k8gege/LadonGo)](https://github.com/k8gege/LadonGo/issues) 
[![Github Stars](https://img.shields.io/github/stars/k8gege/LadonGo)](https://github.com/k8gege/LadonGo) 
[![GitHub forks](https://img.shields.io/github/forks/k8gege/LadonGo)](https://github.com/k8gege/LadonGo)
[![GitHub license](https://img.shields.io/github/license/k8gege/LadonGo)](https://github.com/k8gege/LadonGo)
[![Downloads](https://img.shields.io/github/downloads/k8gege/LadonGo/total?label=Release%20Download)](https://github.com/k8gege/LadonGo/releases/latest)

### Introduction

LadonGo is an open source intranet penetration scanner framework, which can be used to easily detect segment C, B, A live hosts, fingerprint identification, port scanning, password explosion, remote execution, high-risk vulnerability detection, etc. Version 4.0 includes 37 functions, high risk vulnerability detection MS17010, SmbGhost, remote execution of SshCmd, WinrmCmd, PhpShell, JspShell, GoWebShell, L, 12 protocol password explosion Smb/Ssh/Ftp/Mysql/Mssql/Oracle/Sqlplus/Winrm/HttpBasic/Edits/MongoDB/RouterOS, survival detection/information collection/fingerprint identification NbtInfo, OnlinePC, Ping, Icmp, SnmpScan, HttpBanner, HttpTitle, TcpBanner, WeblogicScan, O xidScan, Port scan / service probe portscan, forward to Socks5 proxy< br>

LadonGO 5.2 Pentest Scanner framework 全平台Go开源内网渗透扫描器框架,Windows/Linux/Mac内网渗透，使用它可轻松一键批量探测C段、B段、A段存活主机、高危漏洞检测MS17010、SmbGhost，远程执行SSH/Winrm，密码爆破SMB/SSH/FTP/Mysql/Mssql/Oracle/Winrm/HttpBasic/Redis，端口扫描服务识别PortScan指纹识别/HttpBanner/HttpTitle/TcpBanner/Weblogic/Oxid多网卡主机，端口扫描服务识别PortScan。

### Development environment

OS: Kali 2019 X64<br>

IDE: Mousepad<br>

Go: 1.13 Linux<br>


### Function module

#### Detection

 . | . 
-|-
OnlinePC |         (Using ICMP/SNMP/Ping detect Online hosts)
PingScan |         (Using system ping to detect Online hosts)
IcmpScan |         (Using ICMP Protocol to detect Online hosts)
SnmpScan |         (Using Snmp Protocol to detect Online hosts)
HttpBanner |       (Using HTTP Protocol Scan Web Banner)
HttpTitle |        (Using HTTP protocol Scan Web titles)
T3Scan |           (Using T3 Protocol Scan Weblogic hosts)
PortScan |         (Scan hosts open ports using TCP protocol)
TcpBanner |        (Scan hosts open ports using TCP protocol)
OxidScan |         (Using dcom Protocol enumeration network interfaces)
NbtInfo |        (Scan hosts open ports using NBT protocol)

#### VulDetection

 . | . 
-|-
MS17010 |          (Using SMB Protocol to detect MS17010 hosts)
SmbGhost |         (Using SMB Protocol to detect SmbGhost hosts)
CVE-2021-21972 |   (Check VMware vCenter 6.5 6.7 7.0 Rce Vul)
CVE-2021-26855 |   (Check CVE-2021-26855 Microsoft Exchange SSRF)

 
#### BruteForce

 . | . 
-|-
SmbScan |          (Using SMB Protocol to Brute-For 445 Port)
SshScan |          (Using SSH Protocol to Brute-For 22 Port)
FtpScan |          (Using FTP Protocol to Brute-For 21 Port)
401Scan |          (Using HTTP BasicAuth to Brute-For web Port)
MysqlScan |        (Using Mysql Protocol to Brute-For 3306 Port)
MssqlScan |        (Using Mssql Protocol to Brute-For 1433 Port)
OracleScan |       (Using Oracle Protocol to Brute-For 1521 Port)
MongodbScan |       (Using Mongodb Protocol to Brute-For 27017 Port)
WinrmScan |        (Using Winrm Protocol to Brute-For 5985 Port)
SqlplusScan |      (Using Oracle Sqlplus Brute-For 1521 Port)
RedisScan |      (Using Redis Protocol to Brute-For 6379 Port)

#### RemoteExec

 . | . 
-|-
SshCmd |           (SSH Remote command execution Default 22 Port)
WinrmCmd |         (Winrm Remote command execution Default 5985 Port)
PhpShell |         (Php WebShell command execution Default 80 Port)
GoWebShell  |    (Go WebShell Default http://IP:888/web)
WinJspShell  |   (JSP Shell Remote command execution Default 80 Port)
LnxJspShell  |   (JSP Shell Remote command execution Default 80 Port)
LnxRevShell   |  (Bash Reverse Shell)

#### Exploit

 . | . 
-|-
PhpStudyDoor |     (PhpStudy 2016 & 2018 BackDoor Exploit)
CVE-2018-14847 |   (Export RouterOS Password 6.29 to 6.42)

#### Socks5

 . | . 
-|-
Socks5 |     (Socks5 forward proxy server)

### Build
```Bash
go get github.com/k8gege/LadonGo
go build Ladon.go
```

### Make
```Bash
make windows
make linux
make mac
```

### Install
#### Linux/Mac
```Bash
make install
```

#### Windows
```Bash
go run install.go
```

### Usage

#### help
```Bash
Ladon FuncList
Ladon Detection
Ladon VulDetection
adon BruteFor
Ladon RemoteExec
Ladon Exploit
Ladon Example
```

#### Usage
Ladon IP/pcname/CIDR/URL/txt moudle

```Bash
Ladon 192.168.1.8 MS17010
Ladon 192.168.1.8/24 MS17010
Ladon 192.168.1/c MS17010
Ladon 192.168/b MS17010
Ladon 192/a MS17010

Ladon 192.168.1-192.168.5 MS17010
Ladon http://192.168.1.8:8080 BasicAuthScan
Ladon ip.txt MS17010
Ladon url.txt HttpBanner
```

#### Examples


##### Forward Socks5 proxy server

Ladon Socks5 192.168.1.8 1080


##### Information collection and vulnerability detection


Ping scan segment C live hosts (with any permission) < br >

Ladon 192.168.1.8/24 PingScan<br>


ICMP scanning section C survival host (administrator authority)<br>

Ladon 192.168.1.8/24 IcmpScan<br>


SNMP scanning of surviving hosts and devices in Segment C<br>

Ladon 192.168.1.8/24 SnmpScan<br>


SMB scans C-segment eternal blue ms17010 vulnerable hosts < br >

Ladon 192.168.1.8/24 MS17010<br>


SMB scans C-segment eternal black smbghost vulnerability host < br >

Ladon 192.168.1.8/24 SmbGhost<br>


T3 scans the hosts with open WebLogic in Segment C<br>

Ladon 192.168.1.8/24 T3Scan<br>


HTTP scanning C-segment open Web site Banner<br>

Ladon 192.168.1.8/24 BannerScan<br>


HTTP scanning section C open Web site title<br>

Ladon 192.168.1.8/24 HttpTitle<br>


TCP scans the open port service information of section C. < br >

Ladon 192.168.1.8/24 TcpBanner<br>


TCP scans common open ports of segment C hosts<br>

Ladon 192.168.1.8/24 PortScan<br>


##### Password explosion, weak password


Scan weak password of windows machine on port 445 of section C. < br >

Ladon 192.168.1.8/24 SmbScan<br>


Scan the SSH weak password of the 22 port Linux machine in section C<br>

Ladon 192.168.1.8/24 SshScan<br>


Scan the weak password of FTP server on port 21 of section C. < br >

Ladon 192.168.1.8/24 FtpScan<br>


Scan the weak password of MySQL server on port 3306 in section C<br>

Ladon 192.168.1.8/24 MysqlScan<br>


Scan weak password of Oracle server on port 1521 in segment C<br>

Ladon 192.168.1.8/24 OracleScan<br>


Scan the weak password of MongoDB server on port 27017 in section C<br>

Ladon 192.168.1.8/24 MongodbScan<br>


Scan weak password of Oracle server on port 1521 of section C. < br >

Ladon 192.168.1.8/24 SqlplusScan<br>


Scan section C 5985 port Winrm server weak password<br>

Ladon 192.168.1.8/24 WinrmScan<br>


Scan the empty password of redis server on port 6379 of section C. < br >

Ladon 192.168.1.8/24 RedisScan<br>


Scan C-segment 8728 port Routeros router < br >

Ladon 192.168.1.8/24 RouterOSScan<br>

##### Remote command execution

```Bash
Ladon SshCmd host port user pass cmd
Ladon WinrmCmd host port user pass cmd
Ladon PhpShell url pass cmd
Ladon PhpStudyDoor url cmd
```

SshCmd & WinrmCmd
![image](http://k8gege.org/k8img/LadonGo/LnxSshWinrm.PNG)

PhpShell & PhpStudyDoor
![image](http://k8gege.org/k8img/LadonGo/phpshell.PNG)

#### SCAN IP/24 (192.168.1/c)
 . | . | . 
-|-|-
ICMP |3ms  |1/20s
WebTitle| 10ms| 1/6s
T3Scan |15ms| 1/4s
EthScan |2ms | 1/30s

#### Scan B(192.168/b)
 . | . | . 
-|-|-
EthScan  | 23Min |  1 Port
T3Scan   |  1h |  4 Port
WebTitle | 40Min | 1 Port
MS17010  |12Min | 1 Port
Snmp    |  20Min| 1 Port

PS: the scanning speed is actually similar to the speed of Ladon. Net version, but there is no special record, because the speed of go version is recorded by the way during the rewrite test


### Cross platform / whole platform / whole system

Support old and new operating systems, especially old Linux systems. Many online tools can't be used at all or various errors are reported

#### TestOn

ID | OS 
-|-
0 | WinXP
1 | Win 2003
2 | Win 7
3 | Win 8.1
4 | Win 10
5 | Win 2008 R2
6 | Win 2012 R2
7 | Win 2019
8 | Kali 1.0.2
9 | Kali 2018
10 | Kali 2019
11 | SUSE 10
12 | CentOS 5.8
13 | CentOS 6.3
14 | CentOS 6.8  
15 | Fedora 5
16 | RedHat 5.7 
17 | BT5-R3  
18 | MacOS 10.15
19 | Ubuntu 8
20 | Ubuntu 18


#### MacOS x64 10.15
![image](http://k8gege.org/k8img/LadonGo/MacMS17010.png)

#### Linux
![image](http://k8gege.org/k8img/LadonGo/LnxMS17010.PNG)

#### Windows
![image](http://k8gege.org/k8img/LadonGo/OxidScan.PNG)


### Download

#### LadonGo (ALL OS)
https://github.com/k8gege/LadonGo/releases<br>
http://k8gege.org/Download/LadonGo.rar

#### Ladon (Windows & Cobalt Strike)

History: https://github.com/k8gege/Ladon/releases<br>
911 Ver：http://k8gege.org/Download<br>


<div style="text-align: center; width: 710px; border: green solid 0px;">
<img alt="" src="http://k8gege.org/img/k8team.jpg" style="display: inline-block;width: 250px;height: 300px;" />
</div>


## Stargazers over time

[![Stargazers over time](https://starchart.cc/k8gege/LadonGo.svg)](https://starchart.cc/k8gege/LadonGo)

<img align='right' src="https://profile-counter.glitch.me/LadonGo/count.svg" width="200">
