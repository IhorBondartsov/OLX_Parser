package clients

import (
	"github.com/IhorBondartsov/OLX_Parser/website/entity"

)

func NewOLXParserRPCClient(host, port, prefix string) *olxParserRPCClient {
	return &olxParserRPCClient{
		RPCClient{
			Prefix: prefix,
			Host:   host,
			Port:   port,
		}}
}

type olxParserRPCClient struct {
	RPCClient
}

func (c *olxParserRPCClient) Echo(name string) ([]byte, error) {
	var er entity.EchoRes
	opts := callOpts{
		apiName:   "Echo",
		marshalTo: &er,
		reqBody:   entity.EchoReq{Name: name},
	}
	c.call(&opts)
	return opts.response, opts.err
}

func (c *olxParserRPCClient) MakeOrder(req entity.MakeOrderReq) ([]byte, error) {
	var er entity.MakeOrderRes
	opts := callOpts{
		apiName:   "MakeOrder",
		marshalTo: &er,
		reqBody:  req,
	}
	c.call(&opts)
	return opts.response, opts.err
}

func (c *olxParserRPCClient) ShowAllOder(req  entity.ShowAllOderReq) ([]byte, error) {
	var er entity.ShowAllOderResp
	opts := callOpts{
		apiName:   "ShowAllOder",
		marshalTo: &er,
		reqBody:   req,
	}
	c.call(&opts)
	return opts.response, opts.err
}

func (c *olxParserRPCClient) GetAdvertisementByOrder(req entity.GetAdvertisementByOrderReq) ([]byte, error) {
	var er entity.GetAdvertisementByOrderResp
	opts := callOpts{
		apiName:   "GetAdvertisementByOrder",
		marshalTo: &er,
		reqBody:   req,
	}
	c.call(&opts)
	return opts.response, opts.err
}