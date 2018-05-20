package webrpc

import (
	"math/rand"
	"net/http"
	"net/rpc"
	"strconv"
	"time"

	"github.com/IhorBondartsov/OLX_Parser/userms/cfg"
	"github.com/IhorBondartsov/OLX_Parser/userms/entities"

	"fmt"
	"github.com/IhorBondartsov/OLX_Parser/lib/jwtLib"
	"github.com/IhorBondartsov/OLX_Parser/userms/storage"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"github.com/sirupsen/logrus"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const tokenLength = 256

var log = logrus.New()


type ExampleSvc struct{}
type NameRes struct{ Name string }
// Method with named params and HTTP context.
func (*API) FullName3(_ struct{}, res *NameRes) error {
	fmt.Printf("FullName3(): Remote IP is %s\n", "2342354")
	return nil
}


func Start(cfg CfgAPI) {
	// Server export an object of type ExampleSvc.
	if err :=rpc.Register(NewAPI(cfg)); err != nil {
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
	ats, err := jwtLib.NewJWTSigner(cfg.AccessPrivateKey)
	if err != nil {
		log.Errorf("Cant create AccessTokenSigner. Err %v", err)
		return nil
	}
	return &API{
		AccessTokenParser: atp,
		AccessTokenSigner: ats,
		UserStor:          cfg.UserStor,
		RefreshStor:       cfg.RefreshStor,
	}
}

type CfgAPI struct {
	AccessPublicKey  []byte
	AccessPrivateKey []byte
	UserStor         storage.Storage
	RefreshStor      storage.RefreshToken
}

type API struct {
	AccessTokenParser jwtLib.JWTParser
	AccessTokenSigner jwtLib.JWTSigner
	UserStor          storage.Storage
	RefreshStor       storage.RefreshToken
	TTLAccessToken    time.Duration
}

// Login - check user in database if user is in databse then ganarate and return refrash token
func (a *API) Login(req LoginReq, resp *LoginResp) error {
	user, err := a.UserStor.GetUserByLogin(req.Login)
	if err != nil {
		log.Errorf("[Login][GetUserByLogin] Error %v", err)
		return err
	}

	token := randStringBytesRmndr(tokenLength)
	t := time.Duration(cfg.TTLRefreshToken) * time.Second
	tokenStruct := entities.Token{
		Token:          token,
		ExpirationTime: int(time.Now().Add(t).Unix()),
		UserID:         user.ID,
	}
	err = a.RefreshStor.SetToken(tokenStruct)
	if err != nil {
		log.Errorf("[Login][SetToken] Error %v", err)
		return err
	}
	resp.RefreshToken = token
	return err
}

// GetAcessToken - create access token for user using refrash token
func (a *API) GetAcessToken(req AcessTokenRequest, resp *AcessTokenResponse) error {
	refToken, err := a.RefreshStor.GetTokenByToken(req.RefreshToken)
	if err != nil {
		log.Errorf("[GetAcessToken][GetTokenByToken] Error with database %v", err)
		return err
	}
	claim := jwtLib.Claims{
		ID: strconv.Itoa(refToken.UserID),
	}
	resp.AcessToken, err = a.AccessTokenSigner.Sign(claim, a.TTLAccessToken)
	if err != nil {
		log.Errorf("[GetAcessToken][Sign] Cant sign. Error %v", err)
		return err
	}
	return nil
}

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
