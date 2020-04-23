// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"grpc-test/internal/dao"
	"grpc-test/internal/server/grpc"
	"grpc-test/internal/server/http"
	"grpc-test/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
