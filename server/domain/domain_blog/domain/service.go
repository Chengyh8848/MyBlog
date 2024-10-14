package domain

import (
	"domain_blog/domain/core"
	pb "domain_blog/infrastructure/grpc/pb"
)

type BlogGRPCServer struct {
	*pb.UnimplementedBlogGRPCServer
	CoreService *core.CoreService
}

func NewBlogGrpcServer() *BlogGRPCServer {
	return &BlogGRPCServer{
		UnimplementedBlogGRPCServer: &pb.UnimplementedBlogGRPCServer{},
		CoreService:                 core.NewCoreService(),
	}
}
