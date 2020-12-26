package tcp

import (
	"github.com/cbwfree/micro-game/agent"
	"github.com/cbwfree/micro-game/utils/log"
	"net"
	"sync"
	"time"
)

func init() {
	agent.RegisterAgent["tcp"] = NewServer
}

type Server struct {
	sync.Mutex
	wg       *sync.WaitGroup
	agent    *agent.Agent
	listener net.Listener
	running  bool
	exit     chan chan error
}

// Name 服务器名称
func (s *Server) Name() string {
	return "tcp"
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

// 启动服务
func (s *Server) start() {
	var delay time.Duration
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if delay == 0 {
					delay = 5 * time.Millisecond
				} else {
					delay *= 2
				}
				if max := 1 * time.Second; delay > max {
					delay = max
				}

				log.Warn("accept error: %v; retrying in %v", err, delay)
				time.Sleep(delay)
				continue
			}
			log.Warn("accept error: %v", err)
			return
		}

		delay = 0

		// 最大连接数检查
		if s.Opts().MaxConnNum > 0 && s.agent.Count() > int(s.Opts().MaxConnNum) {
			_ = conn.Close()
			log.Warn("too many connections")
			continue
		}

		go func() {
			s.wg.Add(1)
			defer s.wg.Done()

			ip := conn.RemoteAddr().(*net.TCPAddr).IP.String()
			s.agent.StartClient(NewClient(s, conn, ip))
		}()
	}
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

	// 启动
	go s.start()

	s.exit = make(chan chan error, 1)
	s.running = true

	go func() {
		ch := <-s.exit
		ch <- s.listener.Close()
	}()

	log.Info("[TCP] Server is ready, Listening on %s:%d", agent.Opts.Host, s.Port())

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

	s.wg.Wait()

	log.Info("[TCP] Server is stopping.")
}

func NewServer(agent *agent.Agent) agent.Server {
	s := &Server{
		wg:    new(sync.WaitGroup),
		agent: agent,
	}
	return s
}
