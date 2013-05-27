package models

import (
	"github.com/coopernurse/gorp"
	"time"
)

type Player struct {
	Id      int        `json:"id"`
	User_id int        `json:"user_id"`
	Game_id int        `json:"game_id"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

func (model *Player) PreInsert(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Created = &now
	model.Updated = model.Created
	return nil
}

func (model *Player) PreUpdate(s gorp.SqlExecutor) error {
	now := time.Now().UTC()
	model.Updated = &now
	return nil
}
