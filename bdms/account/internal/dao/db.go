package dao

import (
	pb "account/api"
	"account/internal/model"
	"context"
	dsql "database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
)

const (
	//帐号启用
	_AccountEnable = 1
	//帐号冻结
	_AccountFrozen = 2

	//表
	_sysUserTable = "sys_user"

	//SQL语句
	//查询用户账户
	_queryUser = "select account, password, id, avatar, name, role_id, birthday, sex, email, phone, status, create_time from %s where account = ?"

	//查询用户信息列表
	_queryUserList = "select id, avatar, account, password, name, role_id, birthday, sex, email, phone, status, create_time from %s where account like ? LIMIT ?, ?"

	//查询用户总数
	_queryUserCount = "select count(*) from %s where account like ?"

	//插入用户
	_insertUser = "insert into %s (avatar, account, password, name, birthday, sex, email, phone, create_time) values (?,?,?,?,?,?,?,?,?)"

	//更新用户
	_updateUser      = "update %s set avatar = ?, account = ?, password = ?, name = ?, birthday = ?, sex = ?, email = ?, phone = ? where id = ?"
	_updateUserNoPas = "update %s set avatar = ?, account = ?, name = ?, birthday = ?, sex = ?, email = ?, phone = ? where id = ?"

	//删除用户
	_deleteUser = "delete from %s where id = ?"

	//更新用户角色
	_updateUserStatus = "update %s set status = ? where id = ?"

	//更新用户角色
	_updateUserRole = "update %s set role_id = ? where id = ?"

	//查询用户帐号
	_queryAccountCount = "select count(*) from %s where account = ?"

	//查询用户姓名
	_queryUserName = "select name, account from %s where id = ?"

	//查询用户角色
	_queryUserRole = "select role_id from %s where id = ?"

	//查询用户姓名
	_queryUserNameList = "select id, name, account from %s"
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

//用户登录
func (d *dao) RawUser(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	sqlUser := fmt.Sprintf(_queryUser, _sysUserTable)
	resp = &pb.LoginResp{}
	userInfo := &pb.UserInfo{}
	var password string
	var birthday, createTime xtime.Time
	if err = d.db.QueryRow(ctx, sqlUser, req.Account).Scan(&userInfo.Account, &password, &userInfo.Id,
		&userInfo.Avatar, &userInfo.Name, &userInfo.RoleId, &birthday,
		&userInfo.Sex, &userInfo.Email, &userInfo.Phone, &userInfo.Status, &createTime); err != nil && err != sql.ErrNoRows {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] scan short id record error(%v)", err)
		resp.Result = "db server error"
		return
	}
	userInfo.Birthday = birthday.Time().Unix()
	userInfo.CreateTime = createTime.Time().Unix()
	if err == sql.ErrNoRows {
		resp.Result = "Account aren't existent"
		return
	}
	if userInfo.Status != _AccountEnable {
		resp.Result = "Account are frozen"
		return
	}
	if password == req.Password {
		var expire int64
		//是否保存密码 保存：1，不保存：2
		if req.Save == "1" {
			expire = 365 * 24 * 60 * 60 //一年
		} else {
			expire = 24 * 60 * 60 //一天
		}
		req := &pb.NewTokenReq{
			UserId:   userInfo.Id,
			Operator: strconv.FormatInt(userInfo.RoleId, 10),
			Expire:   expire,
		}
		var tokenResp = &pb.NewTokenResp{}
		if tokenResp, err = d.tokenClient.Request(ctx, req); err != nil {
			log.Error("d.tokenClient.Request error(%v)", err)
			resp.Result = "Token server error"
			return
		}
		resp.AccessToken = tokenResp.AccessToken
		resp.RandomKey = tokenResp.RandomKey
		resp.UserInfo = userInfo
		resp.Result = "success"
		return
	}
	resp.Result = "Password error"
	return
}

//查询用户信息列表
func (d *dao) RawUserPage(ctx context.Context, req *pb.GetUserPageReq) (resp *pb.GetUserPageResp, err error) {
	sqlUserList := fmt.Sprintf(_queryUserList, _sysUserTable)
	sqlAccountCount := fmt.Sprintf(_queryUserCount, _sysUserTable)
	query := fmt.Sprintf("%%%s%%", req.Query)
	resp = &pb.GetUserPageResp{}
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	if err = d.db.QueryRow(ctx, sqlAccountCount, query).Scan(&resp.Total); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] scan all record error(%v)", err)
		return
	}
	start := (req.PageNum - 1) * req.PageSize
	size := resp.PageSize
	rows, err := d.db.Query(ctx, sqlUserList, query, start, size)
	if err != nil {
		log.Error("select user error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var birthday, createTime xtime.Time
		userInfo := &pb.UserInfo{}
		if err = rows.Scan(&userInfo.Id, &userInfo.Avatar, &userInfo.Account, &userInfo.Password,
			&userInfo.Name, &userInfo.RoleId, &birthday, &userInfo.Sex, &userInfo.Email,
			&userInfo.Phone, &userInfo.Status, &createTime); err != nil {
			log.Error("[dao.dao-anchor.mysql|sys_user] scan short id record error(%v)", err)
			return
		}
		userInfo.Birthday = birthday.Time().Unix()
		userInfo.CreateTime = createTime.Time().Unix()
		resp.UserList = append(resp.UserList, userInfo)
	}
	return
}

