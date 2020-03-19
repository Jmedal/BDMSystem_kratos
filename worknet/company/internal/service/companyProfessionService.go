package service

import (
	pb "company/api"
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
)


func (s *Service) WorknetLocationList(ctx context.Context, e *empty.Empty) (resp *pb.WorknetLocationListResp, err error) {
	resp, err = s.dao.RawLocationNumber(ctx)
	if err != nil {
		log.Error("WorknetLocationList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCompanyList(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCompanyListResp, err error) {
	resp, err = s.dao.RawCompanyList(ctx)
	if err != nil {
		log.Error("WorknetCompanyList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCompanyProfessionChange(ctx context.Context, req *pb.WorknetProfessionChangeReq) (resp *pb.WorknetProfessionChangeResp, err error) {
	resp, err = s.dao.RawCompanyProfessionNumberChange(ctx ,req)
	if err != nil {
		log.Error("WorknetCompanyProfessionChange发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetProfessionNumberRank(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCompanyRankResp, err error) {
	resp, err = s.dao.RawCompanyProfessionRank(ctx)
	if err != nil {
		log.Error("WorknetProfessionNumberRank发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCvNumberRank(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCompanyRankResp, err error) {
	resp, err = s.dao.RawCompanyCvRank(ctx)
	if err != nil {
		log.Error("WorknetCvNumberRank发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCvStatusNumber(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCvStatusNumberResp, err error) {
	resp, err = s.dao.RawCompanyCvStatusNumber(ctx)
	if err != nil {
		log.Error("WorknetCvStatusNumber发现错误,错误信息：", err)
	}
	return
}







