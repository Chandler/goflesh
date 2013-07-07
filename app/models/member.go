package models

type Member struct {
	Id      int `json:"id"`
	User_id int `json:"user_id"`
	Organization_id int `json:"organization_id"`
	TimeTrackedModel
}
