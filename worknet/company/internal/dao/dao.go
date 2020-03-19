package dao

import (
	pb "company/api"
	"company/internal/model"
	"context"
	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"
	"time"

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
	RawLocationNumber(ctx context.Context) (resp *pb.WorknetLocationListResp, err error)
	RawCompanyList(ctx context.Context) (resp *pb.WorknetCompanyListResp, err error)
	RawCompanyProfessionNumberChange(ctx context.Context, req *pb.WorknetProfessionChangeReq) (resp *pb.WorknetProfessionChangeResp, err error)
	RawCompanyProfessionRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error)
	RawCompanyCvRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error)
	RawCompanyCvStatusNumber(ctx context.Context) (resp *pb.WorknetCvStatusNumberResp, err error)
	RawCompanyContestNumberChange(ctx context.Context, req *pb.WorknetContestChangeReq) (resp *pb.WorknetContestChangeResp, err error)
	RawCompanyContestRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error)
	RawCompanyContestingRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error)
	RawCompanyContestApplyRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error)
}

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	mc         *memcache.Memcache
	cache      *fanout.Fanout
	demoExpire int32
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
	cf = d.Close
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
