package protocol

import (
	"fmt"
	"github.com/cbwfree/micro-game/agent"
	"github.com/cbwfree/micro-game/utils/errors"
	"reflect"
	"strconv"
	"strings"
)

type Route struct {
	cmd     uint32
	name    string
	hdlr    reflect.Value
	method  reflect.Method
	reqType reflect.Type
	rspType reflect.Type
}

func (h *Route) Cmd() uint32 {
	return h.cmd
}

func (h *Route) Name() string {
	return h.name
}

func (h *Route) NewReqValue() reflect.Value {
	return reflect.New(h.reqType)
}

func (h *Route) NewRspValue() reflect.Value {
	return reflect.New(h.rspType)
}

func (h *Route) Call(gmt *agent.Meta, req, rsp interface{}) error {
	values := h.method.Func.Call([]reflect.Value{
		h.hdlr,
		reflect.ValueOf(gmt),
		reflect.ValueOf(req),
		reflect.ValueOf(rsp),
	})
	if err := values[0].Interface(); err != nil {
		return err.(error)
	}
	return nil
}

func ParseRoutes(handler interface{}) []*Route {
	var routes []*Route

	typ := reflect.TypeOf(handler)
	hdlr := reflect.ValueOf(handler)
	name := reflect.Indirect(hdlr).Type().Name()

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)

		cmd, err := parseProtocolCmd(method.Name)
		if err != nil {
			panic(err.Error())
		}

		mtype := method.Type
		reqType := mtype.In(2)
		rspType := mtype.In(3)

		if reqType.Kind() == reflect.Ptr {
			reqType = reqType.Elem()
		}
		if rspType.Kind() == reflect.Ptr {
			rspType = rspType.Elem()
		}

		routes = append(routes, &Route{
			cmd:     cmd,
			name:    fmt.Sprintf("%s.%s", name, method.Name),
			hdlr:    hdlr,
			method:  method,
			reqType: reqType,
			rspType: rspType,
		})
	}

	return routes
}

func parseProtocolCmd(name string) (uint32, error) {
	names := strings.Split(name, "_")
	if len(names) != 2 {
		return 0, errors.Invalid("%s.%s format error", name, name)
	}

	cmd, err := strconv.ParseUint(names[1], 10, 32)
	if err != nil {
		return 0, errors.Invalid("%s.%s invalid command: %s", name, name, names[1])
	}

	return uint32(cmd), nil
}
