package models

type Oz struct {
	Id        int  `json:"id"`
	Confirmed bool `json:"confirmed"`
	TimeTrackedModel
}
