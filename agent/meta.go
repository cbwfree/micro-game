package agent

import (
	"context"
	"github.com/cbwfree/micro-game/app"
	"github.com/cbwfree/micro-game/meta"
	"github.com/cbwfree/micro-game/utils/dtype"
)

const (
	MetaClientId   = "Client-Id"   // 客户端连接ID
	MetaClientIp   = "Client-Ip"   // IPv4地址
	MetaClientVer  = "Client-Ver"  // 客户端版本
	MetaAccountId  = "Account-Id"  // 账户ID
	MetaChannelUid = "Channel-Uid" // 渠道账户
	MetaServerId   = "Server-Id"   // 服务器ID
	MetaRoleId     = "Role-Id"     // 角色ID
)

// 网关上下文
type Meta struct {
	*meta.Meta
}

func (ctx *Meta) ClientId() string {
	return ctx.Get(MetaClientId)
}

func (ctx *Meta) ClientVer() string {
	return ctx.Get(MetaClientVer)
}

func (ctx *Meta) ClientIp() string {
	return ctx.Get(MetaClientIp)
}

func (ctx *Meta) AccountId() int64 {
	return dtype.ParseInt64(ctx.Get(MetaAccountId))
}

func (ctx *Meta) ChannelUid() string {
	return ctx.Get(MetaChannelUid)
}

func (ctx *Meta) ServerId() int32 {
	return dtype.ParseInt32(ctx.Get(MetaServerId))
}

func (ctx *Meta) RoleId() int64 {
	return dtype.ParseInt64(ctx.Get(MetaRoleId))
}

func (ctx *Meta) IsOnline() bool {
	return ctx.Get(MetaRoleId) != ""
}

// NewMeta 实例化metadata数据
func NewMeta(clientId string) *Meta {
	mt := meta.NewMeta(app.Name(), app.Id(), map[string]string{
		MetaClientId:   clientId,
		MetaClientIp:   "",
		MetaClientVer:  "",
		MetaAccountId:  "",
		MetaChannelUid: "",
		MetaServerId:   "",
		MetaRoleId:     "",
	})
	return &Meta{Meta: mt}
}

// FromMeta 从Meta解析metadata数据
func FromMeta(ctx context.Context) (*Meta, error) {
	mt, err := meta.FromMeta(ctx)
	if err != nil {
		return nil, err
	}
	return &Meta{Meta: mt}, nil
}
