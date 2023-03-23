package config

import (
	handlers "GRPC-AUTH/cmd/handler"
	implement_user "GRPC-AUTH/internal/domain/repository/implement/user"

	"log"

	pb_user "GRPC-AUTH/internal/infra/proto/user"

	"flag"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

var router *echo.Echo

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "../data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	Setup(configPath)
}

func Run(s *grpc.Server, configPath string) *grpc.Server {
	conf := GetConfig()

	router = echo.New()
	router.Start(":" + conf.Server.Port)
	log.Println("Server started on [::]:" + conf.Server.Port)

	setupDB(conf)
	setupStorageClient(conf)
	pb_user.RegisterUserServiceServer(s, handlers.NewServerUser(implement_user.UserRepository(DB), DB))

	return s
}
