# qloger

qloger is a logger which rotates log file by date and size

```bash
go get -u -v github.com/domac/qloger
```

### Examples

1. normal logger example

```go
func main() {
    logger, err := NewQLogger("/tmp/test.log", "info")
	if err != nil {
        return
    }
    logger.Infoln("hello")
}
rotates
```

 2. rotatable log example

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