package service

import (
	"context"
	pb "curriculum/api"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Service) WorknetCurriculumGraph(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCurriculumGraphResp, err error) {
	resp, err = s.dao.RawCurriculumTree(ctx)
	if err != nil {
		log.Error("WorknetCurriculumGraph发现错误,错误信息：", err)
	}
	return
}

func (s *Service) WorknetCurriculumCourse(ctx context.Context, e *empty.Empty) (resp *pb.WorknetCurriculumCourseResp, err error) {
	resp, err = s.dao.RawCurriculumCourse(ctx)
	if err != nil {
		log.Error("WorknetCurriculumCourse发现错误,错误信息：", err)
	}
	return

}
