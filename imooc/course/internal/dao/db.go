package dao

import (
	"context"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	"strconv"
	"strings"

	pb "course/api"
	"course/internal/model"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
)

const (
	//表
	//慕课网课程数据表
	_sysCourseTable = "imooc__0__course_info"

	//查询课程评分列表
	_queryCourseScore = "select comprehensive_score, utility_score, concise_score, logic_score from %s "

	//查询课程 学习人数/评分数 列表
	_queryCourseNumber = "select learner_number, evaluation_number from %s "

	//查询课程类别列表
	_queryCourseClass = "select class, count(*) from %s group by class"

	//查询课程难度列表
	_queryCourseDifficulty = "select difficulty, count(*) from %s group by difficulty"

	//查询课程时长列表
	_queryCourseDuration = "select duration from %s"

	//查询课程综合评分列表
	_queryCourseScoreList = "select name, comprehensive_score from %s order by comprehensive_score*1 desc LIMIT ?, ?"

	//查询课程学习人数列表
	_queryCourseLearnerList = "select name, learner_number from %s order by learner_number*1 desc LIMIT ?, ?"

	//查询课程信息列表
	_queryCourseList = "select id, name, class, introduction, learner_number, description, path, difficulty, duration, comprehensive_score, utility_score, concise_score, logic_score, evaluation_number, Url, ParentUrl, DownloadTime from %s where name like ? LIMIT ?, ?"

	//查询课程总数
	_queryCourseCount = "select count(*) from %s where name like ?"
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

//查询课程评分比例
func (d *dao) RawCourseScoreSection(ctx context.Context) (resp *pb.ImoocCourseScoreSectionResp, err error) {
	sql := fmt.Sprintf(_queryCourseScore, _sysCourseTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.ImoocCourseScoreSectionResp{}
	comprehensiveMap := make(map[string]int64)
	utilityMap := make(map[string]int64)
	conciseMap := make(map[string]int64)
	logicMap := make(map[string]int64)
	for rows.Next() {
		var comprehensive string
		var utility string
		var concise string
		var logic string
		if err = rows.Scan(&comprehensive, &utility, &concise, &logic); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|imooc__0__course_info]] scan all record error(%v)", err)
			return
		}
		score, err1 := strconv.ParseFloat(comprehensive, 10)
		if err1 == nil {
			comprehensiveMap = scoreSectionAdd(comprehensiveMap, score)
		}
		score, err1 = strconv.ParseFloat(utility, 10)
		if err1 == nil {
			utilityMap = scoreSectionAdd(utilityMap, score)
		}
		score, err1 = strconv.ParseFloat(concise, 10)
		if err1 == nil {
			conciseMap = scoreSectionAdd(conciseMap, score)
		}
		score, err1 = strconv.ParseFloat(logic, 10)
		if err1 == nil {
			logicMap = scoreSectionAdd(logicMap, score)
		}
	}
	resp.Comprehensive = scoreSectionOutput(comprehensiveMap)
	resp.Utility = scoreSectionOutput(utilityMap)
	resp.Concise = scoreSectionOutput(conciseMap)
	resp.Logic = scoreSectionOutput(logicMap)
	return
}

//查询课程学习人数/评论数
func (d *dao) RawCourseNumber(ctx context.Context) (resp *pb.ImoocCourseAllNumberResp, err error) {
	sql := fmt.Sprintf(_queryCourseNumber, _sysCourseTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.ImoocCourseAllNumberResp{}
	for rows.Next() {
		var learner string
		var evaluation string
		if err = rows.Scan(&learner, &evaluation); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|imooc__0__course_info]] scan all record error(%v)", err)
			return
		}
		learnerNumber, _ := strconv.ParseInt(learner, 10, 64)
		resp.AllLearnerNumber += learnerNumber
		evaluationNumber, _ := strconv.ParseInt(evaluation, 10, 64)
		resp.AllEvaluationNumber += evaluationNumber
	}
	return
}

//查询课程比例
func (d *dao) RawCourseSection(ctx context.Context) (resp *pb.ImoocCourseSectionResp, err error) {
	resp = &pb.ImoocCourseSectionResp{}
	//课程类别
	sql := fmt.Sprintf(_queryCourseClass, _sysCourseTable)
	rows, err := d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	class := &pb.ImoocCourseSectionResp_Lists{}
	resp.Class = class
	for rows.Next() {
		var name string
		var number int64
		if err = rows.Scan(&name, &number); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|imooc__0__course_info]] scan all record error(%v)", err)
			return
		}
		if name == "" {
			name = "其它"
		}
		class.Name = append(class.Name, name)
		class.Number = append(class.Number, number)
	}
	//课程难度
	sql = fmt.Sprintf(_queryCourseDifficulty, _sysCourseTable)
	rows, err = d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	difficulty := &pb.ImoocCourseSectionResp_Lists{}
	resp.Difficulty = difficulty
	for rows.Next() {
		var name string
		var number int64
		if err = rows.Scan(&name, &number); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|imooc__0__course_info]] scan all record error(%v)", err)
			return
		}
		if name == "" {
			name = "其它"
		}
		difficulty.Name = append(difficulty.Name, name)
		difficulty.Number = append(difficulty.Number, number)
	}
	//课程时长
	sql = fmt.Sprintf(_queryCourseDuration, _sysCourseTable)
	rows, err = d.db.Query(ctx, sql)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	duration := &pb.ImoocCourseSectionResp_Lists{}
	resp.Duration = duration
	durationMap := make(map[string]int64)
	for rows.Next() {
		var duration string
		if err = rows.Scan(&duration); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|imooc__0__course_info]] scan all record error(%v)", err)
			return
		}
		time := strings.Split(duration, "小时")
		if len(time) == 1 {
			durationMap["小于1小时"]++
		} else {
			time = strings.Split(time[0], " ")
			if len(time) == 2 {
				time[0] = time[1]
			}
			hour, err1 := strconv.ParseInt(time[0], 10, 64)
			if err1 != nil {
				log.Error("string to int64 err(%v)", err1)
				continue
			}
			durationMap[fmt.Sprintf("%d小时~%d小时", hour, hour+1)]++
		}
	}
	for name, number := range durationMap {
		duration.Name = append(duration.Name, name)
		duration.Number = append(duration.Number, number)
	}
	return
}

