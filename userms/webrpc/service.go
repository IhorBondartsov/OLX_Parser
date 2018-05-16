package webrpc

import (
	"math/rand"

	"github.com/IhorBondartsov/OLX_Parser/userMS/cfg"
	"github.com/IhorBondartsov/OLX_Parser/userms/entities"

	"github.com/IhorBondartsov/OLX_Parser/lib/jwtLib"
	"github.com/IhorBondartsov/OLX_Parser/userms/storage"
	"github.com/sirupsen/logrus"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const tokenLength = 256

var log = logrus.New()

func NewAPI(cfg CfgAPI) *API {
	atp, err := jwtLib.NewJWTParser(cfg.AccessPrivateKey)
	if err != nil {
		log.Errorf("Cant create AccessTokenParser. Err %v", err)
		return nil
	}
	ats, err := jwtLib.NewJWTSigner(cfg.AccessPublicKey)
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
}

func (a *API) Login(req LoginReq, resp *LoginResp) error {
	user, err := a.UserStor.GetUserByLogin(req.Login)
	if err != nil {
		log.Errorf("[Login][GetUserByLogin] Error %v", err)
		return err
	}

	token := randStringBytesRmndr(tokenLength)
	tokenStruct := entities.Token{
		Token:  token,
		TTL:    cfg.TTLRefradhToken,
		UserID: user.ID,
	}
	err = a.RefreshStor.SetToken(tokenStruct)
	if err != nil {
		log.Errorf("[Login][SetToken] Error %v", err)
		return err
	}
	resp.RefreshToken = token
	return err
}

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
