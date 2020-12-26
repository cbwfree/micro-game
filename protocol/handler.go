package protocol

import (
	"fmt"
	"github.com/cbwfree/micro-game/agent"
	"github.com/golang/protobuf/proto"
	"reflect"
	"strconv"
	"strings"
)

type Handler struct {
	cmd    uint32
	name   string
	hdlr   reflect.Value
	method reflect.Method
}

func (h *Handler) Cmd() uint32 {
	return h.cmd
}

func (h *Handler) Name() string {
	return h.name
}

func (h *Handler) Call(gmt *agent.Meta, c2s proto.Message, s2c proto.Message) (err error) {
	values := []reflect.Value{
		h.hdlr,
		reflect.ValueOf(gmt),
		reflect.ValueOf(c2s),
		reflect.ValueOf(s2c),
	}

	returnValues := h.method.Func.Call(values)
	if err := returnValues[0].Interface(); err != nil {
		return err.(error)
	}

	return nil
}

func NewHandlers(handler interface{}) []*Handler {
	var handlers []*Handler

	typ := reflect.TypeOf(handler)
	hdlr := reflect.ValueOf(handler)
	name := reflect.Indirect(hdlr).Type().Name()

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		names := strings.Split(method.Name, "_")
		if len(names) != 2 {
			continue
		}
		cmd, _ := strconv.ParseUint(names[1], 10, 32)
		handlers = append(handlers, &Handler{
			cmd:    uint32(cmd),
			name:   fmt.Sprintf("%s.%s", name, method.Name),
			hdlr:   hdlr,
			method: method,
		})
	}

	return handlers
}
