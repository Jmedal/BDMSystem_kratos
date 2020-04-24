package service

import (
	"context"
	pb "course/api"
	"github.com/bilibili/kratos/pkg/log"
)

func (s *Service) GetCoursePage(ctx context.Context, req *pb.GetCoursePageReq) (resp *pb.GetCoursePageResp, err error) {
	resp, err = s.dao.RawCoursePage(ctx, req)
	if err != nil {
		log.Error("GetCoursePage发现错误,错误信息：", err)
	}
	return
}
