package app

import (
	"github.com/micro/cli/v2"
)

var (
	Opts = new(struct {
		PsAddr string // gops分析
		Lang   string // 设置语言环境
	})

	defaultFlags = []cli.Flag{
		&cli.StringFlag{
			Name:        "gops",
			Usage:       "设置gops分析地址",
			EnvVars:     []string{"GAME_GOPS"},
			Destination: &Opts.PsAddr,
		},
		&cli.StringFlag{
			Name:        "lang",
			Value:       "zh-cn",
			Usage:       "设置当前语言环境.",
			EnvVars:     []string{"GAME_LANG"},
			Destination: &Opts.Lang,
		},
		&cli.StringFlag{
			Name:    "profile",
			Usage:   "Debug profiler for cpu and memory stats",
			EnvVars: []string{"GAME_DEBUG_PROFILE"},
		},
		&cli.StringFlag{
			Name:    "runtime",
			Usage:   "Runtime for building and running services e.g local, kubernetes",
			EnvVars: []string{"GAME_RUNTIME"},
			Value:   "local",
		},
		&cli.StringFlag{
			Name:    "runtime_source",
			Usage:   "Runtime source for building and running services e.g github.com/micro/service",
			EnvVars: []string{"GAME_RUNTIME_SOURCE"},
			Value:   "github.com/micro/services",
		},
		&cli.StringFlag{
			Name:    "client_request_timeout",
			Value:   "10s",
			EnvVars: []string{"GAME_CLIENT_REQUEST_TIMEOUT"},
			Usage:   "设置客户端请求超时. e.g 500ms, 5s, 1m. Default: 5s",
		},
		&cli.IntFlag{
			Name:    "client_pool_size",
			EnvVars: []string{"GAME_CLIENT_POOL_SIZE"},
			Usage:   "设置客户端连接池大小. Default: 1",
		},
		&cli.StringFlag{
			Name:    "client_pool_ttl",
			EnvVars: []string{"GAME_CLIENT_POOL_TTL"},
			Usage:   "设置客户端连接TTL. e.g 500ms, 5s, 1m. Default: 1m",
		},
		&cli.StringFlag{
			Name:    "registry_address",
			EnvVars: []string{"GAME_REGISTRY_ADDRESS"},
			Usage:   "服务发现地址. 以逗号分隔",
		},
		&cli.IntFlag{
			Name:    "register_ttl",
			EnvVars: []string{"GAME_REGISTER_TTL"},
			Value:   60,
			Usage:   "服务发现注册TTL",
		},
		&cli.IntFlag{
			Name:    "register_interval",
			EnvVars: []string{"GAME_REGISTER_INTERVAL"},
			Value:   30,
			Usage:   "服务发现注册时间间隔",
		},
		&cli.StringFlag{
			Name:    "broker_address",
			EnvVars: []string{"GAME_BROKER_ADDRESS"},
			Usage:   "设置代理地址. 以逗号分隔",
		},
		&cli.StringFlag{
			Name:    "selector",
			EnvVars: []string{"GAME_SELECTOR"},
			Usage:   "设置用于选择节点进行查询的选择器",
		},
		&cli.StringFlag{
			Name:    "tracer",
			EnvVars: []string{"GAME_TRACER"},
			Usage:   "设置分布式跟踪的跟踪器, e.g. memory, jaeger",
		},
		&cli.StringFlag{
			Name:    "tracer_address",
			EnvVars: []string{"GAME_TRACER_ADDRESS"},
			Usage:   "设置分布式跟踪的跟踪器地址. 以逗号分隔",
		},
	}
)
