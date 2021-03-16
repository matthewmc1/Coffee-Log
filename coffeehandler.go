package main

import (
	"encoding/json"
	"fmt"
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

	for _, element := range CoffeeList {
		if element.ID == coffee.ID {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Record has not been created, as the id provided already exists in the database"))
			c.l.Printf("Duplicate created received was %#v", coffee)
			return
		}
	}

	c.l.Printf("Coffee details that are being created are %s", coffee.NAME)
	addCoffee(&coffee)

	rw.WriteHeader(http.StatusCreated)
	str := fmt.Sprintf("Coffee log has been created for id %d at %s", coffee.ID, time.Now().UTC().String())
	rw.Write([]byte(str))
}

func (c *Coffees) updateCoffee(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		c.l.Panic("Error message as body cannot be parsed")
		return
	}

	c.l.Printf("Request to update coffee log at %s with content %s and from host %s", time.Now().UTC().String(), body, r.Host)

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

	for idx, element := range CoffeeList {
		if element.ID == id {
			CoffeeList[idx] = &coffee
			rw.WriteHeader(http.StatusAccepted)

			c.l.Printf("Update request has been processed for id %d at %s", id, time.Now().UTC().String())
			c.l.Printf("Previous Coffee record is %#v", element)
			c.l.Printf("New Coffee record is %#v", coffee)

			str := fmt.Sprintf("Update request has been processed for id %d at %s", id, time.Now().UTC().String())
			rw.Write([]byte(str))
			return
		}
	}
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write([]byte("Record not found"))
}
