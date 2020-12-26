package pb

import _struct "github.com/golang/protobuf/ptypes/struct"

type ProtoStruct struct {
	*_struct.Struct
}

func (p *ProtoStruct) Proto() *_struct.Struct {
	return p.Struct
}

func (p *ProtoStruct) Value(field string) *ProtoValue {
	if p.Fields != nil {
		if val, ok := p.Fields[field]; ok {
			return NewProtoValue(val)
		}
	}
	return NewProtoValue(new(_struct.Value))
}

func (p *ProtoStruct) Interface() interface{} {
	var res = make(map[string]interface{})
	for k, v := range p.Fields {
		res[k] = NewProtoValue(v).Interface()
	}
	return res
}

func NewProtoStruct(s *_struct.Struct) *ProtoStruct {
	return &ProtoStruct{s}
}
