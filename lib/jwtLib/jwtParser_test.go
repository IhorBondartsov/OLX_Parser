package jwtLib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
MIIEpQIBAAKCAQEAwEXBRCwisurukRcgKDfTpEHlG0lZOjNgPiS3vDorVv5k8pk6
iERM0Q5Bi9ok9RLEuIuxY10b5ODp5qtIXODhg3a/hNye1gaQ1a2JhixTC0DUxYL0
GsaGlUdGd6I3jYxrSjUGFGCubbcllBFnu4BsLxLcy/3sm/ym5sL3aYgjbjB8j/R5
T+RJKn/06FdhhxbjVrOQ+ySCvTzAizF+n7Iu/iiVW+0LrWru5GqnjkDp4h3iF9PQ
EoeaCLFP+XhMEsNF1cuWpZo4JcZODPyP9uhNOmzXR6C5Fd9nsTfrLm1bggMqZvZT
vctOOiP8d2rkiLV0iPNV8KID/kWiGAWcwJ4bJQIDAQABAoIBAQCVED9i0/jez2a/
k7c5lvZ9HR07R2VmytfttdfAlTRukHHA52zKonEPjsbgXvJSEgfzr1GKYHA0xO5y
CA7k+u+7VT/sCIMYGSUGlDDPleOYEN4kdn87lvhWGVkXfOsm0VuIv40EFWSF6Py4
S7opUsoMwMnvWOCsmnbm5vabmZEmZhVY11s68Bj4gpkSPbg2PX1/9TDi3t81XOpX
Hs0SUMuLqu+RBv3zcKKE3mzpAEJIt4a+M73dfIhAbcMlCIk6qeBS4Ef9d37n2sGP
tUf3uZEr2J2ubjKcsqARZOh7ik6H6/FrKLs9dayMA11YJ1JQeRo6gsZmtv+1K79b
r9HgZUfBAoGBAMoPMUZoRGpR8I8ZUWQeff3oX/ft5DnbNe3O0zGqaNPnMBPkGW8l
0IJcNl3GE/6IYSOfKDfbZj5WeX5qSwd6BWyZGzXu1SEoF6Mt094zzxD/Vdwp/er6
7zOlovaMOYYWVP0Phg2mdozZQw8edMDkgwOs3z44/OL1yiJfDHRC9d41AoGBAPOZ
uX/Z1mrh35jlBb7pWNK9sbenuNMoX41hiLLM4O3VQvKz99VKf6ecVpnybqzRb/XZ
w6xSoNkdOFmLNUcesvXjGMVLFcjAb681Hwg/XrOHoV7AXUD1NX+oVs59dTpLeiIK
LDomI7xqx5jPL52mL2TimnGv3l/EMEgn8eE4cqcxAoGBALhUB7WdJNyBB5zjkdZt
4q4gsHuyr0MAcUASL6PdO2gkPT3IFlPJAWAK0yXUryTCMSuX486g/bLcnFTwjqWx
cQVS3vnieNeYhYe9JMxOIxX0eNUyEYutg01wSyvzW/6wnLOG19nYA8oV4REHzaay
T7qol7dhsYEyjEWyw9/DvhwJAoGBAMYOYkeUsURNSSJicn1584HEQU30y8zCso2c
kyvsb6RvE/OIZyO/YnLAJZkdBubzkXfuCa19sNFVdI85I6QjTZWlgmpGVfvWmfd1
1Tk7bJ/C0mzbMfayZ0DXeVyBnGI9I5uxcawFfmCe4JFZQnmGuFnp83AjcUWjmS3R
bFOVNYLBAoGAYFFMVyVmyhG1Bi2om/oughmIPwYW7N8Au3xfeegSbHXgc0fnAJAT
Wf2WkxXAOWRpK6dK4u/Xe00Ja28ze8GXtFlSn0hURMzaKjlq4eEof8/8bad3qLqe
0X+atKV2A3CJCoCMgscNo7UZ1zYkZxxbWEOLYiB3GPQYplouK0tN/4w=
-----END PRIVATE KEY-----
`)
var testPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBCgKCAQEAwEXBRCwisurukRcgKDfTpEHlG0lZOjNgPiS3vDorVv5k8pk6iERM
0Q5Bi9ok9RLEuIuxY10b5ODp5qtIXODhg3a/hNye1gaQ1a2JhixTC0DUxYL0GsaG
lUdGd6I3jYxrSjUGFGCubbcllBFnu4BsLxLcy/3sm/ym5sL3aYgjbjB8j/R5T+RJ
Kn/06FdhhxbjVrOQ+ySCvTzAizF+n7Iu/iiVW+0LrWru5GqnjkDp4h3iF9PQEoea
CLFP+XhMEsNF1cuWpZo4JcZODPyP9uhNOmzXR6C5Fd9nsTfrLm1bggMqZvZTvctO
OiP8d2rkiLV0iPNV8KID/kWiGAWcwJ4bJQIDAQAB
-----END PUBLIC KEY-----
`)


func TestRsaParser_Parse(t *testing.T) {
	claim := Claims{
		ID: "10",
	}
	a := assert.New(t)

	si, err:= NewJWTSigner(testPrivateKey)
	a.NoError(err)

	token, err := si.Sign(claim, 42)
	a.NoError(err)

	pa, err := NewJWTParser(testPublicKey)
	a.NoError(err)

	cl, err := pa.Parse(token)
	a.NoError(err)
	a.Equal(claim.ID, cl.ID)
}


func TestRsaParser_ParseError(t *testing.T) {
	// Token from another example.  This token is expired
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
	a := assert.New(t)

	pa, err := NewJWTParser(testPublicKey)
	a.NoError(err)

	cl, err := pa.Parse(tokenString)
	a.Error(err)
	a.Equal(ErrTokenExpired, err)
	a.Nil(cl)
}



