package jwtLib

import "crypto/rsa"

type JWTParser interface {
	Parse()
}
type RSAParser struct{
	publicKey *rsa.PublicKey
}