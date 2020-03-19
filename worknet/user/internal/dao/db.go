package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	"strconv"
	"time"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	pb "user/api"
	"user/internal/model"
)

const (
	//标识
	//激活
	_activityLabel = "1"
	//未激活
	_inactivityLabel = "0"
	//男
	_manLabel = "男"
	//女
	_humanLabel = "女"

	//表名
	//用户表
	_userTable = "sys_user"
	//普通用户表
	_commonUserTable = "sys_learner_info"
	//教师用户表
	_teacherTable = "sys_teacher_info"
	//公司用户表
	_companyTable = "sys_company"

	//SQL语句
	//查询某类用户数量
	_queryUserNumber = "select count(*) AS number from %s"
	//查询是否激活用户数量
	_queryActivityUserNumber = "select count(*) AS number from  %s where activity = ?"
	//查询用户性别数量
	_querySexUserNumber = "select count(*) AS number from %s where sex = ?"

	//查询首个注册用户的注册时间
	_queryFirstRegisterTime = "select MIN(createtime) from %s "

	//查询某年某类用户数量
	_queryYearAndTypeRegisterUserNumber = "select count(*) AS number from %s where user_id in(select id from %s where  createtime >= ? and createtime < ?)"
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

//查询普通用户，教师用户，公司用户数量
func (d *dao) RawUserNumber(ctx context.Context) (resp *pb.WorknetUserNumberResp, err error) {
	sql := make([]string, 0, 3)
	sql = append(sql, fmt.Sprintf(_queryUserNumber, _commonUserTable))
	sql = append(sql, fmt.Sprintf(_queryUserNumber, _teacherTable))
	sql = append(sql, fmt.Sprintf(_queryUserNumber, _companyTable))
	num := make([]int64, 3, 3)
	for i, v := range sql {
		if err = d.db.QueryRow(ctx, v).Scan(&num[i]); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_learner_info|sys_teacher_info|sys_company]] scan record error(%v)", err)
			return
		}
	}
	resp = &pb.WorknetUserNumberResp{
		CommonUser:  num[0],
		TeacherUser: num[1],
		CompanyUser: num[2],
	}
	return
}

//查询激活和未激活用户数量
func (d *dao) RawActivityUserNumber(ctx context.Context) (resp *pb.WorknetUserActivityResp, err error) {
	sql := fmt.Sprintf(_queryActivityUserNumber, _userTable)
	num := make([]int64, 2, 2)
	for i := 0; i < 2; i++ {
		if err = d.db.QueryRow(ctx, sql, i).Scan(&num[i]); err != nil {
			log.Error("[dao.dao-anchor.mysql|dbUser] scan record error(%v)", err)
			return
		}
	}
	resp = &pb.WorknetUserActivityResp{
		Activity:   num[1],
		Inactivity: num[0],
	}
	return
}

//查询男性和女性数量用户数量
func (d *dao) RawSexUserNumber(ctx context.Context) (resp *pb.WorknetUserSexResp, err error) {
	sql := fmt.Sprintf(_querySexUserNumber, _commonUserTable)
	sex := []string{_manLabel, _humanLabel}
	num := make([]int64, 2, 2)
	for i, v := range sex {
		if err = d.db.QueryRow(ctx, sql, v).Scan(&num[i]); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_learner_info]] scan record error(%v)", err)
			return
		}
	}
	resp = &pb.WorknetUserSexResp{
		Man:   num[0],
		Human: num[1],
	}
	return
}

//查询普通用户，教师用户，公司用户数量变化
func (d *dao) RawUserNumberChange(ctx context.Context) (resp *pb.WorknetUserChangeResp, err error) {
	var year xtime.Time
	sql := fmt.Sprintf(_queryFirstRegisterTime, _userTable)
	if err = d.db.QueryRow(ctx, sql).Scan(&year); err != nil {
		log.Error("[dao.dao-anchor.mysql|dbUser] scan record error(%v)", err)
		return
	}
	firstYear := year.Time().Format("2006")
	nowYear := time.Now().Format("2006")
	now, err := strconv.ParseInt(nowYear, 10, 64)
	if err != nil {
		log.Error("字符串转换成整数失败")
	}
	first, err := strconv.ParseInt(firstYear, 10, 64)
	if err != nil {
		log.Error("字符串转换成整数失败")
		first = now
	}
	sqlQuery := make([]string, 0, 3)
	sqlQuery = append(sqlQuery, fmt.Sprintf(_queryYearAndTypeRegisterUserNumber, _commonUserTable, _userTable)) //普通用户
	sqlQuery = append(sqlQuery, fmt.Sprintf(_queryYearAndTypeRegisterUserNumber, _companyTable, _userTable))    //公司用户
	sqlQuery = append(sqlQuery, fmt.Sprintf(_queryYearAndTypeRegisterUserNumber, _teacherTable, _userTable))    //老师用户
	resp = &pb.WorknetUserChangeResp{}
	size := now-first+1
	resp.CommonUser = make([]int64, size, size)
	resp.CompanyUser = make([]int64, size, size)
	resp.TeacherUser = make([]int64, size, size)
	for i := first; i <= now; i++ {
		resp.Time = append(resp.Time, strconv.FormatInt(i, 10))
		num := make([]int64, 3, 3)
		for j, v := range sqlQuery {
			if err = d.db.QueryRow(ctx, v, strconv.FormatInt(i, 10), strconv.FormatInt(i+1, 10)).Scan(&num[j]); err != nil {
				log.Error("[dao.dao-anchor.mysql|dbUser] scan short id record error(%v)", err)
				return
			}
		}
		resp.CommonUser[i-first] = num[0]
		resp.TeacherUser[i-first] = num[1]
		resp.CompanyUser[i-first] = num[2]
	}
	return
}
