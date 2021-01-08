package meta

import (
	"context"
	"github.com/micro/go-micro/v2/metadata"
)

const (
	KeyDataId = "Data-Id"
)

// 数据节点
type DataNodeMeta interface {
	NodeId() string   // 节点ID
	NodeName() string // 节点名称
	DataId() string   // 数据ID
	Context() context.Context
}

// 数据节点
type DataMeta struct {
	*Meta
}

func (m *DataMeta) IsValid() bool {
	return m.Len() > 0 && m.Get(KeyDataId) != ""
}

func (m *DataMeta) DataId() string {
	return m.Get(KeyDataId)
}

func (m *DataMeta) SetDataId(id interface{}) {
	m.Set(KeyDataId, ToMetaValue(id))
}

// 实例化数据节点
func NewDataMeta(name, id string, dataId interface{}) *DataMeta {
	mt := &DataMeta{
		Meta: NewMeta(name, id, metadata.Metadata{
			KeyDataId: ToMetaValue(dataId),
		}),
	}
	return mt
}

func FromDataMeta(ctx context.Context) (*DataMeta, error) {
	mt, err := FromMeta(ctx)
	if err != nil {
		return nil, err
	}
	return &DataMeta{Meta: mt}, nil
}

func ToDataMeta(values map[string]string) *DataMeta {
	if values == nil {
		values = make(map[string]string)
	}
	return &DataMeta{Meta: ToMeta(values)}
}
