package models

type Organization struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Location         string `json:"location"`
	Default_timezone string `json:"default_timezone"`
	Description      string `json:"description"`
	TimeTrackedModel
}
