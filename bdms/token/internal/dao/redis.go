package dao

import (
	"context"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
)

func NewRedis() (r *redis.Pool, cf func(), err error) {
	var (
		cfg redis.Config
		ct  paladin.Map
	)
	if err = paladin.Get("redis.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Redis").UnmarshalTOML(&cfg); err != nil {
		return
	}
	r = redis.NewPool(&cfg)
	cf = func() { r.Close() }
	return
}

func (d *dao) PingRedis(ctx context.Context) (err error) {
	conn := d.redis.Get(ctx)
	if _, err = conn.Do("SET", "ping", "pong"); err != nil {
		log.Error("conn.Set(PING) error(%v)", err)
	}
	return
}
