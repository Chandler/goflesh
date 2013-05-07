package models

type Organization struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Default_timezone string `json:"default_timezone"`
}
