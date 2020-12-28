package meta

import (
	"context"
	"github.com/cbwfree/micro-game/app"
	"github.com/cbwfree/micro-game/utils/errors"
	"github.com/micro/go-micro/v2/metadata"
	"strconv"
)

const (
	KeyNodeId   = "Node-Id"
	KeyNodeName = "Node-Name"
)

// 普通节点
type NodeMeta interface {
	NodeId() string
	NodeName() string
	Context() context.Context
}

// 网关上下文
type Meta metadata.Metadata

func (m Meta) NodeId() string {
	return m[KeyNodeId]
}

func (m Meta) SetNodeId(id string) {
	m[KeyNodeId] = id
}

func (m Meta) NodeName() string {
	return m[KeyNodeName]
}

func (m Meta) SetNodeName(name string) {
	m[KeyNodeName] = name
}

func (m Meta) Metadata() metadata.Metadata {
	return metadata.Metadata(m)
}

func (m Meta) SetMetadata(values metadata.Metadata) {
	for key, val := range values {
		m[key] = val
	}
}

func (m Meta) Context() context.Context {
	return metadata.NewContext(context.TODO(), m.Metadata())
}

func (m Meta) Get(key string) string {
	return m[key]
}

func (m Meta) Set(key string, val interface{}) {
	m[key] = ToMetaValue(val)
}

func (m Meta) SetValues(values map[string]interface{}) {
	for key, val := range values {
		m[key] = ToMetaValue(val)
	}
}

func (m Meta) Bool(key string) bool {
	v, _ := strconv.ParseBool(m.Get(key))
	return v
}

func (m Meta) Int(key string) int {
	v, _ := strconv.Atoi(m.Get(key))
	return v
}

func (m Meta) Int32(key string) int32 {
	v, _ := strconv.Atoi(m.Get(key))
	return int32(v)
}

func (m Meta) Int64(key string) int64 {
	v, _ := strconv.ParseInt(m.Get(key), 10, 64)
	return v
}

func (m Meta) Uint32(key string) uint32 {
	v, _ := strconv.ParseUint(m.Get(key), 10, 32)
	return uint32(v)
}

func (m Meta) Uint64(key string) uint64 {
	v, _ := strconv.ParseUint(m.Get(key), 10, 64)
	return v
}

func (m Meta) Float32(key string) float32 {
	v, _ := strconv.ParseFloat(m.Get(key), 32)
	return float32(v)
}

func (m Meta) Float64(key string) float64 {
	v, _ := strconv.ParseFloat(m.Get(key), 64)
	return v
}

func NewMeta(name, id string, meta metadata.Metadata) Meta {
	mt := Meta{
		KeyNodeName: name,
		KeyNodeId:   id,
	}
	for k, v := range meta {
		mt[k] = v
	}
	return mt
}

func FromMeta(ctx context.Context) (Meta, error) {
	mt, b := metadata.FromContext(ctx)
	if !b {
		return nil, errors.Invalid("no meta context")
	}
	return Meta(mt), nil
}

func NewContext(values map[string]string) context.Context {
	if len(values) == 0 {
		return context.TODO()
	}
	return Meta(values).Context()
}

// IsSelf 检查是否自身节点
func IsSelf(mt NodeMeta) bool {
	return mt.NodeName() == app.Name() && mt.NodeId() == app.Id()
}

// IsValid 检查节点是否有效
func IsValid(mt NodeMeta) bool {
	if mt.NodeId() == "" {
		return false
	}
	return app.CheckServiceNode(mt.NodeName(), mt.NodeId())
}
