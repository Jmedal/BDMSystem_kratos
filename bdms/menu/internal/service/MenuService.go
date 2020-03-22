package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
	pb "menu/api"
)

const (
	//顶级菜单Id
	_topMenuId = 1

	//顶级菜单Pid
	_topMenuPid = 0

	//顶级菜单层级
	_topMenuLevels = 0
)

func (s *Service) GetMenuList(ctx context.Context, e *empty.Empty) (resp *pb.GetMenuListResp, err error) {
	resp, err = s.dao.RawMenuList(ctx, _topMenuId, _topMenuLevels)
	if err != nil {
		log.Error("GetMenuList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) AddMenu(ctx context.Context, req *pb.AddMenuReq) (resp *pb.AddMenuResp, err error) {
	resp, err = s.dao.InsertMenu(ctx, req)
	if err != nil {
		log.Error("AddMenu发现错误,错误信息：", err)
	}
	return
}

func (s *Service) UpdateMenu(ctx context.Context, req *pb.UpdateMenuReq) (resp *pb.UpdateMenuResp, err error) {
	resp, err = s.dao.UpdateMenu(ctx, req)
	if err != nil {
		log.Error("UpdateMenu发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteMenu(ctx context.Context, req *pb.DeleteMenuReq) (resp *pb.DeleteMenuResp, err error) {
	resp, err = s.dao.DeleteMenu(ctx, req)
	if err != nil {
		log.Error("DeleteMenu发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetMenus(ctx context.Context, req *pb.GetMenusReq) (resp *pb.GetMenusResp, err error) {
	resp, err = s.dao.RawMenus(ctx, _topMenuPid, _topMenuLevels, req)
	if err != nil {
		log.Error("GetMenus发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetMenuOptions(ctx context.Context, req *pb.GetMenuOptionsReq) (resp *pb.GetMenuOptionsResp, err error) {
	resp, err = s.dao.RawMenuOptions(ctx, _topMenuPid, _topMenuLevels, req)
	if err != nil {
		log.Error("GetMenuOptions发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetAllMenuOptions(ctx context.Context, e *empty.Empty) (resp *pb.GetMenuOptionsResp, err error) {
	resp, err = s.dao.RawAllMenuOptions(ctx, _topMenuPid, _topMenuLevels)
	if err != nil {
		log.Error("GetAllMenuOptions发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetChildrenMenuList(ctx context.Context, req *pb.GetChildrenMenuListReq) (resp *pb.GetChildrenMenuListResp, err error) {
	resp, err = s.dao.RawChildrenMenuList(ctx, req)
	if err != nil {
		log.Error("GetChildrenMenuList发现错误,错误信息：", err)
	}
	return
}
