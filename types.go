package main

import (
	"math/rand"
	"time"
)

type CreateExpanseRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Expanse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Amount    float64   `json:"amount"`
	CreatedOn time.Time `json:"createdOn"`
}

func NewExpanse(firstName, lastname string) *Expanse {
	return &Expanse{
		FirstName: firstName,
		LastName:  lastname,
		Amount:    rand.Float64() * 1000,
		CreatedOn: time.Now().UTC(),
	}
}
