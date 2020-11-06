# LadonGo 1.0

http://k8gege.org/Ladon/LadonGo.html

### 简介
LadonGo一款开源渗透扫描器框架，使用它可轻松批量探测C段、B段存活主机、指纹识别、端口扫描、密码爆破、高危漏洞检测等。
1.0版本包含11个模块功能，高危漏洞检测MS17010、SmbGhost，密码爆破SmbScan、SshScan、FtpScan、MysqlScan，存活探测/信息收集/指纹识别PingScan、IcmpScan，BannerScan、WeblogicScan，端口扫描PortScan。

### Detection:
PingScan        (Using system ping to detect Online hosts)<br>
IcmpScan        (Using ICMP Protocol to detect Online hosts)<br>
BannerScan      (Using HTTP Protocol to detect Banner hosts)<br>
WeblogicScan    (Using T3 Protocol to detect Weblogic hosts)<br>
PortScan        (Scan hosts open ports using TCP protocol)<br>
MS17010         (Using SMB Protocol to detect MS17010 hosts))<br>
SmbGhost        (Using SMB Protocol to detect SmbGhost hosts))<br>
<br>
### Brute-Force:
SmbScan         (Using SMB Protocol to Brute-For 445 Port))<br>
SshScan         (Using SSH Protocol to Brute-For 22 Port))<br>
FtpScan         (Using FTP Protocol to Brute-For 21 Port))<br>
MysqlScan       (Using Mysql Protocol to Brute-For 3306 Port))<br>
<br>
### Example:
Ladon 192.168.1.8/24 MS17010

![image](http://k8gege.org/k8img/LadonGo/MacMS17010.png)
