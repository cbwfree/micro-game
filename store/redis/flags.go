package redis

import (
	"github.com/micro/cli/v2"
)

var (
	Opts = &struct {
		RedisUrl      string // Redis Proxy URI地址
		RedisDb       int    // Redis Db
		RedisIdeConns int    // Redis 最小空闲连接数
		RedisMaxPool  int    // Redis 最大连接数
	}{}

	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "redis",
			Value:       "",
			Usage:       "设置redis连接地址. 格式: [:password@]hostname:port/[db]",
			EnvVars:     []string{"GAME_REDIS_URL"},
			Destination: &Opts.RedisUrl,
		},
		&cli.IntFlag{
			Name:        "redis_db",
			Value:       0,
			Usage:       "设置redis数据库",
			EnvVars:     []string{"GAME_REDIS_DB"},
			Destination: &Opts.RedisDb,
		},
		&cli.IntFlag{
			Name:        "redis_ide_conns",
			Value:       40,
			Usage:       "设置redis最大空闲连接数",
			EnvVars:     []string{"GAME_REDIS_IDE_CONNS"},
			Destination: &Opts.RedisIdeConns,
		},
		&cli.IntFlag{
			Name:        "redis_max_pool",
			Value:       80,
			Usage:       "设置redis最大连接数",
			EnvVars:     []string{"GAME_REDIS_MAX_POOL"},
			Destination: &Opts.RedisMaxPool,
		},
	}
)
