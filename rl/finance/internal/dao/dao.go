package dao

import (
	"context"
	pb "finance/api"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"time"

	"finance/internal/model"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, NewDB, NewRedis, NewMC)

//go:generate kratos tool genbts
// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	Article(c context.Context, id int64) (*model.Article, error)
	RawProjectPage(ctx context.Context, req *pb.GetProjectPageReq) (resp *pb.GetProjectPageResp, err error)
	InsertProject(ctx context.Context, req *pb.AddProjectReq) (resp *pb.AddProjectResp, err error)
	UpdateProject(ctx context.Context, req *pb.UpdateProjectReq) (resp *pb.UpdateProjectResp, err error)
	FrozenStateProject(ctx context.Context, req *pb.TerminationProjectReq) (resp *pb.TerminationProjectResp, err error)
	ActivityStateProject(ctx context.Context, req *pb.RestartProjectReq) (resp *pb.RestartProjectResp, err error)
	RawBillPage(ctx context.Context, req *pb.GetBillPageReq) (resp *pb.GetBillPageResp, err error)
	InsertBill(ctx context.Context, req *pb.AddBillReq) (resp *pb.AddBillResp, err error)
	UpdateBill(ctx context.Context, req *pb.UpdateBillReq) (resp *pb.UpdateBillResp, err error)
	FrozenStateBill(ctx context.Context, req *pb.DeleteBillReq) (resp *pb.DeleteBillResp, err error)
}

// dao dao.
type dao struct {
	db            *sql.DB
	redis         *redis.Redis
	mc            *memcache.Memcache
	cache         *fanout.Fanout
	demoExpire    int32
	accountClient pb.AccountClient
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
	go initAccountClient(d)
	cf = d.Close
	return
}

//初始化token注册客户端
func initAccountClient(d *dao) (err error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 60),
		Timeout:           xtime.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
	}
	var accountClient pb.AccountClient
	if accountClient, err = pb.RegisterAccountClient(grpccfg); err != nil {
		log.Error("RegisterAccountClient error:(%v)", err)
	}
	d.accountClient = accountClient
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
