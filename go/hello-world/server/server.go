package server

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

type Server struct {
	http.Handler
}

func New() *Server  {

	r := mux.NewRouter()

	r.HandleFunc("/", homehandler)

	return &Server{
		Handler: r,
	}
}

func(s *Server) Start(){
	err := http.ListenAndServe(":8080", s.Handler)
	if err != nil {
		log.Fatalf("Failure to start Server %v", err.Error())
	}
}

func homehandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello World"))
}