package auditlogger

import (
	"time"
)

type AuditEvent struct {
	timestamp time.Time `json:"timestamp"`
	message   string    `json:"message"`
}