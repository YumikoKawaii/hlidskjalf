package utils

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ProtoJSONMarshaler return the marshaler option with support serialization data with json_name specific
func ProtoJSONMarshaler() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true})
}

// StandardisedProtoJSONMarshaler returns a marshaler option which serialize GRPC response to HTTP response with the standardised format.
func StandardisedProtoJSONMarshaler() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true})
}
