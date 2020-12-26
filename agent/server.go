package agent

// 网关网络服务接口
type Server interface {
	Name() string   // 服务器名称
	Agent() *Agent  // 关联服务端
	Opts() *Options // 参数
	Port() int      // 监听端口
	Run() error     // 启动服务
	Close()         // 停止服务
}
