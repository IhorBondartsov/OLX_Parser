package clients

import (
	"encoding/json"

	"github.com/IhorBondartsov/OLX_Parser/website/entity"
	"github.com/Sirupsen/logrus"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

var log = logrus.New()

func NewUserRPCClient(host, port, prefix string) *rpcUserClient {
	return &rpcUserClient{
		Prefix: prefix,
		Host:   host,
		Port:   port,
	}
}

type rpcUserClient struct {
	Prefix string
	Port   string
	Host   string
}

type callOpts struct {
	marshalTo interface{}
	reqBody   interface{}
	err       error
	response  []byte
	apiName   string
}

func (c *rpcUserClient) call(opts *callOpts) {
	client := jsonrpc2.NewHTTPClient(c.Host + ":" + c.Port + "/rpc")
	defer client.Close()

	err := client.Call(c.Prefix+"."+opts.apiName, opts.reqBody, &opts.marshalTo)
	if err != nil {
		log.Errorf("[rpcUserClient][%v]Cant call Err: %v", opts.apiName, err)
		opts.err = err
		return
	}

	body, err := json.Marshal(opts.marshalTo)
	if err != nil {
		log.Errorf("[rpcUserClient][Echo] Cant marshal Err: %v", err)
		opts.err = err
		return
	}

	opts.response = body
	return
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
			Login: login,
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
			RefreshToken:refreshToken,
		},
	}
	c.call(&opts)
	return opts.response, opts.err
}