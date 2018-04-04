# Data validation

Accept only valid data using a switch.  This also can be done using regexp package.
```go
func letterOp(code int) bool {
	switch chars[code].category {
	case "Lu", "Ll", "Lt", "Lm", "Lo":
		return true
	}
	return false
}
```


source: https://github.com/golang/go/wiki/Switch
