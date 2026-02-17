package constants

import "time"

const (
	HTTPScheme = "http"
)

const (
	Identity = "skidbladnir"
	Limiter  = "limiter"
)

const (
	FollowerTTL      = 2 * time.Second
	FollowerInterval = 500 * time.Millisecond
)
