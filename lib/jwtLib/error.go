package jwtLib

import "github.com/powerman/rpc-codec/jsonrpc2"

var (
	ErrTokenMalformed = jsonrpc2.NewError(1, "Token is malformed")
	ErrTokenExpired   = jsonrpc2.NewError(2, "Token expired")
)
