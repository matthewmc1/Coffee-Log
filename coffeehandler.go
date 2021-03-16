package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Coffees struct {
	l *log.Logger
}

func NewCoffee(l *log.Logger) *Coffees {
	return &Coffees{l}
}

func (c *Coffees) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c.l.Printf("Request to service at %s", time.Now().UTC().String())

	lc := AllCoffee()
	c.l.Printf("size of list that has been created is %d", len(lc))

	json.NewEncoder(rw).Encode(lc)
}
