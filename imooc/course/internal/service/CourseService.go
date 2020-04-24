package service

import (
	"context"
	pb "course/api"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Service) ImoocCourseScore(ctx context.Context, req *empty.Empty) (resp *pb.ImoocCourseScoreSectionResp, err error) {
	resp, err = s.dao.RawCourseScoreSection(ctx)
	if err != nil {
		log.Error("ImoocCourseScore发现错误,错误信息：", err)
	}
	return
}

func (s *Service) ImoocCourseAllNumber(ctx context.Context, req *empty.Empty) (resp *pb.ImoocCourseAllNumberResp, err error) {
	resp, err = s.dao.RawCourseNumber(ctx)
	if err != nil {
		log.Error("ImoocCourseAllNumber发现错误,错误信息：", err)
	}
	return
}

func (s *Service) ImoocCourseSection(ctx context.Context, req *empty.Empty) (resp *pb.ImoocCourseSectionResp, err error) {
	resp, err = s.dao.RawCourseSection(ctx)
	if err != nil {
		log.Error("ImoocCourseSection发现错误,错误信息：", err)
	}
	return
}

func (s *Service) ImoocCourseScoreRank(ctx context.Context, req *pb.ImoocCourseRankReq) (resp *pb.ImoocCourseRankResp, err error) {
	resp, err = s.dao.RawCourseScoreRank(ctx, req)
	if err != nil {
		log.Error("ImoocCourseScoreRank发现错误,错误信息：", err)
	}
	return
}
func (s *Service) ImoocCourseLearnerRank(ctx context.Context, req *pb.ImoocCourseRankReq) (resp *pb.ImoocCourseRankResp, err error) {
	resp, err = s.dao.RawCourseLearnerRank(ctx, req)
	if err != nil {
		log.Error("ImoocCourseLearnerRank发现错误,错误信息：", err)
	}
	return
}
