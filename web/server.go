package web

import (
	"github.com/cbwfree/micro-game/utils/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net"
	"sync"
	"time"
)

var (
	DefaultMaxHeaderBytes = 1024 // 默认Header大小限制
)

// Web服务
type Server struct {
	sync.Mutex
	name    string
	addr    string
	echo    *echo.Echo
	running bool
	exit    chan chan error
}

func (s *Server) With(with ...ServerWith) {
	for _, v := range with {
		v(s)
	}
}

func (s *Server) Start() error {
	s.Lock()
	defer s.Unlock()

	if s.running {
		return nil
	}

	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.echo.Listener = l

	go s.echo.StartServer(s.echo.Server)

	s.exit = make(chan chan error, 1)
	s.running = true

	go func() {
		ch := <-s.exit
		ch <- s.echo.Listener.Close()
	}()

	log.Info("[%s] Web Server Start on %v", s.name, l.Addr().String())

	return nil
}

func (s *Server) Stop() {
	s.Lock()
	defer s.Unlock()

	if !s.running {
		return
	}

	ch := make(chan error, 1)
	s.exit <- ch
	s.running = false

	log.Info("[%s] Web Server Close ... ", s.name)
}

func NewServer(name string, validator ...echo.Validator) *Server {
	s := &Server{
		name: name,
		echo: echo.New(),
	}

	// Http Server 设置
	s.echo.Server.ReadTimeout = time.Duration(Opts.ReadTimeout) * time.Second
	s.echo.Server.WriteTimeout = time.Duration(Opts.WriteTimeout) * time.Second
	s.echo.Server.MaxHeaderBytes = DefaultMaxHeaderBytes

	// Echo 设置
	s.echo.HideBanner = true // 隐藏Echo的Banner
	s.echo.HidePort = true
	s.echo.HTTPErrorHandler = ErrorHandler // 统一错误处理

	if len(validator) > 0 {
		s.echo.Validator = validator[0]
	}

	if log.IsDev() {
		s.echo.Debug = true
	}

	// 记录请求日志
	s.echo.Use(middleware.Logger())
	// 从 panic 链中的任意位置恢复程序， 打印堆栈的错误信息，并将错误集中交给 HTTPErrorHandler 处理。
	s.echo.Use(middleware.Recover())

	return s
}
