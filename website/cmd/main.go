package main

import (
	"github.com/IhorBondartsov/OLX_Parser/website/cfg"
	"github.com/IhorBondartsov/OLX_Parser/website/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfgServer := server.CfgServer{
		Port:  cfg.Port,
		Route: cfg.Route,
	}
	serv := server.NewServer(cfgServer)
	log.Info("Website server is starting...")
	serv.Start()
}
