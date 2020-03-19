package main

import (
	"context"
	"course/internal/di"
	"flag"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initDiscovery(ip, port, appID string) (cancelFunc context.CancelFunc, err error) {
	// NOTE: 必须拿到您实例节点的真实IP，
	// NOTE: 必须拿到您实例grpc监听的真实端口，warden默认监听9000
	hn, _ := os.Hostname()
	dis := discovery.New(nil)
	ins := &naming.Instance{
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		AppID:    appID,
		Hostname: hn,
		Addrs: []string{
			"grpc://" + ip + ":" + port,
		},
	}
	return dis.Register(context.Background(), ins)
}

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("course start")
	paladin.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	//cancel, err := initDiscovery("0.0.0.0","9000", "****.service")
	//if err != nil {
	//	panic(err)
	//}
	//defer cancel()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("course exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

//测试test,时间转换和比较

//func main()  {
//	toBeCharge := "2020-02-23 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
//	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
//	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
//	secondTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
//
//	fmt.Println(secondTime.Unix())
//
//
//	//secondTime := time.Now()
//	fmt.Println(secondTime)
//	currentYear, currentMonth, _ := secondTime.Date()
//	fmt.Println(currentYear,currentMonth)
//	currentLocation := secondTime.Location()
//	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
//	fmt.Println(firstOfMonth.Format("2006-01-02"))
//	//下个月第一天
//	fmt.Println(firstOfMonth.AddDate(0, 1, 0).Format("2006-01-02"))
//	//这个月最后一天
//	fmt.Println(firstOfMonth.AddDate(0, 1, -1).Format("2006-01-02"))
//
//
//	//one的时间在secondTime之前返回false，one的时间在secondTime之后或者同一天返回true
//	one :=time.Now()
//	fmt.Println(one.After(secondTime))
//
//
//
//}
