package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/IhorBondartsov/OLX_Parser/userms/cfg"
	"github.com/IhorBondartsov/OLX_Parser/userms/webrpc"
)

var log = logrus.New()

func main() {
	//Make connection with db
	//user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/dbname?tls=skip-verify&autocommit=true",
		cfg.Storage.Login,
		cfg.Storage.Password,
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.DBName)
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("[MAIN] Cant create db connection %v", err)
	}
	fmt.Println(db)

	//userStor := userSQL.NewUserMyClientMySQL(db)
	//tokenStor := userSQL.NewTokenClientMySQL(db)

	apiCfg := webrpc.CfgAPI{
		AccessPublicKey:  []byte(cfg.PublicKey),
		AccessPrivateKey: []byte(cfg.PrivateKey),
		UserStor:         userStor,
		RefreshStor:      tokenStor,
	}
	df
}
