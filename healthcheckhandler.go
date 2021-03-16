package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthChecks struct {
	l *log.Logger
}

func NewHealthCheck(l *log.Logger) *HealthChecks {
	return &HealthChecks{l}
}

var version = "0.0.1"

func (hc *HealthChecks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hc.l.Printf("Server process id: %d", os.Getpid())

	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	hd := Health()

	json.NewEncoder(w).Encode(hd)
}
