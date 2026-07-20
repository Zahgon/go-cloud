package main

import (
	"time"
)

type Order struct {
	ID               string
	Email            string
	InImage          string
	OutImage         string
	CreateTime       time.Time
	FinishTime       time.Time
	Note             string
	DocstoreRevision any
}

type OrderRequest struct {
	ID         string
	Email      string
	InImage    string
	CreateTime time.Time
}
