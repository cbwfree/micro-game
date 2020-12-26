package mongo

import "time"

const (
	DefaultConnectTimeout   = 5 * time.Second // 连接超时时间
	DefaultSocketTimeout    = 5 * time.Second
	DefaultMaxConnIdleTime  = 3 * time.Second // 最大空闲时间
	DefaultReadWriteTimeout = 3 * time.Second // 读写超时时间
)

type Option func(o *Options)

type Options struct {
	RawUrl           string
	DbName           string
	MinPoolSize      uint64        // 最小连接池大小
	MaxPoolSize      uint64        // 最大连接池大小
	ConnectTimeout   time.Duration // 连接超时时间
	SocketTimeout    time.Duration
	MaxConnIdleTime  time.Duration // 最大空闲时间
	ReadWriteTimeout time.Duration // 读写超时时间
}

func (o *Options) Init(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func newOptions(opts ...Option) *Options {
	o := &Options{
		RawUrl:           Opts.MongoUrl,
		MinPoolSize:      Opts.MongoMinPool,
		MaxPoolSize:      Opts.MongoMaxPool,
		ConnectTimeout:   DefaultConnectTimeout,
		SocketTimeout:    DefaultSocketTimeout,
		MaxConnIdleTime:  DefaultMaxConnIdleTime,
		ReadWriteTimeout: DefaultReadWriteTimeout,
	}
	o.Init(opts...)
	return o
}

func WithUrl(url string) Option {
	return func(o *Options) {
		o.RawUrl = url
	}
}

func WithDbName(db string) Option {
	return func(o *Options) {
		o.DbName = db
	}
}

func WithMinPoolSize(size uint64) Option {
	return func(o *Options) {
		o.MinPoolSize = size
	}
}

func WithMaxPoolSize(size uint64) Option {
	return func(o *Options) {
		o.MaxPoolSize = size
	}
}

func WithConnectTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.ConnectTimeout = t
	}
}

func WithSocketTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.SocketTimeout = t
	}
}

func WithMaxConnIdleTime(t time.Duration) Option {
	return func(o *Options) {
		o.MaxConnIdleTime = t
	}
}

func WithReadWriteTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.ReadWriteTimeout = t
	}
}
