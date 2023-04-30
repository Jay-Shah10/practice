package main

import (
	"context"
	"github.com/Jay-Shah10/practice/go/audit-log/server"
	"github.com/Jay-Shah10/practice/go/audit-log/services/auditlogger"
)

func main()  {
	ctx := context.Background()
	
	server := server.New()

	auditlog := Auditlogservice.New()
	go func(){
		auditlog.Run()
	}
	
	options := &server.ServerOption{
		AuditLogger: auditlogger.Auditlogservice
	}
	
	server.Start()
}