// Code generated by kratos tool genmc. DO NOT EDIT.

/*
  Package dao is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// mc: -key=keyArt -type=get
		CacheArticle(c context.Context, id int64) (*model.Article, error)
		// mc: -key=keyArt -expire=d.demoExpire
		AddCacheArticle(c context.Context, id int64, art *model.Article) (err error)
		// mc: -key=keyArt
		DeleteArticleCache(c context.Context, id int64) (err error)
	}
*/

package dao

import (
	"context"
	"fmt"

	"github.com/bilibili/kratos/pkg/cache/memcache"
	"github.com/bilibili/kratos/pkg/log"
	"menu/internal/model"
)

var _ _mc

// CacheArticle get data from mc
func (d *dao) CacheArticle(c context.Context, id int64) (res *model.Article, err error) {
	key := keyArt(id)
	res = &model.Article{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheArticle", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheArticle Set data to mc
func (d *dao) AddCacheArticle(c context.Context, id int64, val *model.Article) (err error) {
	if val == nil {
		return
	}
	key := keyArt(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheArticle", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DeleteArticleCache delete data from mc
func (d *dao) DeleteArticleCache(c context.Context, id int64) (err error) {
	key := keyArt(id)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DeleteArticleCache", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
