package http

import (
	"log"
	"net/http"
	"time"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		start := time.Now()
		next.ServeHTTP(lrw, r)
		log.Printf("[%s] %s %s %v STATUS_CODE: %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start), lrw.StatusCode)
	})
}
