package modelutil

import "time"

type RequestLog struct {
	AcessTime time.Time
	Latency   time.Duration
	ClientIP  string
	Method    string
	Code      int
	Path      string
	UserAgent string
}