# micro-game

基于 [go-micro v2.9](https://github.com/asim/go-micro) 实现的游戏服务端框架

## 简介

RPC和微服务的框架很多, 不想重复造轮子, 选择了基于 go-micro 来做游戏服务端.

为什么选择 go-micro 呢, 因为 go-micro 的可插拔性非常高, 源码学习起来也相对简单. 再一个就是我比较不喜欢写配置文件, 能用参数或者环境变量解决的, 就坚决不写配置文件. 😉 

配合 [direnv](https://github.com/direnv/direnv) 工具使用, 可以非常方便的切换环境变量 

在不改变 go-micro 源码的情况下, 我对其进行了二次封装, 去掉一些在游戏服务端用不上的功能, 简化了编码方式. 目前开源版框架主要实现了一下功能:

- 网关 (包括 Websocket / TCP / QUIC)
- Web服务 (Web服务主要是用来做管理服务的, 配合 [go-bindata](github.com/shuLhan/go-bindata/cmd/go-bindata) 工具, 可以将静态页面打包为 .go 文件)
- 游戏协议编码 (协议编码使用的是差异化模式, 即客户端请求头和服务端响应头不一致. 同时也支持自定义协议头混淆和校验)
- 游戏协议请求 (主要实现了游戏协议处理函数的注册和自动调用, 会自动解析请求内容为具体的协议结构)
- 数据库 (目前仅实现了 MongoDB 的支持, 仅支持 4.2+ 版本. 驱动 [go.mongodb.org/mongo-driver/mongo](go.mongodb.org/mongo-driver/mongo))
- 缓存 (目前仅实现了 Redis 的支持, 支持单节点和 Redis Proxy. 驱动 [github.com/go-redis/redis/v8](github.com/go-redis/redis/v8))
- 分布式锁 (基于 redis 的分布式锁)
- 内存数据管理器 ( store/memory )

在 utils 包下实现了一些常用的工具:

- auth (token鉴权, 仅有 jwt)
- dtype (数据类型转换)
- errors (错误封装, 在RPC请求时请使用此错误返回)
- formula (数据公式解析引擎, 来自 [https://github.com/dengsgo/math-engine](https://github.com/dengsgo/math-engine))
- log (日志输出)
- multi (并发执行多任务)
- tool (一些常用的工具函数库)

## 关于 Meta

Meta 是对 go-micro 的 metadata.Metadata 的再次封装, 实际上它的底层依旧是 context.Context

Meta 主要是用来查找指定的服务节点的, 需要配合 `proto-gen-rpc` 使用

DbMeta 是 Meta 的拓展, 主要表示的是数据服务, 其中携带了标识该数据的 `Data-Id` 信息

## 运行环境

在 `app` 包中, 固定了 go-micro 的 `Broker`, `Registry` , `Server` , `Client` 类型. 

在实际使用中对于热插拔的要求并不高. 同时为了减少运行时参数数量, 所以这里把这些类型都给固定了.

- `Broker` 固定为 `nats` [http://nats.io/](http://nats.io/)
- `Registry` 固定为 `etcd` [http://etcd.io/](http://etcd.io/)
- `Server` 和 `Client` 均固定为 `grpc`

## 安装

```shell
go get github.com/cbwfree/micro-game
```

## 创建服务

```golang
package main

import (
	"context"
	"github.com/cbwfree/micro-game/app"
	"github.com/cbwfree/micro-game/example/libs/def"
	mgo "github.com/cbwfree/micro-game/store/mongo"
	rds "github.com/cbwfree/micro-game/store/redis"
	"github.com/cbwfree/micro-game/utils/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/micro/go-micro/v2"
)

type Service struct {}

func (*Service) Call(ctx context.Context, _ *empty.Empty, _ *empty.Empty) error {
	log.Debug("rpc request ...")
	return nil
}

func main() {
	// 初始化服务
    app.New("ServiceName", "latest", mgo.Flags, rds.Flags)
    
    // 初始化
    app.Init(
        micro.BeforeStart(beforeStart),
        micro.AfterStart(afterStart),
        micro.BeforeStop(beforeStop),
        micro.AfterStop(afterStop),
    )
    
    // 注册RPC服务
    app.AddHandler()
    
    // 启动服务
    if err := app.Run(); err != nil {
        log.Fatal("%+v", err)
    }
}


// beforeStart 启动之前执行
func beforeStart() error {
	// 连接Redis
	if err := rds.Connect(); err != nil {
		return err
	}
	// 连接MongoDB
	if err := mgo.Connect(); err != nil {
		return err
	}

	return nil
}

func afterStart() error {
	// 发起RPC请求
	app.CallCtx(context.TODO(), "ServiceName", "Service.Call", &empty.Empty{}, &empty.Empty{})
	return nil
}

func beforeStop() error {
	return nil
}

func afterStop() error {
	// 关闭Redis
	if err := rds.Disconnect(); err != nil {
		log.Error("关闭Redis连接出错: %s", err.Error())
	}

	// 关闭MongoDB连接
	if err := mgo.Disconnect(); err != nil {
		log.Error("关闭MongoDB连接出错: %s", err.Error())
	}

	return nil
}
```

## proto-gen-rpc

[proto-gen-rpc](https://github.com/cbwfree/micro-game/tree/main/proto-gen-rpc) 是用来生成自定义的 RPC 代码的

使用时, 需要把 `proto-gen-rpc` 文件放置在 `$GOROOT/bin` 或 `$GOPATH/bin` 目录下, 同时在 `protoc` 命令后增加参数:

```shell
protoc --rpc_out=paths=source_relative:.
	
# 完整示例
protoc --go_out=paths=source_relative:. --rpc_out=paths=source_relative:. --proto_path=./ ./proto/*.proto
```

例如 `forward.proto` 文件
```protobuf
syntax = "proto3";

package proto;

// 网关消息转发
service Forward {
    rpc Protocol(InForwardProtocol) returns (OutForwardProtocol);                   // 转发协议
}

message InForwardProtocol {
    uint32 Cmd = 1;
    bytes Data = 2;
}
message OutForwardProtocol {
    uint32 Code = 1;
    bytes Data = 2;
}
```

将会生成 `forward.pb.rpc.go` 文件:
```go
// Code generated by protoc-gen-rpc. DO NOT EDIT.
// source: proto/forward.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import "context"
import "github.com/cbwfree/micro-game/app"
import "github.com/cbwfree/micro-game/meta"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Forward RPC Method constant definition
const (
	ForwardMethod_Protocol = "Forward.Protocol"
)

// Forward RPC Service
type ForwardService interface {
	Protocol(ctx context.Context, req *InForwardProtocol, rsp *OutForwardProtocol) error
}

// Forward RPC Client
type ForwardClient struct {
	name string
}

// CallNode General call method
func (c *ForwardClient) CallNode(ctx context.Context, method string, in interface{}, out interface{}, nodeId ...string) error {
	return app.CallNode(ctx, c.name, method, in, out, nodeId...)
}
func (c *ForwardClient) Protocol(mt meta.NodeMeta, in *InForwardProtocol) (*OutForwardProtocol, error) {
	out := new(OutForwardProtocol)
	err := app.CallNode(mt.Context(), c.name, ForwardMethod_Protocol, in, out, mt.Id())
	return out, err
}

func NewForwardClient(name string) *ForwardClient {
	return &ForwardClient{
		name: name,
	}
}

```

## 示例代码

[example](https://github.com/cbwfree/micro-game/tree/main/example)