package pb

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
)

type ProtoValue struct {
	*_struct.Value
}

func (s *ProtoValue) Proto() *_struct.Value {
	return s.Value
}

func (s *ProtoValue) String() string {
	return s.GetStringValue()
}

func (s *ProtoValue) Bool() bool {
	return s.GetBoolValue()
}

func (s *ProtoValue) Number() float64 {
	return s.GetNumberValue()
}

func (s *ProtoValue) Null() _struct.NullValue {
	return s.GetNullValue()
}

func (s *ProtoValue) Struct() *ProtoStruct {
	return NewProtoStruct(s.GetStructValue())
}

func (s *ProtoValue) List() *ProtoListValue {
	return NewProtoListValue(s.GetListValue())
}

func (s *ProtoValue) Int() int {
	return int(s.GetNumberValue())
}

func (s *ProtoValue) Int32() int32 {
	return int32(s.GetNumberValue())
}

func (s *ProtoValue) Int64() int64 {
	return int64(s.GetNumberValue())
}

func (s *ProtoValue) Float32() float32 {
	return float32(s.GetNumberValue())
}

func (s *ProtoValue) Float64() float64 {
	return s.Number()
}

func (s *ProtoValue) Interface() interface{} {
	switch s.Value.GetKind().(type) {
	case *_struct.Value_BoolValue:
		return s.Bool()
	case *_struct.Value_NumberValue:
		return s.Number()
	case *_struct.Value_StringValue:
		return s.String()
	case *_struct.Value_ListValue:
		return s.List().Interface()
	case *_struct.Value_StructValue:
		return s.Struct().Interface()
	}
	return nil
}

func NewProtoValue(v *_struct.Value) *ProtoValue {
	return &ProtoValue{v}
}
