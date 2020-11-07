# Ladon For Golang

http://k8gege.org/Ladon/LadonGo.html

### 简介
LadonGo一款开源渗透扫描器框架，使用它可轻松批量探测C段、B段存活主机、指纹识别、端口扫描、密码爆破、高危漏洞检测等。1.0版本包含11个模块功能，高危漏洞检测MS17010、SmbGhost，密码爆破SmbScan、SshScan、FtpScan、MysqlScan，存活探测/信息收集/指纹识别PingScan、IcmpScan，BannerScan、WeblogicScan，端口扫描PortScan。

### 功能模块

Detection:<br>
PingScan        (Using system ping to detect Online hosts)<br>
IcmpScan        (Using ICMP Protocol to detect Online hosts)<br>
BannerScan      (Using HTTP Protocol to detect Banner hosts)<br>
WeblogicScan    (Using T3 Protocol to detect Weblogic hosts)<br>
PortScan        (Scan hosts open ports using TCP protocol)<br>
MS17010         (Using SMB Protocol to detect MS17010 hosts))<br>
SmbGhost        (Using SMB Protocol to detect SmbGhost hosts))<br>
<br>
Brute-Force:<br>
SmbScan         (Using SMB Protocol to Brute-For 445 Port))<br>
SshScan         (Using SSH Protocol to Brute-For 22 Port))<br>
FtpScan         (Using FTP Protocol to Brute-For 21 Port))<br>
MysqlScan       (Using Mysql Protocol to Brute-For 3306 Port))<br>
<br>
Example:<br>
Ladon 192.168.1.8/24 MS17010

### 源码编译
```Bash
go get github.com/k8gege/LadonGo
go build Ladon.go
```

### 使用教程

#### 帮助
```Bash
Ladon help
Ladon Detection
Ladon BruteForce
```

#### 用法
Ladon IP/机器名/CIDR 扫描模块

#### 例子

##### 信息收集、漏洞检测

Ping扫描C段存活主机（任意权限）<br>
Ladon 192.168.1.8/24 PingScan<br>
<br>
ICMP扫描C段存活主机(管理员权限)<br>
Ladon 192.168.1.8/24 IcmpScan<br>
<br>
SMB扫描C段永恒之蓝MS17010漏洞主机<br>
Ladon 192.168.1.8/24 MS17010<br>
<br>
SMB扫描C段永恒之黑SmbGhost漏洞主机<br>
Ladon 192.168.1.8/24 SmbGhost<br>
<br>
T3扫描C段开放WebLogic的主机<br>
Ladon 192.168.1.8/24 T3Scan<br>
<br>
HTTP扫描C段开放Web站点Banner<br>
Ladon 192.168.1.8/24 BannerScan


##### 密码爆破、弱口令
扫描C段445端口Windows机器弱口令<br>
Ladon 192.168.1.8/24 SmbScan<br>
<br>
扫描C段22端口Linux机器SSH弱口令<br>
Ladon 192.168.1.8/24 SshScan<br>
<br>
扫描C段21端口FTP服务器弱口令<br>
Ladon 192.168.1.8/24 FtpScan<br>
<br>
扫描C段3306端口Mysql服务器弱口令<br>
Ladon 192.168.1.8/24 MysqlScan


### 扫描速度
1.和Ladon一样，ICMP探测C段仅需1秒<br>
2.Ping扫描C段大约11秒，支持任意权限<br>
3.其它模块自行测试<br>

### 跨平台/全平台/全系统

#### TestOn

Kali 2019<br>
SUSE 10<br>
CentOS 5.8<br>
CentOS 6.8  <br>
Fedora5<br>
XP、2003、Win7、Win8.1、Win10、2008 R2、2012 R2<br>
RedHat5.7 <br>
BT5-R3  (Ubuntu 8)<br>
MacOS 10.15<br>
<br>
以上系统测试成功，其它系统未测，若某些系统不支持可自行编译


#### MacOS x64 10.15
![image](http://k8gege.org/k8img/LadonGo/MacMS17010.png)

#### Linux
![image](http://k8gege.org/k8img/LadonGo/LnxMS17010.PNG)

#### Windows
![image](http://k8gege.org/k8img/LadonGo/WinMS17010.PNG)

### Download

#### LadonGo (ALL OS)
https://github.com/k8gege/LadonGo

#### Ladon (Windows & Cobalt Strike)

历史版本: https://github.com/k8gege/Ladon/releases
7.0版本：http://k8gege.org/Download
7.5版本：K8小密圈


<div style="text-align: center; width: 710px; border: green solid 0px;">
<img alt="" src="http://k8gege.org/img/k8team.jpg" style="display: inline-block;width: 250px;height: 300px;" />
</div>
