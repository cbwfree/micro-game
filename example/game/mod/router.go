// 协议路由
package mod

import (
	"github.com/cbwfree/micro-game/protocol"
	"sync"
)

var (
	routerInstance *protocol.Router
	routerOnce     sync.Once
)

func Router() *protocol.Router {
	routerOnce.Do(func() {
		routerInstance = protocol.NewRouter()
	})
	return routerInstance
}
