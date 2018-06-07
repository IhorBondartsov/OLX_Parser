package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IhorBondartsov/OLX_Parser/website/clients"
	"github.com/IhorBondartsov/OLX_Parser/website/entity"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"io/ioutil"
	"github.com/IhorBondartsov/OLX_Parser/userms/cfg"
)

var log = logrus.New()

// server - create new server
type server struct {
	Route string
	Port  int
}

// CfgServer - config struct which helps to create new server
type CfgServer struct {
	Route string
	Port  int
}

// NewServer - create new server
func NewServer(cfg CfgServer) *server {
	return &server{
		Route: cfg.Route,
		Port:  cfg.Port,
	}
}

// Start - start http server
func (s *server) Start() {
	r := mux.NewRouter()
	r.HandleFunc("/info", Info).Methods("GET")
	r.HandleFunc("/userms/rpc", UserMS).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./view/")))

	addr := fmt.Sprintf("%v:%d", s.Route, s.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("Server Start on %v", addr)
	log.Fatal(srv.ListenAndServe())
}

func Info(w http.ResponseWriter, req *http.Request) {
	log.Info("Hello world!")
	fmt.Fprintln(w, "Hello world!")
}

type Request struct {
	Method string
	Data   json.RawMessage
}

func UserMS(w http.ResponseWriter, req *http.Request) {
	log.Info("[Server][UserMS]")
	cli := clients.NewUserRPCClient(cfg.UserMS.Host, cfg.UserMS.Port, cfg.UserMS.Prefix)

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Errorf("[Server][UserMS] Cant read body Err %v", err)
	}

	var reqJSON Request
	var bodyResp []byte
	err = json.Unmarshal(b, &reqJSON)
	if err != nil {
		log.Errorf("[Server][UserMS] Cant unmarshal Err %v", err)
	}

	switch reqJSON.Method {
	case "Echo":
		var er entity.EchoReq
		err = json.Unmarshal(reqJSON.Data, &er)
		if err != nil {
			log.Errorf("[Server][UserMS] Cant unmarshal Err %v", err)
		}
		bodyResp, err = cli.Echo(er.Name)
		if err != nil {
			log.Errorf("[Server][UserMS] Err %v", err)
		}
	case "GetAcessToken":
		var er entity.AcessTokenRequest
		err = json.Unmarshal(reqJSON.Data, &er)
		if err != nil {
			log.Errorf("[Server][UserMS] Cant unmarshal Err %v", err)
		}
		bodyResp, err = cli.GetAcessToken(er.RefreshToken)
		if err != nil {
			log.Errorf("[Server][UserMS] Err %v", err)
		}
	case "Login":
		var er entity.LoginReq
		err = json.Unmarshal(reqJSON.Data, &er)
		if err != nil {
			log.Errorf("[Server][UserMS] Cant unmarshal Err %v", err)
		}
		bodyResp, err = cli.Login(er.Login, er.Password)
		if err != nil {
			log.Errorf("[Server][UserMS] Err %v", err)
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bodyResp)

}
