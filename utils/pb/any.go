package pb

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func MarshalAny(m proto.Message) (*any.Any, error) {
	return ptypes.MarshalAny(m)
}

func UnmarshalAny(a *any.Any, m proto.Message) error {
	return ptypes.UnmarshalAny(a, m)
}
