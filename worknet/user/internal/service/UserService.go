package service

import (
	"github.com/prometheus/common/log"
	pb "user/api"

	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Service) WorknetUserNumber(ctx context.Context, req *empty.Empty) (resp *pb.WorknetUserNumberResp, err error) {
	resp, err = s.dao.RawUserNumber(ctx)
	if err != nil {
		log.Error("WorknetUserNumber发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetUserActivity(ctx context.Context, req *empty.Empty) (resp *pb.WorknetUserActivityResp, err error) {
	resp, err = s.dao.RawActivityUserNumber(ctx)
	if err != nil {
		log.Error("WorknetUserActivity发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetUserSex(ctx context.Context, req *empty.Empty) (resp *pb.WorknetUserSexResp, err error) {
	resp, err = s.dao.RawSexUserNumber(ctx)
	if err != nil {
		log.Error("WorknetUserSex发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetUserNumberChange(ctx context.Context, req *empty.Empty) (resp *pb.WorknetUserChangeResp, err error) {
	resp, err = s.dao.RawUserNumberChange(ctx)
	if err != nil {
		log.Error("WorknetUserNumberChange发现错误,错误信息：", err)
	}
	return
}
