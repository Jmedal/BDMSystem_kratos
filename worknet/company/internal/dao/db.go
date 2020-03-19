package dao

import (
	pb "company/api"
	"company/internal/model"
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	"time"
)

const (
	//表名
	//公司表
	_companyTable = "sys_company"
	//公司招聘表
	_companyProfessionTable = "sys_company_profession"
	//公司简历表表
	_companyCvTable = "sys_company_cv"
	//公司笔试表
	_companyContestTable = "sys_company_contest"
	//公司笔试报名表
	_companyContestApplyTable = "sys_company_contest_apply"

	//SQL语句
	//查询各地互联网公司分布数量
	_queryLocationNumber = "select sys_company.address, count(*) as number from %s group by sys_company.address order by number desc"

	//查询公司列表
	_queryCompanyList = "select id, name from %s"

	//查询公司名称
	_queryCompanyName = "select name from %s where id = ?"

	//查询某年某月中某个公司新增公司招聘数量										                                                       '2019-6-1' '2020-1-1'
	_queryCompanyProfessionNumber = "select count(*) as number from %s where sys_company_profession.company_id = ? and start_time between ? and ? "

	//查询各个公司的公司招聘数量
	_queryCompanyProfessionList = "select sys_company.name, count(*) as number from %s inner join %s on sys_company_profession.company_id = sys_company.id group by sys_company.id order by number desc"

	//查询各个公司的公司简历数量
	_queryCompanyCvList = "select sys_company.name, count(*) as number from (%s inner join %s on sys_company_cv.company_profession_id = sys_company_profession.id) inner join %s on sys_company_profession.company_id = sys_company.id group by sys_company.id order by number desc"

	//查询所有公司的公司简历状态数量
	_queryCompanyCvStatusNumber = "select sys_company_cv.status, count(*) as number from %s group by sys_company_cv.status order by sys_company_cv.status"

	//查询某年某月中某个公司新增公司笔试数量
	_queryCompanyContestNumber = "select count(*) as number from %s where sys_company_contest.company_id = ? and start_time between ? and ?"

	//查询各个公司的公司笔试数量
	_queryCompanyContestList = "select sys_company.name, count(*) as number from %s inner join %s on sys_company_contest.company_id = sys_company.id group by sys_company.id order by number desc"

	//查询各个公司的公司笔试（正在进行）数量
	_queryCompanyContestingList = "select sys_company.name, count(*) as number from %s inner join %s on sys_company_contest.company_id = sys_company.id where end_time > ? group by sys_company.id order by number desc"

	//查询各个公司的公司笔试报名数量
	_queryCompanyContestApplyList = "select sys_company.name, count(*) as number from (%s inner join %s on sys_company_contest_apply.contest_id = sys_company_contest.id) inner join %s on sys_company_contest.company_id = sys_company.id group by sys_company.id order by number desc"

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

//查询各地互联网公司分布数量
func (d *dao) RawLocationNumber(ctx context.Context) (resp *pb.WorknetLocationListResp, err error) {
	//sql语句
	sql := fmt.Sprintf(_queryLocationNumber, _companyTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetLocationListResp{}
	for rows.Next() {
		locationInfo := &pb.WorknetLocationListResp_LocationInfo{}
		if err = rows.Scan(&locationInfo.CityName, &locationInfo.CompanyNumber); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company]] scan all record error(%v)", err)
			return
		}
		resp.LocationList = append(resp.LocationList, locationInfo)
	}
	return
}

//查询公司列表
func (d *dao) RawCompanyList(ctx context.Context) (resp *pb.WorknetCompanyListResp, err error) {
	sql := fmt.Sprintf(_queryCompanyList, _companyTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCompanyListResp{}
	for rows.Next() {
		company := &pb.WorknetCompanyListResp_WorknetCompany{}
		if err = rows.Scan(&company.CompanyId, &company.CompanyName); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company]] scan all record error(%v)", err)
			return
		}
		resp.CompanyList = append(resp.CompanyList, company)
	}
	return
}

