# Generate unsigned certificate
`go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com`

### Generate unsigned certificate on WINDOWS
windows may have issues with go env GOROOT
`go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost`

Instead of go env GOROOT you can just use the path to the GO SDK wherever it is on your computer

### Generate unsigned certificate on Debian 8 (jessie)
`sudo apt-get install certbot -t jessie-backports`

`sudo certbot certonly --webroot -w /var/www/example -d example.com -d www.example.com -w /var/www/thing -d thing.is -d m.thing.is`

### To generate self-signed SSL using Open SSL:
`openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem`
