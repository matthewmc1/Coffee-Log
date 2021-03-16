package main

import "net/http"

type HealthCheck struct {
	STATUS  int    `json:"status"`
	HEALTH  string `json:"health"`
	VERSION string `json:"version"`
}

func Health() *HealthCheck {
	return &HealthCheck{
		STATUS:  http.StatusOK,
		HEALTH:  "UP",
		VERSION: version,
	}
}
