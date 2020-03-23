package dao

import (
	pb "account/api"
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/pkg/errors"
	"time"

	"account/internal/model"
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
	RawUser(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error)
	RawUserPage(ctx context.Context, req *pb.GetUserPageReq) (resp *pb.GetUserPageResp, err error)
	InsertUser(ctx context.Context, req *pb.AddUserReq) (resp *pb.AddUserResp, err error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (resp *pb.UpdateUserResp, err error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (resp *pb.DeleteUserResp, err error)
	SetUserStatus(ctx context.Context, req *pb.SetUserStatusReq) (resp *pb.SetUserStatusResp, err error)
	SetUserRole(ctx context.Context, req *pb.SetUserRoleReq) (resp *pb.SetUserRoleResp, err error)
	RawAccount(ctx context.Context, req *pb.CheckAccountReq) (resp *pb.CheckAccountResp, err error)
	VerifyToken(ctx context.Context, token string) (userId int64, randomKey string, err error)
}

// dao dao.
type dao struct {
	db          *sql.DB
	redis       *redis.Redis
	mc          *memcache.Memcache
	cache       *fanout.Fanout
	demoExpire  int32
	tokenClient pb.TokenClient
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
	go initTokenClient(d)
	cf = d.Close
	return
}

//初始化token注册客户端
func initTokenClient(d *dao) (err error) {
	grpccfg := &warden.ClientConfig{
		Dial:              xtime.Duration(time.Second * 10),
		Timeout:           xtime.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: xtime.Duration(time.Second * 60),
		KeepAliveTimeout:  xtime.Duration(time.Second * 20),
	}
	var tokenClient pb.TokenClient
	if tokenClient, err = pb.NewClient(grpccfg); err != nil {
		log.Error("NewClient error:(%v)", err)
	}
	d.tokenClient = tokenClient
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

func (d *dao) VerifyToken(ctx context.Context, token string) (userId int64, randomKey string, err error) {
	req := &pb.VerifyTokenReq{
		AccessToken: token,
	}
	var resp = &pb.VerifyTokenResp{}
	if resp, err = d.tokenClient.Verify(ctx, req); err != nil {
		err = errors.Wrapf(err, "%s", req.AccessToken)
	}
	userId = resp.UserId
	randomKey = resp.RandomKey
	return
}
