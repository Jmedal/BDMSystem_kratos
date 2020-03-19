package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	pb "profession/api"
	"profession/internal/model"
	"strconv"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
)

const (
	//表名
	//用户选择职业表
	_userProfessionTable = "sys_user_profession"
	//职业表
	_professionTable = "sys_profession"
	//职业大类表
	_professionTypeTable = "sys_profession_type"

	//sql语句
	//查询用户选择职业数量
	_queryUserProfessionNumber = "select name, count(*) as number from %s inner join %s on sys_profession.id = sys_user_profession.profession_id group by sys_profession.id"

	//查询职业薪水
	_queryProfessionSalary = "select name, salary from %s"

	//查询职业大类拥有职业数量
	_queryProfessionTypeNumber = "select type_name, count(*) as number from %s inner join %s on sys_profession_type.id = sys_profession.type_id  group by sys_profession_type.id"

	//查询职业大类平均薪水情况
	_queryProfessionTypeSalary = "select type_name, avg(salary) from %s inner join %s on sys_profession_type.id = sys_profession.type_id  group by sys_profession_type.id"
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

//查询用户选择职业数量列表
func (d *dao) RawProfessionChoice(ctx context.Context) (resp *pb.WorknetProfessionChoiceResp, err error) {
	sql := fmt.Sprintf(_queryUserProfessionNumber, _professionTable, _userProfessionTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetProfessionChoiceResp{}
	for rows.Next() {
		var name string
		var number int64
		if err = rows.Scan(&name, &number); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_profession|sys_user_profession]] scan all record error(%v)", err)
			return
		}
		resp.ProfessionName = append(resp.ProfessionName, name)
		resp.Number = append(resp.Number, strconv.FormatInt(number, 10))
	}
	return
}

//查询职业薪水列表
func (d *dao) RawProfessionSalary(ctx context.Context) (resp *pb.WorknetProfessionSalaryResp, err error) {
	sql := fmt.Sprintf(_queryProfessionSalary, _professionTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetProfessionSalaryResp{}
	for rows.Next() {
		var name string
		var salary float32
		if err = rows.Scan(&name, &salary); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_profession]] scan all record error(%v)", err)
			return
		}
		resp.ProfessionName = append(resp.ProfessionName, name)
		resp.Salary = append(resp.Salary, fmt.Sprintf("%.1f", salary/1000))
	}
	return
}

//查询职业大类拥有职业数量列表
func (d *dao) RawProfessionTypeNumber(ctx context.Context) (resp *pb.WorknetProfessionChoiceResp, err error) {
	sql := fmt.Sprintf(_queryProfessionTypeNumber, _professionTypeTable, _professionTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetProfessionChoiceResp{}
	for rows.Next() {
		var name string
		var number int64
		if err = rows.Scan(&name, &number); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_profession_type|sys_profession]] scan all record error(%v)", err)
			return
		}
		resp.ProfessionName = append(resp.ProfessionName, name)
		resp.Number = append(resp.Number, strconv.FormatInt(number, 10))
	}
	return
}

//查询职业大类薪水列表
func (d *dao) RawProfessionTypeSalary(ctx context.Context) (resp *pb.WorknetProfessionSalaryResp, err error) {
	sql := fmt.Sprintf(_queryProfessionTypeSalary, _professionTypeTable, _professionTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetProfessionSalaryResp{}
	for rows.Next() {
		var name string
		var salary float32
		if err = rows.Scan(&name, &salary); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_profession]] scan all record error(%v)", err)
			return
		}
		resp.ProfessionName = append(resp.ProfessionName, name)
		resp.Salary = append(resp.Salary, fmt.Sprintf("%.1f", salary/1000))
	}
	return
}
