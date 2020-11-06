# LadonGo 1.0

http://k8gege.org/Ladon/LadonGo.html

### 简介
LadonGo一款开源渗透扫描器框架，使用它可轻松批量探测C段、B段存活主机、指纹识别、端口扫描、密码爆破、高危漏洞检测等。
1.0版本包含11个模块功能，高危漏洞检测MS17010、SmbGhost，密码爆破SmbScan、SshScan、FtpScan、MysqlScan，存活探测/信息收集/指纹识别PingScan、IcmpScan，BannerScan、WeblogicScan，端口扫描PortScan。

### 功能

Detection:
PingScan        (Using system ping to detect Online hosts)
IcmpScan        (Using ICMP Protocol to detect Online hosts)
BannerScan      (Using HTTP Protocol to detect Banner hosts)
WeblogicScan    (Using T3 Protocol to detect Weblogic hosts)
PortScan        (Scan hosts open ports using TCP protocol)
MS17010         (Using SMB Protocol to detect MS17010 hosts))
SmbGhost        (Using SMB Protocol to detect SmbGhost hosts))

Brute-Force:
SmbScan         (Using SMB Protocol to Brute-For 445 Port))
SshScan         (Using SSH Protocol to Brute-For 22 Port))
FtpScan         (Using FTP Protocol to Brute-For 21 Port))
MysqlScan       (Using Mysql Protocol to Brute-For 3306 Port))

### 使用教程
和Ladon一样，用法非常简单
Ladon 目标IP/机器名/CIDR 扫描模块

例子：扫描C段是否存在永恒之蓝漏洞
Ladon 192.168.1.8/24 MS17010

### 扫描速度
1.和Ladon一样，ICMP探测C段仅需1秒
2.Ping扫描C段大约11秒，支持任意权限
3.其它模块自行测试

### 跨平台/全平台/全系统

#### TestOn

Kali 2019
SUSE 10
CentOS 5.8
CentOS 6.8  
Fedora5
XP、2003、Win7、Win8.1、Win10、2008 R2、2012 R2
RedHat5.7 
BT5-R3  (Ubuntu 8)
MacOS 10.15

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
