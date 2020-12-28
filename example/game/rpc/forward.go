package server

import (
	"context"
	"fmt"
	"github.com/cbwfree/micro-game/agent"
	"github.com/cbwfree/micro-game/example/game/mod"
	pgame "github.com/cbwfree/micro-game/example/proto/game"
	"github.com/cbwfree/micro-game/utils/color"
	"github.com/cbwfree/micro-game/utils/debug"
	"github.com/cbwfree/micro-game/utils/errors"
	"github.com/cbwfree/micro-game/utils/log"
	"net/http"
)

type Forward struct{}

// 调用游戏协议
func (r *Forward) Protocol(ctx context.Context, req *pgame.InForwardProtocol, rsp *pgame.OutForwardProtocol) error {
	gmt, err := agent.FromMeta(ctx)
	if err != nil {
		return err
	}

	// 校验账号是否已登录
	if req.Cmd != 10001 && gmt.AccountId() == 0 {
		log.Warn("[Call] [%d] was invalid, because the account id is nil, forced offline ...", req.Cmd)
		return errors.New(http.StatusUnauthorized, "Game Login Unauthorized")
	}

	var cLog = log.Logger
	var prof *debug.Prof
	if log.IsTrace() {
		fields := map[string]interface{}{
			"cmd":    req.Cmd,
			"ip":     gmt.ClientIp(),
			"client": gmt.ClientId(),
		}
		if gmt.RoleId() != 0 {
			fields["role"] = gmt.RoleId()
		}
		cLog = log.Logger.WithFields(fields)
		prof = debug.NewProf(cLog, fmt.Sprintf("[%d]", req.Cmd))
	}

	cLog.Debugf("[Call] command [%d] ...", req.Cmd)

	// 请求协议路由
	var status string
	if s2c, err := mod.Router().Call(gmt, req.Cmd, req.Data); err != nil {
		ee := errors.Parse(err)
		status = ee.Status
		rsp.Code = uint32(ee.Code)
	} else {
		rsp.Data = s2c
	}

	if rsp.Code > 0 {
		cLog.Warn(color.Question.Text("[Call] command [%d] error: [%d] %s, %+v", req.Cmd, rsp.Code, status, err))
	}

	if prof != nil {
		prof.Result()
	}

	return nil
}

// 离线操作
func (r *Forward) Offline(ctx context.Context, _ *pgame.InForwardOffline, _ *pgame.OutForwardOffline) error {
	gmt, err := agent.FromMeta(ctx)
	if err != nil {
		return err
	}

	if gmt.RoleId() == 0 {
		return nil
	}

	// TODO 离线处理逻辑

	return nil
}
