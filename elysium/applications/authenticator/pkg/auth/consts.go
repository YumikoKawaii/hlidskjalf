package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInternal = status.Errorf(codes.Internal, "internal error")
var ErrNotExistAccount = status.Errorf(codes.NotFound, "account not found")
var ErrWrongAccountInfo = status.Errorf(codes.InvalidArgument, "wrong account info")
var ErrExistAccount = status.Errorf(codes.AlreadyExists, "account exists")
