package agent

import (
	"github.com/micro/cli/v2"
)

var (
	Opts = &struct {
		Type              string // 网关类型, websocket / tcp / quic
		Host              string // 主机IP地址
		Port              string // 网关监听端口
		ConnMaxNum        uint   // 网关最大连接
		HeartbeatInterval int64  // 网关 WebSocket 心跳间隔
		HeartbeatDeadline int64  // 网关 WebSocket 心跳等待
		AuthTime          int64  // 网关 鉴权认证 有效时间
	}{}

	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "agent_type",
			Value:       "websocket",
			Usage:       "设置网关类型. 目前支持 websocket, tcp, quic",
			EnvVars:     []string{"GAME_AGENT_TYPE"},
			Destination: &Opts.Type,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "agent_host",
			Value:       "",
			Usage:       "设置当前网关的外网IP地址",
			EnvVars:     []string{"GAME_AGENT_HOST"},
			Destination: &Opts.Host,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "agent_port",
			Value:       ":0",
			Usage:       "设置当前网关的端口. default :0",
			Destination: &Opts.Port,
		},
		&cli.Int64Flag{
			Name:        "agent_heartbeat_interval",
			Value:       30,
			Usage:       "设置当前网关的心跳的时间间隔 (单位秒)",
			EnvVars:     []string{"GAME_AGENT_HEARTBEAT_INTERVAL"},
			Destination: &Opts.HeartbeatInterval,
		},
		&cli.Int64Flag{
			Name:        "agent_heartbeat_deadline",
			Value:       10,
			Usage:       "设置当前网关的心跳的等待时间 (单位秒)",
			EnvVars:     []string{"GAME_AGENT_HEARTBEAT_DEADLINE"},
			Destination: &Opts.HeartbeatDeadline,
		},
		&cli.UintFlag{
			Name:        "agent_conn_max",
			Value:       0,
			Usage:       "设置当前网关的最大连接数",
			EnvVars:     []string{"GAME_AGENT_CONN_MAX"},
			Destination: &Opts.ConnMaxNum,
		},
		&cli.Int64Flag{
			Name:        "agent_auth_time",
			Value:       5,
			Usage:       "设置当前网关的鉴权认证有效时间 (单位秒)",
			EnvVars:     []string{"GAME_AGENT_AUTH_TIME"},
			Destination: &Opts.AuthTime,
		},
	}
)
