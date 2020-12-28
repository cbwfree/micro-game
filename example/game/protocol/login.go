package protocol

import (
	"github.com/cbwfree/micro-game/agent"
	"github.com/cbwfree/micro-game/example/proto/msg"
	"github.com/golang/protobuf/ptypes/empty"
)

type Login struct{}

func (*Login) LoginServer_10001(gmt *agent.Meta, c2s *msg.C2S_10001, s2c *msg.S2C_10001) error {

	return nil
}

func (*Login) SelectRole_10002(gmt *agent.Meta, c2s *msg.C2S_10002, s2c *msg.S2C_10002) error {

	return nil
}

func (*Login) CreateRole_10003(gmt *agent.Meta, c2s *msg.C2S_10003, s2c *msg.S2C_10003) error {

	return nil
}

func (*Login) NoReturn_10004(gmt *agent.Meta, c2s *msg.C2S_10004, _ *empty.Empty) error {

	return nil
}
