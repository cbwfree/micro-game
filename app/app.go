package app

import (
	"context"
	"fmt"
	"github.com/google/gops/agent"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/transport"
	"sync"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"

	cgrpc "github.com/micro/go-micro/v2/client/grpc"
	sgrpc "github.com/micro/go-micro/v2/server/grpc"
	tgrpc "github.com/micro/go-micro/v2/transport/grpc"

	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/cbwfree/micro-game/utils/log"
	"github.com/cbwfree/micro-game/utils/pb"
)

var (
	appInstance *app
	srvOnce     sync.Once
)

// 服务单例
func APP() *app {
	srvOnce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		appInstance = &app{
			ctx:       ctx,
			cancel:    cancel,
			publisher: make(map[string]micro.Publisher),
		}
	})
	return appInstance
}

// 当前服务
type app struct {
	sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	srv       micro.Service
	publisher map[string]micro.Publisher // 订阅
}

func (a *app) Service() micro.Service {
	return a.srv
}

func (a *app) Cancel() {
	a.cancel()
}

func (a *app) newService(name string, version string, flags []cli.Flag) {
	var opts = []micro.Option{
		micro.Context(a.ctx),
		micro.Cmd(cmd.NewCmd(
			cmd.Name(name),
			cmd.Description(fmt.Sprintf("a %s service", name)),
			cmd.Version(version),
		)),
		micro.Client(cgrpc.NewClient()),
		micro.Server(sgrpc.NewServer(
			server.Name(name),
			server.Version(version),
		)),
		micro.Transport(tgrpc.NewTransport(transport.Secure(true))),
		micro.Broker(nats.NewBroker()),     // nats
		micro.Registry(etcd.NewRegistry()), // ectd
		micro.BeforeStart(func() error {
			if Opts.PsAddr != "" {
				if err := agent.Listen(agent.Options{Addr: Opts.PsAddr}); err != nil {
					return err
				}
			}
			return nil
		}),
	}
	a.srv = micro.NewService(opts...)

	// 重载Cmd参数, 省略用不上的参数
	a.srv.Options().Cmd.App().Flags = flags
}

func (a *app) initService(opts ...micro.Option) {
	a.srv.Init(opts...)
}

// New 创建服务
func New(name string, version string, extra ...[]cli.Flag) {
	if APP().srv != nil {
		return
	}

	// 合并flags
	var flags = defaultFlags
	if len(extra) > 0 {
		for _, ex := range extra {
			flags = append(flags, ex...)
		}
	}

	// 创建服务
	APP().newService(name, version, flags)

	// 关闭事件
	_ = AddSub("cancel", func(_ context.Context, c *pb.Cancel) error {
		if c.Name != Name() || (c.NodeId != "" && c.NodeId != Id()) {
			return nil
		}

		log.Debug("[ServiceCancel] service [%s][%s] is being stopped ...", Name(), Id())

		// 停止服务
		APP().Cancel()

		return nil
	})
}

// 启动服务
func Init(opts ...micro.Option) {
	APP().initService(opts...)
}

// Service 获取服务对象
func Srv() micro.Service {
	return APP().srv
}

// Client 获取客户端
func Client() client.Client {
	return Srv().Client()
}

// Server 获取服务端
func Server() server.Server {
	return Srv().Server()
}

// 启动服务
func Run() error {
	return Srv().Run()
}

// 关闭服务
func Close() {
	APP().Cancel()
}

// 取消服务
func Cancel(name string, id ...string) error {
	in := &pb.Cancel{Name: name}
	if len(id) > 0 {
		in.NodeId = id[0]
	}
	return Pub("cancel", in)
}

// Version 获取服务版本
func Version() string {
	return Server().Options().Version
}

// Id 获取服务UUID
func Id() string {
	return Server().Options().Id
}

// Name 获取服务名称
func Name() string {
	return Server().Options().Name
}

// NameId 获取服务节点ID
func NameId() string {
	return fmt.Sprintf("%s-%s", Name(), Id())
}

// 获取服务发现注册信息
func Registry() registry.Registry {
	return APP().Service().Options().Registry
}

// 获取服务节点列表
func GetServices(name string) []*registry.Service {
	res, _ := Registry().GetService(name)
	return res
}

// 注册发布者
func AddPub(names ...string) {
	for _, name := range names {
		APP().publisher[name] = micro.NewEvent(name, Client())
	}
}

// 获取发布者
func GetPub(name string) micro.Publisher {
	return APP().publisher[name]
}

// 发布消息
func Pub(name string, msg interface{}, opts ...client.PublishOption) error {
	return PubCtx(context.TODO(), name, msg, opts...)
}
func PubCtx(ctx context.Context, name string, msg interface{}, opts ...client.PublishOption) error {
	return GetPub(name).Publish(ctx, msg, opts...)
}

// 注册订阅
func AddSub(name string, h interface{}, queue ...bool) error {
	var opts = []server.SubscriberOption{
		server.InternalSubscriber(true),
	}
	if len(queue) > 0 && queue[0] {
		opts = append(opts, server.SubscriberQueue(name))
	}

	return micro.RegisterSubscriber(name, Server(), h, opts...)
}

func AddSubQueue(name string, h interface{}) error {
	return AddSub(name, h, true)
}

// 注册RPC服务
func AddHandler(handles ...interface{}) {
	opts := []server.HandlerOption{
		server.InternalHandler(true),
	}

	for _, h := range handles {
		if err := micro.RegisterHandler(Server(), h, opts...); err != nil {
			log.Error("RegisterHandler Error: %s", err.Error())
		}
	}
}

// 增加meta信息
func AddMetadata(meta map[string]string) {
	for k, v := range meta {
		Server().Options().Metadata[k] = v
	}
}

// Call 通过名称调用RPC
// 	@name 服务名称
// 	@method rpc方法名称. 即 serviceName.rpcName
// 	@in 请求参数
// 	@out 返回数据
// 	@nodeId 指定节点ID
func CallCtx(ctx context.Context, srvName string, method string, in interface{}, out interface{}, opts ...client.CallOption) error {
	req := Client().NewRequest(srvName, method, in)
	return Client().Call(ctx, req, out, opts...)
}

// CallCtx 通过名称调用RPC
// 	@name 服务名称
// 	@method rpc方法名称. 即 serviceName.rpcName
// 	@in 请求参数
// 	@out 返回数据
// 	@nodeId 指定节点ID
func CallNode(ctx context.Context, srvName string, method string, in interface{}, out interface{}, nodeId ...string) error {
	var opts []client.CallOption
	if len(nodeId) > 0 {
		opts = append(opts, FilterServiceNode(srvName, nodeId[0]))
	}
	return CallCtx(ctx, srvName, method, in, out, opts...)
}
