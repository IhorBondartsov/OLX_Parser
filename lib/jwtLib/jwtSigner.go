package jwtLib

import "crypto/rsa"

type jwtSigner interface{
	Sign()
}

type RSASigner struct{
	privateKey *rsa.PrivateKey
}

func