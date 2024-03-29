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

var PathWorknetUserPing = "/service.v1.WorknetUser/Ping"
var PathWorknetUserWorknetUserNumber = "/service.v1.WorknetUser/WorknetUserNumber"
var PathWorknetUserWorknetUserActivity = "/service.v1.WorknetUser/WorknetUserActivity"
var PathWorknetUserWorknetUserSex = "/service.v1.WorknetUser/WorknetUserSex"
var PathWorknetUserWorknetUserNumberChange = "/service.v1.WorknetUser/WorknetUserNumberChange"

var PathTokenVerify = "/service.v1.Token/Verify"

// WorknetUserBMServer is the server API for WorknetUser service.
type WorknetUserBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	// `method:"POST"`
	WorknetUserNumber(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetUserNumberResp, err error)

	// `method:"POST"`
	WorknetUserActivity(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetUserActivityResp, err error)

	// `method:"POST"`
	WorknetUserSex(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetUserSexResp, err error)

	// `method:"POST"`
	WorknetUserNumberChange(ctx context.Context, req *google_protobuf1.Empty) (resp *WorknetUserChangeResp, err error)
}

var WorknetUserSvc WorknetUserBMServer

func worknetUserPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetUserSvc.Ping(c, p)
	c.JSON(resp, err)
}

func worknetUserWorknetUserNumber(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetUserSvc.WorknetUserNumber(c, p)
	c.JSON(resp, err)
}

func worknetUserWorknetUserActivity(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetUserSvc.WorknetUserActivity(c, p)
	c.JSON(resp, err)
}

func worknetUserWorknetUserSex(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetUserSvc.WorknetUserSex(c, p)
	c.JSON(resp, err)
}

func worknetUserWorknetUserNumberChange(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WorknetUserSvc.WorknetUserNumberChange(c, p)
	c.JSON(resp, err)
}

// RegisterWorknetUserBMServer Register the blademaster route
func RegisterWorknetUserBMServer(e *bm.Engine, server WorknetUserBMServer) {
	WorknetUserSvc = server
	e.GET("/service.v1.WorknetUser/Ping", worknetUserPing)
	e.POST("/service.v1.WorknetUser/WorknetUserNumber", worknetUserWorknetUserNumber)
	e.POST("/service.v1.WorknetUser/WorknetUserActivity", worknetUserWorknetUserActivity)
	e.POST("/service.v1.WorknetUser/WorknetUserSex", worknetUserWorknetUserSex)
	e.POST("/service.v1.WorknetUser/WorknetUserNumberChange", worknetUserWorknetUserNumberChange)
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
