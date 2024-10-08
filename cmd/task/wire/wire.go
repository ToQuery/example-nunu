//go:build wireinject
// +build wireinject

package wire

import (
	"example-nunu/internal/repository"
	"example-nunu/internal/server"
	"example-nunu/internal/service"
	"example-nunu/pkg/app"
	"example-nunu/pkg/log"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	//repository.NewUserRepository,
	repository.NewTqAppRepository,
	repository.NewTqDeveloperRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	//service.NewUserService,
	service.NewTqAppService,
	service.NewTqDeveloperService,
)

var serverSet = wire.NewSet(
	server.NewTask,
)

// build App
func newApp(
	task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		serverSet,
		newApp,
	))
}
