package auditlogger

import (
	"time"
)

type auditEvent struct {
	timestamp time.Time `json:"timestamp"`
	message   string    `json:"message"`
}