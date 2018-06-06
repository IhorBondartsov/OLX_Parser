package webrpc

import (
	"fmt"
	"net/http"
	"net/rpc"

	"github.com/IhorBondartsov/OLX_Parser/lib/jwtLib"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/entities"
	"github.com/IhorBondartsov/OLX_Parser/olxParserMS/storage"

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
	AccessPublicKey []byte
}

type API struct {
	AccessTokenParser jwtLib.JWTParser
	Storage           storage.Storage
}

// Echo method for checking service
func (a *API) Echo(req EchoReq, res *EchoRes) error {
	fmt.Println("I called")
	res.Answer = fmt.Sprintf("Hello %s!!!", req.Name)
	return nil
}

// MakeOrder - make row to db
func (a *API) MakeOrder(req MakeOrderReq, res *MakeOrderRes) error {
	_, err := a.AccessTokenParser.Parse(req.Token)
	if err != nil {
		return err
	}

	order := entities.Order{
		DeliveryMethod: req.DeliveryMethod,
		ExpirationTime: req.DateTo,
		Frequency:      req.Frequency,
		PageLimit:      req.PageLimit,
		URL:            req.URL,
		UserID:         req.UserID,
	}


	return nil
}
