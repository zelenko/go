## Declare and define struct
```
	Senddata := struct {
		prev	int
		now		int
		next	int
	}{
		next - 1,
		next,
		next + 1,
	}
```

## Download packages
`go get -u gopkg.in/mgo.v2`

---
## Systemd
`cd /etc/systemd/system/`

`nano vas.service`

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
* `systemctl enable vas.service`
* `systemctl start vas.service`
* `systemctl status vas.service`
* `systemctl stop vas.service`

## Build on Linux
* `export GOPATH=/var/go/web/`
```
GOOS=linux GOARCH=amd64 go build -o web
```

## Build on Windows
* `set GOARCH=amd64`
* `set GOARCH=386`
* `set GOOS=linux`
* `set GOOS=windows`
```
go build -o hello.exe hello.go
```