//查询课程综合评分列表
func (d *dao) RawCourseScoreRank(ctx context.Context, req *pb.ImoocCourseRankReq) (resp *pb.ImoocCourseRankResp, err error) {
	first := req.Location[0]
	last := req.Location[1]
	sql := fmt.Sprintf(_queryCourseScoreList, _sysCourseTable)
	rows, err := d.db.Query(ctx, sql, first, last)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.ImoocCourseRankResp{}
	for i := first; rows.Next(); i++ {
		var name string
		var score string
		if err = rows.Scan(&name, &score); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|imooc__0__course_info]] scan all record error(%v)", err)
			return
		}
		resp.Name = append(resp.Name, fmt.Sprintf("Top%d:%s", i+1, name))
		resp.Number = append(resp.Number, score)
	}
	return
}

//查询课程学习人数列表
func (d *dao) RawCourseLearnerRank(ctx context.Context, req *pb.ImoocCourseRankReq) (resp *pb.ImoocCourseRankResp, err error) {
	first := req.Location[0]
	last := req.Location[1]
	sql := fmt.Sprintf(_queryCourseLearnerList, _sysCourseTable)
	rows, err := d.db.Query(ctx, sql, first, last)
	if err != nil {
		log.Error("select course err(%v)", err)
		return
	}
	defer rows.Close()
	resp = &pb.ImoocCourseRankResp{}
	for i := first; rows.Next(); i++ {
		var name string
		var learner string
		if err = rows.Scan(&name, &learner); err != nil {
			log.Error("[dao.dao-anchor.mysql|db[sys_course|imooc__0__course_info]] scan all record error(%v)", err)
			return
		}
		resp.Name = append(resp.Name, fmt.Sprintf("Top%d:%s", i+1, name))
		resp.Number = append(resp.Number, learner)
	}
	return
}

//查询课程信息列表
func (d *dao) RawCoursePage(ctx context.Context, req *pb.GetCoursePageReq) (resp *pb.GetCoursePageResp, err error) {
	sqlCourseList := fmt.Sprintf(_queryCourseList, _sysCourseTable)
	sqlCourseCount := fmt.Sprintf(_queryCourseCount, _sysCourseTable)
	query := fmt.Sprintf("%%%s%%", req.Query)
	resp = &pb.GetCoursePageResp{}
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	if err = d.db.QueryRow(ctx, sqlCourseCount, query).Scan(&resp.Total); err != nil {
		log.Error("[dao.dao-anchor.mysql|db[imooc__0__course_info]] scan all record error(%v)", err)
		return
	}
	start := (req.PageNum - 1) * req.PageSize
	size := req.PageSize
	rows, err := d.db.Query(ctx, sqlCourseList, query, start, size)
	if err != nil {
		log.Error("select course error(%v)", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		course := &pb.Course{}
		if err = rows.Scan(&course.Id, &course.Name, &course.Class, &course.Introduction,
			&course.LearnerNumber, &course.Description, &course.Path, &course.Difficulty,
			&course.Duration, &course.ComprehensiveScore, &course.UtilityScore, &course.ConciseScore,
			&course.LogicScore, &course.EvaluationNumber, &course.Url, &course.ParentUrl,
			&course.DownloadTime); err != nil {
			log.Error("[dao.dao-anchor.mysql|imooc__0__course_info] scan short id record error(%v)", err)
			return
		}
		resp.CourseList = append(resp.CourseList, course)
	}
	return
}

func scoreSectionAdd(sign map[string]int64, score float64) map[string]int64 {
	switch {
	case 9 < score && score <= 10:
		sign["10分~9分"]++
	case 8 < score && score <= 9:
		sign["9分~8分"]++
	case 7 < score && score <= 8:
		sign["8分~7分"]++
	case 6 < score && score <= 7:
		sign["7分~6分"]++
	case 3 < score && score <= 6:
		sign["6分~3分"]++
	case 0 < score && score <= 3:
		sign["3分~0分"]++
	default:
		sign["无评分"]++
	}
	return sign
}

func scoreSectionOutput(sign map[string]int64) (number []int64) {
	number = append(number, sign["10分~9分"])
	number = append(number, sign["9分~8分"])
	number = append(number, sign["8分~7分"])
	number = append(number, sign["7分~6分"])
	number = append(number, sign["6分~3分"])
	number = append(number, sign["3分~0分"])
	number = append(number, sign["无评分"])
	return
}
