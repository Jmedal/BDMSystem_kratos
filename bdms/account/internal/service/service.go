package service

import (
	pb "account/api"
	"account/internal/dao"
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.AccountServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

//登录
func (s *Service) Login(ctx context.Context, req *pb.LoginReq) (reply *pb.LoginResp, err error) {
	fmt.Println("帐号：", req.Account)
	fmt.Println("密码：", req.Password)
	fmt.Println("save：", req.Save)
	user, err := s.dao.RawUserAccount(ctx, req.Account)
	if err != nil {
		reply = &pb.LoginResp{
			AccessToken: "",
			RandomKey:   "",
			Result:      "db server error",
		}
		return
	}
	if user.Account == req.Account &&
		user.Password == req.Password {
		var token string
		var randomKey string
		var expire int64
		//是否保存密码 保存：1，不保存：2
		if req.Save == "1" {
			expire = 365 * 24 * 60 * 60 //一年
		} else {
			expire = 24 * 60 * 60 //一天
		}
		if token, randomKey, err = s.dao.RequestToken(ctx, user.ID, "CK", expire); err != nil {
			reply = &pb.LoginResp{
				AccessToken: "",
				RandomKey:   "",
				Result:      "token server error",
			}
			return
		}
		reply = &pb.LoginResp{
			AccessToken: token,
			RandomKey:   randomKey,
			Result:      "success",
		}
		return
	}
	reply = &pb.LoginResp{
		AccessToken: "",
		RandomKey:   "",
		Result:      "password error",
	}
	return
}

//注册
func (s *Service) Register(ctx context.Context, req *pb.RegisterReq) (reply *pb.RegisterResp, err error) {
	//reply = &pb.HelloResp{
	//	Content: "hello " + req.Name,
	//}
	//fmt.Printf("hello url %s", req.Name)
	return
}

//用户信息
func (s *Service) UserInfo(ctx context.Context, req *pb.UserInfoReq) (reply *pb.UserInfoResp, err error) {
	user, err := s.dao.RawUserInfo(ctx, req.UserId)
	if err != nil {
		return
	}
	reply = &pb.UserInfoResp{
		Avatar:     user.Avatar,
		Account:    user.Account,
		Name:       user.Name,
		Birthday:   user.Birthday.Time().Format("2006-01-02"),
		Sex:        user.Sex,
		Email:      user.Email,
		Phone:      user.Phone,
		DeptId:     user.DeptId,
		Status:     user.Status,
		CreateTime: user.CreateTime.Time().Format("2006-01-02"),
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
