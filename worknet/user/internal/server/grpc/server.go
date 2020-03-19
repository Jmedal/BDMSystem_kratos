package grpc

import (
	pb "user/api"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
)

// New new a grpc server.
func New(svc pb.WorknetUserServer) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(&cfg)
	ws.Use(checkToken())
	pb.RegisterWorknetUserServer(ws.Server(), svc)
	ws, err = ws.Start()
	return
}
