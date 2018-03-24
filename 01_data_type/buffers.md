There are two ways to make a buffer:

```
buffer := bytes.NewBuffer(make([]byte, 0))
buffer := &bytes.Buffer{}
```

buffer can take string
`buffer.WriteString(line);`

or buffer can read from io.Writer

`buffer.ReadFrom() // takes io.Writer`


buffer can take slice of byte or string
```
buffer.Write([]byte(strings.Join(record, "\t") + "\n"))
buffer.WriteString(line)
```

