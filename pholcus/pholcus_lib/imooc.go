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

var chapterin = make(map[string]string)
var Score = make(map[string]string)
var chapterout = make(map[string]interface{})
var chapterall = make(map[string]interface{})

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
					query := ctx.GetDom().Find(".course-card-container")
					query.Each(func(i int, goq *goquery.Selection) {
						CourseTitle := goq.Find(".course-card-name").Text()
						TechStack := goq.Find(".course-card-top span").Text()
						Introduction := goq.Find("p").Text()
						Attr, ok := goq.Find(".course-card").Attr("href")
						CourseUrlNumber := strings.Join(regexp.MustCompile("[0-9]").FindAllString(Attr, -1), "")
						url := "http://www.imooc.com/course/AjaxCourseMembers?ids=" + CourseUrlNumber
						resp, err := http.Get(url)

						/***

						  if err != nil {
						      log.Println("ERROR:", err)
						      return
						  }
						  doc1, _ := ioutil.ReadAll(resp.Body)
						  ajaxCourseMembers := &AjaxCourseMembers{}
						  if err:= json.Unmarshal([]byte(string(doc1)), &ajaxCourseMembers);err!=nil{
						      log.Println("ERROR:", err)
						      return
						  }

						  ***/

						//myjson, _ := ioutil.ReadAll(resp.Body)
						//fmt.Println(string(myjson))  //resp的body内容OK

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
								Rule: "课程详细信息",
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
			"课程详细信息": {

				ItemFields: []string{
					"课程名称",
					"课程分类",
					"课程简介",
					"学习人数",
					"课程介绍",
					"课程路径",
					"难度级别",
					"课程时长",
					"评分",
					"章节",
				},
				ParseFunc: func(ctx *Context) {
					dom := ctx.GetDom()
					query := dom.Find(".mod-chapters > div")
					Summary := cleanHTML(dom.Find("div.course-brief").Text())
					CoursePath := cleanHTML(dom.Find(".course-infos").Find(".path").Text())
					Difficulty := dom.Find(".course-infos").Find("div.static-item").Eq(1).Find(".meta-value").Text()
					Duration := dom.Find(".course-infos").Find("div.static-item").Eq(2).Find(".meta-value").Text()
					scoretmp := dom.Find(".course-infos").Find(".score-btn")
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
					Score[vScore4] = vScore3
					Score[vScore6] = vScore5
					Score[vScore8] = vScore7
					query.Each(func(i int, goq *goquery.Selection) {
						ChapterH1 := cleanHTML(goq.Find("strong").After("i").Text())
						ctx.SetTemp("ChapterH1", ChapterH1)
						Chapter2_html := goq.Find("ul.video>li")
						Chapter2_html.Each(func(_ int, goq1 *goquery.Selection) {
							Chapter2_url, _ := goq1.Find("a").Attr("href")
							Chapter2 := cleanHTML(goq1.Find("a").After("button").Text())
							chapterin[Chapter2] = "www.imooc.com" + cleanHTML(Chapter2_url)
							ctx.SetTemp("JsonChapterH1", chapterin)
						})
						chapterout[ctx.GetTemp("ChapterH1", "").(string)] = ctx.GetTemp("JsonChapterH1", "")
						chapterall[ctx.GetTemp("CourseTitle", "").(string)] = chapterout
						chapterin = make(map[string]string)

					})
					chapterout = make(map[string]interface{})
					ctx.Output(map[int]interface{}{
						0: ctx.GetTemp("CourseTitle", "").(string),
						1: ctx.GetTemp("TechStack", "").(string),
						2: ctx.GetTemp("Introduction", "").(string),
						3: ctx.GetTemp("LearnerNumber", "").(string),
						4: Summary,
						5: CoursePath,
						6: Difficulty,
						7: Duration,
						8: Score,
						9: chapterall[ctx.GetTemp("CourseTitle", "").(string)],
					})
					chapterall = make(map[string]interface{})
					Score = make(map[string]string)
				},
			},
		},
	},
}

func (s *AjaxCourseMembers) getnumbers() string {
	return s.Data[0].Numbers
}
