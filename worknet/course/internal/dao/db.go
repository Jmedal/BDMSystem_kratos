package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	"time"

	pb "course/api"
	"course/internal/model"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	xtime "github.com/bilibili/kratos/pkg/time"
)

const (
	//标识
	//激活
	_activityLabel = "1"
	//未激活
	_inactivityLabel = "0"
	//完成
	_finishLabel = "1"
	//未完成
	_unFinishLabel = "0"

	//表名
	//课程表
	_courseTable = "sys_course"
	//课程学员学习表
	_courseStudiedTable = "sys_course_studied"
	//课程评价表
	_courseCommentTable = "sys_course_comment"
	//课程点击表
	_courseClickTable = "sys_course_click"
	//课程目录表
	_courseIndexTable = "sys_course_index"
	//课程视频表
	_courseVideoTable = "sys_course_video"

	//SQL语句
	//查询是否激活课程数量
	_queryActivityCourseNumber = "select count(*) AS number from %s where activity = ?"

	//查询某年某月中新增课程数量										'2019-6-1' '2020-1-1'
	_queryCourseNumber = "select count(*) from %s where upload_time between ? and ? "

	//查询课程列表
	_queryCourseList = "select id, name from %s"

	//查询某个课程名称
	_queryCourseName = "select name  from %s where id = ?"

	//查询某个课程某个时间段新增学员数量												'2019-6-1' '2020-1-1'
	_queryStudentNumber = "select count(*) from %s where course_id = ? and (start_time between ? and ? )"

	//查询某课程某时间段开始学习的学员完成课程学习数量										                             '2019-6-1' '2020-1-1'
	_queryStudentFinishNumber = "select count(*) from %s where course_id = ? and is_finished = ? and (start_time between ? and ? )"

	//查询所有课程名和平均分
	_queryCourseAverageScore1 = "select sys_course.name, avg(score) AS avgScore from %s inner join %s on sys_course.id = sys_course_studied.course_id where sys_course_studied.is_finished = ? group by sys_course.id order by avgScore desc"

	//查询所有课程平均分
	_queryCourseAverageScore2 = "select avg(score) AS avgScore from %s inner join %s on sys_course.id = sys_course_studied.course_id where sys_course_studied.is_finished = ? group by sys_course.id order by avgScore desc"

	//查询所有课程综合评分
	_queryCourseStars = "select sys_course.name, avg(sys_course_comment.stars) AS stars from %s inner join %s on sys_course.id = sys_course_comment.course_id group by sys_course.id order by stars desc"

	//查询所有课程点击量
	_queryCourseClick = "select sys_course.name, count(*) AS click from  %s inner join  %s on sys_course.id = sys_course_click.course_id group by sys_course.id order by click desc"

	//查询所有课程视频数量
	_queryCourseVideo = "select sys_course.name, count(*) AS video from (%s inner join %s on sys_course.id = sys_course_index.course_id) inner join %s on sys_course_index.id = sys_course_video.course_Index_id group by sys_course.id order by video desc"
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

//查询激活和未激活课程数量
func (d *dao) RawActivityCourseNumber(ctx context.Context) (resp *pb.WorknetCourseActivityResp, err error) {
	sql := fmt.Sprintf(_queryActivityCourseNumber, _courseTable)
	num := make([]int64, 2, 2)
	for i := 0; i < 2; i++ {
		if err = d.db.QueryRow(ctx, sql, i).Scan(&num[i]); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course]] scan record error(%v)", err)
			return
		}
	}
	resp = &pb.WorknetCourseActivityResp{
		Activity:   num[1],
		Inactivity: num[0],
	}
	return
}

//查询某时间段新增课程数量变化
func (d *dao) RawCourseNumberChange(ctx context.Context, req *pb.WorknetCourseChangeReq) (resp *pb.WorknetCourseChangeResp, err error) {
	//时间区间
	first, last := getTimeDuration(xtime.Time(req.Time[0]), xtime.Time(req.Time[1]))
	//sql语句
	sql := fmt.Sprintf(_queryCourseNumber, _courseTable)
	resp = &pb.WorknetCourseChangeResp{}
	for ; last.After(first); first = first.AddDate(0, 1, 0) {
		resp.Time = append(resp.Time, first.Format("2006-01"))
		firstOfMonth := first.Format("2006-01-02")
		lastOfMonth := first.AddDate(0, 1, -1).Format("2006-01-02")
		var num int64
		if err = d.db.QueryRow(ctx, sql, firstOfMonth, lastOfMonth).Scan(&num); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course]] scan short id record error(%v)", err)
			return
		}
		resp.Course = append(resp.Course, num)
	}
	return
}