//查询某时间段某公司新增公司招聘数量变化
func (d *dao) RawCompanyProfessionNumberChange(ctx context.Context, req *pb.WorknetProfessionChangeReq) (resp *pb.WorknetProfessionChangeResp, err error) {
	resp = &pb.WorknetProfessionChangeResp{}
	sqlCompany := fmt.Sprintf(_queryCompanyName, _companyTable)
	for _, v := range req.CompanyId {
		profession := &pb.WorknetProfessionChangeResp_WorknetProfession{}
		var name string
		if err = d.db.QueryRow(ctx, sqlCompany, v).Scan(&name); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company]] scan short id record error(%v)", err)
			return
		}
		profession.CompanyName = name
		resp.Profession = append(resp.Profession, profession)
	}
	//时间区间
	first, last := getTimeDuration(xtime.Time(req.Time[0]), xtime.Time(req.Time[1]))
	//sql语句
	sqlProfession := fmt.Sprintf(_queryCompanyProfessionNumber, _companyProfessionTable)
	for ; last.After(first); first = first.AddDate(0, 1, 0) {
		resp.Time = append(resp.Time, first.Format("2006-01"))
		firstOfMonth := first.Format("2006-01-02")
		lastOfMonth := first.AddDate(0, 1, -1).Format("2006-01-02")
		for i, v := range req.CompanyId {
			var num int64
			if err = d.db.QueryRow(ctx, sqlProfession, v, firstOfMonth, lastOfMonth).Scan(&num); err != nil {
				log.Error("[dao.dao-anchor.mysql|db[sys_company_profession]] scan short id record error(%v)", err)
				return
			}
			resp.Profession[i].ProfessionChange = append(resp.Profession[i].ProfessionChange, num)
		}
	}
	return
}

//查询公司招聘数量列表
func (d *dao) RawCompanyProfessionRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error) {
	sql := fmt.Sprintf(_queryCompanyProfessionList, _companyProfessionTable, _companyTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCompanyRankResp{}
	for i := 1; i <= 10 && rows.Next(); i++ {
		var name string
		var num int64
		if err = rows.Scan(&name, &num); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company_profession|sys_company]] scan all record error(%v)", err)
			return
		}
		resp.CompanyName = append(resp.CompanyName, fmt.Sprintf("Top%d:%s", i, name))
		resp.Number = append(resp.Number, num)
	}
	return
}

//查询公司简历数量列表
func (d *dao) RawCompanyCvRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error) {
	sql := fmt.Sprintf(_queryCompanyCvList, _companyCvTable, _companyProfessionTable, _companyTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCompanyRankResp{}
	for i := 1; i <= 10 && rows.Next(); i++ {
		var name string
		var num int64
		if err = rows.Scan(&name, &num); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company_cv|sys_company_profession|sys_company]] scan all record error(%v)", err)
			return
		}
		resp.CompanyName = append(resp.CompanyName, fmt.Sprintf("Top%d:%s", i, name))
		resp.Number = append(resp.Number, num)
	}
	return
}

//查询公司简历状态数量
func (d *dao) RawCompanyCvStatusNumber(ctx context.Context) (resp *pb.WorknetCvStatusNumberResp, err error) {
	sql := fmt.Sprintf(_queryCompanyCvStatusNumber, _companyCvTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCvStatusNumberResp{}
	resp.StatusNumber = make([]int64, 4, 4)
	for rows.Next() {
		var status int32
		var num int64
		if err = rows.Scan(&status, &num); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company_cv|sys_company_profession|sys_company]] scan all record error(%v)", err)
			return
		}
		resp.StatusNumber[status] = num
	}
	return
}

