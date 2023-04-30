package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Jay-Shah10/practice/go/audit-log/services/auditlogger"
)

type Server struct{
	http.Handler
	ServerOption
}

type ServerOption struct {
	AuditLogger: auditlogger.Auditlogservice

}

func New() *Server {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/", homeHandler)


	return &Server{
		Handler: r,
	}
}

func (s *Server) Start() {
	go func() {
		err := http.ListenAndServe(":8080", s.Handler)
		if err != nil {
			log.Fatalf("Failure to start server %v", err.Error())
		}
	}()
	
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request)  {

	event := s.auditEvent{
		time: time.Now(),
		message: "From Home Handler"
	}
	s.AuditLogger.Publish(event)
	
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Atlast Home"))
}
