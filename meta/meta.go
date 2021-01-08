package meta

import (
	"context"
	"github.com/cbwfree/micro-game/app"
	"github.com/cbwfree/micro-game/utils/errors"
	"github.com/micro/go-micro/v2/metadata"
	"strconv"
	"sync"
)

const (
	KeyNodeId   = "Node-Id"
	KeyNodeName = "Node-Name"
)

// 普通节点
type NodeMeta interface {
	NodeId() string   // 节点ID
	NodeName() string // 节点名称
	Context() context.Context
}

// 网关上下文
type Meta struct {
	sync.RWMutex
	data metadata.Metadata
}

func (m *Meta) Len() int {
	m.RLock()
	defer m.RUnlock()

	return len(m.data)
}

func (m *Meta) Id() string {
	return m.Get(KeyNodeId)
}

func (m *Meta) SetId(id string) {
	m.Set(KeyNodeId, id)
}

func (m *Meta) Name() string {
	return m.Get(KeyNodeName)
}

func (m *Meta) SetName(name string) {
	m.Set(KeyNodeName, name)
}

func (m *Meta) Metadata() metadata.Metadata {
	m.RLock()
	defer m.RUnlock()

	var mtd = make(metadata.Metadata, len(m.data))
	for k, v := range m.data {
		mtd[k] = v
	}

	return mtd
}

func (m *Meta) SetMetadata(values metadata.Metadata) {
	for key, val := range values {
		m.data[key] = val
	}
}

func (m *Meta) Context() context.Context {
	return metadata.NewContext(context.TODO(), m.Metadata())
}

func (m *Meta) Get(key string) string {
	m.RLock()
	defer m.RUnlock()

	return m.data[key]
}

func (m *Meta) Set(key string, val interface{}) {
	m.Lock()
	defer m.Unlock()

	m.data[key] = ToMetaValue(val)
}

func (m *Meta) SetValues(values map[string]interface{}) {
	m.Lock()
	defer m.Unlock()

	for key, val := range values {
		m.data[key] = ToMetaValue(val)
	}
}

func (m *Meta) Bool(key string) bool {
	v, _ := strconv.ParseBool(m.Get(key))
	return v
}

func (m *Meta) Int(key string) int {
	v, _ := strconv.Atoi(m.Get(key))
	return v
}

func (m *Meta) Int32(key string) int32 {
	v, _ := strconv.Atoi(m.Get(key))
	return int32(v)
}

func (m *Meta) Int64(key string) int64 {
	v, _ := strconv.ParseInt(m.Get(key), 10, 64)
	return v
}

func (m *Meta) Uint32(key string) uint32 {
	v, _ := strconv.ParseUint(m.Get(key), 10, 32)
	return uint32(v)
}

func (m *Meta) Uint64(key string) uint64 {
	v, _ := strconv.ParseUint(m.Get(key), 10, 64)
	return v
}

func (m *Meta) Float32(key string) float32 {
	v, _ := strconv.ParseFloat(m.Get(key), 32)
	return float32(v)
}

func (m *Meta) Float64(key string) float64 {
	v, _ := strconv.ParseFloat(m.Get(key), 64)
	return v
}

func NewMeta(name, id string, meta metadata.Metadata) *Meta {
	mt := &Meta{
		data: metadata.Metadata{
			KeyNodeName: name,
			KeyNodeId:   id,
		},
	}
	for k, v := range meta {
		mt.data[k] = v
	}
	return mt
}

func ToMeta(meta metadata.Metadata) *Meta {
	return &Meta{data: meta}
}

func FromMeta(ctx context.Context) (*Meta, error) {
	mt, b := metadata.FromContext(ctx)
	if !b {
		return nil, errors.Invalid("no meta context")
	}
	return ToMeta(mt), nil
}

func NewContext(values map[string]string) context.Context {
	if len(values) == 0 {
		return context.TODO()
	}
	return ToMeta(values).Context()
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
