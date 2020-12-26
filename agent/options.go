package agent

import (
	"time"
)

var (
	DefaultReadTimeout  = 15 * time.Second // 默认请求超时时间
	DefaultWriteTimeout = 15 * time.Second // 默认请求超时时间
)

type Option func(o *Options)

// 服务器参数结果提
type Options struct {
	Address           string        // 服务器监听地址
	MaxConnNum        uint          // 最大连接数限制
	WaitAuthTime      time.Duration // 等待认证时间
	HeartbeatInterval time.Duration // 心跳间隔
	HeartbeatDeadline time.Duration // 心跳等待
	ReadTimeout       time.Duration // 读超时
	WriteTimeout      time.Duration // 写超时
}

func (o *Options) Init(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithAddress(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

func WithMaxConnNum(num uint) Option {
	return func(o *Options) {
		o.MaxConnNum = num
	}
}

func WithWaitAuthTime(t time.Duration) Option {
	return func(o *Options) {
		o.WaitAuthTime = t
	}
}

func WithHeartbeatInterval(t time.Duration) Option {
	return func(o *Options) {
		o.HeartbeatInterval = t
	}
}

func WithHeartbeatDeadline(t time.Duration) Option {
	return func(o *Options) {
		o.HeartbeatDeadline = t
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

func NewOptions(opts ...Option) *Options {
	o := &Options{
		Address:           Opts.Port,
		MaxConnNum:        Opts.ConnMaxNum,
		WaitAuthTime:      time.Duration(Opts.AuthTime) * time.Second,
		HeartbeatInterval: time.Duration(Opts.HeartbeatInterval) * time.Second,
		HeartbeatDeadline: time.Duration(Opts.HeartbeatDeadline) * time.Second,
		ReadTimeout:       DefaultReadTimeout,
		WriteTimeout:      DefaultWriteTimeout,
	}
	o.Init(opts...)
	return o
}
