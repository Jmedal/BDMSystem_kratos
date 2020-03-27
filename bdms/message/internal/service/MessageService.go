package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	pb "message/api"
)

func (s *Service) GetMessagePage(ctx context.Context, req *pb.GetMessagePageReq) (resp *pb.GetMessagePageResp, err error) {
	resp, err = s.dao.RawMessagePage(ctx, req)
	if err != nil {
		log.Error("GetUserList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) AddMessage(ctx context.Context, req *pb.AddMessageReq) (resp *pb.AddMessageResp, err error) {
	resp, err = s.dao.InsertMessage(ctx, req)
	if err != nil {
		log.Error("AddMessage发现错误,错误信息：", err)
	}
	return
}

func (s *Service) UpdateMessage(ctx context.Context, req *pb.UpdateMessageReq) (resp *pb.UpdateMessageResp, err error) {
	resp, err = s.dao.UpdateMessage(ctx, req)
	if err != nil {
		log.Error("UpdateMessage发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteMessage(ctx context.Context, req *pb.DeleteMessageReq) (resp *pb.DeleteMessageResp, err error) {
	resp, err = s.dao.DeleteMessage(ctx, req)
	if err != nil {
		log.Error("DeleteMessage发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetMessageList(ctx context.Context, req *pb.GetMessageListReq) (resp *pb.GetMessageListResp, err error) {
	resp, err = s.dao.RawMessageList(ctx, req)
	if err != nil {
		log.Error("GetMessageList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) GetMessage(ctx context.Context, req *pb.GetMessageReq) (resp *pb.GetMessageResp, err error) {
	resp, err = s.dao.RawMessage(ctx, req)
	if err != nil {
		log.Error("GetMessage发现错误,错误信息：", err)
	}
	return
}

func (s *Service) SetMessageUserRead(ctx context.Context, req *pb.SetMessageUserReadReq) (resp *pb.SetMessageUserReadResp, err error) {
	resp, err = s.dao.SetMessageUserRead(ctx, req)
	if err != nil {
		log.Error("SetMessageUserRead发现错误,错误信息：", err)
	}
	return
}

func (s *Service) DeleteMessageUser(ctx context.Context, req *pb.DeleteMessageUserReq) (resp *pb.DeleteMessageUserResp, err error) {
	resp, err = s.dao.DeleteMessageUser(ctx, req)
	if err != nil {
		log.Error("DeleteMessageUser发现错误,错误信息：", err)
	}
	return
}
