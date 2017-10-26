# qloger

Qlogger is a logger which rotates log file by date and size

```bash
go get github.com/domac/qlogger
```

examples

```go

func main() {
    maxSize := 1024 * 1024 *1024 //1G
	logger, err := NewRotatorQLogger("/tmp/test.log", "debug", true, false, true, maxSize)
	if err != nil {
        return
	}
	logger.Infoln("hello")
}

```