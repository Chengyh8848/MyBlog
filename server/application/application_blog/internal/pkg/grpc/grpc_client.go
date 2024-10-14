package grpc

import (
	"application_blog/internal/protocal/pb/blog"
	"google.golang.org/grpc"
)

type grpcClients struct {
	BlogHostClient blog.BlogGRPCClient
}

var GrpcClients grpcClients

func InitGrpcClients(conn *grpc.ClientConn) {
	GrpcClients = grpcClients{
		BlogHostClient: blog.NewBlogGRPCClient(conn),
	}
}
