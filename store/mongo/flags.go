package mongo

import (
	"github.com/micro/cli/v2"
)

var (
	Opts = &struct {
		MongoUrl     string // MongoDB Proxy URI地址
		MongoMinPool uint64 // MongoDB 最小连接数
		MongoMaxPool uint64 // MongoDB 最大连接数
	}{}

	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "mongo",
			Value:       "",
			Usage:       "设置MongoDB连接地址. 格式: [username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]]",
			EnvVars:     []string{"GAME_MONGO_URL"},
			Destination: &Opts.MongoUrl,
		},
		&cli.Uint64Flag{
			Name:        "mongo_min_pool",
			Value:       20,
			Usage:       "设置MongoDB最小连接数",
			EnvVars:     []string{"GAME_MONGO_MIN_POOL"},
			Destination: &Opts.MongoMinPool,
		},
		&cli.Uint64Flag{
			Name:        "mongo_max_pool",
			Value:       100,
			Usage:       "设置MongoDB最大连接数",
			EnvVars:     []string{"GAME_MONGO_MAX_POOL"},
			Destination: &Opts.MongoMaxPool,
		},
	}
)
