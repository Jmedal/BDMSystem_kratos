package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"time"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/google/wire"
	pb "role/api"
	"role/internal/model"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	Article(c context.Context, id int64) (*model.Article, error)
	RawRoleList(ctx context.Context) (resp *pb.GetRoleListResp, err error)
	InsertRole(ctx context.Context, req *pb.AddRoleReq) (resp *pb.AddRoleResp, err error)
	UpdateRole(ctx context.Context, req *pb.UpdateRoleReq) (resp *pb.UpdateRoleResp, err error)
	DeleteRole(ctx context.Context, req *pb.DeleteRoleReq) (resp *pb.DeleteRoleResp, err error)
	GetRoleMenus(ctx context.Context, req *pb.GetRoleRightsReq) (resp *pb.GetRoleRightsResp, err error)
	SetRoleMenus(ctx context.Context, req *pb.SetRoleRightsReq) (resp *pb.SetRoleRightsResp, err error)
	DeleteRoleMenus(ctx context.Context, req *pb.DeleteRoleRightsReq) (resp *pb.DeleteRoleRightsResp, err error)
	DeleteRoleNullMenus(ctx context.Context, req *pb.DeleteRoleNullRightsReq) (resp *pb.DeleteRoleNullRightsResp, err error)
}

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	demoExpire int32
	menuClient pb.MenuClient
}

// New new a dao and return.
func New(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, mc, db)
}

func newDao(r *redis.Redis, mc *memcache.Memcache, db *sql.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:         db,
		redis:      r,
		mc:         mc,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	go initMenuClient(d)
	cf = d.Close
	return
}

//初始化菜单服务客户端
func initMenuClient(d *dao) (err error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 60),
		Timeout:           xtime.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
	}
	var menuClient pb.MenuClient
	if menuClient, err = pb.RegisterMenuClient(grpccfg); err != nil {
		log.Error("RegisterMenuClient error:(%v)", err)
	}
	d.menuClient = menuClient
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}
