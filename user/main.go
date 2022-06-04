package main

import (
	"user/conf"
	"user/core"
	services "user/services"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	conf.Init()
	//etcd 注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),
	)

	microService.Init()
	//服务注册
	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))

	//启动服务
	microService.Run()
}
