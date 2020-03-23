package service

import (
	pb "account/api"
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
)

//登录
func (s *Service) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	fmt.Println("帐号：", req.Account)
	fmt.Println("密码：", req.Password)
	fmt.Println("save：", req.Save)
	resp, err = s.dao.RawUser(ctx, req)
	if err != nil {
		log.Error("Login发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetUserPage(ctx context.Context, req *pb.GetUserPageReq) (resp *pb.GetUserPageResp, err error) {
	resp, err = s.dao.RawUserPage(ctx, req)
	if err != nil {
		log.Error("GetUserList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) AddUser(ctx context.Context, req *pb.AddUserReq) (resp *pb.AddUserResp, err error) {
	resp, err = s.dao.InsertUser(ctx, req)
	if err != nil {
		log.Error("AddUser发现错误,错误信息：", err)
	}
	return
}

func (s *Service) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (resp *pb.UpdateUserResp, err error) {
	resp, err = s.dao.UpdateUser(ctx, req)
	if err != nil {
		log.Error("UpdateUser发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (resp *pb.DeleteUserResp, err error) {
	resp, err = s.dao.DeleteUser(ctx, req)
	if err != nil {
		log.Error("DeleteUser发现错误,错误信息：", err)
	}
	return
}

func (s *Service) SetUserStatus(ctx context.Context, req *pb.SetUserStatusReq) (resp *pb.SetUserStatusResp, err error) {
	resp, err = s.dao.SetUserStatus(ctx, req)
	if err != nil {
		log.Error("SetUserStatus发现错误,错误信息：", err)
	}
	return
}

func (s *Service) SetUserRole(ctx context.Context, req *pb.SetUserRoleReq) (resp *pb.SetUserRoleResp, err error) {
	resp, err = s.dao.SetUserRole(ctx, req)
	if err != nil {
		log.Error("SetUserRole发现错误,错误信息：", err)
	}
	return
}

func (s *Service) CheckAccount(ctx context.Context, req *pb.CheckAccountReq) (resp *pb.CheckAccountResp, err error) {
	resp, err = s.dao.RawAccount(ctx, req)
	if err != nil {
		log.Error("CheckAccount发现错误,错误信息：", err)
	}
	return
}
