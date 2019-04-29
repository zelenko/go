
### Validate Input
```go
func validInput(input string) bool {
	switch input {
	case "Option1", "OPtion2", "valid1", "valid2", "valid3":
		return true
	}
	return false
}
```

### AlphaNumeric
```go
func alphaNumeric(input string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(input, "")
}
```
