package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	"strconv"

	pb "curriculum/api"
	"curriculum/internal/model"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
)

const (
	//表名
	//学科表
	_curriculumTable = "sys_curriculum"

	//学科课程表
	_curriculumCourseTable = "sys_curriculum_course"

	//学科依赖树表
	_curriculumTreeTable = "sys_curriculum_tree"

	//SQL语句
	//查询某学科树
	_queryCurriculumList = "select id, name from %s"

	//查询某学科树
	_queryCurriculumTree = "select id, pre_curriculum_id, curriculum_id from %s"

	//查询学科课程数量列表
	_queryCurriculumCourseNumber = "select name, count(*) AS number from %s inner join %s on sys_curriculum.id = sys_curriculum_course.curriculum_id group by sys_curriculum.id"
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

//查询学科课程数量列表
func (d *dao) RawCurriculumTree(ctx context.Context) (resp *pb.WorknetCurriculumGraphResp, err error) {
	sql := fmt.Sprintf(_queryCurriculumList, _curriculumTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select curriculum err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCurriculumGraphResp{}
	for rows.Next() {
		curriculum := &pb.WorknetCurriculumGraphResp_Curriculum{}
		if err = rows.Scan(&curriculum.Id, &curriculum.Name); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_curriculum]] scan all record error(%v)", err)
			return
		}
		resp.Nodes = append(resp.Nodes, curriculum)
	}
	sql = fmt.Sprintf(_queryCurriculumTree, _curriculumTreeTable)
	rows, err = d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select curriculum err(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		link := &pb.WorknetCurriculumGraphResp_Link{}
		if err = rows.Scan(&link.Id, &link.Source, &link.Target); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_curriculum_tree]] scan all record error(%v)", err)
			return
		}
		resp.Links = append(resp.Links, link)
	}
	return
}

//查询学科课程数量列表
func (d *dao) RawCurriculumCourse(ctx context.Context) (resp *pb.WorknetCurriculumCourseResp, err error) {
	sql := fmt.Sprintf(_queryCurriculumCourseNumber, _curriculumTable, _curriculumCourseTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select curriculum err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCurriculumCourseResp{}
	for rows.Next() {
		var name string
		var number int64
		if err = rows.Scan(&name, &number); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_curriculum|sys_curriculum_course]] scan all record error(%v)", err)
			return
		}
		resp.CurriculumName = append(resp.CurriculumName, name)
		resp.CourseNumber = append(resp.CourseNumber, strconv.FormatInt(number, 10))
	}
	return
}
