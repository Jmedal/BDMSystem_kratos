package main

import (
	"context"
	"flag"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"os"
	"os/signal"
	"syscall"
	"time"

	"account/internal/di"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
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
	log.Info("account start")
	paladin.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	cancel, err := initDiscovery("0.0.0.0","5502", "account.service")
	if err != nil {
		panic(err)
	}
	defer cancel()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("account exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
