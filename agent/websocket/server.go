package websocket

import (
	"github.com/cbwfree/micro-game/agent"
	"github.com/cbwfree/micro-game/utils/log"
	"github.com/cbwfree/micro-game/utils/tool"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"sync"
)

var (
	DefaultMaxHeaderBytes = 1024 // 默认Header大小
)

func init() {
	agent.RegisterAgent["websocket"] = NewServer
}

type Server struct {
	sync.Mutex
	agent    *agent.Agent
	listener net.Listener
	upgrader *websocket.Upgrader
	running  bool
	exit     chan chan error
}

// Name 服务器名称
func (s *Server) Name() string {
	return "websocket"
}

// Agent 网关对象
func (s *Server) Agent() *agent.Agent {
	return s.agent
}

// Opts 网关参数
func (s *Server) Opts() *agent.Options {
	return s.agent.Opts()
}

// Port 监听端口
func (s *Server) Port() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}

// 处理 WebSocket 连接
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	// 将HTTP请求升级为WebSocket
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade error: %v", err)
		return
	}

	// 最大连接数检查
	if s.Opts().MaxConnNum > 0 && s.agent.Count() > int(s.Opts().MaxConnNum) {
		_ = conn.Close()
		log.Warn("too many connections")
		return
	}

	// 启动客户端
	s.agent.StartClient(NewClient(s, conn, tool.GetHttpRealIP(r)))
}

// 启动
func (s *Server) Run() error {
	s.Lock()
	defer s.Unlock()

	l, err := net.Listen("tcp", s.Opts().Address)
	if err != nil {
		return err
	}

	s.listener = l
	s.upgrader = &websocket.Upgrader{
		HandshakeTimeout: s.Opts().ReadTimeout,
		CheckOrigin:      func(_ *http.Request) bool { return true },
	}

	hs := &http.Server{
		Handler:        s,
		ReadTimeout:    s.Opts().ReadTimeout,
		WriteTimeout:   s.Opts().WriteTimeout,
		MaxHeaderBytes: DefaultMaxHeaderBytes,
	}

	// 启动
	go hs.Serve(s.listener)

	s.exit = make(chan chan error, 1)
	s.running = true

	go func() {
		ch := <-s.exit
		ch <- s.listener.Close()
	}()

	log.Info("[WebSocket] Server is ready, Listening on %s:%d", agent.Opts.Host, s.Port())

	return nil
}

// 关闭
func (s *Server) Close() {
	s.Lock()
	defer s.Unlock()

	if !s.running {
		return
	}

	ch := make(chan error, 1)
	s.exit <- ch
	s.running = false

	log.Info("[WebSocket] Server is stopping.")
}

func NewServer(agent *agent.Agent) agent.Server {
	s := &Server{
		agent: agent,
	}
	return s
}
