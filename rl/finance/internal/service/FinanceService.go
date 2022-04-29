package service

import (
	"context"
	pb "finance/api"
	"github.com/bilibili/kratos/pkg/log"
)

func (s *Service) GetProjectPage(ctx context.Context, req *pb.GetProjectPageReq) (resp *pb.GetProjectPageResp, err error) {
	resp, err = s.dao.RawProjectPage(ctx, req)
	if err != nil {
		log.Error("GetProjectPage发现错误,错误信息：", err)
	}
	return
}

func (s *Service) AddProject(ctx context.Context, req *pb.AddProjectReq) (resp *pb.AddProjectResp, err error) {
	resp, err = s.dao.InsertProject(ctx, req)
	if err != nil {
		log.Error("AddProject发现错误,错误信息：", err)
	}
	return
}

func (s *Service) UpdateProject(ctx context.Context, req *pb.UpdateProjectReq) (resp *pb.UpdateProjectResp, err error) {
	resp, err = s.dao.UpdateProject(ctx, req)
	if err != nil {
		log.Error("UpdateProject发现错误,错误信息：", err)
	}
	return
}

func (s *Service) TerminationProject(ctx context.Context, req *pb.TerminationProjectReq) (resp *pb.TerminationProjectResp, err error) {
	resp, err = s.dao.FrozenStateProject(ctx, req)
	if err != nil {
		log.Error("TerminationProject发现错误,错误信息：", err)
	}
	return
}

func (s *Service) RestartProject(ctx context.Context, req *pb.RestartProjectReq) (resp *pb.RestartProjectResp, err error) {
	resp, err = s.dao.ActivityStateProject(ctx, req)
	if err != nil {
		log.Error("RestartProject发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetBillPage(ctx context.Context, req *pb.GetBillPageReq) (resp *pb.GetBillPageResp, err error) {
	resp, err = s.dao.RawBillPage(ctx, req)
	if err != nil {
		log.Error("GetBillPage发现错误,错误信息：", err)
	}
	return
}

func (s *Service) AddBill(ctx context.Context, req *pb.AddBillReq) (resp *pb.AddBillResp, err error) {
	resp, err = s.dao.InsertBill(ctx, req)
	if err != nil {
		log.Error("AddBill发现错误,错误信息：", err)
	}
	return
}

func (s *Service) UpdateBill(ctx context.Context, req *pb.UpdateBillReq) (resp *pb.UpdateBillResp, err error) {
	resp, err = s.dao.UpdateBill(ctx, req)
	if err != nil {
		log.Error("UpdateBill发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteBill(ctx context.Context, req *pb.DeleteBillReq) (resp *pb.DeleteBillResp, err error) {
	resp, err = s.dao.FrozenStateBill(ctx, req)
	if err != nil {
		log.Error("DeleteBill发现错误,错误信息：", err)
	}
	return
}
