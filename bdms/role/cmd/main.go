package main

import (
	"flag"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"role/internal/di"
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
	log.Info("role start")
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
			log.Info("role exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
