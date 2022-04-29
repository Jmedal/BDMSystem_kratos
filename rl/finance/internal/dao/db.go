package dao

import (
	"context"
	dsql "database/sql"
	pb "finance/api"
	"finance/internal/model"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	xtime "github.com/bilibili/kratos/pkg/time"
	"time"
)

const (
	//激活
	_activity = 1

	//冻结
	_frozen = 2

	//表
	_rlFinanceProjectTable = "rl_finance_project"

	_rlFinanceBillTable = "rl_finance_bill"

	//查询项目列表
	_queryProjectList = "SELECT rl_finance_project.id, rl_finance_project.`name`, rl_finance_project.development, rl_finance_project.construction, rl_finance_project.`schedule`, rl_finance_project.contract, rl_finance_project.signTime, rl_finance_project.prepayment, rl_finance_project.progress, rl_finance_project.judge, rl_finance_project.remark, ( SELECT IFNULL( SUM(rl_finance_bill.amount), 0) FROM %s WHERE rl_finance_project.id = rl_finance_bill.project_id AND rl_finance_bill.state = '1' AND rl_finance_bill.type = '1') AS transfer, (SELECT COUNT(*) FROM %s WHERE rl_finance_project.id = rl_finance_bill.project_id AND rl_finance_bill.state = '1' AND rl_finance_bill.type = '1') AS transfer_times, ( SELECT IFNULL( SUM(rl_finance_bill.amount), 0) FROM %s WHERE rl_finance_project.id = rl_finance_bill.project_id AND rl_finance_bill.state = '1' AND rl_finance_bill.type = '2') AS expense, (SELECT COUNT(*) FROM %s WHERE rl_finance_project.id = rl_finance_bill.project_id AND rl_finance_bill.state = '1' AND rl_finance_bill.type = '2') AS expense_times, rl_finance_project.state, rl_finance_project.create_user_id, rl_finance_project.create_time FROM %s WHERE `name` LIKE ? OR development LIKE ? OR construction LIKE ? ORDER BY rl_finance_project.state, create_time DESC LIMIT ?, ?"

	//查询项目数目
	_queryProjectCount = "SELECT COUNT(*) FROM %s WHERE `name` LIKE ? OR development LIKE ? OR construction LIKE ?"

	//插入项目
	_insertProject = "INSERT INTO %s (`name`, development, construction, `schedule`, contract, signTime, prepayment, progress, judge, remark, state, create_user_id, create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

	//更新项目
	_updateProject = "UPDATE %s SET `name` = ?, development = ?, construction = ?, `schedule` = ?, contract = ?, signTime = ?, prepayment = ?, progress = ?, judge = ?, remark = ? WHERE id = ?"

	//完成/重启项目
	_updateStateProject = "UPDATE %s SET state = ? WHERE id = ?"

	//查询账目列表
	_queryBillList = "SELECT id, time, amount, direction, type FROM %s WHERE direction LIKE ? AND project_id = ? AND state = '1' ORDER BY time DESC LIMIT ?, ?"

	//查询账目数目
	_queryBillCount = "SELECT count(*) FROM %s WHERE direction LIKE ? AND project_id = ? AND state = '1'"

	//插入账目
	_insertBill = "INSERT INTO %s (project_id, time, amount, direction, type, state) VALUES (?,?,?,?,?,?)"

	//更新账目
	_updateBill = "UPDATE %s SET time = ?, amount = ?, direction = ?, type = ? WHERE id = ?"

	//删除账目
	_deleteStateBill = "UPDATE %s SET state = '2' WHERE id = ?"
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

//查询项目分页
func (d *dao) RawProjectPage(ctx context.Context, req *pb.GetProjectPageReq) (resp *pb.GetProjectPageResp, err error) {

	sqlProjectList := fmt.Sprintf(_queryProjectList, _rlFinanceBillTable, _rlFinanceBillTable, _rlFinanceBillTable, _rlFinanceBillTable, _rlFinanceProjectTable)
	sqlProjectCount := fmt.Sprintf(_queryProjectCount, _rlFinanceProjectTable)
	query := fmt.Sprintf("%%%s%%", req.Query)
	resp = &pb.GetProjectPageResp{}
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	if err = d.db.QueryRow(ctx, sqlProjectCount, query, query, query).Scan(&resp.Total); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] scan all record error(%v)", err)
		return
	}
	start := (req.PageNum - 1) * req.PageSize
	size := req.PageSize
	rows, err := d.db.Query(ctx, sqlProjectList, query, query, query, start, size)
	if err != nil {
		log.Error("select finance error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var signTime xtime.Time
		var createTime xtime.Time
		var userId int64
		projectInfo := &pb.ProjectInfo{}
		if err = rows.Scan(&projectInfo.Id, &projectInfo.Name, &projectInfo.Development,
			&projectInfo.Construction, &projectInfo.Schedule, &projectInfo.Contract,
			&signTime, &projectInfo.Prepayment, &projectInfo.Progress, &projectInfo.Judge,
			&projectInfo.Remark, &projectInfo.Transfer, &projectInfo.TransferTimes,
			&projectInfo.Expense, &projectInfo.ExpenseTimes, &projectInfo.State, &userId, &createTime); err != nil {
			log.Error("[dao.dao-anchor.mysql|rl_finance] scan short id record error(%v)", err)
			return
		}
		userNameReq := &pb.GetUserNameReq{}
		userNameReq.Id = userId
		var userNameResp *pb.GetUserNameResp
		if userNameResp, err = d.accountClient.GetUserName(ctx, userNameReq); err != nil {
			log.Error("d.accountClient.GetUserName err(%v)", err)
			return
		}
		projectInfo.CreateUser = userNameResp.Name
		projectInfo.SignTime = signTime.Time().Unix()
		projectInfo.CreateTime = createTime.Time().Unix()
		resp.ProjectList = append(resp.ProjectList, projectInfo)
	}
	return
}

