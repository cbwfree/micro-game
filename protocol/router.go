package protocol

import (
	"github.com/cbwfree/micro-game/agent"
	"github.com/cbwfree/micro-game/utils/errors"
	"github.com/cbwfree/micro-game/utils/pb"
	"github.com/golang/protobuf/proto"
	"net/http"
)

// 游戏协议路由
type Router struct {
	protocols map[uint32]*Protocol
	routes    map[uint32]*Handler
}

func (r *Router) AddCmd(cmd uint32, client, server interface{}, comment string) {
	c := NewProtocol(cmd, client, server, comment)
	r.protocols[c.Cmd] = c
}

func (r *Router) Cmd(cmd uint32) *Protocol {
	if c, ok := r.protocols[cmd]; ok {
		return c
	}
	return nil
}

func (r *Router) NewClientProto(cmd uint32) proto.Message {
	if c, ok := r.protocols[cmd]; ok {
		return c.NewClient()
	}
	return nil
}

func (r *Router) NewServerProto(cmd uint32) proto.Message {
	if c, ok := r.protocols[cmd]; ok {
		return c.NewServer()
	}
	return nil
}

func (r *Router) ParseClientProto(cmd uint32, req []byte) (proto.Message, error) {
	if c, ok := r.protocols[cmd]; ok {
		c2s := c.NewClient()
		if err := proto.Unmarshal(req, c2s); err != nil {
			return nil, err
		}
		return c2s, nil
	}
	return nil, errors.New(http.StatusNotFound, "not found protocol %d", cmd)
}

// 调用
func (r *Router) Call(gmt *agent.Meta, cmd uint32, req []byte) (rsp []byte, err error) {
	route, ok := r.routes[cmd]
	if !ok {
		return nil, errors.New(http.StatusNotFound, "not found protocol %d", cmd)
	}

	c2s, err := r.ParseClientProto(cmd, req)
	if err != nil {
		return nil, err
	}

	s2c := r.NewServerProto(cmd)
	if s2c == nil {
		s2c = new(pb.None)
	}

	if err := route.Call(gmt, c2s, s2c); err != nil {
		return nil, err
	}

	return proto.Marshal(s2c)
}

// 注册
func (r *Router) Handler(handles ...interface{}) {
	for _, handler := range handles {
		hs := NewHandlers(handler)
		for _, h := range hs {
			r.routes[h.cmd] = h
		}
	}
}

// 注册
func (r *Router) Routes() map[uint32]*Handler {
	routes := make(map[uint32]*Handler, len(r.routes))
	for _, route := range r.routes {
		routes[route.cmd] = route
	}
	return routes
}

func NewRouter() *Router {
	return &Router{
		protocols: make(map[uint32]*Protocol),
		routes:    make(map[uint32]*Handler),
	}
}
