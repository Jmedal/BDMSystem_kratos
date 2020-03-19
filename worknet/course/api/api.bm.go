// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

package api

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathWorknetCoursePing = "/service.v1.WorknetCourse/Ping"
var PathWorknetCourseWorknetCourseActivity = "/service.v1.WorknetCourse/WorknetCourseActivity"
var PathWorknetCourseWorknetCourseNumberChange = "/service.v1.WorknetCourse/WorknetCourseNumberChange"
var PathWorknetCourseWorknetCourseList = "/service.v1.WorknetCourse/WorknetCourseList"
var PathWorknetCourseWorknetStudentChange = "/service.v1.WorknetCourse/WorknetStudentChange"
var PathWorknetCourseWorknetStudentFinish = "/service.v1.WorknetCourse/WorknetStudentFinish"
var PathWorknetCourseWorknetAverageScoreList = "/service.v1.WorknetCourse/WorknetAverageScoreList"
var PathWorknetCourseWorknetAverageScoreSection = "/service.v1.WorknetCourse/WorknetAverageScoreSection"
var PathWorknetCourseWorknetCourseStarsRank = "/service.v1.WorknetCourse/WorknetCourseStarsRank"
var PathWorknetCourseWorknetCourseClickRank = "/service.v1.WorknetCourse/WorknetCourseClickRank"
var PathWorknetCourseWorknetCourseVideoRank = "/service.v1.WorknetCourse/WorknetCourseVideoRank"

var PathTokenVerify = "/service.v1.Token/Verify"

// WorknetCourseBMServer is the server API for WorknetCourse service.
type WorknetCourseBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	// `method:"POST"`
	WorknetCourseActivity(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetCourseActivityResp, err error)

	// `method:"POST"`
	WorknetCourseNumberChange(ctx context.Context, req *WorknetCourseChangeReq) (resp *WorknetCourseChangeResp, err error)

	// `method:"POST"`
	WorknetCourseList(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetCourseListResp, err error)

	// `method:"POST"`
	WorknetStudentChange(ctx context.Context, req *WorknetStudentChangeReq) (resp *WorknetStudentChangeResp, err error)

	// `method:"POST"`
	WorknetStudentFinish(ctx context.Context, req *WorknetStudentFinishReq) (resp *WorknetStudentFinishResp, err error)

	// `method:"POST"`
	WorknetAverageScoreList(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetAverageScoreListResp, err error)

	// `method:"POST"`
	WorknetAverageScoreSection(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetAverageScoreSectionResp, err error)

	// `method:"POST"`
	WorknetCourseStarsRank(ctx context.Context, req *WorknetCourseRankReq) (resp *WorknetCourseRankResp, err error)

	// `method:"POST"`
	WorknetCourseClickRank(ctx context.Context, req *WorknetCourseRankReq) (resp *WorknetCourseRankResp, err error)

	// `method:"POST"`
	WorknetCourseVideoRank(ctx context.Context, req *WorknetCourseRankReq) (resp *WorknetCourseRankResp, err error)
}

var WorknetCourseSvc WorknetCourseBMServer

func worknetCoursePing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.Ping(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetCourseActivity(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetCourseActivity(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetCourseNumberChange(c *bm.Context) {
	p := new(WorknetCourseChangeReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetCourseNumberChange(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetCourseList(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetCourseList(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetStudentChange(c *bm.Context) {
	p := new(WorknetStudentChangeReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetStudentChange(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetStudentFinish(c *bm.Context) {
	p := new(WorknetStudentFinishReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetStudentFinish(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetAverageScoreList(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetAverageScoreList(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetAverageScoreSection(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetAverageScoreSection(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetCourseStarsRank(c *bm.Context) {
	p := new(WorknetCourseRankReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetCourseStarsRank(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetCourseClickRank(c *bm.Context) {
	p := new(WorknetCourseRankReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetCourseClickRank(c, p)
	c.JSON(resp, err)
}

func worknetCourseWorknetCourseVideoRank(c *bm.Context) {
	p := new(WorknetCourseRankReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetCourseSvc.WorknetCourseVideoRank(c, p)
	c.JSON(resp, err)
}

// RegisterWorknetCourseBMServer Register the blademaster route
func RegisterWorknetCourseBMServer(e *bm.Engine, server WorknetCourseBMServer) {
	WorknetCourseSvc = server
	e.GET("/service.v1.WorknetCourse/Ping", worknetCoursePing)
	e.POST("/service.v1.WorknetCourse/WorknetCourseActivity", worknetCourseWorknetCourseActivity)
	e.POST("/service.v1.WorknetCourse/WorknetCourseNumberChange", worknetCourseWorknetCourseNumberChange)
	e.POST("/service.v1.WorknetCourse/WorknetCourseList", worknetCourseWorknetCourseList)
	e.POST("/service.v1.WorknetCourse/WorknetStudentChange", worknetCourseWorknetStudentChange)
	e.POST("/service.v1.WorknetCourse/WorknetStudentFinish", worknetCourseWorknetStudentFinish)
	e.POST("/service.v1.WorknetCourse/WorknetAverageScoreList", worknetCourseWorknetAverageScoreList)
	e.POST("/service.v1.WorknetCourse/WorknetAverageScoreSection", worknetCourseWorknetAverageScoreSection)
	e.POST("/service.v1.WorknetCourse/WorknetCourseStarsRank", worknetCourseWorknetCourseStarsRank)
	e.POST("/service.v1.WorknetCourse/WorknetCourseClickRank", worknetCourseWorknetCourseClickRank)
	e.POST("/service.v1.WorknetCourse/WorknetCourseVideoRank", worknetCourseWorknetCourseVideoRank)
}

// TokenBMServer is the server API for Token service.
type TokenBMServer interface {
	// `method:"POST" internal:"true"`
	Verify(ctx context.Context, req *VerifyTokenReq) (resp *VerifyTokenResp, err error)
}

var TokenSvc TokenBMServer

func tokenVerify(c *bm.Context) {
	p := new(VerifyTokenReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := TokenSvc.Verify(c, p)
	c.JSON(resp, err)
}

// RegisterTokenBMServer Register the blademaster route
func RegisterTokenBMServer(e *bm.Engine, server TokenBMServer) {
	TokenSvc = server
	e.POST("/service.v1.Token/Verify", tokenVerify)
}