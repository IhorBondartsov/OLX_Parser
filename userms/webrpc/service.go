package webrpc

import (
	"github.com/IhorBondartsov/OLX_Parser/lib/jwtLib"
	"github.com/sirupsen/logrus"
)
var log = logrus.New()


func NewAPI(cfg CfgAPI)*API{
	rtp, err := jwtLib.NewJWTParser(cfg.RefreshPublicKey)
	if err != nil{
		log.Errorf("Cant create RefreshTokenParser. Err %v", err)
		return nil
	}
	rts, err := jwtLib.NewJWTSigner(cfg.RefreshPrivateKey)
	if err != nil{
		log.Errorf("Cant create RefreshTokenSigner. Err %v", err)
		return nil
	}
	atp, err := jwtLib.NewJWTParser(cfg.AccessPrivateKey)
	if err != nil{
		log.Errorf("Cant create AccessTokenParser. Err %v", err)
		return nil
	}
	ats, err := jwtLib.NewJWTSigner(cfg.AccessPublicKey)
	if err != nil{
		log.Errorf("Cant create AccessTokenSigner. Err %v", err)
		return nil
	}
	return &API{
		RefreshTokenSigner:rts,
		RefreshTokenParser:rtp,
		AccessTokenParser:atp,
		AccessTokenSigner:ats,
	}
}

type CfgAPI struct{
	RefreshPublicKey []byte
	RefreshPrivateKey []byte
	AccessPublicKey []byte
	AccessPrivateKey []byte
}

type API struct{
	RefreshTokenParser jwtLib.JWTParser
	RefreshTokenSigner jwtLib.JWTSigner
	AccessTokenParser jwtLib.JWTParser
	AccessTokenSigner jwtLib.JWTSigner
}

func (a *API) Login(req LoginReq, resp *LoginResp){

}