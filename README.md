### Simple log for multiple output


#### added a output will output to file and stdout, without InitLog will output to stdout
```
logfile, err := os.OpenFile("LogFile", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    fmt.Println(err)
    return
}
defer logfile.Close()

ewlog.InitLog(logfile, 2)
```
