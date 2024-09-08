package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AccNo     int    `json:"accountNo"`
	Balance   int    `json:"balance"`
}

func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		FirstName: FirstName,
		LastName:  LastName,
		AccNo:     rand.Intn(10000),
		ID:        rand.Intn(1000000),
	}
}
