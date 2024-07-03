package model

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Book struct {
	ID    int
	Title string
	Price float64
}

type Edit struct {
	Title *string
	Price *float64
}
