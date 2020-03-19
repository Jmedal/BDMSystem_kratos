package dao

import (
	"context"
	"time"
	"token/internal/model"

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

	RequestUploadToken(ctx context.Context, userId int64, operator string, now, exp int64) (token, randomKey string, err error)

	VerifyUploadToken(ctx context.Context, token string) (userId int64, randomKey string, err error)
}

// dao dao.
type dao struct {
	db          *sql.DB
	redis       *redis.Pool
	mc          *memcache.Memcache
	cache       *fanout.Fanout
	cacheExpire int32
}

// New new a dao and return.
func New(r *redis.Pool, mc *memcache.Memcache, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, mc, db)
}

func newDao(r *redis.Pool, mc *memcache.Memcache, db *sql.DB) (d *dao, cf func(), err error) {
	var cfg struct {
		CacheExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:          db,
		redis:       r,
		mc:          mc,
		cache:       fanout.New("cache"),
		cacheExpire: int32(time.Duration(cfg.CacheExpire) / time.Second),
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
