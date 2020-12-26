package protocol

import (
	"github.com/golang/protobuf/proto"
	"reflect"
)

// 游戏协议结构
type Protocol struct {
	Cmd     uint32
	Client  reflect.Type
	Server  reflect.Type
	Comment string
}

func (r *Protocol) NewClient() proto.Message {
	if r.Client != nil {
		return reflect.New(r.Client).Interface().(proto.Message)
	}
	return nil
}

func (r *Protocol) NewServer() proto.Message {
	if r.Server != nil {
		return reflect.New(r.Server).Interface().(proto.Message)
	}
	return nil
}

func NewProtocol(cmd uint32, client, server interface{}, comment string) *Protocol {
	return &Protocol{
		Cmd:     cmd,
		Client:  reflect.TypeOf(client),
		Server:  reflect.TypeOf(server),
		Comment: comment,
	}
}
