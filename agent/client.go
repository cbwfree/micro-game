package agent

import (
	"github.com/cbwfree/micro-game/codec"
	"github.com/micro/go-micro/v2/logger"
)

// 客户端
type Client interface {
	Id() string                               // Client Id
	Server() Server                           // 关联服务端
	Meta() *Meta                              // 客户端上下文
	Log() *logger.Helper                      // 日志对象
	Closed() bool                             // 判断是否关闭
	Read() (*codec.ClientHead, []byte, error) // 读取消息
	Write([]byte)                             // 发送消息
	Close()                                   // 关闭连接
	Destroy()                                 // 销毁连接 (丢弃任何未发送或未确认的数据)
	SetAuthState(state bool)                  // 设置认证状态 (建立Socket连接后, 需要发送Token进行认证)
}
