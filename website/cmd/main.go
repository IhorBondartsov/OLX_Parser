package main

import (
	"github.com/IhorBondartsov/OLX_Parser/website/cfg"
	"github.com/IhorBondartsov/OLX_Parser/website/server"
)

func main() {
	cfgServer := server.CfgServer{
		Port:  cfg.Port,
		Route: cfg.Route,
	}
	serv := server.NewServer(cfgServer)
	serv.Start()
}