//查询课程列表
func (d *dao) RawCourseList(ctx context.Context) (resp *pb.WorknetCourseListResp, err error) {
	sql := fmt.Sprintf(_queryCourseList, _courseTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCourseListResp{}
	for rows.Next() {
		course := &pb.WorknetCourseListResp_WorknetCourse{}
		if err = rows.Scan(&course.CourseId, &course.CourseName); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course]] scan all record error(%v)", err)
			return
		}
		resp.CourseList = append(resp.CourseList, course)
	}
	return
}

//查询某时间段某课程新增学员数量变化
func (d *dao) RawStudentNumberChange(ctx context.Context, req *pb.WorknetStudentChangeReq) (resp *pb.WorknetStudentChangeResp, err error) {
	resp = &pb.WorknetStudentChangeResp{}
	sqlCourse := fmt.Sprintf(_queryCourseName, _courseTable)
	for _, v := range req.CourseId {
		course := &pb.WorknetStudentChangeResp_WorknetCourse{}
		var name string
		if err = d.db.QueryRow(ctx, sqlCourse, v).Scan(&name); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course]] scan short id record error(%v)", err)
			return
		}
		course.CourseName = name
		resp.Student = append(resp.Student, course)
	}
	//时间区间
	first, last := getTimeDuration(xtime.Time(req.Time[0]), xtime.Time(req.Time[1]))
	//sql语句
	sqlStudent := fmt.Sprintf(_queryStudentNumber, _courseStudiedTable)
	for ; last.After(first); first = first.AddDate(0, 1, 0) {
		resp.Time = append(resp.Time, first.Format("2006-01"))
		firstOfMonth := first.Format("2006-01-02")
		lastOfMonth := first.AddDate(0, 1, -1).Format("2006-01-02")
		for i, v := range req.CourseId {
			var num int64
			if err = d.db.QueryRow(ctx, sqlStudent, v, firstOfMonth, lastOfMonth).Scan(&num); err != nil {
				log.Error("[dao.dao-anchor.mysql|db[sys_course_studied]] scan short id record error(%v)", err)
				return
			}
			resp.Student[i].StudentChange = append(resp.Student[i].StudentChange, num)
		}
	}
	return
}

//查询某课程某时间段开始学习的学员完成课程数量
func (d *dao) RawStudentFinishNumber(ctx context.Context, req *pb.WorknetStudentFinishReq) (resp *pb.WorknetStudentFinishResp, err error) {
	resp = &pb.WorknetStudentFinishResp{}
	sqlCourse := fmt.Sprintf(_queryCourseName, _courseTable)
	for _, v := range req.CourseId {
		var name string
		if err = d.db.QueryRow(ctx, sqlCourse, v).Scan(&name); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course]] scan short id record error(%v)", err)
			return
		}
		resp.CourseName = append(resp.CourseName, name)
	}
	//时间区间
	first, last := getTimeDuration(xtime.Time(req.Time[0]), xtime.Time(req.Time[1]))
	//sql语句
	sqlStudent := fmt.Sprintf(_queryStudentFinishNumber, _courseStudiedTable)
	start := first.Format("2006-01-02")
	end := last.AddDate(0, 1, -1).Format("2006-01-02")
	for _, v := range req.CourseId {
		var finish int64
		if err = d.db.QueryRow(ctx, sqlStudent, v, _finishLabel, start, end).Scan(&finish); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course_studied]] scan short id record error(%v)", err)
			return
		}
		var unFinish int64
		if err = d.db.QueryRow(ctx, sqlStudent, v, _unFinishLabel, start, end).Scan(&unFinish); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course_studied]] scan short id record error(%v)", err)
			return
		}
		resp.Finish = append(resp.Finish, finish)
		resp.UnFinish = append(resp.UnFinish, unFinish)
	}
	return
}

