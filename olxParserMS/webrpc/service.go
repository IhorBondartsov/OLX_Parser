package webrpc

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/IhorBondartsov/OLX_Parser/lib/jwtLib"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Start(cfg CfgAPI) {
	// Server export an object of type ExampleSvc.
	if err := rpc.Register(NewAPI(cfg)); err != nil {
		log.Panic(err)
	}

	// Server provide a HTTP transport on /rpc endpoint.
	http.Handle("/rpc", jsonrpc2.HTTPHandler(nil))

}

func NewAPI(cfg CfgAPI) *API {
	atp, err := jwtLib.NewJWTParser(cfg.AccessPublicKey)
	if err != nil {
		log.Errorf("Cant create AccessTokenParser. Err %v", err)
		return nil
	}
	return &API{
		AccessTokenParser: atp,

	}
}

type CfgAPI struct {
	AccessPublicKey  []byte
}

type API struct {
	AccessTokenParser jwtLib.JWTParser

}

// Echo method for checking service
func (a *API) Echo(req EchoReq, res *EchoRes) error {
	fmt.Println("I called")
	res.Answer = fmt.Sprintf("Hello %s!!!", req.Name)
	return nil
}


func MakeOrder(req , res *)




