package main

import (
	"fmt"
	//"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/IhorBondartsov/OLX_Parser/userms/cfg"
	"github.com/IhorBondartsov/OLX_Parser/userms/webrpc"
	"github.com/IhorBondartsov/OLX_Parser/userms/storage/userSQL"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"net/http"
)

var log = logrus.New()

func main() {
	//Make connection with db
	//user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%s?timeout=5s",
		cfg.Storage.Login,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.DBName)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("[MAIN] Cant create db connection %v", err)
	}
	userStor := userSQL.NewUserMyClientMySQL(db)
	tokenStor := userSQL.NewTokenClientMySQL(db)

	apiCfg := webrpc.CfgAPI{
		AccessPublicKey:  []byte(cfg.PublicKey),
		AccessPrivateKey: []byte(cfg.PrivateKey),
		UserStor:         userStor,
		RefreshStor:      tokenStor,
	}
	 webrpc.Start(apiCfg)


//	go http.ListenAndServe("127.0.0.1:8001", nil)

	log.Info("Listening on ", (cfg.Route+":"+cfg.Port))
	go http.ListenAndServe((cfg.Route+":"+cfg.Port), nil)


	clientHTTP := jsonrpc2.NewHTTPClient("http://127.0.0.1:8001/rpc")
	defer clientHTTP.Close()

	// Synchronous call using named params and HTTP with context.
	clientHTTP.Call("API.FullName3", nil, nil)
}
