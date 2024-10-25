package conversion

import (
	"domain_blog/infrastructure/database/entity"
	pb "domain_blog/infrastructure/grpc/pb"
)

func AboutsToPb(abouts []entity.About) []*pb.AboutDetail {
	var result []*pb.AboutDetail
	for _, about := range abouts {
		aboutPb := &pb.AboutDetail{
			Id:     int64(about.ID),
			NameEn: about.NameEn,
			NameZh: about.NameZh,
			Value:  about.Value,
		}
		result = append(result, aboutPb)
	}
	return result
}

func PbToAbouts(aboutsPb []*pb.AboutDetail) []entity.About {
	var result []entity.About
	for _, aboutPb := range aboutsPb {
		about := entity.About{
			NameEn: aboutPb.NameEn,
			NameZh: aboutPb.NameZh,
			Value:  aboutPb.Value,
		}
		about.ID = uint(aboutPb.Id)
		result = append(result, about)
	}
	return result
}
