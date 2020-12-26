package web

import (
	"github.com/micro/cli/v2"
)

var (
	Opts = &struct {
		ReadTimeout  int64 // 读取超时
		WriteTimeout int64 // 写入超时
	}{}

	Flags = []cli.Flag{
		&cli.Int64Flag{
			Name:        "web_read_timeout",
			Usage:       "Web读取超时. 单位秒",
			EnvVars:     []string{"WEB_READ_TIMEOUT"},
			Value:       30,
			Destination: &Opts.ReadTimeout,
		},
		&cli.Int64Flag{
			Name:        "web_white_timeout",
			Usage:       "Web写入超时. 单位秒",
			EnvVars:     []string{"WEB_WHITE_TIMEOUT"},
			Value:       30,
			Destination: &Opts.WriteTimeout,
		},
	}
)
