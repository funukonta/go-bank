package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Account struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		// ID:        rand.Intn(10000),
		Firstname: firstname,
		LastName:  lastname,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC(),
	}
}
