# Ladon For Golang

Whttp://k8gege.org/Ladon/LadonGo.html

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

以上系统测试成功，其它系统未测,若不支持可自行编译


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

历史版本: https://github.com/k8gege/Ladon/releases<br>
7.0版本：http://k8gege.org/Download<br>
7.5版本：K8小密圈<br>

<div style="text-align: center; width: 710px; border: green solid 0px;">
<img alt="" src="http://k8gege.org/img/k8team.jpg" style="display: inline-block;width: 250px;height: 300px;" />
</div>

### 为什么使用GO
现有Ladon版本无法兼容一些系统，代理又丢包非常蛋疼。虽然Python版也是跨平台，但是编译体积大，二是有些依赖包（依赖底层库）在某些系统安装非常麻烦甚至装不上，有些编译后不能执行等原因。所以这几天重新学了下GO，现学现卖使用Golang重写Ladon框架，先加一些功能看看效果，GO和PY差不多很简单，框架弄好后，使用开源库一下就可以添加好几个功能模块，然后再测14个操作系统下程序的兼容性，无论性能、体积、兼容性都远甩Python几条街，最主要是编译的程序可在一些旧操作系统上执行，Python可能受限于py版本、相关依赖包或GCC、GLID、SSL等库版本影响，Go可以很好解决这些问题。

缺点： 很多API库没有人封装或者根本不能用，想要实现Ladon的所有功能或者说一半的功能，两三个月都搞不定，如Ladon的OsScan模块用到的协议就已比这个LadonGo 1.0现在的11个功能还要多。如果用Python的话就非常快了，各种依赖库、各种现有POC，写好扫描框架，稍微改一下集成起来就是功能非常多的扫描器，但是目标PY版本低、操作系统老，本地编译再丢过去都不定能运行，本地都未必能编译，所以选用GO。
