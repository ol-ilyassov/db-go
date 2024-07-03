package main

import "errors"

var (
	ErrNoRecord = errors.New("no record")
)

type Vegetable struct {
	Name  string
	Count int
	Price float64
}
