package service

import (
	"context"
	pb "course/api"
	"github.com/bilibili/kratos/pkg/log"
)

func (s *Service) WorknetCourseStarsRank(ctx context.Context, req *pb.WorknetCourseRankReq) (resp *pb.WorknetCourseRankResp, err error) {
	resp, err = s.dao.RawCourseStarsRank(ctx, req)
	if err != nil {
		log.Error("WorknetCourseStarsRank发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCourseClickRank(ctx context.Context, req *pb.WorknetCourseRankReq) (resp *pb.WorknetCourseRankResp, err error) {
	resp, err = s.dao.RawCourseClickRank(ctx, req)
	if err != nil {
		log.Error("WorknetCourseClickRank发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCourseVideoRank(ctx context.Context, req *pb.WorknetCourseRankReq) (resp *pb.WorknetCourseRankResp, err error) {
	resp, err = s.dao.RawCourseVideoRank(ctx, req)
	if err != nil {
		log.Error("WorknetCourseVideoRank发现错误,错误信息：", err)
	}
	return
}
