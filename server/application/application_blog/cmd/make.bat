set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-w -s" -o application_blog main.go
