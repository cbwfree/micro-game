// Code generated by protoc-gen-rpc. DO NOT EDIT.
// source: proto/gate/gate.proto

package gate

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

//GateService RPC Method constant definition
const (
	GateServiceMethod_Push = "GateService.Push"
)

//GateService RPC Service
type GateServiceService interface {
	Push(ctx context.Context, req *InGatePush, rsp *OutGatePush) error
}

//GateService RPC Client
type GateServiceClient struct {
	name string
}

//CallNode General call method
func (c *GateServiceClient) CallNode(ctx context.Context, method string, in interface{}, out interface{}, nodeId ...string) error {
	return app.CallNode(ctx, c.name, method, in, out, nodeId...)
}

//Push
func (c *GateServiceClient) Push(mt meta.NodeMeta, in *InGatePush) (*OutGatePush, error) {
	out := new(OutGatePush)
	err := app.CallNode(mt.Context(), c.name, GateServiceMethod_Push, in, out, mt.NodeId())
	return out, err
}

//NewGateServiceClient
func NewGateServiceClient(name string) *GateServiceClient {
	return &GateServiceClient{
		name: name,
	}
}