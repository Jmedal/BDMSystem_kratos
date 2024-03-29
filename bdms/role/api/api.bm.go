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

var PathRolePing = "/service.v1.Role/Ping"
var PathRoleGetRoleList = "/service.v1.Role/GetRoleList"
var PathRoleAddRole = "/service.v1.Role/AddRole"
var PathRoleUpdateRole = "/service.v1.Role/UpdateRole"
var PathRoleDeleteRole = "/service.v1.Role/DeleteRole"
var PathRoleGetRoleRights = "/service.v1.Role/GetRoleRights"
var PathRoleSetRoleRights = "/service.v1.Role/SetRoleRights"
var PathRoleDeleteRoleRights = "/service.v1.Role/DeleteRoleRights"
var PathRoleDeleteRoleNullRights = "/service.v1.Role/DeleteRoleNullRights"
var PathRoleGetRoleOptions = "/service.v1.Role/GetRoleOptions"

var PathMenuGetMenus = "/service.v1.Menu/GetMenus"
var PathMenuGetChildrenMenuList = "/service.v1.Menu/GetChildrenMenuList"

var PathTokenVerify = "/service.v1.Token/Verify"

// RoleBMServer is the server API for Role service.
type RoleBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	// `method:"POST"`
	GetRoleList(ctx context.Context, req *google_protobuf1.Empty) (resp *GetRoleListResp, err error)

	// `method:"POST"`
	AddRole(ctx context.Context, req *AddRoleReq) (resp *AddRoleResp, err error)

	// `method:"POST"`
	UpdateRole(ctx context.Context, req *UpdateRoleReq) (resp *UpdateRoleResp, err error)

	// `method:"POST"`
	DeleteRole(ctx context.Context, req *DeleteRoleReq) (resp *DeleteRoleResp, err error)

	// `method:"POST"`
	GetRoleRights(ctx context.Context, req *GetRoleRightsReq) (resp *GetRoleRightsResp, err error)

	// `method:"POST"`
	SetRoleRights(ctx context.Context, req *SetRoleRightsReq) (resp *SetRoleRightsResp, err error)

	// `method:"POST"`
	DeleteRoleRights(ctx context.Context, req *DeleteRoleRightsReq) (resp *DeleteRoleRightsResp, err error)

	// `method:"POST"`
	DeleteRoleNullRights(ctx context.Context, req *DeleteRoleNullRightsReq) (resp *DeleteRoleNullRightsResp, err error)

	// `method:"POST"`
	GetRoleOptions(ctx context.Context, req *google_protobuf1.Empty) (resp *GetRoleOptionsResp, err error)
}

var RoleSvc RoleBMServer

func rolePing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.Ping(c, p)
	c.JSON(resp, err)
}

func roleGetRoleList(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.GetRoleList(c, p)
	c.JSON(resp, err)
}

func roleAddRole(c *bm.Context) {
	p := new(AddRoleReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.AddRole(c, p)
	c.JSON(resp, err)
}

func roleUpdateRole(c *bm.Context) {
	p := new(UpdateRoleReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.UpdateRole(c, p)
	c.JSON(resp, err)
}

func roleDeleteRole(c *bm.Context) {
	p := new(DeleteRoleReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.DeleteRole(c, p)
	c.JSON(resp, err)
}

func roleGetRoleRights(c *bm.Context) {
	p := new(GetRoleRightsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.GetRoleRights(c, p)
	c.JSON(resp, err)
}

func roleSetRoleRights(c *bm.Context) {
	p := new(SetRoleRightsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.SetRoleRights(c, p)
	c.JSON(resp, err)
}

func roleDeleteRoleRights(c *bm.Context) {
	p := new(DeleteRoleRightsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.DeleteRoleRights(c, p)
	c.JSON(resp, err)
}

func roleDeleteRoleNullRights(c *bm.Context) {
	p := new(DeleteRoleNullRightsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.DeleteRoleNullRights(c, p)
	c.JSON(resp, err)
}

func roleGetRoleOptions(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RoleSvc.GetRoleOptions(c, p)
	c.JSON(resp, err)
}

// RegisterRoleBMServer Register the blademaster route
func RegisterRoleBMServer(e *bm.Engine, server RoleBMServer) {
	RoleSvc = server
	e.GET("/service.v1.Role/Ping", rolePing)
	e.POST("/service.v1.Role/GetRoleList", roleGetRoleList)
	e.POST("/service.v1.Role/AddRole", roleAddRole)
	e.POST("/service.v1.Role/UpdateRole", roleUpdateRole)
	e.POST("/service.v1.Role/DeleteRole", roleDeleteRole)
	e.POST("/service.v1.Role/GetRoleRights", roleGetRoleRights)
	e.POST("/service.v1.Role/SetRoleRights", roleSetRoleRights)
	e.POST("/service.v1.Role/DeleteRoleRights", roleDeleteRoleRights)
	e.POST("/service.v1.Role/DeleteRoleNullRights", roleDeleteRoleNullRights)
	e.POST("/service.v1.Role/GetRoleOptions", roleGetRoleOptions)
}

// MenuBMServer is the server API for Menu service.
type MenuBMServer interface {
	// `method:"POST"`
	GetMenus(ctx context.Context, req *GetMenusReq) (resp *GetMenusResp, err error)

	// `method:"POST"`
	GetChildrenMenuList(ctx context.Context, req *GetChildrenMenuListReq) (resp *GetChildrenMenuListResp, err error)
}

var MenuSvc MenuBMServer

func menuGetMenus(c *bm.Context) {
	p := new(GetMenusReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MenuSvc.GetMenus(c, p)
	c.JSON(resp, err)
}

func menuGetChildrenMenuList(c *bm.Context) {
	p := new(GetChildrenMenuListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := MenuSvc.GetChildrenMenuList(c, p)
	c.JSON(resp, err)
}

// RegisterMenuBMServer Register the blademaster route
func RegisterMenuBMServer(e *bm.Engine, server MenuBMServer) {
	MenuSvc = server
	e.POST("/service.v1.Menu/GetMenus", menuGetMenus)
	e.POST("/service.v1.Menu/GetChildrenMenuList", menuGetChildrenMenuList)
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
