package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
	pb "profession/api"
)


func (s *Service)WorknetProfessionChoice(ctx context.Context, req *empty.Empty) (resp *pb.WorknetProfessionChoiceResp, err error){
	resp, err = s.dao.RawProfessionChoice(ctx)
	if err != nil {
		log.Error("WorknetProfessionChoice发现错误,错误信息：", err)
	}
	return
}


func (s *Service)WorknetProfessionSalary(ctx context.Context, req *empty.Empty) (resp *pb.WorknetProfessionSalaryResp, err error){
	resp, err = s.dao.RawProfessionSalary(ctx)
	if err != nil {
		log.Error("WorknetProfessionSalary发现错误,错误信息：", err)
	}
	return
}


func (s *Service)WorknetProfessionTypeNumber(ctx context.Context, req *empty.Empty) (resp *pb.WorknetProfessionChoiceResp, err error){
	resp, err = s.dao.RawProfessionTypeNumber(ctx)
	if err != nil {
		log.Error("WorknetProfessionTypeNumber发现错误,错误信息：", err)
	}
	return
}


func (s *Service)WorknetProfessionTypeSalary(ctx context.Context, req *empty.Empty) (resp *pb.WorknetProfessionSalaryResp, err error){
	resp, err = s.dao.RawProfessionTypeSalary(ctx)
	if err != nil {
		log.Error("WorknetProfessionTypeSalary发现错误,错误信息：", err)
	}
	return
}