//添加用户
func (d *dao) InsertUser(ctx context.Context, req *pb.AddUserReq) (resp *pb.AddUserResp, err error) {
	sqlUser := fmt.Sprintf(_insertUser, _sysUserTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlUser, req.User.Avatar, req.User.Account,
		req.User.Password, req.User.Name, xtime.Time(req.User.Birthday).Time(), req.User.Sex,
		req.User.Email, req.User.Phone, time.Now()); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] failed to insert error(%v)", err)
		return
	}
	resp = &pb.AddUserResp{}
	if _, err = res.LastInsertId(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//更新用户
func (d *dao) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (resp *pb.UpdateUserResp, err error) {
	var res dsql.Result
	if req.User.Password != "" {
		sqlUser := fmt.Sprintf(_updateUser, _sysUserTable)
		if res, err = d.db.Exec(ctx, sqlUser, req.User.Avatar, req.User.Account, req.User.Password,
			req.User.Name, xtime.Time(req.User.Birthday).Time(),
			req.User.Sex, req.User.Email, req.User.Phone, req.Id); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_user]] failed to update: (%v), error(%v)", req.Id, err)
		}
	} else {
		sqlUser := fmt.Sprintf(_updateUserNoPas, _sysUserTable)
		if res, err = d.db.Exec(ctx, sqlUser, req.User.Avatar, req.User.Account,
			req.User.Name, xtime.Time(req.User.Birthday).Time(),
			req.User.Sex, req.User.Email, req.User.Phone, req.Id); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_user]] failed to update: (%v), error(%v)", req.Id, err)
		}
	}
	resp = &pb.UpdateUserResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除用户
func (d *dao) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (resp *pb.DeleteUserResp, err error) {
	sqlUser := fmt.Sprintf(_deleteUser, _sysUserTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlUser, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] failed to delete: (%v), error(%v)", req.Id, err)
		return
	}
	resp = &pb.DeleteUserResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//设置用户状态
func (d *dao) SetUserStatus(ctx context.Context, req *pb.SetUserStatusReq) (resp *pb.SetUserStatusResp, err error) {
	sqlUserRole := fmt.Sprintf(_updateUserStatus, _sysUserTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlUserRole, req.Status, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] failed to update: (%v), (%v), error(%v)", req.Status, req.Id)
	}
	resp = &pb.SetUserStatusResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//设置用户角色
func (d *dao) SetUserRole(ctx context.Context, req *pb.SetUserRoleReq) (resp *pb.SetUserRoleResp, err error) {
	sqlUserRole := fmt.Sprintf(_updateUserRole, _sysUserTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlUserRole, req.RoleId, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] failed to update: (%v), (%v), error(%v)", req.RoleId, req.Id, err)
	}
	resp = &pb.SetUserRoleResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//检查用户帐号是否存在
func (d *dao) RawAccount(ctx context.Context, req *pb.CheckAccountReq) (resp *pb.CheckAccountResp, err error) {
	sqlAccount := fmt.Sprintf(_queryAccountCount, _sysUserTable)
	var count int
	if err = d.db.QueryRow(ctx, sqlAccount, req.Account).Scan(&count); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] scan short id record error(%v)", err)
		return
	}
	resp = &pb.CheckAccountResp{}
	if count == 0 {
		resp.Result = "success"
	} else {
		resp.Result = "error"
	}
	return
}

//查询用户姓名[帐号]
func (d *dao) RawUserName(ctx context.Context, req *pb.GetUserNameReq) (resp *pb.GetUserNameResp, err error) {
	sqlUserName := fmt.Sprintf(_queryUserName, _sysUserTable)
	resp = &pb.GetUserNameResp{}
	var name, account string
	if err = d.db.QueryRow(ctx, sqlUserName, req.Id).Scan(&name, &account); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] scan short id record error(%v)", err)
		return
	}
	resp.Name = fmt.Sprintf("%s[%s]", name, account)
	return
}

//查询用户角色
func (d *dao) RawUserRole(ctx context.Context, req *pb.GetUserRoleReq) (resp *pb.GetUserRoleResp, err error) {
	sqlUserName := fmt.Sprintf(_queryUserRole, _sysUserTable)
	resp = &pb.GetUserRoleResp{}
	if err = d.db.QueryRow(ctx, sqlUserName, req.Id).Scan(&resp.RoleId); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[sys_user]] scan short id record error(%v)", err)
		return
	}
	return
}

//查询用户姓名
func (d *dao) RawUserNameOptions(ctx context.Context) (resp *pb.GetUserNameOptionsResp, err error) {
	sqlUserName := fmt.Sprintf(_queryUserNameList, _sysUserTable)
	resp = &pb.GetUserNameOptionsResp{}
	rows, err := d.db.Query(ctx, sqlUserName)
	if err != nil {
		log.Error("select user error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := &pb.GetUserNameOptionsResp_UserOption{}
		var name, account string
		if err = rows.Scan(&user.Id, &name, &account); err != nil {
			log.Error("[dao.dao-anchor.mysql|sys_user] scan short id record error(%v)", err)
			return
		}
		user.Name = fmt.Sprintf("%s[%s]", name, account)
		resp.UserOptions = append(resp.UserOptions, user)
	}
	return
}