//查询某时间段某公司新增公司招聘数量变化
func (d *dao) RawCompanyContestNumberChange(ctx context.Context, req *pb.WorknetContestChangeReq) (resp *pb.WorknetContestChangeResp, err error) {
	resp = &pb.WorknetContestChangeResp{}
	sqlCompany := fmt.Sprintf(_queryCompanyName, _companyTable)
	for _, v := range req.CompanyId {
		contest := &pb.WorknetContestChangeResp_WorknetContest{}
		var name string
		if err = d.db.QueryRow(ctx, sqlCompany, v).Scan(&name); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company]] scan short id record error(%v)", err)
			return
		}
		contest.CompanyName = name
		resp.Contest = append(resp.Contest, contest)
	}
	//时间区间
	first, last := getTimeDuration(xtime.Time(req.Time[0]), xtime.Time(req.Time[1]))
	//sql语句
	sqlContest := fmt.Sprintf(_queryCompanyContestNumber, _companyContestTable)
	for ; last.After(first); first = first.AddDate(0, 1, 0) {
		resp.Time = append(resp.Time, first.Format("2006-01"))
		firstOfMonth := first.Format("2006-01-02")
		lastOfMonth := first.AddDate(0, 1, -1).Format("2006-01-02")
		for i, v := range req.CompanyId {
			var num int64
			if err = d.db.QueryRow(ctx, sqlContest, v, firstOfMonth, lastOfMonth).Scan(&num); err != nil {
				log.Error("[dao.dao-anchor.mysql|db[sys_company_contest]] scan short id record error(%v)", err)
				return
			}
			resp.Contest[i].ContestChange = append(resp.Contest[i].ContestChange, num)
		}
	}
	return
}

//查询公司笔试数量列表
func (d *dao) RawCompanyContestRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error) {
	sql := fmt.Sprintf(_queryCompanyContestList, _companyContestTable, _companyTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCompanyRankResp{}
	for i := 1; i <= 15 && rows.Next(); i++ {
		var name string
		var num int64
		if err = rows.Scan(&name, &num); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company_cv|sys_company_contest|sys_company]] scan all record error(%v)", err)
			return
		}
		resp.CompanyName = append(resp.CompanyName, fmt.Sprintf("Top%d:%s", i, name))
		resp.Number = append(resp.Number, num)
	}
	return
}

//查询公司笔试（正在进行）数量列表
func (d *dao) RawCompanyContestingRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error) {
	sql := fmt.Sprintf(_queryCompanyContestingList, _companyContestTable, _companyTable)
	rows, err := d.db.Query(ctx, sql, time.Now().Format("2006-01-02"))
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCompanyRankResp{}
	for i := 1; i <= 15 && rows.Next(); i++ {
		var name string
		var num int64
		if err = rows.Scan(&name, &num); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company_contest|sys_company]] scan all record error(%v)", err)
			return
		}
		resp.CompanyName = append(resp.CompanyName, fmt.Sprintf("Top%d:%s", i, name))
		resp.Number = append(resp.Number, num)
	}
	return
}

//查询公司笔试报名数量列表
func (d *dao) RawCompanyContestApplyRank(ctx context.Context) (resp *pb.WorknetCompanyRankResp, err error) {
	sql := fmt.Sprintf(_queryCompanyContestApplyList, _companyContestApplyTable, _companyContestTable, _companyTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCompanyRankResp{}
	for i := 1; i <= 15 && rows.Next(); i++ {
		var name string
		var num int64
		if err = rows.Scan(&name, &num); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_company_contest_apply|sys_company_contest|sys_company]] scan all record error(%v)", err)
			return
		}
		resp.CompanyName = append(resp.CompanyName, fmt.Sprintf("Top%d:%s", i, name))
		resp.Number = append(resp.Number, num)
	}
	return
}




//获得时间区间
func getTimeDuration(time1, time2 xtime.Time) (first time.Time, last time.Time) {
	//开始时间
	firstTime := time1.Time()
	firstYear, firstMonth, _ := firstTime.Date()
	firstLocation := firstTime.Location()
	first = time.Date(firstYear, firstMonth, 1, 0, 0, 0, 0, firstLocation)
	//结束时间
	secondTime := time2.Time()
	secondYear, secondMonth, _ := secondTime.Date()
	secondLocation := secondTime.Location()
	secondOfMonth := time.Date(secondYear, secondMonth, 1, 0, 0, 0, 0, secondLocation)
	last = secondOfMonth.AddDate(0, 1, -1)
	return
}
