package main

import (
	"encoding/json"
	"io/ioutil"
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
	switch r.Method {
	case http.MethodGet:
		c.getCoffees(rw, r)
		return
	case http.MethodPost:
		c.createCoffee(rw, r)
		return
	case http.MethodPut:
		c.updateCoffee(rw, r)
		return
	default:
		c.l.Printf("Method used is not supported")
		return
	}

}

func (c *Coffees) getCoffees(rw http.ResponseWriter, r *http.Request) {
	c.l.Printf("Request to get coffee log at %s from host %s", time.Now().UTC().String(), r.Host)

	lc := AllCoffee()

	c.l.Printf("size of list that has been created is %d", len(lc))

	json.NewEncoder(rw).Encode(lc)
}

func (c *Coffees) createCoffee(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		c.l.Panic("Error message as body cannot be parsed")
		return
	}

	c.l.Printf("Request to create coffee log at %s with content %s and from host %s", time.Now().UTC().String(), body, r.Host)

	var coffee Coffee
	err = json.Unmarshal([]byte(body), &coffee)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Record has not been created, please correct data and try again or raise a support ticket"))
		c.l.Printf("Request received was %#v", coffee)
		return
	}

	c.l.Printf("Coffee details that are being created are %s", coffee.NAME)
	addCoffee(&coffee)

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("Record created"))
}

func (c *Coffees) updateCoffee(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		c.l.Panic("Error message as body cannot be parsed")
		return
	}

	c.l.Printf("Request to create coffee log at %s with content %s and from host %s", time.Now().UTC().String(), body, r.Host)

	var coffee Coffee
	err = json.Unmarshal([]byte(body), &coffee)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Record has not been created, please correct data and try again or raise a support ticket"))
		c.l.Printf("Request received was %#v", coffee)
		return
	}

	c.l.Printf("Coffee details that are being created are %s", coffee.NAME)

	id := coffee.ID

	for ix, element := range CoffeeList {
		if element.ID == id {
			CoffeeList[ix] = &coffee
			rw.WriteHeader(http.StatusAccepted)
			rw.Write([]byte("Record updated"))
			return
		}
	}
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write([]byte("Record not found"))
}
