package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"

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
func (s *server) Start(){
	r := mux.NewRouter()
	r.HandleFunc("/info", Info).Methods("GET")

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




