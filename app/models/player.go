package models

type Player struct {
	Id      int `json:"id"`
	User_id int `json:"user_id"`
	Game_id int `json:"game_id"`
	TimeTrackedModel
}
