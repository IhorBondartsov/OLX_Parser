package clients

import (
	"github.com/IhorBondartsov/OLX_Parser/website/entity"
)



func NewUserRPCClient(host, port, prefix string) *rpcUserClient {
	return &rpcUserClient{
		RPCClient{
			Prefix: prefix,
			Host:   host,
			Port:   port,
		}}
}

type rpcUserClient struct {
	RPCClient
}

func (c *rpcUserClient) Echo(name string) ([]byte, error) {
	var er entity.EchoRes
	opts := callOpts{
		apiName:   "Echo",
		marshalTo: &er,
		reqBody:   entity.EchoReq{Name: name},
	}
	c.call(&opts)
	return opts.response, opts.err
}

func (c *rpcUserClient) Login(login, password string) ([]byte, error) {
	var er entity.LoginResp
	opts := callOpts{
		apiName:   "Login",
		marshalTo: &er,
		reqBody: entity.LoginReq{
			Login:    login,
			Password: password,
		},
	}
	c.call(&opts)
	return opts.response, opts.err
}

func (c *rpcUserClient) GetAcessToken(refreshToken string) ([]byte, error) {
	var er entity.AcessTokenResponse
	opts := callOpts{
		apiName:   "GetAcessToken",
		marshalTo: &er,
		reqBody: entity.AcessTokenRequest{
			RefreshToken: refreshToken,
		},
	}
	c.call(&opts)
	return opts.response, opts.err
}
