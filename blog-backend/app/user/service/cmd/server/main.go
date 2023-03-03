package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-blog/pkg/bootstrap"
	"kratos-blog/pkg/service"
)

// go build -ldflags "-X main.Service.Version=x.y.z"

var (
	Service = bootstrap.NewServiceInfo(
		service.UserService,
		"1.0.0",
		"",
	)
)

func newApp(ll log.Logger, rr registry.Registrar, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Registrar(rr),
		kratos.Logger(ll),
		kratos.Server(
			gs,
		),
	)
}

func main() {
	// bootstrap
	cfg, ll, reg := bootstrap.Bootstrap(Service)

	app, cleanup, err := initApp(ll, reg, cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
