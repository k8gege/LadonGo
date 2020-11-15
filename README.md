# Ladon Scanner For Golang
### Wiki
http://k8gege.org/Ladon/LadonGo.html<br>

### 简介
LadonGo一款开源渗透扫描器框架，使用它可轻松批量探测C段、B段存活主机、指纹识别、端口扫描、密码爆破、远程执行、高危漏洞检测等。2.0版本包含15个模块功能，高危漏洞检测MS17010、SmbGhost，远程执行SshCmd、WinrmCmd，密码爆破SmbScan、SshScan、FtpScan、MysqlScan、WinrmScan，存活探测/信息收集/指纹识别PingScan、IcmpScan，BannerScan、HttpTitle、WeblogicScan，端口扫描PortScan。<br>


### 功能模块

 . | . 
-|-
PingScan |         (Using system ping to detect Online hosts)
IcmpScan |         (Using ICMP Protocol to detect Online hosts)
HttpBanner |       (Using HTTP Protocol Scan Web Banner)
HttpTitle |        (Using HTTP protocol Scan Web titles)
T3Scan |           (Using T3 Protocol Scan Weblogic hosts)
PortScan |         (Scan hosts open ports using TCP protocol)
MS17010 |          (Using SMB Protocol to detect MS17010 hosts))
SmbGhost |          (Using SMB Protocol to detect SmbGhost hosts))
SmbScan |          (Using SMB Protocol to Brute-For 445 Port))
SshScan |          (Using SSH Protocol to Brute-For 22 Port))
FtpScan |          (Using FTP Protocol to Brute-For 21 Port))
MysqlScan |        (Using Mysql Protocol to Brute-For 3306 Port))
WinrmScan |        (Using Winrm Protocol to Brute-For 5985 Port))
SshCmd |           (SSH Remote command execution Default 22 Port))
WinrmCmd |         (Winrm Remote command execution Default 5985 Port))

![image](http://k8gege.org/k8img/LadonGo/Help.PNG)

### 源码编译
```Bash
go get github.com/k8gege/LadonGo
go build Ladon.go
```

### 使用教程

#### 帮助
```Bash
Ladon help<br>
Ladon Detection<br>
Ladon BruteForce<br>
```

#### 用法
Ladon IP/机器名/CIDR 扫描模块<br>
```Bash
Ladon 192.168.1.8/24 MS17010
Ladon 192.168.1/c MS17010
Ladon 192.168/b MS17010
Ladon 192/a MS17010
```

#### 例子

##### 信息收集、漏洞检测

Ping扫描C段存活主机（任意权限）<br>
Ladon 192.168.1.8/24 PingScan<br>

ICMP扫描C段存活主机(管理员权限)<br>
Ladon 192.168.1.8/24 IcmpScan<br>

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

##### 密码爆破、弱口令

扫描C段445端口Windows机器弱口令<br>
Ladon 192.168.1.8/24 SmbScan<br>

扫描C段22端口Linux机器SSH弱口令<br>
Ladon 192.168.1.8/24 SshScan<br>

扫描C段21端口FTP服务器弱口令<br>
Ladon 192.168.1.8/24 FtpScan<br>

扫描C段3306端口Mysql服务器弱口令<br>
Ladon 192.168.1.8/24 MysqlScan<br>

扫描C段5985端口Winrm服务器弱口令<br>
Ladon 192.168.1.8/24 WinrmScan<br>

##### 远程命令执行

```Bash
Ladon SshCmd host port user pass cmd
Ladon WinrmCmd host port user pass cmd
```

![image](http://k8gege.org/k8img/LadonGo/LnxSshWinrm.PNG)

### 扫描速度
1.和Ladon一样，ICMP探测C段仅需1秒<br>
2.Ping扫描C段大约11秒，支持任意权限<br>
3.其它模块自行测试<br>

### 跨平台/全平台/全系统

#### TestOn

ID | OS <br>
-|-
0 | WinXP<br>
1 | Win 2003<br>
2 | Win 7<br>
3 | Win 8.1<br>
4 | Win 10<br>
5 | Win 2008 R2<br>
6 | Win 2012 R2<br>
7 | Win 2019<br>
8 | Kali 1.0.2<br>
9 | Kali 2018<br>
10 | Kali 2019<br>
11 | SUSE 10<br>
12 | CentOS 5.8<br>
13 | CentOS 6.3<br>
14 | CentOS 6.8  <br>
15 | Fedora 5<br>
16 | RedHat 5.7 <br>
17 | BT5-R3  <br>
18 | MacOS 10.15<br>
19 | Ubuntu 8<br>
20 | Ubuntu 18<br>

以上系统测试成功，其它系统未测，若不支持可自行编译<br>

#### MacOS x64 10.15
![image](http://k8gege.org/k8img/LadonGo/MacMS17010.png)

#### Linux
![image](http://k8gege.org/k8img/LadonGo/LnxMS17010.PNG)

#### Windows
![image](http://k8gege.org/k8img/LadonGo/WinMS17010.PNG)


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
