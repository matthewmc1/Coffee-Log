package main

import (
	"time"
)

type Coffee struct {
	ID        int      `json:"id"`
	NAME      string   `json:"name"`
	TASTE     []string `json:"tasteNotes,omitempty"`
	ISO       string   `json:"isoCountryCOde,omitempty"`
	VARIETY   string   `json:"variety,omitempty"`
	GRINDSIZE int      `json:"grindSize,omitempty"`
	IN        int      `json:"doseIn,omitempty"`
	OUT       int      `json:"doseOut,omitempty"`
	PICTURE   []byte   `json:"-"`
	LINK      string   `json:"purchaseLink,omitempty"`
	CREATEDON string   `json:"createdOn,omitempty"`
	CREATEDBY *User    `json:"createdBy,omitempty"`
	UPDATEDON string   `json:"-"`
	ARCHIVED  bool     `json:"-"`
}

func AllCoffee() []*Coffee {
	return CoffeeList
}

var CoffeeList = []*Coffee{
	{
		ID:   1,
		NAME: "Chelchele",
		TASTE: []string{
			"fruity", "nougat",
		},
		ISO:       "ETH",
		VARIETY:   "Heirloom",
		GRINDSIZE: 3,
		IN:        18,
		OUT:       45,
		CREATEDON: time.Now().UTC().String(),
		CREATEDBY: Matthew,
	},
	{
		ID:   2,
		NAME: "Chelchele",
		TASTE: []string{
			"fruity", "nougat",
		},
		ISO:       "ETH",
		VARIETY:   "Heirloom",
		GRINDSIZE: 3,
		IN:        16,
		OUT:       32,
		CREATEDON: time.Now().UTC().String(),
		CREATEDBY: Matthew,
	},
}
