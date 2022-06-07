# Dome
microv2+vue+pytest


- 用户注册登录 ( jwt-go鉴权 )

# 项目主要依赖：

**Golang V1.16**

- Gin
- Gorm
- mysql
- go-micro
- protobuf
- grpc
- amqp
- ini
- hystrix
- jwt-go
- crypto

# 项目结构

## 1. gateway 网关部分

```
gateway/
├── pkg
│  ├── e
│  ├── logging
│  └── util
├── services
│  └── proto
├── weblib
│  ├── handlers
│  └──  middleware
└── wrappers
```
- pkg/e : 封装错误码
- pkg/logging : 日志文件
- pkg/util : 工具函数
- service/proto : 放置proto文件以及生成的pb文件
- weblib/handlers : 各个服务的接口
- weblib/middleware : http服务器的中间件
- wrappers : 放置服务熔断的配置

## 2. vue_web vue前端开发

```
├─src
|  ├─views
|  ├─utils
|  ├─store
|  ├─router
|  ├─config
|  ├─components
|  ├─assets
|  ├─api
├─public
```



## 3. user

```
user/
├── conf
├── core
├── model
└── service
```

- conf：配置信息
- core：业务逻辑
- model：数据库模型
- service：proto文件以及各服务

# 运行简要说明
1. 保证rabbitMQ开启状态
2. 保证etcd开启状态
3. 依次执行各模块下的main.go文件
4. 执行user,task,api-gateway的时候需要后面加上这个，注册到etcd并且注册地址是这个地址。
```go
//go run main.go --registry=etcd --registry_address=127.0.0.1:2379
```
# protoc cmd
protoc --proto_path=. --micro_out=. --go_out=. userModels.proto
protoc --proto_path=. --micro_out=. --go_out=. userService.proto

protoc --proto_path=. --micro_out=. --go_out=. taskModels.proto
protoc --proto_path=. --micro_out=. --go_out=. taskService.proto

# protoc-go-inject-tag 转化
protoc-go-inject-tag -input="*.pb.go"


**如果出错一定要注意打开etcd的keeper查看服务是否注册到etcd中。**









