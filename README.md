ewlog
=====

Simple log for multiple output

Doc
-----

- [API Reference](http://godoc.org/github.com/ender-wan/ewlog)
- [Examples](https://godoc.org/github.com/ender-wan/ewlog#example-AddLogOutput)

Installation
-----

go get github.com/ender-wan/ewlog

Example
-----
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
 output example
```
2017/06/16 17:19:45 main.go:13 main.main [Info] - ewlog
```
