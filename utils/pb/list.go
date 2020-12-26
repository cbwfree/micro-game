package pb

import _struct "github.com/golang/protobuf/ptypes/struct"

type ProtoListValue struct {
	*_struct.ListValue
}

func (p *ProtoListValue) Proto() *_struct.ListValue {
	return p.ListValue
}

func (p *ProtoListValue) Index(i int) *ProtoValue {
	if i >= 0 && i < len(p.ListValue.Values) {
		return NewProtoValue(p.ListValue.Values[i])
	}
	return NewProtoValue(new(_struct.Value))
}

func (p *ProtoListValue) Values() []*ProtoValue {
	var values []*ProtoValue
	for _, v := range p.ListValue.Values {
		values = append(values, NewProtoValue(v))
	}
	return values
}

func (p *ProtoListValue) Interface() interface{} {
	var values []interface{}
	for _, v := range p.ListValue.Values {
		values = append(values, NewProtoValue(v).Interface())
	}
	return values
}

func NewProtoListValue(v *_struct.ListValue) *ProtoListValue {
	return &ProtoListValue{v}
}
