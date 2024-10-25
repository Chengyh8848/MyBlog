package core

import (
	"context"
	"domain_blog/common"
	"domain_blog/domain/core/conversion"
	"domain_blog/infrastructure/database/persistence"
	pb "domain_blog/infrastructure/grpc/pb"
)

type AboutDomainI interface {
	GetAboutInfo(ctx context.Context, request *pb.GetAboutInfoRequest) (*pb.GetAboutInfoReply, error)
	GetAboutSetting(ctx context.Context, request *pb.GetAboutSettingRequest) (*pb.GetAboutSettingReply, error)
	UpdateAbout(ctx context.Context, request *pb.UpdateAboutRequest) (*pb.UpdateAboutReply, error)
	GetAboutCommentEnabled(ctx context.Context, request *pb.GetAboutCommentEnabledRequest) (*pb.GetAboutCommentEnabledReply, error)
}

type AboutService struct {
	AboutPersistence persistence.AboutPersistenceI
}

func NewAboutService() *AboutService {
	return &AboutService{
		AboutPersistence: persistence.NewAboutPersistence(),
	}
}

func (s *AboutService) GetAboutInfo(ctx context.Context, request *pb.GetAboutInfoRequest) (*pb.GetAboutInfoReply, error) {
	abouts, _ := common.Client.GetAbouts(common.REDIS_ABOUTS)
	if abouts != nil {
		return &pb.GetAboutInfoReply{AboutDetails: conversion.AboutsToPb(abouts)}, nil
	}
	abouts, err := s.AboutPersistence.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetAboutInfoReply{AboutDetails: conversion.AboutsToPb(abouts)}, nil
}

func (s *AboutService) GetAboutSetting(ctx context.Context, request *pb.GetAboutSettingRequest) (*pb.GetAboutSettingReply, error) {
	abouts, err := s.AboutPersistence.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetAboutSettingReply{AboutDetails: conversion.AboutsToPb(abouts)}, nil
}

func (s *AboutService) UpdateAbout(ctx context.Context, request *pb.UpdateAboutRequest) (*pb.UpdateAboutReply, error) {
	abouts := conversion.PbToAbouts(request.AboutDetails)
	for _, about := range abouts {
		err := s.AboutPersistence.UpdateAbout(ctx, &about)
		if err != nil {
			return nil, err
		}
	}
	return &pb.UpdateAboutReply{}, nil
}

func (s *AboutService) GetAboutCommentEnabled(ctx context.Context, request *pb.GetAboutCommentEnabledRequest) (*pb.GetAboutCommentEnabledReply, error) {
	enabled, err := s.AboutPersistence.GetAboutCommentEnable(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetAboutCommentEnabledReply{Enabled: enabled}, nil
}
