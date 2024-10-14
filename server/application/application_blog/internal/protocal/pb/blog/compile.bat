:: Install proto3.
:: https://github.com/protocolbuffers/protobuf/releases
:: Update protoc Go bindings via
::go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
::go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

::
:: See also
::  https://github.com/grpc/grpc-go/tree/master/examples

protoc --go_out=. --go-grpc_out=. *.proto