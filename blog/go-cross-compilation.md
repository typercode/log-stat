mac下编译linux可执行程序：

GOOS=linux GOARCH=amd64 go build -o req src/stat/req.go
