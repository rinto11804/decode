package main

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type CourseCreateBody struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
}
