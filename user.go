package main

import (
	"time"
)

type User struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	HANDLE    string `json:"handle,omitempty"`
	CREATEDON string `json:"createdOn,omitempty"`
	UPDATEDON string `json:"-"`
	ARCHIVED  bool   `json:"-"`
}

var Matthew = &User{
	ID:        1,
	NAME:      "Matthew McGibbon",
	HANDLE:    "@matt",
	CREATEDON: time.Now().UTC().String(),
}
