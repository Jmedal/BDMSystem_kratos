package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
	pb "role/api"
)

func (s *Service) GetRoleList(ctx context.Context, e *empty.Empty) (resp *pb.GetRoleListResp, err error) {
	resp, err = s.dao.RawRoleList(ctx)
	if err != nil {
		log.Error("GetRoleList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) AddRole(ctx context.Context, req *pb.AddRoleReq) (resp *pb.AddRoleResp, err error) {
	resp, err = s.dao.InsertRole(ctx, req)
	if err != nil {
		log.Error("AddRole发现错误,错误信息：", err)
	}
	return
}

func (s *Service) UpdateRole(ctx context.Context, req *pb.UpdateRoleReq) (resp *pb.UpdateRoleResp, err error) {
	resp, err = s.dao.UpdateRole(ctx, req)
	if err != nil {
		log.Error("UpdateRole发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (resp *pb.DeleteRoleResp, err error) {
	resp, err = s.dao.DeleteRole(ctx, req)
	if err != nil {
		log.Error("DeleteRole发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetRoleRights(ctx context.Context, req *pb.GetRoleRightsReq) (resp *pb.GetRoleRightsResp, err error) {
	resp, err = s.dao.RawRoleMenus(ctx, req)
	if err != nil {
		log.Error("GetRoleRights发现错误,错误信息：", err)
	}
	return
}

//记得加顶级菜单
func (s *Service) SetRoleRights(ctx context.Context, req *pb.SetRoleRightsReq) (resp *pb.SetRoleRightsResp, err error) {
	resp, err = s.dao.SetRoleMenus(ctx, req)
	if err != nil {
		log.Error("SetRoleRights发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteRoleRights(ctx context.Context, req *pb.DeleteRoleRightsReq) (resp *pb.DeleteRoleRightsResp, err error) {
	resp, err = s.dao.DeleteRoleMenus(ctx, req)
	if err != nil {
		log.Error("DeleteRoleRights发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteRoleNullRights(ctx context.Context, req *pb.DeleteRoleNullRightsReq) (resp *pb.DeleteRoleNullRightsResp, err error) {
	resp, err = s.dao.DeleteRoleNullMenus(ctx, req)
	if err != nil {
		log.Error("DeleteRoleNullRights发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetRoleOptions(ctx context.Context, e *empty.Empty) (resp *pb.GetRoleOptionsResp, err error) {
	resp, err = s.dao.RawRoleOptions(ctx)
	if err != nil {
		log.Error("GetRoleOptions发现错误,错误信息：", err)
	}
	return
}
