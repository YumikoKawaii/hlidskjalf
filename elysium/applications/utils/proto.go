package utils

import "github.com/gogo/protobuf/types"

type SortOrder int32

const (
	ASC  SortOrder = 0
	DESC SortOrder = 1
)

func ToSortOrder(value int32) SortOrder {
	if value == int32(ASC) {
		return ASC
	}

	return DESC
}

func (s SortOrder) Value() string {
	if s == ASC {
		return "ASC"
	}

	return "DESC"
}

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
