package web

import (
	"github.com/cbwfree/micro-game/utils/tool"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"path"
)

type ServerWith func(s *Server)

func WithDebug(debug bool) ServerWith {
	return func(s *Server) {
		s.echo.Debug = debug // 调试模式
	}
}

// 跨域设置
func WithCORS(origins ...string) ServerWith {
	return func(s *Server) {
		cors := middleware.CORSConfig{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
			AllowHeaders:  []string{echo.HeaderAccept, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization, "Time", "Sign"},
			ExposeHeaders: []string{echo.HeaderAuthorization},
		}
		if len(origins) > 0 {
			cors.AllowOrigins = origins
			if len(origins) == 1 && origins[0] != "*" && origins[0] != "" {
				cors.AllowCredentials = true
			}
		}
		s.echo.Use(middleware.CORSWithConfig(cors))
	}
}

func WithAddr(addr string) ServerWith {
	return func(s *Server) {
		s.addr = addr
	}
}

// 启用GZIP
func WithGzip() ServerWith {
	return func(s *Server) {
		s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
	}
}

func WithSecure() ServerWith {
	return func(s *Server) {
		s.echo.Use(middleware.Secure())
	}
}

func WithValidator(validator echo.Validator) ServerWith {
	return func(s *Server) {
		s.echo.Validator = validator
	}
}

// WithRouter 注册路由
func WithRouter(prefix string, routes ...func(g *echo.Group)) ServerWith {
	return func(s *Server) {
		g := s.echo.Group(prefix)
		for _, route := range routes {
			route(g)
		}
	}
}

// 绑定视图
func WithView(prefix string, fs http.FileSystem, uri ...string) ServerWith {
	return func(s *Server) {
		handler := echo.WrapHandler(http.StripPrefix(prefix, http.FileServer(fs)))
		s.echo.GET(prefix, handler)
		for _, u := range uri {
			s.echo.GET(path.Join(prefix, u), handler)
		}
	}
}

// 绑定资源访问
func WithAssets(prefix string, handler echo.HandlerFunc) ServerWith {
	return func(s *Server) {
		s.echo.Group(prefix).GET("/*", handler)
	}
}

// 启用Session
func WithSession(store string, secret string) ServerWith {
	return func(s *Server) {
		if err := tool.InitFolder(store, 0755); err != nil {
			log.Fatalf("Enable Web Session Error: %s", err)
			return
		}

		s.echo.Use(session.Middleware(
			sessions.NewFilesystemStore(store, []byte(secret)),
		))
	}
}
