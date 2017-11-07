## Download GO Packages
Packages will be downloaded into the `$GOPATH` system varible
`
go get -u gopkg.in/mgo.v2
`

---
## Systemd
`cd /etc/systemd/system/`

`nano golang.service`

```
[Unit]
Description=Go Server

[Service]
ExecStart=/var/www/web
WorkingDirectory=/var/www/
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
```

### Start the service
* `systemctl enable golang.service`
* `systemctl start golang.service`
* `systemctl status golang.service`
* `systemctl stop golang.service`

## Build on Linux
* `export GOPATH=/var/go/web/`
* `echo $GOPATH`
```
GOOS=linux GOARCH=amd64 go build -o web
```

## Build on Windows
* `set GOARCH=amd64`
* `set GOARCH=386`
* `set GOOS=linux`
* `set GOOS=windows`
* `echo %GOROOT%`
```
go build -o hello.exe hello.go
```


## Installing go on Debian
```
cd /usr/local
curl -LO https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
shasum -a 256 go1.9.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
```

`vi ~/.profile`
```
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=/var/go
```
`source ~/.profile`

`echo $GOROOT`