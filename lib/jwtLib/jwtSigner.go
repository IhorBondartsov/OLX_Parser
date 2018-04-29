package jwtLib

import (
	"crypto/rsa"

	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTSigner interface {
	Sign(c Claims, t time.Duration) (string, error)
}

func NewJWTSigner(key []byte) (JWTSigner, error) {
	k, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return nil, err
	}
	return &rsaSigner{
		privateKey: k,
	}, nil
}

type rsaSigner struct {
	privateKey *rsa.PrivateKey
}

func (s *rsaSigner) Sign(c Claims, t time.Duration) (string, error) {
	c.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(t).Unix(),
		Issuer:    "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	ss, err := token.SignedString(s.privateKey)
	return ss, err
}

type Claims struct {
	jwt.StandardClaims
	ID string `json:"foo"`
}
