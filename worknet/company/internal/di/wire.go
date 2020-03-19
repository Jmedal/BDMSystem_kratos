// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"company/internal/dao"
	"company/internal/service"
	"company/internal/server/grpc"
	"company/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
