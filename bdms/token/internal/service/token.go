package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"time"
	pb "token/api"
)

func (s *Service) Request(ctx context.Context, req *pb.NewTokenReq) (resp *pb.NewTokenResp, err error) {
	var token string
	var randomKey string
	if token, randomKey, err = s.dao.RequestUploadToken(ctx, req.UserId, req.Operator, time.Now().Unix(), req.Expire); err != nil {
		log.Error("New a upload token failure: %v", err)
		return
	}
	resp = &pb.NewTokenResp{
		AccessToken: token,
		RandomKey:   randomKey,
	}
	return
}

func (s *Service) Verify(ctx context.Context, req *pb.VerifyTokenReq) (resp *pb.VerifyTokenResp, err error) {
	var userId int64
	var randomKey string
	userId, randomKey, err = s.dao.VerifyUploadToken(ctx, req.AccessToken)
	resp = &pb.VerifyTokenResp{
		UserId:    userId,
		RandomKey: randomKey,
	}
	return
}
