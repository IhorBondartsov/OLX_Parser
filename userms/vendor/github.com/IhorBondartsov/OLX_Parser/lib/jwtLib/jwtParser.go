package jwtLib

import (
	"crypto/rsa"
	"encoding/asn1"
	"encoding/pem"

	"github.com/dgrijalva/jwt-go"
)

type JWTParser interface {
	Parse(token string) (*Claims, error)
}

func NewJWTParser(publicKey []byte) (JWTParser, error) {
	block, _ := pem.Decode(publicKey)
	var pk rsa.PublicKey
	_, err := asn1.Unmarshal(block.Bytes, &pk)
	if err != nil {
		return nil, err
	}

	return &rsaParser{
		publicKey: &pk,
	}, nil
}

type rsaParser struct {
	publicKey *rsa.PublicKey
}

func (p *rsaParser) Parse(token string) (*Claims, error) {
	var claim Claims
	t, err := jwt.ParseWithClaims(token, &claim, func(t *jwt.Token) (interface{}, error) {
		return p.publicKey, nil
	})
	if t.Valid {
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, ErrTokenMalformed
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, ErrTokenExpired
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
	claims, ok := t.Claims.(*Claims)
	if !ok && !t.Valid {
		return nil, err
	}
	return claims, nil
}
