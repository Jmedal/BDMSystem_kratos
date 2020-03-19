package service

import (
	"context"
	pb "course/api"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Service) WorknetStudentFinish(ctx context.Context, req *pb.WorknetStudentFinishReq) (resp *pb.WorknetStudentFinishResp, err error) {
	resp, err = s.dao.RawStudentFinishNumber(ctx, req)
	if err != nil {
		log.Error("WorknetStudentFinish发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetAverageScoreList(ctx context.Context, req *empty.Empty) (resp *pb.WorknetAverageScoreListResp, err error) {
	resp, err = s.dao.RawCourseAverageScoreList(ctx)
	if err != nil {
		log.Error("WorknetStudentFinish发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetAverageScoreSection(ctx context.Context, req *empty.Empty) (resp *pb.WorknetAverageScoreSectionResp, err error) {
	resp, err = s.dao.RawCourseAverageScoreSection(ctx)
	if err != nil {
		log.Error("WorknetStudentFinish发现错误,错误信息：", err)
	}
	return
}
