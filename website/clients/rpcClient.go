package clients

import (
	"encoding/json"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

type RPCClient struct{
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


func (c *RPCClient) call(opts *callOpts) {
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
