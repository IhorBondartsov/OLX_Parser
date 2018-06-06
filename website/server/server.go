package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

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

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./view/")))

	addr := fmt.Sprintf("%v:%d", s.Route, s.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func Info(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
