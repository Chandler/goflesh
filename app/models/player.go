package models

import (
	"errors"
	"github.com/robfig/revel"
)

type Player struct {
	Id      int `json:"id"`
	User_id int `json:"user_id"`
	Game_id int `json:"game_id"`
	TimeTrackedModel
}

func (p *Player) HumanCode() *HumanCode {
	human, err := Dbm.Get(HumanCode{}, p.Id)
	if err != nil {
		revel.ERROR.Print("Could not get human code", err)
	}
	return human.(*HumanCode)
}

func PlayerFromId(id int) (*Player, error) {
	player, err := Dbm.Get(Player{}, id)
	return player.(*Player), err
}

func (p *Player) isZombie() bool {
	query := `
		SELECT COUNT(1)
		FROM player
		LEFT OUTER JOIN "oz"
			ON player.id = oz.id
		LEFT OUTER JOIN "tag"
			ON player.id = taggee_id
		WHERE player.id = $1
			AND (oz.id IS NULL OR oz.confirmed = FALSE)
			AND taggee_id IS NULL
	`
	numFound, err := Dbm.SelectInt(query, p.Id)
	if err != nil {
		panic(err)
	}
	isZombie := numFound == 0
	return isZombie
}

func (p *Player) isHuman() bool {
	return !p.isZombie()
}

func (p *Player) CanTag() (bool, error) {
	if !p.isZombie() {
		return false, errors.New("player is not a zombie!")
	}
	return true, nil
}

func (p *Player) CanBeTagged() (bool, error) {
	if !p.isHuman() {
		return false, errors.New("player is not a human!")
	}
	return true, nil
}
