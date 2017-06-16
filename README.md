### ewlog
Simple log for multiple output

### Doc

- [API Reference](http://godoc.org/github.com/ender-wan/ewlog)
- [Examples](https://godoc.org/github.com/ender-wan/ewlog#example-AddLogOutput)

### Installation

go get github.com/ender-wan/ewlog

### Example
```
func main() {
    logfile, err := os.OpenFile("LogFile", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer logfile.Close()

    ewlog.SetLogLevel(1)
    ewlog.AddLogOutput(logfile)

    ewlog.Info("ewlog")
}

```
 output
```
2017/06/15 08:48:59 Info ewlog
```
