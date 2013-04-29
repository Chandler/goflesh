package models

import (
// "github.com/robfig/revel"
)

type Organization struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
