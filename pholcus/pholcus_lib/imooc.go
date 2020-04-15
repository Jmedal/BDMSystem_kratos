package pholcus_lib

import (
	"github.com/henrylee2cn/pholcus/logs"
	"strconv"
	"strings"

	"encoding/json"
	//"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/henrylee2cn/pholcus/app/downloader/request"
	. "github.com/henrylee2cn/pholcus/app/spider"
	"github.com/henrylee2cn/pholcus/common/goquery"
)

func init() {
	Imooc.Register()
}

func cleanHTML(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	return s
}

type Data struct {
	id      string `json:"id"`
	Numbers string `json:"numbers"`
}

type AjaxCourseMembers struct {
	result int    `json:"result"`
	Data   []Data `json:"data"`
	msg    string `json:"msg"`
}

var Score = make(map[string]string)

var Imooc = &Spider{
	Name:         "Imooc",
	Description:  "慕课网课程,[Auto Page][imooc.com]",
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{Url: "http://www.imooc.com/course/list", Rule: "首页请求"})
		},
		Trunk: map[string]*Rule{
			"首页请求": {
				ParseFunc: func(ctx *Context) {
					logs.Log.Informational("首页请求\n")
					tmpMp, _ := ctx.GetDom().Find("div.page").Find("a").Eq(8).Attr("href")
					maxPage, _ := strconv.Atoi(strings.TrimLeft(tmpMp, "/course/list?page="))
					ctx.Aid(map[string]interface{}{"loop": [2]int{1, maxPage}, "Rule": "所有课程"}, "所有课程")
				},
			},

			"所有课程": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					for loop := aid["loop"].([2]int); loop[0] < loop[1]; loop[0]++ {
						ctx.AddQueue(&request.Request{
							Url:  "http://www.imooc.com/course/list?page=" + strconv.Itoa(loop[0]),
							Rule: aid["Rule"].(string),
						})
					}
					return nil
				},

				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom().Find("div.course-card-container")
					query.Each(func(i int, goq *goquery.Selection) {
						//课程标题
						CourseTitle := goq.Find(".course-card-name").Text()
						//技术栈
						TechStack := goq.Find("div.course-card-top").Find(".course-label").Find("label").Text()
						//课程简介
						Introduction := goq.Find("p").Text()
						//
						Attr, ok := goq.Find("a.course-card").Attr("href")
						CourseUrlNumber := strings.Join(regexp.MustCompile("[0-9]").FindAllString(Attr, -1), "")
						url := "http://www.imooc.com/course/AjaxCourseMembers?ids=" + CourseUrlNumber
						resp, err := http.Get(url)

						if err != nil {
							log.Println("ERROR:", err)
						}
						defer resp.Body.Close()
						ajaxCourseMembers := &AjaxCourseMembers{}
						if err := json.NewDecoder(resp.Body).Decode(&ajaxCourseMembers); err != nil {
							log.Println("ERROR:", err)
						}
						//LearnerNumber:=ajaxCourseMembers.data[0].Numbers   //此处，这么写更加方便。
						LearnerNumber := ajaxCourseMembers.getnumbers()

						if ok == true {
							ctx.AddQueue(&request.Request{
								Url:  "http://www.imooc.com" + Attr,
								Rule: "Course_Info",
								Temp: map[string]interface{}{
									"CourseTitle":   CourseTitle,
									"TechStack":     TechStack,
									"Introduction":  Introduction,
									"LearnerNumber": LearnerNumber,
								},
							})
						}

					})

				},
			},
			"Course_Info": {

				ItemFields: []string{
					"name",                //课程名称
					"class",               //课程分类
					"introduction",        //课程简介
					"learner_number",      //学习人数
					"description",         //课程介绍
					"path",                //课程路径
					"difficulty",          //难度级别
					"duration",            //课程时长
					"comprehensive_score", //综合评分
					"evaluation_number",   //评价数
					"utility_score",       //内容实用
					"concise_score",       //简洁易懂
					"logic_score",         //逻辑清晰
				},
				ParseFunc: func(ctx *Context) {
					dom := ctx.GetDom()
					Summary := cleanHTML(dom.Find("div.course-description.course-wrap").Text())
					if strings.Contains(Summary, "简介：") {
						Summary = strings.Split(Summary, "简介：")[1]
					}
					CoursePath := cleanHTML(dom.Find("div.course-infos").Find("div.path").Text())
					Difficulty := dom.Find("div.course-infos").Find("div.static-item.l").Eq(0).Find(".meta-value").Text()
					Duration := dom.Find("div.course-infos").Find("div.static-item.l").Eq(1).Find(".meta-value").Text()
					scoretmp := dom.Find("div.course-infos").Find("div.static-item.l.score-btn")
					vScore0 := scoretmp.Find("span").Eq(0).Text()
					vScore1 := scoretmp.Find("span").Eq(1).Text()
					vScore2 := scoretmp.Find("span").Eq(2).Text()
					vScore3 := scoretmp.Find("span").Eq(3).Text()
					vScore4 := scoretmp.Find("span").Eq(4).Text()
					vScore5 := scoretmp.Find("span").Eq(5).Text()
					vScore6 := scoretmp.Find("span").Eq(6).Text()
					vScore7 := scoretmp.Find("span").Eq(7).Text()
					vScore8 := scoretmp.Find("span").Eq(8).Text()
					Score[vScore0] = vScore1
					Score["评价数"] = vScore2
					if strings.Contains(vScore2, "人评价") {
						vScore2 = strings.Split(vScore2, "人评价")[0]
					}
					Score[vScore4] = vScore3
					Score[vScore6] = vScore5
					Score[vScore8] = vScore7
					ctx.Output(map[int]interface{}{
						0:  ctx.GetTemp("CourseTitle", "").(string),
						1:  ctx.GetTemp("TechStack", "").(string),
						2:  ctx.GetTemp("Introduction", "").(string),
						3:  ctx.GetTemp("LearnerNumber", "").(string),
						4:  Summary,
						5:  CoursePath,
						6:  Difficulty,
						7:  Duration,
						8:  vScore1,
						9:  vScore2,
						10: vScore3,
						11: vScore5,
						12: vScore7,
					})
					Score = make(map[string]string)
				},
			},
		},
	},
}

func (s *AjaxCourseMembers) getnumbers() string {
	return s.Data[0].Numbers
}
