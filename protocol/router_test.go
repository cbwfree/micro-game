package protocol

import (
	"context"
	"fmt"
	"github.com/cbwfree/micro-game/agent"
	"github.com/golang/protobuf/proto"
	"testing"
)

type testHandle struct{}

func (t *testHandle) Test_10001(ctx context.Context, c2s proto.Message, s2c proto.Message) error {
	fmt.Printf("call Test_10001, c2s: %+v, s2c: %+v\n", c2s, s2c)
	return nil
}

func TestRouter_Call(t *testing.T) {
	r := NewRouter()
	r.AddRoute(new(testHandle))

	if s2c, err := r.Call(agent.NewMeta(""), 10001, nil); err != nil {
		fmt.Printf("Err: %s\n", err.Error())
	} else {
		fmt.Printf("Res: %+v\n", s2c)
	}
}

func TestRouter_Handler(t *testing.T) {
	r := NewRouter()
	r.AddRoute(new(testHandle))

	fmt.Printf("Routes: %+v\n", r.Routes())
}
