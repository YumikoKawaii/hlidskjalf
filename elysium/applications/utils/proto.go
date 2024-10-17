package utils

import "github.com/gogo/protobuf/types"

func UInt32PointerToProto(value *uint32) *types.UInt32Value {
	if value == nil {
		return nil
	}

	return &types.UInt32Value{
		Value: *value,
	}
}