//插入项目
func (d *dao) InsertProject(ctx context.Context, req *pb.AddProjectReq) (resp *pb.AddProjectResp, err error) {
	sqlProject := fmt.Sprintf(_insertProject, _rlFinanceProjectTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlProject, req.Project.Name, req.Project.Development, req.Project.Construction, req.Project.Schedule,
		req.Project.Contract, xtime.Time(req.Project.SignTime).Time(), req.Project.Prepayment, req.Project.Progress,req.Project.Judge,
		req.Project.Remark, req.Project.State, req.UserId, time.Now().Format("2006-01-02 15:04:05")); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] failed to insert error(%v)", err)
		return
	}

	resp = &pb.AddProjectResp{}
	if _, err = res.LastInsertId(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//更新项目
func (d *dao) UpdateProject(ctx context.Context, req *pb.UpdateProjectReq) (resp *pb.UpdateProjectResp, err error) {
	sqlProject := fmt.Sprintf(_updateProject, _rlFinanceProjectTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlProject, req.Project.Name, req.Project.Development,
		req.Project.Construction, req.Project.Schedule, req.Project.Contract,
		xtime.Time(req.Project.SignTime).Time(), req.Project.Prepayment,req.Project.Progress,
		req.Project.Judge,req.Project.Remark, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.UpdateProjectResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//完成项目
func (d *dao) FrozenStateProject(ctx context.Context, req *pb.TerminationProjectReq) (resp *pb.TerminationProjectResp, err error) {
	sqlProject := fmt.Sprintf(_updateStateProject, _rlFinanceProjectTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlProject, _frozen, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.TerminationProjectResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//重启项目
func (d *dao) ActivityStateProject(ctx context.Context, req *pb.RestartProjectReq) (resp *pb.RestartProjectResp, err error) {
	sqlProject := fmt.Sprintf(_updateStateProject, _rlFinanceProjectTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlProject, _activity, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.RestartProjectResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//查询账目分页
func (d *dao) RawBillPage(ctx context.Context, req *pb.GetBillPageReq) (resp *pb.GetBillPageResp, err error) {
	sqlBillList := fmt.Sprintf(_queryBillList, _rlFinanceBillTable)
	sqlBillCount := fmt.Sprintf(_queryBillCount, _rlFinanceBillTable)
	query := fmt.Sprintf("%%%s%%", req.Query)
	resp = &pb.GetBillPageResp{}
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	if err = d.db.QueryRow(ctx, sqlBillCount, query, req.ProjectId).Scan(&resp.Total); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] scan all record error(%v)", err)
		return
	}
	start := (req.PageNum - 1) * req.PageSize
	size := req.PageSize
	rows, err := d.db.Query(ctx, sqlBillList, query, req.ProjectId, start, size)
	if err != nil {
		log.Error("select finance error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var billTime xtime.Time
		billInfo := &pb.BillInfo{}
		if err = rows.Scan(&billInfo.Id, &billTime, &billInfo.Amount, &billInfo.Direction, &billInfo.Type); err != nil {
			log.Error("[dao.dao-anchor.mysql|rl_finance] scan short id record error(%v)", err)
			return
		}
		billInfo.Time = billTime.Time().Unix()
		resp.BillList = append(resp.BillList, billInfo)
	}
	return
}

//插入账目
func (d *dao) InsertBill(ctx context.Context, req *pb.AddBillReq) (resp *pb.AddBillResp, err error) {
	sqlBill := fmt.Sprintf(_insertBill, _rlFinanceBillTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlBill, req.Bill.ProjectId, xtime.Time(req.Bill.Time).Time(),
		req.Bill.Amount, req.Bill.Direction, req.Bill.Type, req.Bill.State); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] failed to insert error(%v)", err)
		return
	}
	resp = &pb.AddBillResp{}
	if _, err = res.LastInsertId(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//更新账目
func (d *dao) UpdateBill(ctx context.Context, req *pb.UpdateBillReq) (resp *pb.UpdateBillResp, err error) {
	sqlBill := fmt.Sprintf(_updateBill, _rlFinanceBillTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlBill, xtime.Time(req.Bill.Time).Time(),
		req.Bill.Amount, req.Bill.Direction, req.Bill.Type, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.UpdateBillResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}

//删除账目
func (d *dao) FrozenStateBill(ctx context.Context, req *pb.DeleteBillReq) (resp *pb.DeleteBillResp, err error) {
	sqlBill := fmt.Sprintf(_deleteStateBill, _rlFinanceBillTable)
	var res dsql.Result
	if res, err = d.db.Exec(ctx, sqlBill, req.Id); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[rl_finance]] failed to update: (%v), error(%v)", req.Id, err)
	}
	resp = &pb.DeleteBillResp{}
	if _, err = res.RowsAffected(); err == nil {
		resp.Result = "success"
		return
	}
	resp.Result = "error"
	return
}