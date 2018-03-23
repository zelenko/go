
# Persisting
Start your go app when the server starts.  In other word _daemonize_ your golang program.

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
