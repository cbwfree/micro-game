package agent

import (
	"fmt"
	"github.com/cbwfree/micro-game/codec"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/pkg/errors"
	"sync"
)

var (
	RegisterAgent = map[string]func(agent *Agent) Server{}
)

type Agent struct {
	sync.RWMutex
	opts        *Options
	server      Server
	clientCodec *codec.Client
	serverCodec *codec.Server
	clients     map[string]Client

	OnReceive    func(Client, *codec.ClientHead, []byte) (*codec.ServerHead, []byte, error) // 收到数据调用
	OnDisconnect func(Client)                                                               // 连接断开时调用
}

func (g *Agent) Opts() *Options {
	return g.opts
}

func (g *Agent) Server() Server {
	return g.server
}

// Address 网关地址
func (g *Agent) Address() string {
	return fmt.Sprintf("%s:%d", Opts.Host, g.server.Port())
}

func (g *Agent) ClientCodec() *codec.Client {
	return g.clientCodec
}

func (g *Agent) ServerCodec() *codec.Server {
	return g.serverCodec
}

func (g *Agent) SetOnReceive(fn func(Client, *codec.ClientHead, []byte) (*codec.ServerHead, []byte, error)) {
	g.OnReceive = fn
}

func (g *Agent) SetOnDisconnect(fn func(Client)) {
	g.OnDisconnect = fn
}

func (g *Agent) StartClient(client Client) {
	g.Lock()
	g.clients[client.Id()] = client
	g.Unlock()

	client.Log().Debugf("connected ...")

	// 接收消息处理
	for {
		// 接收消息
		cHead, cData, err := client.Read()
		if err != nil {
			client.Log().Warn(color.Warn.Text("read data error: %s", err))
			break
		}

		// 处理接收的消息
		sHead, sData, err := g.OnReceive(client, cHead, cData)
		if err != nil {
			break
		}

		// 响应消息
		if sHead.Code > 0 || len(sData) > 0 {
			b, err := g.ServerCodec().Marshal(sHead, sData)
			if err != nil {
				client.Log().Warn(color.Warn.Text("marshal data error: %s", err))
				break
			}

			client.Write(b)
		}
	}

	client.Log().Debug("disconnected ...")

	client.Close()

	g.Lock()
	delete(g.clients, client.Id())
	g.Unlock()

	g.OnDisconnect(client) // 连接断开处理
}

// 获取客户端连接对象
func (g *Agent) GetClient(val string, by ...string) Client {
	g.RLock()
	defer g.RUnlock()

	if len(by) > 0 {
		byKey := by[0]
		for _, client := range g.clients {
			if client.Meta().Get(byKey) == val {
				return client
			}
		}
	} else {
		if client, ok := g.clients[val]; ok {
			return client
		}
	}

	return nil
}

// 连接数
func (g *Agent) Count() int {
	g.RLock()
	defer g.RUnlock()

	return len(g.clients)
}

// 全部客户端
func (g *Agent) All() map[string]Client {
	g.RLock()
	defer g.RUnlock()

	clients := make(map[string]Client, len(g.clients))
	for k, v := range g.clients {
		clients[k] = v
	}

	return clients
}

// 迭代客户端
func (g *Agent) ForEach(fn func(client Client) bool) {
	g.RLock()
	defer g.RUnlock()

	for _, client := range g.clients {
		if !fn(client) {
			break
		}
	}
}

// 过滤角色对象
func (g *Agent) Filter(filter func(client Client) bool) map[string]Client {
	g.RLock()
	defer g.RUnlock()

	var clients = make(map[string]Client)
	for id, client := range g.clients {
		if filter == nil || filter(client) {
			clients[id] = client
		}
	}

	return clients
}

// 广播
func (g *Agent) Broadcast(msg []byte, filter func(client Client) bool) {
	g.RLock()
	defer g.RUnlock()

	for _, client := range g.clients {
		if !client.Meta().IsOnline() {
			continue
		}

		if filter == nil || filter(client) {
			client := client
			go client.Write(msg)
		}
	}
}

// 启动网关服务
func (g *Agent) Run() error {
	g.Lock()
	defer g.Unlock()

	if g.server != nil {
		return errors.Errorf("agent server is set %s ...", g.server.Name())
	}

	if Opts.Type == "ws" || Opts.Type == "wss" {
		Opts.Type = "websocket"
	}

	if s, ok := RegisterAgent[Opts.Type]; ok {
		g.server = s(g)
	} else {
		return errors.Errorf("Unsupported agent server type: %s", Opts.Type)
	}

	return g.server.Run()
}

// 关闭网关服务
func (g *Agent) Close() {
	g.Lock()
	defer g.Unlock()

	if g.server == nil {
		return
	}

	g.server.Close()

	for _, client := range g.clients {
		g.OnDisconnect(client)
		client.Destroy()
		delete(g.clients, client.Id())
	}
}

// NewAgent
func NewAgent(codecMix []uint8, opts ...Option) *Agent {
	g := &Agent{
		opts:        NewOptions(opts...),
		clientCodec: codec.NewClient(codecMix...),
		serverCodec: codec.NewServer(codecMix...),
		clients:     make(map[string]Client),
	}
	return g
}
