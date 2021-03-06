package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	db "shopping/user/database"
	"shopping/user/handler"
	"shopping/user/repository"

	proto "shopping/user/proto/user"
)

func main() {

	db, err := db.InitDB()
	if err != nil {
		log.Fatal("connection error : %v \n", err)
	}
	defer db.Close()

	repo := &repository.User{Db: db}

	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterUserServiceHandler(service.Server(), &handler.User{Repo: repo})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
