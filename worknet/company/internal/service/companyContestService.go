package service

import (
	pb "company/api"
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Service) WorknetCompanyContestChange(ctx context.Context, req *pb.WorknetContestChangeReq) (resp *pb.WorknetContestChangeResp, err error) {
	resp, err = s.dao.RawCompanyContestNumberChange(ctx, req)
	if err != nil {
		log.Error("WorknetCompanyContestChange发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetContestNumberRank(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCompanyRankResp, err error) {
	resp, err = s.dao.RawCompanyContestRank(ctx)
	if err != nil {
		log.Error("WorknetContestNumberRank发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetContestingNumberRank(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCompanyRankResp, err error) {
	resp, err = s.dao.RawCompanyContestingRank(ctx)
	if err != nil {
		log.Error("WorknetContestingNumberRank发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetContestApplyNumberRank(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCompanyRankResp, err error) {
	resp, err = s.dao.RawCompanyContestApplyRank(ctx)
	if err != nil {
		log.Error("WorknetContestApplyNumberRank发现错误,错误信息：", err)
	}
	return
}

