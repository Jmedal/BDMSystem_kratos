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
	pb "menu/api"
	"menu/internal/model"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	Article(c context.Context, id int64) (*model.Article, error)
	RawMenuList(ctx context.Context, pid int64, levels int64) (resp *pb.GetMenuListResp, err error)
	InsertMenu(ctx context.Context, req *pb.AddMenuReq) (resp *pb.AddMenuResp, err error)
	UpdateMenu(c context.Context, req *pb.UpdateMenuReq) (resp *pb.UpdateMenuResp, err error)
	DeleteMenu(c context.Context, req *pb.DeleteMenuReq) (resp *pb.DeleteMenuResp, err error)
	RawMenus(ctx context.Context, pid int64, levels int64, req *pb.GetMenusReq) (resp *pb.GetMenusResp, err error)
	RawMenuOptions(ctx context.Context, pid int64, levels int64, req *pb.GetMenuOptionsReq) (resp *pb.GetMenuOptionsResp, err error)
	RawAllMenuOptions(ctx context.Context, pid int64, levels int64) (resp *pb.GetMenuOptionsResp, err error)
	RawChildrenMenuList(ctx context.Context, req *pb.GetChildrenMenuListReq) (resp *pb.GetChildrenMenuListResp, err error)
}

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	demoExpire int32
	roleClient pb.RoleClient
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
	go initRoleClient(d)
	cf = d.Close
	return
}

//初始化角色服务客户端
func initRoleClient(d *dao) (err error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 60),
		Timeout:           xtime.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
	}
	var roleClient pb.RoleClient
	if roleClient, err = pb.RegisterRoleClient(grpccfg); err != nil {
		log.Error("RegisterRoleClient error:(%v)", err)
	}
	d.roleClient = roleClient
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
