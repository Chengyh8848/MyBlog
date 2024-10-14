set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -mod=vendor -ldflags "-w -s" -o domain_blog main.go


