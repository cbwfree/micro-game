package redis

import "time"

const (
	DefaultMaxRetries      = 3
	DefaultReadTimeout     = 3 * time.Second
	DefaultWriteTimeout    = 3 * time.Second
	DefaultIdleTimeout     = 5 * time.Minute
	DefaultMinRetryBackoff = 200 * time.Millisecond
	DefaultMaxRetryBackoff = 1 * time.Second
)

type Option func(o *Options)

type Options struct {
	RawUrl          string
	Db              int           // 数据库
	MinIdleConns    int           // 最小空闲连接数
	PoolSize        int           // 最大连接数
	MaxRetries      int           // 最大重试次数
	ReadTimeout     time.Duration // 读取超时
	WriteTimeout    time.Duration // 写入超时
	IdleTimeout     time.Duration // 空闲连接关闭时间
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration
}

func (o *Options) Init(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func newOptions(opts ...Option) *Options {
	o := &Options{
		RawUrl:          Opts.RedisUrl,
		Db:              Opts.RedisDb,
		MinIdleConns:    Opts.RedisIdeConns,
		PoolSize:        Opts.RedisMaxPool,
		MaxRetries:      DefaultMaxRetries,
		ReadTimeout:     DefaultReadTimeout,
		WriteTimeout:    DefaultWriteTimeout,
		IdleTimeout:     DefaultIdleTimeout,
		MinRetryBackoff: DefaultMinRetryBackoff,
		MaxRetryBackoff: DefaultMaxRetryBackoff,
	}
	o.Init(opts...)
	return o
}

func WithUrl(url string) Option {
	return func(o *Options) {
		o.RawUrl = url
	}
}

func WithDb(db int) Option {
	return func(o *Options) {
		o.Db = db
	}
}

func WithMinIdleConns(size int) Option {
	return func(o *Options) {
		o.MinIdleConns = size
	}
}

func WithPoolSize(size int) Option {
	return func(o *Options) {
		o.PoolSize = size
	}
}

func WithMaxRetries(size int) Option {
	return func(o *Options) {
		o.MaxRetries = size
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.WriteTimeout = t
	}
}

func WithIdleTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.IdleTimeout = t
	}
}

func WithMinRetryBackoff(t time.Duration) Option {
	return func(o *Options) {
		o.MinRetryBackoff = t
	}
}

func WithMaxRetryBackoff(t time.Duration) Option {
	return func(o *Options) {
		o.MaxRetryBackoff = t
	}
}
