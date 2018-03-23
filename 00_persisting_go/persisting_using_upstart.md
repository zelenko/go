# Starting GO binary using UPSTART
_Upstart is used by Amazon Machine Images (AMI)_

Find out the version of Upstart:
`initctl version`

Upstart configuration files are located in this folder:
* `/etc/init` folder is for Upstart
* `/etc/init.d` folder is for SysVinit

**Go to folder: /etc/init**
`cd /etc/init`

create file: go.conf
`sudo nano go.conf`

"go.conf" file contents
```
description "go binary"
author "Developer"

# start on (local-filesystems and net-device-up)
start on runlevel [2345]
stop on runlevel [!2345]

respawn
chdir /home/user-name/go
exec /home/user-name/go/web-binary
```
---

### Reload configuration:
`sudo initctl reload-configuration`


### To start the service:
`sudo start go`
or
`sudo initctl start go`

To see the status:
`sudo initctl list | grep go`
or
`sudo status go`

### Update binary
If binary is updated, or overridden by newer version, then need to restart the service:
`sudo restart go`
or
`sudo initctl restart go`



---
# Build commands
generate build using specific Go file. If file not specified, all files are used.
`go build -o web-binary http.go`

test to see if the executable is working
`sudo /home/user-name/go/web-binary`
