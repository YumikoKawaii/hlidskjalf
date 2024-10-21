package utils

import "github.com/gogo/protobuf/types"

type SortOrder int

const (
	ASC  = 0
	DESC = 1
)

func UInt32PointerToProto(value *uint32) *types.UInt32Value {
	if value == nil {
		return nil
	}

	return &types.UInt32Value{
		Value: *value,
	}
}

func ProtoToUInt32Pointer(value *types.UInt32Value) *uint32 {
	if value == nil {
		return nil
	}

	return &value.Value
}

func StringPointerToProto(value *string) *types.StringValue {
	if value == nil {
		return nil
	}

	return &types.StringValue{
		Value: *value,
	}
}

func ProtoToStringPointer(value *types.StringValue) *string {
	if value == nil {
		return nil
	}

	return &value.Value
}
