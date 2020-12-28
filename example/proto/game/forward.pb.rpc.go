// Code generated by protoc-gen-rpc. DO NOT EDIT.
// source: proto/game/forward.proto

package game

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import "context"
import "github.com/cbwfree/micro-game/app"
import "github.com/cbwfree/micro-game/meta"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//Forward RPC Method constant definition
const (
	ForwardMethod_Protocol = "Forward.Protocol"
	ForwardMethod_Offline  = "Forward.Offline"
)

//Forward RPC Service
type ForwardService interface {
	Protocol(ctx context.Context, req *InForwardProtocol, rsp *OutForwardProtocol) error
	Offline(ctx context.Context, req *InForwardOffline, rsp *OutForwardOffline) error
}

//Forward RPC Client
type ForwardClient struct {
	name string
}

//CallNode General call method
func (c *ForwardClient) CallNode(ctx context.Context, method string, in interface{}, out interface{}, nodeId ...string) error {
	return app.CallNode(ctx, c.name, method, in, out, nodeId...)
}

//Protocol
func (c *ForwardClient) Protocol(mt meta.NodeMeta, in *InForwardProtocol) (*OutForwardProtocol, error) {
	out := new(OutForwardProtocol)
	err := app.CallNode(mt.Context(), c.name, ForwardMethod_Protocol, in, out, mt.NodeId())
	return out, err
}

//Offline
func (c *ForwardClient) Offline(mt meta.NodeMeta, in *InForwardOffline) (*OutForwardOffline, error) {
	out := new(OutForwardOffline)
	err := app.CallNode(mt.Context(), c.name, ForwardMethod_Offline, in, out, mt.NodeId())
	return out, err
}

//NewForwardClient
func NewForwardClient(name string) *ForwardClient {
	return &ForwardClient{
		name: name,
	}
}