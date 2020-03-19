package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"

	"account/internal/model"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
)

const (
	//表
	_userTable = "sys_user"

	//SQL语句
	//查询用户账户
	_queryAccount = "select `id`,`account`,`password` from %s where `account` = ?"
	//查询用户信息
	_queryUserInfo = "select `avatar`,`account`,`name`,`birthday`,`sex`,`email`,`phone`,`deptid`,`status`,`createtime` from %s where `id` = ?"
)

func NewDB() (db *sql.DB, cf func(), err error) {
	var (
		cfg sql.Config
		ct  paladin.TOML
	)
	if err = paladin.Get("db.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Mysql").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(&cfg)
	cf = func() { db.Close() }
	return
}

func (d *dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	// get data from db
	return
}

func (d *dao) RawUserAccount(ctx context.Context, account string) (u *model.User, err error) {
	sql := fmt.Sprintf(_queryAccount, _userTable)
	rows, err := d.db.Query(ctx, sql, account)
	if err != nil {
		return
	}
	defer rows.Close()
	u = &model.User{}
	rows.Next()
	err = rows.Scan(&u.ID, &u.Account, &u.Password)
	if err != nil {
		log.Error("[dao.dao-anchor.mysql|dbUser] scan short id record error(%v)", err)
		return nil, err
	}
	return
}

func (d *dao) RawUserInfo(ctx context.Context, userId int64) (u *model.User, err error) {
	sql := fmt.Sprintf(_queryUserInfo, _userTable)
	rows, err := d.db.Query(ctx, sql, userId)
	if err != nil {
		return
	}
	defer rows.Close()
	u = &model.User{}
	rows.Next()
	err = rows.Scan(&u.Avatar, &u.Account, &u.Name, &u.Birthday,
		&u.Sex, &u.Email, &u.Phone, &u.DeptId, &u.Status, &u.CreateTime)
	if err != nil {
		log.Error("[dao.dao-anchor.mysql|dbUser] scan short id record error(%v)", err)
		return nil, err
	}
	return
}
