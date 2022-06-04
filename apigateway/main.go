package main

import (
	apiservice "apigateway/services"
	"apigateway/weblib"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//user
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
	)

	// 用户服务调用实例
	userService := apiservice.NewUserService("rpcUserService", userMicroService.Client())

	//创建微服务实例，使用gin暴露http接口并注册到etcd
	service := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		//服务调用实例使用gin处理
		web.Handler(weblib.NewRouter(userService)), //设置路由
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),      //设置注册服务的过期时间
		web.RegisterInterval(time.Second*15), //设置间隔多久再次注册服务
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	//运行吧
	_ = service.Init()
	_ = service.Run()
}