//查询所有课程名和平均分
func (d *dao) RawCourseAverageScoreList(ctx context.Context) (resp *pb.WorknetAverageScoreListResp, err error) {
	sql := fmt.Sprintf(_queryCourseAverageScore1, _courseTable, _courseStudiedTable)
	rows, err := d.db.Query(ctx, sql, _finishLabel)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetAverageScoreListResp{}
	for i := 1; i <= 10 && rows.Next(); i++ {
		var name string
		var score float32
		if err = rows.Scan(&name, &score); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|sys_course_studied]] scan all record error(%v)", err)
			return
		}
		resp.CourseName = append(resp.CourseName, fmt.Sprintf("Top%d:%s", i, name))
		resp.AverageScore = append(resp.AverageScore, fmt.Sprintf("%.2f", score))
	}
	return
}

//查询所有课程平均分
func (d *dao) RawCourseAverageScoreSection(ctx context.Context) (resp *pb.WorknetAverageScoreSectionResp, err error) {
	sql := fmt.Sprintf(_queryCourseAverageScore2, _courseTable, _courseStudiedTable)
	rows, err := d.db.Query(ctx, sql, _finishLabel)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetAverageScoreSectionResp{}
	resp.CourseNumber = make([]int64, 5, 5)
	for rows.Next() {
		var score float32
		if err = rows.Scan(&score); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|sys_course_studied]] scan all record error(%v)", err)
			return
		}
		switch {
		case score >= 90 && score <= 100:
			resp.CourseNumber[0]++
		case score >= 80 && score < 90:
			resp.CourseNumber[1]++
		case score >= 70 && score < 80:
			resp.CourseNumber[2]++
		case score >= 60 && score < 70:
			resp.CourseNumber[3]++
		case score < 60:
			resp.CourseNumber[4]++
		}
	}
	return
}

//查询课程综合评分列表
func (d *dao) RawCourseStarsRank(ctx context.Context, req *pb.WorknetCourseRankReq) (resp *pb.WorknetCourseRankResp, err error) {
	first := req.Location[0]
	last := req.Location[1]
	sql := fmt.Sprintf(_queryCourseStars, _courseTable, _courseCommentTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCourseRankResp{}
	for i := int64(1); i <= last && rows.Next(); i++ {
		if i >= first {
			var name string
			var stars float32
			if err = rows.Scan(&name, &stars); err != nil {
				log.Error("[dao.dao-anchor.mysql|db[sys_course|sys_course_comment]] scan all record error(%v)", err)
				return
			}
			resp.CourseName = append(resp.CourseName, fmt.Sprintf("Top%d:%s", i, name))
			resp.Number = append(resp.Number, fmt.Sprintf("%.2f", stars))
		}
	}
	return
}

//查询课程点击量列表
func (d *dao) RawCourseClickRank(ctx context.Context, req *pb.WorknetCourseRankReq) (resp *pb.WorknetCourseRankResp, err error) {
	first := req.Location[0]
	last := req.Location[1]
	sql := fmt.Sprintf(_queryCourseClick, _courseTable, _courseClickTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCourseRankResp{}
	for i := int64(1); i <= last && rows.Next(); i++ {
		if i >= first {
			var name string
			var click float32
			if err = rows.Scan(&name, &click); err != nil {
				log.Error("[dao.dao-anchor.mysql|db[sys_course|sys_course_click]] scan all record error(%v)", err)
				return
			}
			resp.CourseName = append(resp.CourseName, fmt.Sprintf("Top%d:%s", i, name))
			resp.Number = append(resp.Number, fmt.Sprintf("%.2f", click))
		}
	}
	return
}

//查询课程视频数量列表
func (d *dao) RawCourseVideoRank(ctx context.Context, req *pb.WorknetCourseRankReq) (resp *pb.WorknetCourseRankResp, err error) {
	first := req.Location[0]
	last := req.Location[1]
	sql := fmt.Sprintf(_queryCourseVideo, _courseTable, _courseIndexTable, _courseVideoTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.WorknetCourseRankResp{}
	for i := int64(1); i <= last && rows.Next(); i++ {
		if i >= first {
			var name string
			var video float32
			if err = rows.Scan(&name, &video); err != nil {
				log.Error("[dao.dao-anchor.mysql|db[sys_course|sys_course_index|sys_course_video]] scan all record error(%v)", err)
				return
			}
			resp.CourseName = append(resp.CourseName, fmt.Sprintf("Top%d:%s", i, name))
			resp.Number = append(resp.Number, fmt.Sprintf("%.2f", video))
		}
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
