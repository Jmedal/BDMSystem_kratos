package service

import (
	"context"
	pb "course/api"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Service) WorknetCourseActivity(ctx context.Context, req *empty.Empty) (resp *pb.WorknetCourseActivityResp, err error) {
	resp, err = s.dao.RawActivityCourseNumber(ctx)
	if err != nil {
		log.Error("WorknetCourseActivity发现错误,错误信息：", err)
	}
	return
}
func (s *Service) WorknetCourseNumberChange(ctx context.Context, req *pb.WorknetCourseChangeReq) (resp *pb.WorknetCourseChangeResp, err error) {
	resp, err = s.dao.RawCourseNumberChange(ctx,req)
	if err != nil {
		log.Error("WorknetCourseNumberChange发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCourseList(ctx context.Context, req *empty.Empty) (resp *pb.WorknetCourseListResp, err error) {
	resp, err = s.dao.RawCourseList(ctx)
	if err != nil {
		log.Error("WorknetCourseList发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetStudentChange(ctx context.Context, req *pb.WorknetStudentChangeReq) (resp *pb.WorknetStudentChangeResp, err error) {
	resp, err = s.dao.RawStudentNumberChange(ctx,req)
	if err != nil {
		log.Error("WorknetStudentChange发现错误,错误信息：", err)
	}
	return
}
