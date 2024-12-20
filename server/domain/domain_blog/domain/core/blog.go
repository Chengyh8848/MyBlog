package core

import (
	"context"
	"domain_blog/infrastructure/database/persistence"
	pb "domain_blog/infrastructure/grpc/pb"
)

type BlogDomainI interface {
	GetGroupYearMonthByIsPublished(ctx context.Context, request *pb.GetGroupYearMonthByIsPublishedRequest) (*pb.GetGroupYearMonthByIsPublishedReply, error)
}

type BlogService struct {
	BlogPersistence persistence.BlogPersistenceI
}

func NewBlogService() *BlogService {
	return &BlogService{
		BlogPersistence: persistence.NewBlogPersistence(),
	}
}

func (s *BlogService) GetGroupYearMonthByIsPublished(ctx context.Context, request *pb.GetGroupYearMonthByIsPublishedRequest) (*pb.GetGroupYearMonthByIsPublishedReply, error) {
	result, err := s.BlogPersistence.GetGroupYearMonthByIsPublished(ctx)
	return &pb.GetGroupYearMonthByIsPublishedReply{
		Value: result}, err
}
