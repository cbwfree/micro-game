package main

import (
	"context"
	"github.com/cbwfree/micro-game/agent"
	"github.com/cbwfree/micro-game/app"
	"github.com/cbwfree/micro-game/codec"
	"github.com/cbwfree/micro-game/example/libs/def"
	pcom "github.com/cbwfree/micro-game/example/proto/com"
	pgame "github.com/cbwfree/micro-game/example/proto/game"
	pgate "github.com/cbwfree/micro-game/example/proto/gate"
	rds "github.com/cbwfree/micro-game/store/redis"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/cbwfree/micro-game/utils/log"
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"

	_ "github.com/cbwfree/micro-game/agent/quic"
	_ "github.com/cbwfree/micro-game/agent/tcp"
	_ "github.com/cbwfree/micro-game/agent/websocket"
)

var (
	gate *agent.Agent // 网关服务
)

// 入口
func main() {
	// 创建服务
	app.New(def.SrvGateName, def.Version, rds.Flags, agent.Flags)

	// 初始化
	app.Init(
		micro.BeforeStart(beforeStart),
		micro.AfterStart(afterStart),
		micro.BeforeStop(beforeStop),
		micro.AfterStop(afterStop),
	)

	// 注册RPC服务
	app.AddHandler(new(GateService))

	// 启动服务
	if err := app.Run(); err != nil {
		log.Fatal("%+v", err)
	}
}

func beforeStart() error {
	// 启用并连接Redis
	if err := rds.Connect(); err != nil {
		return errors.Wrap(err, "redis connect error")
	}

	// 启动网关服务
	if err := agentRun(); err != nil {
		return err
	}

	// 注册网关信息
	app.AddMetadata(map[string]string{
		"type":  agent.Opts.Type,
		"agent": gate.Address(),
	})

	return nil
}

func afterStart() error {
	return nil
}

func beforeStop() error {
	// 关闭服务器
	gate.Close()

	// 关闭Redis
	if err := rds.Disconnect(); err != nil {
		log.Error("关闭Redis服务出错: %s", err.Error())
	}

	return nil
}

func afterStop() error {
	return nil
}

// --------------
type GateService struct{}

// 推送消息
func (*GateService) Push(_ context.Context, req *pgate.InGatePush, _ *pgate.OutGatePush) error {
	client := gate.GetClient(req.ClientId)
	if client == nil {
		return errors.Errorf("client [%s] not found", req.ClientId)
	}

	b, err := gate.ServerCodec().Marshal(&codec.ServerHead{Cmd: req.Cmd, Code: req.Code}, req.Data)
	if err != nil {
		return err
	}

	client.Write(b)
	return nil
}

// ---------------
func agentRun() error {
	gate = agent.NewAgent(nil)
	gate.SetOnReceive(onReceiveHandler)
	gate.SetOnDisconnect(onDisconnectHandler)
	return gate.Run()
}

// OnReceiveHandler 收到数据处理 (返回错误就会断开连接)
func onReceiveHandler(client agent.Client, head *codec.ClientHead, data []byte) (*codec.ServerHead, []byte, error) {
	// 转发消息到路由
	in := &pgame.InForwardProtocol{
		Cmd:  head.Cmd,
		Data: data,
	}
	out := new(pgame.OutForwardProtocol)
	if err := app.CallCtx(client.Meta().Context(), def.SrvGameName, pgame.ForwardMethod_Protocol, in, out); err != nil {
		client.Log().Warn(color.Warn.Text("[OnReceive] Forward Protocol [%d] error: %s", head.Cmd, err))
		return nil, nil, err
	}

	return &codec.ServerHead{
		Serial: head.Serial,
		Cmd:    head.Cmd,
		Code:   out.Code,
	}, out.Data, nil
}

// OnDisconnectHandler 连接断开时触发
func onDisconnectHandler(client agent.Client) {
	if !client.Meta().IsOnline() {
		return
	}

	if err := app.CallCtx(client.Meta().Context(), def.SrvGameName, pgame.ForwardMethod_Offline, &pcom.None{}, &pcom.None{}); err != nil {
		client.Log().Warn(color.Warn.Text("forward offline error: %s", err))
	}
}
