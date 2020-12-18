
# Ladon Scanner For Golang
### Wiki
http://k8gege.org/Ladon/LadonGo.html<br>

### 简介
LadonGo一款开源内网渗透扫描器框架，使用它可轻松一键探测C段、B段、A段存活主机、指纹识别、端口扫描、密码爆破、远程执行、高危漏洞检测等。3.2版本包含24个模块功能，高危漏洞检测MS17010、SmbGhost，远程执行SshCmd、WinrmCmd，密码爆破SmbScan、SshScan、FtpScan、MysqlScan、MssqlScan、OracleScan、SqlplusScan、WinrmScan、HttpBasicScan，存活探测/信息收集/指纹识别OnlinePC、Ping、Icmp、SnmpScan，HttpBanner、HttpTitle、TcpBanner、WeblogicScan、OxidScan，端口扫描PortScan、服务探测PortScanBanner。<br>

### 开发环境
OS: Kali 2019 X64<br>
IDE: 记事本Mousepad<br>
Go:  1.13 Linux<br>

### 功能模块

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


#### VulDetection

 . | . 
-|-
MS17010 |          (Using SMB Protocol to detect MS17010 hosts)
SmbGhost |         (Using SMB Protocol to detect SmbGhost hosts)


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
WinrmScan |        (Using Winrm Protocol to Brute-For 5985 Port)
SqlplusScan |      (Using Oracle Sqlplus Brute-For 1521 Port)


#### RemoteExec

 . | . 
-|-
SshCmd |           (SSH Remote command execution Default 22 Port)
WinrmCmd |         (Winrm Remote command execution Default 5985 Port)


#### Exploit

 . | . 
-|-
PhpStudyDoor |     (PhpStudy 2016 & 2018 BackDoor Exploit)

### 源码编译
```Bash
go get github.com/k8gege/LadonGo
go build Ladon.go
```

### 使用教程

#### 帮助
```Bash
Ladon FuncList
Ladon Detection
Ladon VulDetection
adon BruteFor
Ladon RemoteExec
Ladon Exploit
Ladon Example
```

#### 用法
Ladon IP/机器名/CIDR/URL/txt 扫描模块

```Bash
IP Ladon 192.168.1.8 MS17010
C段 Ladon 192.168.1.8/24 MS17010
C段 Ladon 192.168.1/c MS17010
B段 Ladon 192.168/b MS17010
A段 Ladon 192/a MS17010

C段(1-5) Ladon 192.168.1-192.168.5 MS17010
URL Ladon http://192.168.1.8:8080 BasicAuthScan
IP列表  Ladon ip.txt MS17010
URL列表 Ladon url.txt HttpBanner
```

#### 例子

##### 信息收集、漏洞检测

Ping扫描C段存活主机（任意权限）<br>
Ladon 192.168.1.8/24 PingScan<br>

ICMP扫描C段存活主机(管理员权限)<br>
Ladon 192.168.1.8/24 IcmpScan<br>

SNMP扫描C段存活主机、设备信息<br>
Ladon 192.168.1.8/24 SnmpScan<br>

SMB扫描C段永恒之蓝MS17010漏洞主机<br>
Ladon 192.168.1.8/24 MS17010<br>

SMB扫描C段永恒之黑SmbGhost漏洞主机<br>
Ladon 192.168.1.8/24 SmbGhost<br>

T3扫描C段开放WebLogic的主机<br>
Ladon 192.168.1.8/24 T3Scan<br>

HTTP扫描C段开放Web站点Banner<br>
Ladon 192.168.1.8/24 BannerScan<br>

HTTP扫描C段开放Web站点标题<br>
Ladon 192.168.1.8/24 HttpTitle<br>

TCP扫描C段开放端口服务信息<br>
Ladon 192.168.1.8/24 TcpBanner<br>

TCP扫描C段主机常见开放端口<br>
Ladon 192.168.1.8/24 PortScan<br>

##### 密码爆破、弱口令

扫描C段445端口Windows机器弱口令<br>
Ladon 192.168.1.8/24 SmbScan<br>

扫描C段22端口Linux机器SSH弱口令<br>
Ladon 192.168.1.8/24 SshScan<br>

扫描C段21端口FTP服务器弱口令<br>
Ladon 192.168.1.8/24 FtpScan<br>

扫描C段3306端口Mysql服务器弱口令<br>
Ladon 192.168.1.8/24 MysqlScan<br>

扫描C段1521端口Oracle服务器弱口令<br>
Ladon 192.168.1.8/24 OracleScan<br>

扫描C段1521端口Oracle服务器弱口令<br>
Ladon 192.168.1.8/24 SqlplusScan<br>

扫描C段5985端口Winrm服务器弱口令<br>
Ladon 192.168.1.8/24 WinrmScan<br>

##### 远程命令执行

```Bash
Ladon SshCmd host port user pass cmd
Ladon WinrmCmd host port user pass cmd
```

![image](http://k8gege.org/k8img/LadonGo/LnxSshWinrm.PNG)

#### 扫C段(192.168.1/c)
 . | . | . 
-|-|-
ICMP |3毫秒  |1/20秒
WebTitle| 10毫秒| 1/6秒
T3Scan |15毫秒| 1/4秒
EthScan |2毫秒 | 1/30秒

#### 扫B段(192.168/b)
 . | . | . 
-|-|-
EthScan  | 23分钟 |  1个端口
T3Scan   |  1小时 |  4个端口
WebTitle | 40分钟 | 1个端口
MS17010  |12分钟 | 1个端口
Snmp    |  20分钟| 1个端口

PS：扫描速度实际上和Ladon .net版速度也差不多，只是没专门记录，因为重写测试过程中顺便记录一下GO版速度

### 跨平台/全平台/全系统

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

以上系统测试成功，其它系统未测，若不支持可自行编译<br>

#### MacOS x64 10.15
![image](http://k8gege.org/k8img/LadonGo/MacMS17010.png)

#### Linux
![image](http://k8gege.org/k8img/LadonGo/LnxMS17010.PNG)

#### Windows
![image](http://k8gege.org/k8img/LadonGo/OxidScan.PNG)


### Download

#### LadonGo (ALL OS)
https://github.com/k8gege/LadonGo/releases<br>

#### Ladon (Windows & Cobalt Strike)

历史版本: https://github.com/k8gege/Ladon/releases<br>
7.0版本：http://k8gege.org/Download<br>
7.5版本：K8小密圈<br>


<div style="text-align: center; width: 710px; border: green solid 0px;">
<img alt="" src="http://k8gege.org/img/k8team.jpg" style="display: inline-block;width: 250px;height: 300px;" />
</div>

