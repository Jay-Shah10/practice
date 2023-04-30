package auditlogger

import (
	"fmt"
	"encoding/json"
	"context"
	
)

type Auditlogservice struct {
	eventChannel chan AuditEvent
}

func New() *Auditlogservice {
	ch := make(chan AuditEvent, 64)

	return &Auditlogservice{
		eventChannel: ch,
	}
}

func (a *Auditlogservice) Run(ctx context.Context){
	for {
		select {
		case <- ctx.Done():
			return
		case auditLogEvent := <- a.eventChannel:
			a.processEvent(auditLogEvent)
		}
	}
}

func (a *Auditlogservice) processEvent(event AuditEvent){
	j, e := json.MarshalIndent(event, "", " ")
	if e != nil {
		fmt.Println("Cannot print json")
	}
	fmt.Println(j)

}

// func (a *Auditlogservice) Publish(event AuditEvent){
// 	switch {
// 	case a.eventChannel <- event:
// 		fmt.Println("Got event")
// 	default: 
// 		fmt.Println("Cannot process event	")
// 	}
// }
