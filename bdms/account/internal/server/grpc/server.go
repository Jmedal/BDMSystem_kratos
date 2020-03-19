package grpc

import (
	pb "account/api"
	"account/internal/dao"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

var svd dao.Dao

// New new a grpc server.
func New(svc pb.AccountServer, d dao.Dao) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svd = d
	ws = warden.NewServer(&cfg)
	ws.Use(checkToken())
	pb.RegisterAccountServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
