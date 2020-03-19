package dao

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/log"
	"math/rand"
	"time"
)

const (
	_defaultExpiration = 24 * 60 * 60 //24h

	_randomKeyLength = 8

	_tokenNamespace = "bdms_token:"

	_tokenRandomKeyNamespace = "bdms_token_randomKey:"
)

// RequestUploadToken generates a token for subsequent upload.
// Token will expire in a specific duration.
func (d *dao) RequestUploadToken(ctx context.Context, userId int64, operator string, now, exp int64) (token, randomKey string, err error) {
	token = genToken(userId, operator, now)
	randomKey = randomString(_randomKeyLength)

	conn := d.redis.Get(ctx)
	defer conn.Close()

	if exp <= 0 {
		exp = _defaultExpiration
	}
	//token
	nsToken := namespaceToken(token)
	//randomKey
	nsTokenRandomkey := namespaceTokenRandomKey(token)
	if _, err = conn.Do("SETEX", nsToken, exp, userId); err != nil {
		log.Error("conn.Do(SETEX %s %d) failure(%v)", nsToken, exp, err)
	}
	if _, err = conn.Do("SETEX", nsTokenRandomkey, exp, randomKey); err != nil {
		log.Error("conn.Do(SETEX %s %d) failure(%v)", nsTokenRandomkey, exp, err)
	}
	return
}

// VerifyUploadToken verifies if a token is legal.
func (d *dao) VerifyUploadToken(ctx context.Context, token string) (userId int64, randomKey string, err error) {
	conn := d.redis.Get(ctx)
	defer conn.Close()

	nsToken := namespaceToken(token)
	if userId, err = redis.Int64(conn.Do("GET", nsToken)); err != nil {
		if err == redis.ErrNil {
			err = nil
		} else {
			log.Warn("conn.Do(GET %s) failure(%v)", nsToken, err)
		}
		userId = -1
	}
	nsTokenRandomkey := namespaceTokenRandomKey(token)
	if randomKey, err = redis.String(conn.Do("GET", nsTokenRandomkey)); err != nil {
		if err == redis.ErrNil {
			err = nil
		} else {
			log.Warn("conn.Do(GET %s) failure(%v)", nsTokenRandomkey, err)
		}
		randomKey = ""
	}
	return
}

func genToken(userId int64, operator string, now int64) string {
	sha := sha1.New()
	sha.Write([]byte(fmt.Sprintf("i love bdms:%d:%s:%d", userId, operator, now)))
	return fmt.Sprintf("%s:%d", hex.EncodeToString(sha.Sum([]byte(""))), now)
}

// Avoid key collision.
func namespaceToken(token string) string {
	return _tokenNamespace + token
}

// Avoid key collision.
func namespaceTokenRandomKey(token string) string {
	return _tokenRandomKeyNamespace + token
}

//create randomString
func randomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
